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

type GeneratorBoolAttribute struct {
	schema.BoolAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.BoolDefault
	PlanModifiers []specschema.BoolPlanModifier
	Validators    []specschema.BoolValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in.
func (g GeneratorBoolAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports[g.CustomType.Import.Path] = struct{}{}
		}
	} else {
		imports[generatorschema.TypesImport] = struct{}{}
	}

	if g.Default != nil {
		if g.Default.Static != nil {
			imports[defaultBoolImport] = struct{}{}
		} else if g.Default.Custom != nil && g.Default.Custom.HasImport() {
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

	if _, err = addCommonAttributeTemplate(t); err != nil {
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

func (g GeneratorBoolAttribute) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.BoolValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
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
		if !customValidatorsEqual(v.Custom, y[k].Custom) {
			return false
		}
	}

	return true
}

func customValidatorsEqual(x, y *specschema.CustomValidator) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil || y == nil {
		return false
	}

	if len(x.Imports) != len(y.Imports) {
		return false
	}

	//TODO: Sort before comparing.
	for k, v := range x.Imports {
		if v.Path != y.Imports[k].Path {
			return false
		}

		if v.Alias != nil && y.Imports[k].Alias == nil {
			return false
		}

		if v.Alias == nil && y.Imports[k].Alias != nil {
			return false
		}

		if v.Alias != nil && y.Imports[k].Alias != nil {
			if *v.Alias != *y.Imports[k].Alias {
				return false
			}
		}
	}

	return x.SchemaDefinition == y.SchemaDefinition
}
