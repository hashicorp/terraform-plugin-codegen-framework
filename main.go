package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	"github/hashicorp/terraform-provider-code-generator/internal/cmd"
)

//go:generate go run main.go
func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	commands := map[string]cli.CommandFactory{
		"all": func() (cli.Command, error) {
			return cmd.AllCommand{
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
