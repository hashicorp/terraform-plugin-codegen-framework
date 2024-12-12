package ncloud_datasource

import (
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/datasource"
	generatorschema "github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/schema"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
)

func NewSchemas(spec util.NcloudSpecification) (map[string]generatorschema.GeneratorSchema, error) {
	dataSourceSchemas := make(map[string]generatorschema.GeneratorSchema, len(spec.DataSources))

	for _, v := range spec.DataSources {
		s, err := NewSchema(v)
		if err != nil {
			return nil, err
		}

		dataSourceSchemas[v.Name] = s
	}

	return dataSourceSchemas, nil
}

func NewSchema(d util.DataSource) (generatorschema.GeneratorSchema, error) {
	var s generatorschema.GeneratorSchema

	attributes := make(generatorschema.GeneratorAttributes, len(d.Schema.Attributes))
	blocks := make(generatorschema.GeneratorBlocks, len(d.Schema.Blocks))

	for _, v := range d.Schema.Attributes {
		a, err := datasource.NewAttribute(v)

		if err != nil {
			return s, err
		}

		attributes[v.Name] = a
	}

	s.Attributes = attributes

	for _, v := range d.Schema.Blocks {
		b, err := datasource.NewBlock(v)

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
