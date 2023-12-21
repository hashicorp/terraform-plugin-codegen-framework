// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"strings"

	"github.com/hashicorp/cli"
)

type ScaffoldCommand struct {
	UI cli.Ui
}

func (cmd *ScaffoldCommand) Help() string {
	helpText := `
	Usage: tfplugingen-framework scaffold <subcommand> [<args>]
	
	  This command has subcommands for scaffolding Terraform Plugin Framework code.
	
	`
	return strings.TrimSpace(helpText)
}

func (a *ScaffoldCommand) Synopsis() string {
	return "Terraform Plugin Framework code scaffolding commands"
}

func (cmd *ScaffoldCommand) Run(args []string) int {
	return cli.RunResultHelp
}
