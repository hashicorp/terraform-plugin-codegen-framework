// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/input"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/validate"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/mitchellh/cli"
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
	fs.StringVar(&cmd.flagPackageName, "package", "provider", "name of Go package for generated code files")

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

	strBuilder.WriteString("\nUsage: terraform-plugin-codegen-framework generate provider [<args>]\n\n")
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

	err = generateProviderCode(spec, cmd.flagOutputPath, cmd.flagPackageName)
	if err != nil {
		return fmt.Errorf("error generating provider code: %w", err)
	}

	return nil
}

func generateProviderCode(spec spec.Specification, outputPath, packageName string) error {
	// convert IR to framework schema
	providerSchemaConverter := provider_convert.NewConverter(spec)
	providerSchemas, err := providerSchemaConverter.ToGeneratorProviderSchema()
	if err != nil {
		return fmt.Errorf("error converting IR to Plugin Framework schema: %w", err)
	}

	// convert framework schema to []byte
	providerSchemaGenerator := provider_generate.NewGeneratorProviderSchemas(providerSchemas)
	providerSchemaBytes, err := providerSchemaGenerator.ToBytes(packageName)
	if err != nil {
		return fmt.Errorf("error converting Plugin Framework schema to Go code: %w", err)
	}

	// format schema code
	formattedProvidersSchema, err := format.Format(providerSchemaBytes)
	if err != nil {
		return fmt.Errorf("error formatting Go code: %w", err)
	}

	// write code
	err = output.WriteProviders(formattedProvidersSchema, outputPath)
	if err != nil {
		return fmt.Errorf("error writing Go code to output: %w", err)
	}

	return nil
}
