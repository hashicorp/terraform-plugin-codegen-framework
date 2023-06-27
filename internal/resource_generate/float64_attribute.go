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

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorFloat64Attribute struct {
	schema.Float64Attribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.Float64Default
	PlanModifiers []specschema.Float64PlanModifier
	Validators    []specschema.Float64Validator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in.
func (g GeneratorFloat64Attribute) Imports() map[string]struct{} {
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
			imports[defaultFloat64Import] = struct{}{}
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

func (g GeneratorFloat64Attribute) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.Float64ValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
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
		if !customValidatorsEqual(v.Custom, y[k].Custom) {
			return false
		}
	}

	return true
}
