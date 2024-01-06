// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSingleNestedAttribute struct {
	AssociatedExternalType   *generatorschema.AssocExtType
	Attributes               generatorschema.GeneratorAttributes
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomType               *specschema.CustomType
	CustomTypeNestedObject   convert.CustomTypeNestedObject
	Default                  *specschema.ObjectDefault
	DefaultCustom            convert.DefaultCustom
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	PlanModifiers            specschema.ObjectPlanModifiers
	PlanModifiersCustom      convert.PlanModifiersCustom
	Sensitive                convert.Sensitive
	Validators               specschema.ObjectValidators
	ValidatorsCustom         convert.ValidatorsCustom
}

func NewGeneratorSingleNestedAttribute(name string, a *resource.SingleNestedAttribute) (GeneratorSingleNestedAttribute, error) {
	if a == nil {
		return GeneratorSingleNestedAttribute{}, fmt.Errorf("*resource.SingleNestedAttribute is nil")
	}

	attributes, err := NewAttributes(a.Attributes)

	if err != nil {
		return GeneratorSingleNestedAttribute{}, err
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ct := convert.NewCustomTypeNestedObject(a.CustomType, name)

	dc := convert.NewDefaultCustom(a.Default.CustomDefault())

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	pm := convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, a.PlanModifiers.CustomPlanModifiers())

	s := convert.NewSensitive(a.Sensitive)

	vc := convert.NewValidatorsCustom(convert.ValidatorTypeObject, a.Validators.CustomValidators())

	return GeneratorSingleNestedAttribute{
		AssociatedExternalType:   generatorschema.NewAssocExtType(a.AssociatedExternalType),
		Attributes:               attributes,
		ComputedOptionalRequired: c,
		CustomType:               a.CustomType,
		CustomTypeNestedObject:   ct,
		Default:                  a.Default,
		DefaultCustom:            dc,
		DeprecationMessage:       dm,
		Description:              d,
		PlanModifiers:            a.PlanModifiers,
		PlanModifiersCustom:      pm,
		Sensitive:                s,
		Validators:               a.Validators,
		ValidatorsCustom:         vc,
	}, nil
}

func (g GeneratorSingleNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSingleNestedAttribute
}

func (g GeneratorSingleNestedAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	if g.Default != nil {
		customDefaultImports := generatorschema.CustomDefaultImports(g.Default.Custom)
		imports.Append(customDefaultImports)
	}

	for _, v := range g.PlanModifiers {
		customPlanModifierImports := generatorschema.CustomPlanModifierImports(v.Custom)
		imports.Append(customPlanModifierImports)
	}

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	for _, v := range g.Attributes {
		imports.Append(v.Imports())
	}

	// TODO: This should only be added if custom types (models) are being generated.
	imports.Append(generatorschema.AttrImports())

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorSingleNestedAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
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

	if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.CustomTypeNestedObject.Equal(h.CustomTypeNestedObject) {
		return false
	}

	if !g.Default.Equal(h.Default) {
		return false
	}

	if !g.DefaultCustom.Equal(h.DefaultCustom) {
		return false
	}

	if !g.DeprecationMessage.Equal(h.DeprecationMessage) {
		return false
	}

	if !g.Description.Equal(h.Description) {
		return false
	}

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.PlanModifiersCustom.Equal(h.PlanModifiersCustom) {
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

func (g GeneratorSingleNestedAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	attributesSchema, err := g.Attributes.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.SingleNestedAttribute{\n", name))
	b.WriteString("Attributes: map[string]schema.Attribute{")
	b.WriteString(attributesSchema)
	b.WriteString("\n},\n")
	b.Write(g.CustomTypeNestedObject.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.PlanModifiersCustom.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.Write(g.DefaultCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorSingleNestedAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: name.ToPascalCase() + "Value",
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorSingleNestedAttribute) GetAttributes() generatorschema.GeneratorAttributes {
	return g.Attributes
}

func (g GeneratorSingleNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	objectType := generatorschema.NewCustomNestedObjectType(name, attributeAttrValues)

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

	objectValue := generatorschema.NewCustomNestedObjectValue(name, attributeTypes, attributeAttrTypes, attributeAttrValues, attributeCollectionTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.Attributes.SortedKeys()

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

func (g GeneratorSingleNestedAttribute) To() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}

func (g GeneratorSingleNestedAttribute) From() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("single nested type is not yet implemented"))
}
