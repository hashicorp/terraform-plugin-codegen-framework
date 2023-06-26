// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

type GeneratorSchema interface {
	ToBytes() (map[string][]byte, error)
}

type GeneratorResourceSchemas struct {
	schemas map[string]GeneratorResourceSchema
	// TODO: Could add a field to hold custom templates that are used in calls to
	// getAttributes() and getBlocks() funcs.
}

type GeneratorResourceSchema struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
}

func NewGeneratorResourceSchemas(schemas map[string]GeneratorResourceSchema) GeneratorResourceSchemas {
	return GeneratorResourceSchemas{
		schemas: schemas,
	}
}

func (g GeneratorResourceSchemas) ToBytes(packageName string) (map[string][]byte, error) {
	schemasBytes := make(map[string][]byte, len(g.schemas))

	for k, s := range g.schemas {
		b, err := g.toBytes(k, s, packageName)

		if err != nil {
			return nil, err
		}

		schemasBytes[k] = b
	}

	return schemasBytes, nil
}

func (g GeneratorResourceSchemas) toBytes(name string, s GeneratorResourceSchema, packageName string) ([]byte, error) {
	funcMap := template.FuncMap{
		"getImports":    getImports,
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
	}

	t, err := template.New("schema").Funcs(funcMap).Parse(schemaGoTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		Name string
		GeneratorResourceSchema
		PackageName string
	}{
		Name:                    name,
		GeneratorResourceSchema: s,
		PackageName:             packageName,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getImports(schema GeneratorResourceSchema) (string, error) {
	var s strings.Builder

	var imports = make(map[string]struct{})

	for _, v := range schema.Attributes {
		for k := range v.Imports() {
			imports[k] = struct{}{}
		}
	}

	for _, v := range schema.Blocks {
		for k := range v.Imports() {
			imports[k] = struct{}{}
		}
	}

	for a := range imports {
		s.WriteString(fmt.Sprintf("%q\n", a))
	}

	return s.String(), nil
}

func getAttributes(attributes map[string]GeneratorAttribute) (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var keys = make([]string, 0, len(attributes))

	for k := range attributes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if attributes[k] == nil {
			continue
		}

		str, err := attributes[k].ToString(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

func getBlocks(blocks map[string]GeneratorBlock) (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee block order as maps are unordered in Go.
	var keys = make([]string, 0, len(blocks))

	for k := range blocks {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if blocks[k] == nil {
			continue
		}

		str, err := blocks[k].ToString(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

type GeneratorImport interface {
	Imports() map[string]struct{}
}

type GeneratorModel interface {
	ModelField(string) (model.Field, error)
}

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	ToString(string) (string, error)
	GeneratorImport
	GeneratorModel
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	ToString(string) (string, error)
	GeneratorImport
	GeneratorModel
}

type GeneratorNestedAttributeObject struct {
	Attributes    map[string]GeneratorAttribute
	CustomType    *specschema.CustomType
	PlanModifiers []specschema.ObjectPlanModifier
	Validators    []specschema.ObjectValidator
}

type GeneratorNestedBlockObject struct {
	Attributes    map[string]GeneratorAttribute
	Blocks        map[string]GeneratorBlock
	CustomType    *specschema.CustomType
	PlanModifiers []specschema.ObjectPlanModifier
	Validators    []specschema.ObjectValidator
}

func customTypeEqual(x, y *specschema.CustomType) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if x.Import == nil && y.Import != nil {
		return false
	}

	if x.Import != nil && y.Import == nil {
		return false
	}

	if x.Import != nil && y.Import != nil {
		if *x.Import != *y.Import {
			return false
		}
	}

	if x.Type != y.Type {
		return false
	}

	if x.ValueType != y.ValueType {
		return false
	}

	return true
}

func objectTypeEqual(x, y []specschema.ObjectAttributeType) bool {
	for k, v := range x {
		if v.Name != y[k].Name {
			return false
		}

		a := specschema.ElementType{
			Bool:    v.Bool,
			Float64: v.Float64,
			Int64:   v.Int64,
			List:    v.List,
			Map:     v.Map,
			Number:  v.Number,
			Object:  v.Object,
			Set:     v.Set,
			String:  v.String,
		}

		b := specschema.ElementType{
			Bool:    y[k].Bool,
			Float64: y[k].Float64,
			Int64:   y[k].Int64,
			List:    y[k].List,
			Map:     y[k].Map,
			Number:  y[k].Number,
			Object:  y[k].Object,
			Set:     y[k].Set,
			String:  y[k].String,
		}

		if !elementTypeEqual(a, b) {
			return false
		}
	}

	return true
}

func elementTypeEqual(x, y specschema.ElementType) bool {
	if x.Bool != nil && y.Bool != nil {
		return customTypeEqual(x.Bool.CustomType, y.Bool.CustomType)
	}

	if x.Float64 != nil && y.Float64 != nil {
		return customTypeEqual(x.Float64.CustomType, y.Float64.CustomType)
	}

	if x.Int64 != nil && y.Float64 != nil {
		return customTypeEqual(x.Int64.CustomType, y.Int64.CustomType)
	}

	if x.List != nil && y.List != nil {
		if !customTypeEqual(x.List.CustomType, y.List.CustomType) {
			return false
		}

		return elementTypeEqual(x.List.ElementType, y.List.ElementType)
	}

	if x.Map != nil && y.Map != nil {
		if !customTypeEqual(x.Map.CustomType, y.Map.CustomType) {
			return false
		}

		return elementTypeEqual(x.Map.ElementType, y.Map.ElementType)
	}

	if x.Number != nil && y.Number != nil {
		return customTypeEqual(x.Number.CustomType, y.Number.CustomType)
	}

	if x.Object != nil && y.Object != nil {
		return objectTypeEqual(x.Object, y.Object)
	}

	if x.Set != nil && y.Set != nil {
		if !customTypeEqual(x.Set.CustomType, y.Set.CustomType) {
			return false
		}

		return elementTypeEqual(x.Set.ElementType, y.Set.ElementType)
	}

	if x.String != nil && y.String != nil {
		return customTypeEqual(x.String.CustomType, y.String.CustomType)
	}

	return false
}

type ResourcesModelsGenerator struct {
}

func NewResourcesModelsGenerator() ResourcesModelsGenerator {
	return ResourcesModelsGenerator{}
}

func (d ResourcesModelsGenerator) Process(schemas map[string]GeneratorResourceSchema) (map[string][]byte, error) {
	resourcesModels := make(map[string][]byte, len(schemas))

	for name, schema := range schemas {
		var buf bytes.Buffer

		generatorResourceSchema := GeneratorResourceSchema{
			Attributes: schema.Attributes,
			Blocks:     schema.Blocks,
		}

		models, err := generatorResourceSchema.Model(name)
		if err != nil {
			return nil, err
		}

		for _, m := range models {
			buf.WriteString("\n" + m.String() + "\n")
		}

		resourcesModels[name] = buf.Bytes()
	}

	return resourcesModels, nil
}

func (g GeneratorResourceSchema) Model(name string) ([]model.Model, error) {
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
			generatorResourceSchema := GeneratorResourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorResourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorMapNestedAttribute:
			generatorResourceSchema := GeneratorResourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorResourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorSetNestedAttribute:
			generatorResourceSchema := GeneratorResourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorResourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorSingleNestedAttribute:
			generatorResourceSchema := GeneratorResourceSchema{
				Attributes: t.Attributes,
			}

			nestedModels, err = generatorResourceSchema.Model(attributeName)
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
			generatorResourceSchema := GeneratorResourceSchema{
				Attributes: t.NestedObject.Attributes,
				Blocks:     t.NestedObject.Blocks,
			}

			nestedModels, err = generatorResourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		case GeneratorSetNestedBlock:
			generatorResourceSchema := GeneratorResourceSchema{
				Attributes: t.NestedObject.Attributes,
				Blocks:     t.NestedObject.Blocks,
			}

			nestedModels, err = generatorResourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		case GeneratorSingleNestedBlock:
			generatorResourceSchema := GeneratorResourceSchema{
				Attributes: t.Attributes,
				Blocks:     t.Blocks,
			}

			nestedModels, err = generatorResourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		}

		models = append(models, nestedModels...)
	}

	return models, nil
}

func (g GeneratorResourceSchema) ModelFields() ([]model.Field, error) {
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
