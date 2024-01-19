// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorInt64Attribute struct {
	AssociatedExternalType *generatorschema.AssocExtType
	OptionalRequired       convert.OptionalRequired
	CustomTypePrimitive    convert.CustomTypePrimitive
	DeprecationMessage     convert.DeprecationMessage
	Description            convert.Description
	Sensitive              convert.Sensitive
	ValidatorsCustom       convert.ValidatorsCustom
}

func NewGeneratorInt64Attribute(name string, a *provider.Int64Attribute) (GeneratorInt64Attribute, error) {
	if a == nil {
		return GeneratorInt64Attribute{}, fmt.Errorf("*provider.Int64Attribute is nil")
	}

	c := convert.NewOptionalRequired(a.OptionalRequired)

	ctp := convert.NewCustomTypePrimitive(a.CustomType, a.AssociatedExternalType, name)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	s := convert.NewSensitive(a.Sensitive)

	vc := convert.NewValidatorsCustom(convert.ValidatorTypeInt64, a.Validators.CustomValidators())

	return GeneratorInt64Attribute{
		AssociatedExternalType: generatorschema.NewAssocExtType(a.AssociatedExternalType),
		OptionalRequired:       c,
		CustomTypePrimitive:    ctp,
		DeprecationMessage:     dm,
		Description:            d,
		Sensitive:              s,
		ValidatorsCustom:       vc,
	}, nil
}

func (g GeneratorInt64Attribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorInt64Attribute
}

func (g GeneratorInt64Attribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(g.CustomTypePrimitive.Imports())

	imports.Append(g.ValidatorsCustom.Imports())

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

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.OptionalRequired.Equal(h.OptionalRequired) {
		return false
	}

	if !g.CustomTypePrimitive.Equal(h.CustomTypePrimitive) {
		return false
	}

	if !g.DeprecationMessage.Equal(h.DeprecationMessage) {
		return false
	}

	if !g.Description.Equal(h.Description) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.ValidatorsCustom.Equal(h.ValidatorsCustom)
}

func (g GeneratorInt64Attribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.Int64Attribute{\n", name))
	b.Write(g.CustomTypePrimitive.Schema())
	b.Write(g.OptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorInt64Attribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.Int64ValueType,
	}

	customValueType := g.CustomTypePrimitive.ValueType()

	if customValueType != "" {
		field.ValueType = customValueType
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
func (g GeneratorInt64Attribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase()), nil
	}

	return "basetypes.Int64Type{}", nil
}

// AttrValue returns a string representation of a basetypes.Int64Valuable type.
func (g GeneratorInt64Attribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.Int64Value"
}

func (g GeneratorInt64Attribute) To() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "ValueInt64Pointer",
	}, nil
}

func (g GeneratorInt64Attribute) From() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return generatorschema.ToFromConversion{
		Default: "Int64PointerValue",
	}, nil
}
