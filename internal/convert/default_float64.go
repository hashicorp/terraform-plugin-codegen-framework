// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type DefaultFloat64 struct {
	float64Default *specschema.Float64Default
}

func NewDefaultFloat64(b *specschema.Float64Default) DefaultFloat64 {
	return DefaultFloat64{
		float64Default: b,
	}
}

func (d DefaultFloat64) Equal(other DefaultFloat64) bool {
	return d.float64Default.Equal(other.float64Default)
}

func (d DefaultFloat64) Schema() []byte {
	if d.float64Default == nil {
		return nil
	}

	if d.float64Default.Static != nil {
		return []byte(fmt.Sprintf("Default: float64default.StaticFloat64(%g),\n", *d.float64Default.Static))
	}

	if d.float64Default.Custom != nil && d.float64Default.Custom.SchemaDefinition != "" {
		return []byte(fmt.Sprintf("Default: %s,\n", d.float64Default.Custom.SchemaDefinition))
	}

	return nil
}
