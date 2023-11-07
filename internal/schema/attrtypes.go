// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"errors"
	"fmt"
	"strings"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

// GetAttrTypes generates the strings for use within templates for specifying the types to use with
// object attribute types.
func GetAttrTypes(attrTypes specschema.ObjectAttributeTypes) string {
	var aTypes strings.Builder

	for _, v := range attrTypes {
		if aTypes.Len() > 0 {
			aTypes.WriteString("\n")
		}

		aTypes.WriteString(fmt.Sprintf("%q: ", v.Name))

		switch {
		case v.Bool != nil:
			if v.Bool.CustomType != nil {
				aTypes.WriteString(v.Bool.CustomType.Type)
			} else {
				aTypes.WriteString("types.BoolType")
			}
		case v.Float64 != nil:
			if v.Float64.CustomType != nil {
				aTypes.WriteString(v.Float64.CustomType.Type)
			} else {
				aTypes.WriteString("types.Float64Type")
			}
		case v.Int64 != nil:
			if v.Int64.CustomType != nil {
				aTypes.WriteString(v.Int64.CustomType.Type)
			} else {
				aTypes.WriteString("types.Int64Type")
			}
		case v.List != nil:
			if v.List.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("%s{\nElemType: %s,\n}", v.List.CustomType.Type, GetElementType(v.List.ElementType)))
			} else {
				aTypes.WriteString(fmt.Sprintf("types.ListType{\nElemType: %s,\n}", GetElementType(v.List.ElementType)))
			}
		case v.Map != nil:
			if v.Map.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("%s{\nElemType: %s,\n}", v.Map.CustomType.Type, GetElementType(v.Map.ElementType)))
			} else {
				aTypes.WriteString(fmt.Sprintf("types.MapType{\nElemType: %s,\n}", GetElementType(v.Map.ElementType)))
			}
		case v.Number != nil:
			if v.Number.CustomType != nil {
				aTypes.WriteString(v.Number.CustomType.Type)
			} else {
				aTypes.WriteString("types.NumberType")
			}
		case v.Object != nil:
			if v.Object.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("%s{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", v.Object.CustomType.Type, GetAttrTypes(v.Object.AttributeTypes)))
			} else {
				aTypes.WriteString(fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", GetAttrTypes(v.Object.AttributeTypes)))
			}
		case v.Set != nil:
			if v.Set.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("%s{\nElemType: %s,\n}", v.Set.CustomType.Type, GetElementType(v.Set.ElementType)))
			} else {
				aTypes.WriteString(fmt.Sprintf("types.SetType{\nElemType: %s,\n}", GetElementType(v.Set.ElementType)))
			}
		case v.String != nil:
			if v.String.CustomType != nil {
				aTypes.WriteString(v.String.CustomType.Type)
			} else {
				aTypes.WriteString("types.StringType")
			}
		}

		aTypes.WriteString(",")
	}

	return aTypes.String()
}

type AttrTypesToFuncs struct {
	AttrValue string
	ToFunc    string
}

// GetAttrTypesToFuncs returns string representations of the function that is used
// for converting to an API Go type from a framework type.
// TODO: Handle custom type, and types other than primitives.
func GetAttrTypesToFuncs(a specschema.ObjectAttributeTypes) (map[string]AttrTypesToFuncs, error) {
	attrTypesFuncs := make(map[string]AttrTypesToFuncs, len(a))

	for _, v := range a {
		switch {
		case v.Bool != nil:
			attrTypesFuncs[v.Name] = AttrTypesToFuncs{
				AttrValue: "types.Bool",
				ToFunc:    "ValueBoolPointer",
			}
		case v.Float64 != nil:
			attrTypesFuncs[v.Name] = AttrTypesToFuncs{
				AttrValue: "types.Float64",
				ToFunc:    "ValueFloat64Pointer",
			}
		case v.Int64 != nil:
			attrTypesFuncs[v.Name] = AttrTypesToFuncs{
				AttrValue: "types.Int64",
				ToFunc:    "ValueInt64Pointer",
			}
		case v.List != nil:
			return nil, NewUnimplementedError(errors.New("list attribute type is not yet implemented"))
		case v.Map != nil:
			return nil, NewUnimplementedError(errors.New("map attribute type is not yet implemented"))

		case v.Number != nil:
			attrTypesFuncs[v.Name] = AttrTypesToFuncs{
				AttrValue: "types.Number",
				ToFunc:    "ValueBigFloat",
			}
		case v.Object != nil:
			return nil, NewUnimplementedError(errors.New("object attribute type is not yet implemented"))
		case v.Set != nil:
			return nil, NewUnimplementedError(errors.New("set attribute type is not yet implemented"))
		case v.String != nil:
			attrTypesFuncs[v.Name] = AttrTypesToFuncs{
				AttrValue: "types.String",
				ToFunc:    "ValueStringPointer",
			}
		}
	}

	return attrTypesFuncs, nil
}

// GetAttrTypesFromFuncs returns string representations of the function that is used
// for converting from an API Go type to a framework type.
// TODO: Handle custom type, and types other than primitives.
func GetAttrTypesFromFuncs(a specschema.ObjectAttributeTypes) (map[string]string, error) {
	attrTypesFuncs := make(map[string]string, len(a))

	for _, v := range a {
		switch {
		case v.Bool != nil:
			attrTypesFuncs[v.Name] = "types.BoolPointerValue"
		case v.Float64 != nil:
			attrTypesFuncs[v.Name] = "types.Float64PointerValue"
		case v.Int64 != nil:
			attrTypesFuncs[v.Name] = "types.Int64PointerValue"
		case v.List != nil:
			return nil, NewUnimplementedError(errors.New("list attribute type is not yet implemented"))
		case v.Map != nil:
			return nil, NewUnimplementedError(errors.New("map attribute type is not yet implemented"))
		case v.Number != nil:
			attrTypesFuncs[v.Name] = "types.NumberValue"
		case v.Object != nil:
			return nil, NewUnimplementedError(errors.New("object attribute type is not yet implemented"))
		case v.Set != nil:
			return nil, NewUnimplementedError(errors.New("set attribute type is not yet implemented"))
		case v.String != nil:
			attrTypesFuncs[v.Name] = "types.StringPointerValue"
		}
	}

	return attrTypesFuncs, nil
}
