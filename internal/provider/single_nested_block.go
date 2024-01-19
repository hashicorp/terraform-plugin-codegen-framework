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
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSingleNestedBlock struct {
	AssociatedExternalType *generatorschema.AssocExtType
	Attributes             generatorschema.GeneratorAttributes
	Blocks                 generatorschema.GeneratorBlocks
	OptionalRequired       convert.OptionalRequired
	CustomTypeNestedObject convert.CustomTypeNestedObject
	DeprecationMessage     convert.DeprecationMessage
	Description            convert.Description
	Sensitive              convert.Sensitive
	ValidatorsCustom       convert.ValidatorsCustom
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

	vc := convert.NewValidatorsCustom(convert.ValidatorTypeObject, b.Validators.CustomValidators())

	return GeneratorSingleNestedBlock{
		AssociatedExternalType: generatorschema.NewAssocExtType(b.AssociatedExternalType),
		Attributes:             attributes,
		Blocks:                 blocks,
		OptionalRequired:       c,
		CustomTypeNestedObject: ct,
		DeprecationMessage:     dm,
		Description:            d,
		Sensitive:              s,
		ValidatorsCustom:       vc,
	}, nil
}

func (g GeneratorSingleNestedBlock) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSingleNestedBlock
}

func (g GeneratorSingleNestedBlock) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(g.CustomTypeNestedObject.Imports())

	imports.Append(g.ValidatorsCustom.Imports())

	imports.Append(g.Attributes.Imports())

	imports.Append(g.Blocks.Imports())

	imports.Append(generatorschema.AttrImports())

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSingleNestedBlock) Equal(ga generatorschema.GeneratorBlock) bool {
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

	if !g.CustomTypeNestedObject.Equal(h.CustomTypeNestedObject) {
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

	return g.ValidatorsCustom.Equal(h.ValidatorsCustom)
}

func (g GeneratorSingleNestedBlock) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
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
	b.Write(g.CustomTypeNestedObject.Schema())
	b.Write(g.OptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorSingleNestedBlock) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	f := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: name.ToPascalCase() + "Value",
	}

	customValueType := g.CustomTypeNestedObject.ValueType()

	if customValueType != "" {
		f.ValueType = customValueType
	}

	return f, nil
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

	objectType := generatorschema.NewCustomNestedObjectType(name, attributesBlocksAttrValues)

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

	objectValue := generatorschema.NewCustomNestedObjectValue(name, attributesBlocksTypes, attributesBlocksAttrTypes, attributesBlocksAttrValues, attributeCollectionTypes)

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

	toFrom := generatorschema.NewToFromNestedObject(name, g.AssociatedExternalType, toFuncs, fromFuncs)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.Attributes.SortedKeys()

	// Recursively call ToFromFunctions() for each attribute that implements
	// ToFrom interface.
	for _, k := range attributeKeys {
		if c, ok := g.Attributes[k].(generatorschema.ToFrom); ok {
			b, err := c.ToFromFunctions(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}

func (g GeneratorSingleNestedBlock) To() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}

func (g GeneratorSingleNestedBlock) From() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}
