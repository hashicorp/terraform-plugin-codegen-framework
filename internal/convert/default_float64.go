// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

const defaultFloat64Import = "github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"

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

func (d DefaultFloat64) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	if d.float64Default == nil {
		return imports
	}

	if d.float64Default.Static != nil {
		imports.Add(code.Import{
			Path: defaultFloat64Import,
		})
	}

	if d.float64Default.Custom != nil {
		for _, i := range d.float64Default.Custom.Imports {
			if len(i.Path) > 0 {
				imports.Add(i)
			}
		}
	}

	return imports
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
