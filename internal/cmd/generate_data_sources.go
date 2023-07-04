// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/mitchellh/cli"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/input"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/validate"
)

type GenerateDataSourcesCommand struct {
	UI              cli.Ui
	flagIRInputPath string
	flagOutputPath  string
	flagPackageName string
}

func (cmd *GenerateDataSourcesCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("generate data-sources", flag.ExitOnError)
	fs.StringVar(&cmd.flagIRInputPath, "input", "./ir.json", "path to intermediate representation (JSON)")
	fs.StringVar(&cmd.flagOutputPath, "output", "./output", "directory path to output generated code files")
	fs.StringVar(&cmd.flagPackageName, "package", "provider", "name of Go package for generated code files")

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

	strBuilder.WriteString("\nUsage: terraform-plugin-codegen-framework generate data-sources [<args>]\n\n")
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

	fs := cmd.Flags()
	err := fs.Parse(args)
	if err != nil {
		cmd.UI.Error(fmt.Sprintf("error parsing command flags: %s", err))
		return 1
	}

	err = cmd.runInternal(ctx)
	if err != nil {
		cmd.UI.Error(fmt.Sprintf("Error executing command: %s\n", err))
		return 1
	}

	return 0
}

func (cmd *GenerateDataSourcesCommand) runInternal(ctx context.Context) error {
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

	err = generateDataSourceCode(spec, cmd.flagOutputPath, cmd.flagPackageName)
	if err != nil {
		return fmt.Errorf("error generating data source code: %w", err)
	}

	return nil
}

func generateDataSourceCode(spec spec.Specification, outputPath, packageName string) error {
	// convert IR to framework schema
	c := datasource_convert.NewConverter(spec)
	schema, err := c.ToGeneratorDataSourceSchema()
	if err != nil {
		return fmt.Errorf("error converting IR to Plugin Framework schema: %w", err)
	}

	// convert framework schema to []byte
	g := datasource_generate.NewGeneratorDataSourceSchemas(schema)
	schemaBytes, err := g.SchemasBytes(packageName)
	if err != nil {
		return fmt.Errorf("error converting Plugin Framework schema to Go code: %w", err)
	}

	// generate model code
	modelsBytes, err := g.ModelsBytes()
	if err != nil {
		log.Fatal(err)
	}

	// generate model object helpers code
	modelsObjectHelpersBytes, err := g.ModelsObjectHelpersBytes()
	if err != nil {
		log.Fatal(err)
	}

	// format schema code
	formattedDataSourcesSchema, err := format.Format(schemaBytes)
	if err != nil {
		return fmt.Errorf("error formatting Go code: %w", err)
	}

	// format model code
	formattedDataSourcesModels, err := format.Format(modelsBytes)
	if err != nil {
		log.Fatal(err)
	}

	// format model object helpers code
	formattedDataSourcesModelObjectHelpers, err := format.Format(modelsObjectHelpersBytes)
	if err != nil {
		log.Fatal(err)
	}

	// write code
	err = output.WriteDataSources(formattedDataSourcesSchema, formattedDataSourcesModels, formattedDataSourcesModelObjectHelpers, outputPath)
	if err != nil {
		return fmt.Errorf("error writing Go code to output: %w", err)
	}

	return nil
}
