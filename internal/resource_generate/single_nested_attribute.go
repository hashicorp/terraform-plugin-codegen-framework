// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSingleNestedAttribute struct {
	schema.SingleNestedAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	Attributes    GeneratorAttributes
	CustomType    *specschema.CustomType
	Default       *specschema.ObjectDefault
	PlanModifiers []specschema.ObjectPlanModifier
	Validators    []specschema.ObjectValidator
}

func (g GeneratorSingleNestedAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	if g.Default != nil {
		customDefaultImports := generatorschema.CustomDefaultImports(g.Default.Custom)
		imports.Append(customDefaultImports)
	}

	for _, v := range g.PlanModifiers {
		customPlanModifierImports := generatorschema.CustomPlanModifierImports(v.Custom)
		imports.Append(customPlanModifierImports)
	}

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	for _, v := range g.Attributes {
		imports.Append(v.Imports())
	}

	return imports
}

func (g GeneratorSingleNestedAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorSingleNestedAttribute)
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

func (g GeneratorSingleNestedAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"AttributesString": g.Attributes.String,
		"getObjectDefault": getObjectDefault,
	}

	t, err := template.New("single_nested_attribute").Funcs(funcMap).Parse(singleNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSingleNestedAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSingleNestedAttribute) ModelField(name string) (model.Field, error) {
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

func (g GeneratorSingleNestedAttribute) GetAttributes() GeneratorAttributes {
	return g.Attributes
}

func (g GeneratorSingleNestedAttribute) validatorsEqual(x, y []specschema.ObjectValidator) bool {
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
