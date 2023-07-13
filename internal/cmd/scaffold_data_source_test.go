// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
	"github.com/mitchellh/cli"
)

func TestScaffoldDataSourceCommand(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		goldenFileDir string
	}{
		"data source scaffold": {
			goldenFileDir: "testdata/scaffold/data_source",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testOutputDir := t.TempDir()
			mockUi := cli.NewMockUi()
			c := cmd.ScaffoldDataSourceCommand{
				UI: mockUi,
			}

			args := []string{
				"--name", "thing",
				"--package", "scaffold",
				"--output-dir", testOutputDir,
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `scaffold data-source` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}
