// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type ToFromBool struct {
	Name         FrameworkIdentifier
	AssocExtType *AssocExtType
	templates    map[string]string
}

func NewToFromBool(name string, assocExtType *AssocExtType) ToFromBool {
	t := map[string]string{
		"from": BoolFromTemplate,
		"to":   BoolToTemplate,
	}

	return ToFromBool{
		Name:         FrameworkIdentifier(name),
		AssocExtType: assocExtType,
		templates:    t,
	}
}

func (o ToFromBool) Render() ([]byte, error) {
	var buf bytes.Buffer

	renderFuncs := []func() ([]byte, error){
		o.renderTo,
		o.renderFrom,
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

func (o ToFromBool) renderTo() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(o.templates["to"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name         string
		AssocExtType *AssocExtType
	}{
		Name:         o.Name.ToPascalCase(),
		AssocExtType: o.AssocExtType,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o ToFromBool) renderFrom() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(o.templates["from"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name         string
		AssocExtType *AssocExtType
	}{
		Name:         o.Name.ToPascalCase(),
		AssocExtType: o.AssocExtType,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
