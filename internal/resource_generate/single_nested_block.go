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

type GeneratorSingleNestedBlock struct {
	schema.SingleNestedBlock

	AssociatedExternalType *generatorschema.AssocExtType
	Attributes             generatorschema.GeneratorAttributes
	Blocks                 generatorschema.GeneratorBlocks
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	PlanModifiers specschema.ObjectPlanModifiers
	Validators    specschema.ObjectValidators
}

func (g GeneratorSingleNestedBlock) AssocExtType() *generatorschema.AssocExtType {
	return g.AssociatedExternalType
}

func (g GeneratorSingleNestedBlock) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSingleNestedBlock
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

	// TODO: This should only be added if custom types (models) are being generated.
	imports.Append(generatorschema.AttrImports())

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSingleNestedBlock) Equal(ga generatorschema.GeneratorBlock) bool {
	h, ok := ga.(GeneratorSingleNestedBlock)

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

	for k := range g.Blocks {
		if _, ok := h.Blocks[k]; !ok {
			return false
		}

		if !g.Blocks[k].Equal(h.Blocks[k]) {
			return false
		}
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.SingleNestedBlock.Equal(h.SingleNestedBlock)
}

func (g GeneratorSingleNestedBlock) ToString(name string) (string, error) {
	type singleNestedBlock struct {
		Name                       string
		TypeValueName              string
		Attributes                 string
		Blocks                     string
		GeneratorSingleNestedBlock GeneratorSingleNestedBlock
	}

	attributesStr, err := g.Attributes.String()

	if err != nil {
		return "", err
	}

	blocksStr, err := g.Blocks.String()

	if err != nil {
		return "", err
	}

	l := singleNestedBlock{
		Name:                       name,
		TypeValueName:              model.SnakeCaseToCamelCase(name),
		Attributes:                 attributesStr,
		Blocks:                     blocksStr,
		GeneratorSingleNestedBlock: g,
	}

	t, err := template.New("single_nested_block").Parse(singleNestedBlockGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonBlockTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, l)
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
