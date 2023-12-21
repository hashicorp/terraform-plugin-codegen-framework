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

type ScaffoldResourceCommand struct {
	UI                    cli.Ui
	flagResourceNameSnake string
	flagOutputDir         string
	flagOutputFile        string
	flagPackageName       string
	flagForceOverwrite    bool
}

func (cmd *ScaffoldResourceCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("scaffold resource", flag.ExitOnError)

	fs.StringVar(&cmd.flagResourceNameSnake, "name", "", "name of resource in snake case without the provider type prefix, required")
	fs.BoolVar(&cmd.flagForceOverwrite, "force", false, "force creation, overwriting existing files")
	fs.StringVar(&cmd.flagOutputDir, "output-dir", ".", "directory path to output scaffolded code file")
	fs.StringVar(&cmd.flagOutputFile, "output-file", "", "file name and extension to write scaffolded code to, default will use the --name flag with '_resource.go' suffix")
	fs.StringVar(&cmd.flagPackageName, "package", "provider", "name of Go package for scaffolded code file")
	return fs
}

func (cmd *ScaffoldResourceCommand) Help() string {
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

	strBuilder.WriteString("\nUsage: tfplugingen-framework scaffold resource [<args>]\n\n")
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

func (a *ScaffoldResourceCommand) Synopsis() string {
	return "Create scaffolding code for a Terraform Plugin Framework resource."
}

func (cmd *ScaffoldResourceCommand) Run(args []string) int {
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

func (cmd *ScaffoldResourceCommand) runInternal(_ context.Context) error {
	if cmd.flagResourceNameSnake == "" {
		return errors.New("--name flag is required")
	}

	resourceIdentifier := schema.FrameworkIdentifier(cmd.flagResourceNameSnake)
	if !resourceIdentifier.Valid() {
		return fmt.Errorf("'%s' is not a valid Terraform resource identifier", cmd.flagResourceNameSnake)
	}

	goBytes, err := scaffold.ResourceBytes(resourceIdentifier, cmd.flagPackageName)
	if err != nil {
		return fmt.Errorf("error creating scaffolding resource Go code: %w", err)
	}

	formattedGoBytes, err := format.Source(goBytes)
	if err != nil {
		return fmt.Errorf("error formatting scaffolding resource Go code: %w", err)
	}

	err = output.WriteBytes(cmd.getOutputFilePath(), formattedGoBytes, cmd.flagForceOverwrite)
	if err != nil {
		return fmt.Errorf("error writing scaffolding resource Go code: %w", err)
	}

	return nil
}

func (cmd *ScaffoldResourceCommand) getOutputFilePath() string {
	filename := fmt.Sprintf("%s_resource.go", cmd.flagResourceNameSnake)
	if cmd.flagOutputFile != "" {
		filename = cmd.flagOutputFile
	}

	return filepath.Join(cmd.flagOutputDir, filename)
}
