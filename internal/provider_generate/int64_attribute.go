package provider_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

type GeneratorInt64Attribute struct {
	schema.Int64Attribute

	CustomType *specschema.CustomType
	Validators []specschema.Int64Validator
}

func (g GeneratorInt64Attribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorInt64Attribute); !ok {
		return false
	}

	gba := ga.(GeneratorInt64Attribute)

	if !customTypeEqual(g.CustomType, gba.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, gba.Validators) {
		return false
	}

	return g.Int64Attribute.Equal(gba.Int64Attribute)
}

func (g GeneratorInt64Attribute) ToString(name string) (string, error) {
	t, err := template.New("int64_attribute").Parse(int64AttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorInt64Attribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorInt64Attribute) validatorsEqual(x, y []specschema.Int64Validator) bool {
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
