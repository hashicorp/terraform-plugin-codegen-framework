package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	"github/hashicorp/terraform-provider-code-generator/internal/config"
	"github/hashicorp/terraform-provider-code-generator/internal/format"
	"github/hashicorp/terraform-provider-code-generator/internal/generate"
	"github/hashicorp/terraform-provider-code-generator/internal/input"
	"github/hashicorp/terraform-provider-code-generator/internal/output"
	"github/hashicorp/terraform-provider-code-generator/internal/transform"
	"github/hashicorp/terraform-provider-code-generator/internal/validate"
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
	errs := validate.Schema(conf.Input, conf.Schema)
	if errs != nil {
		log.Println("The document is not valid. see errors :")

		for _, e := range errs {
			log.Println(e)
		}

		log.Fatal("The document is not valid. Terminating execution.")
	}

	// unmarshal JSON
	ir, err := transform.Unmarshal(src)
	if err != nil {
		log.Fatal(err)
	}

	// generate code
	dataSourcesSchema, err := generate.DataSourcesSchema(ir, conf.Output)
	if err != nil {
		log.Fatal(err)
	}

	// format code
	formattedDataSourcesSchema, err := format.Format(dataSourcesSchema)
	if err != nil {
		log.Fatal(err)
	}

	// write code
	err = output.WriteDataSources(formattedDataSourcesSchema, conf.Output)
	if err != nil {
		log.Fatal(err)
	}

	return 0
}
