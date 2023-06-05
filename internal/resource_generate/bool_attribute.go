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

type GeneratorBoolAttribute struct {
	schema.BoolAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.BoolDefault
	PlanModifiers []specschema.BoolPlanModifier
	Validators    []specschema.BoolValidator
}

func (g GeneratorBoolAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorBoolAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.BoolAttribute.Equal(h.BoolAttribute)
}

func getBoolDefault(boolDefault specschema.BoolDefault) string {
	if boolDefault.Static != nil {
		return fmt.Sprintf("booldefault.StaticBool(%t)", *boolDefault.Static)
	}

	if boolDefault.Custom != nil {
		return boolDefault.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorBoolAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getBoolDefault": getBoolDefault,
	}

	t, err := template.New("bool_attribute").Funcs(funcMap).Parse(boolAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorBoolAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorBoolAttribute) validatorsEqual(x, y []specschema.BoolValidator) bool {
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
