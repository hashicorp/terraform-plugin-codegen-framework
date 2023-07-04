// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"sort"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

type DataSourcesModelsGenerator struct {
}

func NewDataSourcesModelsGenerator() DataSourcesModelsGenerator {
	return DataSourcesModelsGenerator{}
}

func (d DataSourcesModelsGenerator) Process(schemas map[string]GeneratorDataSourceSchema) (map[string][]byte, error) {
	dataSourcesModels := make(map[string][]byte, len(schemas))

	for name, schema := range schemas {
		var buf bytes.Buffer

		generatorDataSourceSchema := GeneratorDataSourceSchema{
			Attributes: schema.Attributes,
			Blocks:     schema.Blocks,
		}

		models, err := generatorDataSourceSchema.Model(name)
		if err != nil {
			return nil, err
		}

		for _, m := range models {
			buf.WriteString("\n" + m.String() + "\n")
		}

		dataSourcesModels[name] = buf.Bytes()
	}

	return dataSourcesModels, nil
}

func (g GeneratorDataSourceSchema) Model(name string) ([]model.Model, error) {
	var models []model.Model

	fields, err := g.ModelFields()
	if err != nil {
		return nil, err
	}

	m := model.Model{
		Name:   model.SnakeCaseToCamelCase(name),
		Fields: fields,
	}

	models = append(models, m)

	// Using sorted attributeNames to guarantee attribute order as maps are unordered in Go.
	var attributeNames = make([]string, 0, len(g.Attributes))

	for attributeName := range g.Attributes {
		attributeNames = append(attributeNames, attributeName)
	}

	sort.Strings(attributeNames)

	// If there are any nested attributes, generate model.
	for _, attributeName := range attributeNames {
		var nestedModels []model.Model

		switch t := g.Attributes[attributeName].(type) {
		case GeneratorListNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorMapNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorSetNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorSingleNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		}

		models = append(models, nestedModels...)
	}

	// Using sorted blockNames to guarantee block order as maps are unordered in Go.
	var blockNames = make([]string, 0, len(g.Blocks))

	for blockName := range g.Blocks {
		blockNames = append(blockNames, blockName)
	}

	sort.Strings(blockNames)

	// If there are any nested blocks, generate model.
	for _, blockName := range blockNames {
		var nestedModels []model.Model

		switch t := g.Blocks[blockName].(type) {
		case GeneratorListNestedBlock:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
				Blocks:     t.NestedObject.Blocks,
			}

			nestedModels, err = generatorDataSourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		case GeneratorSetNestedBlock:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
				Blocks:     t.NestedObject.Blocks,
			}

			nestedModels, err = generatorDataSourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		case GeneratorSingleNestedBlock:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.Attributes,
				Blocks:     t.Blocks,
			}

			nestedModels, err = generatorDataSourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		}

		models = append(models, nestedModels...)
	}

	return models, nil
}

func (g GeneratorDataSourceSchema) ModelFields() ([]model.Field, error) {
	var modelFields []model.Field

	// Using sorted attributeKeys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	for _, k := range attributeKeys {
		if g.Attributes[k] == nil {
			continue
		}

		modelField, err := g.Attributes[k].ModelField(k)

		if err != nil {
			return nil, err
		}

		modelFields = append(modelFields, modelField)
	}

	// Using sorted blockKeys to guarantee block order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(g.Blocks))

	for k := range g.Blocks {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	for _, k := range blockKeys {
		if g.Blocks[k] == nil {
			continue
		}

		modelField, err := g.Blocks[k].ModelField(k)

		if err != nil {
			return nil, err
		}

		modelFields = append(modelFields, modelField)
	}

	return modelFields, nil
}

type DataSourcesModelObjectHelpersGenerator struct {
}

func NewDataSourcesModelObjectHelpersGenerator() DataSourcesModelObjectHelpersGenerator {
	return DataSourcesModelObjectHelpersGenerator{}
}

func (d DataSourcesModelObjectHelpersGenerator) Process(schemas map[string]GeneratorDataSourceSchema) (map[string][]byte, error) {
	dataSourcesModelObjectHelpers := make(map[string][]byte, len(schemas))

	for name, s := range schemas {
		var buf bytes.Buffer

		g := GeneratorDataSourceSchema{
			Attributes: s.Attributes,
			Blocks:     s.Blocks,
		}

		var attributeKeys = make([]string, 0, len(g.Attributes))

		for k := range g.Attributes {
			attributeKeys = append(attributeKeys, k)
		}

		sort.Strings(attributeKeys)

		for _, k := range attributeKeys {
			if g.Attributes[k] == nil {
				continue
			}

			switch t := g.Attributes[k].(type) {
			case GeneratorListNestedAttribute:
				var hasNestedAttribute bool

				for _, v := range t.NestedObject.Attributes {
					switch v.(type) {
					case GeneratorListNestedAttribute:
						hasNestedAttribute = true
						break
					}
				}

				if hasNestedAttribute {
					modelObjectHelpers, err := t.ModelObjectHelpersString(k)

					if err != nil {
						return nil, err
					}

					buf.WriteString(modelObjectHelpers)
				}
			}
		}

		dataSourcesModelObjectHelpers[name] = buf.Bytes()
	}

	return dataSourcesModelObjectHelpers, nil
}
