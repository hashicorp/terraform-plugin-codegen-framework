// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorInt64Attribute struct {
	schema.Int64Attribute

	AssociatedExternalType *generatorschema.AssocExtType
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.Int64Default
	PlanModifiers specschema.Int64PlanModifiers
	Validators    specschema.Int64Validators
}

func (g GeneratorInt64Attribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorInt64Attribute
}

func (g GeneratorInt64Attribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	if g.Default != nil {
		if g.Default.Static != nil {
			imports.Add(code.Import{
				Path: defaultInt64Import,
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

	if g.AssociatedExternalType != nil {
		imports.Append(generatorschema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorInt64Attribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorInt64Attribute)
	if !ok {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Default.Equal(h.Default) {
		return false
	}

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.Int64Attribute.Equal(h.Int64Attribute)
}

func int64Default(d *specschema.Int64Default) string {
	if d == nil {
		return ""
	}

	if d.Static != nil {
		return fmt.Sprintf("int64default.StaticInt64(%d)", *d.Static)
	}

	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorInt64Attribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                    string
		CustomType              string
		Default                 string
		GeneratorInt64Attribute GeneratorInt64Attribute
	}

	a := attribute{
		Name:                    name.ToString(),
		Default:                 int64Default(g.Default),
		GeneratorInt64Attribute: g,
	}

	switch {
	case g.CustomType != nil:
		a.CustomType = g.CustomType.Type
	case g.AssociatedExternalType != nil:
		a.CustomType = fmt.Sprintf("%sType{}", name.ToPascalCase())
	}

	t, err := template.New("int64_attribute").Parse(int64AttributeTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, a)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorInt64Attribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.Int64ValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return field, nil
}

func (g GeneratorInt64Attribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	int64Type := generatorschema.NewCustomInt64Type(name)

	b, err := int64Type.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	int64Value := generatorschema.NewCustomInt64Value(name)

	b, err = int64Value.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorInt64Attribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	toFrom := generatorschema.NewToFromInt64(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.Int64Typable type.
func (g GeneratorInt64Attribute) AttrType(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase())
	}

	return "basetypes.Int64Type{}"
}

// AttrValue returns a string representation of a basetypes.Int64Valuable type.
func (g GeneratorInt64Attribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.Int64Value"
}

func (g GeneratorInt64Attribute) To() generatorschema.ToFromConversion {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}
	}

	return generatorschema.ToFromConversion{
		Default: "ValueInt64Pointer",
	}
}

func (g GeneratorInt64Attribute) From() generatorschema.ToFromConversion {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}
	}

	return generatorschema.ToFromConversion{
		Default: "Int64PointerValue",
	}
}
