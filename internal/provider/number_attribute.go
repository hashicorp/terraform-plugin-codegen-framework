// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorNumberAttribute struct {
	AssociatedExternalType *schema.AssocExtType
	OptionalRequired       convert.OptionalRequired
	CustomType             convert.CustomTypePrimitive
	DeprecationMessage     convert.DeprecationMessage
	Description            convert.Description
	Sensitive              convert.Sensitive
	Validators             convert.Validators
}

func NewGeneratorNumberAttribute(name string, a *provider.NumberAttribute) (GeneratorNumberAttribute, error) {
	if a == nil {
		return GeneratorNumberAttribute{}, fmt.Errorf("*provider.NumberAttribute is nil")
	}

	c := convert.NewOptionalRequired(a.OptionalRequired)

	ctp := convert.NewCustomTypePrimitive(a.CustomType, a.AssociatedExternalType, name)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	s := convert.NewSensitive(a.Sensitive)

	v := convert.NewValidators(convert.ValidatorTypeNumber, a.Validators.CustomValidators())

	return GeneratorNumberAttribute{
		AssociatedExternalType: schema.NewAssocExtType(a.AssociatedExternalType),
		OptionalRequired:       c,
		CustomType:             ctp,
		DeprecationMessage:     dm,
		Description:            d,
		Sensitive:              s,
		Validators:             v,
	}, nil
}

func (g GeneratorNumberAttribute) GeneratorSchemaType() schema.Type {
	return schema.GeneratorNumberAttribute
}

func (g GeneratorNumberAttribute) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.Validators.Imports())

	if g.AssociatedExternalType != nil {
		imports.Append(schema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorNumberAttribute) Equal(ga schema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorNumberAttribute)

	if !ok {
		return false
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.OptionalRequired.Equal(h.OptionalRequired) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
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

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorNumberAttribute) Schema(name schema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.NumberAttribute{\n", name))
	b.Write(g.CustomType.Schema())
	b.Write(g.OptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.Validators.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorNumberAttribute) ModelField(name schema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.NumberValueType,
	}

	customValueType := g.CustomType.ValueType()

	if customValueType != "" {
		field.ValueType = customValueType
	}

	return field, nil
}

func (g GeneratorNumberAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	numberType := schema.NewCustomNumberType(name)

	b, err := numberType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	numberValue := schema.NewCustomNumberValue(name)

	b, err = numberValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorNumberAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	toFrom := schema.NewToFromNumber(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.NumberTypable type.
func (g GeneratorNumberAttribute) AttrType(name schema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase()), nil
	}

	return "basetypes.NumberType{}", nil
}

// AttrValue returns a string representation of a basetypes.NumberValuable type.
func (g GeneratorNumberAttribute) AttrValue(name schema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.NumberValue"
}

func (g GeneratorNumberAttribute) To() (schema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return schema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return schema.ToFromConversion{
		Default: "ValueBigFloat",
	}, nil
}

func (g GeneratorNumberAttribute) From() (schema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return schema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return schema.ToFromConversion{
		Default: "NumberValue",
	}, nil
}
