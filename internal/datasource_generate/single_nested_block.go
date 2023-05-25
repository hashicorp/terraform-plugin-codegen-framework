package datasource_generate

import (
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type GeneratorSingleNestedBlock struct {
	schema.SingleNestedBlock

	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

func (g GeneratorSingleNestedBlock) Equal(ga GeneratorBlock) bool {
	if _, ok := ga.(GeneratorSingleNestedBlock); !ok {
		return false
	}

	gsna := ga.(GeneratorSingleNestedBlock)

	if !customTypeEqual(g.CustomType, gsna.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, gsna.Validators) {
		return false
	}

	for k, a := range g.Attributes {
		if !a.Equal(gsna.Attributes[k]) {
			return false
		}
	}

	return true
}

func (g GeneratorSingleNestedBlock) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
	}

	t, err := template.New("single_nested_block").Funcs(funcMap).Parse(singleNestedBlockGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_block").Parse(commonBlockGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorSingleNestedBlock{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorSingleNestedBlock) validatorsEqual(x, y []specschema.ObjectValidator) bool {
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
