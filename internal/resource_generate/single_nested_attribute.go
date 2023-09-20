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

	AssociatedExternalType *generatorschema.AssocExtType
	Attributes             generatorschema.GeneratorAttributes
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.ObjectDefault
	PlanModifiers specschema.ObjectPlanModifiers
	Validators    specschema.ObjectValidators
}

func (g GeneratorSingleNestedAttribute) AssocExtType() *generatorschema.AssocExtType {
	return g.AssociatedExternalType
}

func (g GeneratorSingleNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSingleNestedAttribute
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

	// TODO: This should only be added if custom types (models) are being generated.
	imports.Append(generatorschema.AttrImports())

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSingleNestedAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorSingleNestedAttribute)

	if !ok {
		return false
	}

	for k := range g.Attributes {
		if _, ok := h.Attributes[k]; !ok {
			return false
		}

		if !g.Attributes[k].Equal(h.Attributes[k]) {
			return false
		}
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Default.Equal(h.Default) {
		return false
	}

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.SingleNestedAttribute.Equal(h.SingleNestedAttribute)
}

func (g GeneratorSingleNestedAttribute) ToString(name string) (string, error) {
	type singleNestedAttribute struct {
		Name                           string
		TypeValueName                  string
		Attributes                     string
		GeneratorSingleNestedAttribute GeneratorSingleNestedAttribute
	}

	attributesStr, err := g.Attributes.String()

	if err != nil {
		return "", err
	}

	l := singleNestedAttribute{
		Name:                           name,
		TypeValueName:                  model.SnakeCaseToCamelCase(name),
		Attributes:                     attributesStr,
		GeneratorSingleNestedAttribute: g,
	}

	t, err := template.New("single_nested_attribute").Parse(singleNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, l)
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

func (g GeneratorSingleNestedAttribute) GetAttributes() generatorschema.GeneratorAttributes {
	return g.Attributes
}
