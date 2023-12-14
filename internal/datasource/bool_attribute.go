// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorBoolAttribute struct {
	AttributeType            convert.AttributeType
	AssociatedExternalType   *generatorschema.AssocExtType
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomType               *specschema.CustomType
	Description              convert.Description
	DeprecationMessage       convert.DeprecationMessage
	Sensitive                convert.Sensitive
	Validators               specschema.BoolValidators
	ValidatorsCustom         convert.ValidatorsCustom
}

func NewGeneratorBoolAttribute(name string, a *datasource.BoolAttribute) (GeneratorBoolAttribute, error) {
	if a == nil {
		return GeneratorBoolAttribute{}, fmt.Errorf("*datasource.BoolAttribute is nil")
	}

	at := convert.NewAttributeType(a.CustomType, a.AssociatedExternalType, name)

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	s := convert.NewSensitive(a.Sensitive)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	// TODO: codegen-spec: Add interface to <Type>Validators []<Type>Validator to allow retrieval of []*specschema.CustomValidator
	// TODO: codegen-spec: Will need equivalent for <Type>PlanModifiers and Default
	var custom []*specschema.CustomValidator

	for _, v := range a.Validators {
		custom = append(custom, v.Custom)
	}

	vc := convert.NewValidatorsCustom(convert.ValidatorTypeBool, custom)

	return GeneratorBoolAttribute{
		AttributeType:            at,
		AssociatedExternalType:   generatorschema.NewAssocExtType(a.AssociatedExternalType),
		ComputedOptionalRequired: c,
		CustomType:               a.CustomType,
		Description:              d,
		DeprecationMessage:       dm,
		Sensitive:                s,
		Validators:               a.Validators,
		ValidatorsCustom:         vc,
	}, nil
}

func (g GeneratorBoolAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorBoolAttribute
}

func (g GeneratorBoolAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

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

func (g GeneratorBoolAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorBoolAttribute)

	if !ok {
		return false
	}

	//TODO: Add equality checks for all struct types

	// TODO: Equality functions that operate on specschema types should be added to codegen-spec repo.
	//if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
	//	return false
	//}

	// TODO: codegen-spec: Add Equality function to allow comparison of []*specschema.CustomValidator

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return true
}

func (g GeneratorBoolAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.BoolAttribute{\n", name))
	b.Write(g.AttributeType.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorBoolAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.BoolValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
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
