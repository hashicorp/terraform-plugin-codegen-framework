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
		return fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", GetAttrTypes(e.Object))
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

		switch {
		case v.Bool != nil:
			if v.Bool.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s,", v.Name, v.Bool.CustomType.Type))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.BoolType,", v.Name))
			}
		case v.Float64 != nil:
			if v.Float64.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s,", v.Name, v.Float64.CustomType.Type))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.Float64Type,", v.Name))
			}
		case v.Int64 != nil:
			if v.Int64.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s,", v.Name, v.Int64.CustomType.Type))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.Int64Type,", v.Name))
			}
		case v.List != nil:
			if v.List.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s{\nElemType: %s,\n},", v.Name, v.List.CustomType.Type, GetElementType(v.List.ElementType)))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.ListType{\nElemType: %s,\n},", v.Name, GetElementType(v.List.ElementType)))
			}
		case v.Map != nil:
			if v.Map.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s{\nElemType: %s,\n},", v.Name, v.Map.CustomType.Type, GetElementType(v.Map.ElementType)))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.MapType{\nElemType: %s,\n},", v.Name, GetElementType(v.Map.ElementType)))
			}
		case v.Number != nil:
			if v.Number.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s,", v.Name, v.Number.CustomType.Type))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.NumberType,", v.Name))
			}
		case v.Object != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n},", v.Name, GetAttrTypes(v.Object)))
		case v.Set != nil:
			if v.Set.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s{\nElemType: %s,\n},", v.Name, v.Set.CustomType.Type, GetElementType(v.Set.ElementType)))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.SetType{\nElemType: %s,\n},", v.Name, GetElementType(v.Set.ElementType)))
			}
		case v.String != nil:
			if v.String.CustomType != nil {
				aTypes.WriteString(fmt.Sprintf("\"%s\": %s,", v.Name, v.String.CustomType.Type))
			} else {
				aTypes.WriteString(fmt.Sprintf("\"%s\": types.StringType,", v.Name))
			}
		}
	}

	return aTypes.String()
}

// GetElementTypeImports generates a map of import declarations for use in generated schema with
// collection (i.e., list, map and set) element types.
func GetElementTypeImports(e specschema.ElementType, imports map[string]struct{}) map[string]struct{} {
	switch {
	case e.Bool != nil:
		if e.Bool.CustomType != nil && e.Bool.CustomType.HasImport() {
			imports[*e.Bool.CustomType.Import] = struct{}{}
			return imports
		}
		imports[TypesImport] = struct{}{}
		return imports
	case e.Float64 != nil:
		if e.Float64.CustomType != nil && e.Float64.CustomType.HasImport() {
			imports[*e.Float64.CustomType.Import] = struct{}{}
			return imports
		}
		imports[TypesImport] = struct{}{}
		return imports
	case e.Int64 != nil:
		if e.Int64.CustomType != nil && e.Int64.CustomType.HasImport() {
			imports[*e.Int64.CustomType.Import] = struct{}{}
			return imports
		}
		imports[TypesImport] = struct{}{}
		return imports
	case e.List != nil:
		return GetElementTypeImports(e.List.ElementType, imports)
	case e.Map != nil:
		return GetElementTypeImports(e.Map.ElementType, imports)
	case e.Number != nil:
		if e.Number.CustomType != nil && e.Number.CustomType.HasImport() {
			imports[*e.Number.CustomType.Import] = struct{}{}
			return imports
		}
		imports[TypesImport] = struct{}{}
		return imports
	case e.Object != nil:
		return GetAttrTypesImports(e.Object, imports)
	case e.Set != nil:
		return GetElementTypeImports(e.Set.ElementType, imports)
	case e.String != nil:
		if e.String.CustomType != nil && e.String.CustomType.HasImport() {
			imports[*e.String.CustomType.Import] = struct{}{}
			return imports
		}
		imports[TypesImport] = struct{}{}
		return imports
	default:
		return imports
	}
}

// GetAttrTypesImports generates a map of import declarations for use in generated schema with
// object attribute types.
func GetAttrTypesImports(attrTypes []specschema.ObjectAttributeType, imports map[string]struct{}) map[string]struct{} {
	if len(attrTypes) == 0 {
		return imports
	}

	for _, v := range attrTypes {
		switch {
		case v.Bool != nil:
			if v.Bool.CustomType != nil && v.Bool.CustomType.HasImport() {
				imports[*v.Bool.CustomType.Import] = struct{}{}
				continue
			}
			imports[AttrImport] = struct{}{}
			imports[TypesImport] = struct{}{}
		case v.Float64 != nil:
			if v.Float64.CustomType != nil && v.Float64.CustomType.HasImport() {
				imports[*v.Float64.CustomType.Import] = struct{}{}
				continue
			}
			imports[AttrImport] = struct{}{}
			imports[TypesImport] = struct{}{}
		case v.Int64 != nil:
			if v.Int64.CustomType != nil && v.Int64.CustomType.HasImport() {
				imports[*v.Int64.CustomType.Import] = struct{}{}
				continue
			}
			imports[AttrImport] = struct{}{}
			imports[TypesImport] = struct{}{}
		case v.List != nil:
			if v.List.CustomType != nil && v.List.CustomType.HasImport() {
				imports[*v.List.CustomType.Import] = struct{}{}
			}
			i := GetElementTypeImports(v.List.ElementType, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.Map != nil:
			if v.Map.CustomType != nil && v.Map.CustomType.HasImport() {
				imports[*v.Map.CustomType.Import] = struct{}{}
			}
			i := GetElementTypeImports(v.Map.ElementType, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.Number != nil:
			if v.Number.CustomType != nil && v.Number.CustomType.HasImport() {
				imports[*v.Number.CustomType.Import] = struct{}{}
				continue
			}
			imports[AttrImport] = struct{}{}
			imports[TypesImport] = struct{}{}
		case v.Object != nil:
			i := GetAttrTypesImports(v.Object, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.Set != nil:
			if v.Set.CustomType != nil && v.Set.CustomType.HasImport() {
				imports[*v.Set.CustomType.Import] = struct{}{}
			}
			i := GetElementTypeImports(v.Set.ElementType, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.String != nil:
			if v.String.CustomType != nil && v.String.CustomType.HasImport() {
				imports[*v.Float64.CustomType.Import] = struct{}{}
				continue
			}
			imports[AttrImport] = struct{}{}
			imports[TypesImport] = struct{}{}
		}
	}

	return imports
}
