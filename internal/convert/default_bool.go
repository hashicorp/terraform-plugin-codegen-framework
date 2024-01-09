// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type DefaultBool struct {
	boolDefault *specschema.BoolDefault
}

func NewDefaultBool(b *specschema.BoolDefault) DefaultBool {
	return DefaultBool{
		boolDefault: b,
	}
}

func (d DefaultBool) Equal(other DefaultBool) bool {
	return d.boolDefault.Equal(other.boolDefault)
}

func (d DefaultBool) Schema() []byte {
	if d.boolDefault == nil {
		return nil
	}

	if d.boolDefault.Static != nil {
		return []byte(fmt.Sprintf("Default: booldefault.StaticBool(%t),\n", *d.boolDefault.Static))
	}

	if d.boolDefault.Custom != nil && d.boolDefault.Custom.SchemaDefinition != "" {
		return []byte(fmt.Sprintf("Default: %s,\n", d.boolDefault.Custom.SchemaDefinition))
	}

	return nil
}
