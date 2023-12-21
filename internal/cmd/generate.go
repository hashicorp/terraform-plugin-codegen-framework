// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"strings"

	"github.com/hashicorp/cli"
)

type GenerateCommand struct {
	UI cli.Ui
}

func (cmd *GenerateCommand) Help() string {
	helpText := `
	Usage: tfplugingen-framework generate <subcommand> [<args>]
	
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
