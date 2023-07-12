// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorNestedAttributeObject struct {
	Attributes schema.GeneratorAttributes
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

type GeneratorNestedBlockObject struct {
	Attributes schema.GeneratorAttributes
	Blocks     schema.GeneratorBlocks
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}
