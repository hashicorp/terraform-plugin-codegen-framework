// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type CustomNestedObjectType struct {
	Name       FrameworkIdentifier
	AttrValues map[FrameworkIdentifier]string
	templates  map[string]string
}

func NewCustomNestedObjectType(name string, attrValues map[string]string) CustomNestedObjectType {
	t := map[string]string{
		"equal":              NestedObjectTypeEqualTemplate,
		"string":             NestedObjectTypeStringTemplate,
		"typable":            NestedObjectTypeTypableTemplate,
		"type":               NestedObjectTypeTypeTemplate,
		"value":              NestedObjectTypeValueTemplate,
		"valueFromObject":    NestedObjectTypeValueFromObjectTemplate,
		"valueFromTerraform": NestedObjectTypeValueFromTerraformTemplate,
		"valueMust":          NestedObjectTypeValueMustTemplate,
		"valueNull":          NestedObjectTypeValueNullTemplate,
		"valueType":          NestedObjectTypeValueTypeTemplate,
		"valueUnknown":       NestedObjectTypeValueUnknownTemplate,
	}

	a := make(map[FrameworkIdentifier]string, len(attrValues))

	for k, v := range attrValues {
		a[FrameworkIdentifier(k)] = v
	}

	return CustomNestedObjectType{
		Name:       FrameworkIdentifier(name),
		AttrValues: a,
		templates:  t,
	}
}

func (c CustomNestedObjectType) Render() ([]byte, error) {
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

func (c CustomNestedObjectType) renderEqual() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["equal"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderString() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["string"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderTypable() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["typable"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderType() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["type"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["value"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name       string
		AttrValues map[FrameworkIdentifier]string
	}{
		Name:       c.Name.ToPascalCase(),
		AttrValues: c.AttrValues,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderValueFromObject() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueFromObject"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name       string
		AttrValues map[FrameworkIdentifier]string
	}{
		Name:       c.Name.ToPascalCase(),
		AttrValues: c.AttrValues,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderValueFromTerraform() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueFromTerraform"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderValueMust() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueMust"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderValueNull() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueNull"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderValueType() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueType"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectType) renderValueUnknown() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueUnknown"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type CustomNestedObjectValue struct {
	Name            FrameworkIdentifier
	AttributeTypes  map[FrameworkIdentifier]string
	AttrTypes       map[FrameworkIdentifier]string
	AttrValues      map[FrameworkIdentifier]string
	CollectionTypes map[FrameworkIdentifier]map[string]string
	templates       map[string]string
}

func NewCustomNestedObjectValue(name string, attributeTypes, attrTypes, attrValues map[string]string, collectionTypes map[string]map[string]string) CustomNestedObjectValue {
	t := map[string]string{
		"attributeTypes":   NestedObjectValueAttributeTypesTemplate,
		"equal":            NestedObjectValueEqualTemplate,
		"isNull":           NestedObjectValueIsNullTemplate,
		"isUnknown":        NestedObjectValueIsUnknownTemplate,
		"string":           NestedObjectValueStringTemplate,
		"toObjectValue":    NestedObjectValueToObjectValueTemplate,
		"toTerraformValue": NestedObjectValueToTerraformValueTemplate,
		"type":             NestedObjectValueTypeTemplate,
		"valuable":         NestedObjectValueValuableTemplate,
		"value":            NestedObjectValueValueTemplate,
	}

	attribTypes := make(map[FrameworkIdentifier]string, len(attributeTypes))

	for k, v := range attributeTypes {
		attribTypes[FrameworkIdentifier(k)] = v
	}

	attrTyps := make(map[FrameworkIdentifier]string, len(attrTypes))

	for k, v := range attrTypes {
		attrTyps[FrameworkIdentifier(k)] = v
	}

	attrVals := make(map[FrameworkIdentifier]string, len(attrValues))

	for k, v := range attrValues {
		attrVals[FrameworkIdentifier(k)] = v
	}

	collectionTyps := make(map[FrameworkIdentifier]map[string]string, len(collectionTypes))

	for k, v := range collectionTypes {
		collectionTyps[FrameworkIdentifier(k)] = v
	}

	return CustomNestedObjectValue{
		Name:            FrameworkIdentifier(name),
		AttributeTypes:  attribTypes,
		AttrTypes:       attrTyps,
		AttrValues:      attrVals,
		CollectionTypes: collectionTyps,
		templates:       t,
	}
}

func (c CustomNestedObjectValue) Render() ([]byte, error) {
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

func (c CustomNestedObjectValue) renderAttributeTypes() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["attributeTypes"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name      string
		AttrTypes map[FrameworkIdentifier]string
	}{
		Name:      c.Name.ToPascalCase(),
		AttrTypes: c.AttrTypes,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderEqual() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["equal"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name       string
		AttrValues map[FrameworkIdentifier]string
	}{
		Name:       c.Name.ToPascalCase(),
		AttrValues: c.AttrValues,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderIsNull() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["isNull"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderIsUnknown() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["isUnknown"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderString() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["string"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderToObjectValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["toObjectValue"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name            string
		AttributeTypes  map[FrameworkIdentifier]string
		AttrTypes       map[FrameworkIdentifier]string
		CollectionTypes map[FrameworkIdentifier]map[string]string
	}{
		Name:            c.Name.ToPascalCase(),
		AttributeTypes:  c.AttributeTypes,
		AttrTypes:       c.AttrTypes,
		CollectionTypes: c.CollectionTypes,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderToTerraformValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["toTerraformValue"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name      string
		AttrTypes map[FrameworkIdentifier]string
	}{
		Name:      c.Name.ToPascalCase(),
		AttrTypes: c.AttrTypes,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderType() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["type"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderValuable() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valuable"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name string
	}{
		Name: c.Name.ToPascalCase(),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomNestedObjectValue) renderValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["value"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name       string
		AttrValues map[FrameworkIdentifier]string
	}{
		Name:       c.Name.ToPascalCase(),
		AttrValues: c.AttrValues,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
