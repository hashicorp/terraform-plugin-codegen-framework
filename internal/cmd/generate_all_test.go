// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd_test

import (
	"testing"

	"github.com/hashicorp/cli"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
)

func TestGenerateAllCommand(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		irInputPath   string
		templatesPath string
		pkgName       string
		goldenFileDir string
	}{
		"specified_pkg_name": {
			irInputPath:   "testdata/custom_and_external/ir.json",
			pkgName:       "specified",
			goldenFileDir: "testdata/custom_and_external/all_output/specified_pkg_name",
		},
		"default_pkg_name": {
			irInputPath:   "testdata/custom_and_external/ir.json",
			goldenFileDir: "testdata/custom_and_external/all_output/default_pkg_name",
		},
		"templates_test": {
			irInputPath:   "testdata/templates_test/ir.json",
			templatesPath: "testdata/templates_test/codegen_templates",
			goldenFileDir: "testdata/templates_test/output",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testOutputDir := t.TempDir()
			mockUi := cli.NewMockUi()
			c := cmd.GenerateAllCommand{
				UI: mockUi,
			}

			args := []string{
				"--input", testCase.irInputPath,
				"--package", testCase.pkgName,
				"--output", testOutputDir,
				"--templates", testCase.templatesPath,
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `generate all` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}
