// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
)

// TODO: Field(s) could be added to handle end-user supplying their own templates to allow overriding.
type GeneratorSchemas struct {
	schemas map[string]GeneratorSchema
}

func NewGeneratorSchemas(schemas map[string]GeneratorSchema) GeneratorSchemas {
	return GeneratorSchemas{
		schemas: schemas,
	}
}

func (g GeneratorSchemas) SchemasBytes(packageName, generatorType string) (map[string][]byte, error) {
	schemasBytes := make(map[string][]byte, len(g.schemas))

	for k, s := range g.schemas {

		b, err := s.SchemaBytes(k, packageName, generatorType)

		if err != nil {
			return nil, err
		}

		schemasBytes[k] = b
	}

	return schemasBytes, nil
}

func (g GeneratorSchemas) ModelsBytes() (map[string][]byte, error) {
	modelsBytes := make(map[string][]byte, len(g.schemas))

	for name, schema := range g.schemas {
		var buf bytes.Buffer

		generatorSchema := GeneratorSchema{
			Attributes: schema.Attributes,
			Blocks:     schema.Blocks,
		}

		models, err := generatorSchema.Models(name)
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

func (g GeneratorSchemas) ModelsObjectHelpersBytes() (map[string][]byte, error) {
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

func (g GeneratorSchemas) ModelsToFromBytes() (map[string][]byte, error) {
	modelsExpandFlattenBytes := make(map[string][]byte, len(g.schemas))

	for name, s := range g.schemas {
		b, err := s.ModelsToFromBytes()
		if err != nil {
			return nil, err
		}

		modelsExpandFlattenBytes[name] = b
	}

	return modelsExpandFlattenBytes, nil
}
