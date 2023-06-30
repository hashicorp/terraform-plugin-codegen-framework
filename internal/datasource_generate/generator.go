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

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	fwtypes "github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorSchema interface {
	ToBytes() (map[string][]byte, error)
}

type GeneratorDataSourceSchemas struct {
	schemas map[string]GeneratorDataSourceSchema
	// TODO: Could add a field to hold custom templates that are used in calls toBytes() to allow overriding.
	// getAttributes() and getBlocks() funcs.
}

type GeneratorDataSourceSchema struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
}

func NewGeneratorDataSourceSchemas(schemas map[string]GeneratorDataSourceSchema) GeneratorDataSourceSchemas {
	return GeneratorDataSourceSchemas{
		schemas: schemas,
	}
}

func (g GeneratorDataSourceSchemas) ToBytes(packageName string) (map[string][]byte, error) {
	schemasBytes := make(map[string][]byte, len(g.schemas))

	for k, s := range g.schemas {
		b, err := g.toBytes(k, s, packageName)

		if err != nil {
			return nil, err
		}

		schemasBytes[k] = b
	}

	return schemasBytes, nil
}

func (g GeneratorDataSourceSchemas) toBytes(name string, s GeneratorDataSourceSchema, packageName string) ([]byte, error) {
	funcMap := template.FuncMap{
		"getImports":    getImports,
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
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
		GeneratorDataSourceSchema: s,
		PackageName:               packageName,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getImports(s GeneratorDataSourceSchema) (string, error) {
	imports := schema.NewImports()

	for _, v := range s.Attributes {
		imports.Add(v.Imports().All()...)
	}

	for _, v := range s.Blocks {
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

func getAttributes(attributes map[string]GeneratorAttribute) (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var keys = make([]string, 0, len(attributes))

	for k := range attributes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if attributes[k] == nil {
			continue
		}

		str, err := attributes[k].ToString(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

func getBlocks(blocks map[string]GeneratorBlock) (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee block order as maps are unordered in Go.
	var keys = make([]string, 0, len(blocks))

	for k := range blocks {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if blocks[k] == nil {
			continue
		}

		str, err := blocks[k].ToString(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

type AttrType interface {
	GeneratorAttrType() (GeneratorAttrType, error)
}

type GeneratorImport interface {
	Imports() *schema.Imports
}

type GeneratorModel interface {
	ModelField(string) (model.Field, error)
}

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	ToString(string) (string, error)
	GeneratorModel
	GeneratorImport
	AttrType
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	ToString(string) (string, error)
	GeneratorModel
	GeneratorImport
	AttrType
}

type GeneratorNestedAttributeObject struct {
	Attributes map[string]GeneratorAttribute
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

type GeneratorNestedBlockObject struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

type GeneratorAttrType struct {
	attr.Type
}

func (g GeneratorAttrType) Equal(t attr.Type) bool {
	return t.Equal(g.Type)
}

func customTypeEqual(x, y *specschema.CustomType) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if x.Import == nil && y.Import != nil {
		return false
	}

	if x.Import != nil && y.Import == nil {
		return false
	}

	if x.Import != nil && y.Import != nil {
		if *x.Import != *y.Import {
			return false
		}
	}

	if x.Type != y.Type {
		return false
	}

	if x.ValueType != y.ValueType {
		return false
	}

	return true
}

func objectTypeEqual(x, y *specschema.ObjectType) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil || y == nil {
		return false
	}

	if !customTypeEqual(x.CustomType, y.CustomType) {
		return false
	}

	for k, v := range x.AttributeTypes {
		if v.Name != y.AttributeTypes[k].Name {
			return false
		}

		a := specschema.ElementType{
			Bool:    v.Bool,
			Float64: v.Float64,
			Int64:   v.Int64,
			List:    v.List,
			Map:     v.Map,
			Number:  v.Number,
			Object:  v.Object,
			Set:     v.Set,
			String:  v.String,
		}

		b := specschema.ElementType{
			Bool:    y.AttributeTypes[k].Bool,
			Float64: y.AttributeTypes[k].Float64,
			Int64:   y.AttributeTypes[k].Int64,
			List:    y.AttributeTypes[k].List,
			Map:     y.AttributeTypes[k].Map,
			Number:  y.AttributeTypes[k].Number,
			Object:  y.AttributeTypes[k].Object,
			Set:     y.AttributeTypes[k].Set,
			String:  y.AttributeTypes[k].String,
		}

		if !elementTypeEqual(a, b) {
			return false
		}
	}

	return true
}

func elementTypeEqual(x, y specschema.ElementType) bool {
	if x.Bool != nil && y.Bool != nil {
		return customTypeEqual(x.Bool.CustomType, y.Bool.CustomType)
	}

	if x.Float64 != nil && y.Float64 != nil {
		return customTypeEqual(x.Float64.CustomType, y.Float64.CustomType)
	}

	if x.Int64 != nil && y.Float64 != nil {
		return customTypeEqual(x.Int64.CustomType, y.Int64.CustomType)
	}

	if x.List != nil && y.List != nil {
		if !customTypeEqual(x.List.CustomType, y.List.CustomType) {
			return false
		}

		return elementTypeEqual(x.List.ElementType, y.List.ElementType)
	}

	if x.Map != nil && y.Map != nil {
		if !customTypeEqual(x.Map.CustomType, y.Map.CustomType) {
			return false
		}

		return elementTypeEqual(x.Map.ElementType, y.Map.ElementType)
	}

	if x.Number != nil && y.Number != nil {
		return customTypeEqual(x.Number.CustomType, y.Number.CustomType)
	}

	if x.Object != nil && y.Object != nil {
		return objectTypeEqual(x.Object, y.Object)
	}

	if x.Set != nil && y.Set != nil {
		if !customTypeEqual(x.Set.CustomType, y.Set.CustomType) {
			return false
		}

		return elementTypeEqual(x.Set.ElementType, y.Set.ElementType)
	}

	if x.String != nil && y.String != nil {
		return customTypeEqual(x.String.CustomType, y.String.CustomType)
	}

	return false
}

type DataSourcesModelsGenerator struct {
}

func NewDataSourcesModelsGenerator() DataSourcesModelsGenerator {
	return DataSourcesModelsGenerator{}
}

func (d DataSourcesModelsGenerator) Process(schemas map[string]GeneratorDataSourceSchema) (map[string][]byte, error) {
	dataSourcesModels := make(map[string][]byte, len(schemas))

	for name, schema := range schemas {
		var buf bytes.Buffer

		generatorDataSourceSchema := GeneratorDataSourceSchema{
			Attributes: schema.Attributes,
			Blocks:     schema.Blocks,
		}

		models, err := generatorDataSourceSchema.Model(name)
		if err != nil {
			return nil, err
		}

		for _, m := range models {
			buf.WriteString("\n" + m.String() + "\n")
		}

		dataSourcesModels[name] = buf.Bytes()
	}

	return dataSourcesModels, nil
}

func (g GeneratorDataSourceSchema) Model(name string) ([]model.Model, error) {
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

	// If there are any nested attributes, generate model.
	for _, attributeName := range attributeNames {
		var nestedModels []model.Model

		switch t := g.Attributes[attributeName].(type) {
		case GeneratorListNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorMapNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorSetNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		case GeneratorSingleNestedAttribute:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.Attributes,
			}

			nestedModels, err = generatorDataSourceSchema.Model(attributeName)
			if err != nil {
				return nil, err
			}
		}

		models = append(models, nestedModels...)
	}

	// Using sorted blockNames to guarantee block order as maps are unordered in Go.
	var blockNames = make([]string, 0, len(g.Blocks))

	for blockName := range g.Blocks {
		blockNames = append(blockNames, blockName)
	}

	sort.Strings(blockNames)

	// If there are any nested blocks, generate model.
	for _, blockName := range blockNames {
		var nestedModels []model.Model

		switch t := g.Blocks[blockName].(type) {
		case GeneratorListNestedBlock:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
				Blocks:     t.NestedObject.Blocks,
			}

			nestedModels, err = generatorDataSourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		case GeneratorSetNestedBlock:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.NestedObject.Attributes,
				Blocks:     t.NestedObject.Blocks,
			}

			nestedModels, err = generatorDataSourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		case GeneratorSingleNestedBlock:
			generatorDataSourceSchema := GeneratorDataSourceSchema{
				Attributes: t.Attributes,
				Blocks:     t.Blocks,
			}

			nestedModels, err = generatorDataSourceSchema.Model(blockName)
			if err != nil {
				return nil, err
			}
		}

		models = append(models, nestedModels...)
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

func (g GeneratorDataSourceSchema) ModelObjectHelpers(name string) ([]model.ObjectHelper, error) {
	var modelObjectHelpers []model.ObjectHelper

	//Assumption is that if this is called then we have already detected nested attribute or block.
	//Could do the recursion (below) first to ensure that this is the case, would also avoid generation of
	//model object helpers for top-level.

	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	modelObjectHelperAttrTypes := make(map[string]string)

	for _, k := range attributeKeys {
		if g.Attributes[k] == nil {
			continue
		}

		//g.Attributes[k].ModelField()

		modelObjectHelperAttrTypes[k] = ""
	}

	//// TODO: Following is required for recursion
	//// Using sorted attributeKeys to guarantee attribute order as maps are unordered in Go.
	//var attributeKeys = make([]string, 0, len(g.Attributes))
	//
	//for k := range g.Attributes {
	//	attributeKeys = append(attributeKeys, k)
	//}
	//
	//sort.Strings(attributeKeys)
	//
	//for _, k := range attributeKeys {
	//	if g.Attributes[k] == nil {
	//		continue
	//	}
	//
	//	switch g.Attributes[k].(type) {
	//	case GeneratorListNestedAttribute:
	//
	//	}
	//}

	return modelObjectHelpers, nil
}

func ElementTypeGeneratorAttrType(e specschema.ElementType) (GeneratorAttrType, error) {
	switch {
	case e.Bool != nil:
		return GeneratorAttrType{
			fwtypes.BoolType,
		}, nil
	case e.Float64 != nil:
		return GeneratorAttrType{
			fwtypes.Float64Type,
		}, nil
	case e.Int64 != nil:
		return GeneratorAttrType{
			fwtypes.Int64Type,
		}, nil
	case e.List != nil:
		elemType, err := ElementTypeGeneratorAttrType(e.List.ElementType)
		if err != nil {
			return GeneratorAttrType{}, err
		}

		return GeneratorAttrType{
			fwtypes.ListType{
				ElemType: elemType,
			},
		}, nil
	case e.Map != nil:
		elemType, err := ElementTypeGeneratorAttrType(e.Map.ElementType)
		if err != nil {
			return GeneratorAttrType{}, err
		}

		return GeneratorAttrType{
			fwtypes.MapType{
				ElemType: elemType,
			},
		}, nil
	case e.Number != nil:
		return GeneratorAttrType{
			fwtypes.NumberType,
		}, nil
	case e.Object != nil:
		objAttrTypes, err := ObjectAttributeTypesGeneratorAttrType(e.Object.AttributeTypes)
		if err != nil {
			return GeneratorAttrType{}, err
		}

		return GeneratorAttrType{
			fwtypes.ObjectType{
				AttrTypes: objAttrTypes,
			},
		}, nil
	case e.Set != nil:
		elemType, err := ElementTypeGeneratorAttrType(e.Set.ElementType)
		if err != nil {
			return GeneratorAttrType{}, err
		}

		return GeneratorAttrType{
			fwtypes.SetType{
				ElemType: elemType,
			},
		}, nil
	case e.String != nil:
		return GeneratorAttrType{
			fwtypes.StringType,
		}, nil
	}

	return GeneratorAttrType{}, errors.New("element type is not set")
}

func ObjectAttributeTypesGeneratorAttrType(o specschema.ObjectAttributeTypes) (map[string]attr.Type, error) {
	objAttrTypes := make(map[string]attr.Type)

	for _, v := range o {
		switch {
		case v.Bool != nil:
			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.BoolType,
			}
		case v.Int64 != nil:
			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.Int64Type,
			}
		case v.Float64 != nil:
			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.Float64Type,
			}
		case v.List != nil:
			elemType, err := ElementTypeGeneratorAttrType(v.List.ElementType)
			if err != nil {
				return nil, err
			}

			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.ListType{
					ElemType: elemType,
				},
			}
		case v.Map != nil:
			elemType, err := ElementTypeGeneratorAttrType(v.Map.ElementType)
			if err != nil {
				return nil, err
			}

			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.MapType{
					ElemType: elemType,
				},
			}
		case v.Number != nil:
			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.NumberType,
			}
		case v.Object != nil:
			objectAttrTypes, err := ObjectAttributeTypesGeneratorAttrType(v.Object.AttributeTypes)
			if err != nil {
				return nil, err
			}

			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.ObjectType{
					AttrTypes: objectAttrTypes,
				},
			}
		case v.Set != nil:
			elemType, err := ElementTypeGeneratorAttrType(v.Set.ElementType)
			if err != nil {
				return nil, err
			}

			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.SetType{
					ElemType: elemType,
				},
			}
		case v.String != nil:
			objAttrTypes[v.Name] = GeneratorAttrType{
				fwtypes.StringType,
			}
		}
	}

	return objAttrTypes, nil
}
