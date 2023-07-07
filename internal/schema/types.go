// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

type Attributes interface {
	GetAttributes() GeneratorAttributes
}

type Blocks interface {
	Attributes
	GetBlocks() GeneratorBlocks
}

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	Imports() *Imports
	ModelField(string) (model.Field, error)
	ToString(string) (string, error)
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	Imports() *Imports
	ModelField(string) (model.Field, error)
	ToString(string) (string, error)
}
