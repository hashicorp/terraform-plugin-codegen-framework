// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

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
	CustomType *specschema.CustomType
	Validators specschema.ObjectValidators
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

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.SingleNestedBlock.Equal(h.SingleNestedBlock)
}

func (g GeneratorSingleNestedBlock) Schema(name string) (string, error) {
	type block struct {
		Name                       string
		TypeValueName              string
		Attributes                 string
		Blocks                     string
		GeneratorSingleNestedBlock GeneratorSingleNestedBlock
	}

	attributesStr, err := g.Attributes.Schema()

	if err != nil {
		return "", err
	}

	blocksStr, err := g.Blocks.Schema()

	if err != nil {
		return "", err
	}

	b := block{
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

	err = t.Execute(&buf, b)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSingleNestedBlock) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.SnakeCaseToCamelCase(name) + "Value",
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

func (g GeneratorSingleNestedBlock) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	blockAttrValues, err := g.Blocks.AttrValues()

	if err != nil {
		return nil, err
	}

	attributesBlocksAttrValues := make(map[string]string, len(g.Attributes)+len(g.Blocks))

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

	attributeTypes, err := g.Attributes.AttributeTypes()

	if err != nil {
		return nil, err
	}

	blockTypes, err := g.Blocks.AttributeTypes()

	if err != nil {
		return nil, err
	}

	attributesBlocksTypes := make(map[string]string, len(g.Attributes)+len(g.Blocks))

	for k, v := range attributeTypes {
		attributesBlocksTypes[k] = v
	}

	for k, v := range blockTypes {
		attributesBlocksTypes[k] = v
	}

	attributeAttrTypes, err := g.Attributes.AttrTypes()

	if err != nil {
		return nil, err
	}

	blockAttrTypes, err := g.Blocks.AttrTypes()

	if err != nil {
		return nil, err
	}

	attributesBlocksAttrTypes := make(map[string]string, len(g.Attributes)+len(g.Blocks))

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
	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(g.Blocks))

	for k := range g.Blocks {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	// Recursively call CustomTypeAndValue() for each attribute that implements
	// CustomTypeAndValue interface (i.e, nested attributes).
	for _, k := range attributeKeys {
		if c, ok := g.Attributes[k].(generatorschema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	for _, k := range blockKeys {
		if c, ok := g.Blocks[k].(generatorschema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}
