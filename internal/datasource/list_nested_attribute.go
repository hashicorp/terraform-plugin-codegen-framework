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

type GeneratorListNestedAttribute struct {
	schema.ListNestedAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType   *specschema.CustomType
	NestedObject GeneratorNestedAttributeObject
	Validators   specschema.ListValidators
}

func NewGeneratorListNestedAttribute(a *datasource.ListNestedAttribute) (GeneratorListNestedAttribute, error) {
	if a == nil {
		return GeneratorListNestedAttribute{}, fmt.Errorf("*datasource.ListNestedAttribute is nil")
	}

	attributes := make(generatorschema.GeneratorAttributes, len(a.NestedObject.Attributes))

	for _, v := range a.NestedObject.Attributes {
		attribute, err := NewAttribute(v)

		if err != nil {
			return GeneratorListNestedAttribute{}, err
		}

		attributes[v.Name] = attribute
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	s := convert.NewSensitive(a.Sensitive)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	return GeneratorListNestedAttribute{
		ListNestedAttribute: schema.ListNestedAttribute{
			Required:            c.IsRequired(),
			Optional:            c.IsOptional(),
			Computed:            c.IsComputed(),
			Sensitive:           s.IsSensitive(),
			Description:         d.Description(),
			MarkdownDescription: d.Description(),
			DeprecationMessage:  dm.DeprecationMessage(),
		},

		CustomType: a.CustomType,
		NestedObject: GeneratorNestedAttributeObject{
			AssociatedExternalType: generatorschema.NewAssocExtType(a.NestedObject.AssociatedExternalType),
			Attributes:             attributes,
			CustomType:             a.NestedObject.CustomType,
			Validators:             a.NestedObject.Validators,
		},
		Validators: a.Validators,
	}, nil
}

func (g GeneratorListNestedAttribute) AssocExtType() *generatorschema.AssocExtType {
	return g.NestedObject.AssociatedExternalType
}

func (g GeneratorListNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorListNestedAttribute
}

func (g GeneratorListNestedAttribute) Imports() *generatorschema.Imports {
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

func (g GeneratorListNestedAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListNestedAttribute)

	if !ok {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	if !g.NestedObject.Equal(h.NestedObject) {
		return false
	}

	return g.ListNestedAttribute.Equal(h.ListNestedAttribute)
}

func (g GeneratorListNestedAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                         string
		TypeValueName                string
		Attributes                   string
		GeneratorListNestedAttribute GeneratorListNestedAttribute
	}

	attributesStr, err := g.NestedObject.Attributes.Schema()

	if err != nil {
		return "", err
	}

	a := attribute{
		Name:                         name.ToString(),
		TypeValueName:                name.ToPascalCase(),
		Attributes:                   attributesStr,
		GeneratorListNestedAttribute: g,
	}

	t, err := template.New("list_nested_attribute").Parse(listNestedAttributeGoTemplate)
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

func (g GeneratorListNestedAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	f := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.ListValueType,
	}

	if g.CustomType != nil {
		f.ValueType = g.CustomType.ValueType
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
	// CustomTypeAndValue interface.
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
