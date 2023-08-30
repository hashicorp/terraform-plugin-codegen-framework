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
		Name:            model.SnakeCaseToCamelCase(name),
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
	type aType struct {
		AttrType      string
		AttrValue     string
		AttributeName string
		AttributeType string
		VarName       string
	}

	attrTypeStrings := make(map[string]aType)

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var attributeKeys = make([]string, 0, len(g.Attributes))

	for k := range g.Attributes {
		attributeKeys = append(attributeKeys, k)
	}

	sort.Strings(attributeKeys)

	// Populate attrTypeStrings map for use in template.
	for _, k := range attributeKeys {
		if g.Attributes[k].AttrType() == types.BoolType {
			attrTypeStrings[k] = aType{
				AttrType:  "basetypes.BoolType{}",
				AttrValue: "basetypes.BoolValue",
			}
		}

		if g.Attributes[k].AttrType() == types.Float64Type {
			attrTypeStrings[k] = aType{
				AttrType:  "basetypes.Float64Type{}",
				AttrValue: "basetypes.Float64Value",
			}
		}

		if g.Attributes[k].AttrType() == types.Int64Type {
			attrTypeStrings[k] = aType{
				AttrType:  "basetypes.Int64Type{}",
				AttrValue: "basetypes.Int64Value",
			}
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
					attrTypeStrings[k] = aType{
						AttrType:  fmt.Sprintf("basetypes.ListType{\nElemType: %s,\n}", elemType),
						AttrValue: "basetypes.ListValue",
					}
				} else {
					return nil, fmt.Errorf("%s.%s attribute is a ListType but does not implement Elements interface", name, k)
				}
			} else {
				attrTypeStrings[k] = aType{
					AttrType:      fmt.Sprintf("basetypes.ListType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
					AttrValue:     "basetypes.ListValue",
					AttributeType: "ListNestedAttribute",
				}
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
					attrTypeStrings[k] = aType{
						AttrType:  fmt.Sprintf("basetypes.MapType{\nElemType: %s,\n}", elemType),
						AttrValue: "basetypes.MapValue",
					}
				} else {
					return nil, fmt.Errorf("%s.%s attribute is a MapType but does not implement Elements interface", name, k)
				}
			} else {
				attrTypeStrings[k] = aType{
					AttrType:  fmt.Sprintf("basetypes.MapType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
					AttrValue: "basetypes.MapValue",
				}
			}
		}

		if g.Attributes[k].AttrType() == types.NumberType {
			attrTypeStrings[k] = aType{
				AttrType:  "basetypes.NumberType{}",
				AttrValue: "basetypes.NumberValue",
			}
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
					attrTypeStrings[k] = aType{
						AttrType:  fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s,\n},\n}", attrTypes),
						AttrValue: "basetypes.ObjectValue",
					}
				} else {
					return nil, fmt.Errorf("%s.%s attribute is an ObjectType but does not implement Attrs interface", name, k)
				}
			} else {
				attrTypeStrings[k] = aType{
					AttrType:  fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k)),
					AttrValue: "basetypes.ObjectValue",
				}
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
					attrTypeStrings[k] = aType{
						AttrType:  fmt.Sprintf("basetypes.SetType{\nElemType: %s,\n}", elemType),
						AttrValue: "basetypes.SetValue",
					}
				} else {
					return nil, fmt.Errorf("%s.%s attribute is a SetType but does not implement Elements interface", name, k)
				}
			} else {
				attrTypeStrings[k] = aType{
					AttrType:  fmt.Sprintf("basetypes.SetType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
					AttrValue: "basetypes.SetValue",
				}
			}
		}

		if g.Attributes[k].AttrType() == types.StringType {
			attrTypeStrings[k] = aType{
				AttrType:  "basetypes.StringType{}",
				AttrValue: "basetypes.StringValue",
			}
		}

		a := attrTypeStrings[k]

		camelCaseName := model.SnakeCaseToCamelCase(k)
		lcFirstName := strings.ToLower(camelCaseName[:1]) + camelCaseName[1:]

		a.AttributeName = camelCaseName
		a.VarName = lcFirstName

		attrTypeStrings[k] = a
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
			attrTypeStrings[k] = aType{
				AttrType:  fmt.Sprintf("basetypes.ListType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
				AttrValue: "basetypes.ListValue",
			}
		}

		if _, ok := g.Blocks[k].AttrType().(basetypes.SetType); ok {
			attrTypeStrings[k] = aType{
				AttrType:  fmt.Sprintf("basetypes.SetType{\nElemType: %sValue{}.Type(ctx),\n}", model.SnakeCaseToCamelCase(k)),
				AttrValue: "basetypes.SetValue",
			}
		}

		if _, ok := g.Blocks[k].AttrType().(basetypes.ObjectType); ok {
			attrTypeStrings[k] = aType{
				AttrType:  fmt.Sprintf("basetypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n}", model.SnakeCaseToCamelCase(k)),
				AttrValue: "basetypes.ObjectValue",
			}
		}

		a := attrTypeStrings[k]

		camelCaseName := model.SnakeCaseToCamelCase(k)
		lcFirstName := strings.ToLower(camelCaseName[:1] + camelCaseName[1:])

		a.AttributeName = camelCaseName
		a.VarName = lcFirstName

		attrTypeStrings[k] = a
	}

	t, err := template.New("model_object_helpers").Parse(templates.ModelObjectHelpersTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		Name      string
		AttrTypes map[string]aType
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

// ModelsToFromBytes generates code for expand and flatten functions.
// Whilst associated external types can be defined on any attribute
// type, the only types which are processed are list, map, set and
// single nested attributes, and list, set and single nested blocks.
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

		// Only process attributes implementing GeneratorAttributeAssocExtType.
		var attributeAssocExtType GeneratorAttributeAssocExtType
		var ok bool

		if attributeAssocExtType, ok = g.Attributes[k].(GeneratorAttributeAssocExtType); !ok {
			continue
		}

		// Only process if AssocExtType() is not nil.
		assocExtType := attributeAssocExtType.AssocExtType()

		if assocExtType == nil {
			continue
		}

		// Only process if attribute implements Attributes (i.e., list, map, set, single
		// nested attributes).
		a, ok := g.Attributes[k].(Attributes)

		if !ok {
			continue
		}

		var fields []objectField

		attributeAttributes := a.GetAttributes()

		// Using sorted attributeKeys to guarantee attribute order as maps are unordered in Go.
		var attributeAttributeKeys = make([]string, 0, len(attributeAttributes))

		for aa := range attributeAttributes {
			attributeAttributeKeys = append(attributeAttributeKeys, aa)
		}

		sort.Strings(attributeAttributeKeys)

		for _, x := range attributeAttributeKeys {
			switch attributeAttributes[x].AttrType().(type) {
			case basetypes.BoolTypable:
				fields = append(fields, boolObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.Int64Typable:
				fields = append(fields, int64ObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.Float64Typable:
				fields = append(fields, float64ObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.NumberTypable:
				fields = append(fields, numberObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.StringTypable:
				fields = append(fields, stringObjectField(model.SnakeCaseToCamelCase(x)))
			}
		}

		var t *template.Template
		var err error

		switch attributeAssocExtType.AttrType().(type) {
		case basetypes.ListTypable:
			t, err = template.New("to_from").Parse(templates.ToFromTemplate)
			if err != nil {
				return nil, err
			}
		case basetypes.MapTypable:
			t, err = template.New("map_nested_object_to_from").Parse(templates.MapNestedObjectToFromTemplate)
			if err != nil {
				return nil, err
			}
		case basetypes.ObjectTypable:
			t, err = template.New("single_nested_object_to_from").Parse(templates.SingleNestedObjectToFromTemplate)
			if err != nil {
				return nil, err
			}
		case basetypes.SetTypable:
			t, err = template.New("set_nested_object_to_from").Parse(templates.SetNestedObjectToFromTemplate)
			if err != nil {
				return nil, err
			}
		}

		if t == nil {
			return nil, fmt.Errorf("no matching template for type: %T", attributeAssocExtType.AttrType())
		}

		var templateBuf bytes.Buffer

		templateData := struct {
			Name          string
			Type          string
			TypeReference string
			TypeName      string
			Fields        []objectField
		}{
			Name:          model.SnakeCaseToCamelCase(k),
			Type:          assocExtType.Type(),
			TypeReference: assocExtType.TypeReference(),
			TypeName:      model.DotNotationToCamelCase(assocExtType.TypeReference()),
			Fields:        fields,
		}

		err = t.Execute(&templateBuf, templateData)
		if err != nil {
			return nil, err
		}

		buf.Write(templateBuf.Bytes())
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

		// Only process blocks implementing GeneratorBlockAssocExtType.
		var blockAssocExtType GeneratorBlockAssocExtType
		var ok bool

		if blockAssocExtType, ok = g.Blocks[k].(GeneratorBlockAssocExtType); !ok {
			continue
		}

		// Only process if AssocExtType() is not nil.
		assocExtType := blockAssocExtType.AssocExtType()

		if assocExtType == nil {
			continue
		}

		// Only process if block implements Attributes (i.e., list, set, single
		// nested blocks).
		a, ok := g.Blocks[k].(Attributes)

		if !ok {
			continue
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
			switch blockAttributes[x].AttrType().(type) {
			case basetypes.BoolTypable:
				fields = append(fields, boolObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.Int64Typable:
				fields = append(fields, int64ObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.Float64Typable:
				fields = append(fields, float64ObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.NumberTypable:
				fields = append(fields, numberObjectField(model.SnakeCaseToCamelCase(x)))
			case basetypes.StringTypable:
				fields = append(fields, stringObjectField(model.SnakeCaseToCamelCase(x)))
			}
		}

		var t *template.Template
		var err error

		switch blockAssocExtType.AttrType().(type) {
		case basetypes.ListTypable:
			t, err = template.New("list_nested_object_to_from").Parse(templates.ListNestedObjectToFromTemplate)
			if err != nil {
				return nil, err
			}
		case basetypes.ObjectTypable:
			t, err = template.New("single_nested_object_to_from").Parse(templates.SingleNestedObjectToFromTemplate)
			if err != nil {
				return nil, err
			}
		case basetypes.SetTypable:
			t, err = template.New("set_nested_object_to_from").Parse(templates.SetNestedObjectToFromTemplate)
			if err != nil {
				return nil, err
			}
		}

		if t == nil {
			return nil, fmt.Errorf("no matching template for type: %T", blockAssocExtType.AttrType())
		}

		var templateBuf bytes.Buffer

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

		err = t.Execute(&templateBuf, templateData)
		if err != nil {
			return nil, err
		}

		buf.Write(templateBuf.Bytes())
	}

	return buf.Bytes(), nil
}

type field struct {
	DefaultTo   string
	DefaultFrom string
}

type objectField struct {
	Name string
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

func float64Field() field {
	return field{
		DefaultTo:   "ValueFloat64Pointer",
		DefaultFrom: "Float64PointerValue",
	}
}

func numberField() field {
	return field{
		DefaultTo:   "ValueBigFloat",
		DefaultFrom: "NumberValue",
	}
}

func stringField() field {
	return field{
		DefaultTo:   "ValueStringPointer",
		DefaultFrom: "StringPointerValue",
	}
}

func boolObjectField(name string) objectField {
	return objectField{
		Name:  name,
		field: boolField(),
	}
}

func int64ObjectField(name string) objectField {
	return objectField{
		Name:  name,
		field: int64Field(),
	}
}

func float64ObjectField(name string) objectField {
	return objectField{
		Name:  name,
		field: float64Field(),
	}
}

func numberObjectField(name string) objectField {
	return objectField{
		Name:  name,
		field: numberField(),
	}
}

func stringObjectField(name string) objectField {
	return objectField{
		Name:  name,
		field: stringField(),
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
