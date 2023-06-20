// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorMapNestedAttribute struct {
	schema.MapNestedAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.MapDefault
	NestedObject  GeneratorNestedAttributeObject
	PlanModifiers []specschema.MapPlanModifier
	Validators    []specschema.MapValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in.  The same
// logic is applied to the NestedObject. Further imports are then retrieved by
// calling Imports on each of the nested attributes.
func (g GeneratorMapNestedAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports[*g.CustomType.Import] = struct{}{}
		}
	}

	if g.Default != nil {
		if g.Default.Custom != nil && g.Default.Custom.HasImport() {
			imports[*g.Default.Custom.Import] = struct{}{}
		}
	}

	for _, v := range g.PlanModifiers {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		imports[planModifierImport] = struct{}{}
		imports[*v.Custom.Import] = struct{}{}
	}

	for _, v := range g.Validators {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		imports[generatorschema.ValidatorImport] = struct{}{}
		imports[*v.Custom.Import] = struct{}{}
	}

	if g.NestedObject.CustomType != nil {
		if g.NestedObject.CustomType.HasImport() {
			imports[*g.NestedObject.CustomType.Import] = struct{}{}
		}
	}

	for _, v := range g.NestedObject.PlanModifiers {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		imports[planModifierImport] = struct{}{}
		imports[*v.Custom.Import] = struct{}{}
	}

	for _, v := range g.NestedObject.Validators {
		if v.Custom == nil {
			continue
		}

		if !v.Custom.HasImport() {
			continue
		}

		imports[generatorschema.ValidatorImport] = struct{}{}
		imports[*v.Custom.Import] = struct{}{}
	}

	for _, v := range g.NestedObject.Attributes {
		for k := range v.Imports() {
			imports[k] = struct{}{}
		}
	}

	return imports
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

	if _, err = addCommonAttributeTemplate(t); err != nil {
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
