// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type ToFromObject struct {
	Name         FrameworkIdentifier
	AssocExtType *AssocExtType
	FromFuncs    map[FrameworkIdentifier]string
	ToFuncs      map[FrameworkIdentifier]string
	templates    map[string]string
}

func NewToFromObject(name string, assocExtType *AssocExtType, toFuncs, fromFuncs map[string]string) ToFromObject {
	t := map[string]string{
		"from": ObjectFromTemplate,
		"to":   ObjectToTemplate,
	}

	tf := make(map[FrameworkIdentifier]string, len(toFuncs))

	for k, v := range toFuncs {
		tf[FrameworkIdentifier(k)] = v
	}

	ff := make(map[FrameworkIdentifier]string, len(fromFuncs))

	for k, v := range fromFuncs {
		ff[FrameworkIdentifier(k)] = v
	}

	return ToFromObject{
		Name:         FrameworkIdentifier(name),
		AssocExtType: assocExtType,
		FromFuncs:    ff,
		ToFuncs:      tf,
		templates:    t,
	}
}

func (o ToFromObject) Render() ([]byte, error) {
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

func (o ToFromObject) renderTo() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(o.templates["to"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name         string
		AssocExtType *AssocExtType
		ToFuncs      map[FrameworkIdentifier]string
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

func (o ToFromObject) renderFrom() ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("").Parse(o.templates["from"])

	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, struct {
		Name         string
		AssocExtType *AssocExtType
		FromFuncs    map[FrameworkIdentifier]string
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
