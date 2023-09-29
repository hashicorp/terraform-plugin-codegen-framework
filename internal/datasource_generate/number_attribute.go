// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorNumberAttribute struct {
	schema.NumberAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType *specschema.CustomType
	Validators specschema.NumberValidators
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

func (g GeneratorNumberAttribute) Schema(name string) (string, error) {
	type attribute struct {
		Name                     string
		GeneratorNumberAttribute GeneratorNumberAttribute
	}

	a := attribute{
		Name:                     name,
		GeneratorNumberAttribute: g,
	}

	t, err := template.New("number_attribute").Parse(numberAttributeGoTemplate)
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

func (g GeneratorNumberAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.NumberValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}
