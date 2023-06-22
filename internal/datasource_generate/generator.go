// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"
	"unicode"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type GeneratorSchema interface {
	ToBytes() (map[string][]byte, error)
}

type GeneratorDataSourceSchemas struct {
	schemas map[string]GeneratorDataSourceSchema
	// TODO: Could add a field to hold custom templates that are used in calls toBytes() to allow overriding.
	// getAttributes() and getBlocks() funcs.
}

type GeneratorDataSourceSchema struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
}

func NewGeneratorDataSourceSchemas(schemas map[string]GeneratorDataSourceSchema) GeneratorDataSourceSchemas {
	return GeneratorDataSourceSchemas{
		schemas: schemas,
	}
}

func (g GeneratorDataSourceSchemas) ToBytes() (map[string][]byte, error) {
	schemasBytes := make(map[string][]byte, len(g.schemas))

	for k, s := range g.schemas {
		b, err := g.toBytes(k, s)

		if err != nil {
			return nil, err
		}

		schemasBytes[k] = b
	}

	return schemasBytes, nil
}

func (g GeneratorDataSourceSchemas) toBytes(name string, s GeneratorDataSourceSchema) ([]byte, error) {
	funcMap := template.FuncMap{
		"getImports":    getImports,
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
	}

	t, err := template.New("schema").Funcs(funcMap).Parse(
		schemaGoTemplate,
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		Name string
		GeneratorDataSourceSchema
	}{
		Name:                      name,
		GeneratorDataSourceSchema: s,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getImports(schema GeneratorDataSourceSchema) (string, error) {
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
	ToModel(string) (string, error)
}

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	ToString(string) (string, error)
	GeneratorImport
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	ToString(string) (string, error)
	GeneratorImport
}

type GeneratorNestedAttributeObject struct {
	Attributes map[string]GeneratorAttribute
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

type GeneratorNestedBlockObject struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
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

type DataSourcesModelsGenerator struct {
}

func NewDataSourcesModelsGenerator() DataSourcesModelsGenerator {
	return DataSourcesModelsGenerator{}
}

func (d DataSourcesModelsGenerator) Process(schemas map[string]GeneratorDataSourceSchema) (map[string][]byte, error) {
	funcMap := template.FuncMap{
		"getModel": getModel,
	}

	datasourceModelTemplate, err := template.New("datasource_model.gotmpl").Funcs(funcMap).Parse(
		datasourceModelTmpl,
	)
	if err != nil {
		return nil, err
	}

	dataSourcesModels := make(map[string][]byte, len(schemas))

	for k, s := range schemas {
		var buf bytes.Buffer

		templateData := struct {
			Name       string
			Attributes map[string]GeneratorAttribute
			Blocks     map[string]GeneratorBlock
		}{
			Name:       k,
			Attributes: s.Attributes,
			Blocks:     s.Blocks,
		}

		// Generate model
		err = datasourceModelTemplate.Execute(&buf, templateData)
		if err != nil {
			return nil, err
		}

		nested, err := handleNested(s.Attributes, s.Blocks)
		if err != nil {
			return nil, err
		}

		buf.Write(nested)

		dataSourcesModels[k] = buf.Bytes()
	}

	return dataSourcesModels, nil
}

func handleNested(attributes map[string]GeneratorAttribute, blocks map[string]GeneratorBlock) ([]byte, error) {
	var buf bytes.Buffer

	// Using sorted attributeKeys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(attributes))

	for x := range attributes {
		attributeKeys = append(attributeKeys, x)
	}

	sort.Strings(attributeKeys)

	// If there are any nested attributes, generate model.
	for _, x := range attributeKeys {
		switch t := attributes[x].(type) {
		case GeneratorListNestedAttribute:
			model, err := nestedModel(x, t.NestedObject.Attributes, nil)
			if err != nil {
				return nil, err
			}

			buf.Write(model)
		case GeneratorMapNestedAttribute:
			model, err := nestedModel(x, t.NestedObject.Attributes, nil)
			if err != nil {
				return nil, err
			}

			buf.Write(model)
		case GeneratorSetNestedAttribute:
			model, err := nestedModel(x, t.NestedObject.Attributes, nil)
			if err != nil {
				return nil, err
			}

			buf.Write(model)
		case GeneratorSingleNestedAttribute:
			model, err := nestedModel(x, t.Attributes, nil)
			if err != nil {
				return nil, err
			}

			buf.Write(model)
		}
	}

	// Using sorted blockKeys to guarantee block order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(blocks))

	for x := range blocks {
		blockKeys = append(blockKeys, x)
	}

	sort.Strings(blockKeys)

	// If there are any nested blocks, generate model.
	for _, x := range blockKeys {
		switch t := blocks[x].(type) {
		case GeneratorListNestedBlock:
			model, err := nestedModel(x, t.NestedObject.Attributes, t.NestedObject.Blocks)
			if err != nil {
				return nil, err
			}

			buf.Write(model)
		case GeneratorSetNestedBlock:
			model, err := nestedModel(x, t.NestedObject.Attributes, t.NestedObject.Blocks)
			if err != nil {
				return nil, err
			}

			buf.Write(model)
		case GeneratorSingleNestedBlock:
			model, err := nestedModel(x, t.Attributes, t.Blocks)
			if err != nil {
				return nil, err
			}

			buf.Write(model)
		}
	}

	return buf.Bytes(), nil
}

func nestedModel(attribName string, attributes map[string]GeneratorAttribute, blocks map[string]GeneratorBlock) ([]byte, error) {
	var buf bytes.Buffer

	funcMap := template.FuncMap{
		"getModel": getModel,
		"lcFirst":  lcFirst,
	}

	datasourceModelTemplate, err := template.New("nested_model.gotmpl").Funcs(funcMap).Parse(
		nestedModelTmpl,
	)
	if err != nil {
		return nil, err
	}

	templateData := struct {
		Name       string
		Attributes map[string]GeneratorAttribute
		Blocks     map[string]GeneratorBlock
	}{
		Name:       attribName,
		Attributes: attributes,
		Blocks:     blocks,
	}

	err = datasourceModelTemplate.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	nested, err := handleNested(attributes, blocks)
	if err != nil {
		return nil, err
	}

	buf.Write(nested)

	return buf.Bytes(), nil

}

// snakeCaseToCamelCase relies on the convention of using snake-case
// names in configuration.
// TODO: A more robust approach is likely required here.
func snakeCaseToCamelCase(input string) string {
	inputSplit := strings.Split(input, "_")

	var ucName string

	for _, v := range inputSplit {
		if len(v) < 1 {
			continue
		}

		firstChar := v[0:1]
		ucFirstChar := strings.ToUpper(firstChar)

		if len(v) < 2 {
			ucName += ucFirstChar
			continue
		}

		ucName += ucFirstChar + v[1:]
	}

	return ucName
}

func lcFirst(input string) string {
	str := snakeCaseToCamelCase(input)

	for i, v := range str {
		preparedStr := string(unicode.ToLower(v)) + str[i+1:]
		return preparedStr
	}

	return str
}

func getModel(attributes map[string]GeneratorAttribute, blocks map[string]GeneratorBlock) (string, error) {
	var s strings.Builder

	// Using sorted attributeKeys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(attributes))

	for k := range attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	for _, k := range attributeKeys {
		if attributes[k] == nil {
			continue
		}

		// TODO: Remove once implemented across all generator attributes and blocks
		if m, ok := attributes[k].(GeneratorModel); ok {
			str, err := m.ToModel(k)

			if err != nil {
				return "", err
			}

			s.WriteString(str)
		}
	}

	// Using sorted blockKeys to guarantee block order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(blocks))

	for k := range blocks {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	for _, k := range blockKeys {
		if blocks[k] == nil {
			continue
		}

		// TODO: Remove once implemented across all generator attributes and blocks
		if m, ok := blocks[k].(GeneratorModel); ok {
			str, err := m.ToModel(k)

			if err != nil {
				return "", err
			}

			s.WriteString(str)
		}
	}

	return s.String(), nil
}
