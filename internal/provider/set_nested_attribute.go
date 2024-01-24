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

type GeneratorSetNestedAttribute struct {
	OptionalRequired      convert.OptionalRequired
	CustomType            convert.CustomTypeNestedCollection
	DeprecationMessage    convert.DeprecationMessage
	Description           convert.Description
	NestedObject          GeneratorNestedAttributeObject
	NestedAttributeObject convert.NestedAttributeObject
	Sensitive             convert.Sensitive
	Validators            convert.Validators
}

func NewGeneratorSetNestedAttribute(name string, a *provider.SetNestedAttribute) (GeneratorSetNestedAttribute, error) {
	if a == nil {
		return GeneratorSetNestedAttribute{}, fmt.Errorf("*provider.SetNestedAttribute is nil")
	}

	attributes, err := NewAttributes(a.NestedObject.Attributes)

	if err != nil {
		return GeneratorSetNestedAttribute{}, err
	}

	c := convert.NewOptionalRequired(a.OptionalRequired)

	ct := convert.NewCustomTypeNestedCollection(a.CustomType)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	vo := convert.NewValidators(convert.ValidatorTypeObject, a.NestedObject.Validators.CustomValidators())

	nat := convert.NewNestedAttributeObject(attributes, a.NestedObject.CustomType, vo, name)

	s := convert.NewSensitive(a.Sensitive)

	vs := convert.NewValidators(convert.ValidatorTypeSet, a.Validators.CustomValidators())

	return GeneratorSetNestedAttribute{
		OptionalRequired:   c,
		CustomType:         ct,
		DeprecationMessage: dm,
		Description:        d,
		NestedObject: GeneratorNestedAttributeObject{
			AssociatedExternalType: schema.NewAssocExtType(a.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			CustomType:             a.NestedObject.CustomType,
			Validators:             a.NestedObject.Validators,
		},
		NestedAttributeObject: nat,
		Sensitive:             s,
		Validators:            vs,
	}, nil
}

func (g GeneratorSetNestedAttribute) GeneratorSchemaType() schema.Type {
	return schema.GeneratorSetNestedAttribute
}

func (g GeneratorSetNestedAttribute) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.Validators.Imports())

	imports.Append(g.NestedAttributeObject.Imports())

	imports.Append(schema.AttrImports())

	imports.Append(g.NestedObject.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSetNestedAttribute) Equal(ga schema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorSetNestedAttribute)

	if !ok {
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

	if !g.NestedObject.Equal(h.NestedObject) {
		return false
	}

	if !g.NestedAttributeObject.Equal(h.NestedAttributeObject) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorSetNestedAttribute) Schema(name schema.FrameworkIdentifier) (string, error) {
	nestedObjectSchema, err := g.NestedAttributeObject.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.SetNestedAttribute{\n", name))
	b.Write(nestedObjectSchema)
	b.Write(g.CustomType.Schema())
	b.Write(g.OptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.Validators.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorSetNestedAttribute) ModelField(name schema.FrameworkIdentifier) (model.Field, error) {
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

func (g GeneratorSetNestedAttribute) GetAttributes() schema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorSetNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.NestedObject.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	objectType := schema.NewCustomNestedObjectType(name, attributeAttrValues)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeTypes, err := g.NestedObject.Attributes.AttributeTypes()

	if err != nil {
		return nil, err
	}

	attributeAttrTypes, err := g.NestedObject.Attributes.AttrTypes()

	if err != nil {
		return nil, err
	}

	attributeCollectionTypes, err := g.NestedObject.Attributes.CollectionTypes()

	if err != nil {
		return nil, err
	}

	objectValue := schema.NewCustomNestedObjectValue(name, attributeTypes, attributeAttrTypes, attributeAttrValues, attributeCollectionTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.NestedObject.Attributes.SortedKeys()

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

	return buf.Bytes(), nil
}

func (g GeneratorSetNestedAttribute) ToFromFunctions(name string) ([]byte, error) {
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

func (g GeneratorSetNestedAttribute) To() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("set nested type is not yet implemented"))
}

func (g GeneratorSetNestedAttribute) From() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("set nested type is not yet implemented"))
}
