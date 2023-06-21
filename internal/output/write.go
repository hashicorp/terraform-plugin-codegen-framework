// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package output

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteDataSources(dataSourcesSchema, dataSourcesModels map[string][]byte, outputDir string) error {
	//func WriteDataSources(dataSourcesSchema, dataSourcesModels, dataSourcesHelpers map[string][]byte, outputDir string) error {
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

		//_, err = f.Write(dataSourcesHelpers[k])
		//if err != nil {
		//	return err
		//}
	}

	return nil
}

func WriteResources(resourcesSchema map[string][]byte, outputDir string) error {
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

		//_, err = f.DataSourcesSchema(buf.Bytes())
		//if err != nil {
		//return err
		//}
	}

	return nil
}

func WriteProviders(providersSchema map[string][]byte, outputDir string) error {
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

		//_, err = f.DataSourcesSchema(buf.Bytes())
		//if err != nil {
		//return err
		//}
	}

	return nil
}
