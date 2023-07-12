// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

// GetElementType generates the strings for use within templates for specifying the types to use with
// collection (i.e., list, map and set) element types.
func GetElementType(e specschema.ElementType) string {
	switch {
	case e.Bool != nil:
		if e.Bool.CustomType != nil {
			return e.Bool.CustomType.Type
		}
		return "types.BoolType"
	case e.Float64 != nil:
		if e.Float64.CustomType != nil {
			return e.Float64.CustomType.Type
		}
		return "types.Float64Type"
	case e.Int64 != nil:
		if e.Int64.CustomType != nil {
			return e.Int64.CustomType.Type
		}
		return "types.Int64Type"
	case e.List != nil:
		if e.List.CustomType != nil {
			return fmt.Sprintf("%s{\nElemType: %s,\n}", e.List.CustomType.Type, GetElementType(e.List.ElementType))
		}
		return fmt.Sprintf("types.ListType{\nElemType: %s,\n}", GetElementType(e.List.ElementType))
	case e.Map != nil:
		if e.Map.CustomType != nil {
			return fmt.Sprintf("%s{\nElemType: %s,\n}", e.Map.CustomType.Type, GetElementType(e.Map.ElementType))
		}
		return fmt.Sprintf("types.MapType{\nElemType: %s,\n}", GetElementType(e.Map.ElementType))
	case e.Number != nil:
		if e.Number.CustomType != nil {
			return e.Number.CustomType.Type
		}
		return "types.NumberType"
	case e.Object != nil:
		if e.Object.CustomType != nil {
			return fmt.Sprintf("%s{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", e.Object.CustomType.Type, GetAttrTypes(e.Object.AttributeTypes))
		}
		return fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", GetAttrTypes(e.Object.AttributeTypes))
	case e.Set != nil:
		if e.Set.CustomType != nil {
			return fmt.Sprintf("%s{\nElemType: %s,\n}", e.Set.CustomType.Type, GetElementType(e.Set.ElementType))
		}
		return fmt.Sprintf("types.SetType{\nElemType: %s,\n}", GetElementType(e.Set.ElementType))
	case e.String != nil:
		if e.String.CustomType != nil {
			return e.String.CustomType.Type
		}
		return "types.StringType"
	}

	return ""
}
