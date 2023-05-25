package datasource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type GeneratorListNestedAttribute struct {
	schema.ListNestedAttribute

	CustomType   *specschema.CustomType
	NestedObject GeneratorNestedAttributeObject
	Validators   []specschema.ListValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the datasourceSchemaImport will be used. The same
// logic is applied to the NestedObject. Further imports are then retrieved by
// calling Imports on each of the nested attributes.
func (g GeneratorListNestedAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.Import != nil && *g.CustomType.Import != "" {
			imports[*g.CustomType.Import] = struct{}{}
		}
	} else {
		imports[datasourceSchemaImport] = struct{}{}
	}

	if g.NestedObject.CustomType != nil {
		if g.NestedObject.CustomType.Import != nil && *g.NestedObject.CustomType.Import != "" {
			imports[*g.NestedObject.CustomType.Import] = struct{}{}
		}
	} else {
		imports[datasourceSchemaImport] = struct{}{}
	}

	for _, v := range g.NestedObject.Attributes {
		for k := range v.Imports() {
			imports[k] = struct{}{}
		}
	}

	return imports
}

func (g GeneratorListNestedAttribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorListNestedAttribute); !ok {
		return false
	}

	glna := ga.(GeneratorListNestedAttribute)

	if !customTypeEqual(g.CustomType, glna.CustomType) {
		return false
	}

	if !g.listValidatorsEqual(g.Validators, glna.Validators) {
		return false
	}

	if !customTypeEqual(g.NestedObject.CustomType, glna.NestedObject.CustomType) {
		return false
	}

	if !g.objectValidatorsEqual(g.NestedObject.Validators, glna.NestedObject.Validators) {
		return false
	}

	for k, a := range g.NestedObject.Attributes {
		if !a.Equal(glna.NestedObject.Attributes[k]) {
			return false
		}
	}

	return true
}

func (g GeneratorListNestedAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": getAttributes,
	}

	t, err := template.New("list_nested_attribute").Funcs(funcMap).Parse(listNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorListNestedAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorListNestedAttribute) listValidatorsEqual(x, y []specschema.ListValidator) bool {
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

func (g GeneratorListNestedAttribute) objectValidatorsEqual(x, y []specschema.ObjectValidator) bool {
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
