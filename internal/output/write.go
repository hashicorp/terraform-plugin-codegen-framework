// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package output

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteDataSources(dataSourcesSchema, dataSourcesModels, dataSourcesModelObjectHelpers map[string][]byte, outputDir string) error {
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

		_, err = f.Write(dataSourcesModelObjectHelpers[k])
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
