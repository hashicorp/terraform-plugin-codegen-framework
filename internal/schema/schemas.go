// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/logging"
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

func (g GeneratorSchemas) Schemas(packageName, generatorType string) (map[string]GoCode, error) {
	schemasCode := make(map[string]GoCode, len(g.schemas))

	for k, s := range g.schemas {

		pkgName := packageName
		if pkgName == "" {
			pkgName = fmt.Sprintf("%s_%s", strings.ToLower(generatorType), k)
		}

		goCode, err := s.Schema(k, pkgName, generatorType)

		if err != nil {
			return nil, err
		}

		schemasCode[k] = *goCode
	}

	return schemasCode, nil
}

func (g GeneratorSchemas) Models() (map[string]GoCode, error) {
	modelsCode := make(map[string]GoCode, len(g.schemas))

	for name, schema := range g.schemas {
		var buf bytes.Buffer

		generatorSchema := GeneratorSchema{
			Attributes: schema.Attributes,
			Blocks:     schema.Blocks,
		}

		schemaModel, err := generatorSchema.Models(name)
		if err != nil {
			return nil, err
		}

		buf.WriteString("\n" + schemaModel.String() + "\n")

		modelsCode[name] = GoCode{
			NotableExports: map[NotableExport]string{
				ExportSchemaModelType: schemaModel.ModelType(),
			},
			Bytes: buf.Bytes(),
		}
	}

	return modelsCode, nil
}

func (g GeneratorSchemas) CustomTypeValue() (map[string]GoCode, error) {
	customTypeValueCode := make(map[string]GoCode, len(g.schemas))

	for name, s := range g.schemas {
		b, err := s.CustomTypeValueBytes()
		if err != nil {
			return nil, err
		}

		customTypeValueCode[name] = GoCode{Bytes: b}
	}

	return customTypeValueCode, nil
}

func (g GeneratorSchemas) ToFromFunctions(ctx context.Context, logger *slog.Logger) (map[string]GoCode, error) {
	modelsExpandFlattenCode := make(map[string]GoCode, len(g.schemas))

	for name, s := range g.schemas {
		ctxWithPath := logging.SetPathInContext(ctx, name)

		b, err := s.ToFromFunctions(ctxWithPath, logger)
		if err != nil {
			return nil, err
		}

		modelsExpandFlattenCode[name] = GoCode{Bytes: b}
	}

	return modelsExpandFlattenCode, nil
}
