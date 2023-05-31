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

type GeneratorObjectAttribute struct {
	schema.ObjectAttribute

	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the datasourceSchemaImport will be used.
// The imports required for the object attribute types are retrieved by calling
// getAttrTypesImports.
func (g GeneratorObjectAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		// TODO: Refactor once HasImport() helpers have been added to spec Go bindings.
		if g.CustomType.Import != nil && *g.CustomType.Import != "" {
			imports[*g.CustomType.Import] = struct{}{}
		}
	} else {
		imports[datasourceSchemaImport] = struct{}{}
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

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
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

func getAttrTypes(attrTypes map[string]attr.Type) string {
	var aTypes strings.Builder

	for k, v := range attrTypes {
		switch t := v.(type) {
		case basetypes.BoolType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.BoolType,", k))
		case basetypes.Float64Type:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.Float64Type,", k))
		case basetypes.Int64Type:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.Int64Type,", k))
		case types.ListType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.ListType{\nElemType: %s,\n},", k, getElementType(t.ElementType())))
		case types.MapType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.MapType{\nElemType: %s,\n},", k, getElementType(t.ElementType())))
		case basetypes.NumberType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.NumberType,", k))
		case types.ObjectType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n},", k, getAttrTypes(t.AttrTypes)))
		case types.SetType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.SetType{\nElemType: %s,\n},", k, getElementType(t.ElementType())))
		case basetypes.StringType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.StringType,", k))
		}
	}

	return aTypes.String()
}

func getAttrTypesImports(attrTypes map[string]attr.Type, imports map[string]struct{}) map[string]struct{} {
	if len(attrTypes) == 0 {
		return imports
	}

	imports[attrImport] = struct{}{}
	imports[typesImport] = struct{}{}

	return imports
}
