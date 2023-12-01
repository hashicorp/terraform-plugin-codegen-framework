// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

type DeprecationMessage struct {
	deprecationMessage *string
}

func NewDeprecationMessage(d *string) DeprecationMessage {
	return DeprecationMessage{
		deprecationMessage: d,
	}
}

func (s DeprecationMessage) DeprecationMessage() string {
	if s.deprecationMessage == nil {
		return ""
	}

	return *s.deprecationMessage
}

type Description struct {
	description *string
}

func NewDescription(d *string) Description {
	return Description{
		description: d,
	}
}

func (d Description) Description() string {
	if d.description == nil {
		return ""
	}

	return *d.description
}

type ComputedOptionalRequired struct {
	computedOptionalRequired specschema.ComputedOptionalRequired
}

func NewComputedOptionalRequired(c specschema.ComputedOptionalRequired) ComputedOptionalRequired {
	return ComputedOptionalRequired{
		computedOptionalRequired: c,
	}
}

func (c ComputedOptionalRequired) IsRequired() bool {
	return c.computedOptionalRequired == specschema.Required
}

func (c ComputedOptionalRequired) IsOptional() bool {
	if c.computedOptionalRequired == specschema.Optional || c.computedOptionalRequired == specschema.ComputedOptional {
		return true
	}

	return false
}

func (c ComputedOptionalRequired) IsComputed() bool {
	if c.computedOptionalRequired == specschema.Computed || c.computedOptionalRequired == specschema.ComputedOptional {
		return true
	}

	return false
}

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

type Sensitive struct {
	sensitive *bool
}

func NewSensitive(s *bool) Sensitive {
	return Sensitive{
		sensitive: s,
	}
}

func (s Sensitive) IsSensitive() bool {
	if s.sensitive == nil {
		return false
	}

	return *s.sensitive
}
