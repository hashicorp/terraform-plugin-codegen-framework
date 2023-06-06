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

type GeneratorListAttribute struct {
	schema.ListAttribute

	CustomType  *specschema.CustomType
	ElementType specschema.ElementType
	Validators  []specschema.ListValidator
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

// Equal does not delegate to g.ListAttribute.Equal(h.ListAttribute) as the
// call returns false owing to !a.GetType().Equal(b.GetType()) returning false
// when the ElementType is nil.
func (g GeneratorListAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorListAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !elementTypeEqual(g.ElementType, h.ElementType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	if g.Required != h.Required {
		return false
	}

	if g.Optional != h.Optional {
		return false
	}

	if g.Computed != h.Computed {
		return false
	}

	if g.Sensitive != h.Sensitive {
		return false
	}

	if g.Description != h.Description {
		return false
	}

	if g.MarkdownDescription != h.MarkdownDescription {
		return false
	}

	if g.DeprecationMessage != h.DeprecationMessage {
		return false
	}

	return true
}

func (g GeneratorListAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": getElementType,
	}

	t, err := template.New("list_attribute").Funcs(funcMap).Parse(listAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
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

func getElementType(e specschema.ElementType) string {
	switch {
	case e.Bool != nil:
		return "types.BoolType"
	case e.Float64 != nil:
		return "types.Float64Type"
	case e.Int64 != nil:
		return "types.Int64Type"
	case e.List != nil:
		return fmt.Sprintf("types.ListType{\nElemType: %s,\n}", getElementType(e.List.ElementType))
	case e.Map != nil:
		return fmt.Sprintf("types.MapType{\nElemType: %s,\n}", getElementType(e.Map.ElementType))
	case e.Number != nil:
		return "types.NumberType"
	case e.Object != nil:
		return fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n}", getAttrTypes(e.Object))
	case e.Set != nil:
		return fmt.Sprintf("types.SetType{\nElemType: %s,\n}", getElementType(e.Set.ElementType))
	case e.String != nil:
		return "types.StringType"
	}

	return ""
}

func getElementTypeImports(e specschema.ElementType, imports map[string]struct{}) map[string]struct{} {
	switch {
	case e.Bool != nil:
		if e.Bool.CustomType != nil && e.Bool.CustomType.HasImport() {
			imports[*e.Bool.CustomType.Import] = struct{}{}
			return imports
		}
		imports[typesImport] = struct{}{}
		return imports
	case e.Float64 != nil:
		if e.Float64.CustomType != nil && e.Float64.CustomType.HasImport() {
			imports[*e.Float64.CustomType.Import] = struct{}{}
			return imports
		}
		imports[typesImport] = struct{}{}
		return imports
	case e.Int64 != nil:
		if e.Int64.CustomType != nil && e.Int64.CustomType.HasImport() {
			imports[*e.Int64.CustomType.Import] = struct{}{}
			return imports
		}
		imports[typesImport] = struct{}{}
		return imports
	case e.List != nil:
		return getElementTypeImports(e.List.ElementType, imports)
	case e.Map != nil:
		return getElementTypeImports(e.Map.ElementType, imports)
	case e.Number != nil:
		if e.Number.CustomType != nil && e.Number.CustomType.HasImport() {
			imports[*e.Number.CustomType.Import] = struct{}{}
			return imports
		}
		imports[typesImport] = struct{}{}
		return imports
	case e.Object != nil:
		return getAttrTypesImports(e.Object, imports)
	case e.Set != nil:
		return getElementTypeImports(e.Set.ElementType, imports)
	case e.String != nil:
		if e.String.CustomType != nil && e.String.CustomType.HasImport() {
			imports[*e.Float64.CustomType.Import] = struct{}{}
			return imports
		}
		imports[typesImport] = struct{}{}
		return imports
	default:
		return imports
	}
}
