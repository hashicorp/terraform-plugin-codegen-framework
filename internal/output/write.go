// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package output

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func WriteDataSources(dataSourcesSchema, dataSourcesModels map[string][]byte, outputDir string) error {
	for k, v := range dataSourcesSchema {
		filename := fmt.Sprintf("%s_data_source_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, filename))
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
	}

	return nil
}

func WriteResources(resourcesSchema, resourcesModels map[string][]byte, outputDir string) error {
	for k, v := range resourcesSchema {
		filename := fmt.Sprintf("%s_resource_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, filename))
		if err != nil {
			return err
		}

		_, err = f.Write(v)
		if err != nil {
			return err
		}

		_, err = f.Write(resourcesModels[k])
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteProviders(providersSchema, providerModels map[string][]byte, outputDir string) error {
	for k, v := range providersSchema {
		filename := fmt.Sprintf("%s_provider_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, filename))
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
