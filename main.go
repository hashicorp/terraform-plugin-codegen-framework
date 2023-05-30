package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	"github/hashicorp/terraform-provider-code-generator/internal/cmd"
	"github/hashicorp/terraform-provider-code-generator/internal/validate"
)

func main() {
	// TODO: Move set-up of ui, validators etc to boostrap so that
	// we can just call c.Run()
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	irValidator := validate.NewIntermediateRepresentationValidator()
	specValidator := validate.NewSpecValidator()

	commands := map[string]cli.CommandFactory{
		"all": func() (cli.Command, error) {
			return cmd.AllCommand{
				Ui:            ui,
				IRValidator:   irValidator,
				SpecValidator: specValidator,
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
