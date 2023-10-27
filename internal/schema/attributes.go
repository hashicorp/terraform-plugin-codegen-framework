// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"fmt"
	"sort"
	"strings"
)

type GeneratorAttributes map[string]GeneratorAttribute

// AttributeTypes returns a mapping of attribute names to string representations of the
// attribute type.
func (g GeneratorAttributes) AttributeTypes() (map[string]string, error) {
	attributeKeys := g.SortedKeys()

	attributeTypes := make(map[string]string, len(g))

	for _, k := range attributeKeys {
		switch g[k].GeneratorSchemaType() {
		case GeneratorBoolAttribute:
			attributeTypes[k] = "Bool"
		case GeneratorFloat64Attribute:
			attributeTypes[k] = "Float64"
		case GeneratorInt64Attribute:
			attributeTypes[k] = "Int64"
		case GeneratorListAttribute:
			attributeTypes[k] = "List"
		case GeneratorListNestedAttribute:
			attributeTypes[k] = "ListNested"
		case GeneratorMapAttribute:
			attributeTypes[k] = "Map"
		case GeneratorMapNestedAttribute:
			attributeTypes[k] = "MapNested"
		case GeneratorNumberAttribute:
			attributeTypes[k] = "Number"
		case GeneratorObjectAttribute:
			attributeTypes[k] = "Object"
		case GeneratorSetAttribute:
			attributeTypes[k] = "Set"
		case GeneratorSetNestedAttribute:
			attributeTypes[k] = "SetNested"
		case GeneratorSingleNestedAttribute:
			attributeTypes[k] = "SingleNested"
		case GeneratorStringAttribute:
			attributeTypes[k] = "String"
		}
	}

	return attributeTypes, nil
}

// AttrTypes returns a mapping of attribute names to string representations of the
// underlying attr.Type.
func (g GeneratorAttributes) AttrTypes() (map[string]string, error) {
	attributeKeys := g.SortedKeys()

	attrTypes := make(map[string]string, len(g))

	for _, k := range attributeKeys {
		name := FrameworkIdentifier(k)

		if a, ok := g[k].(AttrType); ok {
			attrType, err := a.AttrType(name)

			if err != nil {
				return nil, err
			}

			attrTypes[k] = attrType

			continue
		}

		switch g[k].GeneratorSchemaType() {
		case GeneratorListNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.ListType{\nElemType: %sValue{}.Type(ctx),\n}", name.ToPascalCase())
		case GeneratorMapNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.MapType{\nElemType: %sValue{}.Type(ctx),\n}", name.ToPascalCase())
		case GeneratorObjectAttribute:
			if o, ok := g[k].(Attrs); ok {
				aTypes, err := AttrTypesString(o.AttrTypes())
				if err != nil {
					return nil, err
				}
				attrTypes[k] = fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", aTypes)
			} else {
				return nil, fmt.Errorf("%s attribute is an ObjectType but does not implement Attrs interface", k)
			}
		case GeneratorSetNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.SetType{\nElemType: %sValue{}.Type(ctx),\n}", name.ToPascalCase())
		case GeneratorSingleNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}", name.ToPascalCase())
		}
	}

	return attrTypes, nil
}

// AttrValues returns a mapping of attribute names to string representations of the
// underlying attr.Value.
func (g GeneratorAttributes) AttrValues() (map[string]string, error) {
	attributeKeys := g.SortedKeys()

	attrValues := make(map[string]string, len(g))

	for _, k := range attributeKeys {
		if a, ok := g[k].(AttrValue); ok {
			attrValues[k] = a.AttrValue(FrameworkIdentifier(k))
			continue
		}

		switch g[k].GeneratorSchemaType() {
		case GeneratorListNestedAttribute:
			attrValues[k] = "basetypes.ListValue"
		case GeneratorMapNestedAttribute:
			attrValues[k] = "basetypes.MapValue"
		case GeneratorObjectAttribute:
			attrValues[k] = "basetypes.ObjectValue"
		case GeneratorSetNestedAttribute:
			attrValues[k] = "basetypes.SetValue"
		case GeneratorSingleNestedAttribute:
			attrValues[k] = "basetypes.ObjectValue"
		}
	}

	return attrValues, nil
}

// CollectionTypes returns a mapping of attribute names to string representations of the
// element type (e.g., types.BoolType), and type value function (e.g., types.ListValue)
// for collection types that do not have an associated external type.
func (g GeneratorAttributes) CollectionTypes() (map[string]map[string]string, error) {
	attributeKeys := g.SortedKeys()

	collectionTypes := make(map[string]map[string]string, len(g))

	for _, k := range attributeKeys {
		c, ok := g[k].(CollectionType)

		if !ok {
			continue
		}

		collectionType, err := c.CollectionType()

		if err != nil {
			return nil, err
		}

		if collectionType == nil {
			continue
		}

		collectionTypes[k] = collectionType
	}

	return collectionTypes, nil
}

// FromFuncs returns a mapping of attribute names to string representations of the
// function that converts a Go value to a framework value.
func (g GeneratorAttributes) FromFuncs() map[string]ToFromConversion {
	attributeKeys := g.SortedKeys()

	fromFuncs := make(map[string]ToFromConversion, len(g))

	for _, k := range attributeKeys {
		if a, ok := g[k].(From); ok {
			fromFuncs[k], _ = a.From()
		}
	}

	return fromFuncs
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

		str, err := g[k].Schema(FrameworkIdentifier(k))

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

// ToFuncs returns a mapping of attribute names to string representations of the
// function that converts a framework value to a Go value.
func (g GeneratorAttributes) ToFuncs() map[string]ToFromConversion {
	attributeKeys := g.SortedKeys()

	toFuncs := make(map[string]ToFromConversion, len(g))

	for _, k := range attributeKeys {
		if a, ok := g[k].(To); ok {
			toFuncs[k], _ = a.To()
		}
	}

	return toFuncs
}

func (g GeneratorAttributes) SortedKeys() []string {
	var attributeKeys = make([]string, 0, len(g))

	for k := range g {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	return attributeKeys
}
