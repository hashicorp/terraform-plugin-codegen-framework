// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

const (
	ValidatorTypeBool    ValidatorType = "Bool"
	ValidatorTypeFloat64 ValidatorType = "Float64"
	ValidatorTypeInt64   ValidatorType = "Int64"
	ValidatorTypeList    ValidatorType = "List"
	ValidatorTypeMap     ValidatorType = "Map"
	ValidatorTypeNumber  ValidatorType = "Number"
	ValidatorTypeObject  ValidatorType = "Object"
	ValidatorTypeSet     ValidatorType = "Set"
	ValidatorTypeString  ValidatorType = "String"
)

type ValidatorType string

type ValidatorsCustom struct {
	validatorType ValidatorType
	custom        specschema.CustomValidators
}

func NewValidatorsCustom(t ValidatorType, c specschema.CustomValidators) ValidatorsCustom {
	return ValidatorsCustom{
		validatorType: t,
		custom:        c,
	}
}

func (v ValidatorsCustom) Equal(other ValidatorsCustom) bool {
	if v.validatorType != other.validatorType {
		return false
	}

	if len(v.custom) == 0 && len(other.custom) == 0 {
		return true
	}

	if len(v.custom) != len(other.custom) {
		return false
	}

	v.custom.Sort()

	other.custom.Sort()

	for i := 0; i < len(v.custom); i++ {
		if !v.custom[i].Equal(other.custom[i]) {
			return false
		}
	}

	return true
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
