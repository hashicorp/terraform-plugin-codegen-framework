// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/logging"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

type GeneratorSchema struct {
	Attributes          GeneratorAttributes
	Blocks              GeneratorBlocks
	Description         *string
	MarkdownDescription *string
	DeprecationMessage  *string
}

func (g GeneratorSchema) Imports() (string, error) {
	imports := NewImports()

	imports.Add(
		code.Import{
			Path: ContextImport,
		},
	)

	for _, a := range g.Attributes {
		if a == nil {
			continue
		}

		if _, ok := a.(Attributes); ok {
			imports.Add([]code.Import{
				{
					Path: FmtImport,
				},
				{
					Path: StringsImport,
				},
				{
					Path: DiagImport,
				},
				{
					Path: AttrImport,
				},
				{
					Path: TfTypesImport,
				},
				{
					Path: BaseTypesImport,
				},
			}...)
		}
	}

	for _, b := range g.Blocks {
		if b == nil {
			continue
		}

		imports.Add([]code.Import{
			{
				Path: ContextImport,
			},
		}...)

		if _, ok := b.(Blocks); ok {
			imports.Add([]code.Import{
				{
					Path: FmtImport,
				},
				{
					Path: StringsImport,
				},
				{
					Path: DiagImport,
				},
				{
					Path: AttrImport,
				},
				{
					Path: TfTypesImport,
				},
				{
					Path: BaseTypesImport,
				},
			}...)
		}
	}

	for _, v := range g.Attributes {
		imports.Add(v.Imports().All()...)
	}

	for _, v := range g.Blocks {
		imports.Add(v.Imports().All()...)
	}

	var sb strings.Builder

	for _, i := range imports.All() {
		var alias string

		if i.Alias != nil {
			alias = *i.Alias + " "
		}

		sb.WriteString(fmt.Sprintf("%s%q\n", alias, i.Path))
	}

	return sb.String(), nil
}

func (g GeneratorSchema) Schema(name, packageName, generatorType string) ([]byte, error) {
	attributes, err := g.Attributes.Schema()

	if err != nil {
		return nil, err
	}

	blocks, err := g.Blocks.Schema()

	if err != nil {
		return nil, err
	}

	imports, err := g.Imports()

	if err != nil {
		return nil, err
	}

	description := ""
	if g.Description != nil {
		description = *g.Description
	}

	markdownDescription := ""
	if g.MarkdownDescription != nil {
		markdownDescription = *g.MarkdownDescription
	}

	deprecationMessage := ""
	if g.DeprecationMessage != nil {
		deprecationMessage = *g.DeprecationMessage
	}

	templateData := struct {
		Name                string
		PackageName         string
		GeneratorType       string
		Attributes          string
		Blocks              string
		Description         string
		Imports             string
		MarkdownDescription string
		DeprecationMessage  string
	}{
		Name:                FrameworkIdentifier(name).ToPascalCase(),
		PackageName:         packageName,
		GeneratorType:       generatorType,
		Attributes:          attributes,
		Blocks:              blocks,
		Description:         description,
		Imports:             imports,
		MarkdownDescription: markdownDescription,
		DeprecationMessage:  deprecationMessage,
	}

	t, err := template.New("schema").Parse(SchemaGoTemplate)

	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g GeneratorSchema) Models(name string) ([]model.Model, error) {
	var models []model.Model

	var modelFields []model.Field

	attributeKeys := g.Attributes.SortedKeys()

	for _, k := range attributeKeys {
		if g.Attributes[k] == nil {
			continue
		}

		modelField, err := g.Attributes[k].ModelField(FrameworkIdentifier(k))

		if err != nil {
			return nil, err
		}

		modelFields = append(modelFields, modelField)
	}

	blockKeys := g.Blocks.SortedKeys()

	for _, k := range blockKeys {
		if g.Blocks[k] == nil {
			continue
		}

		modelField, err := g.Blocks[k].ModelField(FrameworkIdentifier(k))

		if err != nil {
			return nil, err
		}

		modelFields = append(modelFields, modelField)
	}

	m := model.Model{
		Name:   FrameworkIdentifier(name).ToPascalCase(),
		Fields: modelFields,
	}

	models = append(models, m)

	return models, nil
}

// CustomTypeValueBytes iterates over all the attributes and blocks to generate code
// for custom type and value types for use in the schema and data models.
func (g GeneratorSchema) CustomTypeValueBytes() ([]byte, error) {
	var buf bytes.Buffer

	attributeKeys := g.Attributes.SortedKeys()

	for _, k := range attributeKeys {
		if g.Attributes[k] == nil {
			continue
		}

		if c, ok := g.Attributes[k].(CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	blockKeys := g.Blocks.SortedKeys()

	for _, k := range blockKeys {
		if g.Blocks[k] == nil {
			continue
		}

		if c, ok := g.Blocks[k].(CustomTypeAndValue); ok {
			b, err := c.CustomTypeAndValue(k)

			if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	if buf.Len() > 0 {
		buf.WriteString("\n")
	}

	return buf.Bytes(), nil
}

// ToFromFunctions generates code for converting to an associated
// external type from a framework type, and from an associated
// external type to a framework type.
func (g GeneratorSchema) ToFromFunctions(ctx context.Context, logger *slog.Logger) ([]byte, error) {
	var buf bytes.Buffer

	attributeKeys := g.Attributes.SortedKeys()

	for _, k := range attributeKeys {
		if g.Attributes[k] == nil {
			continue
		}

		if t, ok := g.Attributes[k].(ToFrom); ok {
			b, err := t.ToFromFunctions(k)

			var unimplErr *UnimplementedError

			if errors.As(err, &unimplErr) {
				logger.Error("error generating to/from methods", "path", fmt.Sprintf("%s.%s.%s", logging.GetPathFromContext(ctx), k, unimplErr.Path()), "err", err)
			} else if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	blockKeys := g.Blocks.SortedKeys()

	for _, k := range blockKeys {
		if g.Blocks[k] == nil {
			continue
		}

		if g.Blocks[k] == nil {
			continue
		}

		if t, ok := g.Blocks[k].(ToFrom); ok {
			b, err := t.ToFromFunctions(k)

			var unimplErr *UnimplementedError

			if errors.As(err, &unimplErr) {
				logger.Error("error generating to/from methods", "path", fmt.Sprintf("%s.%s.%s", logging.GetPathFromContext(ctx), k, unimplErr.Path()), "err", err)
			} else if err != nil {
				return nil, err
			}

			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}

func ElementTypeString(elementType specschema.ElementType) (string, error) {
	switch {
	case elementType.Bool != nil:
		return "types.BoolType", nil
	case elementType.Float64 != nil:
		return "types.Float64Type", nil
	case elementType.Int64 != nil:
		return "types.Int64Type", nil
	case elementType.List != nil:
		elemType, err := ElementTypeString(elementType.List.ElementType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.ListType{\nElemType: %s,\n}", elemType), nil
	case elementType.Map != nil:
		elemType, err := ElementTypeString(elementType.Map.ElementType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.MapType{\nElemType: %s,\n}", elemType), nil
	case elementType.Number != nil:
		return "types.NumberType", nil
	case elementType.Object != nil:
		attrTypesStr, err := AttrTypesString(elementType.Object.AttributeTypes)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", attrTypesStr), nil
	case elementType.Set != nil:
		elemType, err := ElementTypeString(elementType.Set.ElementType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.SetType{\nElemType: %s,\n}", elemType), nil
	case elementType.String != nil:
		return "types.StringType", nil
	}

	return "", errors.New("no matching element type found")
}

// ElementTypeGoType defaults to the defined pointer types on the basis of the
// supplied elementType.
// TODO: Provide a mechanism to allow mapping to be configured. For instance elementType.Float64 => float32
// TODO: Implement for list, map, object, and set.
func ElementTypeGoType(elementType specschema.ElementType) (string, error) {
	switch {
	case elementType.Bool != nil:
		return "*bool", nil
	case elementType.Float64 != nil:
		return "*float64", nil
	case elementType.Int64 != nil:
		return "*int64", nil
	case elementType.List != nil:
		return "", NewUnimplementedError(errors.New("list element type is not yet implemented"))
	case elementType.Map != nil:
		return "", NewUnimplementedError(errors.New("map element type is not yet implemented"))
	case elementType.Number != nil:
		return "*big.Float", nil
	case elementType.Object != nil:
		return "", NewUnimplementedError(errors.New("object element type is not yet implemented"))
	case elementType.Set != nil:
		return "", NewUnimplementedError(errors.New("set element type is not yet implemented"))
	case elementType.String != nil:
		return "*string", nil
	}

	return "", errors.New("no matching element type found")
}

func AttrTypesString(attrTypes specschema.ObjectAttributeTypes) (string, error) {
	var attrTypesStr []string

	for _, v := range attrTypes {
		switch {
		case v.Bool != nil:
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.BoolType", v.Name))
		case v.Float64 != nil:
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.Float64Type", v.Name))
		case v.Int64 != nil:
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.Int64Type", v.Name))
		case v.List != nil:
			elemType, err := ElementTypeString(v.List.ElementType)
			if err != nil {
				return "", err
			}
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.ListType{\nElemType: %s,\n}", v.Name, elemType))
		case v.Map != nil:
			elemType, err := ElementTypeString(v.Map.ElementType)
			if err != nil {
				return "", err
			}
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.MapType{\nElemType: %s,\n}", v.Name, elemType))
		case v.Number != nil:
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.NumberType", v.Name))
		case v.Object != nil:
			objAttrTypesStr, err := AttrTypesString(v.Object.AttributeTypes)
			if err != nil {
				return "", err
			}
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n}\n}", v.Name, objAttrTypesStr))
		case v.Set != nil:
			elemType, err := ElementTypeString(v.Set.ElementType)
			if err != nil {
				return "", err
			}
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.SetType{\nElemType: %s,\n}", v.Name, elemType))
		case v.String != nil:
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.StringType", v.Name))
		}
	}

	return strings.Join(attrTypesStr, ",\n"), nil
}

func ObjectFieldTo(o specschema.ObjectAttributeType) (ObjectField, error) {
	switch {
	case o.Bool != nil:
		return ObjectField{
			GoType: "*bool",
			Type:   "types.Bool",
			ToFunc: "ValueBoolPointer",
		}, nil
	case o.Float64 != nil:
		return ObjectField{
			GoType: "*float64",
			Type:   "types.Float64",
			ToFunc: "ValueFloat64Pointer",
		}, nil
	case o.Int64 != nil:
		return ObjectField{
			GoType: "*int64",
			Type:   "types.Int64",
			ToFunc: "ValueInt64Pointer",
		}, nil
	case o.List != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("list attribute type is not yet implemented"))
	case o.Map != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("map attribute type is not yet implemented"))
	case o.Number != nil:
		return ObjectField{
			GoType: "*big.Float",
			Type:   "types.Number",
			ToFunc: "ValueBigFloat",
		}, nil
	case o.Object != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("object attribute type is not yet implemented"))
	case o.Set != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("set attribute type is not yet implemented"))
	case o.String != nil:
		return ObjectField{
			GoType: "*string",
			Type:   "types.String",
			ToFunc: "ValueStringPointer",
		}, nil
	}

	return ObjectField{}, errors.New("no matching object attribute type found")
}

func ObjectFieldFrom(o specschema.ObjectAttributeType) (ObjectField, error) {
	switch {
	case o.Bool != nil:
		return ObjectField{
			Type:     "types.BoolType",
			FromFunc: "BoolPointerValue",
		}, nil
	case o.Float64 != nil:
		return ObjectField{
			Type:     "types.Float64Type",
			FromFunc: "Float64PointerValue",
		}, nil
	case o.Int64 != nil:
		return ObjectField{
			Type:     "types.Int64Type",
			FromFunc: "Int64PointerValue",
		}, nil
	case o.List != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("list attribute type is not yet implemented"))
	case o.Map != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("map attribute type is not yet implemented"))
	case o.Number != nil:
		return ObjectField{
			Type:     "types.NumberType",
			FromFunc: "NumberValue",
		}, nil
	case o.Object != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("object attribute type is not yet implemented"))
	case o.Set != nil:
		return ObjectField{}, NewUnimplementedError(errors.New("set attribute type is not yet implemented"))
	case o.String != nil:
		return ObjectField{
			Type:     "types.StringType",
			FromFunc: "StringPointerValue",
		}, nil
	}

	return ObjectField{}, errors.New("no matching object attribute type found")
}
