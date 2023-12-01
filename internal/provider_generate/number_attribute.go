// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorNumberAttribute struct {
	schema.NumberAttribute

	AssociatedExternalType *generatorschema.AssocExtType
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType *specschema.CustomType
	Validators specschema.NumberValidators
}

func NewGeneratorNumberAttribute(a *provider.NumberAttribute) (GeneratorNumberAttribute, error) {
	if a == nil {
		return GeneratorNumberAttribute{}, fmt.Errorf("*provider.NumberAttribute is nil")
	}

	c := convert.NewOptionalRequired(a.OptionalRequired)

	s := convert.NewSensitive(a.Sensitive)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	return GeneratorNumberAttribute{
		NumberAttribute: schema.NumberAttribute{
			Required:            c.IsRequired(),
			Optional:            c.IsOptional(),
			Sensitive:           s.IsSensitive(),
			Description:         d.Description(),
			MarkdownDescription: d.Description(),
			DeprecationMessage:  dm.DeprecationMessage(),
		},

		AssociatedExternalType: generatorschema.NewAssocExtType(a.AssociatedExternalType),
		CustomType:             a.CustomType,
		Validators:             a.Validators,
	}, nil
}

func (g GeneratorNumberAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorNumberAttribute
}

func (g GeneratorNumberAttribute) Imports() *generatorschema.Imports {
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

func (g GeneratorNumberAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorNumberAttribute)
	if !ok {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.NumberAttribute.Equal(h.NumberAttribute)
}

func (g GeneratorNumberAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                     string
		CustomType               string
		GeneratorNumberAttribute GeneratorNumberAttribute
	}

	a := attribute{
		Name:                     name.ToString(),
		GeneratorNumberAttribute: g,
	}

	switch {
	case g.CustomType != nil:
		a.CustomType = g.CustomType.Type
	case g.AssociatedExternalType != nil:
		a.CustomType = fmt.Sprintf("%sType{}", name.ToPascalCase())
	}

	t, err := template.New("number_attribute").Parse(numberAttributeTemplate)
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

func (g GeneratorNumberAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.NumberValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return field, nil
}

func (g GeneratorNumberAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	numberType := generatorschema.NewCustomNumberType(name)

	b, err := numberType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	numberValue := generatorschema.NewCustomNumberValue(name)

	b, err = numberValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorNumberAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	toFrom := generatorschema.NewToFromNumber(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.NumberTypable type.
func (g GeneratorNumberAttribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase()), nil
	}

	return "basetypes.NumberType{}", nil
}

// AttrValue returns a string representation of a basetypes.NumberValuable type.
func (g GeneratorNumberAttribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.NumberValue"
}

func (g GeneratorNumberAttribute) To() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "ValueBigFloat",
	}, nil
}

func (g GeneratorNumberAttribute) From() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "NumberValue",
	}, nil
}
