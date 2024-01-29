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

type ObjectAttributeTypes struct {
	objectAttributeTypes specschema.ObjectAttributeTypes
}

func NewObjectAttributeTypes(o specschema.ObjectAttributeTypes) ObjectAttributeTypes {
	return ObjectAttributeTypes{
		objectAttributeTypes: o,
	}
}

func (o ObjectAttributeTypes) Equal(other ObjectAttributeTypes) bool {
	return o.objectAttributeTypes.Equal(other.objectAttributeTypes)
}

func (o ObjectAttributeTypes) AttributeTypes() []byte {
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
				b.WriteString(fmt.Sprintf("%q: types.MapType{\nElemType: %s,\n},", v.Name, NewElementType(v.Map.ElementType).ElementType()))
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
				b.WriteString(fmt.Sprintf("%q: types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n},", v.Name, NewObjectAttributeTypes(v.Object.AttributeTypes).AttributeTypes()))
			}
		case v.Set != nil:
			if v.Set.CustomType != nil {
				b.WriteString(fmt.Sprintf("%q: %s,", v.Name, v.Set.CustomType.Type))
			} else {
				b.WriteString(fmt.Sprintf("%q: types.SetType{\nElemType: %s,\n},", v.Name, NewElementType(v.Set.ElementType).ElementType()))
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

func (o ObjectAttributeTypes) Imports() *schema.Imports {
	imports := schema.NewImports()

	if len(o.objectAttributeTypes) == 0 {
		return imports
	}

	for _, v := range o.objectAttributeTypes {
		switch {
		case v.Bool != nil:
			if v.Bool.CustomType != nil && v.Bool.CustomType.HasImport() {
				imports.Add(*v.Bool.CustomType.Import)
				continue
			}
			imports.Add(code.Import{Path: schema.AttrImport}, code.Import{Path: schema.TypesImport})
		case v.Float64 != nil:
			if v.Float64.CustomType != nil && v.Float64.CustomType.HasImport() {
				imports.Add(*v.Float64.CustomType.Import)
				continue
			}
			imports.Add(code.Import{Path: schema.AttrImport}, code.Import{Path: schema.TypesImport})
		case v.Int64 != nil:
			if v.Int64.CustomType != nil && v.Int64.CustomType.HasImport() {
				imports.Add(*v.Int64.CustomType.Import)
				continue
			}
			imports.Add(code.Import{Path: schema.AttrImport}, code.Import{Path: schema.TypesImport})
		case v.List != nil:
			if v.List.CustomType != nil && v.List.CustomType.HasImport() {
				imports.Add(*v.List.CustomType.Import)
			}
			imports.Add(NewElementType(v.List.ElementType).Imports().All()...)
		case v.Map != nil:
			if v.Map.CustomType != nil && v.Map.CustomType.HasImport() {
				imports.Add(*v.Map.CustomType.Import)
			}
			imports.Add(NewElementType(v.Map.ElementType).Imports().All()...)
		case v.Number != nil:
			if v.Number.CustomType != nil && v.Number.CustomType.HasImport() {
				imports.Add(*v.Number.CustomType.Import)
				continue
			}
			imports.Add(code.Import{Path: schema.AttrImport}, code.Import{Path: schema.TypesImport})
		case v.Object != nil:
			imports.Add(NewObjectAttributeTypes(v.Object.AttributeTypes).Imports().All()...)
		case v.Set != nil:
			if v.Set.CustomType != nil && v.Set.CustomType.HasImport() {
				imports.Add(*v.Set.CustomType.Import)
			}
			imports.Add(NewElementType(v.Set.ElementType).Imports().All()...)
		case v.String != nil:
			if v.String.CustomType != nil && v.String.CustomType.HasImport() {
				imports.Add(*v.String.CustomType.Import)
				continue
			}
			imports.Add(code.Import{Path: schema.AttrImport}, code.Import{Path: schema.TypesImport})
		}
	}

	return imports
}

func (o ObjectAttributeTypes) Schema() []byte {
	var b, at bytes.Buffer

	at.Write(o.AttributeTypes())

	if at.Len() > 0 {
		b.WriteString("AttributeTypes: map[string]attr.Type{\n")
		b.Write(at.Bytes())
		b.WriteString("\n},\n")
	}

	return b.Bytes()
}
