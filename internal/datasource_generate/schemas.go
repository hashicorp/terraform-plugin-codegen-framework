package datasource_generate

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
