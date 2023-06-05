// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
)

func JSON(input []byte) error {
	if !json.Valid(input) {
		return errors.New("invalid JSON")
	}

	return nil
}

// Schema
// TODO: Check for duplicate keys at same nesting level in JSON.
// TODO: Handle schema when supplied as URL
func Schema(input []byte) error {
	return spec.Validate(context.TODO(), input)
}
