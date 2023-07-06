// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templates"
)

type Attributes interface {
	GetAttributes() GeneratorAttributes
}

type Blocks interface {
	Attributes
	GetBlocks() GeneratorBlocks
}

type GeneratorDataSourceSchema struct {
	Attributes GeneratorAttributes
	Blocks     GeneratorBlocks
}

func (g GeneratorDataSourceSchema) ImportsString() (string, error) {
	imports := schema.NewImports()

	for _, v := range g.Attributes {
		imports.Add(v.Imports().All()...)
	}

	for _, v := range g.Blocks {
		imports.Add(v.Imports().All()...)
	}

	// Both context and terraform-plugin-framework/diag packages are required if
	// model object helpers are generated. Refer to the logic in
	// ModelsObjectHelpersBytes() method.
	for _, a := range g.Attributes {
		if a == nil {
			continue
		}

		if _, ok := a.(Attributes); ok {
			imports.Add([]code.Import{
				{
					Path: schema.ContextImport,
				},
				{
					Path: schema.DiagImport,
				},
			}...)
		}
	}

	// Both context and terraform-plugin-framework/diag packages are required if
	// model object helpers are generated. Refer to the logic in
	// ModelsObjectHelpersBytes() method.
	for _, b := range g.Blocks {
		if b == nil {
			continue
		}

		if _, ok := b.(Blocks); ok {
			imports.Add([]code.Import{
				{
					Path: schema.ContextImport,
				},
				{
					Path: schema.DiagImport,
				},
			}...)
		}
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

func (g GeneratorDataSourceSchema) SchemaBytes(name, packageName string) ([]byte, error) {
	funcMap := template.FuncMap{
		"ImportsString":    g.ImportsString,
		"AttributesString": g.Attributes.String,
		"BlocksString":     g.Blocks.String,
	}

	t, err := template.New("schema").Funcs(funcMap).Parse(
		schemaGoTemplate,
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		Name string
		GeneratorDataSourceSchema
		PackageName string
	}{
		Name:                      name,
		GeneratorDataSourceSchema: g,
		PackageName:               packageName,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g GeneratorDataSourceSchema) Models(name string) ([]model.Model, error) {
	var models []model.Model

	fields, err := g.ModelFields()
	if err != nil {
		return nil, err
	}

	m := model.Model{
		Name:   model.SnakeCaseToCamelCase(name),
		Fields: fields,
	}

	models = append(models, m)

	// Using sorted attributeNames to guarantee attribute order as maps are unordered in Go.
	var attributeNames = make([]string, 0, len(g.Attributes))

	for attributeName := range g.Attributes {
		attributeNames = append(attributeNames, attributeName)
	}

	sort.Strings(attributeNames)

	// If there are any attributes which implement the Attributes interface
	// (i.e., nested attributes), generate model.
	for _, attributeName := range attributeNames {
		if nestedAttribute, ok := g.Attributes[attributeName].(Attributes); ok {
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: nestedAttribute.GetAttributes(),
			}

			nestedModels, err := generatorDataSourceSchema.Models(attributeName)
			if err != nil {
				return nil, err
			}

			models = append(models, nestedModels...)
		}
	}

	// Using sorted blockNames to guarantee block order as maps are unordered in Go.
	var blockNames = make([]string, 0, len(g.Blocks))

	for blockName := range g.Blocks {
		blockNames = append(blockNames, blockName)
	}

	sort.Strings(blockNames)

	// If there are any nested blocks, generate model.
	for _, blockName := range blockNames {
		if nestedBlock, ok := g.Blocks[blockName].(Blocks); ok {
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: nestedBlock.GetAttributes(),
				Blocks:     nestedBlock.GetBlocks(),
			}

			nestedModels, err := generatorDataSourceSchema.Models(blockName)
			if err != nil {
				return nil, err
			}

			models = append(models, nestedModels...)
		}
	}

	return models, nil
}

func (g GeneratorDataSourceSchema) ModelFields() ([]model.Field, error) {
	var modelFields []model.Field

	// Using sorted attributeKeys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	for _, k := range attributeKeys {
		if g.Attributes[k] == nil {
			continue
		}

		modelField, err := g.Attributes[k].ModelField(k)

		if err != nil {
			return nil, err
		}

		modelFields = append(modelFields, modelField)
	}

	// Using sorted blockKeys to guarantee block order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(g.Blocks))

	for k := range g.Blocks {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	for _, k := range blockKeys {
		if g.Blocks[k] == nil {
			continue
		}

		modelField, err := g.Blocks[k].ModelField(k)

		if err != nil {
			return nil, err
		}

		modelFields = append(modelFields, modelField)
	}

	return modelFields, nil
}

// ModelsObjectHelpersBytes iterates over all the attributes and blocks to determine whether
// any of them implement the Attributes interface (i.e., they are nested attributes or
// nested blocks). If any of the attributes or blocks fill the Attributes interface,
// then ModelObjectHelpersTemplate is called.
func (g GeneratorDataSourceSchema) ModelsObjectHelpersBytes() ([]byte, error) {
	var buf bytes.Buffer

	// Using sorted attributeKeys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	for _, k := range attributeKeys {
		if g.Attributes[k] == nil {
			continue
		}

		if a, ok := g.Attributes[k].(Attributes); ok {
			ng := GeneratorDataSourceSchema{
				Attributes: a.GetAttributes(),
			}

			modelObjectHelpers, err := ng.ModelObjectHelpersTemplate(k)
			if err != nil {
				return nil, err
			}

			buf.Write(modelObjectHelpers)
		}
	}

	// Using sorted blockKeys to guarantee block order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(g.Blocks))

	for k := range g.Blocks {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	for _, k := range blockKeys {
		if g.Blocks[k] == nil {
			continue
		}

		if b, ok := g.Blocks[k].(Blocks); ok {
			ng := GeneratorDataSourceSchema{
				Attributes: b.GetAttributes(),
				Blocks:     b.GetBlocks(),
			}

			modelObjectHelpers, err := ng.ModelObjectHelpersTemplate(k)
			if err != nil {
				return nil, err
			}

			buf.Write(modelObjectHelpers)
		}
	}

	if buf.Len() > 0 {
		buf.WriteString("\n")
	}

	return buf.Bytes(), nil
}

// ModelObjectHelpersTemplate iterates over all the attributes and blocks adding the string representation of
// the attr.Type for each attribute or block. A template is then used to generate the model object helpers code.
// If any of the attributes or blocks are nested attributes or nested blocks, respectively, then
// ModelObjectHelpersTemplate is called recursively.
func (g GeneratorDataSourceSchema) ModelObjectHelpersTemplate(name string) ([]byte, error) {
	attrTypeStrings := make(map[string]string)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	// Populate attrTypeStrings map for use in template.
	for _, k := range attributeKeys {
		switch t := g.Attributes[k].(type) {
		case GeneratorBoolAttribute:
			attrTypeStrings[k] = "types.BoolType"
		case GeneratorFloat64Attribute:
			attrTypeStrings[k] = "types.Float64Type"
		case GeneratorInt64Attribute:
			attrTypeStrings[k] = "types.Int64Type"
		case GeneratorListAttribute:
			elemType, err := elementTypeString(t.ElementType)
			if err != nil {
				return nil, err
			}
			attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %s,\n}", elemType)
		case GeneratorListNestedAttribute:
			attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorMapAttribute:
			elemType, err := elementTypeString(t.ElementType)
			if err != nil {
				return nil, err
			}
			attrTypeStrings[k] = fmt.Sprintf("types.MapType{\nElemType: %s,\n}", elemType)
		case GeneratorMapNestedAttribute:
			attrTypeStrings[k] = fmt.Sprintf("types.MapType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorNumberAttribute:
			attrTypeStrings[k] = "types.NumberType"
		case GeneratorObjectAttribute:
			attrTypes, err := attrTypesString(t.AttributeTypes)
			if err != nil {
				return nil, err
			}
			attrTypeStrings[k] = fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", attrTypes)
		case GeneratorSetAttribute:
			elemType, err := elementTypeString(t.ElementType)
			if err != nil {
				return nil, err
			}
			attrTypeStrings[k] = fmt.Sprintf("types.SetType{\nElemType: %s,\n}", elemType)
		case GeneratorSetNestedAttribute:
			attrTypeStrings[k] = fmt.Sprintf("types.SetType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorSingleNestedAttribute:
			attrTypeStrings[k] = fmt.Sprintf("types.ObjectType{\nAttrTypes: %sModel{}.ObjectAttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorStringAttribute:
			attrTypeStrings[k] = "types.StringType"
		}
	}

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var blockKeys = make([]string, 0, len(g.Blocks))

	for k := range g.Blocks {
		blockKeys = append(blockKeys, k)
	}

	sort.Strings(blockKeys)

	// Populate attrTypeStrings map for use in template.
	for _, k := range blockKeys {
		switch g.Blocks[k].(type) {
		case GeneratorListNestedBlock:
			attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorSetNestedBlock:
			attrTypeStrings[k] = fmt.Sprintf("types.SetType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
		case GeneratorSingleNestedBlock:
			attrTypeStrings[k] = fmt.Sprintf("types.ObjectType{\nAttrTypes: %sModel{}.ObjectAttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k))
		}
	}

	t, err := template.New("model_object_helpers").Parse(templates.ModelObjectHelpersTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		Name      string
		AttrTypes map[string]string
	}{
		Name:      model.SnakeCaseToCamelCase(name),
		AttrTypes: attrTypeStrings,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	// Recursively call ModelObjectHelpersTemplate() for each attribute that implements
	// Attributes interface (i.e, nested attributes).
	for _, k := range attributeKeys {
		if a, ok := g.Attributes[k].(Attributes); ok {
			ng := GeneratorDataSourceSchema{
				Attributes: a.GetAttributes(),
			}

			b, err := ng.ModelObjectHelpersTemplate(k)
			if err != nil {
				return nil, err
			}

			buf.WriteString("\n")
			buf.Write(b)
		}
	}

	// Recursively call ModelObjectHelpersTemplate() for each block that implements
	// Blocks interface (i.e, nested blocks).
	for _, k := range blockKeys {
		if b, ok := g.Blocks[k].(Blocks); ok {
			ng := GeneratorDataSourceSchema{
				Attributes: b.GetAttributes(),
				Blocks:     b.GetBlocks(),
			}

			byt, err := ng.ModelObjectHelpersTemplate(k)
			if err != nil {
				return nil, err
			}

			buf.WriteString("\n")
			buf.Write(byt)
		}
	}

	return buf.Bytes(), nil
}

func elementTypeString(elementType specschema.ElementType) (string, error) {
	switch {
	case elementType.Bool != nil:
		return "types.BoolType", nil
	case elementType.Float64 != nil:
		return "types.Float64Type", nil
	case elementType.Int64 != nil:
		return "types.Int64Type", nil
	case elementType.List != nil:
		elemType, err := elementTypeString(elementType.List.ElementType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.ListType{\nElemType: %s,\n}", elemType), nil
	case elementType.Map != nil:
		elemType, err := elementTypeString(elementType.Map.ElementType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.MapType{\nElemType: %s,\n}", elemType), nil
	case elementType.Number != nil:
		return "types.NumberType", nil
	case elementType.Object != nil:
		attrTypesStr, err := attrTypesString(elementType.Object.AttributeTypes)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", attrTypesStr), nil
	case elementType.Set != nil:
		elemType, err := elementTypeString(elementType.Set.ElementType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("types.SetType{\nElemType: %s,\n}", elemType), nil
	case elementType.String != nil:
		return "types.StringType", nil
	}

	return "", errors.New("no matching element type found")
}

func attrTypesString(attrTypes specschema.ObjectAttributeTypes) (string, error) {
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
			elemType, err := elementTypeString(v.List.ElementType)
			if err != nil {
				return "", err
			}
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.ListType{\nElemType: %s,\n}", v.Name, elemType))
		case v.Map != nil:
			elemType, err := elementTypeString(v.Map.ElementType)
			if err != nil {
				return "", err
			}
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.MapType{\nElemType: %s,\n}", v.Name, elemType))
		case v.Number != nil:
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.NumberType", v.Name))
		case v.Object != nil:
			objAttrTypesStr, err := attrTypesString(v.Object.AttributeTypes)
			if err != nil {
				return "", err
			}
			attrTypesStr = append(attrTypesStr, fmt.Sprintf("%q: types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n}\n}", v.Name, objAttrTypesStr))
		case v.Set != nil:
			elemType, err := elementTypeString(v.Set.ElementType)
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
