// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

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
	PlanModifiers specschema.ListPlanModifiers
	Validators    specschema.ListValidators
}

func (g GeneratorListNestedBlock) AssocExtType() *generatorschema.AssocExtType {
	return g.NestedObject.AssociatedExternalType
}

func (g GeneratorListNestedBlock) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorListNestedBlock
}

func (g GeneratorListNestedBlock) Imports() *generatorschema.Imports {
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

	customTypeImports = generatorschema.CustomTypeImports(g.NestedObject.CustomType)
	imports.Append(customTypeImports)

	for _, v := range g.NestedObject.PlanModifiers {
		customPlanModifierImports := generatorschema.CustomPlanModifierImports(v.Custom)
		imports.Append(customPlanModifierImports)
	}

	for _, v := range g.NestedObject.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	for _, v := range g.NestedObject.Attributes {
		imports.Append(v.Imports())
	}

	for _, v := range g.NestedObject.Blocks {
		imports.Append(v.Imports())
	}

	// TODO: This should only be added if custom types (models) are being generated.
	imports.Append(generatorschema.AttrImports())

	imports.Append(g.NestedObject.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorListNestedBlock) Equal(ga generatorschema.GeneratorBlock) bool {
	h, ok := ga.(GeneratorListNestedBlock)

	if !ok {
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

	if !g.NestedObject.Equal(h.NestedObject) {
		return false
	}

	return g.ListNestedBlock.Equal(h.ListNestedBlock)
}

func (g GeneratorListNestedBlock) Schema(name string) (string, error) {
	type block struct {
		Name                     string
		TypeValueName            string
		Attributes               string
		Blocks                   string
		GeneratorListNestedBlock GeneratorListNestedBlock
	}

	attributesStr, err := g.NestedObject.Attributes.Schema()

	if err != nil {
		return "", err
	}

	blocksStr, err := g.NestedObject.Blocks.Schema()

	if err != nil {
		return "", err
	}

	b := block{
		Name:                     name,
		TypeValueName:            model.SnakeCaseToCamelCase(name),
		Attributes:               attributesStr,
		Blocks:                   blocksStr,
		GeneratorListNestedBlock: g,
	}

	t, err := template.New("list_nested_block").Parse(listNestedBlockGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonBlockTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, b)
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

func (g GeneratorListNestedBlock) GetAttributes() generatorschema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorListNestedBlock) GetBlocks() generatorschema.GeneratorBlocks {
	return g.NestedObject.Blocks
}

func (g GeneratorListNestedBlock) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.NestedObject.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	blockAttrValues, err := g.NestedObject.Blocks.AttrValues()

	if err != nil {
		return nil, err
	}

	attributesBlocksAttrValues := make(map[string]string, len(g.NestedObject.Attributes)+len(g.NestedObject.Blocks))

	for k, v := range attributeAttrValues {
		attributesBlocksAttrValues[k] = v
	}

	for k, v := range blockAttrValues {
		attributesBlocksAttrValues[k] = v
	}

	objectType := generatorschema.NewCustomObjectType(name, attributesBlocksAttrValues)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeTypes, err := g.NestedObject.Attributes.AttributeTypes()

	if err != nil {
		return nil, err
	}

	blockTypes, err := g.NestedObject.Blocks.AttributeTypes()

	if err != nil {
		return nil, err
	}

	attributesBlocksTypes := make(map[string]string, len(g.NestedObject.Attributes)+len(g.NestedObject.Blocks))

	for k, v := range attributeTypes {
		attributesBlocksTypes[k] = v
	}

	for k, v := range blockTypes {
		attributesBlocksTypes[k] = v
	}

	attributeAttrTypes, err := g.NestedObject.Attributes.AttrTypes()

	if err != nil {
		return nil, err
	}

	blockAttrTypes, err := g.NestedObject.Blocks.AttrTypes()

	if err != nil {
		return nil, err
	}

	attributesBlocksAttrTypes := make(map[string]string, len(g.NestedObject.Attributes)+len(g.NestedObject.Blocks))

	for k, v := range attributeAttrTypes {
		attributesBlocksAttrTypes[k] = v
	}

	for k, v := range blockAttrTypes {
		attributesBlocksAttrTypes[k] = v
	}

	objectValue := generatorschema.NewCustomObjectValue(name, attributesBlocksTypes, attributesBlocksAttrTypes, attributesBlocksAttrValues)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.NestedObject.Attributes))

	for k := range g.NestedObject.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(g.NestedObject.Blocks))

	for k := range g.NestedObject.Blocks {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	// Recursively call CustomTypeAndValue() for each attribute that implements
	// CustomTypeAndValue interface (i.e, nested attributes).
	for _, k := range attributeKeys {
		if c, ok := g.NestedObject.Attributes[k].(generatorschema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	for _, k := range blockKeys {
		if c, ok := g.NestedObject.Blocks[k].(generatorschema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)

			continue
		}
	}

	return buf.Bytes(), nil
}
