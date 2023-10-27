// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorMapAttribute struct {
	schema.MapAttribute

	AssociatedExternalType *generatorschema.AssocExtType
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType    *specschema.CustomType
	Default       *specschema.MapDefault
	ElementType   specschema.ElementType
	PlanModifiers specschema.MapPlanModifiers
	Validators    specschema.MapValidators
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

	if g.Default != nil {
		customDefaultImports := generatorschema.CustomDefaultImports(g.Default.Custom)
		imports.Append(customDefaultImports)
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

// Equal does not delegate to g.MapAttribute.Equal(h.MapAttribute) as the
// call returns false when the ElementType is nil.
func (g GeneratorMapAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorMapAttribute)
	if !ok {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Default.Equal(h.Default) {
		return false
	}

	if !g.ElementType.Equal(h.ElementType) {
		return false
	}

	if !g.PlanModifiers.Equal(h.PlanModifiers) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	if g.Required != h.Required {
		return false
	}

	if g.Optional != h.Optional {
		return false
	}

	if g.Computed != h.Computed {
		return false
	}

	if g.Sensitive != h.Sensitive {
		return false
	}

	if g.Description != h.Description {
		return false
	}

	if g.MarkdownDescription != h.MarkdownDescription {
		return false
	}

	if g.DeprecationMessage != h.DeprecationMessage {
		return false
	}

	return true
}

func mapDefault(d *specschema.MapDefault) string {
	if d == nil {
		return ""
	}

	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorMapAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                  string
		CustomType            string
		Default               string
		ElementType           string
		GeneratorMapAttribute GeneratorMapAttribute
	}

	a := attribute{
		Name:                  name.ToString(),
		Default:               mapDefault(g.Default),
		ElementType:           generatorschema.GetElementType(g.ElementType),
		GeneratorMapAttribute: g,
	}

	switch {
	case g.CustomType != nil:
		a.CustomType = g.CustomType.Type
	case g.AssociatedExternalType != nil:
		a.CustomType = fmt.Sprintf("%sType{\ntypes.MapType{\nElemType: %s,\n},\n}", name.ToPascalCase(), generatorschema.GetElementType(g.ElementType))
	}

	t, err := template.New("map_attribute").Parse(mapAttributeTemplate)
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
	elementFrom := generatorschema.GetElementFromFunc(g.ElementType)

	toFrom := generatorschema.NewToFromMap(name, g.AssociatedExternalType, elementTypeType, elementTypeValue, elementFrom)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AttrType returns a string representation of a basetypes.MapTypable type.
func (g GeneratorMapAttribute) AttrType(name generatorschema.FrameworkIdentifier) (string, error) {
	if g.AssociatedExternalType != nil {
		return fmt.Sprintf("%sType{}", name.ToPascalCase()), nil
	}

	elemType, err := generatorschema.ElementTypeString(g.ElemType())

	if err != nil {
		return "", err
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
