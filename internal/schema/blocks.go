// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

type GeneratorBlocks map[string]GeneratorBlock

// BlockTypes returns a mapping of block names to string representations of the
// block type.
func (g GeneratorBlocks) BlockTypes() (map[string]string, error) {
	blockKeys := g.SortedKeys()

	blockTypes := make(map[string]string, len(g))

	for _, k := range blockKeys {
		switch g[k].GeneratorSchemaType() {
		case GeneratorListNestedBlock:
			blockTypes[k] = "ListNested"
		case GeneratorSetNestedBlock:
			blockTypes[k] = "SetNested"
		case GeneratorSingleNestedBlock:
			blockTypes[k] = "SingleNested"
		}
	}

	return blockTypes, nil
}

// AttrTypes returns a mapping of block names to string representations of the
// underlying attr.Type.
func (g GeneratorBlocks) AttrTypes() (map[string]string, error) {
	blockKeys := g.SortedKeys()

	attrTypes := make(map[string]string, len(g))

	for _, k := range blockKeys {
		switch g[k].GeneratorSchemaType() {
		case GeneratorListNestedBlock:
			attrTypes[k] = fmt.Sprintf("basetypes.ListType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorSetNestedBlock:
			attrTypes[k] = fmt.Sprintf("basetypes.SetType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorSingleNestedBlock:
			attrTypes[k] = fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k))
		}
	}

	return attrTypes, nil
}

// AttrValues returns a mapping of block names to string representations of the
// underlying attr.Value.
func (g GeneratorBlocks) AttrValues() (map[string]string, error) {
	blockKeys := g.SortedKeys()

	attrValues := make(map[string]string, len(g))

	for _, k := range blockKeys {
		switch g[k].GeneratorSchemaType() {
		case GeneratorListNestedBlock:
			attrValues[k] = "basetypes.ListValue"
		case GeneratorSetNestedBlock:
			attrValues[k] = "basetypes.SetValue"
		case GeneratorSingleNestedBlock:
			attrValues[k] = "basetypes.ObjectValue"
		}
	}

	return attrValues, nil
}

func (g GeneratorBlocks) Schema() (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee block order as maps are unordered in Go.
	var keys = make([]string, 0, len(g))

	for k := range g {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if g[k] == nil {
			continue
		}

		str, err := g[k].Schema(FrameworkIdentifier(k))

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

func (g GeneratorBlocks) SortedKeys() []string {
	var blockKeys = make([]string, 0, len(g))

	for k := range g {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	return blockKeys
}
