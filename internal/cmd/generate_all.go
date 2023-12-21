// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/input"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/validate"
)

type GenerateAllCommand struct {
	UI              cli.Ui
	flagIRInputPath string
	flagOutputPath  string
	flagPackageName string
}

func (cmd *GenerateAllCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("generate all", flag.ExitOnError)
	fs.StringVar(&cmd.flagIRInputPath, "input", "", "path to intermediate representation (JSON)")
	fs.StringVar(&cmd.flagOutputPath, "output", "./output", "directory path to output generated code files")
	fs.StringVar(&cmd.flagPackageName, "package", "", "name of Go package for generated code files")

	return fs
}

func (cmd *GenerateAllCommand) Help() string {
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

	strBuilder.WriteString("\nUsage: tfplugingen-framework generate all [<args>]\n\n")
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

func (cmd *GenerateAllCommand) Synopsis() string {
	return "Generate code for provider, resources, and data sources from an Intermediate Representation (IR) JSON file."
}

func (cmd *GenerateAllCommand) Run(args []string) int {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))

	fs := cmd.Flags()
	err := fs.Parse(args)
	if err != nil {
		logger.Error("error parsing command flags", "err", err)
		return 1
	}

	err = cmd.runInternal(ctx, logger)
	if err != nil {
		logger.Error("error executing command", "err", err)
		return 1
	}

	return 0
}

func (cmd *GenerateAllCommand) runInternal(ctx context.Context, logger *slog.Logger) error {
	// read input file
	src, err := input.Read(cmd.flagIRInputPath)
	if err != nil {
		return fmt.Errorf("error reading IR JSON: %w", err)
	}

	// validate JSON
	err = validate.JSON(src)
	if err != nil {
		return fmt.Errorf("error validating IR JSON: %w", err)
	}

	// parse and validate IR against specification
	spec, err := spec.Parse(ctx, src)
	if err != nil {
		return fmt.Errorf("error parsing IR JSON: %w", err)
	}

	err = generateDataSourceCode(ctx, spec, cmd.flagOutputPath, cmd.flagPackageName, "DataSource", logger)
	if err != nil {
		return fmt.Errorf("error generating data source code: %w", err)
	}

	err = generateResourceCode(ctx, spec, cmd.flagOutputPath, cmd.flagPackageName, "Resource", logger)
	if err != nil {
		return fmt.Errorf("error generating resource code: %w", err)
	}

	err = generateProviderCode(ctx, spec, cmd.flagOutputPath, cmd.flagPackageName, "Provider", logger)
	if err != nil {
		return fmt.Errorf("error generating provider code: %w", err)
	}

	return nil
}
