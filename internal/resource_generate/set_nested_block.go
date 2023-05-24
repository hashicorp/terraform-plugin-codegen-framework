package resource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type GeneratorSetNestedBlock struct {
	schema.SetNestedBlock

	CustomType    *specschema.CustomType
	NestedObject  GeneratorNestedBlockObject
	PlanModifiers []specschema.SetPlanModifier
	Validators    []specschema.SetValidator
}

func (g GeneratorSetNestedBlock) Equal(ga GeneratorBlock) bool {
	if _, ok := ga.(GeneratorSetNestedBlock); !ok {
		return false
	}

	glna := ga.(GeneratorSetNestedBlock)

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

func (g GeneratorSetNestedBlock) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
	}

	t, err := template.New("set_nested_block").Funcs(funcMap).Parse(setNestedBlockGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_block").Parse(commonBlockGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSetNestedBlock{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSetNestedBlock) setValidatorsEqual(x, y []specschema.SetValidator) bool {
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

func (g GeneratorSetNestedBlock) objectValidatorsEqual(x, y []specschema.ObjectValidator) bool {
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
