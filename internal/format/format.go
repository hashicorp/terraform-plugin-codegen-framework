// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package format

import (
	"go/format"
)

func Format(schemas map[string][]byte) (map[string][]byte, error) {
	formattedSchemas := make(map[string][]byte, len(schemas))

	for k, v := range schemas {
		formattedSchema, err := format.Source(v)
		if err != nil {
			return nil, err
		}

		formattedSchemas[k] = formattedSchema
	}

	return formattedSchemas, nil
}
