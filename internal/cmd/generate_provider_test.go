// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd_test

import (
	"testing"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
)

func TestGenerateProviderCommand(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		irInputPath   string
		goldenFileDir string
	}{
		"custom_and_external": {
			irInputPath:   "testdata/custom_and_external/ir.json",
			goldenFileDir: "testdata/custom_and_external/provider_output",
		},
		"provider_no_attributes": {
			irInputPath:   "testdata/provider_no_attributes/ir.json",
			goldenFileDir: "testdata/provider_no_attributes/provider_output",
		},
	}
	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testOutputDir := t.TempDir()
			mockUi := cli.NewMockUi()
			c := cmd.GenerateProviderCommand{
				UI: mockUi,
			}

			args := []string{
				"--input", testCase.irInputPath,
				"--package", "generated",
				"--output", testOutputDir,
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `generate provider` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}
