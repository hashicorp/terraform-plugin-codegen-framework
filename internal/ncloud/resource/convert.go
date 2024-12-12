package ncloud_resource

import (
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/resource"
	generatorschema "github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/schema"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
)

func NewSchemas(spec util.NcloudSpecification) (map[string]generatorschema.GeneratorSchema, error) {
	resourceSchemas := make(map[string]generatorschema.GeneratorSchema, len(spec.Resources))

	for _, v := range spec.Resources {
		s, err := NewSchema(v)
		if err != nil {
			return nil, err
		}

		resourceSchemas[v.Name] = s
	}

	return resourceSchemas, nil
}

func NewSchema(d util.Resource) (generatorschema.GeneratorSchema, error) {
	var s generatorschema.GeneratorSchema

	attributes := make(generatorschema.GeneratorAttributes, len(d.Schema.Attributes))
	blocks := make(generatorschema.GeneratorBlocks, len(d.Schema.Blocks))

	for _, v := range d.Schema.Attributes {
		a, err := resource.NewAttribute(v)

		if err != nil {
			return s, err
		}

		attributes[v.Name] = a
	}

	s.Attributes = attributes

	for _, v := range d.Schema.Blocks {
		b, err := resource.NewBlock(v)

		if err != nil {
			return s, err
		}

		blocks[v.Name] = b
	}

	s.Blocks = blocks

	s.Description = d.Schema.Description

	s.MarkdownDescription = d.Schema.MarkdownDescription

	s.DeprecationMessage = d.Schema.DeprecationMessage

	return s, nil
}
