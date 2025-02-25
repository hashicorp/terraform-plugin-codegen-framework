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
	}
	for name, testCase := range testCases {

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
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `generate all` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}
