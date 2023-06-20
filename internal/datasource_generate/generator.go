// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

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

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
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
