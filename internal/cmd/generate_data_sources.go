// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/input"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/logging"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templating"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/validate"
)

type GenerateDataSourcesCommand struct {
	UI                cli.Ui
	flagIRInputPath   string
	flagOutputPath    string
	flagPackageName   string
	flagTemplatesPath string
}

func (cmd *GenerateDataSourcesCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("generate data-sources", flag.ExitOnError)
	fs.StringVar(&cmd.flagIRInputPath, "input", "./ir.json", "path to intermediate representation (JSON)")
	fs.StringVar(&cmd.flagOutputPath, "output", "./output", "directory path to output generated code files")
	fs.StringVar(&cmd.flagPackageName, "package", "", "name of Go package for generated code files")
	fs.StringVar(&cmd.flagTemplatesPath, "templates", "", "directory path for templates (*.gotmpl files) to be processed after schema generation")

	return fs
}

func (cmd *GenerateDataSourcesCommand) Help() string {
	strBuilder := &strings.Builder{}

	longestName := 0
	longestUsage := 0
	cmd.Flags().VisitAll(func(f *flag.Flag) {
		if len(f.Name) > longestName {
			longestName = len(f.Name)
		}
		if len(f.Usage) > longestUsage {
			longestUsage = len(f.Usage)
		}
	})

	strBuilder.WriteString("\nUsage: tfplugingen-framework generate data-sources [<args>]\n\n")
	cmd.Flags().VisitAll(func(f *flag.Flag) {
		if f.DefValue != "" {
			strBuilder.WriteString(fmt.Sprintf("    --%s <ARG> %s%s%s  (default: %q)\n",
				f.Name,
				strings.Repeat(" ", longestName-len(f.Name)+2),
				f.Usage,
				strings.Repeat(" ", longestUsage-len(f.Usage)+2),
				f.DefValue,
			))
		} else {
			strBuilder.WriteString(fmt.Sprintf("    --%s <ARG> %s%s%s\n",
				f.Name,
				strings.Repeat(" ", longestName-len(f.Name)+2),
				f.Usage,
				strings.Repeat(" ", longestUsage-len(f.Usage)+2),
			))
		}
	})
	strBuilder.WriteString("\n")

	return strBuilder.String()
}

func (a *GenerateDataSourcesCommand) Synopsis() string {
	return "Generate code for data sources from an Intermediate Representation (IR) JSON file."
}

func (cmd *GenerateDataSourcesCommand) Run(args []string) int {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))

	fs := cmd.Flags()
	err := fs.Parse(args)
	if err != nil {
		cmd.UI.Error(fmt.Sprintf("error parsing command flags: %s", err))
		return 1
	}

	err = cmd.runInternal(ctx, logger)
	if err != nil {
		cmd.UI.Error(fmt.Sprintf("Error executing command: %s\n", err))
		return 1
	}

	return 0
}

func (cmd *GenerateDataSourcesCommand) runInternal(ctx context.Context, logger *slog.Logger) error {
	// read input file
	src, err := input.Read(cmd.flagIRInputPath)
	if err != nil {
		return fmt.Errorf("error reading IR JSON: %w", err)
	}

	// validate JSON
	err = validate.JSON(src)
	if err != nil {
		return fmt.Errorf("error validating IR JSON: %w", err)
	}

	// parse and validate IR against specification
	spec, err := spec.Parse(ctx, src)
	if err != nil {
		return fmt.Errorf("error parsing IR JSON: %w", err)
	}

	templateData, err := generateDataSourceCode(ctx, spec, cmd.flagOutputPath, cmd.flagPackageName, "DataSource", logger)
	if err != nil {
		return fmt.Errorf("error generating data source code: %w", err)
	}

	if cmd.flagTemplatesPath != "" {
		templator := templating.NewTemplator(logger, os.DirFS(cmd.flagTemplatesPath))

		dOutput, err := templator.ProcessDataSources(templateData)
		if err != nil {
			return fmt.Errorf("error processing data source templates: %w", err)
		}

		for fileName, fileBytes := range dOutput {
			outputFile := path.Join(cmd.flagOutputPath, fileName)
			err := output.WriteBytes(outputFile, fileBytes, true)
			if err != nil {
				return fmt.Errorf("error writing processed template to output dir: %w", err)
			}
		}
	}

	return nil
}

func generateDataSourceCode(ctx context.Context, spec spec.Specification, outputPath, packageName, generatorType string, logger *slog.Logger) (map[string]templating.DataSourceTemplateData, error) {
	ctxWithPath := logging.SetPathInContext(ctx, "data_source")

	// convert IR to framework schema
	s, err := datasource.NewSchemas(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting IR to Plugin Framework schema: %w", err)
	}

	// convert framework schema to []byte
	g := schema.NewGeneratorSchemas(s)
	schemas, err := g.Schemas(packageName, generatorType)
	if err != nil {
		return nil, fmt.Errorf("error converting Plugin Framework schema to Go code: %w", err)
	}

	// generate model code
	models, err := g.Models()
	if err != nil {
		return nil, err
	}

	// generate custom type and value types code
	customTypeValue, err := g.CustomTypeValue()
	if err != nil {
		return nil, err
	}

	// generate "expand" and "flatten" code
	toFromFunctions, err := g.ToFromFunctions(ctxWithPath, logger)
	if err != nil {
		return nil, err
	}

	// write code
	templateData, err := output.WriteDataSources(schemas, models, customTypeValue, toFromFunctions, outputPath, packageName)
	if err != nil {
		return nil, fmt.Errorf("error writing Go code to output: %w", err)
	}

	return templateData, nil
}
