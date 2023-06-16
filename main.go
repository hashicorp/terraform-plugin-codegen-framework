// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"encoding/json"
	"log"
	"os"

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

//go:generate go run main.go
func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	commands := map[string]cli.CommandFactory{
		"schema": func() (cli.Command, error) {
			return SchemaModelsCommand{
				Ui: ui,
			}, nil
		},
	}

	c := &cli.CLI{
		Args:     os.Args[1:],
		Commands: commands,
	}

	code, err := c.Run()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

type SchemaModelsCommand struct {
	Ui cli.Ui
}

func (a SchemaModelsCommand) Help() string {
	return "Both -input and -output can be specified. " +
		"-input defaults to input/example.json. " +
		"-schema defaults to input/schema.json. " +
		"-output defaults to output. " +
		"A subset of schema can be generated using -include."
}

func (a SchemaModelsCommand) Synopsis() string {
	return "Generates schema."
}

func (a SchemaModelsCommand) Run(args []string) int {
	// parse flags
	conf, err := config.New()
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

	// validate against IR Schema
	err = validate.Schema(src)
	if err != nil {
		log.Println("The document is not valid. see errors :")
		log.Println(err)
		log.Fatal("The document is not valid. Terminating execution.")
	}

	// unmarshal JSON
	var s spec.Specification
	err = json.Unmarshal(src, &s)
	if err != nil {
		log.Fatal(err)
	}

	// ********************
	// DataSources
	// ********************
	// convert IR to framework schema
	c := datasource_convert.NewConverter(s)
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
	resourceSchemaConverter := resource_convert.NewConverter(s)
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
	providerSchemaConverter := provider_convert.NewConverter(s)
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
