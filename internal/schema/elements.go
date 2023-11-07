// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"errors"
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

// GetElementValueType generates the strings for use within templates for specifying the value types
// to use with collection (i.e., list, map and set) element types.
func GetElementValueType(e specschema.ElementType) string {
	switch {
	case e.Bool != nil:
		if e.Bool.CustomType != nil {
			return e.Bool.CustomType.ValueType
		}
		return "types.Bool"
	case e.Float64 != nil:
		if e.Float64.CustomType != nil {
			return e.Float64.CustomType.ValueType
		}
		return "types.Float64"
	case e.Int64 != nil:
		if e.Int64.CustomType != nil {
			return e.Int64.CustomType.ValueType
		}
		return "types.Int64"
	case e.List != nil:
		if e.List.CustomType != nil {
			return e.List.CustomType.ValueType
		}
		return "types.List"
	case e.Map != nil:
		if e.Map.CustomType != nil {
			return e.Map.CustomType.ValueType
		}
		return "types.Map"
	case e.Number != nil:
		if e.Number.CustomType != nil {
			return e.Number.CustomType.ValueType
		}
		return "types.Number"
	case e.Object != nil:
		if e.Object.CustomType != nil {
			return e.Object.CustomType.ValueType
		}
		return "types.Object"
	case e.Set != nil:
		if e.Set.CustomType != nil {
			return e.Set.CustomType.ValueType
		}
		return "types.Set"
	case e.String != nil:
		if e.String.CustomType != nil {
			return e.String.CustomType.ValueType
		}
		return "types.String"
	}

	return ""
}

// GetElementFromFunc returns a string representation of the function that is used
// for converting from an API Go type to a framework type.
// TODO: Handle custom type, and types other than primitives.
func GetElementFromFunc(e specschema.ElementType) (string, error) {
	switch {
	case e.Bool != nil:
		return "types.BoolPointerValue", nil
	case e.Float64 != nil:
		return "types.Float64PointerValue", nil
	case e.Int64 != nil:
		return "types.Int64PointerValue", nil
	case e.List != nil:
		return "", NewUnimplementedError(errors.New("list element type is not yet implemented"))
	case e.Map != nil:
		return "", NewUnimplementedError(errors.New("map element type is not yet implemented"))
	case e.Number != nil:
		return "types.NumberValue", nil
	case e.Object != nil:
		return "", NewUnimplementedError(errors.New("object element type is not yet implemented"))
	case e.Set != nil:
		return "", NewUnimplementedError(errors.New("set element type is not yet implemented"))
	case e.String != nil:
		return "types.StringPointerValue", nil
	}

	return "", nil
}
