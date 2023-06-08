// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"encoding/json"
	"errors"
)

func JSON(input []byte) error {
	if !json.Valid(input) {
		return errors.New("invalid JSON")
	}

	return nil
}
