package datasource_generate

import (
	"bytes"
	"sort"
)

// TODO: Field(s) could be added to handle end-user supplying their own templates to allow overriding.
type GeneratorDataSourceSchemas struct {
	schemas map[string]GeneratorDataSourceSchema
}

func NewGeneratorDataSourceSchemas(schemas map[string]GeneratorDataSourceSchema) GeneratorDataSourceSchemas {
	return GeneratorDataSourceSchemas{
		schemas: schemas,
	}
}

func (g GeneratorDataSourceSchemas) SchemasBytes(packageName string) (map[string][]byte, error) {
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

func (g GeneratorDataSourceSchemas) ModelsBytes() (map[string][]byte, error) {
	modelsBytes := make(map[string][]byte, len(g.schemas))

	for name, schema := range g.schemas {
		var buf bytes.Buffer

		generatorDataSourceSchema := GeneratorDataSourceSchema{
			Attributes: schema.Attributes,
			Blocks:     schema.Blocks,
		}

		models, err := generatorDataSourceSchema.Models(name)
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

func (g GeneratorDataSourceSchemas) ModelsObjectHelpersBytes() (map[string][]byte, error) {
	modelsObjectHelpersBytes := make(map[string][]byte, len(g.schemas))

	for name, s := range g.schemas {
		var buf bytes.Buffer

		schema := GeneratorDataSourceSchema{
			Attributes: s.Attributes,
			Blocks:     s.Blocks,
		}

		var attributeKeys = make([]string, 0, len(schema.Attributes))

		for k := range schema.Attributes {
			attributeKeys = append(attributeKeys, k)
		}

		sort.Strings(attributeKeys)

		for _, k := range attributeKeys {
			if schema.Attributes[k] == nil {
				continue
			}

			switch t := schema.Attributes[k].(type) {
			case GeneratorListNestedAttribute:
				var hasNestedAttribute bool

				for _, v := range t.NestedObject.Attributes {
					switch v.(type) {
					case GeneratorListNestedAttribute:
						hasNestedAttribute = true
						break
					}
				}

				if hasNestedAttribute {
					modelObjectHelpers, err := t.ModelObjectHelpersString(k)

					if err != nil {
						return nil, err
					}

					buf.WriteString(modelObjectHelpers)
				}
			}
		}

		modelsObjectHelpersBytes[name] = buf.Bytes()
	}

	return modelsObjectHelpersBytes, nil
}
