// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templates"
)

type GeneratorSchema struct {
	Attributes GeneratorAttributes
	Blocks     GeneratorBlocks
}

func (g GeneratorSchema) ImportsString() (string, error) {
	imports := NewImports()

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
					Path: ContextImport,
				},
				{
					Path: DiagImport,
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
					Path: ContextImport,
				},
				{
					Path: DiagImport,
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

func (g GeneratorSchema) SchemaBytes(name, packageName, generatorType string) ([]byte, error) {
	funcMap := template.FuncMap{
		"ImportsString":    g.ImportsString,
		"AttributesString": g.Attributes.String,
		"BlocksString":     g.Blocks.String,
	}

	t, err := template.New("schema").Funcs(funcMap).Parse(
		templates.SchemaGoTemplate,
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		Name string
		GeneratorSchema
		PackageName   string
		GeneratorType string
	}{
		Name:            name,
		GeneratorSchema: g,
		PackageName:     packageName,
		GeneratorType:   generatorType,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g GeneratorSchema) Models(name string) ([]model.Model, error) {
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
			generatorSchema := GeneratorSchema{
				Attributes: nestedAttribute.GetAttributes(),
			}

			nestedModels, err := generatorSchema.Models(attributeName)
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
			generatorSchema := GeneratorSchema{
				Attributes: nestedBlock.GetAttributes(),
				Blocks:     nestedBlock.GetBlocks(),
			}

			nestedModels, err := generatorSchema.Models(blockName)
			if err != nil {
				return nil, err
			}

			models = append(models, nestedModels...)
		}
	}

	return models, nil
}

func (g GeneratorSchema) ModelFields() ([]model.Field, error) {
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
func (g GeneratorSchema) ModelsObjectHelpersBytes() ([]byte, error) {
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
			ng := GeneratorSchema{
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
			ng := GeneratorSchema{
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
func (g GeneratorSchema) ModelObjectHelpersTemplate(name string) ([]byte, error) {
	attrTypeStrings := make(map[string]string)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	// Populate attrTypeStrings map for use in template.
	for _, k := range attributeKeys {
		if g.Attributes[k].AttrType() == types.BoolType {
			attrTypeStrings[k] = "types.BoolType"
		}

		if g.Attributes[k].AttrType() == types.Float64Type {
			attrTypeStrings[k] = "types.Float64Type"
		}

		if g.Attributes[k].AttrType() == types.Int64Type {
			attrTypeStrings[k] = "types.Int64Type"
		}

		// ListType could either be a ListAttribute or a ListNestedAttribute.
		if _, ok := g.Attributes[k].AttrType().(basetypes.ListType); ok {
			// If attribute does not implement Attributes interface it is a ListAttribute else it is a ListNestedAttribute.
			if _, ok := g.Attributes[k].(Attributes); !ok {
				if e, ok := g.Attributes[k].(Elements); ok {
					elemType, err := elementTypeString(e.ElemType())
					if err != nil {
						return nil, err
					}
					attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %s,\n}", elemType)
				} else {
					return nil, fmt.Errorf("%s.%s attribute is a ListType but does not implement Elements interface", name, k)
				}
			} else {
				attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
			}
		}

		// MapType could either be a MapAttribute or a MapNestedAttribute.
		if _, ok := g.Attributes[k].AttrType().(basetypes.MapType); ok {
			// If attribute does not implement Attributes interface it is a MapAttribute else it is a MapNestedAttribute.
			if _, ok := g.Attributes[k].(Attributes); !ok {
				if e, ok := g.Attributes[k].(Elements); ok {
					elemType, err := elementTypeString(e.ElemType())
					if err != nil {
						return nil, err
					}
					attrTypeStrings[k] = fmt.Sprintf("types.MapType{\nElemType: %s,\n}", elemType)
				} else {
					return nil, fmt.Errorf("%s.%s attribute is a MapType but does not implement Elements interface", name, k)
				}
			} else {
				attrTypeStrings[k] = fmt.Sprintf("types.MapType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
			}
		}

		if g.Attributes[k].AttrType() == types.NumberType {
			attrTypeStrings[k] = "types.NumberType"
		}

		// ObjectType could either be an ObjectAttribute or a SingleNestedAttribute.
		if _, ok := g.Attributes[k].AttrType().(basetypes.ObjectType); ok {
			// If attribute does not implement Attributes interface it is an ObjectAttribute else it is a SingleNestedAttribute.
			if _, ok := g.Attributes[k].(Attributes); !ok {
				if o, ok := g.Attributes[k].(Attrs); ok {
					attrTypes, err := attrTypesString(o.AttrTypes())
					if err != nil {
						return nil, err
					}
					attrTypeStrings[k] = fmt.Sprintf("types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", attrTypes)
				} else {
					return nil, fmt.Errorf("%s.%s attribute is an ObjectType but does not implement Attrs interface", name, k)
				}
			} else {
				attrTypeStrings[k] = fmt.Sprintf("types.ObjectType{\nAttrTypes: %sModel{}.ObjectAttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k))
			}
		}

		// SetType could either be a SetAttribute or a SetNestedAttribute.
		if _, ok := g.Attributes[k].AttrType().(basetypes.SetType); ok {
			// If attribute does not implement Attributes interface it is a SetAttribute else it is a SetNestedAttribute.
			if _, ok := g.Attributes[k].(Attributes); !ok {
				if e, ok := g.Attributes[k].(Elements); ok {
					elemType, err := elementTypeString(e.ElemType())
					if err != nil {
						return nil, err
					}
					attrTypeStrings[k] = fmt.Sprintf("types.SetType{\nElemType: %s,\n}", elemType)
				} else {
					return nil, fmt.Errorf("%s.%s attribute is a SetType but does not implement Elements interface", name, k)
				}
			} else {
				attrTypeStrings[k] = fmt.Sprintf("types.SetType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
			}
		}

		if g.Attributes[k].AttrType() == types.StringType {
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
		if _, ok := g.Blocks[k].AttrType().(basetypes.ListType); ok {
			attrTypeStrings[k] = fmt.Sprintf("types.ListType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
		}

		if _, ok := g.Blocks[k].AttrType().(basetypes.SetType); ok {
			attrTypeStrings[k] = fmt.Sprintf("types.SetType{\nElemType: %sModel{}.ObjectType(ctx),\n}", model.SnakeCaseToCamelCase(k))
		}

		if _, ok := g.Blocks[k].AttrType().(basetypes.ObjectType); ok {
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
			ng := GeneratorSchema{
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
			ng := GeneratorSchema{
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

// TODO: Call recursively to generate "expand" and "flatten" functions for all nested blocks and attributes
// which may have an associated external type.
func (g GeneratorSchema) ModelsToFromBytes() ([]byte, error) {
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

		// TODO: Type check is only required until all blocks and attributes
		// implement AssocExtType().
		var attributeAssocExtType GeneratorAttributeAssocExtType
		var ok bool

		if attributeAssocExtType, ok = g.Attributes[k].(GeneratorAttributeAssocExtType); !ok {
			continue
		}

		assocExtType := attributeAssocExtType.AssocExtType()

		if assocExtType == nil {
			continue
		}

		if _, ok := attributeAssocExtType.(Attributes); ok {
			// TODO: Handle objects - list, map, set, single nested object
		} else {
			var templateData attributeField

			switch attributeAssocExtType.AttrType() {
			case types.BoolType:
				templateData = boolAttributeField(model.SnakeCaseToCamelCase(k), assocExtType.Type(), assocExtType.TypeReference())
			}

			t, err := template.New("primitive_to_from").Parse(templates.PrimitiveToFromTemplate)
			if err != nil {
				return nil, err
			}

			var tBuf bytes.Buffer

			err = t.Execute(&tBuf, templateData)
			if err != nil {
				return nil, err
			}

			if tBuf.Len() > 0 {
				buf.WriteString("\n")
				buf.Write(tBuf.Bytes())
			}
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

		// TODO: Type check is only required until all blocks and attributes implement AssocExtType().
		var blockAssocExtType GeneratorBlockAssocExtType
		var ok bool

		if blockAssocExtType, ok = g.Blocks[k].(GeneratorBlockAssocExtType); !ok {
			continue
		}

		assocExtType := blockAssocExtType.AssocExtType()

		if assocExtType == nil {
			continue
		}

		a, ok := g.Blocks[k].(Attributes)

		if !ok {
			return nil, fmt.Errorf("all block types must implement Attributes, %s does not", k)
		}

		// TODO: Need to process blocks as well as attributes in the template
		b, ok := g.Blocks[k].(Blocks)

		if !ok {
			return nil, fmt.Errorf("all block types must implement Blocks, %s does not", k)
		}

		var fields []objectField

		blockAttributes := a.GetAttributes()

		// Using sorted blockKeys to guarantee block order as maps are unordered in Go.
		var blockAttributeKeys = make([]string, 0, len(blockAttributes))

		for ba := range blockAttributes {
			blockAttributeKeys = append(blockAttributeKeys, ba)
		}

		sort.Strings(blockAttributeKeys)

		for _, x := range blockAttributeKeys {
			switch blockAttributes[x].AttrType() {
			case types.BoolType:
				// TODO: Remove type assertion once all attributes and blocks implement AssocExtType()
				if y, ok := blockAttributes[x].(GeneratorAttributeAssocExtType); ok {
					if y.AssocExtType() != nil {
						fields = append(fields, boolObjectField(model.SnakeCaseToCamelCase(x), true))
						continue
					}

					fields = append(fields, boolObjectField(model.SnakeCaseToCamelCase(x), false))
				}
			case types.Int64Type:
				fields = append(fields, int64ObjectField(model.SnakeCaseToCamelCase(x), false))
			}
		}

		// now need to know if we're dealing with list, set or single nested block
		// as that determines whether we're handling a single object in the "expand"
		// "flatten" or a slice of objects (list, set) in "expand" and "flatten".
		// This can be determined by using the attr.Type and a case.

		switch blockAssocExtType.AttrType().(type) {
		case basetypes.ObjectTypable:
			t, err := template.New("model_object_to_from").Parse(templates.ModelObjectToFromTemplate)
			if err != nil {
				return nil, err
			}

			var objBuf bytes.Buffer

			templateData := struct {
				Name          string
				Type          string
				TypeReference string
				Fields        []objectField
			}{
				Name:          model.SnakeCaseToCamelCase(k),
				Type:          assocExtType.Type(),
				TypeReference: assocExtType.TypeReference(),
				Fields:        fields,
			}

			err = t.Execute(&objBuf, templateData)
			if err != nil {
				return nil, err
			}

			buf.WriteString("\n")
			buf.Write(objBuf.Bytes())
		}

		s := GeneratorSchema{
			Attributes: a.GetAttributes(),
			Blocks:     b.GetBlocks(),
		}

		toFromBytes, err := s.ModelsToFromBytes()
		if err != nil {
			return nil, err
		}

		if len(toFromBytes) > 0 {
			buf.WriteString("\n")
			buf.Write(toFromBytes)
		}
	}

	if buf.Len() > 0 {
		buf.WriteString("\n")
	}

	return buf.Bytes(), nil
}

type field struct {
	DefaultTo   string
	DefaultFrom string
}

type attributeField struct {
	Name          string
	Type          string
	TypeReference string
	TfType        string
	field
}

type objectField struct {
	Name            string
	HasAssocExtType bool
	field
}

func boolField() field {
	return field{
		DefaultTo:   "ValueBoolPointer",
		DefaultFrom: "BoolPointerValue",
	}
}

func int64Field() field {
	return field{
		DefaultTo:   "ValueInt64Pointer",
		DefaultFrom: "Int64PointerValue",
	}
}

func boolAttributeField(name, assocType, typeReference string) attributeField {
	return attributeField{
		Name:          name,
		Type:          assocType,
		TypeReference: typeReference,
		TfType:        "types.Bool",
		field:         boolField(),
	}
}

func boolObjectField(name string, hasAssocType bool) objectField {
	return objectField{
		Name:            name,
		HasAssocExtType: hasAssocType,
		field:           boolField(),
	}
}

func int64ObjectField(name string, hasAssocType bool) objectField {
	return objectField{
		Name:            name,
		HasAssocExtType: hasAssocType,
		field:           int64Field(),
	}
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
