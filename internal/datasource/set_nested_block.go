// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSetNestedBlock struct {
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomType               convert.CustomTypeNestedCollection
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	NestedObject             GeneratorNestedBlockObject
	NestedBlockObject        convert.NestedBlockObject
	Sensitive                convert.Sensitive
	Validators               convert.Validators
}

func NewGeneratorSetNestedBlock(name string, b *datasource.SetNestedBlock) (GeneratorSetNestedBlock, error) {
	if b == nil {
		return GeneratorSetNestedBlock{}, fmt.Errorf("*datasource.SetNestedBlock is nil")
	}

	attributes, err := NewAttributes(b.NestedObject.Attributes)

	if err != nil {
		return GeneratorSetNestedBlock{}, err
	}

	blocks, err := NewBlocks(b.NestedObject.Blocks)

	if err != nil {
		return GeneratorSetNestedBlock{}, err
	}

	c := convert.NewComputedOptionalRequired(b.ComputedOptionalRequired)

	ct := convert.NewCustomTypeNestedCollection(b.CustomType)

	d := convert.NewDescription(b.Description)

	dm := convert.NewDeprecationMessage(b.DeprecationMessage)

	vo := convert.NewValidators(convert.ValidatorTypeObject, b.NestedObject.Validators.CustomValidators())

	nbo := convert.NewNestedBlockObject(attributes, blocks, b.NestedObject.CustomType, vo, name)

	s := convert.NewSensitive(b.Sensitive)

	vs := convert.NewValidators(convert.ValidatorTypeSet, b.Validators.CustomValidators())

	return GeneratorSetNestedBlock{
		ComputedOptionalRequired: c,
		CustomType:               ct,
		DeprecationMessage:       dm,
		Description:              d,
		NestedObject: GeneratorNestedBlockObject{
			AssociatedExternalType: schema.NewAssocExtType(b.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			Blocks:                 blocks,
			CustomType:             b.NestedObject.CustomType,
			Validators:             b.NestedObject.Validators,
		},
		NestedBlockObject: nbo,
		Sensitive:         s,
		Validators:        vs,
	}, nil
}

func (g GeneratorSetNestedBlock) GeneratorSchemaType() schema.Type {
	return schema.GeneratorSetNestedBlock
}

func (g GeneratorSetNestedBlock) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.Validators.Imports())

	imports.Append(g.NestedBlockObject.Imports())

	imports.Append(schema.AttrImports())

	imports.Append(g.NestedObject.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSetNestedBlock) Equal(ga schema.GeneratorBlock) bool {
	h, ok := ga.(GeneratorSetNestedBlock)

	if !ok {
		return false
	}

	if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
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

	if !g.NestedObject.Equal(h.NestedObject) {
		return false
	}

	if !g.NestedBlockObject.Equal(h.NestedBlockObject) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorSetNestedBlock) Schema(name schema.FrameworkIdentifier) (string, error) {
	nestedObjectSchema, err := g.NestedBlockObject.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.SetNestedBlock{\n", name))
	b.Write(nestedObjectSchema)
	b.Write(g.CustomType.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.Validators.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorSetNestedBlock) ModelField(name schema.FrameworkIdentifier) (model.Field, error) {
	f := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.SetValueType,
	}

	customValueType := g.CustomType.ValueType()

	if customValueType != "" {
		f.ValueType = customValueType
	}

	return f, nil
}

func (g GeneratorSetNestedBlock) GetAttributes() schema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorSetNestedBlock) GetBlocks() schema.GeneratorBlocks {
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

	objectType := schema.NewCustomNestedObjectType(name, attributesBlocksAttrValues)

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

	// Only attributes need to be processed here as we're only concerned with List, Map, and Set.
	attributeCollectionTypes, err := g.NestedObject.Attributes.CollectionTypes()

	if err != nil {
		return nil, err
	}

	objectValue := schema.NewCustomNestedObjectValue(name, attributesBlocksTypes, attributesBlocksAttrTypes, attributesBlocksAttrValues, attributeCollectionTypes)

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
		if c, ok := g.NestedObject.Attributes[k].(schema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	for _, k := range blockKeys {
		if c, ok := g.NestedObject.Blocks[k].(schema.CustomTypeAndValue); ok {
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

	toFrom := schema.NewToFromNestedObject(name, g.NestedObject.AssociatedExternalType, toFuncs, fromFuncs)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.NestedObject.Attributes.SortedKeys()

	// Recursively call ToFromFunctions() for each attribute that implements
	// ToFrom interface.
	for _, k := range attributeKeys {
		if c, ok := g.NestedObject.Attributes[k].(schema.ToFrom); ok {
			b, err := c.ToFromFunctions(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}

func (g GeneratorSetNestedBlock) To() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("set nested type is not yet implemented"))
}

func (g GeneratorSetNestedBlock) From() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("set nested type is not yet implemented"))
}
