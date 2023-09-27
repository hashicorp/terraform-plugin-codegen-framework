// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSetNestedAttribute struct {
	schema.SetNestedAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType   *specschema.CustomType
	NestedObject GeneratorNestedAttributeObject
	Validators   specschema.SetValidators
}

func (g GeneratorSetNestedAttribute) AssocExtType() *generatorschema.AssocExtType {
	return g.NestedObject.AssociatedExternalType
}

func (g GeneratorSetNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSetNestedAttribute
}

func (g GeneratorSetNestedAttribute) Imports() *generatorschema.Imports {
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

func (g GeneratorSetNestedAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorSetNestedAttribute)

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

	return g.SetNestedAttribute.Equal(h.SetNestedAttribute)
}

func (g GeneratorSetNestedAttribute) Schema(name string) (string, error) {
	type attribute struct {
		Name                        string
		TypeValueName               string
		Attributes                  string
		GeneratorSetNestedAttribute GeneratorSetNestedAttribute
	}

	attributesStr, err := g.NestedObject.Attributes.Schema()

	if err != nil {
		return "", err
	}

	a := attribute{
		Name:                        name,
		TypeValueName:               model.SnakeCaseToCamelCase(name),
		Attributes:                  attributesStr,
		GeneratorSetNestedAttribute: g,
	}

	t, err := template.New("set_nested_attribute").Parse(setNestedAttributeGoTemplate)
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

func (g GeneratorSetNestedAttribute) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.SetValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorSetNestedAttribute) GetAttributes() generatorschema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

func (g GeneratorSetNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.NestedObject.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	objectType := generatorschema.NewCustomObjectType(name, attributeAttrValues)

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

	objectValue := generatorschema.NewCustomObjectValue(name, attributeTypes, attributeAttrTypes, attributeAttrValues)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	// TODO: Remove once refactored to Generator<Type>Attribute|Block
	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.NestedObject.Attributes))

	for k := range g.NestedObject.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	// Recursively call CustomTypeAndValue() for each attribute that implements
	// CustomTypeAndValue interface (i.e, nested attributes).
	for _, k := range attributeKeys {
		// TODO: Also need to consider how to handle instances in which an associated_external_type
		// has been defined on a type which does not implement CustomTypeAndValue (e.g., bool)
		// If To/From methods are going to be hung off custom value type, then will to generate
		// "wrapped" / embedded types that embed bool in a type that can have To/From methods
		// added to it.
		if c, ok := g.NestedObject.Attributes[k].(generatorschema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)

			continue
		}

		// TODO: Remove once refactored to Generator<Type>Attribute|Block
		if a, ok := g.NestedObject.Attributes[k].(generatorschema.Attributes); ok {
			ng := generatorschema.GeneratorSchema{
				Attributes: a.GetAttributes(),
			}

			b, err := ng.ModelObjectHelpersTemplate(k)
			if err != nil {
				return nil, err
			}

			buf.WriteString("\n")
			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}
