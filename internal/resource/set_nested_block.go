// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSetNestedBlock struct {
	schema.SetNestedBlock

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	NestedObject  GeneratorNestedBlockObject
	PlanModifiers specschema.SetPlanModifiers
	Validators    specschema.SetValidators
}

func NewGeneratorSetNestedBlock(b *resource.SetNestedBlock) (GeneratorSetNestedBlock, error) {
	if b == nil {
		return GeneratorSetNestedBlock{}, fmt.Errorf("*resource.SetNestedBlock is nil")
	}

	attributes := make(generatorschema.GeneratorAttributes, len(b.NestedObject.Attributes))

	for _, v := range b.NestedObject.Attributes {
		attribute, err := NewAttribute(v)

		if err != nil {
			return GeneratorSetNestedBlock{}, err
		}

		attributes[v.Name] = attribute
	}

	blocks := make(generatorschema.GeneratorBlocks, len(b.NestedObject.Blocks))

	for _, v := range b.NestedObject.Blocks {
		block, err := NewBlock(v)

		if err != nil {
			return GeneratorSetNestedBlock{}, err
		}

		blocks[v.Name] = block
	}

	d := convert.NewDescription(b.Description)

	dm := convert.NewDeprecationMessage(b.DeprecationMessage)

	return GeneratorSetNestedBlock{
		SetNestedBlock: schema.SetNestedBlock{
			Description:         d.Description(),
			MarkdownDescription: d.Description(),
			DeprecationMessage:  dm.DeprecationMessage(),
		},

		CustomType: b.CustomType,
		NestedObject: GeneratorNestedBlockObject{
			AssociatedExternalType: generatorschema.NewAssocExtType(b.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			Blocks:                 blocks,
			CustomType:             b.NestedObject.CustomType,
			PlanModifiers:          b.NestedObject.PlanModifiers,
			Validators:             b.NestedObject.Validators,
		},
		PlanModifiers: b.PlanModifiers,
		Validators:    b.Validators,
	}, nil
}

func (g GeneratorSetNestedBlock) AssocExtType() *generatorschema.AssocExtType {
	return g.NestedObject.AssociatedExternalType
}

func (g GeneratorSetNestedBlock) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSetNestedBlock
}

func (g GeneratorSetNestedBlock) Imports() *generatorschema.Imports {
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

func (g GeneratorSetNestedBlock) Equal(ga generatorschema.GeneratorBlock) bool {
	h, ok := ga.(GeneratorSetNestedBlock)

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

	return g.SetNestedBlock.Equal(h.SetNestedBlock)
}

func (g GeneratorSetNestedBlock) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type block struct {
		Name                    string
		TypeValueName           string
		Attributes              string
		Blocks                  string
		GeneratorSetNestedBlock GeneratorSetNestedBlock
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
		Name:                    name.ToString(),
		TypeValueName:           name.ToPascalCase(),
		Attributes:              attributesStr,
		Blocks:                  blocksStr,
		GeneratorSetNestedBlock: g,
	}

	t, err := template.New("set_nested_block").Parse(setNestedBlockGoTemplate)
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

func (g GeneratorSetNestedBlock) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.SetValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorSetNestedBlock) GetAttributes() generatorschema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorSetNestedBlock) GetBlocks() generatorschema.GeneratorBlocks {
	return g.NestedObject.Blocks
}

func (g GeneratorSetNestedBlock) CustomTypeAndValue(name string) ([]byte, error) {
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

	objectType := generatorschema.NewCustomNestedObjectType(name, attributesBlocksAttrValues)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeTypes, err := g.NestedObject.Attributes.AttributeTypes()

	if err != nil {
		return nil, err
	}

	blockTypes, err := g.NestedObject.Blocks.BlockTypes()

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

	// Only attributes need to be processed here as we're only concerned with Set, Map, and Set.
	attributeCollectionTypes, err := g.NestedObject.Attributes.CollectionTypes()

	if err != nil {
		return nil, err
	}

	objectValue := generatorschema.NewCustomNestedObjectValue(name, attributesBlocksTypes, attributesBlocksAttrTypes, attributesBlocksAttrValues, attributeCollectionTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.NestedObject.Attributes.SortedKeys()

	blockKeys := g.NestedObject.Blocks.SortedKeys()

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
		}
	}

	return buf.Bytes(), nil
}

func (g GeneratorSetNestedBlock) ToFromFunctions(name string) ([]byte, error) {
	if g.NestedObject.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	toFuncs, err := g.NestedObject.Attributes.ToFuncs()

	if err != nil {
		return nil, err
	}

	fromFuncs, err := g.NestedObject.Attributes.FromFuncs()

	if err != nil {
		return nil, err
	}

	toFrom := generatorschema.NewToFromNestedObject(name, g.NestedObject.AssociatedExternalType, toFuncs, fromFuncs)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.NestedObject.Attributes.SortedKeys()

	// Recursively call ToFromFunctions() for each attribute that implements
	// ToFrom interface.
	for _, k := range attributeKeys {
		if c, ok := g.NestedObject.Attributes[k].(generatorschema.ToFrom); ok {
			b, err := c.ToFromFunctions(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}

func (g GeneratorSetNestedBlock) To() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("set nested type is not yet implemented"))
}

func (g GeneratorSetNestedBlock) From() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("set nested type is not yet implemented"))
}
