// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

const defaultInt64Import = "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"

type DefaultInt64 struct {
	int64Default *specschema.Int64Default
}

func NewDefaultInt64(b *specschema.Int64Default) DefaultInt64 {
	return DefaultInt64{
		int64Default: b,
	}
}

func (d DefaultInt64) Equal(other DefaultInt64) bool {
	return d.int64Default.Equal(other.int64Default)
}

func (d DefaultInt64) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	if d.int64Default == nil {
		return imports
	}

	if d.int64Default.Static != nil {
		imports.Add(code.Import{
			Path: defaultInt64Import,
		})
	}

	if d.int64Default.Custom != nil {
		for _, i := range d.int64Default.Custom.Imports {
			if len(i.Path) > 0 {
				imports.Add(i)
			}
		}
	}

	return imports
}

func (d DefaultInt64) Schema() []byte {
	if d.int64Default == nil {
		return nil
	}

	if d.int64Default.Static != nil {
		return []byte(fmt.Sprintf("Default: int64default.StaticInt64(%d),\n", *d.int64Default.Static))
	}

	if d.int64Default.Custom != nil && d.int64Default.Custom.SchemaDefinition != "" {
		return []byte(fmt.Sprintf("Default: %s,\n", d.int64Default.Custom.SchemaDefinition))
	}

	return nil
}
