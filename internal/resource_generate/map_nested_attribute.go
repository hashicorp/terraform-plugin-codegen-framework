// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorMapNestedAttribute struct {
	schema.MapNestedAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.MapDefault
	NestedObject  GeneratorNestedAttributeObject
	PlanModifiers []specschema.MapPlanModifier
	Validators    []specschema.MapValidator
}

func (g GeneratorMapNestedAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorMapNestedAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.mapValidatorsEqual(g.Validators, h.Validators) {
		return false
	}

	if !customTypeEqual(g.NestedObject.CustomType, h.NestedObject.CustomType) {
		return false
	}

	if !g.objectValidatorsEqual(g.NestedObject.Validators, h.NestedObject.Validators) {
		return false
	}

	for k, a := range g.NestedObject.Attributes {
		if !a.Equal(h.NestedObject.Attributes[k]) {
			return false
		}
	}

	return true
}

func (g GeneratorMapNestedAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": getAttributes,
		"getMapDefault": getMapDefault,
	}

	t, err := template.New("map_nested_attribute").Funcs(funcMap).Parse(mapNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorMapNestedAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorMapNestedAttribute) mapValidatorsEqual(x, y []specschema.MapValidator) bool {
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
		if v.Custom == nil && y[k].Custom != nil {
			return false
		}

		if v.Custom != nil && y[k].Custom == nil {
			return false
		}

		if v.Custom != nil && y[k].Custom != nil {
			if *v.Custom.Import != *y[k].Custom.Import {
				return false
			}
		}

		if v.Custom.SchemaDefinition != y[k].Custom.SchemaDefinition {
			return false
		}
	}

	return true
}

func (g GeneratorMapNestedAttribute) objectValidatorsEqual(x, y []specschema.ObjectValidator) bool {
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
		if v.Custom == nil && y[k].Custom != nil {
			return false
		}

		if v.Custom != nil && y[k].Custom == nil {
			return false
		}

		if v.Custom != nil && y[k].Custom != nil {
			if *v.Custom.Import != *y[k].Custom.Import {
				return false
			}
		}

		if v.Custom.SchemaDefinition != y[k].Custom.SchemaDefinition {
			return false
		}
	}

	return true
}
