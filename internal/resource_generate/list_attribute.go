// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorListAttribute struct {
	schema.ListAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.ListDefault
	ElementType   specschema.ElementType
	PlanModifiers []specschema.ListPlanModifier
	Validators    []specschema.ListValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the datasourceSchemaImport will be used.
func (g GeneratorListAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports[g.CustomType.Import.Path] = struct{}{}
		}
	} else {
		imports[generatorschema.TypesImport] = struct{}{}
	}

	elemTypeImports := generatorschema.GetElementTypeImports(g.ElementType, make(map[string]struct{}))

	for k := range elemTypeImports {
		imports[k] = struct{}{}
	}

	if g.Default != nil {
		if g.Default.Custom != nil && g.Default.Custom.HasImport() {
			for _, i := range g.Default.Custom.Imports {
				if len(i.Path) > 0 {
					imports[i.Path] = struct{}{}
				}
			}
		}
	}

	for _, v := range g.PlanModifiers {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		for _, i := range v.Custom.Imports {
			if len(i.Path) > 0 {
				imports[planModifierImport] = struct{}{}
				imports[i.Path] = struct{}{}
			}
		}
	}

	for _, v := range g.Validators {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		for _, i := range v.Custom.Imports {
			if len(i.Path) > 0 {
				imports[generatorschema.ValidatorImport] = struct{}{}
				imports[i.Path] = struct{}{}
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

func getListDefault(d specschema.ListDefault) string {
	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorListAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": generatorschema.GetElementType,
		"getListDefault": getListDefault,
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
