// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type GeneratorObjectAttribute struct {
	schema.ObjectAttribute

	AttributeTypes []specschema.ObjectAttributeType
	CustomType     *specschema.CustomType
	Validators     []specschema.ObjectValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the schemaImport will be used.
// The imports required for the object attribute types are retrieved by calling
// getAttrTypesImports.
func (g GeneratorObjectAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.HasImport() {
			imports[*g.CustomType.Import] = struct{}{}
		}
	} else {
		imports[schemaImport] = struct{}{}
	}

	attrTypesImports := getAttrTypesImports(g.AttributeTypes, make(map[string]struct{}))

	for k := range attrTypesImports {
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

func (g GeneratorObjectAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorObjectAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.ObjectAttribute.Equal(h.ObjectAttribute)
}

func (g GeneratorObjectAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttrTypes": getAttrTypes,
	}

	t, err := template.New("object_attribute").Funcs(funcMap).Parse(objectAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addCommonAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorObjectAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorObjectAttribute) validatorsEqual(x, y []specschema.ObjectValidator) bool {
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

func getAttrTypes(attrTypes []specschema.ObjectAttributeType) string {
	var aTypes strings.Builder

	for _, v := range attrTypes {
		switch {
		case v.Bool != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.BoolType,", v.Name))
		case v.Float64 != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.Float64Type,", v.Name))
		case v.Int64 != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.Int64Type,", v.Name))
		case v.List != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.ListType{\nElemType: %s,\n},", v.Name, getElementType(v.List.ElementType)))
		case v.Map != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.MapType{\nElemType: %s,\n},", v.Name, getElementType(v.Map.ElementType)))
		case v.Number != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.NumberType,", v.Name))
		case v.Object != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n},", v.Name, getAttrTypes(v.Object)))
		case v.Set != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.SetType{\nElemType: %s,\n},", v.Name, getElementType(v.Set.ElementType)))
		case v.String != nil:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.StringType,", v.Name))
		}
	}

	return aTypes.String()
}

func getAttrTypesImports(attrTypes []specschema.ObjectAttributeType, imports map[string]struct{}) map[string]struct{} {
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
			imports[attrImport] = struct{}{}
			imports[typesImport] = struct{}{}
		case v.Float64 != nil:
			if v.Float64.CustomType != nil && v.Float64.CustomType.HasImport() {
				imports[*v.Float64.CustomType.Import] = struct{}{}
				continue
			}
			imports[attrImport] = struct{}{}
			imports[typesImport] = struct{}{}
		case v.Int64 != nil:
			if v.Int64.CustomType != nil && v.Int64.CustomType.HasImport() {
				imports[*v.Int64.CustomType.Import] = struct{}{}
				continue
			}
			imports[attrImport] = struct{}{}
			imports[typesImport] = struct{}{}
		case v.List != nil:
			if v.List.CustomType != nil && v.List.CustomType.HasImport() {
				imports[*v.List.CustomType.Import] = struct{}{}
			}
			i := getElementTypeImports(v.List.ElementType, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.Map != nil:
			if v.Map.CustomType != nil && v.Map.CustomType.HasImport() {
				imports[*v.Map.CustomType.Import] = struct{}{}
			}
			i := getElementTypeImports(v.Map.ElementType, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.Number != nil:
			if v.Number.CustomType != nil && v.Number.CustomType.HasImport() {
				imports[*v.Number.CustomType.Import] = struct{}{}
				continue
			}
			imports[attrImport] = struct{}{}
			imports[typesImport] = struct{}{}
		case v.Object != nil:
			i := getAttrTypesImports(v.Object, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.Set != nil:
			if v.Set.CustomType != nil && v.Set.CustomType.HasImport() {
				imports[*v.Set.CustomType.Import] = struct{}{}
			}
			i := getElementTypeImports(v.Set.ElementType, imports)
			for k, v := range i {
				imports[k] = v
			}
		case v.String != nil:
			if v.String.CustomType != nil && v.String.CustomType.HasImport() {
				imports[*v.Float64.CustomType.Import] = struct{}{}
				continue
			}
			imports[attrImport] = struct{}{}
			imports[typesImport] = struct{}{}
		}
	}

	return imports
}
