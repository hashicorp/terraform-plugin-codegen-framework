// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type ToFromInt64 struct {
	Name         FrameworkIdentifier
	AssocExtType *AssocExtType
	templates    map[string]string
}

func NewToFromInt64(name string, assocExtType *AssocExtType) ToFromInt64 {
	t := map[string]string{
		"from": Int64FromTemplate,
		"to":   Int64ToTemplate,
	}

	return ToFromInt64{
		Name:         FrameworkIdentifier(name),
		AssocExtType: assocExtType,
		templates:    t,
	}
}

func (o ToFromInt64) Render() ([]byte, error) {
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

func (o ToFromInt64) renderTo() ([]byte, error) {
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

func (o ToFromInt64) renderFrom() ([]byte, error) {
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
