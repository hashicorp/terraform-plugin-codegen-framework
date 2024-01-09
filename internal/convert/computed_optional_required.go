// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type ComputedOptionalRequired struct {
	computedOptionalRequired specschema.ComputedOptionalRequired
}

func NewComputedOptionalRequired(c specschema.ComputedOptionalRequired) ComputedOptionalRequired {
	return ComputedOptionalRequired{
		computedOptionalRequired: c,
	}
}

func (c ComputedOptionalRequired) Equal(other ComputedOptionalRequired) bool {
	return c.computedOptionalRequired.Equal(other.computedOptionalRequired)
}

func (c ComputedOptionalRequired) IsComputed() bool {
	if c.computedOptionalRequired == specschema.Computed || c.computedOptionalRequired == specschema.ComputedOptional {
		return true
	}

	return false
}

func (c ComputedOptionalRequired) IsOptional() bool {
	if c.computedOptionalRequired == specschema.Optional || c.computedOptionalRequired == specschema.ComputedOptional {
		return true
	}

	return false
}

func (c ComputedOptionalRequired) IsRequired() bool {
	return c.computedOptionalRequired == specschema.Required
}

func (c ComputedOptionalRequired) Schema() []byte {
	var b bytes.Buffer

	if c.IsRequired() {
		b.WriteString("Required: true,\n")
	}

	if c.IsOptional() {
		b.WriteString("Optional: true,\n")
	}

	if c.IsComputed() {
		b.WriteString("Computed: true,\n")
	}

	return b.Bytes()
}
