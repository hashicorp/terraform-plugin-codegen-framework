// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"bytes"
)

// TODO: Field(s) could be added to handle end-user supplying their own templates to allow overriding.
type GeneratorProviderSchemas struct {
	schemas map[string]GeneratorProviderSchema
}

func NewGeneratorProviderSchemas(schemas map[string]GeneratorProviderSchema) GeneratorProviderSchemas {
	return GeneratorProviderSchemas{
		schemas: schemas,
	}
}

func (g GeneratorProviderSchemas) SchemasBytes(packageName string) (map[string][]byte, error) {
	schemasBytes := make(map[string][]byte, len(g.schemas))

	for k, s := range g.schemas {

		b, err := s.SchemaBytes(k, packageName)

		if err != nil {
			return nil, err
		}

		schemasBytes[k] = b
	}

	return schemasBytes, nil
}

func (g GeneratorProviderSchemas) ModelsBytes() (map[string][]byte, error) {
	modelsBytes := make(map[string][]byte, len(g.schemas))

	for name, schema := range g.schemas {
		var buf bytes.Buffer

		generatorProviderSchema := GeneratorProviderSchema{
			Attributes: schema.Attributes,
			Blocks:     schema.Blocks,
		}

		models, err := generatorProviderSchema.Models(name)
		if err != nil {
			return nil, err
		}

		for _, m := range models {
			buf.WriteString("\n" + m.String() + "\n")
		}

		modelsBytes[name] = buf.Bytes()
	}

	return modelsBytes, nil
}

func (g GeneratorProviderSchemas) ModelsObjectHelpersBytes() (map[string][]byte, error) {
	modelsObjectHelpersBytes := make(map[string][]byte, len(g.schemas))

	for name, s := range g.schemas {
		b, err := s.ModelsObjectHelpersBytes()
		if err != nil {
			return nil, err
		}

		modelsObjectHelpersBytes[name] = b
	}

	return modelsObjectHelpersBytes, nil
}
