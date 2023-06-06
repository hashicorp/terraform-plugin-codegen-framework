// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GeneratorListAttribute struct {
	schema.ListAttribute

	CustomType *specschema.CustomType
	Validators []specschema.ListValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the datasourceSchemaImport will be used. Further
// imports are retrieved by calling getElementTypeImports.
func (g GeneratorListAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		// TODO: Refactor once HasImport() helpers have been added to spec Go bindings.
		if g.CustomType.Import != nil && *g.CustomType.Import != "" {
			imports[*g.CustomType.Import] = struct{}{}
		}
	} else {
		imports[datasourceSchemaImport] = struct{}{}
	}

	elemTypeImports := getElementTypeImports(g.ElementType, make(map[string]struct{}))

	for k := range elemTypeImports {
		imports[k] = struct{}{}
	}

	for _, v := range g.Validators {
		if v.Custom == nil {
			continue
		}

		if v.Custom.Import == nil {
			continue
		}

		if *v.Custom.Import == "" {
			continue
		}

		imports[validatorImport] = struct{}{}
		imports[*v.Custom.Import] = struct{}{}
	}

	return imports
}

func (g GeneratorListAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.ListAttribute.Equal(h.ListAttribute)
}

func (g GeneratorListAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": getElementType,
	}

	t, err := template.New("list_attribute").Funcs(funcMap).Parse(listAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorListAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorListAttribute) validatorsEqual(x, y []specschema.ListValidator) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	//TODO: Sort before comparing.
	for k, v := range x {
		if v.Custom == nil && y[k].Custom != nil {
			return false
		}

		if v.Custom != nil && y[k].Custom == nil {
			return false
		}

		if v.Custom != nil && y[k].Custom != nil {
			if *v.Custom.Import != *y[k].Custom.Import {
				return false
			}
		}

		if v.Custom.SchemaDefinition != y[k].Custom.SchemaDefinition {
			return false
		}
	}

	return true
}

func getElementType(elementType attr.Type) string {
	switch t := elementType.(type) {
	case basetypes.BoolType:
		return "types.BoolType"
	case basetypes.Float64Type:
		return "types.Float64Type"
	case basetypes.Int64Type:
		return "types.Int64Type"
	case types.ListType:
		return fmt.Sprintf("types.ListType{\nElemType: %s,\n}", getElementType(t.ElementType()))
	case types.MapType:
		return fmt.Sprintf("types.MapType{\nElemType: %s,\n}", getElementType(t.ElementType()))
	case basetypes.NumberType:
		return "types.NumberType"
	case types.ObjectType:
		return fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", getAttrTypes(t.AttrTypes))
	case types.SetType:
		return fmt.Sprintf("types.SetType{\nElemType: %s,\n}", getElementType(t.ElementType()))
	case basetypes.StringType:
		return "types.StringType"
	}

	return ""
}

// TODO: Handle custom types
func getElementTypeImports(elementType attr.Type, imports map[string]struct{}) map[string]struct{} {
	if elementType == nil {
		return imports
	}

	imports[typesImport] = struct{}{}

	switch t := elementType.(type) {
	case basetypes.BoolType,
		basetypes.Float64Type,
		basetypes.Int64Type,
		basetypes.NumberType,
		basetypes.StringType:
		return imports
	case types.ListType:
		return getElementTypeImports(t.ElementType(), imports)
	case types.MapType:
		return getElementTypeImports(t.ElementType(), imports)
	case types.ObjectType:
		return getAttrTypesImports(t.AttrTypes, imports)
	case types.SetType:
		return getElementTypeImports(t.ElementType(), imports)
	}

	return imports
}
