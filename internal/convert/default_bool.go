// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

const defaultBoolImport = "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"

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

func (d DefaultBool) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	if d.boolDefault == nil {
		return imports
	}

	if d.boolDefault.Static != nil {
		imports.Add(code.Import{
			Path: defaultBoolImport,
		})
	}

	if d.boolDefault.Custom != nil {
		for _, i := range d.boolDefault.Custom.Imports {
			if len(i.Path) > 0 {
				imports.Add(i)
			}
		}
	}

	return imports
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
