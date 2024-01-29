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

type GeneratorSingleNestedAttribute struct {
	AssociatedExternalType *schema.AssocExtType
	Attributes             schema.GeneratorAttributes
	OptionalRequired       convert.OptionalRequired
	CustomType             convert.CustomTypeNestedObject
	DeprecationMessage     convert.DeprecationMessage
	Description            convert.Description
	Sensitive              convert.Sensitive
	Validators             convert.Validators
}

func NewGeneratorSingleNestedAttribute(name string, a *provider.SingleNestedAttribute) (GeneratorSingleNestedAttribute, error) {
	if a == nil {
		return GeneratorSingleNestedAttribute{}, fmt.Errorf("*provider.SingleNestedAttribute is nil")
	}

	attributes, err := NewAttributes(a.Attributes)

	if err != nil {
		return GeneratorSingleNestedAttribute{}, err
	}

	c := convert.NewOptionalRequired(a.OptionalRequired)

	ct := convert.NewCustomTypeNestedObject(a.CustomType, name)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	s := convert.NewSensitive(a.Sensitive)

	v := convert.NewValidators(convert.ValidatorTypeObject, a.Validators.CustomValidators())

	return GeneratorSingleNestedAttribute{
		AssociatedExternalType: schema.NewAssocExtType(a.AssociatedExternalType),
		Attributes:             attributes,
		OptionalRequired:       c,
		CustomType:             ct,
		DeprecationMessage:     dm,
		Description:            d,
		Sensitive:              s,
		Validators:             v,
	}, nil
}

func (g GeneratorSingleNestedAttribute) GeneratorSchemaType() schema.Type {
	return schema.GeneratorSingleNestedAttribute
}

func (g GeneratorSingleNestedAttribute) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.Validators.Imports())

	imports.Append(g.Attributes.Imports())

	imports.Append(schema.AttrImports())

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSingleNestedAttribute) Equal(ga schema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorSingleNestedAttribute)

	if !ok {
		return false
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.Attributes.Equal(h.Attributes) {
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

func (g GeneratorSingleNestedAttribute) Schema(name schema.FrameworkIdentifier) (string, error) {
	attributesSchema, err := g.Attributes.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.SingleNestedAttribute{\n", name))
	b.WriteString("Attributes: map[string]schema.Attribute{")
	b.WriteString(attributesSchema)
	b.WriteString("\n},\n")
	b.Write(g.CustomType.Schema())
	b.Write(g.OptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.Validators.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorSingleNestedAttribute) ModelField(name schema.FrameworkIdentifier) (model.Field, error) {
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

func (g GeneratorSingleNestedAttribute) GetAttributes() schema.GeneratorAttributes {
	return g.Attributes
}

func (g GeneratorSingleNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	objectType := schema.NewCustomNestedObjectType(name, attributeAttrValues)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeTypes, err := g.Attributes.AttributeTypes()

	if err != nil {
		return nil, err
	}

	attributeAttrTypes, err := g.Attributes.AttrTypes()

	if err != nil {
		return nil, err
	}

	attributeCollectionTypes, err := g.Attributes.CollectionTypes()

	if err != nil {
		return nil, err
	}

	objectValue := schema.NewCustomNestedObjectValue(name, attributeTypes, attributeAttrTypes, attributeAttrValues, attributeCollectionTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.Attributes.SortedKeys()

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

	return buf.Bytes(), nil
}

func (g GeneratorSingleNestedAttribute) ToFromFunctions(name string) ([]byte, error) {
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

func (g GeneratorSingleNestedAttribute) To() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}

func (g GeneratorSingleNestedAttribute) From() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}
