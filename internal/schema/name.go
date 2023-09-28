// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "strings"

type Name string

func (n Name) CamelCase() string {
	split := strings.Split(string(n), "_")

	var camelCased string

	for _, v := range split {
		if len(v) < 1 {
			continue
		}

		firstChar := v[0:1]
		ucFirstChar := strings.ToUpper(firstChar)

		if len(v) < 2 {
			camelCased += ucFirstChar
			continue
		}

		camelCased += ucFirstChar + v[1:]
	}

	return camelCased
}

func (n Name) CamelCaseLCFirst() string {
	camelCased := n.CamelCase()

	if len(camelCased) < 2 {
		return strings.ToLower(camelCased)
	}

	return strings.ToLower(camelCased[:1]) + camelCased[1:]
}

func (n Name) String() string {
	return string(n)
}
