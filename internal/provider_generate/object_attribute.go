// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorObjectAttribute struct {
	schema.ObjectAttribute

	AssociatedExternalType *generatorschema.AssocExtType
	// The "specschema" types are used instead of the types within the attribute
	// because support for extracting custom import information is required.
	AttributeTypes specschema.ObjectAttributeTypes
	CustomType     *specschema.CustomType
	Validators     specschema.ObjectValidators
}

func (g GeneratorObjectAttribute) GeneratorSchemaType() generatorschema.Type {
	return generatorschema.GeneratorObjectAttribute
}

func (g GeneratorObjectAttribute) AttrTypes() specschema.ObjectAttributeTypes {
	return g.AttributeTypes
}

func (g GeneratorObjectAttribute) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	customTypeImports := generatorschema.CustomTypeImports(g.CustomType)
	imports.Append(customTypeImports)

	attrTypesImports := generatorschema.GetAttrTypesImports(g.CustomType, g.AttributeTypes)
	imports.Append(attrTypesImports)

	for _, v := range g.Validators {
		customValidatorImports := generatorschema.CustomValidatorImports(v.Custom)
		imports.Append(customValidatorImports)
	}

	if g.AssociatedExternalType != nil {
		imports.Append(generatorschema.AssociatedExternalTypeImports())
	}

	imports.Append(g.AssociatedExternalType.Imports())

	return imports
}

func (g GeneratorObjectAttribute) Equal(ga generatorschema.GeneratorAttribute) bool {
	h, ok := ga.(GeneratorObjectAttribute)

	if !ok {
		return false
	}

	if !g.AttributeTypes.Equal(h.AttributeTypes) {
		return false
	}

	if !g.CustomType.Equal(h.CustomType) {
		return false
	}

	if !g.Validators.Equal(h.Validators) {
		return false
	}

	return g.ObjectAttribute.Equal(h.ObjectAttribute)
}

func (g GeneratorObjectAttribute) Schema(name generatorschema.FrameworkIdentifier) (string, error) {
	type attribute struct {
		Name                     string
		AttributeTypes           string
		CustomType               string
		GeneratorObjectAttribute GeneratorObjectAttribute
	}

	a := attribute{
		Name:                     name.ToString(),
		AttributeTypes:           generatorschema.GetAttrTypes(g.AttributeTypes),
		GeneratorObjectAttribute: g,
	}

	switch {
	case g.CustomType != nil:
		a.CustomType = g.CustomType.Type
	case g.AssociatedExternalType != nil:
		a.CustomType = fmt.Sprintf("%sType{\ntypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n},\n}", name.ToPascalCase(), name.ToPascalCase())
	}

	t, err := template.New("object_attribute").Parse(objectAttributeTemplate)
	if err != nil {
		return "", err
	}

	if _, err = addAttributeTemplate(t); err != nil {
		return "", err
	}

	var buf strings.Builder

	err = t.Execute(&buf, a)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorObjectAttribute) ModelField(name generatorschema.FrameworkIdentifier) (model.Field, error) {
	field := model.Field{
		Name:      name.ToPascalCase(),
		TfsdkName: name.ToString(),
		ValueType: model.ObjectValueType,
	}

	switch {
	case g.CustomType != nil:
		field.ValueType = g.CustomType.ValueType
	case g.AssociatedExternalType != nil:
		field.ValueType = fmt.Sprintf("%sValue", name.ToPascalCase())
	}

	return field, nil
}

func (g GeneratorObjectAttribute) CustomTypeAndValue(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	var buf bytes.Buffer

	objectType := generatorschema.NewCustomObjectType(name)

	b, err := objectType.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	attrTypes := generatorschema.GetAttrTypes(g.AttrTypes())

	objectValue := generatorschema.NewCustomObjectValue(name, attrTypes)

	b, err = objectValue.Render()

	if err != nil {
		return nil, err
	}

	buf.Write(b)

	return buf.Bytes(), nil
}

func (g GeneratorObjectAttribute) ToFromFunctions(name string) ([]byte, error) {
	if g.AssociatedExternalType == nil {
		return nil, nil
	}

	attrTypesToFuncs := generatorschema.GetAttrTypesToFuncs(g.AttributeTypes)

	attrTypesFromFuncs := generatorschema.GetAttrTypesFromFuncs(g.AttributeTypes)

	toFrom := generatorschema.NewToFromObject(name, g.AssociatedExternalType, attrTypesToFuncs, attrTypesFromFuncs)

	b, err := toFrom.Render()

	if err != nil {
		return nil, err
	}

	return b, nil
}
