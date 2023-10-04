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

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/input"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/validate"
)

type GenerateProviderCommand struct {
	UI              cli.Ui
	flagIRInputPath string
	flagOutputPath  string
	flagPackageName string
}

func (cmd *GenerateProviderCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("generate provider", flag.ExitOnError)
	fs.StringVar(&cmd.flagIRInputPath, "input", "./ir.json", "path to intermediate representation (JSON)")
	fs.StringVar(&cmd.flagOutputPath, "output", "./output", "directory path to output generated code files")
	fs.StringVar(&cmd.flagPackageName, "package", "", "name of Go package for generated code files")

	return fs
}

func (cmd *GenerateProviderCommand) Help() string {
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

	strBuilder.WriteString("\nUsage: tfplugingen-framework generate provider [<args>]\n\n")
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

func (a *GenerateProviderCommand) Synopsis() string {
	return "Generate code for the provider from an Intermediate Representation (IR) JSON file."
}

func (cmd *GenerateProviderCommand) Run(args []string) int {
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

func (cmd *GenerateProviderCommand) runInternal(ctx context.Context) error {
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

	err = generateProviderCode(spec, cmd.flagOutputPath, cmd.flagPackageName, "Provider")
	if err != nil {
		return fmt.Errorf("error generating provider code: %w", err)
	}

	return nil
}

func generateProviderCode(spec spec.Specification, outputPath, packageName, generatorType string) error {
	// convert IR to framework schema
	c := provider_convert.NewConverter(spec)
	s, err := c.ToGeneratorProviderSchema()
	if err != nil {
		return fmt.Errorf("error converting IR to Plugin Framework schema: %w", err)
	}

	// convert framework schema to []byte
	g := schema.NewGeneratorSchemas(s)
	schemaBytes, err := g.SchemasBytes(packageName, generatorType)
	if err != nil {
		return fmt.Errorf("error converting Plugin Framework schema to Go code: %w", err)
	}

	// generate model code
	modelsBytes, err := g.ModelsBytes()
	if err != nil {
		log.Fatal(err)
	}

	// generate custom type and value types code
	customTypeValueBytes, err := g.CustomTypeValueBytes()
	if err != nil {
		log.Fatal(err)
	}

	// generate "expand" and "flatten" code
	toFromFunctions, err := g.ToFromFunctions()
	if err != nil {
		log.Fatal(err)
	}

	// format schema code
	formattedProvidersSchema, err := format.Format(schemaBytes)
	if err != nil {
		return fmt.Errorf("error formatting Go code: %w", err)
	}

	// format model code
	formattedProvidersModels, err := format.Format(modelsBytes)
	if err != nil {
		log.Fatal(err)
	}

	// format custom type and value types code
	formattedCustomTypeValue, err := format.Format(customTypeValueBytes)
	if err != nil {
		log.Fatal(err)
	}

	// format "expand" and "flatten" code
	formattedToFromFunctions, err := format.Format(toFromFunctions)
	if err != nil {
		log.Fatal(err)
	}

	// write code
	err = output.WriteProviders(formattedProvidersSchema, formattedProvidersModels, formattedCustomTypeValue, formattedToFromFunctions, outputPath, packageName)
	if err != nil {
		return fmt.Errorf("error writing Go code to output: %w", err)
	}

	return nil
}
