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

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return err
		}

		_, err = f.Write(v)
		if err != nil {
			return err
		}

		_, err = f.Write(dataSourcesModels[k])
		if err != nil {
			return err
		}

		_, err = f.Write(customTypeValue[k])
		if err != nil {
			return err
		}

		_, err = f.Write(dataSourcesToFrom[k])
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

			f, err := os.Create(filepath.Join(outputDir, dirName, filename))
			if err != nil {
				return err
			}

			// Deduplication logic
			written := make(map[string]struct{})

			writeUnique := func(code []byte) error {
				lines := strings.Split(string(code), "\n")
				var buffer []string
				for _, line := range lines {
					trim := strings.TrimSpace(line)
					if strings.HasPrefix(trim, "type ") || strings.HasPrefix(trim, "func ") || strings.HasPrefix(trim, "struct ") {
						name := trim
						if _, exists := written[name]; exists {
							continue // skip duplicate
						}
						written[name] = struct{}{}
					}
					buffer = append(buffer, line)
				}
				if len(buffer) > 0 {
					_, err := f.Write([]byte(strings.Join(buffer, "\n") + "\n"))
					if err != nil {
						return err
					}
				}
				return nil
			}

			if err := writeUnique(v); err != nil {
				return err
			}
			if err := writeUnique(resourcesModels[k]); err != nil {
				return err
			}
			if err := writeUnique(customTypeValue[k]); err != nil {
				return err
			}
			if err := writeUnique(resourcesToFrom[k]); err != nil {
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
