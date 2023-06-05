// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

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

type GeneratorProviderSchemas struct {
	schemas map[string]GeneratorProviderSchema
	// TODO: Could add a field to hold custom templates that are used in calls to
	// attributeStringsFromGeneratorAttributes() and blockStringsFromGeneratorBlocks() funcs.
}

type GeneratorProviderSchema struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
}

func NewGeneratorProviderSchemas(schemas map[string]GeneratorProviderSchema) GeneratorProviderSchemas {
	return GeneratorProviderSchemas{
		schemas: schemas,
	}
}

func (g GeneratorProviderSchemas) ToBytes() (map[string][]byte, error) {
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

func (g GeneratorProviderSchemas) toBytes(name string, s GeneratorProviderSchema) ([]byte, error) {
	funcMap := template.FuncMap{
		"getAttributes": attributeStringsFromGeneratorAttributes,
		"getBlocks":     blockStringsFromGeneratorBlocks,
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
		GeneratorProviderSchema
	}{
		Name:                    name,
		GeneratorProviderSchema: s,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func attributeStringsFromGeneratorAttributes(attributes map[string]GeneratorAttribute) (string, error) {
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

func blockStringsFromGeneratorBlocks(blocks map[string]GeneratorBlock) (string, error) {
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
