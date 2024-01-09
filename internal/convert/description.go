// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"
	"fmt"
	"strconv"
)

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

func (d Description) Equal(other Description) bool {
	if d.description == nil && other.description == nil {
		return true
	}

	if d.description == nil || other.description == nil {
		return false
	}

	return *d.description == *other.description
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
