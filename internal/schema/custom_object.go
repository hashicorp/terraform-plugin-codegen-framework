// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

type CustomObjectType struct {
	Name       string
	AttrValues map[Name]string
	templates  map[string]string
}

func NewCustomObjectType(name string, attrValues map[string]string) CustomObjectType {
	t := map[string]string{
		"equal":              ObjectTypeEqualTemplate,
		"string":             ObjectTypeStringTemplate,
		"typable":            ObjectTypeTypableTemplate,
		"type":               ObjectTypeTypeTemplate,
		"value":              ObjectTypeValueTemplate,
		"valueFromObject":    ObjectTypeValueFromObjectTemplate,
		"valueFromTerraform": ObjectTypeValueFromTerraformTemplate,
		"valueMust":          ObjectTypeValueMustTemplate,
		"valueNull":          ObjectTypeValueNullTemplate,
		"valueType":          ObjectTypeValueTypeTemplate,
		"valueUnknown":       ObjectTypeValueUnknownTemplate,
	}

	a := make(map[Name]string, len(attrValues))

	for k, v := range attrValues {
		a[Name(k)] = v
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

		buf.Write([]byte("\n"))

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
		AttrValues map[Name]string
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
		AttrValues map[Name]string
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
	AttributeTypes map[Name]string
	AttrTypes      map[Name]string
	AttrValues     map[Name]string
	templates      map[string]string
}

func NewCustomObjectValue(name string, attributeTypes, attrTypes, attrValues map[string]string) CustomObjectValue {
	t := map[string]string{
		"attributeTypes":   ObjectValueAttributeTypesTemplate,
		"equal":            ObjectValueEqualTemplate,
		"isNull":           ObjectValueIsNullTemplate,
		"isUnknown":        ObjectValueIsUnknownTemplate,
		"string":           ObjectValueStringTemplate,
		"toObjectValue":    ObjectValueToObjectValueTemplate,
		"toTerraformValue": ObjectValueToTerraformValueTemplate,
		"type":             ObjectValueTypeTemplate,
		"valuable":         ObjectValueValuableTemplate,
		"value":            ObjectValueValueTemplate,
	}

	attribTypes := make(map[Name]string, len(attributeTypes))

	for k, v := range attributeTypes {
		attribTypes[Name(k)] = v
	}

	attrTyps := make(map[Name]string, len(attrTypes))

	for k, v := range attrTypes {
		attrTyps[Name(k)] = v
	}

	attrVals := make(map[Name]string, len(attrValues))

	for k, v := range attrValues {
		attrVals[Name(k)] = v
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

		buf.Write([]byte("\n"))

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
		AttrTypes map[Name]string
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
		AttrValues map[Name]string
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
		AttributeTypes map[Name]string
		AttrTypes      map[Name]string
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
		AttrTypes map[Name]string
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
		AttrValues map[Name]string
	}{
		Name:       model.SnakeCaseToCamelCase(c.Name),
		AttrValues: c.AttrValues,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
