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

func (g GeneratorSchemas) Schemas(packageName, generatorType string) (map[string][]byte, error) {
	schemasBytes := make(map[string][]byte, len(g.schemas))

	for k, s := range g.schemas {

		pkgName := packageName
		if pkgName == "" {
			pkgName = fmt.Sprintf("%s_%s", strings.ToLower(generatorType), k)
		}

		b, err := s.Schema(k, pkgName, generatorType)

		if err != nil {
			return nil, err
		}

		schemasBytes[k] = b
	}

	return schemasBytes, nil
}

func (g GeneratorSchemas) Models() (map[string][]byte, error) {
	modelsBytes := make(map[string][]byte, len(g.schemas))

	for name, schema := range g.schemas {
		var buf bytes.Buffer

		generatorSchema := GeneratorSchema{
			Attributes:          schema.Attributes,
			Blocks:              schema.Blocks,
			Description:         schema.Description,
			MarkdownDescription: schema.MarkdownDescription,
			DeprecationMessage:  schema.DeprecationMessage,
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

func (g GeneratorSchemas) CustomTypeValue() (map[string][]byte, error) {
	customTypeValueBytes := make(map[string][]byte, len(g.schemas))

	for name, s := range g.schemas {
		b, err := s.CustomTypeValueBytes()
		if err != nil {
			return nil, err
		}

		customTypeValueBytes[name] = b
	}

	return customTypeValueBytes, nil
}

func (g GeneratorSchemas) ToFromFunctions(ctx context.Context, logger *slog.Logger) (map[string][]byte, error) {
	modelsExpandFlattenBytes := make(map[string][]byte, len(g.schemas))

	for name, s := range g.schemas {
		ctxWithPath := logging.SetPathInContext(ctx, name)

		b, err := s.ToFromFunctions(ctxWithPath, logger)
		if err != nil {
			return nil, err
		}

		modelsExpandFlattenBytes[name] = b
	}

	return modelsExpandFlattenBytes, nil
}
