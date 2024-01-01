// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type OptionalRequired struct {
	optionalRequired specschema.OptionalRequired
}

func NewOptionalRequired(c specschema.OptionalRequired) OptionalRequired {
	return OptionalRequired{
		optionalRequired: c,
	}
}

func (o OptionalRequired) Equal(other OptionalRequired) bool {
	return o.optionalRequired.Equal(other.optionalRequired)
}

func (o OptionalRequired) IsRequired() bool {
	return o.optionalRequired == specschema.Required
}

func (o OptionalRequired) IsOptional() bool {
	return o.optionalRequired == specschema.Optional
}

func (o OptionalRequired) Schema() []byte {
	var b bytes.Buffer

	if o.IsRequired() {
		b.WriteString("Required: true,\n")
	}

	if o.IsOptional() {
		b.WriteString("Optional: true,\n")
	}

	return b.Bytes()
}
