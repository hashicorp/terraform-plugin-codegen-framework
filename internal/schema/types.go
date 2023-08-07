// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"

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
	AttrType() attr.Type
	Equal(GeneratorAttribute) bool
	Imports() *Imports
	ModelField(string) (model.Field, error)
	ToString(string) (string, error)
}

type GeneratorBlock interface {
	AttrType() attr.Type
	Equal(GeneratorBlock) bool
	Imports() *Imports
	ModelField(string) (model.Field, error)
	ToString(string) (string, error)
}

// TODO: AssocExtType() can be added to GeneratorAttribute,
// and GeneratorBlock once all attributes and blocks
// implement the function.
type GeneratorBlockAssocExtType interface {
	GeneratorBlock
	AssocExtType() *AssocExtType
}
