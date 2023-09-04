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

func WriteDataSources(dataSourcesSchema, dataSourcesModels, dataSourcesModelObjectHelpers, dataSourcesToFrom map[string][]byte, outputDir string) error {
	for k, v := range dataSourcesSchema {
		dirName := fmt.Sprintf("datasource_%s", k)

		err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
		if err != nil {
			return err
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

		_, err = f.Write(dataSourcesModelObjectHelpers[k])
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

func WriteResources(resourcesSchema, resourcesModels, resourcesModelObjectHelpers, resourcesToFrom map[string][]byte, outputDir string) error {
	for k, v := range resourcesSchema {
		dirName := fmt.Sprintf("resource_%s", k)

		err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("%s_resource_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
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

		_, err = f.Write(resourcesModelObjectHelpers[k])
		if err != nil {
			return err
		}

		_, err = f.Write(resourcesToFrom[k])
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteProviders(providersSchema, providerModels, providerModelObjectHelpers, providerToFrom map[string][]byte, outputDir string) error {
	for k, v := range providersSchema {
		dirName := fmt.Sprintf("provider_%s", k)

		err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
		if err != nil {
			return err
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

		_, err = f.Write(providerModelObjectHelpers[k])
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
