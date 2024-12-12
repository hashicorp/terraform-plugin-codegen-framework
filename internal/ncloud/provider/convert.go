package ncloud_provider

import (
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/provider"
	generatorschema "github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/schema"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
)

func NewSchemas(spec util.NcloudSpecification) (map[string]generatorschema.GeneratorSchema, error) {
	providerSchemas := make(map[string]generatorschema.GeneratorSchema, 1)

	providerSchema, err := NewSchema(spec.Provider)

	if err != nil {
		return nil, err
	}

	providerSchemas[spec.Provider.Name] = providerSchema

	return providerSchemas, nil
}

func NewSchema(p *util.NcloudProvider) (generatorschema.GeneratorSchema, error) {
	var s generatorschema.GeneratorSchema

	if p.Schema == nil {
		return s, nil
	}

	attributes := make(generatorschema.GeneratorAttributes, len(p.Schema.Attributes))
	blocks := make(generatorschema.GeneratorBlocks, len(p.Schema.Blocks))

	for _, v := range p.Schema.Attributes {
		a, err := provider.NewAttribute(v)

		if err != nil {
			return s, err
		}

		attributes[v.Name] = a
	}

	s.Attributes = attributes

	for _, v := range p.Schema.Blocks {
		b, err := provider.NewBlock(v)

		if err != nil {
			return s, err
		}

		blocks[v.Name] = b
	}

	s.Blocks = blocks

	s.Description = p.Schema.Description

	s.MarkdownDescription = p.Schema.MarkdownDescription

	s.DeprecationMessage = p.Schema.DeprecationMessage

	return s, nil
}
