// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd_test

import (
	"testing"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
)

func TestScaffoldProviderCommand(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		goldenFileDir string
	}{
		"provider scaffold": {
			goldenFileDir: "testdata/scaffold/provider",
		},
	}
	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testOutputDir := t.TempDir()
			mockUi := cli.NewMockUi()
			c := cmd.ScaffoldProviderCommand{
				UI: mockUi,
			}

			args := []string{
				"--name", "examplecloud",
				"--package", "scaffold",
				"--output-dir", testOutputDir,
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `scaffold provider` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}
