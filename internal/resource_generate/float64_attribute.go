// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorFloat64Attribute struct {
	schema.Float64Attribute

	AssociatedExternalType *generatorschema.AssocExtType
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.Float64Default
	PlanModifiers specschema.Float64PlanModifiers
	Validators    specschema.Float64Validators
}

func (g GeneratorFloat64Attribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorFloat64Attribute
}

func (g GeneratorFloat64Attribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	if g.Default != nil {
		if g.Default.Static != nil {
			imports.Add(code.Import{
				Path: defaultFloat64Import,
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

func (g GeneratorFloat64Attribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorFloat64Attribute)
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

	return g.Float64Attribute.Equal(h.Float64Attribute)
}

func float64Default(d *specschema.Float64Default) string {
	if d == nil {
		return ""
	}

	if d.Static != nil {
		return fmt.Sprintf("float64default.StaticFloat64(%s)", strconv.FormatFloat(*d.Static, 'f', -1, 64))
	}

	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorFloat64Attribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                      string
		CustomType                string
		Default                   string
		GeneratorFloat64Attribute GeneratorFloat64Attribute
	}

	a := attribute{
		Name:                      name.ToString(),
		Default:                   float64Default(g.Default),
		GeneratorFloat64Attribute: g,
	}

	switch {
	case g.CustomType != nil:
		a.CustomType = g.CustomType.Type
	case g.AssociatedExternalType != nil:
		a.CustomType = fmt.Sprintf("%sType{}", name.ToPascalCase())
	}

	t, err := template.New("float64_attribute").Parse(float64AttributeTemplate)
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

func (g GeneratorFloat64Attribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.Float64ValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return field, nil
}

func (g GeneratorFloat64Attribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	float64Type := generatorschema.NewCustomFloat64Type(name)

	b, err := float64Type.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	float64Value := generatorschema.NewCustomFloat64Value(name)

	b, err = float64Value.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorFloat64Attribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	toFrom := generatorschema.NewToFromFloat64(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.Float64Typable type.
func (g GeneratorFloat64Attribute) AttrType(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase())
	}

	return "basetypes.Float64Type{}"
}

// AttrValue returns a string representation of a basetypes.Float64Valuable type.
func (g GeneratorFloat64Attribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.Float64Value"
}

func (g GeneratorFloat64Attribute) To() generatorschema.ToFromConversion {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}
	}

	return generatorschema.ToFromConversion{
		Default: "ValueFloat64Pointer",
	}
}

func (g GeneratorFloat64Attribute) From() generatorschema.ToFromConversion {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}
	}

	return generatorschema.ToFromConversion{
		Default: "Float64PointerValue",
	}
}
