// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/mitchellh/cli"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
)

type GenerateCommand struct {
	UI cli.Ui
}

func (cmd *GenerateCommand) Help() string {
	helpText := `
	Usage: terraform-plugin-codegen-framework generate <subcommand> [<args>]
	
	  This command has subcommands for Terraform Plugin Framework code generation.
	
	`
	return strings.TrimSpace(helpText)
}

func (a *GenerateCommand) Synopsis() string {
	return "Terraform Plugin Framework code generation commands"
}

func (cmd *GenerateCommand) Run(args []string) int {
	return cli.RunResultHelp
}

func generateDataSourceCode(spec spec.Specification, outputPath string) error {
	// convert IR to framework schema
	c := datasource_convert.NewConverter(spec)
	schema, err := c.ToGeneratorDataSourceSchema()
	if err != nil {
		return fmt.Errorf("error converting IR to Plugin Framework schema: %w", err)
	}

	// convert framework schema to []byte
	g := datasource_generate.NewGeneratorDataSourceSchemas(schema)
	schemaBytes, err := g.ToBytes()
	if err != nil {
		return fmt.Errorf("error converting Plugin Framework schema to Go code: %w", err)
	}

	// format schema code
	formattedDataSourcesSchema, err := format.Format(schemaBytes)
	if err != nil {
		return fmt.Errorf("error formatting Go code: %w", err)
	}

	// write code
	err = output.WriteDataSources(formattedDataSourcesSchema, outputPath)
	if err != nil {
		return fmt.Errorf("error writing Go code to output: %w", err)
	}

	return nil
}

func generateResourceCode(spec spec.Specification, outputPath string) error {
	// convert IR to framework schema
	resourceSchemaConverter := resource_convert.NewConverter(spec)
	resourceSchemas, err := resourceSchemaConverter.ToGeneratorResourceSchema()
	if err != nil {
		return fmt.Errorf("error converting IR to Plugin Framework schema: %w", err)
	}

	// convert framework schema to []byte
	resourceSchemaGenerator := resource_generate.NewGeneratorResourceSchemas(resourceSchemas)
	resourceSchemaBytes, err := resourceSchemaGenerator.ToBytes()
	if err != nil {
		return fmt.Errorf("error converting Plugin Framework schema to Go code: %w", err)
	}

	// format schema code
	formattedResourcesSchema, err := format.Format(resourceSchemaBytes)
	if err != nil {
		return fmt.Errorf("error formatting Go code: %w", err)
	}

	// write code
	err = output.WriteResources(formattedResourcesSchema, outputPath)
	if err != nil {
		return fmt.Errorf("error writing Go code to output: %w", err)
	}

	return nil
}

func generateProviderCode(spec spec.Specification, outputPath string) error {
	// convert IR to framework schema
	providerSchemaConverter := provider_convert.NewConverter(spec)
	providerSchemas, err := providerSchemaConverter.ToGeneratorProviderSchema()
	if err != nil {
		return fmt.Errorf("error converting IR to Plugin Framework schema: %w", err)
	}

	// convert framework schema to []byte
	providerSchemaGenerator := provider_generate.NewGeneratorProviderSchemas(providerSchemas)
	providerSchemaBytes, err := providerSchemaGenerator.ToBytes()
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
