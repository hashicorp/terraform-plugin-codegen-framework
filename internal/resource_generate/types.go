// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	Imports() *schema.Imports
	ModelField(string) (model.Field, error)
	ToString(string) (string, error)
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	Imports() *schema.Imports
	ModelField(string) (model.Field, error)
	ToString(string) (string, error)
}

type GeneratorNestedAttributeObject struct {
	Attributes    GeneratorAttributes
	CustomType    *specschema.CustomType
	PlanModifiers []specschema.ObjectPlanModifier
	Validators    []specschema.ObjectValidator
}

type GeneratorNestedBlockObject struct {
	Attributes    GeneratorAttributes
	Blocks        GeneratorBlocks
	CustomType    *specschema.CustomType
	PlanModifiers []specschema.ObjectPlanModifier
	Validators    []specschema.ObjectValidator
}
