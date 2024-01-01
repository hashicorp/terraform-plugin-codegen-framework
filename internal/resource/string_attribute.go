// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorStringAttribute struct {
	AssociatedExternalType   *generatorschema.AssocExtType
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomType               *specschema.CustomType
	CustomTypePrimitive      convert.CustomTypePrimitive
	Default                  *specschema.StringDefault
	DefaultString            convert.DefaultString
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	PlanModifiers            specschema.StringPlanModifiers
	PlanModifiersCustom      convert.PlanModifiersCustom
	Sensitive                convert.Sensitive
	Validators               specschema.StringValidators
	ValidatorsCustom         convert.ValidatorsCustom
}

func NewGeneratorStringAttribute(name string, a *resource.StringAttribute) (GeneratorStringAttribute, error) {
	if a == nil {
		return GeneratorStringAttribute{}, fmt.Errorf("*resource.StringAttribute is nil")
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ctp := convert.NewCustomTypePrimitive(a.CustomType, a.AssociatedExternalType, name)

	db := convert.NewDefaultString(a.Default)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	d := convert.NewDescription(a.Description)

	pm := convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, a.PlanModifiers.CustomPlanModifiers())

	s := convert.NewSensitive(a.Sensitive)

	vc := convert.NewValidatorsCustom(convert.ValidatorTypeString, a.Validators.CustomValidators())

	return GeneratorStringAttribute{
		AssociatedExternalType:   generatorschema.NewAssocExtType(a.AssociatedExternalType),
		ComputedOptionalRequired: c,
		CustomType:               a.CustomType,
		CustomTypePrimitive:      ctp,
		Default:                  a.Default,
		DefaultString:            db,
		Description:              d,
		DeprecationMessage:       dm,
		PlanModifiers:            a.PlanModifiers,
		PlanModifiersCustom:      pm,
		Sensitive:                s,
		Validators:               a.Validators,
		ValidatorsCustom:         vc,
	}, nil
}

func (g GeneratorStringAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorStringAttribute
}

func (g GeneratorStringAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	if g.Default != nil {
		if g.Default.Static != nil {
			imports.Add(code.Import{
				Path: defaultStringImport,
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

func (g GeneratorStringAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorStringAttribute)

	if !ok {
		return false
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.CustomTypePrimitive.Equal(h.CustomTypePrimitive) {
		return false
	}

	if !g.Default.Equal(h.Default) {
		return false
	}

	if !g.DefaultString.Equal(h.DefaultString) {
		return false
	}

	if !g.DeprecationMessage.Equal(h.DeprecationMessage) {
		return false
	}

	if !g.Description.Equal(h.Description) {
		return false
	}

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.PlanModifiersCustom.Equal(h.PlanModifiersCustom) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.ValidatorsCustom.Equal(h.ValidatorsCustom)
}

func (g GeneratorStringAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.StringAttribute{\n", name))
	b.Write(g.CustomTypePrimitive.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.PlanModifiersCustom.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.Write(g.DefaultString.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorStringAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.StringValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return field, nil
}

func (g GeneratorStringAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	stringType := generatorschema.NewCustomStringType(name)

	b, err := stringType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	stringValue := generatorschema.NewCustomStringValue(name)

	b, err = stringValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorStringAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	toFrom := generatorschema.NewToFromString(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.StringTypable type.
func (g GeneratorStringAttribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase()), nil
	}

	return "basetypes.StringType{}", nil
}

// AttrValue returns a string representation of a basetypes.StringValuable type.
func (g GeneratorStringAttribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.StringValue"
}

func (g GeneratorStringAttribute) To() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "ValueStringPointer",
	}, nil
}

func (g GeneratorStringAttribute) From() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "StringPointerValue",
	}, nil
}
