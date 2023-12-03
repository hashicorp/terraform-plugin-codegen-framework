// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

type OptionalRequired struct {
	optionalRequired specschema.OptionalRequired
}

func NewOptionalRequired(c specschema.OptionalRequired) OptionalRequired {
	return OptionalRequired{
		optionalRequired: c,
	}
}

func (c OptionalRequired) IsRequired() bool {
	return c.optionalRequired == specschema.Required
}

func (c OptionalRequired) IsOptional() bool {
	return c.optionalRequired == specschema.Optional
}
