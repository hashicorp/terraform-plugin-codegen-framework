// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorNestedAttributeObject struct {
	Attributes             schema.GeneratorAttributes
	AssociatedExternalType *schema.AssocExtType
	CustomType             *specschema.CustomType
	Validators             specschema.ObjectValidators
}

func (g GeneratorNestedAttributeObject) Equal(other GeneratorNestedAttributeObject) bool {
	if !g.Attributes.Equal(other.Attributes) {
		return false
	}

	if !g.AssociatedExternalType.Equal(other.AssociatedExternalType) {
		return false
	}

	if !g.CustomType.Equal(other.CustomType) {
		return false
	}

	return g.Validators.Equal(other.Validators)
}

type GeneratorNestedBlockObject struct {
	Attributes schema.GeneratorAttributes
	Blocks     schema.GeneratorBlocks

	AssociatedExternalType *schema.AssocExtType
	CustomType             *specschema.CustomType
	Validators             specschema.ObjectValidators
}

func (g GeneratorNestedBlockObject) Equal(other GeneratorNestedBlockObject) bool {
	for k := range g.Attributes {
		if _, ok := other.Attributes[k]; !ok {
			return false
		}

		if !g.Attributes[k].Equal(other.Attributes[k]) {
			return false
		}
	}

	for k := range g.Blocks {
		if _, ok := other.Blocks[k]; !ok {
			return false
		}

		if !g.Blocks[k].Equal(other.Blocks[k]) {
			return false
		}
	}

	if !g.AssociatedExternalType.Equal(other.AssociatedExternalType) {
		return false
	}

	if !g.CustomType.Equal(other.CustomType) {
		return false
	}

	return g.Validators.Equal(other.Validators)
}
