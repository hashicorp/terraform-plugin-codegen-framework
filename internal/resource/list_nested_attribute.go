// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorListNestedAttribute struct {
	ComputedOptionalRequired   convert.ComputedOptionalRequired
	CustomTypeNestedCollection convert.CustomTypeNestedCollection
	DefaultCustom              convert.DefaultCustom
	DeprecationMessage         convert.DeprecationMessage
	Description                convert.Description
	NestedObject               GeneratorNestedAttributeObject
	NestedAttributeObject      NestedAttributeObject
	PlanModifiersCustom        convert.PlanModifiersCustom
	Sensitive                  convert.Sensitive
	ValidatorsCustom           convert.ValidatorsCustom
}

func NewGeneratorListNestedAttribute(name string, a *resource.ListNestedAttribute) (GeneratorListNestedAttribute, error) {
	if a == nil {
		return GeneratorListNestedAttribute{}, fmt.Errorf("*resource.ListNestedAttribute is nil")
	}

	attributes, err := NewAttributes(a.NestedObject.Attributes)

	if err != nil {
		return GeneratorListNestedAttribute{}, err
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ct := convert.NewCustomTypeNestedCollection(a.CustomType)

	dc := convert.NewDefaultCustom(a.Default.CustomDefault())

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	pmo := convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, a.NestedObject.PlanModifiers.CustomPlanModifiers())

	vco := convert.NewValidatorsCustom(convert.ValidatorTypeObject, a.NestedObject.Validators.CustomValidators())

	nat := NewNestedAttributeObject(attributes, a.NestedObject.CustomType, pmo, vco, name)

	pml := convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, a.PlanModifiers.CustomPlanModifiers())

	s := convert.NewSensitive(a.Sensitive)

	vcl := convert.NewValidatorsCustom(convert.ValidatorTypeList, a.Validators.CustomValidators())

	return GeneratorListNestedAttribute{
		ComputedOptionalRequired:   c,
		CustomTypeNestedCollection: ct,
		DefaultCustom:              dc,
		DeprecationMessage:         dm,
		Description:                d,
		NestedObject: GeneratorNestedAttributeObject{
			AssociatedExternalType: generatorschema.NewAssocExtType(a.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			CustomType:             a.NestedObject.CustomType,
			Validators:             a.NestedObject.Validators,
		},
		NestedAttributeObject: nat,
		PlanModifiersCustom:   pml,
		Sensitive:             s,
		ValidatorsCustom:      vcl,
	}, nil
}

func (g GeneratorListNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorListNestedAttribute
}

func (g GeneratorListNestedAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(g.CustomTypeNestedCollection.Imports())

	imports.Append(g.DefaultCustom.Imports())

	imports.Append(g.PlanModifiersCustom.Imports())

	imports.Append(g.ValidatorsCustom.Imports())

	imports.Append(g.NestedAttributeObject.Imports())

	imports.Append(generatorschema.AttrImports())

	imports.Append(g.NestedObject.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorListNestedAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListNestedAttribute)

	if !ok {
		return false
	}

	if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
		return false
	}

	if !g.CustomTypeNestedCollection.Equal(h.CustomTypeNestedCollection) {
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

	if !g.NestedObject.Equal(h.NestedObject) {
		return false
	}

	if !g.NestedAttributeObject.Equal(h.NestedAttributeObject) {
		return false
	}

	if !g.PlanModifiersCustom.Equal(h.PlanModifiersCustom) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.ValidatorsCustom.Equal(h.ValidatorsCustom)
}

func (g GeneratorListNestedAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	nestedObjectSchema, err := g.NestedAttributeObject.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.ListNestedAttribute{\n", name))
	b.Write(nestedObjectSchema)
	b.Write(g.CustomTypeNestedCollection.Schema())
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

func (g GeneratorListNestedAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	f := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.ListValueType,
	}

	customValueType := g.CustomTypeNestedCollection.ValueType()

	if customValueType != "" {
		f.ValueType = customValueType
	}

	return f, nil
}

func (g GeneratorListNestedAttribute) GetAttributes() generatorschema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorListNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.NestedObject.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	objectType := generatorschema.NewCustomNestedObjectType(name, attributeAttrValues)

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

	objectValue := generatorschema.NewCustomNestedObjectValue(name, attributeTypes, attributeAttrTypes, attributeAttrValues, attributeCollectionTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeKeys := g.NestedObject.Attributes.SortedKeys()

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

func (g GeneratorListNestedAttribute) To() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("list nested type is not yet implemented"))
}

func (g GeneratorListNestedAttribute) From() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("list nested type is not yet implemented"))
}
