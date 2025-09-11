// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package output

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WriteDataSources uses the packageName to determine whether to create a directory and package per data source.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per data source. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteDataSources(dataSourcesSchema, dataSourcesModels, customTypeValue, dataSourcesToFrom map[string][]byte, outputDir, packageName string) error {
		for k, v := range dataSourcesSchema {
			dirName := ""

			if packageName == "" {
				dirName = fmt.Sprintf("datasource_%s", k)

				err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
				if err != nil {
					return err
				}
			}

			filename := fmt.Sprintf("%s_data_source_gen.go", k)

			// Combine all content first
			var allContent []byte
			allContent = append(allContent, v...)
			allContent = append(allContent, dataSourcesModels[k]...)
			allContent = append(allContent, customTypeValue[k]...)
			allContent = append(allContent, dataSourcesToFrom[k]...)

			// Deduplicate the combined content
			deduplicated, err := deduplicateGoCode(allContent)
			if err != nil {
				return err
			}

			f, err := os.Create(filepath.Join(outputDir, dirName, filename))
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = f.Write(deduplicated)
			if err != nil {
				return err
			}
		}

		return nil
}

// WriteResources uses the packageName to determine whether to create a directory and package per resource.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per resource. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteResources(resourcesSchema, resourcesModels, customTypeValue, resourcesToFrom map[string][]byte, outputDir, packageName string) error {
		for k, v := range resourcesSchema {
			dirName := ""

			if packageName == "" {
				dirName = fmt.Sprintf("resource_%s", k)

				err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
				if err != nil {
					return err
				}
			}

			filename := fmt.Sprintf("%s_resource_gen.go", k)

			// Combine all content first
			var allContent []byte
			allContent = append(allContent, v...)
			allContent = append(allContent, resourcesModels[k]...)
			allContent = append(allContent, customTypeValue[k]...)
			allContent = append(allContent, resourcesToFrom[k]...)

			// Deduplicate the combined content
			deduplicated, err := deduplicateGoCode(allContent)
			if err != nil {
				return err
			}

			f, err := os.Create(filepath.Join(outputDir, dirName, filename))
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = f.Write(deduplicated)
			if err != nil {
				return err
			}
		}

		return nil
}

// WriteProviders uses the packageName to determine whether to create a directory and package for the provider.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory for the provider. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteProviders(providersSchema, providerModels, customTypeValue, providerToFrom map[string][]byte, outputDir, packageName string) error {
	for k, v := range providersSchema {
		dirName := ""

		if packageName == "" {
			dirName = fmt.Sprintf("provider_%s", k)

			err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
			if err != nil {
				return err
			}
		}

		filename := fmt.Sprintf("%s_provider_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return err
		}

		_, err = f.Write(v)
		if err != nil {
			return err
		}

		_, err = f.Write(providerModels[k])
		if err != nil {
			return err
		}

		_, err = f.Write(customTypeValue[k])
		if err != nil {
			return err
		}

		_, err = f.Write(providerToFrom[k])
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteBytes(outputFilePath string, outputBytes []byte, forceOverwrite bool) error {
	if _, err := os.Stat(outputFilePath); !errors.Is(err, fs.ErrNotExist) && !forceOverwrite {
		return fmt.Errorf("file (%s) already exists and --force is false", outputFilePath)
	}

	f, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(outputBytes)
	if err != nil {
		return err
	}

	return nil
}

// deduplicateGoCode removes duplicate type and function declarations from Go source code
func deduplicateGoCode(content []byte) ([]byte, error) {
	source := string(content)
	lines := strings.Split(source, "\n")
	
	// Track seen declarations
	seen := make(map[string]bool)
	result := make([]string, 0, len(lines))
	
	i := 0
	for i < len(lines) {
		line := lines[i]
		trimmedLine := strings.TrimSpace(line)
		
		// Check for type declarations
		if strings.HasPrefix(trimmedLine, "type ") {
			// Extract type name
			fields := strings.Fields(trimmedLine)
			if len(fields) >= 2 {
				typeName := fields[1]
				key := "type:" + typeName
				
				if seen[key] {
					// Skip this entire type declaration
					i = skipGoDeclaration(lines, i)
					continue
				} else {
					seen[key] = true
				}
			}
		}
		
		// Check for function declarations
		if strings.HasPrefix(trimmedLine, "func ") {
			// Extract function name
			funcName := extractFunctionName(trimmedLine)
			if funcName != "" {
				key := "func:" + funcName
				
				if seen[key] {
					// Skip this entire function declaration
					i = skipGoDeclaration(lines, i)
					continue
				} else {
					seen[key] = true
				}
			}
		}
		
		// Check for var declarations
		if strings.HasPrefix(trimmedLine, "var _ ") {
			// Extract the type being checked
			parts := strings.Split(trimmedLine, "=")
			if len(parts) > 1 {
				rightPart := strings.TrimSpace(parts[1])
				// Extract type name from "TypeName{}" pattern
				if strings.HasSuffix(rightPart, "{}") {
					typeName := strings.TrimSpace(strings.TrimSuffix(rightPart, "{}"))
					key := "var:" + typeName
					
					if seen[key] {
						// Skip this var declaration
						i++
						continue
					} else {
						seen[key] = true
					}
				}
			}
		}
		
		result = append(result, line)
		i++
	}
	
	return []byte(strings.Join(result, "\n")), nil
}

// extractFunctionName extracts function name from a function declaration line
func extractFunctionName(line string) string {
	// Handle both regular functions and methods
	// func Name(...) or func (receiver) Name(...)
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return ""
	}
	
	if strings.HasPrefix(fields[1], "(") {
		// Method with receiver: func (r Type) Name(...)
		if len(fields) >= 4 {
			// Extract receiver type and method name to create unique identifier
			receiverPart := fields[2] // This should be the type name like "PrincipalType)"
			funcName := fields[3]
			
			// Clean up receiver type (remove closing parenthesis)
			receiverType := strings.TrimSuffix(receiverPart, ")")
			
			// Extract just the function name (remove parameters)
			if idx := strings.Index(funcName, "("); idx > 0 {
				funcName = funcName[:idx]
			}
			
			// Create unique key: ReceiverType.MethodName
			return receiverType + "." + funcName
		}
	} else {
		// Regular function: func Name(...)
		funcName := fields[1]
		if idx := strings.Index(funcName, "("); idx > 0 {
			return funcName[:idx]
		}
	}
	
	return ""
}

// skipGoDeclaration skips over a complete Go declaration (type, func, etc.)
func skipGoDeclaration(lines []string, start int) int {
	if start >= len(lines) {
		return start
	}
	
	line := strings.TrimSpace(lines[start])
	
	// If it's a single-line declaration, just skip it
	if !strings.Contains(line, "{") {
		return start + 1
	}
	
	// For multi-line declarations, count braces to find the end
	braceCount := 0
	i := start
	
	for i < len(lines) {
		currentLine := lines[i]
		
		// Count opening and closing braces
		for _, char := range currentLine {
			switch char {
			case '{':
				braceCount++
			case '}':
				braceCount--
			}
		}
		
		i++
		
		// If we've closed all braces, we're done
		if braceCount == 0 {
			break
		}
	}
	
	// Skip any empty lines after the declaration
	for i < len(lines) && strings.TrimSpace(lines[i]) == "" {
		i++
	}
	
	return i
}
