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

type GeneratorMapNestedAttribute struct {
	ComputedOptionalRequired   convert.ComputedOptionalRequired
	CustomTypeNestedCollection convert.CustomTypeNestedCollection
	DefaultCustom              convert.DefaultCustom
	DeprecationMessage         convert.DeprecationMessage
	Description                convert.Description
	NestedObject               GeneratorNestedAttributeObject
	NestedAttributeObject      NestedAttributeObject
	PlanModifiers              convert.PlanModifiers
	Sensitive                  convert.Sensitive
	Validators                 convert.Validators
}

func NewGeneratorMapNestedAttribute(name string, a *resource.MapNestedAttribute) (GeneratorMapNestedAttribute, error) {
	if a == nil {
		return GeneratorMapNestedAttribute{}, fmt.Errorf("*resource.MapNestedAttribute is nil")
	}

	attributes, err := NewAttributes(a.NestedObject.Attributes)

	if err != nil {
		return GeneratorMapNestedAttribute{}, err
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ct := convert.NewCustomTypeNestedCollection(a.CustomType)

	dc := convert.NewDefaultCustom(a.Default.CustomDefault())

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	pmo := convert.NewPlanModifiers(convert.PlanModifierTypeObject, a.NestedObject.PlanModifiers.CustomPlanModifiers())

	vo := convert.NewValidators(convert.ValidatorTypeObject, a.NestedObject.Validators.CustomValidators())

	nat := NewNestedAttributeObject(attributes, a.NestedObject.CustomType, pmo, vo, name)

	pmm := convert.NewPlanModifiers(convert.PlanModifierTypeMap, a.PlanModifiers.CustomPlanModifiers())

	s := convert.NewSensitive(a.Sensitive)

	vm := convert.NewValidators(convert.ValidatorTypeMap, a.Validators.CustomValidators())

	return GeneratorMapNestedAttribute{
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
		PlanModifiers:         pmm,
		Sensitive:             s,
		Validators:            vm,
	}, nil
}

func (g GeneratorMapNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorMapNestedAttribute
}

func (g GeneratorMapNestedAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(g.CustomTypeNestedCollection.Imports())

	imports.Append(g.DefaultCustom.Imports())

	imports.Append(g.PlanModifiers.Imports())

	imports.Append(g.Validators.Imports())

	imports.Append(g.NestedAttributeObject.Imports())

	imports.Append(generatorschema.AttrImports())

	imports.Append(g.NestedObject.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorMapNestedAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorMapNestedAttribute)

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

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorMapNestedAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	nestedObjectSchema, err := g.NestedAttributeObject.Schema()

	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.MapNestedAttribute{\n", name))
	b.Write(nestedObjectSchema)
	b.Write(g.CustomTypeNestedCollection.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.PlanModifiers.Schema())
	b.Write(g.Validators.Schema())
	b.Write(g.DefaultCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorMapNestedAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	f := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.MapValueType,
	}

	customValueType := g.CustomTypeNestedCollection.ValueType()

	if customValueType != "" {
		f.ValueType = customValueType
	}

	return f, nil
}

func (g GeneratorMapNestedAttribute) GetAttributes() generatorschema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorMapNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
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

func (g GeneratorMapNestedAttribute) ToFromFunctions(name string) ([]byte, error) {
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

func (g GeneratorMapNestedAttribute) To() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("map nested type is not yet implemented"))
}

func (g GeneratorMapNestedAttribute) From() (generatorschema.ToFromConversion, error) {
	return generatorschema.ToFromConversion{}, generatorschema.NewUnimplementedError(errors.New("map nested type is not yet implemented"))
}
