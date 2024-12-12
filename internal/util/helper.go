package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RemovingWhiteSpace(input string) string {
	input = strings.ReplaceAll(input, "\t", "") // 탭 제거
	input = strings.ReplaceAll(input, "\n", "") // 개행 제거
	input = strings.ReplaceAll(input, " ", "")  // 공백 제거
	return input
}

type Paths struct {
	AttributePath string
	ConfigPath    string
}

func InitializePaths(inputSpec, inputConfig, outputDir string) Paths {
	if inputSpec == "" || inputConfig == "" || outputDir == "" {
		log.Fatalf("Error: --input-spec, --input-config, and --output-dir are required.")
	}

	attributePath := MustAbs(inputSpec)
	configPath := MustAbs(inputConfig)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error occurred at making directory: %v", err)
	}

	return Paths{
		AttributePath: attributePath,
		ConfigPath:    configPath,
	}
}

func MustAbs(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Error getting absolute path for %s: %v", path, err)
	}
	return absPath
}

func SliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func ReadAndUnmarshal[T any](filePath string, result *T) error {
	data, err := os.ReadFile(fmt.Sprintf("./testfile/%s", filePath))
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("failed to unmarshal %s JSON: %w", filePath, err)
	}
	return nil
}
