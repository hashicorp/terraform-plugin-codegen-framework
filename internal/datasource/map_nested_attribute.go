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

type GeneratorMapNestedAttribute struct {
	ComputedOptionalRequired   convert.ComputedOptionalRequired
	CustomType                 *specschema.CustomType
	CustomTypeNestedCollection convert.CustomTypeNestedCollection
	DeprecationMessage         convert.DeprecationMessage
	Description                convert.Description
	NestedObject               GeneratorNestedAttributeObject
	NestedAttributeObject      convert.NestedAttributeObject
	Sensitive                  convert.Sensitive
	Validators                 specschema.MapValidators
	ValidatorsCustom           convert.ValidatorsCustom
}

func NewGeneratorMapNestedAttribute(name string, a *datasource.MapNestedAttribute) (GeneratorMapNestedAttribute, error) {
	if a == nil {
		return GeneratorMapNestedAttribute{}, fmt.Errorf("*datasource.MapNestedAttribute is nil")
	}

	attributes, err := NewAttributes(a.NestedObject.Attributes)

	if err != nil {
		return GeneratorMapNestedAttribute{}, err
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ct := convert.NewCustomTypeNestedCollection(a.CustomType)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	vco := convert.NewValidatorsCustom(convert.ValidatorTypeObject, a.NestedObject.Validators.CustomValidators())

	nat := convert.NewNestedAttributeObject(attributes, a.NestedObject.CustomType, vco, name)

	s := convert.NewSensitive(a.Sensitive)

	vcl := convert.NewValidatorsCustom(convert.ValidatorTypeMap, a.Validators.CustomValidators())

	return GeneratorMapNestedAttribute{
		ComputedOptionalRequired:   c,
		CustomType:                 a.CustomType,
		CustomTypeNestedCollection: ct,
		DeprecationMessage:         dm,
		Description:                d,
		NestedObject: GeneratorNestedAttributeObject{
			AssociatedExternalType: generatorschema.NewAssocExtType(a.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			CustomType:             a.NestedObject.CustomType,
			Validators:             a.NestedObject.Validators,
		},
		NestedAttributeObject: nat,
		Sensitive:             s,
		Validators:            a.Validators,
		ValidatorsCustom:      vcl,
	}, nil
}

func (g GeneratorMapNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorMapNestedAttribute
}

func (g GeneratorMapNestedAttribute) Imports() *generatorschema.Imports {
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

	// TODO: This should only be added if custom types (models) are being generated.
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

	if !g.NestedAttributeObject.Equal(h.NestedAttributeObject) {
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
	b.Write(g.ValidatorsCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorMapNestedAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.MapValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
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
