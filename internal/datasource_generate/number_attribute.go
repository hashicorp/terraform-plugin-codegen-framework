package datasource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type GeneratorNumberAttribute struct {
	schema.NumberAttribute

	CustomType *specschema.CustomType
	Validators []specschema.NumberValidator
}

// Imports examines the CustomType and if this is not nil then the CustomType.Import
// will be used if it is not nil. If CustomType.Import is nil then no import will be
// specified as it is assumed that the CustomType.Type and CustomType.ValueType will
// be accessible from the same package that the schema.Schema for the data source is
// defined in. If CustomType is nil, then the datasourceSchemaImport will be used.
func (g GeneratorNumberAttribute) Imports() map[string]struct{} {
	imports := make(map[string]struct{})

	if g.CustomType != nil {
		if g.CustomType.Import != nil && *g.CustomType.Import != "" {
			imports[*g.CustomType.Import] = struct{}{}
		}
	} else {
		imports[datasourceSchemaImport] = struct{}{}
	}

	return imports
}

func (g GeneratorNumberAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorNumberAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.NumberAttribute.Equal(h.NumberAttribute)
}

func (g GeneratorNumberAttribute) ToString(name string) (string, error) {
	t, err := template.New("number_attribute").Parse(numberAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorNumberAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorNumberAttribute) validatorsEqual(x, y []specschema.NumberValidator) bool {
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
