// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func convertSchema(d resource.Resource) (resource_generate.GeneratorResourceSchema, error) {
	var s resource_generate.GeneratorResourceSchema

	attributes := make(map[string]resource_generate.GeneratorAttribute, len(d.Schema.Attributes))
	blocks := make(map[string]resource_generate.GeneratorBlock, len(d.Schema.Blocks))

	for _, v := range d.Schema.Attributes {
		a, err := convertAttribute(v)

		if err != nil {
			return s, err
		}

		attributes[v.Name] = a
	}

	s.Attributes = attributes

	for _, v := range d.Schema.Blocks {
		b, err := convertBlock(v)

		if err != nil {
			return s, err
		}

		blocks[v.Name] = b
	}

	s.Blocks = blocks

	return s, nil
}
