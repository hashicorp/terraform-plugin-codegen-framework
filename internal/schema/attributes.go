// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"sort"
	"strings"
)

type GeneratorAttributes map[string]GeneratorAttribute

// AttrValues returns the attribute name and a string representation of the
// underlying attr.Value.
func (g GeneratorAttributes) AttrValues() (map[string]string, error) {
	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g))

	for k := range g {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	attrValues := make(map[string]string)

	for _, k := range attributeKeys {
		switch g[k].GeneratorSchemaType() {
		case GeneratorBoolAttribute:
			attrValues[k] = "basetypes.BoolValue"
		case GeneratorFloat64Attribute:
			attrValues[k] = "basetypes.Float64Value"
		case GeneratorInt64Attribute:
			attrValues[k] = "basetypes.Int64Value"
		case GeneratorListAttribute:
			attrValues[k] = "basetypes.ListValue"
		case GeneratorListNestedAttribute:
			attrValues[k] = "basetypes.ListValue"
		case GeneratorMapAttribute:
			attrValues[k] = "basetypes.MapValue"
		case GeneratorMapNestedAttribute:
			attrValues[k] = "basetypes.MapValue"
		case GeneratorNumberAttribute:
			attrValues[k] = "basetypes.NumberValue"
		case GeneratorObjectAttribute:
			attrValues[k] = "basetypes.ObjectValue"
		case GeneratorSetAttribute:
			attrValues[k] = "basetypes.SetValue"
		case GeneratorSetNestedAttribute:
			attrValues[k] = "basetypes.SetValue"
		case GeneratorSingleNestedAttribute:
			attrValues[k] = "basetypes.ObjectValue"
		case GeneratorStringAttribute:
			attrValues[k] = "basetypes.StringValue"
		}
	}

	return attrValues, nil
}

func (g GeneratorAttributes) Schema() (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var keys = make([]string, 0, len(g))

	for k := range g {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if g[k] == nil {
			continue
		}

		str, err := g[k].Schema(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}
