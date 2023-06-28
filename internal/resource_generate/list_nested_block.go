// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorListNestedBlock struct {
	schema.ListNestedBlock

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	NestedObject  GeneratorNestedBlockObject
	PlanModifiers []specschema.ListPlanModifier
	Validators    []specschema.ListValidator
}

func (g GeneratorListNestedBlock) Imports() *generatorschema.Imports {
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

	for _, v := range g.PlanModifiers {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		for _, i := range v.Custom.Imports {
			if len(i.Path) > 0 {
				imports.Add(code.Import{
					Path: planModifierImport,
				})

				imports.Add(i)
			}
		}
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

	if g.NestedObject.CustomType != nil {
		if g.NestedObject.CustomType.HasImport() {
			imports.Add(*g.NestedObject.CustomType.Import)
		}
	}

	for _, v := range g.NestedObject.PlanModifiers {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		for _, i := range v.Custom.Imports {
			if len(i.Path) > 0 {
				imports.Add(code.Import{
					Path: planModifierImport,
				})

				imports.Add(i)
			}
		}
	}

	for _, v := range g.NestedObject.Validators {
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

	for _, v := range g.NestedObject.Attributes {
		imports.Add(v.Imports().All()...)
	}

	for _, v := range g.NestedObject.Blocks {
		imports.Add(v.Imports().All()...)
	}

	return imports
}

func (g GeneratorListNestedBlock) Equal(ga GeneratorBlock) bool {
	h, ok := ga.(GeneratorListNestedBlock)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.listValidatorsEqual(g.Validators, h.Validators) {
		return false
	}

	if !customTypeEqual(g.NestedObject.CustomType, h.NestedObject.CustomType) {
		return false
	}

	if !g.objectValidatorsEqual(g.NestedObject.Validators, h.NestedObject.Validators) {
		return false
	}

	for k, a := range g.NestedObject.Attributes {
		if !a.Equal(h.NestedObject.Attributes[k]) {
			return false
		}
	}

	return true
}

func (g GeneratorListNestedBlock) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
	}

	t, err := template.New("list_nested_block").Funcs(funcMap).Parse(listNestedBlockGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonBlockTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorListNestedBlock{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorListNestedBlock) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.ListValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorListNestedBlock) listValidatorsEqual(x, y []specschema.ListValidator) bool {
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

func (g GeneratorListNestedBlock) objectValidatorsEqual(x, y []specschema.ObjectValidator) bool {
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
