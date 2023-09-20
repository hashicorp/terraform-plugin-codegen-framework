// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"sort"
	"strings"
)

type GeneratorAttributes map[string]GeneratorAttribute

func (g GeneratorAttributes) Schema() (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	var keys = make([]string, 0, len(g))

	for k := range g {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if g[k] == nil {
			continue
		}

		str, err := g[k].Schema(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}
