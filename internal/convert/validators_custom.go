// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

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
		if c == nil {
			continue
		}

		if c.SchemaDefinition == "" {
			continue
		}

		cb.WriteString(fmt.Sprintf("%s,\n", c.SchemaDefinition))
	}

	if cb.Len() > 0 {
		b.WriteString(fmt.Sprintf("Validators: []validator.%s{\n", v.validatorType))
		b.Write(cb.Bytes())
		b.WriteString("},\n")
	}

	return b.Bytes()
}
