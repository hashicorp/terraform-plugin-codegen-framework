// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"text/template"
)

type ToFromObject struct {
	Name               FrameworkIdentifier
	AssocExtType       *AssocExtType
	AttrTypesToFuncs   map[FrameworkIdentifier]AttrTypesToFuncs
	AttrTypesFromFuncs map[FrameworkIdentifier]string
	templates          map[string]string
}

func NewToFromObject(name string, assocExtType *AssocExtType, attrTypesToFuncs map[string]AttrTypesToFuncs, attrTypesFromFuncs map[string]string) ToFromObject {
	t := map[string]string{
		"from": ObjectFromTemplate,
		"to":   ObjectToTemplate,
	}

	attf := make(map[FrameworkIdentifier]AttrTypesToFuncs, len(attrTypesToFuncs))

	for k, v := range attrTypesToFuncs {
		attf[FrameworkIdentifier(k)] = v
	}

	atff := make(map[FrameworkIdentifier]string, len(attrTypesFromFuncs))

	for k, v := range attrTypesFromFuncs {
		atff[FrameworkIdentifier(k)] = v
	}

	return ToFromObject{
		Name:               FrameworkIdentifier(name),
		AssocExtType:       assocExtType,
		AttrTypesToFuncs:   attf,
		AttrTypesFromFuncs: atff,
		templates:          t,
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
		Name             string
		AssocExtType     *AssocExtType
		AttrTypesToFuncs map[FrameworkIdentifier]AttrTypesToFuncs
	}{
		Name:             o.Name.ToPascalCase(),
		AssocExtType:     o.AssocExtType,
		AttrTypesToFuncs: o.AttrTypesToFuncs,
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
		Name               string
		AssocExtType       *AssocExtType
		AttrTypesFromFuncs map[FrameworkIdentifier]string
	}{
		Name:               o.Name.ToPascalCase(),
		AssocExtType:       o.AssocExtType,
		AttrTypesFromFuncs: o.AttrTypesFromFuncs,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
