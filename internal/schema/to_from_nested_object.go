// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type ToFromNestedObject struct {
	Name         FrameworkIdentifier
	AssocExtType *AssocExtType
	ToFuncs      map[FrameworkIdentifier]ToFromConversion
	FromFuncs    map[FrameworkIdentifier]ToFromConversion
	templates    map[string]string
}

func NewToFromNestedObject(name string, assocExtType *AssocExtType, toFuncs, fromFuncs map[string]ToFromConversion) ToFromNestedObject {
	t := map[string]string{
		"from": NestedObjectFromTemplate,
		"to":   NestedObjectToTemplate,
	}

	tf := make(map[FrameworkIdentifier]ToFromConversion, len(toFuncs))

	for k, v := range toFuncs {
		tf[FrameworkIdentifier(k)] = v
	}

	ff := make(map[FrameworkIdentifier]ToFromConversion, len(fromFuncs))

	for k, v := range fromFuncs {
		ff[FrameworkIdentifier(k)] = v
	}

	return ToFromNestedObject{
		Name:         FrameworkIdentifier(name),
		AssocExtType: assocExtType,
		FromFuncs:    ff,
		ToFuncs:      tf,
		templates:    t,
	}
}

func (o ToFromNestedObject) Render() ([]byte, error) {
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

func (o ToFromNestedObject) renderTo() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(o.templates["to"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name         string
		AssocExtType *AssocExtType
		ToFuncs      map[FrameworkIdentifier]ToFromConversion
	}{
		Name:         o.Name.ToPascalCase(),
		AssocExtType: o.AssocExtType,
		ToFuncs:      o.ToFuncs,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o ToFromNestedObject) renderFrom() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(o.templates["from"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name         string
		AssocExtType *AssocExtType
		FromFuncs    map[FrameworkIdentifier]ToFromConversion
	}{
		Name:         o.Name.ToPascalCase(),
		AssocExtType: o.AssocExtType,
		FromFuncs:    o.FromFuncs,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
