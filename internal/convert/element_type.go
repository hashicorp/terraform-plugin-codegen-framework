// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type ElementType struct {
	elementType specschema.ElementType
}

func NewElementType(e specschema.ElementType) ElementType {
	return ElementType{
		elementType: e,
	}
}

func (e ElementType) ElementType() []byte {
	var b bytes.Buffer

	switch {
	case e.elementType.Bool != nil:
		if e.elementType.Bool.CustomType != nil {
			b.WriteString(e.elementType.Bool.CustomType.Type)
		} else {
			b.WriteString("types.BoolType")
		}
	case e.elementType.Float64 != nil:
		if e.elementType.Float64.CustomType != nil {
			b.WriteString(e.elementType.Float64.CustomType.Type)
		} else {
			b.WriteString("types.Float64Type")
		}
	case e.elementType.Int64 != nil:
		if e.elementType.Int64.CustomType != nil {
			b.WriteString(e.elementType.Int64.CustomType.Type)
		} else {
			b.WriteString("types.Int64Type")
		}
	case e.elementType.List != nil:
		if e.elementType.List.CustomType != nil {
			b.WriteString(e.elementType.List.CustomType.Type)
		} else {
			b.WriteString(fmt.Sprintf("types.ListType{\nElemType: %s,\n}", NewElementType(e.elementType.List.ElementType).ElementType()))
		}
	case e.elementType.Map != nil:
		if e.elementType.Map.CustomType != nil {
			b.WriteString(e.elementType.Map.CustomType.Type)
		} else {
			b.WriteString(fmt.Sprintf("types.MapType{\nElemType: %s,\n}", NewElementType(e.elementType.Map.ElementType).ElementType()))
		}
	case e.elementType.Number != nil:
		if e.elementType.Number.CustomType != nil {
			b.WriteString(e.elementType.Number.CustomType.Type)
		} else {
			b.WriteString("types.NumberType")
		}
	case e.elementType.Object != nil:
		if e.elementType.Object.CustomType != nil {
			b.WriteString(e.elementType.Object.CustomType.Type)
		} else {
			b.WriteString(fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", NewObjectAttributeTypes(e.elementType.Object.AttributeTypes).AttributeTypes()))
		}
	case e.elementType.Set != nil:
		if e.elementType.Set.CustomType != nil {
			b.WriteString(e.elementType.Set.CustomType.Type)
		} else {
			b.WriteString(fmt.Sprintf("types.SetType{\nElemType: %s,\n}", NewElementType(e.elementType.Set.ElementType).ElementType()))
		}
	case e.elementType.String != nil:
		if e.elementType.String.CustomType != nil {
			b.WriteString(e.elementType.String.CustomType.Type)
		} else {
			b.WriteString("types.StringType")
		}
	}

	return b.Bytes()
}

func (e ElementType) Equal(other ElementType) bool {
	return e.elementType.Equal(other.elementType)
}

func (e ElementType) Imports() *schema.Imports {
	imports := schema.NewImports()

	switch {
	case e.elementType.Bool != nil:
		if e.elementType.Bool.CustomType != nil && e.elementType.Bool.CustomType.HasImport() {
			imports.Add(*e.elementType.Bool.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: schema.TypesImport,
		})
		return imports
	case e.elementType.Float64 != nil:
		if e.elementType.Float64.CustomType != nil && e.elementType.Float64.CustomType.HasImport() {
			imports.Add(*e.elementType.Float64.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: schema.TypesImport,
		})
		return imports
	case e.elementType.Int64 != nil:
		if e.elementType.Int64.CustomType != nil && e.elementType.Int64.CustomType.HasImport() {
			imports.Add(*e.elementType.Int64.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: schema.TypesImport,
		})
		return imports
	case e.elementType.List != nil:
		imports.Add(NewElementType(e.elementType.List.ElementType).Imports().All()...)
		return imports
	case e.elementType.Map != nil:
		imports.Add(NewElementType(e.elementType.Map.ElementType).Imports().All()...)
		return imports
	case e.elementType.Number != nil:
		if e.elementType.Number.CustomType != nil && e.elementType.Number.CustomType.HasImport() {
			imports.Add(*e.elementType.Number.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: schema.TypesImport,
		})
		return imports
	case e.elementType.Object != nil:
		imports.Add(NewObjectAttributeTypes(e.elementType.Object.AttributeTypes).Imports().All()...)
		return imports
	case e.elementType.Set != nil:
		imports.Add(NewElementType(e.elementType.Set.ElementType).Imports().All()...)
		return imports
	case e.elementType.String != nil:
		if e.elementType.String.CustomType != nil && e.elementType.String.CustomType.HasImport() {
			imports.Add(*e.elementType.String.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: schema.TypesImport,
		})
		return imports
	default:
		return imports
	}
}

func (e ElementType) Schema() []byte {
	return []byte(fmt.Sprintf("ElementType: %s,\n", string(e.ElementType())))
}
