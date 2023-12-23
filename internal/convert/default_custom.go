// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type DefaultCustom struct {
	custom *specschema.CustomDefault
}

func NewDefaultCustom(c *specschema.CustomDefault) DefaultCustom {
	return DefaultCustom{
		custom: c,
	}
}

func (d DefaultCustom) Equal(other DefaultCustom) bool {
	return d.custom.Equal(other.custom)
}

func (d DefaultCustom) Schema() []byte {
	if d.custom != nil && d.custom.SchemaDefinition != "" {
		return []byte(fmt.Sprintf("Default: %s,", d.custom.SchemaDefinition))
	}

	return nil
}
