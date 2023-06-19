// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func convertSchema(d datasource.DataSource) (datasource_generate.GeneratorDataSourceSchema, error) {
	var s datasource_generate.GeneratorDataSourceSchema

	attributes := make(map[string]datasource_generate.GeneratorAttribute, len(d.Schema.Attributes))
	blocks := make(map[string]datasource_generate.GeneratorBlock, len(d.Schema.Blocks))

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
