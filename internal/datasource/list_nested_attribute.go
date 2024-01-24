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

type GeneratorListNestedAttribute struct {
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomType               convert.CustomTypeNestedCollection
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	NestedObject             GeneratorNestedAttributeObject
	NestedAttributeObject    convert.NestedAttributeObject
	Sensitive                convert.Sensitive
	Validators               convert.Validators
}

func NewGeneratorListNestedAttribute(name string, a *datasource.ListNestedAttribute) (GeneratorListNestedAttribute, error) {
	if a == nil {
		return GeneratorListNestedAttribute{}, fmt.Errorf("*datasource.ListNestedAttribute is nil")
	}

	attributes, err := NewAttributes(a.NestedObject.Attributes)

	if err != nil {
		return GeneratorListNestedAttribute{}, err
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ct := convert.NewCustomTypeNestedCollection(a.CustomType)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	vo := convert.NewValidators(convert.ValidatorTypeObject, a.NestedObject.Validators.CustomValidators())

	nat := convert.NewNestedAttributeObject(attributes, a.NestedObject.CustomType, vo, name)

	s := convert.NewSensitive(a.Sensitive)

	vl := convert.NewValidators(convert.ValidatorTypeList, a.Validators.CustomValidators())

	return GeneratorListNestedAttribute{
		ComputedOptionalRequired: c,
		CustomType:               ct,
		DeprecationMessage:       dm,
		Description:              d,
		NestedObject: GeneratorNestedAttributeObject{
			AssociatedExternalType: schema.NewAssocExtType(a.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			CustomType:             a.NestedObject.CustomType,
			Validators:             a.NestedObject.Validators,
		},
		NestedAttributeObject: nat,
		Sensitive:             s,
		Validators:            vl,
	}, nil
}

func (g GeneratorListNestedAttribute) GeneratorSchemaType() schema.Type {
	return schema.GeneratorListNestedAttribute
}

func (g GeneratorListNestedAttribute) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.Validators.Imports())

	imports.Append(g.NestedAttributeObject.Imports())

	imports.Append(schema.AttrImports())

	imports.Append(g.NestedObject.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorListNestedAttribute) Equal(ga schema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListNestedAttribute)

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

	if !g.NestedAttributeObject.Equal(h.NestedAttributeObject) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorListNestedAttribute) Schema(name schema.FrameworkIdentifier) (string, error) {
	nestedObjectSchema, err := g.NestedAttributeObject.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.ListNestedAttribute{\n", name))
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

func (g GeneratorListNestedAttribute) ModelField(name schema.FrameworkIdentifier) (model.Field, error) {
	f := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.ListValueType,
	}

	customValueType := g.CustomType.ValueType()

	if customValueType != "" {
		f.ValueType = customValueType
	}

	return f, nil
}

func (g GeneratorListNestedAttribute) GetAttributes() schema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorListNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
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
	// CustomTypeAndValue interface.
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

func (g GeneratorListNestedAttribute) ToFromFunctions(name string) ([]byte, error) {
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

func (g GeneratorListNestedAttribute) To() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("list nested type is not yet implemented"))
}

func (g GeneratorListNestedAttribute) From() (schema.ToFromConversion, error) {
	return schema.ToFromConversion{}, schema.NewUnimplementedError(errors.New("list nested type is not yet implemented"))
}
