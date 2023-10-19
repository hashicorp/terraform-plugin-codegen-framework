// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type CustomFloat64Type struct {
	Name      FrameworkIdentifier
	templates map[string]string
}

func NewCustomFloat64Type(name string) CustomFloat64Type {
	t := map[string]string{
		"equal":              Float64TypeEqualTemplate,
		"string":             Float64TypeStringTemplate,
		"type":               Float64TypeTypeTemplate,
		"typable":            Float64TypeTypableTemplate,
		"valueFromFloat64":   Float64TypeValueFromFloat64Template,
		"valueFromTerraform": Float64TypeValueFromTerraformTemplate,
		"valueType":          Float64TypeValueTypeTemplate,
	}

	return CustomFloat64Type{
		Name:      FrameworkIdentifier(name),
		templates: t,
	}
}

func (c CustomFloat64Type) Render() ([]byte, error) {
	var buf bytes.Buffer

	renderFuncs := []func() ([]byte, error){
		c.renderTypable,
		c.renderType,
		c.renderEqual,
		c.renderString,
		c.renderValueFromFloat64,
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

func (c CustomFloat64Type) renderEqual() ([]byte, error) {
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

func (c CustomFloat64Type) renderString() ([]byte, error) {
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

func (c CustomFloat64Type) renderType() ([]byte, error) {
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

func (c CustomFloat64Type) renderTypable() ([]byte, error) {
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

func (c CustomFloat64Type) renderValueFromFloat64() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(c.templates["valueFromFloat64"])

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

func (c CustomFloat64Type) renderValueFromTerraform() ([]byte, error) {
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

func (c CustomFloat64Type) renderValueType() ([]byte, error) {
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

type CustomFloat64Value struct {
	Name      FrameworkIdentifier
	templates map[string]string
}

func NewCustomFloat64Value(name string) CustomFloat64Value {
	t := map[string]string{
		"equal":    Float64ValueEqualTemplate,
		"type":     Float64ValueTypeTemplate,
		"valuable": Float64ValueValuableTemplate,
		"value":    Float64ValueValueTemplate,
	}

	return CustomFloat64Value{
		Name:      FrameworkIdentifier(name),
		templates: t,
	}
}

func (c CustomFloat64Value) Render() ([]byte, error) {
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

func (c CustomFloat64Value) renderEqual() ([]byte, error) {
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

func (c CustomFloat64Value) renderType() ([]byte, error) {
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

func (c CustomFloat64Value) renderValuable() ([]byte, error) {
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

func (c CustomFloat64Value) renderValue() ([]byte, error) {
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
