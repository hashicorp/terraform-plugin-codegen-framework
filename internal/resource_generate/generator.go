// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
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

func (g GeneratorResourceSchemas) ToBytes() (map[string][]byte, error) {
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

func (g GeneratorResourceSchemas) toBytes(name string, s GeneratorResourceSchema) ([]byte, error) {
	funcMap := template.FuncMap{
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
	}{
		Name:                    name,
		GeneratorResourceSchema: s,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	ToString(string) (string, error)
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	ToString(string) (string, error)
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
