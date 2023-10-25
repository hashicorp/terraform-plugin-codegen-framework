// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type ToFromList struct {
	Name             FrameworkIdentifier
	AssocExtType     *AssocExtType
	ElementTypeType  string
	ElementTypeValue string
	ElementFrom      string
	templates        map[string]string
}

func NewToFromList(name string, assocExtType *AssocExtType, elemTypeType, elemTypeValue, elemFrom string) ToFromList {
	t := map[string]string{
		"from": ListFromTemplate,
		"to":   ListToTemplate,
	}

	return ToFromList{
		Name:             FrameworkIdentifier(name),
		AssocExtType:     assocExtType,
		ElementTypeType:  elemTypeType,
		ElementTypeValue: elemTypeValue,
		ElementFrom:      elemFrom,
		templates:        t,
	}
}

func (o ToFromList) Render() ([]byte, error) {
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

func (o ToFromList) renderTo() ([]byte, error) {
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

func (o ToFromList) renderFrom() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(o.templates["from"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name             string
		AssocExtType     *AssocExtType
		ElementTypeType  string
		ElementTypeValue string
		ElementFrom      string
	}{
		Name:             o.Name.ToPascalCase(),
		AssocExtType:     o.AssocExtType,
		ElementTypeType:  o.ElementTypeType,
		ElementTypeValue: o.ElementTypeValue,
		ElementFrom:      o.ElementFrom,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
