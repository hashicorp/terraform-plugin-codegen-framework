// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
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

func NewCustomObjectType(name string, attrValues map[string]string) CustomObjectType {
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

	a := make(map[key]string, len(attrValues))

	for k, v := range attrValues {
		a[key(k)] = v
	}

	return CustomObjectType{
		Name:       name,
		AttrValues: a,
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

type CustomObjectValue struct {
	Name           string
	AttributeTypes map[key]string
	AttrTypes      map[key]string
	AttrValues     map[key]string
	templates      map[string]string
}

func NewCustomObjectValue(name string, attributeTypes, attrTypes, attrValues map[string]string) CustomObjectValue {
	t := map[string]string{
		"attributeTypes":   templates.ObjectValueAttributeTypesTemplate,
		"equal":            templates.ObjectValueEqualTemplate,
		"isNull":           templates.ObjectValueIsNullTemplate,
		"isUnknown":        templates.ObjectValueIsUnknownTemplate,
		"string":           templates.ObjectValueStringTemplate,
		"toObjectValue":    templates.ObjectValueToObjectValueTemplate,
		"toTerraformValue": templates.ObjectValueToTerraformValueTemplate,
		"type":             templates.ObjectValueTypeTemplate,
		"valuable":         templates.ObjectValueValuableTemplate,
		"value":            templates.ObjectValueValueTemplate,
	}

	attribTypes := make(map[key]string, len(attributeTypes))

	for k, v := range attributeTypes {
		attribTypes[key(k)] = v
	}

	attrTyps := make(map[key]string, len(attrTypes))

	for k, v := range attrTypes {
		attrTyps[key(k)] = v
	}

	attrVals := make(map[key]string, len(attrValues))

	for k, v := range attrValues {
		attrVals[key(k)] = v
	}

	return CustomObjectValue{
		Name:           name,
		AttributeTypes: attribTypes,
		AttrTypes:      attrTyps,
		AttrValues:     attrVals,
		templates:      t,
	}
}

func (c CustomObjectValue) Render() ([]byte, error) {
	var buf bytes.Buffer

	renderFuncs := []func() ([]byte, error){
		c.renderValuable,
		c.renderValue,
		c.renderToTerraformValue,
		c.renderIsNull,
		c.renderIsUnknown,
		c.renderString,
		c.renderToObjectValue,
		c.renderEqual,
		c.renderType,
		c.renderAttributeTypes,
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

func (c CustomObjectValue) renderAttributeTypes() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["attributeTypes"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name      string
		AttrTypes map[key]string
	}{
		Name:      model.SnakeCaseToCamelCase(c.Name),
		AttrTypes: c.AttrTypes,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectValue) renderEqual() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["equal"])

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

func (c CustomObjectValue) renderIsNull() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["isNull"])

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

func (c CustomObjectValue) renderIsUnknown() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["isUnknown"])

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

func (c CustomObjectValue) renderString() ([]byte, error) {
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

func (c CustomObjectValue) renderToObjectValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["toObjectValue"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name           string
		AttributeTypes map[key]string
		AttrTypes      map[key]string
	}{
		Name:           model.SnakeCaseToCamelCase(c.Name),
		AttributeTypes: c.AttributeTypes,
		AttrTypes:      c.AttrTypes,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectValue) renderToTerraformValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["toTerraformValue"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name      string
		AttrTypes map[key]string
	}{
		Name:      model.SnakeCaseToCamelCase(c.Name),
		AttrTypes: c.AttrTypes,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomObjectValue) renderType() ([]byte, error) {
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

func (c CustomObjectValue) renderValuable() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valuable"])

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

func (c CustomObjectValue) renderValue() ([]byte, error) {
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

func (g GeneratorListNestedAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	var buf bytes.Buffer

	attributeAttrValues, err := g.NestedObject.Attributes.AttrValues()

	if err != nil {
		return nil, err
	}

	objectType := NewCustomObjectType(name, attributeAttrValues)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attributeTypes, err := g.NestedObject.Attributes.AttributeTypes()

	if err != nil {
		return nil, err
	}

	attributeAttrTypes, err := g.NestedObject.Attributes.AttrTypes()

	if err != nil {
		return nil, err
	}

	objectValue := NewCustomObjectValue(name, attributeTypes, attributeAttrTypes, attributeAttrValues)

	b, err = objectValue.Render()

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
