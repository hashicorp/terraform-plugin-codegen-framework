package cmd_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
	"github.com/mitchellh/cli"
)

func TestGenerateDataSourcesCommand(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		irInputPath   string
		goldenFileDir string
	}{
		"custom_and_external": {
			irInputPath:   "testdata/custom_and_external/ir.json",
			goldenFileDir: "testdata/custom_and_external/data_sources_output",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testOutputDir := t.TempDir()
			mockUi := cli.NewMockUi()
			c := cmd.GenerateDataSourcesCommand{
				UI: mockUi,
			}

			args := []string{
				"--input", testCase.irInputPath,
				"--output", testOutputDir,
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `generate data-sources` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}
