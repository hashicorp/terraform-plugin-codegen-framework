// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorListNestedBlock struct {
	ComputedOptionalRequired   convert.ComputedOptionalRequired
	CustomType                 *specschema.CustomType
	CustomTypeNestedCollection convert.CustomTypeNestedCollection
	DeprecationMessage         convert.DeprecationMessage
	Description                convert.Description
	NestedObject               GeneratorNestedBlockObject
	NestedBlockObject          convert.NestedBlockObject
	Sensitive                  convert.Sensitive
	Validators                 specschema.ListValidators
	ValidatorsCustom           convert.ValidatorsCustom
}

func NewGeneratorListNestedBlock(name string, b *datasource.ListNestedBlock) (GeneratorListNestedBlock, error) {
	if b == nil {
		return GeneratorListNestedBlock{}, fmt.Errorf("*datasource.ListNestedBlock is nil")
	}

	attributes, err := NewAttributes(b.NestedObject.Attributes)

	if err != nil {
		return GeneratorListNestedBlock{}, err
	}

	blocks, err := NewBlocks(b.NestedObject.Blocks)

	if err != nil {
		return GeneratorListNestedBlock{}, err
	}

	c := convert.NewComputedOptionalRequired(b.ComputedOptionalRequired)

	ct := convert.NewCustomTypeNestedCollection(b.CustomType)

	d := convert.NewDescription(b.Description)

	dm := convert.NewDeprecationMessage(b.DeprecationMessage)

	vco := convert.NewValidatorsCustom(convert.ValidatorTypeObject, b.NestedObject.Validators.CustomValidators())

	nbo := convert.NewNestedBlockObject(attributes, blocks, b.NestedObject.CustomType, vco, name)

	s := convert.NewSensitive(b.Sensitive)

	vcl := convert.NewValidatorsCustom(convert.ValidatorTypeList, b.Validators.CustomValidators())

	return GeneratorListNestedBlock{
		ComputedOptionalRequired:   c,
		CustomType:                 b.CustomType,
		CustomTypeNestedCollection: ct,
		DeprecationMessage:         dm,
		Description:                d,
		NestedObject: GeneratorNestedBlockObject{
			AssociatedExternalType: generatorschema.NewAssocExtType(b.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			Blocks:                 blocks,
			CustomType:             b.NestedObject.CustomType,
			Validators:             b.NestedObject.Validators,
		},
		NestedBlockObject: nbo,
		Sensitive:         s,
		Validators:        b.Validators,
		ValidatorsCustom:  vcl,
	}, nil
}

func (g GeneratorListNestedBlock) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorListNestedBlock
}

func (g GeneratorListNestedBlock) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	customTypeImports = generatorschema.CustomTypeImports(g.NestedObject.CustomType)
	imports.Append(customTypeImports)

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

	if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.CustomTypeNestedCollection.Equal(h.CustomTypeNestedCollection) {
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

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.ValidatorsCustom.Equal(h.ValidatorsCustom)
}

func (g GeneratorListNestedBlock) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	nestedObjectSchema, err := g.NestedBlockObject.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.ListNestedBlock{\n", name))
	b.Write(nestedObjectSchema)
	b.Write(g.CustomTypeNestedCollection.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorListNestedBlock) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
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

	// Only attributes need to be processed here as we're only concerned with List, Map, and Set.
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

func (g GeneratorListNestedBlock) ToFromFunctions(name string) ([]byte, error) {
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

func (g GeneratorListNestedBlock) To() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("list nested type is not yet implemented"))
}

func (g GeneratorListNestedBlock) From() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("list nested type is not yet implemented"))
}
