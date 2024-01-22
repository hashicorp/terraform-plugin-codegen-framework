// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorBoolAttribute struct {
	AssociatedExternalType   *generatorschema.AssocExtType
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomTypePrimitive      convert.CustomTypePrimitive
	DefaultBool              convert.DefaultBool
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	PlanModifiersCustom      convert.PlanModifiersCustom
	Sensitive                convert.Sensitive
	ValidatorsCustom         convert.ValidatorsCustom
}

func NewGeneratorBoolAttribute(name string, a *resource.BoolAttribute) (GeneratorBoolAttribute, error) {
	if a == nil {
		return GeneratorBoolAttribute{}, fmt.Errorf("*resource.BoolAttribute is nil")
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ctp := convert.NewCustomTypePrimitive(a.CustomType, a.AssociatedExternalType, name)

	db := convert.NewDefaultBool(a.Default)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	d := convert.NewDescription(a.Description)

	pm := convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, a.PlanModifiers.CustomPlanModifiers())

	s := convert.NewSensitive(a.Sensitive)

	vc := convert.NewValidatorsCustom(convert.ValidatorTypeBool, a.Validators.CustomValidators())

	return GeneratorBoolAttribute{
		AssociatedExternalType:   generatorschema.NewAssocExtType(a.AssociatedExternalType),
		ComputedOptionalRequired: c,
		CustomTypePrimitive:      ctp,
		DefaultBool:              db,
		Description:              d,
		DeprecationMessage:       dm,
		PlanModifiersCustom:      pm,
		Sensitive:                s,
		ValidatorsCustom:         vc,
	}, nil
}

func (g GeneratorBoolAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorBoolAttribute
}

func (g GeneratorBoolAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(g.CustomTypePrimitive.Imports())

	imports.Append(g.DefaultBool.Imports())

	imports.Append(g.PlanModifiersCustom.Imports())

	imports.Append(g.ValidatorsCustom.Imports())

	if g.AssociatedExternalType != nil {
		imports.Append(generatorschema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorBoolAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorBoolAttribute)

	if !ok {
		return false
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
		return false
	}

	if !g.CustomTypePrimitive.Equal(h.CustomTypePrimitive) {
		return false
	}

	if !g.DefaultBool.Equal(h.DefaultBool) {
		return false
	}

	if !g.DeprecationMessage.Equal(h.DeprecationMessage) {
		return false
	}

	if !g.Description.Equal(h.Description) {
		return false
	}

	if !g.PlanModifiersCustom.Equal(h.PlanModifiersCustom) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.ValidatorsCustom.Equal(h.ValidatorsCustom)
}

func (g GeneratorBoolAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.BoolAttribute{\n", name))
	b.Write(g.CustomTypePrimitive.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.PlanModifiersCustom.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.Write(g.DefaultBool.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorBoolAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.BoolValueType,
	}

	customValueType := g.CustomTypePrimitive.ValueType()

	if customValueType != "" {
		field.ValueType = customValueType
	}

	return field, nil
}

func (g GeneratorBoolAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	boolType := generatorschema.NewCustomBoolType(name)

	b, err := boolType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	boolValue := generatorschema.NewCustomBoolValue(name)

	b, err = boolValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorBoolAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	toFrom := generatorschema.NewToFromBool(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.BoolTypable type.
func (g GeneratorBoolAttribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase()), nil
	}

	return "basetypes.BoolType{}", nil
}

// AttrValue returns a string representation of a basetypes.BoolValuable type.
func (g GeneratorBoolAttribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.BoolValue"
}

func (g GeneratorBoolAttribute) To() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "ValueBoolPointer",
	}, nil
}

func (g GeneratorBoolAttribute) From() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "BoolPointerValue",
	}, nil
}
