// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorListAttribute struct {
	schema.ListAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType  *specschema.CustomType
	ElementType specschema.ElementType
	Validators  []specschema.ListValidator
}

func (g GeneratorListAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports.Add(*g.CustomType.Import)
		}
	} else {
		imports.Add(code.Import{
			Path: generatorschema.TypesImport,
		})
	}

	elemTypeImports := generatorschema.GetElementTypeImportsStruct(g.ElementType)

	imports.Add(elemTypeImports.All()...)

	for _, v := range g.Validators {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		for _, i := range v.Custom.Imports {
			if len(i.Path) > 0 {
				imports.Add(code.Import{
					Path: generatorschema.ValidatorImport,
				})

				imports.Add(i)
			}
		}
	}

	return imports
}

// Equal does not delegate to g.ListAttribute.Equal(h.ListAttribute) as the
// call returns false owing to !a.GetType().Equal(b.GetType()) returning false
// when the ElementType is nil.
func (g GeneratorListAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !elementTypeEqual(g.ElementType, h.ElementType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
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

func (g GeneratorListAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": generatorschema.GetElementType,
	}

	t, err := template.New("list_attribute").Funcs(funcMap).Parse(listAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorListAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorListAttribute) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.ListValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorListAttribute) validatorsEqual(x, y []specschema.ListValidator) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	//TODO: Sort before comparing.
	for k, v := range x {
		if !customValidatorsEqual(v.Custom, y[k].Custom) {
			return false
		}
	}

	return true
}
