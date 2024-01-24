// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorFloat64Attribute struct {
	AssociatedExternalType   *schema.AssocExtType
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomType               convert.CustomTypePrimitive
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	Sensitive                convert.Sensitive
	Validators               convert.Validators
}

func NewGeneratorFloat64Attribute(name string, a *datasource.Float64Attribute) (GeneratorFloat64Attribute, error) {
	if a == nil {
		return GeneratorFloat64Attribute{}, fmt.Errorf("*datasource.Float64Attribute is nil")
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ctp := convert.NewCustomTypePrimitive(a.CustomType, a.AssociatedExternalType, name)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	s := convert.NewSensitive(a.Sensitive)

	v := convert.NewValidators(convert.ValidatorTypeFloat64, a.Validators.CustomValidators())

	return GeneratorFloat64Attribute{
		AssociatedExternalType:   schema.NewAssocExtType(a.AssociatedExternalType),
		ComputedOptionalRequired: c,
		CustomType:               ctp,
		DeprecationMessage:       dm,
		Description:              d,
		Sensitive:                s,
		Validators:               v,
	}, nil
}

func (g GeneratorFloat64Attribute) GeneratorSchemaType() schema.Type {
	return schema.GeneratorFloat64Attribute
}

func (g GeneratorFloat64Attribute) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.Validators.Imports())

	if g.AssociatedExternalType != nil {
		imports.Append(schema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorFloat64Attribute) Equal(ga schema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorFloat64Attribute)

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

func (g GeneratorFloat64Attribute) Schema(name schema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("%q: schema.Float64Attribute{\n", name))
	b.Write(g.CustomType.Schema())
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.Validators.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorFloat64Attribute) ModelField(name schema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.Float64ValueType,
	}

	customValueType := g.CustomType.ValueType()

	if customValueType != "" {
		field.ValueType = customValueType
	}

	return field, nil
}

func (g GeneratorFloat64Attribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	float64Type := schema.NewCustomFloat64Type(name)

	b, err := float64Type.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	float64Value := schema.NewCustomFloat64Value(name)

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

	toFrom := schema.NewToFromFloat64(name, g.AssociatedExternalType)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.Float64Typable type.
func (g GeneratorFloat64Attribute) AttrType(name schema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase()), nil
	}

	return "basetypes.Float64Type{}", nil
}

// AttrValue returns a string representation of a basetypes.Float64Valuable type.
func (g GeneratorFloat64Attribute) AttrValue(name schema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.Float64Value"
}

func (g GeneratorFloat64Attribute) To() (schema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return schema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return schema.ToFromConversion{
		Default: "ValueFloat64Pointer",
	}, nil
}

func (g GeneratorFloat64Attribute) From() (schema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return schema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	return schema.ToFromConversion{
		Default: "Float64PointerValue",
	}, nil
}
