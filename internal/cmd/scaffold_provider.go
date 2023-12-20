// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"path/filepath"
	"strings"

	"github.com/hashicorp/cli"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/scaffold"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type ScaffoldProviderCommand struct {
	UI                    cli.Ui
	flagProviderNameSnake string
	flagOutputDir         string
	flagOutputFile        string
	flagPackageName       string
	flagForceOverwrite    bool
}

func (cmd *ScaffoldProviderCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("scaffold provider", flag.ExitOnError)

	fs.StringVar(&cmd.flagProviderNameSnake, "name", "", "name of provider in snake case, required")
	fs.BoolVar(&cmd.flagForceOverwrite, "force", false, "force creation, overwriting existing files")
	fs.StringVar(&cmd.flagOutputDir, "output-dir", ".", "directory path to output scaffolded code file")
	fs.StringVar(&cmd.flagOutputFile, "output-file", "", "file name and extension to write scaffolded code to, default is 'provider.go'")
	fs.StringVar(&cmd.flagPackageName, "package", "provider", "name of Go package for scaffolded code file")
	return fs
}

func (cmd *ScaffoldProviderCommand) Help() string {
	strBuilder := &strings.Builder{}

	longestName := 0
	longestUsage := 0
	cmd.Flags().VisitAll(func(f *flag.Flag) {
		if len(f.Name) > longestName {
			longestName = len(f.Name)
		}
		if len(f.Usage) > longestUsage {
			longestUsage = len(f.Usage)
		}
	})

	strBuilder.WriteString("\nUsage: tfplugingen-framework scaffold provider [<args>]\n\n")
	cmd.Flags().VisitAll(func(f *flag.Flag) {
		if f.DefValue != "" {
			strBuilder.WriteString(fmt.Sprintf("    --%s <ARG> %s%s%s  (default: %q)\n",
				f.Name,
				strings.Repeat(" ", longestName-len(f.Name)+2),
				f.Usage,
				strings.Repeat(" ", longestUsage-len(f.Usage)+2),
				f.DefValue,
			))
		} else {
			strBuilder.WriteString(fmt.Sprintf("    --%s <ARG> %s%s%s\n",
				f.Name,
				strings.Repeat(" ", longestName-len(f.Name)+2),
				f.Usage,
				strings.Repeat(" ", longestUsage-len(f.Usage)+2),
			))
		}
	})
	strBuilder.WriteString("\n")

	return strBuilder.String()
}

func (a *ScaffoldProviderCommand) Synopsis() string {
	return "Create scaffolding code for a Terraform Plugin Framework provider."
}

func (cmd *ScaffoldProviderCommand) Run(args []string) int {
	ctx := context.Background()

	fs := cmd.Flags()
	err := fs.Parse(args)
	if err != nil {
		cmd.UI.Error(fmt.Sprintf("error parsing command flags: %s", err))
		return 1
	}

	err = cmd.runInternal(ctx)
	if err != nil {
		cmd.UI.Error(fmt.Sprintf("Error executing command: %s\n", err))
		return 1
	}

	return 0
}

func (cmd *ScaffoldProviderCommand) runInternal(_ context.Context) error {
	if cmd.flagProviderNameSnake == "" {
		return errors.New("--name flag is required")
	}

	providerIdentifier := schema.FrameworkIdentifier(cmd.flagProviderNameSnake)
	if !providerIdentifier.Valid() {
		return fmt.Errorf("'%s' is not a valid Terraform provider identifier", cmd.flagProviderNameSnake)
	}

	goBytes, err := scaffold.ProviderBytes(providerIdentifier, cmd.flagPackageName)
	if err != nil {
		return fmt.Errorf("error creating scaffolding provider Go code: %w", err)
	}

	formattedGoBytes, err := format.Source(goBytes)
	if err != nil {
		return fmt.Errorf("error formatting scaffolding provider Go code: %w", err)
	}

	err = output.WriteBytes(cmd.getOutputFilePath(), formattedGoBytes, cmd.flagForceOverwrite)
	if err != nil {
		return fmt.Errorf("error writing scaffolding provider Go code: %w", err)
	}

	return nil
}

func (cmd *ScaffoldProviderCommand) getOutputFilePath() string {
	filename := "provider.go"
	if cmd.flagOutputFile != "" {
		filename = cmd.flagOutputFile
	}

	return filepath.Join(cmd.flagOutputDir, filename)
}
