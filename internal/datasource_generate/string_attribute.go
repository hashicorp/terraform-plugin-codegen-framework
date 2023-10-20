// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorStringAttribute struct {
	schema.StringAttribute

	AssociatedExternalType *generatorschema.AssocExtType
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType *specschema.CustomType
	Validators specschema.StringValidators
}

func (g GeneratorStringAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorStringAttribute
}

func (g GeneratorStringAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	if g.AssociatedExternalType != nil {
		imports.Append(generatorschema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorStringAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorStringAttribute)
	if !ok {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.StringAttribute.Equal(h.StringAttribute)
}

func (g GeneratorStringAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                     string
		CustomType               string
		GeneratorStringAttribute GeneratorStringAttribute
	}

	a := attribute{
		Name:                     name.ToString(),
		GeneratorStringAttribute: g,
	}

	switch {
	case g.CustomType != nil:
		a.CustomType = g.CustomType.Type
	case g.AssociatedExternalType != nil:
		a.CustomType = fmt.Sprintf("%sType{}", name.ToPascalCase())
	}

	t, err := template.New("string_attribute").Parse(stringAttributeTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, a)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorStringAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.StringValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return field, nil
}

func (g GeneratorStringAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	stringType := generatorschema.NewCustomStringType(name)

	b, err := stringType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	stringValue := generatorschema.NewCustomStringValue(name)

	b, err = stringValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorStringAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	toFrom := generatorschema.NewToFromString(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.StringTypable type.
func (g GeneratorStringAttribute) AttrType(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase())
	}

	return "basetypes.StringType{}"
}

// AttrValue returns a string representation of a basetypes.StringValuable type.
func (g GeneratorStringAttribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.StringValue"
}

func (g GeneratorStringAttribute) To() generatorschema.ToFromConversion {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}
	}

	return generatorschema.ToFromConversion{
		Default: "ValueStringPointer",
	}
}

func (g GeneratorStringAttribute) From() generatorschema.ToFromConversion {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}
	}

	return generatorschema.ToFromConversion{
		Default: "StringPointerValue",
	}
}
