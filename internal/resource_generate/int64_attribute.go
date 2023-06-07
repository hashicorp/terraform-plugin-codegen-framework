// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorInt64Attribute struct {
	schema.Int64Attribute

	CustomType    *specschema.CustomType
	Default       *specschema.Int64Default
	PlanModifiers []specschema.Int64PlanModifier
	Validators    []specschema.Int64Validator
}

func (g GeneratorInt64Attribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorInt64Attribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.Int64Attribute.Equal(h.Int64Attribute)
}

func getInt64Default(d specschema.Int64Default) string {
	if d.Static != nil {
		return fmt.Sprintf("int64default.StaticInt64(%d)", *d.Static)
	}

	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorInt64Attribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getInt64Default": getInt64Default,
	}

	t, err := template.New("int64_attribute").Funcs(funcMap).Parse(int64AttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorInt64Attribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorInt64Attribute) validatorsEqual(x, y []specschema.Int64Validator) bool {
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
