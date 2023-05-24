package provider_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

type GeneratorStringAttribute struct {
	schema.StringAttribute

	CustomType *specschema.CustomType
	Validators []specschema.StringValidator
}

func (g GeneratorStringAttribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorStringAttribute); !ok {
		return false
	}

	gba := ga.(GeneratorStringAttribute)

	if !customTypeEqual(g.CustomType, gba.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, gba.Validators) {
		return false
	}

	return g.StringAttribute.Equal(gba.StringAttribute)
}

func (g GeneratorStringAttribute) ToString(name string) (string, error) {
	t, err := template.New("string_attribute").Parse(stringAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorStringAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorStringAttribute) validatorsEqual(x, y []specschema.StringValidator) bool {
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
