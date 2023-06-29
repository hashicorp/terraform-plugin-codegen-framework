// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSingleNestedBlock struct {
	schema.SingleNestedBlock

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

func (g GeneratorSingleNestedBlock) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports.Add(*g.CustomType.Import)
		}
	} else {
		imports.Add(code.Import{
			Path: generatorschema.TypesImport,
		})
	}

	for _, v := range g.Validators {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		for _, i := range v.Custom.Imports {
			if len(i.Path) > 0 {
				imports.Add(code.Import{
					Path: generatorschema.ValidatorImport,
				})

				imports.Add(i)
			}
		}
	}

	for _, v := range g.Attributes {
		imports.Add(v.Imports().All()...)
	}

	for _, v := range g.Blocks {
		imports.Add(v.Imports().All()...)
	}

	return imports
}

func (g GeneratorSingleNestedBlock) Equal(ga GeneratorBlock) bool {
	h, ok := ga.(GeneratorSingleNestedBlock)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	for k, a := range g.Attributes {
		if !a.Equal(h.Attributes[k]) {
			return false
		}
	}

	return true
}

func (g GeneratorSingleNestedBlock) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
	}

	t, err := template.New("single_nested_block").Funcs(funcMap).Parse(singleNestedBlockGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonBlockTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSingleNestedBlock{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSingleNestedBlock) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.ObjectValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorSingleNestedBlock) validatorsEqual(x, y []specschema.ObjectValidator) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	//TODO: Sort before comparing.
	for k, v := range x {
		if !customValidatorsEqual(v.Custom, y[k].Custom) {
			return false
		}
	}

	return true
}
