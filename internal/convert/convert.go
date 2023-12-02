// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"
	"fmt"
	"strconv"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type DeprecationMessage struct {
	deprecationMessage *string
}

func NewDeprecationMessage(d *string) DeprecationMessage {
	return DeprecationMessage{
		deprecationMessage: d,
	}
}

func (d DeprecationMessage) DeprecationMessage() string {
	if d.deprecationMessage == nil {
		return ""
	}

	return *d.deprecationMessage
}

func (d DeprecationMessage) Schema() []byte {
	if d.deprecationMessage != nil {
		return []byte(fmt.Sprintf("DeprecationMessage: %s,\n", strconv.Quote(*d.deprecationMessage)))
	}

	return nil
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

func (d Description) Schema() []byte {
	var b bytes.Buffer

	if d.description != nil {
		quotedDescription := strconv.Quote(*d.description)

		b.WriteString(fmt.Sprintf("Description: %s,\n", quotedDescription))
		b.WriteString(fmt.Sprintf("MarkdownDescription: %s,\n", quotedDescription))
	}

	return b.Bytes()
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

func (s Sensitive) Schema() []byte {
	if s.IsSensitive() {
		return []byte("Sensitive: true,\n")
	}

	return nil
}

const (
	ValidatorTypeBool ValidatorType = "Bool"
)

type ValidatorType string

type ValidatorsCustom struct {
	validatorType ValidatorType
	custom        []*specschema.CustomValidator
}

func NewValidatorsCustom(t ValidatorType, c []*specschema.CustomValidator) ValidatorsCustom {
	return ValidatorsCustom{
		validatorType: t,
		custom:        c,
	}
}

func (v ValidatorsCustom) Schema() []byte {
	var b, cb bytes.Buffer

	for _, c := range v.custom {
		if c != nil {
			cb.WriteString(fmt.Sprintf("%s,\n", c.SchemaDefinition))
		}
	}

	if cb.Len() > 0 {
		b.WriteString(fmt.Sprintf("Validators: []validator.%s{\n", v.validatorType))
		b.Write(cb.Bytes())
		b.WriteString("},\n")
	}

	return b.Bytes()
}
