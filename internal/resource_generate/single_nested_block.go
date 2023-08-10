// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSingleNestedBlock struct {
	schema.SingleNestedBlock

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	AssociatedExternalType *specschema.AssociatedExternalType
	Attributes             generatorschema.GeneratorAttributes
	Blocks                 generatorschema.GeneratorBlocks
	CustomType             *specschema.CustomType
	PlanModifiers          []specschema.ObjectPlanModifier
	Validators             []specschema.ObjectValidator
}

func (g GeneratorSingleNestedBlock) AssocExtType() *generatorschema.AssocExtType {
	return generatorschema.NewAssocExtType(g.AssociatedExternalType)
}

func (g GeneratorSingleNestedBlock) AttrType() attr.Type {
	return types.ObjectType{
		//TODO: Add AttrTypes?
	}
}

func (g GeneratorSingleNestedBlock) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

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

	for _, v := range g.Blocks {
		imports.Append(v.Imports())
	}

	// TODO: This should only be added if model object helper functions are being generated.
	imports.Append(generatorschema.AttrImports())

	return imports
}

func (g GeneratorSingleNestedBlock) Equal(ga generatorschema.GeneratorBlock) bool {
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
		"AttributesString": g.Attributes.String,
		"BlocksString":     g.Blocks.String,
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

func (g GeneratorSingleNestedBlock) GetAttributes() generatorschema.GeneratorAttributes {
	return g.Attributes
}

func (g GeneratorSingleNestedBlock) GetBlocks() generatorschema.GeneratorBlocks {
	return g.Blocks
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
