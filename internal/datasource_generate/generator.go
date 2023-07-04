// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"sort"
)

type DataSourcesModelObjectHelpersGenerator struct {
}

func NewDataSourcesModelObjectHelpersGenerator() DataSourcesModelObjectHelpersGenerator {
	return DataSourcesModelObjectHelpersGenerator{}
}

func (d DataSourcesModelObjectHelpersGenerator) Process(schemas map[string]GeneratorDataSourceSchema) (map[string][]byte, error) {
	dataSourcesModelObjectHelpers := make(map[string][]byte, len(schemas))

	for name, s := range schemas {
		var buf bytes.Buffer

		g := GeneratorDataSourceSchema{
			Attributes: s.Attributes,
			Blocks:     s.Blocks,
		}

		var attributeKeys = make([]string, 0, len(g.Attributes))

		for k := range g.Attributes {
			attributeKeys = append(attributeKeys, k)
		}

		sort.Strings(attributeKeys)

		for _, k := range attributeKeys {
			if g.Attributes[k] == nil {
				continue
			}

			switch t := g.Attributes[k].(type) {
			case GeneratorListNestedAttribute:
				var hasNestedAttribute bool

				for _, v := range t.NestedObject.Attributes {
					switch v.(type) {
					case GeneratorListNestedAttribute:
						hasNestedAttribute = true
						break
					}
				}

				if hasNestedAttribute {
					modelObjectHelpers, err := t.ModelObjectHelpersString(k)

					if err != nil {
						return nil, err
					}

					buf.WriteString(modelObjectHelpers)
				}
			}
		}

		dataSourcesModelObjectHelpers[name] = buf.Bytes()
	}

	return dataSourcesModelObjectHelpers, nil
}
