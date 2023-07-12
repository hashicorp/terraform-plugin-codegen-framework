// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

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

func (g GeneratorBoolAttribute) AttrType() attr.Type {
	return types.BoolType
}

func (g GeneratorBoolAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	if g.Default != nil {
		if g.Default.Static != nil {
			imports.Add(code.Import{
				Path: defaultBoolImport,
			})
		} else {
			customDefaultImports := generatorschema.CustomDefaultImports(g.Default.Custom)
			imports.Append(customDefaultImports)
		}
	}

	for _, v := range g.PlanModifiers {
		customPlanModifierImports := generatorschema.CustomPlanModifierImports(v.Custom)
		imports.Append(customPlanModifierImports)
	}

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	return imports
}

func (g GeneratorBoolAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
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
