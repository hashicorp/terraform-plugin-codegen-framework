// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorListNestedAttribute struct {
	schema.ListNestedAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType   *specschema.CustomType
	NestedObject GeneratorNestedAttributeObject
	Validators   []specschema.ListValidator
}

func (g GeneratorListNestedAttribute) GeneratorAttrType() (GeneratorAttrType, error) {
	attrTypes := make(map[string]attr.Type)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var keys = make([]string, 0, len(g.NestedObject.Attributes))

	for k := range g.NestedObject.Attributes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		generatorAttrType, err := g.NestedObject.Attributes[k].GeneratorAttrType()
		if err != nil {
			return GeneratorAttrType{}, err
		}

		attrTypes[k] = generatorAttrType
	}

	return GeneratorAttrType{
		Type: types.ListType{
			ElemType: GeneratorAttrType{
				Type: types.ObjectType{
					AttrTypes: attrTypes,
				},
			},
		},
	}, nil
}

func (g GeneratorListNestedAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	customTypeImports = generatorschema.CustomTypeImports(g.NestedObject.CustomType)
	imports.Append(customTypeImports)

	for _, v := range g.NestedObject.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	for _, v := range g.NestedObject.Attributes {
		imports.Append(v.Imports())
	}

	return imports
}

func (g GeneratorListNestedAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListNestedAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.listValidatorsEqual(g.Validators, h.Validators) {
		return false
	}

	if !customTypeEqual(g.NestedObject.CustomType, h.NestedObject.CustomType) {
		return false
	}

	if !g.objectValidatorsEqual(g.NestedObject.Validators, h.NestedObject.Validators) {
		return false
	}

	for k, a := range g.NestedObject.Attributes {
		if !a.Equal(h.NestedObject.Attributes[k]) {
			return false
		}
	}

	return true
}

func (g GeneratorListNestedAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"AttributesString": g.NestedObject.Attributes.String,
	}

	t, err := template.New("list_nested_attribute").Funcs(funcMap).Parse(listNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorListNestedAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorListNestedAttribute) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.ListValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}

func (g GeneratorListNestedAttribute) ModelObjectHelpersString(name string) (string, error) {
	attrTypeStrings := make(map[string]string)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var keys = make([]string, 0, len(g.NestedObject.Attributes))

	for k := range g.NestedObject.Attributes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Populate attrTypeStrings map for use in template.
	// TODO: Add in remaining attribute types.
	for _, k := range keys {
		switch t := g.NestedObject.Attributes[k].(type) {
		case GeneratorBoolAttribute:
			attrTypeStrings[k] = "types.BoolType"
		case GeneratorListAttribute:
			var elemType string

			switch {
			case t.ElementType.String != nil:
				elemType = "types.StringType"
			}

			attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %s,\n}", elemType)
		case GeneratorListNestedAttribute:
			attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %sModel{}.objectType(),\n}", model.SnakeCaseToCamelCase(k))
		}
	}

	t, err := template.New("list_nested_attribute").Parse(listNestedAttributeModelObjectHelpers)
	if err != nil {
		return "", err
	}

	var buf strings.Builder

	templateData := struct {
		Name      string
		AttrTypes map[string]string
	}{
		Name:      model.SnakeCaseToCamelCase(name),
		AttrTypes: attrTypeStrings,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return "", err
	}

	// Recursively call ModelObjectHelpersString() for each attribute that is a nested attribute or nested block.
	for _, k := range keys {
		switch t := g.NestedObject.Attributes[k].(type) {
		case GeneratorListNestedAttribute:
			str, err := t.ModelObjectHelpersString(k)
			if err != nil {
				return "", err
			}

			buf.WriteString("\n" + str)
		}
	}

	return buf.String(), nil
}

func (g GeneratorListNestedAttribute) listValidatorsEqual(x, y []specschema.ListValidator) bool {
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

func (g GeneratorListNestedAttribute) objectValidatorsEqual(x, y []specschema.ObjectValidator) bool {
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
