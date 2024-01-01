// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorMapAttribute struct {
	AssociatedExternalType *generatorschema.AssocExtType
	OptionalRequired       convert.OptionalRequired
	CustomType             *specschema.CustomType
	CustomTypeCollection   convert.CustomTypeCollection
	DeprecationMessage     convert.DeprecationMessage
	Description            convert.Description
	ElementType            specschema.ElementType
	ElementTypeCollection  convert.ElementType
	Sensitive              convert.Sensitive
	Validators             specschema.MapValidators
	ValidatorsCustom       convert.ValidatorsCustom
}

func NewGeneratorMapAttribute(name string, a *provider.MapAttribute) (GeneratorMapAttribute, error) {
	if a == nil {
		return GeneratorMapAttribute{}, fmt.Errorf("*provider.MapAttribute is nil")
	}

	et := convert.NewElementType(a.ElementType)

	c := convert.NewOptionalRequired(a.OptionalRequired)

	ctc := convert.NewCustomTypeCollection(a.CustomType, a.AssociatedExternalType, convert.CustomCollectionTypeMap, string(et.ElementType()), name)

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	s := convert.NewSensitive(a.Sensitive)

	vc := convert.NewValidatorsCustom(convert.ValidatorTypeMap, a.Validators.CustomValidators())

	return GeneratorMapAttribute{
		AssociatedExternalType: generatorschema.NewAssocExtType(a.AssociatedExternalType),
		OptionalRequired:       c,
		CustomType:             a.CustomType,
		CustomTypeCollection:   ctc,
		DeprecationMessage:     dm,
		Description:            d,
		ElementType:            a.ElementType,
		ElementTypeCollection:  et,
		Sensitive:              s,
		Validators:             a.Validators,
		ValidatorsCustom:       vc,
	}, nil
}

func (g GeneratorMapAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorMapAttribute
}

func (g GeneratorMapAttribute) ElemType() specschema.ElementType {
	return g.ElementType
}

func (g GeneratorMapAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	elemTypeImports := generatorschema.GetElementTypeImports(g.ElementType)
	imports.Append(elemTypeImports)

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

// Equal does not delegate to g.MapAttribute.Equal(h.MapAttribute) as the
// call returns false when the ElementType is nil.
func (g GeneratorMapAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorMapAttribute)

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

	if !g.CustomTypeCollection.Equal(h.CustomTypeCollection) {
		return false
	}

	if !g.DeprecationMessage.Equal(h.DeprecationMessage) {
		return false
	}

	if !g.Description.Equal(h.Description) {
		return false
	}

	if !g.ElementType.Equal(h.ElementType) {
		return false
	}

	if !g.ElementTypeCollection.Equal(h.ElementTypeCollection) {
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

func (g GeneratorMapAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	customTypeSchema := g.CustomTypeCollection.Schema()

	b.WriteString(fmt.Sprintf("%q: schema.MapAttribute{\n", name))
	b.Write(customTypeSchema)
	if len(customTypeSchema) == 0 {
		b.Write(g.ElementTypeCollection.Schema())
	}
	b.Write(g.OptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.ValidatorsCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorMapAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.MapValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return field, nil
}

func (g GeneratorMapAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	listType := generatorschema.NewCustomMapType(name)

	b, err := listType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	elemType := generatorschema.GetElementType(g.ElementType)

	listValue := generatorschema.NewCustomMapValue(name, elemType)

	b, err = listValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorMapAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	elementTypeType := generatorschema.GetElementType(g.ElementType)
	elementTypeValue := generatorschema.GetElementValueType(g.ElementType)

	elementFrom, err := generatorschema.GetElementFromFunc(g.ElementType)

	if err != nil {
		return nil, err
	}

	toFrom := generatorschema.NewToFromMap(name, g.AssociatedExternalType, elementTypeType, elementTypeValue, elementFrom)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.MapTypable type.
func (g GeneratorMapAttribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	elemType, err := generatorschema.ElementTypeString(g.ElemType())

	if err != nil {
		return "", err
	}

	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{\nbasetypes.MapType{\nElemType: %s,\n}}", name.ToPascalCase(), elemType), nil
	}

	return fmt.Sprintf("basetypes.MapType{\nElemType: %s,\n}", elemType), nil
}

// AttrValue returns a string representation of a basetypes.MapValuable type.
func (g GeneratorMapAttribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.MapValue"
}

func (g GeneratorMapAttribute) To() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	elementGoType, err := generatorschema.ElementTypeGoType(g.ElementType)

	if err != nil {
		return generatorschema.ToFromConversion{}, err
	}

	return generatorschema.ToFromConversion{
		CollectionType: generatorschema.CollectionFields{
			GoType: fmt.Sprintf("map[string]%s", elementGoType),
		},
	}, nil
}

func (g GeneratorMapAttribute) From() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	elementType, err := generatorschema.ElementTypeString(g.ElementType)

	if err != nil {
		return generatorschema.ToFromConversion{}, err
	}

	return generatorschema.ToFromConversion{
		CollectionType: generatorschema.CollectionFields{
			ElementType:   elementType,
			TypeValueFrom: "types.MapValueFrom",
		},
	}, nil
}

// CollectionType returns string representations of the element type (e.g., types.BoolType),
// and type value function (e.g., types.MapValue) if there is no associated external type.
func (g GeneratorMapAttribute) CollectionType() (map[string]string, error) {
	if g.AssociatedExternalType != nil {
		return nil, nil
	}

	elementType, err := generatorschema.ElementTypeString(g.ElemType())

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"ElementType":   elementType,
		"TypeValueFunc": "types.MapValue",
	}, nil
}
