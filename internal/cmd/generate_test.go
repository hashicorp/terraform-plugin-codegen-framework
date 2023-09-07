// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd_test

import (
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
		t.Fatalf("mismatched file count in output directory, golden directory: %d file(s), test directory: %d file(s)", len(wantDirEntries), len(gotDirEntries))
	}

	for i, wantEntry := range wantDirEntries {
		gotEntry := gotDirEntries[i]

		if gotEntry.Name() != wantEntry.Name() {
			t.Errorf("mismatched file name, golden file name: %s, test file name: %s", wantEntry.Name(), gotEntry.Name())
			continue
		}

		if gotEntry.Type() != wantEntry.Type() {
			t.Errorf("mismatched file type, golden file type: %s, test file type: %s", wantEntry.Type(), gotEntry.Type())
			continue
		}

		compareFiles(t, path.Join(gotDirPath, gotEntry.Name()), path.Join(wantDirPath, wantEntry.Name()))
	}
}

func compareFiles(t *testing.T, got, want string) {
	gotFile, err := os.Open(got)

	if err != nil {
		t.Fatalf("unexpected error opening %s: %s", got, err)
	}

	gotFileInfo, err := gotFile.Stat()

	if err != nil {
		t.Fatalf("unexpected error stat %s: %s", got, err)
	}

	if gotFileInfo.IsDir() {
		dirEntries, err := os.ReadDir(got)

		if err != nil {
			t.Fatalf("unexpected error reading dir %s: %s", got, err)
		}

		for _, dirEntry := range dirEntries {
			compareFiles(t, path.Join(got, dirEntry.Name()), path.Join(want, dirEntry.Name()))
		}

		return
	}

	gotFileBytes, err := os.ReadFile(got)

	if err != nil {
		t.Fatalf("unexpected error reading `got` file: %s", err)
	}

	wantFileBytes, err := os.ReadFile(want)

	if err != nil {
		t.Fatalf("unexpected error reading `want` file: %s", err)
	}

	if diff := cmp.Diff(string(gotFileBytes), string(wantFileBytes)); diff != "" {
		t.Errorf("unexpected difference in %s: %s", got, diff)
	}
}
