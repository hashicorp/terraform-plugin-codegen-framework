package datasource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type GeneratorFloat64Attribute struct {
	schema.Float64Attribute

	CustomType *specschema.CustomType
	Validators []specschema.Float64Validator
}

func (g GeneratorFloat64Attribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorFloat64Attribute); !ok {
		return false
	}

	gba := ga.(GeneratorFloat64Attribute)

	if !customTypeEqual(g.CustomType, gba.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, gba.Validators) {
		return false
	}

	return g.Float64Attribute.Equal(gba.Float64Attribute)
}

func (g GeneratorFloat64Attribute) ToString(name string) (string, error) {
	t, err := template.New("float64_attribute").Parse(float64AttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorFloat64Attribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorFloat64Attribute) validatorsEqual(x, y []specschema.Float64Validator) bool {
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
