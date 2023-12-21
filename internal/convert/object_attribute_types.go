package convert

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type ObjectAttributeTypes struct {
	objectAttributeTypes schema.ObjectAttributeTypes
}

func NewObjectAttributeTypes(o schema.ObjectAttributeTypes) ObjectAttributeTypes {
	return ObjectAttributeTypes{
		objectAttributeTypes: o,
	}
}

func (o ObjectAttributeTypes) Schema() []byte {
	var b bytes.Buffer

	for _, v := range o.objectAttributeTypes {
		if b.Len() > 0 {
			b.WriteString("\n")
		}

		switch {
		case v.Bool != nil:
			if v.Bool.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Bool.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.BoolType,", v.Name))
			}
		case v.Float64 != nil:
			if v.Float64.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Float64.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.Float64Type,", v.Name))
			}
		case v.Int64 != nil:
			if v.Int64.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Int64.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.Int64Type,", v.Name))
			}
		case v.List != nil:
			if v.List.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.List.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.ListType{\nElemType: %s,\n},", v.Name, NewElementType(v.List.ElementType).ElementType()))
			}
		case v.Map != nil:
			if v.Map.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Map.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.MapType{\nElemType: %s,\n},", v.Name, NewElementType(v.List.ElementType).ElementType()))
			}
		case v.Number != nil:
			if v.Number.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Number.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.NumberType,", v.Name))
			}
		case v.Object != nil:
			if v.Object.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Object.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n},", v.Name, NewObjectAttributeTypes(v.Object.AttributeTypes).Schema()))
			}
		case v.Set != nil:
			if v.Set.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Set.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.SetType{\nElemType: %s,\n},", v.Name, NewElementType(v.List.ElementType).ElementType()))
			}
		case v.String != nil:
			if v.String.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.String.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.StringType,", v.Name))
			}
		}
	}

	return b.Bytes()
}
