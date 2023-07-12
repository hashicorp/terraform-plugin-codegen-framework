// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorNestedAttributeObject struct {
	Attributes    schema.GeneratorAttributes
	CustomType    *specschema.CustomType
	PlanModifiers []specschema.ObjectPlanModifier
	Validators    []specschema.ObjectValidator
}

type GeneratorNestedBlockObject struct {
	Attributes    schema.GeneratorAttributes
	Blocks        schema.GeneratorBlocks
	CustomType    *specschema.CustomType
	PlanModifiers []specschema.ObjectPlanModifier
	Validators    []specschema.ObjectValidator
}
