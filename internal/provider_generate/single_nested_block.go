// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

type GeneratorSingleNestedBlock struct {
	schema.SingleNestedBlock

	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

func (g GeneratorSingleNestedBlock) Equal(ga GeneratorBlock) bool {
	h, ok := ga.(GeneratorSingleNestedBlock)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	for k, a := range g.Attributes {
		if !a.Equal(h.Attributes[k]) {
			return false
		}
	}

	return true
}

func (g GeneratorSingleNestedBlock) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": attributeStringsFromGeneratorAttributes,
		"getBlocks":     blockStringsFromGeneratorBlocks,
	}

	t, err := template.New("single_nested_block").Funcs(funcMap).Parse(singleNestedBlockGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_block").Parse(commonBlockGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSingleNestedBlock{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSingleNestedBlock) validatorsEqual(x, y []specschema.ObjectValidator) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	//TODO: Sort before comparing.
	for k, v := range x {
		if v.Custom == nil && y[k].Custom != nil {
			return false
		}

		if v.Custom != nil && y[k].Custom == nil {
			return false
		}

		if v.Custom != nil && y[k].Custom != nil {
			if *v.Custom.Import != *y[k].Custom.Import {
				return false
			}
		}

		if v.Custom.SchemaDefinition != y[k].Custom.SchemaDefinition {
			return false
		}
	}

	return true
}
