package datasource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type GeneratorMapAttribute struct {
	schema.MapAttribute

	CustomType *specschema.CustomType
	Validators []specschema.MapValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the datasourceSchemaImport will be used. Further
// imports are retrieved by calling getElementTypeImports.
func (g GeneratorMapAttribute) Imports() map[string]struct{} {
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

func (g GeneratorMapAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorMapAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.MapAttribute.Equal(h.MapAttribute)
}

func (g GeneratorMapAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": getElementType,
	}

	t, err := template.New("map_attribute").Funcs(funcMap).Parse(mapAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorMapAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorMapAttribute) validatorsEqual(x, y []specschema.MapValidator) bool {
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
