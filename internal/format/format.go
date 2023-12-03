// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package format

import (
	"go/format"
	"regexp"
	"strings"
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

func ToPascalCase(str string) string {
	return snakeLetters.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// snakeLetters will match to the first letter and an underscore followed by a letter
var snakeLetters = regexp.MustCompile("(^[a-z])|_[a-z0-9]")
