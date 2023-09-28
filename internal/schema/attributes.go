// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
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
		switch g[k].GeneratorSchemaType() {
		case GeneratorBoolAttribute:
			attrTypes[k] = "basetypes.BoolType{}"
		case GeneratorFloat64Attribute:
			attrTypes[k] = "basetypes.Float64Type{}"
		case GeneratorInt64Attribute:
			attrTypes[k] = "basetypes.Int64Type{}"
		case GeneratorListAttribute:
			if e, ok := g[k].(Elements); ok {
				elemType, err := ElementTypeString(e.ElemType())
				if err != nil {
					return nil, err
				}
				attrTypes[k] = fmt.Sprintf("basetypes.ListType{\nElemType: %s,\n}", elemType)
			} else {
				return nil, fmt.Errorf("%s attribute is a ListType but does not implement Elements interface", k)
			}
		case GeneratorListNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.ListType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorMapAttribute:
			if e, ok := g[k].(Elements); ok {
				elemType, err := ElementTypeString(e.ElemType())
				if err != nil {
					return nil, err
				}
				attrTypes[k] = fmt.Sprintf("basetypes.MapType{\nElemType: %s,\n}", elemType)
			} else {
				return nil, fmt.Errorf("%s attribute is a MapType but does not implement Elements interface", k)
			}
		case GeneratorMapNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.MapType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorNumberAttribute:
			attrTypes[k] = "basetypes.NumberType{}"
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
		case GeneratorSetAttribute:
			if e, ok := g[k].(Elements); ok {
				elemType, err := ElementTypeString(e.ElemType())
				if err != nil {
					return nil, err
				}
				attrTypes[k] = fmt.Sprintf("basetypes.SetType{\nElemType: %s,\n}", elemType)
			} else {
				return nil, fmt.Errorf("%s attribute is a SetType but does not implement Elements interface", k)
			}
		case GeneratorSetNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.SetType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorSingleNestedAttribute:
			attrTypes[k] = fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorStringAttribute:
			attrTypes[k] = "basetypes.StringType{}"
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

func (g GeneratorAttributes) SortedKeys() []string {
	var attributeKeys = make([]string, 0, len(g))

	for k := range g {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	return attributeKeys
}
