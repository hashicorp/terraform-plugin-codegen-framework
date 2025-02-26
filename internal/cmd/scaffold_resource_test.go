// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd_test

import (
	"testing"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
)

func TestScaffoldResourceCommand(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		goldenFileDir string
	}{
		"resource scaffold": {
			goldenFileDir: "testdata/scaffold/resource",
		},
	}
	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testOutputDir := t.TempDir()
			mockUi := cli.NewMockUi()
			c := cmd.ScaffoldResourceCommand{
				UI: mockUi,
			}

			args := []string{
				"--name", "thing",
				"--package", "scaffold",
				"--output-dir", testOutputDir,
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `scaffold resource` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}
