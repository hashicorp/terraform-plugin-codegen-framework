// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorFloat64Attribute struct {
	schema.Float64Attribute

	CustomType    *specschema.CustomType
	Default       *specschema.Float64Default
	PlanModifiers []specschema.Float64PlanModifier
	Validators    []specschema.Float64Validator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the schemaImport will be used.
func (g GeneratorFloat64Attribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports[*g.CustomType.Import] = struct{}{}
		}
	} else {
		imports[schemaImport] = struct{}{}
	}

	if g.Default != nil {
		if g.Default.Static != nil {
			imports[defaultFloat64Import] = struct{}{}
		} else if g.Default.Custom != nil && g.Default.Custom.HasImport() {
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

func (g GeneratorFloat64Attribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorFloat64Attribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.Float64Attribute.Equal(h.Float64Attribute)
}

func getFloat64Default(float64Default specschema.Float64Default) string {
	if float64Default.Static != nil {
		return fmt.Sprintf("float64default.StaticFloat64(%s)", strconv.FormatFloat(*float64Default.Static, 'f', -1, 64))
	}

	if float64Default.Custom != nil {
		return float64Default.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorFloat64Attribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getFloat64Default": getFloat64Default,
	}

	t, err := template.New("float64_attribute").Funcs(funcMap).Parse(float64AttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorFloat64Attribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorFloat64Attribute) validatorsEqual(x, y []specschema.Float64Validator) bool {
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
