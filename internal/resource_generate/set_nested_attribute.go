package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorSetNestedAttribute struct {
	schema.SetNestedAttribute

	CustomType    *specschema.CustomType
	Default       *specschema.SetDefault
	NestedObject  GeneratorNestedAttributeObject
	PlanModifiers []specschema.SetPlanModifier
	Validators    []specschema.SetValidator
}

func (g GeneratorSetNestedAttribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorSetNestedAttribute); !ok {
		return false
	}

	glna := ga.(GeneratorSetNestedAttribute)

	if !customTypeEqual(g.CustomType, glna.CustomType) {
		return false
	}

	if !g.setValidatorsEqual(g.Validators, glna.Validators) {
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

func (g GeneratorSetNestedAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": getAttributes,
		"getSetDefault": getSetDefault,
	}

	t, err := template.New("set_nested_attribute").Funcs(funcMap).Parse(setNestedAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSetNestedAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSetNestedAttribute) setValidatorsEqual(x, y []specschema.SetValidator) bool {
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

func (g GeneratorSetNestedAttribute) objectValidatorsEqual(x, y []specschema.ObjectValidator) bool {
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
