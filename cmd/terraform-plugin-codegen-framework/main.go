// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mitchellh/cli"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
)

func main() {
	// TODO: Temporary name for CLI :)
	name := "terraform-plugin-codegen-framework"
	version := name + " Version " + version
	if commit != "" {
		version += " from commit " + commit
	}

	os.Exit(runCLI(
		name,
		version,
		os.Args[1:],
		os.Stdin,
		colorable.NewColorableStdout(),
		colorable.NewColorableStderr(),
	))
}

func initCommands(ui cli.Ui) map[string]cli.CommandFactory {
	generateFactory := func() (cli.Command, error) {
		return &cmd.GenerateCommand{
			UI: ui,
		}, nil
	}

	generateAllFactory := func() (cli.Command, error) {
		return &cmd.GenerateAllCommand{
			UI: ui,
		}, nil
	}

	generateResourcesFactory := func() (cli.Command, error) {
		return &cmd.GenerateResourcesCommand{
			UI: ui,
		}, nil
	}

	generateDataSourcesFactory := func() (cli.Command, error) {
		return &cmd.GenerateDataSourcesCommand{
			UI: ui,
		}, nil
	}

	generateProviderFactory := func() (cli.Command, error) {
		return &cmd.GenerateProviderCommand{
			UI: ui,
		}, nil
	}

	return map[string]cli.CommandFactory{
		"generate":              generateFactory,
		"generate all":          generateAllFactory,
		"generate resources":    generateResourcesFactory,
		"generate data-sources": generateDataSourcesFactory,
		"generate provider":     generateProviderFactory,
	}
}

func runCLI(name, version string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	ui := &cli.ColoredUi{
		ErrorColor: cli.UiColorRed,
		WarnColor:  cli.UiColorYellow,

		Ui: &cli.BasicUi{
			Reader:      stdin,
			Writer:      stdout,
			ErrorWriter: stderr,
		},
	}

	commands := initCommands(ui)
	frameworkGen := cli.CLI{
		Name:       name,
		Args:       args,
		Commands:   commands,
		HelpFunc:   cli.BasicHelpFunc(name),
		HelpWriter: stderr,
		Version:    version,
	}
	exitCode, err := frameworkGen.Run()
	if err != nil {
		return 1
	}

	return exitCode
}
