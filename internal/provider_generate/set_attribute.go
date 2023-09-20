// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSetAttribute struct {
	schema.SetAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType  *specschema.CustomType
	ElementType specschema.ElementType
	Validators  specschema.SetValidators
}

func (g GeneratorSetAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorSetAttribute
}

func (g GeneratorSetAttribute) ElemType() specschema.ElementType {
	return g.ElementType
}

func (g GeneratorSetAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	elemTypeImports := generatorschema.GetElementTypeImports(g.ElementType)
	imports.Append(elemTypeImports)

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	return imports
}

// Equal does not delegate to g.SetAttribute.Equal(h.SetAttribute) as the
// call returns false when the ElementType is nil.
func (g GeneratorSetAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorSetAttribute)
	if !ok {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.ElementType.Equal(h.ElementType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	if g.Required != h.Required {
		return false
	}

	if g.Optional != h.Optional {
		return false
	}

	if g.Sensitive != h.Sensitive {
		return false
	}

	if g.Description != h.Description {
		return false
	}

	if g.MarkdownDescription != h.MarkdownDescription {
		return false
	}

	if g.DeprecationMessage != h.DeprecationMessage {
		return false
	}

	return true
}

func (g GeneratorSetAttribute) Schema(name string) (string, error) {
	type attribute struct {
		Name                  string
		ElementType           string
		GeneratorSetAttribute GeneratorSetAttribute
	}

	a := attribute{
		Name:                  name,
		ElementType:           generatorschema.GetElementType(g.ElementType),
		GeneratorSetAttribute: g,
	}

	t, err := template.New("set_attribute").Parse(setAttributeGoTemplate)
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

func (g GeneratorSetAttribute) ModelField(name string) (model.Field, error) {
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
