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
	name := "tfplugingen-framework"
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

func initCommands(ui cli.Ui) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		// Code generation commands
		"generate":              commandFactory(&cmd.GenerateCommand{UI: ui}),
		"generate all":          commandFactory(&cmd.GenerateAllCommand{UI: ui}),
		"generate resources":    commandFactory(&cmd.GenerateResourcesCommand{UI: ui}),
		"generate data-sources": commandFactory(&cmd.GenerateDataSourcesCommand{UI: ui}),
		"generate provider":     commandFactory(&cmd.GenerateProviderCommand{UI: ui}),
		// Code scaffolding commands
		"scaffold":             commandFactory(&cmd.ScaffoldCommand{UI: ui}),
		"scaffold resource":    commandFactory(&cmd.ScaffoldResourceCommand{UI: ui}),
		"scaffold data-source": commandFactory(&cmd.ScaffoldDataSourceCommand{UI: ui}),
		"scaffold provider":    commandFactory(&cmd.ScaffoldProviderCommand{UI: ui}),
	}
}

func commandFactory(cmd cli.Command) cli.CommandFactory {
	return func() (cli.Command, error) {
		return cmd, nil
	}
}
