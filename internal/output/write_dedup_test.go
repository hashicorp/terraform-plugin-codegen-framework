package output_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/output"
)

func TestWriteResources_Deduplication(t *testing.T) {
	// Setup: create temp dir
	dir, err := ioutil.TempDir("", "dedup_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	// Simulate duplicate code in input
	resourceName := "myresource"
	code := []byte(`type Duplicated struct {}
type Duplicated struct {}
func Duplicated() {}
func Duplicated() {}
`)
	resourcesSchema := map[string][]byte{resourceName: code}
	resourcesModels := map[string][]byte{resourceName: code}
	customTypeValue := map[string][]byte{resourceName: code}
	resourcesToFrom := map[string][]byte{resourceName: code}

	// Run WriteResources
	err = output.WriteResources(resourcesSchema, resourcesModels, customTypeValue, resourcesToFrom, dir, "")
	if err != nil {
		t.Fatalf("WriteResources failed: %v", err)
	}

	// Read output file
	outFile := filepath.Join(dir, "resource_"+resourceName, resourceName+"_resource_gen.go")
	outBytes, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	outStr := string(outBytes)

	// Check only one struct and one func definition
	structCount := 0
	funcCount := 0
	for _, line := range splitLines(outStr) {
		if line == "type Duplicated struct {}" {
			structCount++
		}
		if line == "func Duplicated() {}" {
			funcCount++
		}
	}
	if structCount != 1 {
		t.Errorf("expected 1 struct, got %d", structCount)
	}
	if funcCount != 1 {
		t.Errorf("expected 1 func, got %d", funcCount)
	}
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := range s {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}
