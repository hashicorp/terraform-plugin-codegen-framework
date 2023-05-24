package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorNumberAttribute struct {
	schema.NumberAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.NumberDefault
	PlanModifiers []specschema.NumberPlanModifier
	Validators    []specschema.NumberValidator
}

func (g GeneratorNumberAttribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorNumberAttribute); !ok {
		return false
	}

	gba := ga.(GeneratorNumberAttribute)

	if !customTypeEqual(g.CustomType, gba.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, gba.Validators) {
		return false
	}

	return g.NumberAttribute.Equal(gba.NumberAttribute)
}

func getNumberDefault(d specschema.NumberDefault) string {
	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorNumberAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getNumberDefault": getNumberDefault,
	}

	t, err := template.New("number_attribute").Funcs(funcMap).Parse(numberAttributeGoTemplate)
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
