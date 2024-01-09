// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"
	"strconv"
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

func (d DeprecationMessage) Equal(other DeprecationMessage) bool {
	if d.deprecationMessage == nil && other.deprecationMessage == nil {
		return true
	}

	if d.deprecationMessage == nil || other.deprecationMessage == nil {
		return false
	}

	return *d.deprecationMessage == *other.deprecationMessage
}

func (d DeprecationMessage) Schema() []byte {
	if d.deprecationMessage != nil {
		return []byte(fmt.Sprintf("DeprecationMessage: %s,\n", strconv.Quote(*d.deprecationMessage)))
	}

	return nil
}
