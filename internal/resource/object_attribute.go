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

type GeneratorObjectAttribute struct {
	AssociatedExternalType   *generatorschema.AssocExtType
	AttributeTypes           specschema.ObjectAttributeTypes
	AttributeTypesObject     convert.ObjectAttributeTypes
	ComputedOptionalRequired convert.ComputedOptionalRequired
	CustomTypeObject         convert.CustomTypeObject
	DefaultCustom            convert.DefaultCustom
	DeprecationMessage       convert.DeprecationMessage
	Description              convert.Description
	PlanModifiersCustom      convert.PlanModifiersCustom
	Sensitive                convert.Sensitive
	Validators               convert.Validators
}

func NewGeneratorObjectAttribute(name string, a *resource.ObjectAttribute) (GeneratorObjectAttribute, error) {
	if a == nil {
		return GeneratorObjectAttribute{}, fmt.Errorf("*resource.ObjectAttribute is nil")
	}

	c := convert.NewComputedOptionalRequired(a.ComputedOptionalRequired)

	cto := convert.NewCustomTypeObject(a.CustomType, a.AssociatedExternalType, name)

	dc := convert.NewDefaultCustom(a.Default.CustomDefault())

	d := convert.NewDescription(a.Description)

	dm := convert.NewDeprecationMessage(a.DeprecationMessage)

	oat := convert.NewObjectAttributeTypes(a.AttributeTypes)

	pm := convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, a.PlanModifiers.CustomPlanModifiers())

	s := convert.NewSensitive(a.Sensitive)

	v := convert.NewValidators(convert.ValidatorTypeObject, a.Validators.CustomValidators())

	return GeneratorObjectAttribute{
		AssociatedExternalType:   generatorschema.NewAssocExtType(a.AssociatedExternalType),
		AttributeTypes:           a.AttributeTypes,
		AttributeTypesObject:     oat,
		ComputedOptionalRequired: c,
		CustomTypeObject:         cto,
		DefaultCustom:            dc,
		DeprecationMessage:       dm,
		Description:              d,
		PlanModifiersCustom:      pm,
		Sensitive:                s,
		Validators:               v,
	}, nil
}

func (g GeneratorObjectAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorObjectAttribute
}

func (g GeneratorObjectAttribute) AttrTypes() specschema.ObjectAttributeTypes {
	return g.AttributeTypes
}

func (g GeneratorObjectAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(g.CustomTypeObject.Imports())

	imports.Append(g.AttributeTypesObject.Imports())

	imports.Append(g.DefaultCustom.Imports())

	imports.Append(g.PlanModifiersCustom.Imports())

	imports.Append(g.Validators.Imports())

	if g.AssociatedExternalType != nil {
		imports.Append(generatorschema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	for _, v := range g.AttrTypes() {
		if v.Number != nil && g.AssociatedExternalType == nil {
			imports.Add(code.Import{
				Path: generatorschema.MathBigImport,
			})
		}
	}

	return imports
}

func (g GeneratorObjectAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorObjectAttribute)

	if !ok {
		return false
	}

	if !g.AssociatedExternalType.Equal(h.AssociatedExternalType) {
		return false
	}

	if !g.AttributeTypes.Equal(h.AttributeTypes) {
		return false
	}

	if !g.AttributeTypesObject.Equal(h.AttributeTypesObject) {
		return false
	}

	if !g.ComputedOptionalRequired.Equal(h.ComputedOptionalRequired) {
		return false
	}

	if !g.CustomTypeObject.Equal(h.CustomTypeObject) {
		return false
	}

	if !g.DefaultCustom.Equal(h.DefaultCustom) {
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

	return g.Validators.Equal(h.Validators)
}

func (g GeneratorObjectAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	var b bytes.Buffer

	customTypeSchema := g.CustomTypeObject.Schema()

	b.WriteString(fmt.Sprintf("%q: schema.ObjectAttribute{\n", name))
	b.Write(customTypeSchema)
	if len(customTypeSchema) == 0 {
		b.Write(g.AttributeTypesObject.Schema())
	}
	b.Write(g.ComputedOptionalRequired.Schema())
	b.Write(g.Sensitive.Schema())
	b.Write(g.Description.Schema())
	b.Write(g.DeprecationMessage.Schema())
	b.Write(g.PlanModifiersCustom.Schema())
	b.Write(g.Validators.Schema())
	b.Write(g.DefaultCustom.Schema())
	b.WriteString("},")

	return b.String(), nil
}

func (g GeneratorObjectAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.ObjectValueType,
	}

	customValueType := g.CustomTypeObject.ValueType()

	if customValueType != "" {
		field.ValueType = customValueType
	}

	return field, nil
}

func (g GeneratorObjectAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	objectType := generatorschema.NewCustomObjectType(name)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attrTypes := generatorschema.GetAttrTypes(g.AttrTypes())

	objectValue := generatorschema.NewCustomObjectValue(name, attrTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorObjectAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	attrTypesToFuncs, err := generatorschema.GetAttrTypesToFuncs(g.AttributeTypes)

	if err != nil {
		return nil, err
	}

	attrTypesFromFuncs, err := generatorschema.GetAttrTypesFromFuncs(g.AttributeTypes)

	if err != nil {
		return nil, err
	}

	toFrom := generatorschema.NewToFromObject(name, g.AssociatedExternalType, attrTypesToFuncs, attrTypesFromFuncs)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.ObjectTypable type.
func (g GeneratorObjectAttribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{\nbasetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}}", name.ToPascalCase(), name.ToPascalCase()), nil
	}

	aTypes, err := generatorschema.AttrTypesString(g.AttrTypes())

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", aTypes), nil
}

// AttrValue returns a string representation of a basetypes.ListValuable type.
func (g GeneratorObjectAttribute) AttrValue(name generatorschema.FrameworkIdentifier) string {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return "basetypes.ObjectValue"
}

func (g GeneratorObjectAttribute) To() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	objectFields := make(map[generatorschema.FrameworkIdentifier]generatorschema.ObjectField, len(g.AttributeTypes))

	for _, v := range g.AttributeTypes {
		objField, err := generatorschema.ObjectFieldTo(v)

		if err != nil {
			return generatorschema.ToFromConversion{}, err
		}

		objectFields[generatorschema.FrameworkIdentifier(v.Name)] = objField
	}

	return generatorschema.ToFromConversion{
		ObjectType: objectFields,
	}, nil
}

func (g GeneratorObjectAttribute) From() (generatorschema.ToFromConversion, error) {
	if g.AssociatedExternalType != nil {
		return generatorschema.ToFromConversion{
			AssocExtType: g.AssociatedExternalType,
		}, nil
	}

	objectFields := make(map[generatorschema.FrameworkIdentifier]generatorschema.ObjectField, len(g.AttributeTypes))

	for _, v := range g.AttributeTypes {
		objField, err := generatorschema.ObjectFieldFrom(v)

		if err != nil {
			return generatorschema.ToFromConversion{}, err
		}

		objectFields[generatorschema.FrameworkIdentifier(v.Name)] = objField
	}

	return generatorschema.ToFromConversion{
		ObjectType: objectFields,
	}, nil
}
