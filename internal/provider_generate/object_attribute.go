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

type GeneratorObjectAttribute struct {
	schema.ObjectAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	AttributeTypes []specschema.ObjectAttributeType
	CustomType     *specschema.CustomType
	Validators     []specschema.ObjectValidator
}

func (g GeneratorObjectAttribute) Imports() *generatorschema.Imports {
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

	attrTypesImports := generatorschema.GetAttrTypesImports(g.CustomType, g.AttributeTypes)

	imports.Add(attrTypesImports.All()...)

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

func (g GeneratorObjectAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorObjectAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.ObjectAttribute.Equal(h.ObjectAttribute)
}

func (g GeneratorObjectAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttrTypes": generatorschema.GetAttrTypes,
	}

	t, err := template.New("object_attribute").Funcs(funcMap).Parse(objectAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorObjectAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorObjectAttribute) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.ObjectValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorObjectAttribute) validatorsEqual(x, y []specschema.ObjectValidator) bool {
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
