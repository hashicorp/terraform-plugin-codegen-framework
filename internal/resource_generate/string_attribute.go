package resource_generate

import (
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorStringAttribute struct {
	schema.StringAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.StringDefault
	PlanModifiers []specschema.StringPlanModifier
	Validators    []specschema.StringValidator
}

func (g GeneratorStringAttribute) Equal(ga GeneratorAttribute) bool {
	h, ok := ga.(GeneratorStringAttribute)
	if !ok {
		return false
	}

	if !customTypeEqual(g.CustomType, h.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, h.Validators) {
		return false
	}

	return g.StringAttribute.Equal(h.StringAttribute)
}

func getStringDefault(d specschema.StringDefault) string {
	if d.Static != nil {
		return fmt.Sprintf("stringdefault.StaticString(%q)", *d.Static)
	}

	if d.Custom != nil {
		return d.Custom.SchemaDefinition
	}

	return ""
}

func (g GeneratorStringAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getStringDefault": getStringDefault,
	}

	t, err := template.New("string_attribute").Funcs(funcMap).Parse(stringAttributeGoTemplate)
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
