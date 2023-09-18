// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorObjectAttribute struct {
	schema.ObjectAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	AttributeTypes specschema.ObjectAttributeTypes
	CustomType     *specschema.CustomType
	Validators     specschema.ObjectValidators
}

func (g GeneratorObjectAttribute) AttrType() attr.Type {
	return types.ObjectType{
		//TODO: Add AttrTypes?
	}
}

func (g GeneratorObjectAttribute) AttrTypes() specschema.ObjectAttributeTypes {
	return g.AttributeTypes
}

func (g GeneratorObjectAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	attrTypesImports := generatorschema.GetAttrTypesImports(g.CustomType, g.AttributeTypes)
	imports.Append(attrTypesImports)

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	return imports
}

func (g GeneratorObjectAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorObjectAttribute)

	if !ok {
		return false
	}

	if !g.AttributeTypes.Equal(h.AttributeTypes) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.ObjectAttribute.Equal(h.ObjectAttribute)
}

func (g GeneratorObjectAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttrTypes": generatorschema.GetAttrTypes,
	}

	t, err := template.New("object_attribute").Funcs(funcMap).Parse(objectAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorObjectAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorObjectAttribute) ModelField(name string) (model.Field, error) {
	field := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.ObjectValueType,
	}

	if g.CustomType != nil {
		field.ValueType = g.CustomType.ValueType
	}

	return field, nil
}
