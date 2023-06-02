// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GeneratorListAttribute struct {
	schema.ListAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.ListDefault
	PlanModifiers []specschema.ListPlanModifier
	Validators    []specschema.ListValidator
}

func (g GeneratorListAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.ListAttribute.Equal(h.ListAttribute)
}

func getListDefault(d specschema.ListDefault) string {
	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorListAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": getElementType,
		"getListDefault": getListDefault,
	}

	t, err := template.New("list_attribute").Funcs(funcMap).Parse(listAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
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

func getElementType(elementType attr.Type) string {
	switch t := elementType.(type) {
	case basetypes.BoolType:
		return "types.BoolType"
	case basetypes.Float64Type:
		return "types.Float64Type"
	case basetypes.Int64Type:
		return "types.Int64Type"
	case types.ListType:
		return fmt.Sprintf("types.ListType{\nElemType: %s,\n}", getElementType(t.ElementType()))
	case types.MapType:
		return fmt.Sprintf("types.MapType{\nElemType: %s,\n}", getElementType(t.ElementType()))
	case basetypes.NumberType:
		return "types.NumberType"
	case types.ObjectType:
		return fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", getAttrTypes(t.AttrTypes))
	case types.SetType:
		return fmt.Sprintf("types.SetType{\nElemType: %s,\n}", getElementType(t.ElementType()))
	case basetypes.StringType:
		return "types.StringType"
	}

	return ""
}
