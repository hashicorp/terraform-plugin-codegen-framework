// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package output

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/ncloud"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
)

// WriteDataSources uses the packageName to determine whether to create a directory and package per data source.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per data source. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteDataSources(dataSourcesSchema map[string][]byte, spec util.NcloudSpecification, outputDir, packageName string) error {
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

		n := ncloud.New(spec, k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return err
		}

		_, err = f.Write(v)
		if err != nil {
			return err
		}

		// CORE - 이곳에 코드를 추가한다.
		_, err = f.Write(n.RenderInitial())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderRead())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderModel())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderRefresh())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderWait())
		if err != nil {
			return err
		}

		filePath := f.Name()

		util.RemoveDuplicates(filePath)
	}

	return nil
}

// WriteResources uses the packageName to determine whether to create a directory and package per resource.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per resource. If packageName is set then all generated code is
// placed into the same directory and package.
// CORE - 여기에 줄을 추가하여 생성하는 것으로 한다.
func WriteResources(resourcesSchema map[string][]byte, spec util.NcloudSpecification, outputDir, packageName string) error {
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

		n := ncloud.New(spec, k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return err
		}

		_, err = f.Write(v)
		if err != nil {
			return err
		}

		// CORE - 이곳에 코드를 추가한다.
		_, err = f.Write(n.RenderInitial())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderImportState())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderCreate())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderRead())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderUpdate())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderDelete())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderModel())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderRefresh())
		if err != nil {
			return err
		}

		_, err = f.Write(n.RenderWait())
		if err != nil {
			return err
		}

		filePath := f.Name()

		util.RemoveDuplicates(filePath)
	}

	return nil
}

// WriteResources uses the packageName to determine whether to create a directory and package per resource.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per resource. If packageName is set then all generated code is
// placed into the same directory and package.
// CORE - 여기에 줄을 추가하여 생성하는 것으로 한다.
func WriteResourceTests(resourcesSchema map[string][]byte, spec util.NcloudSpecification, outputDir, packageName string) error {
	for k := range resourcesSchema {
		dirName := ""

		if packageName == "" {
			dirName = fmt.Sprintf("resource_%s", k)

			err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
			if err != nil {
				return err
			}
		}

		filename := fmt.Sprintf("%s_resource_gen_test.go", k)

		n := ncloud.New(spec, k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return err
		}

		// CORE - 이곳에 코드를 추가한다.
		_, err = f.Write(n.RenderTest())
		if err != nil {
			return err
		}

		filePath := f.Name()

		util.RemoveDuplicates(filePath)
	}

	return nil
}

// WriteDataSources uses the packageName to determine whether to create a directory and package per data source.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per data source. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteDataSourceTests(dataSourcesSchema map[string][]byte, spec util.NcloudSpecification, outputDir, packageName string) error {
	for k := range dataSourcesSchema {
		dirName := ""

		if packageName == "" {
			dirName = fmt.Sprintf("datasource_%s", k)

			err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
			if err != nil {
				return err
			}
		}

		filename := fmt.Sprintf("%s_data_source_gen_test.go", k)

		n := ncloud.New(spec, k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return err
		}

		// CORE - 이곳에 코드를 추가한다.
		_, err = f.Write(n.RenderTest())
		if err != nil {
			return err
		}

		filePath := f.Name()

		util.RemoveDuplicates(filePath)
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

		filePath := f.Name()

		util.RemoveDuplicates(filePath)
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

	filePath := f.Name()

	util.RemoveDuplicates(filePath)

	return nil
}
