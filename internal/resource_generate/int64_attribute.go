// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorInt64Attribute struct {
	schema.Int64Attribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.Int64Default
	PlanModifiers []specschema.Int64PlanModifier
	Validators    []specschema.Int64Validator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in.
func (g GeneratorInt64Attribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports[*g.CustomType.Import] = struct{}{}
		}
	} else {
		imports[generatorschema.TypesImport] = struct{}{}
	}

	if g.Default != nil {
		if g.Default.Static != nil {
			imports[defaultInt64Import] = struct{}{}
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

		imports[generatorschema.ValidatorImport] = struct{}{}
		imports[*v.Custom.Import] = struct{}{}
	}

	return imports
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

func (g GeneratorInt64Attribute) ToModel(name string) (string, error) {
	field := model.StructField{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.Int64ValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field.String(), nil
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
