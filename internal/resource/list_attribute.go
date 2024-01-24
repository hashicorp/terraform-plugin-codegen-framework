// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorListAttribute struct {
	AssociatedExternalType   *generatorschema.AssocExtType
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomType               convert.CustomTypeCollection
	Default                  convert.DefaultCustom
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	ElementType              specschema.ElementType
	ElementTypeCollection    convert.ElementType
	PlanModifiers            convert.PlanModifiers
	Sensitive                convert.Sensitive
	Validators               convert.Validators
}

func NewGeneratorListAttribute(name string, a *resource.ListAttribute) (GeneratorListAttribute, error) {
	if a == nil {
		return GeneratorListAttribute{}, fmt.Errorf("*resource.ListAttribute is nil")
	}

	et := convert.NewElementType(a.ElementType)

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	ctc := convert.NewCustomTypeCollection(a.CustomType, a.AssociatedExternalType, convert.CustomCollectionTypeList, string(et.ElementType()), name)

	dc := convert.NewDefaultCustom(a.Default.CustomDefault())

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	pm := convert.NewPlanModifiers(convert.PlanModifierTypeList, a.PlanModifiers.CustomPlanModifiers())

	s := convert.NewSensitive(a.Sensitive)

	v := convert.NewValidators(convert.ValidatorTypeList, a.Validators.CustomValidators())

	return GeneratorListAttribute{
		AssociatedExternalType:   generatorschema.NewAssocExtType(a.AssociatedExternalType),
		ComputedOptionalRequired: c,
		CustomType:               ctc,
		Default:                  dc,
		DeprecationMessage:       dm,
		Description:              d,
		ElementType:              a.ElementType,
		PlanModifiers:            pm,
		ElementTypeCollection:    et,
		Sensitive:                s,
		Validators:               v,
	}, nil
}

func (g GeneratorListAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorListAttribute
}

func (g GeneratorListAttribute) ElemType() specschema.ElementType {
	return g.ElementType
}

func (g GeneratorListAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(g.CustomType.Imports())

	imports.Append(g.ElementTypeCollection.Imports())

	imports.Append(g.Default.Imports())

	imports.Append(g.PlanModifiers.Imports())

	imports.Append(g.Validators.Imports())

	if g.AssociatedExternalType != nil {
		imports.Append(generatorschema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

// Equal does not delegate to g.ListAttribute.Equal(h.ListAttribute) as the
// call returns false when the ElementType is nil.
func (g GeneratorListAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListAttribute)

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

	if !g.Default.Equal(h.Default) {
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

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.Sensitive.Equal(h.Sensitive) {
		return false
	}

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorListAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	customTypeSchema := g.CustomType.Schema()

	b.WriteString(fmt.Sprintf("%q: schema.ListAttribute{\n", name))
	b.Write(customTypeSchema)
	if len(customTypeSchema) == 0 {
		b.Write(g.ElementTypeCollection.Schema())
	}
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.PlanModifiers.Schema())
	b.Write(g.Validators.Schema())
	b.Write(g.Default.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorListAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.ListValueType,
	}

	customValueType := g.CustomType.ValueType()

	if customValueType != "" {
		field.ValueType = customValueType
	}

	return field, nil
}

func (g GeneratorListAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	listType := generatorschema.NewCustomListType(name)

	b, err := listType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	elemType := generatorschema.GetElementType(g.ElementType)

	listValue := generatorschema.NewCustomListValue(name, elemType)

	b, err = listValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorListAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	elementTypeType := generatorschema.GetElementType(g.ElementType)
	elementTypeValue := generatorschema.GetElementValueType(g.ElementType)

	elementFrom, err := generatorschema.GetElementFromFunc(g.ElementType)

	if err != nil {
		return nil, err
	}

	toFrom := generatorschema.NewToFromList(name, g.AssociatedExternalType, elementTypeType, elementTypeValue, elementFrom)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.ListTypable type.
func (g GeneratorListAttribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	elemType, err := generatorschema.ElementTypeString(g.ElemType())

	if err != nil {
		return "", err
	}

	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{\nbasetypes.ListType{\nElemType: %s,\n}}", name.ToPascalCase(), elemType), nil
	}

	return fmt.Sprintf("basetypes.ListType{\nElemType: %s,\n}", elemType), nil
}

// AttrValue returns a string representation of a basetypes.ListValuable type.
func (g GeneratorListAttribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.ListValue"
}

func (g GeneratorListAttribute) To() (generatorschema.ToFromConversion, error) {
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
			GoType: fmt.Sprintf("[]%s", elementGoType),
		},
	}, nil
}

func (g GeneratorListAttribute) From() (generatorschema.ToFromConversion, error) {
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
			TypeValueFrom: "types.ListValueFrom",
		},
	}, nil
}

// CollectionType returns string representations of the element type (e.g., types.BoolType),
// and type value function (e.g., types.ListValue) if there is no associated external type.
func (g GeneratorListAttribute) CollectionType() (map[string]string, error) {
	if g.AssociatedExternalType != nil {
		return nil, nil
	}

	elementType, err := generatorschema.ElementTypeString(g.ElemType())

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"ElementType":   elementType,
		"TypeValueFunc": "types.ListValue",
	}, nil
}
