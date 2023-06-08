// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorSetAttribute struct {
	schema.SetAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.SetDefault
	ElementType   specschema.ElementType
	PlanModifiers []specschema.SetPlanModifier
	Validators    []specschema.SetValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in.  Further
// imports are retrieved by calling getElementTypeImports.
func (g GeneratorSetAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports[*g.CustomType.Import] = struct{}{}
		}
	}

	elemTypeImports := getElementTypeImports(g.ElementType, make(map[string]struct{}))

	for k := range elemTypeImports {
		imports[k] = struct{}{}
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

		imports[validatorImport] = struct{}{}
		imports[*v.Custom.Import] = struct{}{}
	}

	return imports
}

func (g GeneratorSetAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorSetAttribute)
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

	if g.Computed != h.Computed {
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

func getSetDefault(d specschema.SetDefault) string {
	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorSetAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": getElementType,
		"getSetDefault":  getSetDefault,
	}

	t, err := template.New("set_attribute").Funcs(funcMap).Parse(setAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSetAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSetAttribute) validatorsEqual(x, y []specschema.SetValidator) bool {
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
