// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSingleNestedBlock struct {
	AssociatedExternalType *schema.AssocExtType
	Attributes             schema.GeneratorAttributes
	Blocks                 schema.GeneratorBlocks
	OptionalRequired       convert.OptionalRequired
	CustomType             convert.CustomTypeNestedObject
	DeprecationMessage     convert.DeprecationMessage
	Description            convert.Description
	Sensitive              convert.Sensitive
	Validators             convert.Validators
}

func NewGeneratorSingleNestedBlock(name string, b *provider.SingleNestedBlock) (GeneratorSingleNestedBlock, error) {
	if b == nil {
		return GeneratorSingleNestedBlock{}, fmt.Errorf("*provider.SingleNestedBlock is nil")
	}

	attributes, err := NewAttributes(b.Attributes)

	if err != nil {
		return GeneratorSingleNestedBlock{}, err
	}

	blocks, err := NewBlocks(b.Blocks)

	if err != nil {
		return GeneratorSingleNestedBlock{}, err
	}

	c := convert.NewOptionalRequired(b.OptionalRequired)

	ct := convert.NewCustomTypeNestedObject(b.CustomType, name)

	d := convert.NewDescription(b.Description)

	dm := convert.NewDeprecationMessage(b.DeprecationMessage)

	s := convert.NewSensitive(b.Sensitive)

	v := convert.NewValidators(convert.ValidatorTypeObject, b.Validators.CustomValidators())

	return GeneratorSingleNestedBlock{
		AssociatedExternalType: schema.NewAssocExtType(b.AssociatedExternalType),
		Attributes:             attributes,
		Blocks:                 blocks,
		OptionalRequired:       c,
		CustomType:             ct,
		DeprecationMessage:     dm,
		Description:            d,
		Sensitive:              s,
		Validators:             v,
	}, nil
}

func (g GeneratorSingleNestedBlock) GeneratorSchemaType() schema.Type {
	return schema.GeneratorSingleNestedBlock
}

func (g GeneratorSingleNestedBlock) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.Validators.Imports())

	imports.Append(g.Attributes.Imports())

	imports.Append(g.Blocks.Imports())

	imports.Append(schema.AttrImports())

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSingleNestedBlock) Equal(ga schema.GeneratorBlock) bool {
	h, ok := ga.(GeneratorSingleNestedBlock)

	if !ok {
		return false
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.Attributes.Equal(h.Attributes) {
		return false
	}

	if !g.Blocks.Equal(h.Blocks) {
		return false
	}

	if !g.OptionalRequired.Equal(h.OptionalRequired) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.DeprecationMessage.Equal(h.DeprecationMessage) {
		return false
	}

	if !g.Description.Equal(h.Description) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorSingleNestedBlock) Schema(name schema.FrameworkIdentifier) (string, error) {
	attributesSchema, err := g.Attributes.Schema()

	if err != nil {
		return "", err
	}

	blocksSchema, err := g.Blocks.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.SingleNestedBlock{\n", name))
	if attributesSchema != "" {
		b.WriteString("Attributes: map[string]schema.Attribute{")
		b.WriteString(attributesSchema)
		b.WriteString("\n},\n")
	}
	if blocksSchema != "" {
		b.WriteString("Blocks: map[string]schema.Block{")
		b.WriteString(blocksSchema)
		b.WriteString("\n},\n")
	}
	b.Write(g.CustomType.Schema())
	b.Write(g.OptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.Validators.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorSingleNestedBlock) ModelField(name schema.FrameworkIdentifier) (model.Field, error) {
	f := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: name.ToPascalCase() + "Value",
	}

	customValueType := g.CustomType.ValueType()

	if customValueType != "" {
		f.ValueType = customValueType
	}

	return f, nil
}

func (g GeneratorSingleNestedBlock) GetAttributes() schema.GeneratorAttributes {
	return g.Attributes
}

func (g GeneratorSingleNestedBlock) GetBlocks() schema.GeneratorBlocks {
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

	objectType := schema.NewCustomNestedObjectType(name, attributesBlocksAttrValues)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeTypes, err := g.Attributes.AttributeTypes()

	if err != nil {
		return nil, err
	}

	blockTypes, err := g.Blocks.BlockTypes()

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

	// Only attributes need to be processed here as we're only concerned with List, Map, and Set.
	attributeCollectionTypes, err := g.Attributes.CollectionTypes()

	if err != nil {
		return nil, err
	}

	objectValue := schema.NewCustomNestedObjectValue(name, attributesBlocksTypes, attributesBlocksAttrTypes, attributesBlocksAttrValues, attributeCollectionTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.Attributes.SortedKeys()

	blockKeys := g.Blocks.SortedKeys()

	// Recursively call CustomTypeAndValue() for each attribute that implements
	// CustomTypeAndValue interface (i.e, nested attributes).
	for _, k := range attributeKeys {
		if c, ok := g.Attributes[k].(schema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	for _, k := range blockKeys {
		if c, ok := g.Blocks[k].(schema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}

func (g GeneratorSingleNestedBlock) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	toFuncs, err := g.Attributes.ToFuncs()

	if err != nil {
		return nil, err
	}

	fromFuncs, _ := g.Attributes.FromFuncs()

	toFrom := schema.NewToFromNestedObject(name, g.AssociatedExternalType, toFuncs, fromFuncs)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.Attributes.SortedKeys()

	// Recursively call ToFromFunctions() for each attribute that implements
	// ToFrom interface.
	for _, k := range attributeKeys {
		if c, ok := g.Attributes[k].(schema.ToFrom); ok {
			b, err := c.ToFromFunctions(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}

func (g GeneratorSingleNestedBlock) To() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}

func (g GeneratorSingleNestedBlock) From() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}
