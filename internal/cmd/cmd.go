// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/mitchellh/cli"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/config"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/input"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/validate"
)

type AllCommand struct {
	Ui cli.Ui
}

func (a AllCommand) Help() string {
	return "Both -input and -output can be specified. " +
		"-output defaults to output. "
}

func (a AllCommand) Synopsis() string {
	return "Generates schema."
}

func (a AllCommand) Run(args []string) int {
	ctx := context.Background()
	// parse flags
	conf, err := config.New(args)
	if err != nil {
		log.Fatal(err)
	}

	// read input file
	src, err := input.Read(conf.Input)
	if err != nil {
		log.Fatal(err)
	}

	// validate JSON
	err = validate.JSON(src)
	if err != nil {
		log.Fatal(err)
	}

	spec, err := spec.Parse(ctx, src)
	if err != nil {
		log.Fatal(err)
	}

	// ********************
	// DataSources
	// ********************
	// convert IR to framework schema
	c := datasource_convert.NewConverter(spec)
	schema, err := c.ToGeneratorDataSourceSchema()
	if err != nil {
		log.Fatal(err)
	}

	// convert framework schema to []byte
	g := datasource_generate.NewGeneratorDataSourceSchemas(schema)
	schemaBytes, err := g.ToBytes()
	if err != nil {
		log.Fatal(err)
	}

	// format schema code
	formattedDataSourcesSchema, err := format.Format(schemaBytes)
	if err != nil {
		log.Fatal(err)
	}

	// write code
	err = output.WriteDataSources(formattedDataSourcesSchema, conf.Output)
	if err != nil {
		log.Fatal(err)
	}

	// ********************
	// Resources
	// ********************
	// convert IR to framework schema
	resourceSchemaConverter := resource_convert.NewConverter(spec)
	resourceSchemas, err := resourceSchemaConverter.ToGeneratorResourceSchema()
	if err != nil {
		log.Fatal(err)
	}

	// convert framework schema to []byte
	resourceSchemaGenerator := resource_generate.NewGeneratorResourceSchemas(resourceSchemas)
	resourceSchemaBytes, err := resourceSchemaGenerator.ToBytes()
	if err != nil {
		log.Fatal(err)
	}

	// format schema code
	formattedResourcesSchema, err := format.Format(resourceSchemaBytes)
	if err != nil {
		log.Fatal(err)
	}

	// write code
	err = output.WriteResources(formattedResourcesSchema, conf.Output)
	if err != nil {
		log.Fatal(err)
	}

	// ********************
	// Provider
	// ********************
	// convert IR to framework schema
	providerSchemaConverter := provider_convert.NewConverter(spec)
	providerSchemas, err := providerSchemaConverter.ToGeneratorProviderSchema()
	if err != nil {
		log.Fatal(err)
	}

	// convert framework schema to []byte
	providerSchemaGenerator := provider_generate.NewGeneratorProviderSchemas(providerSchemas)
	providerSchemaBytes, err := providerSchemaGenerator.ToBytes()
	if err != nil {
		log.Fatal(err)
	}

	// format schema code
	formattedProvidersSchema, err := format.Format(providerSchemaBytes)
	if err != nil {
		log.Fatal(err)
	}

	// write code
	err = output.WriteProviders(formattedProvidersSchema, conf.Output)
	if err != nil {
		log.Fatal(err)
	}

	return 0
}
