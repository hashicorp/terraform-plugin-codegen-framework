// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

const (
	PlanModifierTypeBool    PlanModifierType = "Bool"
	PlanModifierTypeFloat64 PlanModifierType = "Float64"
	PlanModifierTypeInt64   PlanModifierType = "Int64"
	PlanModifierTypeList    PlanModifierType = "List"
	PlanModifierTypeMap     PlanModifierType = "Map"
	PlanModifierTypeNumber  PlanModifierType = "Number"
	PlanModifierTypeObject  PlanModifierType = "Object"
	PlanModifierTypeSet     PlanModifierType = "Set"
	PlanModifierTypeString  PlanModifierType = "String"
)

type PlanModifierType string

type PlanModifiers struct {
	planModifierType PlanModifierType
	custom           specschema.CustomPlanModifiers
}

func NewPlanModifiers(t PlanModifierType, c specschema.CustomPlanModifiers) PlanModifiers {
	return PlanModifiers{
		planModifierType: t,
		custom:           c,
	}
}

func (v PlanModifiers) Equal(other PlanModifiers) bool {
	if v.planModifierType != other.planModifierType {
		return false
	}

	if len(v.custom) == 0 && len(other.custom) == 0 {
		return true
	}

	if len(v.custom) != len(other.custom) {
		return false
	}

	v.custom.Sort()

	other.custom.Sort()

	for i := 0; i < len(v.custom); i++ {
		if !v.custom[i].Equal(other.custom[i]) {
			return false
		}
	}

	return true
}

func (v PlanModifiers) Imports() *schema.Imports {
	imports := schema.NewImports()

	if v.custom == nil {
		return imports
	}

	for _, c := range v.custom {
		for _, i := range c.Imports {
			if len(i.Path) > 0 {
				imports.Add(code.Import{
					Path: schema.PlanModifierImport,
				})

				imports.Add(i)
			}
		}
	}

	return imports
}

func (v PlanModifiers) Schema() []byte {
	var b, cb bytes.Buffer

	for _, c := range v.custom {
		if c == nil {
			continue
		}

		if c.SchemaDefinition == "" {
			continue
		}

		cb.WriteString(fmt.Sprintf("%s,\n", c.SchemaDefinition))
	}

	if cb.Len() > 0 {
		b.WriteString(fmt.Sprintf("PlanModifiers: []planmodifier.%s{\n", v.planModifierType))
		b.Write(cb.Bytes())
		b.WriteString("},\n")
	}

	return b.Bytes()
}
