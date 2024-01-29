// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"fmt"
	"sort"
	"strings"
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
		name := FrameworkIdentifier(k)

		switch g[k].GeneratorSchemaType() {
		case GeneratorListNestedBlock:
			attrTypes[k] = fmt.Sprintf("basetypes.ListType{\nElemType: %sValue{}.Type(ctx),\n}", name.ToPascalCase())
		case GeneratorSetNestedBlock:
			attrTypes[k] = fmt.Sprintf("basetypes.SetType{\nElemType: %sValue{}.Type(ctx),\n}", name.ToPascalCase())
		case GeneratorSingleNestedBlock:
			attrTypes[k] = fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}", name.ToPascalCase())
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

func (g GeneratorBlocks) Equal(other GeneratorBlocks) bool {
	if len(g) != len(other) {
		return false
	}

	for k, v := range g {
		otherBlock, ok := other[k]

		if !ok {
			return false
		}

		if !v.Equal(otherBlock) {
			return false
		}
	}

	return true
}

// FromFuncs returns a mapping of block names to string representations of the
// function that converts a Go value to a framework value.
func (g GeneratorBlocks) FromFuncs() map[string]string {
	attributeKeys := g.SortedKeys()

	fromFuncs := make(map[string]string, len(g))

	for _, k := range attributeKeys {
		switch g[k].GeneratorSchemaType() {
		case GeneratorBoolAttribute:
			fromFuncs[k] = "BoolPointerValue"
		case GeneratorFloat64Attribute:
			fromFuncs[k] = "Float64PointerValue"
		case GeneratorInt64Attribute:
			fromFuncs[k] = "Int64PointerValue"
		case GeneratorNumberAttribute:
			fromFuncs[k] = "NumberValue"
		case GeneratorStringAttribute:
			fromFuncs[k] = "StringPointerValue"
		}
	}

	return fromFuncs
}

func (g GeneratorBlocks) Imports() *Imports {
	imports := NewImports()

	for _, v := range g {
		imports.Append(v.Imports())
	}

	return imports
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

		if !strings.HasPrefix(str, "\n") {
			str = "\n" + str
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

// ToFuncs returns a mapping of block names to string representations of the
// function that converts a framework value to a Go value.
func (g GeneratorBlocks) ToFuncs() map[string]string {
	attributeKeys := g.SortedKeys()

	toFuncs := make(map[string]string, len(g))

	for _, k := range attributeKeys {
		switch g[k].GeneratorSchemaType() {
		case GeneratorBoolAttribute:
			toFuncs[k] = "ValueBoolPointer"
		case GeneratorFloat64Attribute:
			toFuncs[k] = "ValueFloat64Pointer"
		case GeneratorInt64Attribute:
			toFuncs[k] = "ValueInt64Pointer"
		case GeneratorNumberAttribute:
			toFuncs[k] = "ValueBigFloat"
		case GeneratorStringAttribute:
			toFuncs[k] = "ValueStringPointer"
		}
	}

	return toFuncs
}

func (g GeneratorBlocks) SortedKeys() []string {
	var blockKeys = make([]string, 0, len(g))

	for k := range g {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	return blockKeys
}
