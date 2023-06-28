// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"fmt"
	"strings"

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

// GetAttrTypes generates the strings for use within templates for specifying the types to use with
// object attribute types.
func GetAttrTypes(attrTypes []specschema.ObjectAttributeType) string {
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
			aTypes.WriteString(fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", GetAttrTypes(v.Object.AttributeTypes)))
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
