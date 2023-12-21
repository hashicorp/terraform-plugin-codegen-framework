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

type ScaffoldDataSourceCommand struct {
	UI                      cli.Ui
	flagDataSourceNameSnake string
	flagOutputDir           string
	flagOutputFile          string
	flagPackageName         string
	flagForceOverwrite      bool
}

func (cmd *ScaffoldDataSourceCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("scaffold data source", flag.ExitOnError)

	fs.StringVar(&cmd.flagDataSourceNameSnake, "name", "", "name of data source in snake case without the provider type prefix, required")
	fs.BoolVar(&cmd.flagForceOverwrite, "force", false, "force creation, overwriting existing files")
	fs.StringVar(&cmd.flagOutputDir, "output-dir", ".", "directory path to output scaffolded code file")
	fs.StringVar(&cmd.flagOutputFile, "output-file", "", "file name and extension to write scaffolded code to, default will use the --name flag with '_data_source.go' suffix")
	fs.StringVar(&cmd.flagPackageName, "package", "provider", "name of Go package for scaffolded code file")
	return fs
}

func (cmd *ScaffoldDataSourceCommand) Help() string {
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

	strBuilder.WriteString("\nUsage: tfplugingen-framework scaffold data-source [<args>]\n\n")
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

func (a *ScaffoldDataSourceCommand) Synopsis() string {
	return "Create scaffolding code for a Terraform Plugin Framework data source."
}

func (cmd *ScaffoldDataSourceCommand) Run(args []string) int {
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

func (cmd *ScaffoldDataSourceCommand) runInternal(_ context.Context) error {
	if cmd.flagDataSourceNameSnake == "" {
		return errors.New("--name flag is required")
	}

	dataSourceIdentifier := schema.FrameworkIdentifier(cmd.flagDataSourceNameSnake)
	if !dataSourceIdentifier.Valid() {
		return fmt.Errorf("'%s' is not a valid Terraform data source identifier", cmd.flagDataSourceNameSnake)
	}

	goBytes, err := scaffold.DataSourceBytes(dataSourceIdentifier, cmd.flagPackageName)
	if err != nil {
		return fmt.Errorf("error creating scaffolding data source Go code: %w", err)
	}

	formattedGoBytes, err := format.Source(goBytes)
	if err != nil {
		return fmt.Errorf("error formatting scaffolding data source Go code: %w", err)
	}

	err = output.WriteBytes(cmd.getOutputFilePath(), formattedGoBytes, cmd.flagForceOverwrite)
	if err != nil {
		return fmt.Errorf("error writing scaffolding data source Go code: %w", err)
	}

	return nil
}

func (cmd *ScaffoldDataSourceCommand) getOutputFilePath() string {
	filename := fmt.Sprintf("%s_data_source.go", cmd.flagDataSourceNameSnake)
	if cmd.flagOutputFile != "" {
		filename = cmd.flagOutputFile
	}

	return filepath.Join(cmd.flagOutputDir, filename)
}
