package provider_convert

import (
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertSchema(d *provider.Provider) (provider_generate.GeneratorProviderSchema, error) {
	var s provider_generate.GeneratorProviderSchema

	if d.Schema == nil {
		return s, nil
	}

	attributes := make(map[string]provider_generate.GeneratorAttribute, len(d.Schema.Attributes))
	blocks := make(map[string]provider_generate.GeneratorBlock, len(d.Schema.Blocks))

	for _, v := range d.Schema.Attributes {
		a, err := convertAttribute(v)

		if err != nil {
			return s, err
		}

		attributes[v.Name] = a
	}

	s.Attributes = attributes

	for _, v := range d.Schema.Blocks {
		b, err := convertBlock(v)

		if err != nil {
			return s, err
		}

		blocks[v.Name] = b
	}

	s.Blocks = blocks

	return s, nil
}
