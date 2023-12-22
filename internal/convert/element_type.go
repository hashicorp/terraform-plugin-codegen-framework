// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type ElementType struct {
	elementType schema.ElementType
}

func NewElementType(e schema.ElementType) ElementType {
	return ElementType{
		elementType: e,
	}
}

// TODO: Indentation (\t) for types with elem / attr types?
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
			b.WriteString(fmt.Sprintf(e.elementType.Map.CustomType.Type))
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
			b.WriteString(fmt.Sprintf(e.elementType.Object.CustomType.Type))
		} else {
			b.WriteString(fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", NewObjectAttributeTypes(e.elementType.Object.AttributeTypes).AttributeTypes()))
		}
	case e.elementType.Set != nil:
		if e.elementType.Set.CustomType != nil {
			b.WriteString(fmt.Sprintf(e.elementType.Set.CustomType.Type))
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

func (e ElementType) Schema() []byte {
	return []byte(fmt.Sprintf("ElementType: %s,\n", string(e.ElementType())))
}
