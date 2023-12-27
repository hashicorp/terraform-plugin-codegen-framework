// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package output

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templating"
)

// WriteDataSources uses the packageName to determine whether to create a directory and package per data source.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per data source. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteDataSources(dataSourcesSchema, dataSourcesModels, customTypeValue, dataSourcesToFrom map[string]schema.GoCode, outputDir, packageName string) (map[string]templating.DataSourceTemplateData, error) {
	templateData := make(map[string]templating.DataSourceTemplateData, len(dataSourcesSchema))

	for k, v := range dataSourcesSchema {
		dirName := ""

		if packageName == "" {
			dirName = fmt.Sprintf("datasource_%s", k)

			err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
			if err != nil {
				return nil, err
			}
		}

		filename := fmt.Sprintf("%s_data_source_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return nil, err
		}

		// Format and write the data source schema
		formattedSchema, err := v.Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedSchema)
		if err != nil {
			return nil, err
		}

		// Format and write the data source model type
		formattedModel, err := dataSourcesModels[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedModel)
		if err != nil {
			return nil, err
		}

		// Format and write the data source custom type and value
		formattedCustomTypeValue, err := customTypeValue[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedCustomTypeValue)
		if err != nil {
			return nil, err
		}

		// Format and write the data source to/from helper methods
		formattedToFrom, err := dataSourcesToFrom[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedToFrom)
		if err != nil {
			return nil, err
		}

		identifier := schema.FrameworkIdentifier(k)
		templateData[k] = templating.DataSourceTemplateData{
			SnakeName:       identifier.ToString(),
			PascalName:      identifier.ToPascalCase(),
			CamelName:       identifier.ToCamelCase(),
			Package:         v.PackageName,
			SchemaFunc:      v.NotableExports[schema.ExportSchemaFunc],
			SchemaModelType: dataSourcesModels[k].NotableExports[schema.ExportSchemaModelType],
		}
	}

	return templateData, nil
}

// WriteResources uses the packageName to determine whether to create a directory and package per resource.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory per resource. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteResources(resourcesSchema, resourcesModels, customTypeValue, resourcesToFrom map[string]schema.GoCode, outputDir, packageName string) (map[string]templating.ResourceTemplateData, error) {
	templateData := make(map[string]templating.ResourceTemplateData, len(resourcesSchema))

	for k, v := range resourcesSchema {
		dirName := ""

		if packageName == "" {
			dirName = fmt.Sprintf("resource_%s", k)

			err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
			if err != nil {
				return nil, err
			}
		}

		filename := fmt.Sprintf("%s_resource_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return nil, err
		}

		// Format and write the resource schema
		formattedSchema, err := v.Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedSchema)
		if err != nil {
			return nil, err
		}

		// Format and write the resource model type
		formattedModel, err := resourcesModels[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedModel)
		if err != nil {
			return nil, err
		}

		// Format and write the resource custom type and value
		formattedCustomTypeValue, err := customTypeValue[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedCustomTypeValue)
		if err != nil {
			return nil, err
		}

		// Format and write the resource to/from helper methods
		formattedToFrom, err := resourcesToFrom[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedToFrom)
		if err != nil {
			return nil, err
		}

		identifier := schema.FrameworkIdentifier(k)
		templateData[k] = templating.ResourceTemplateData{
			SnakeName:       identifier.ToString(),
			PascalName:      identifier.ToPascalCase(),
			CamelName:       identifier.ToCamelCase(),
			Package:         v.PackageName,
			SchemaFunc:      v.NotableExports[schema.ExportSchemaFunc],
			SchemaModelType: resourcesModels[k].NotableExports[schema.ExportSchemaModelType],
		}
	}

	return templateData, nil
}

// WriteProviders uses the packageName to determine whether to create a directory and package for the provider.
// If packageName is an empty string, this indicates that the flag was not set, and the default behaviour is
// then to create a package and directory for the provider. If packageName is set then all generated code is
// placed into the same directory and package.
func WriteProviders(providersSchema, providerModels, customTypeValue, providerToFrom map[string]schema.GoCode, outputDir, packageName string) (map[string]templating.ProviderTemplateData, error) {
	templateData := make(map[string]templating.ProviderTemplateData, len(providersSchema))

	for k, v := range providersSchema {
		dirName := ""

		if packageName == "" {
			dirName = fmt.Sprintf("provider_%s", k)

			err := os.MkdirAll(filepath.Join(outputDir, dirName), os.ModePerm)
			if err != nil {
				return nil, err
			}
		}

		filename := fmt.Sprintf("%s_provider_gen.go", k)

		f, err := os.Create(filepath.Join(outputDir, dirName, filename))
		if err != nil {
			return nil, err
		}

		// Format and write the provider schema
		formattedSchema, err := v.Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedSchema)
		if err != nil {
			return nil, err
		}

		// Format and write the provider model type
		formattedModel, err := providerModels[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedModel)
		if err != nil {
			return nil, err
		}

		// Format and write the provider custom type and value
		formattedCustomTypeValue, err := customTypeValue[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedCustomTypeValue)
		if err != nil {
			return nil, err
		}

		// Format and write the provider to/from helper methods
		formattedToFrom, err := providerToFrom[k].Format()
		if err != nil {
			return nil, fmt.Errorf("error formatting Go code: %w", err)
		}

		_, err = f.Write(formattedToFrom)
		if err != nil {
			return nil, err
		}

		identifier := schema.FrameworkIdentifier(k)
		templateData[k] = templating.ProviderTemplateData{
			SnakeName:       identifier.ToString(),
			PascalName:      identifier.ToPascalCase(),
			CamelName:       identifier.ToCamelCase(),
			Package:         v.PackageName,
			SchemaFunc:      v.NotableExports[schema.ExportSchemaFunc],
			SchemaModelType: providerModels[k].NotableExports[schema.ExportSchemaModelType],
		}
	}

	return templateData, nil
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
