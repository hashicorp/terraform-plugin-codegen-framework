// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd_test

import (
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/cmd"
	"github.com/mitchellh/cli"
)

func TestAllCommand(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		irInputPath   string
		goldenFileDir string
	}{
		"all - custom_and_external": {
			irInputPath:   "testdata/all/custom_and_external/ir.json",
			goldenFileDir: "testdata/all/custom_and_external/output",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testOutputDir := t.TempDir()
			mockUi := cli.NewMockUi()
			c := cmd.GenerateCommand{
				UI:                  mockUi,
				GenerateResources:   true,
				GenerateDataSources: true,
				GenerateProvider:    true,
			}

			args := []string{
				"-input", testCase.irInputPath,
				"-output", testOutputDir,
			}

			exitCode := c.Run(args)
			if exitCode != 0 {
				t.Fatalf("unexpected error running `all` cmd: %s", mockUi.ErrorWriter.String())
			}

			compareDirectories(t, testCase.goldenFileDir, testOutputDir)
		})
	}
}

// TODO: currently doesn't compare nested directory files
func compareDirectories(t *testing.T, wantDirPath, gotDirPath string) {
	t.Helper()

	wantDirEntries, err := os.ReadDir(wantDirPath)
	if err != nil {
		t.Fatalf("unexpected error reading `want` directory: %s", err)
	}

	gotDirEntries, err := os.ReadDir(gotDirPath)
	if err != nil {
		t.Fatalf("unexpected error reading `got` directory: %s", err)
	}

	if len(gotDirEntries) != len(wantDirEntries) {
		t.Fatalf("mismatched files in output directory, wanted: %d file(s), got: %d file(s)", len(wantDirEntries), len(gotDirEntries))
	}

	for i, wantEntry := range wantDirEntries {
		gotEntry := gotDirEntries[i]

		if gotEntry.Name() != wantEntry.Name() {
			t.Errorf("mismatched file name, wanted: %s, got: %s", wantEntry.Name(), gotEntry.Name())
			continue
		}

		if gotEntry.Type() != wantEntry.Type() {
			t.Errorf("mismatched file type, wanted: %s, got: %s", wantEntry.Type(), gotEntry.Type())
			continue
		}

		gotFile, err := os.ReadFile(path.Join(gotDirPath, gotEntry.Name()))
		if err != nil {
			t.Fatalf("unexpected error reading `got` file: %s", err)
		}
		wantFile, _ := os.ReadFile(path.Join(wantDirPath, wantEntry.Name()))
		if err != nil {
			t.Fatalf("unexpected error reading `want` file: %s", err)
		}

		if diff := cmp.Diff(string(gotFile), string(wantFile)); diff != "" {
			t.Errorf("unexpected difference in %s: %s", wantEntry.Name(), diff)
		}
	}
}
