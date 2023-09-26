// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templates"
)

type GeneratorListNestedAttribute struct {
	schema.ListNestedAttribute

	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	CustomType   *specschema.CustomType
	NestedObject GeneratorNestedAttributeObject
	Validators   specschema.ListValidators
}

func (g GeneratorListNestedAttribute) AssocExtType() *generatorschema.AssocExtType {
	return g.NestedObject.AssociatedExternalType
}

func (g GeneratorListNestedAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorListNestedAttribute
}

func (g GeneratorListNestedAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	customTypeImports = generatorschema.CustomTypeImports(g.NestedObject.CustomType)
	imports.Append(customTypeImports)

	for _, v := range g.NestedObject.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	for _, v := range g.NestedObject.Attributes {
		imports.Append(v.Imports())
	}

	// TODO: This should only be added if custom types (models) are being generated.
	imports.Append(generatorschema.AttrImports())

	imports.Append(g.NestedObject.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorListNestedAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListNestedAttribute)

	if !ok {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	if !g.NestedObject.Equal(h.NestedObject) {
		return false
	}

	return g.ListNestedAttribute.Equal(h.ListNestedAttribute)
}

func (g GeneratorListNestedAttribute) Schema(name string) (string, error) {
	type attribute struct {
		Name                         string
		TypeValueName                string
		Attributes                   string
		GeneratorListNestedAttribute GeneratorListNestedAttribute
	}

	attributesStr, err := g.NestedObject.Attributes.Schema()

	if err != nil {
		return "", err
	}

	a := attribute{
		Name:                         name,
		TypeValueName:                model.SnakeCaseToCamelCase(name),
		Attributes:                   attributesStr,
		GeneratorListNestedAttribute: g,
	}

	t, err := template.New("list_nested_attribute").Parse(listNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, a)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorListNestedAttribute) ModelField(name string) (model.Field, error) {
	f := model.Field{
		Name:      model.SnakeCaseToCamelCase(name),
		TfsdkName: name,
		ValueType: model.ListValueType,
	}

	if g.CustomType != nil {
		f.ValueType = g.CustomType.ValueType
	}

	return f, nil
}

func (g GeneratorListNestedAttribute) GetAttributes() generatorschema.GeneratorAttributes {
	return g.NestedObject.Attributes
}

type CustomObjectType struct {
	Name       string
	AttrValues map[key]string
	templates  map[string]string
}

func NewCustomObjectType(name string, fields map[key]field, attrValues map[key]string) CustomObjectType {
	t := map[string]string{
		"equal":              templates.ObjectTypeEqualTemplate,
		"string":             templates.ObjectTypeStringTemplate,
		"typable":            templates.ObjectTypeTypableTemplate,
		"type":               templates.ObjectTypeTypeTemplate,
		"value":              templates.ObjectTypeValueTemplate,
		"valueFromObject":    templates.ObjectTypeValueFromObjectTemplate,
		"valueFromTerraform": templates.ObjectTypeValueFromTerraformTemplate,
		"valueMust":          templates.ObjectTypeValueMustTemplate,
		"valueNull":          templates.ObjectTypeValueNullTemplate,
		"valueType":          templates.ObjectTypeValueTypeTemplate,
		"valueUnknown":       templates.ObjectTypeValueUnknownTemplate,
	}

	return CustomObjectType{
		Name:       name,
		AttrValues: attrValues,
		templates:  t,
	}
}

func (c CustomObjectType) Render() ([]byte, error) {
	var buf bytes.Buffer

	renderFuncs := []func() ([]byte, error){
		c.renderTypable,
		c.renderType,
		c.renderEqual,
		c.renderString,
		c.renderValueFromObject,
		c.renderValueNull,
		c.renderValueUnknown,
		c.renderValue,
		c.renderValueMust,
		c.renderValueFromTerraform,
		c.renderValueType,
	}

	for _, f := range renderFuncs {
		b, err := f()

		if err != nil {
			return nil, err
		}

		buf.Write(b)
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderEqual() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["equal"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderString() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["string"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderTypable() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["typable"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderType() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["type"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["value"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name       string
		AttrValues map[key]string
	}{
		Name:       model.SnakeCaseToCamelCase(c.Name),
		AttrValues: c.AttrValues,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderValueFromObject() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueFromObject"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name       string
		AttrValues map[key]string
	}{
		Name:       model.SnakeCaseToCamelCase(c.Name),
		AttrValues: c.AttrValues,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderValueFromTerraform() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueFromTerraform"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderValueMust() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueMust"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderValueNull() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueNull"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderValueType() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueType"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectType) renderValueUnknown() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueUnknown"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: model.SnakeCaseToCamelCase(c.Name),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// TODO: Separate i.e., extract into AttrTypes - map[string]AttrType, AttrValue - map[string]AttrValue etc
type field struct {
	AttrType         string
	AttrValue        string
	FieldName        string
	FieldType        string
	FieldNameLCFirst string
}

func (g GeneratorListNestedAttribute) AttributeAttrValues() (map[key]string, error) {
	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.NestedObject.Attributes))

	for k := range g.NestedObject.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	attrValues := make(map[key]string)

	for _, k := range attributeKeys {
		switch g.NestedObject.Attributes[k].GeneratorSchemaType() {
		case generatorschema.GeneratorBoolAttribute:
			attrValues[key(k)] = "basetypes.BoolValue"
		case generatorschema.GeneratorFloat64Attribute:
			attrValues[key(k)] = "basetypes.Float64Value"
		case generatorschema.GeneratorInt64Attribute:
			attrValues[key(k)] = "basetypes.Int64Value"
		case generatorschema.GeneratorListAttribute:
			attrValues[key(k)] = "basetypes.ListValue"
		case generatorschema.GeneratorListNestedAttribute:
			attrValues[key(k)] = "basetypes.ListValue"
		case generatorschema.GeneratorMapAttribute:
			attrValues[key(k)] = "basetypes.MapValue"
		case generatorschema.GeneratorMapNestedAttribute:
			attrValues[key(k)] = "basetypes.MapValue"
		case generatorschema.GeneratorNumberAttribute:
			attrValues[key(k)] = "basetypes.NumberValue"
		case generatorschema.GeneratorObjectAttribute:
			attrValues[key(k)] = "basetypes.ObjectValue"
		case generatorschema.GeneratorSetAttribute:
			attrValues[key(k)] = "basetypes.SetValue"
		case generatorschema.GeneratorSetNestedAttribute:
			attrValues[key(k)] = "basetypes.SetValue"
		case generatorschema.GeneratorSingleNestedAttribute:
			attrValues[key(k)] = "basetypes.ObjectValue"
		case generatorschema.GeneratorStringAttribute:
			attrValues[key(k)] = "basetypes.StringValue"
		}
	}

	return attrValues, nil
}

type key string

func (k key) CamelCaseLCFirst() string {
	camelCased := k.CamelCase()

	if len(camelCased) < 2 {
		return strings.ToLower(camelCased)
	}

	return strings.ToLower(camelCased[:1]) + camelCased[1:]
}

func (k key) CamelCase() string {
	split := strings.Split(string(k), "_")

	var camelCased string

	for _, v := range split {
		if len(v) < 1 {
			continue
		}

		firstChar := v[0:1]
		ucFirstChar := strings.ToUpper(firstChar)

		if len(v) < 2 {
			camelCased += ucFirstChar
			continue
		}

		camelCased += ucFirstChar + v[1:]
	}

	return camelCased
}

func (k key) String() string {
	return string(k)
}

func (g GeneratorListNestedAttribute) ExtractFields(name string) (map[key]field, error) {
	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.NestedObject.Attributes))

	for k := range g.NestedObject.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	fields := make(map[key]field)

	// Populate fields map for use in template.
	for _, k := range attributeKeys {
		switch g.NestedObject.Attributes[k].GeneratorSchemaType() {
		case generatorschema.GeneratorBoolAttribute:
			fields[key(k)] = field{
				AttrType:  "basetypes.BoolType{}",
				AttrValue: "basetypes.BoolValue",
			}
		case generatorschema.GeneratorFloat64Attribute:
			fields[key(k)] = field{
				AttrType:  "basetypes.Float64Type{}",
				AttrValue: "basetypes.Float64Value",
			}
		case generatorschema.GeneratorInt64Attribute:
			fields[key(k)] = field{
				AttrType:  "basetypes.Int64Type{}",
				AttrValue: "basetypes.Int64Value",
			}
		case generatorschema.GeneratorListAttribute:
			if e, ok := g.NestedObject.Attributes[k].(generatorschema.Elements); ok {
				elemType, err := generatorschema.ElementTypeString(e.ElemType())
				if err != nil {
					return nil, err
				}
				fields[key(k)] = field{
					AttrType:  fmt.Sprintf("basetypes.ListType{\nElemType: %s,\n}", elemType),
					AttrValue: "basetypes.ListValue",
				}
			} else {
				return nil, fmt.Errorf("%s.%s attribute is a ListType but does not implement Elements interface", name, k)
			}
		case generatorschema.GeneratorListNestedAttribute:
			fields[key(k)] = field{
				AttrType:  fmt.Sprintf("basetypes.ListType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
				AttrValue: "basetypes.ListValue",
				FieldType: "ListNestedAttribute",
			}
		case generatorschema.GeneratorMapAttribute:
			if e, ok := g.NestedObject.Attributes[k].(generatorschema.Elements); ok {
				elemType, err := generatorschema.ElementTypeString(e.ElemType())
				if err != nil {
					return nil, err
				}
				fields[key(k)] = field{
					AttrType:  fmt.Sprintf("basetypes.MapType{\nElemType: %s,\n}", elemType),
					AttrValue: "basetypes.MapValue",
				}
			} else {
				return nil, fmt.Errorf("%s.%s attribute is a MapType but does not implement Elements interface", name, k)
			}
		case generatorschema.GeneratorMapNestedAttribute:
			fields[key(k)] = field{
				AttrType:  fmt.Sprintf("basetypes.MapType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
				AttrValue: "basetypes.MapValue",
				FieldType: "MapNestedAttribute",
			}
		case generatorschema.GeneratorNumberAttribute:
			fields[key(k)] = field{
				AttrType:  "basetypes.NumberType{}",
				AttrValue: "basetypes.NumberValue",
			}
		case generatorschema.GeneratorObjectAttribute:
			if o, ok := g.NestedObject.Attributes[k].(generatorschema.Attrs); ok {
				attrTypes, err := generatorschema.AttrTypesString(o.AttrTypes())
				if err != nil {
					return nil, err
				}
				fields[key(k)] = field{
					AttrType:  fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", attrTypes),
					AttrValue: "basetypes.ObjectValue",
				}
			} else {
				return nil, fmt.Errorf("%s.%s attribute is an ObjectType but does not implement Attrs interface", name, k)
			}
		case generatorschema.GeneratorSetAttribute:
			if e, ok := g.NestedObject.Attributes[k].(generatorschema.Elements); ok {
				elemType, err := generatorschema.ElementTypeString(e.ElemType())
				if err != nil {
					return nil, err
				}
				fields[key(k)] = field{
					AttrType:  fmt.Sprintf("basetypes.SetType{\nElemType: %s,\n}", elemType),
					AttrValue: "basetypes.SetValue",
				}
			} else {
				return nil, fmt.Errorf("%s.%s attribute is a SetType but does not implement Elements interface", name, k)
			}
		case generatorschema.GeneratorSetNestedAttribute:
			fields[key(k)] = field{
				AttrType:  fmt.Sprintf("basetypes.SetType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
				AttrValue: "basetypes.SetValue",
				FieldType: "SetNestedAttribute",
			}
		case generatorschema.GeneratorSingleNestedAttribute:
			fields[key(k)] = field{
				AttrType:  fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k)),
				AttrValue: "basetypes.ObjectValue",
				FieldType: "SingleNestedAttribute",
			}

		case generatorschema.GeneratorStringAttribute:
			fields[key(k)] = field{
				AttrType:  "basetypes.StringType{}",
				AttrValue: "basetypes.StringValue",
			}
		}

		f := fields[key(k)]

		camelCaseName := model.SnakeCaseToCamelCase(k)
		lcFirstName := strings.ToLower(camelCaseName[:1]) + camelCaseName[1:]

		f.FieldName = camelCaseName
		f.FieldNameLCFirst = lcFirstName

		fields[key(k)] = f
	}

	return fields, nil
}

func (g GeneratorListNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	fields, err := g.ExtractFields(name)

	if err != nil {
		return nil, err
	}

	attributeAttrValues, err := g.AttributeAttrValues()

	if err != nil {
		return nil, err
	}

	c := NewCustomObjectType(name, fields, attributeAttrValues)

	b, err := c.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.NestedObject.Attributes))

	for k := range g.NestedObject.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	t, err := template.New("model_object_helpers").Parse(templates.ModelObjectHelpersTemplateEdited)

	if err != nil {
		return nil, err
	}

	templateData := struct {
		Name   string
		Fields map[key]field
	}{
		Name:   model.SnakeCaseToCamelCase(name),
		Fields: fields,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	// Recursively call ModelObjectHelpersTemplate() for each attribute that implements
	// Attributes interface (i.e, nested attributes).
	for _, k := range attributeKeys {

		// TODO: Also need to consider how to handle instances in which an associated_external_type
		// has been defined on a type which does not implement CustomTypeAndValue (e.g., bool)
		// If To/From methods are going to be hung off custom value type, then will to generate
		// "wrapped" / embedded types that embed bool in a type that can have To/From methods
		// added to it.
		if c, ok := g.NestedObject.Attributes[k].(generatorschema.CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)

			continue
		}

		// TODO: Remove once refactored to Generator<Type>Attribute|Block
		if a, ok := g.NestedObject.Attributes[k].(generatorschema.Attributes); ok {
			ng := generatorschema.GeneratorSchema{
				Attributes: a.GetAttributes(),
			}

			b, err := ng.ModelObjectHelpersTemplate(k)
			if err != nil {
				return nil, err
			}

			buf.WriteString("\n")
			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}
