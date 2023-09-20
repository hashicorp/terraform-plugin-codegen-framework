// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

type Attributes interface {
	GetAttributes() GeneratorAttributes
}

type Attrs interface {
	AttrTypes() specschema.ObjectAttributeTypes
}

type Blocks interface {
	Attributes
	GetBlocks() GeneratorBlocks
}

type Elements interface {
	ElemType() specschema.ElementType
}

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	GeneratorSchemaType() Type
	Imports() *Imports
	ModelField(string) (model.Field, error)
	Schema(string) (string, error)
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	GeneratorSchemaType() Type
	Imports() *Imports
	ModelField(string) (model.Field, error)
	Schema(string) (string, error)
}

type GeneratorAttributeAssocExtType interface {
	GeneratorAttribute
	AssocExtType() *AssocExtType
}

type GeneratorBlockAssocExtType interface {
	GeneratorBlock
	AssocExtType() *AssocExtType
}

type Type int64

const (
	InvalidGeneratorSchemaType Type = iota
	GeneratorBoolAttribute
	GeneratorFloat64Attribute
	GeneratorInt64Attribute
	GeneratorListAttribute
	GeneratorListNestedAttribute
	GeneratorListNestedBlock
	GeneratorMapAttribute
	GeneratorMapNestedAttribute
	GeneratorNumberAttribute
	GeneratorObjectAttribute
	GeneratorSetAttribute
	GeneratorSetNestedAttribute
	GeneratorSetNestedBlock
	GeneratorSingleNestedAttribute
	GeneratorSingleNestedBlock
	GeneratorStringAttribute
)
