package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorSetAttribute struct {
	schema.SetAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.SetDefault
	PlanModifiers []specschema.SetPlanModifier
	Validators    []specschema.SetValidator
}

func (g GeneratorSetAttribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorSetAttribute); !ok {
		return false
	}

	gla := ga.(GeneratorSetAttribute)

	if !customTypeEqual(g.CustomType, gla.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, gla.Validators) {
		return false
	}

	return g.SetAttribute.Equal(gla.SetAttribute)
}

func getSetDefault(d specschema.SetDefault) string {
	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorSetAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getElementType": getElementType,
		"getSetDefault":  getSetDefault,
	}

	t, err := template.New("set_attribute").Funcs(funcMap).Parse(setAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSetAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSetAttribute) validatorsEqual(x, y []specschema.SetValidator) bool {
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
