// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type CustomMapType struct {
	Name      FrameworkIdentifier
	templates map[string]string
}

func NewCustomMapType(name string) CustomMapType {
	t := map[string]string{
		"equal":              MapTypeEqualTemplate,
		"string":             MapTypeStringTemplate,
		"type":               MapTypeTypeTemplate,
		"typable":            MapTypeTypableTemplate,
		"valueFromMap":       MapTypeValueFromMapTemplate,
		"valueFromTerraform": MapTypeValueFromTerraformTemplate,
		"valueType":          MapTypeValueTypeTemplate,
	}

	return CustomMapType{
		Name:      FrameworkIdentifier(name),
		templates: t,
	}
}

func (c CustomMapType) Render() ([]byte, error) {
	var buf bytes.Buffer

	renderFuncs := []func() ([]byte, error){
		c.renderTypable,
		c.renderType,
		c.renderEqual,
		c.renderString,
		c.renderValueFromMap,
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

func (c CustomMapType) renderEqual() ([]byte, error) {
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

func (c CustomMapType) renderString() ([]byte, error) {
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

func (c CustomMapType) renderType() ([]byte, error) {
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

func (c CustomMapType) renderTypable() ([]byte, error) {
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

func (c CustomMapType) renderValueFromMap() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueFromMap"])

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

func (c CustomMapType) renderValueFromTerraform() ([]byte, error) {
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

func (c CustomMapType) renderValueType() ([]byte, error) {
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

type CustomMapValue struct {
	Name        FrameworkIdentifier
	ElementType string
	templates   map[string]string
}

func NewCustomMapValue(name, elemType string) CustomMapValue {
	t := map[string]string{
		"equal":    MapValueEqualTemplate,
		"type":     MapValueTypeTemplate,
		"valuable": MapValueValuableTemplate,
		"value":    MapValueValueTemplate,
	}

	return CustomMapValue{
		Name:        FrameworkIdentifier(name),
		ElementType: elemType,
		templates:   t,
	}
}

func (c CustomMapValue) Render() ([]byte, error) {
	var buf bytes.Buffer

	renderFuncs := []func() ([]byte, error){
		c.renderValuable,
		c.renderValue,
		c.renderEqual,
		c.renderType,
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

func (c CustomMapValue) renderEqual() ([]byte, error) {
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

func (c CustomMapValue) renderType() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["type"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name        string
		ElementType string
	}{
		Name:        c.Name.ToPascalCase(),
		ElementType: c.ElementType,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c CustomMapValue) renderValuable() ([]byte, error) {
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

func (c CustomMapValue) renderValue() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["value"])

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
