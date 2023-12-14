// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSingleNestedAttribute struct {
	schema.SingleNestedAttribute

	AssociatedExternalType *generatorschema.AssocExtType
	Attributes             generatorschema.GeneratorAttributes
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType *specschema.CustomType
	Validators specschema.ObjectValidators
}

func NewGeneratorSingleNestedAttribute(a *datasource.SingleNestedAttribute) (GeneratorSingleNestedAttribute, error) {
	if a == nil {
		return GeneratorSingleNestedAttribute{}, fmt.Errorf("*datasource.SingleNestedAttribute is nil")
	}

	attributes := make(generatorschema.GeneratorAttributes, len(a.Attributes))

	for _, v := range a.Attributes {
		attribute, err := NewAttribute(v)

		if err != nil {
			return GeneratorSingleNestedAttribute{}, err
		}

		attributes[v.Name] = attribute
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	s := convert.NewSensitive(a.Sensitive)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	return GeneratorSingleNestedAttribute{
		SingleNestedAttribute: schema.SingleNestedAttribute{
			Required:            c.IsRequired(),
			Optional:            c.IsOptional(),
			Computed:            c.IsComputed(),
			Sensitive:           s.IsSensitive(),
			Description:         d.Description(),
			MarkdownDescription: d.Description(),
			DeprecationMessage:  dm.DeprecationMessage(),
		},

		AssociatedExternalType: generatorschema.NewAssocExtType(a.AssociatedExternalType),
		Attributes:             attributes,
		CustomType:             a.CustomType,
		Validators:             a.Validators,
	}, nil
}

func (g GeneratorSingleNestedAttribute) AssocExtType() *generatorschema.AssocExtType {
	return g.AssociatedExternalType
}

func (g GeneratorSingleNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSingleNestedAttribute
}

func (g GeneratorSingleNestedAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

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

	for k := range g.Attributes {
		if _, ok := h.Attributes[k]; !ok {
			return false
		}

		if !g.Attributes[k].Equal(h.Attributes[k]) {
			return false
		}
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.SingleNestedAttribute.Equal(h.SingleNestedAttribute)
}

func (g GeneratorSingleNestedAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                           string
		TypeValueName                  string
		Attributes                     string
		GeneratorSingleNestedAttribute GeneratorSingleNestedAttribute
	}

	attributesStr, err := g.Attributes.Schema()

	if err != nil {
		return "", err
	}

	a := attribute{
		Name:                           name.ToString(),
		TypeValueName:                  name.ToPascalCase(),
		Attributes:                     attributesStr,
		GeneratorSingleNestedAttribute: g,
	}

	t, err := template.New("single_nested_attribute").Parse(singleNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, a)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
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
