// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type DefaultString struct {
	stringDefault *specschema.StringDefault
}

func NewDefaultString(b *specschema.StringDefault) DefaultString {
	return DefaultString{
		stringDefault: b,
	}
}

func (d DefaultString) Equal(other DefaultString) bool {
	return d.stringDefault.Equal(other.stringDefault)
}

func (d DefaultString) Schema() []byte {
	if d.stringDefault == nil {
		return nil
	}

	if d.stringDefault.Static != nil {
		return []byte(fmt.Sprintf("Default: stringdefault.StaticString(%q),\n", *d.stringDefault.Static))
	}

	if d.stringDefault.Custom != nil && d.stringDefault.Custom.SchemaDefinition != "" {
		return []byte(fmt.Sprintf("Default: %s,\n", d.stringDefault.Custom.SchemaDefinition))
	}

	return nil
}
