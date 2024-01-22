// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
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

func (d DefaultCustom) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	if d.custom == nil {
		return imports
	}

	for _, i := range d.custom.Imports {
		if len(i.Path) > 0 {
			imports.Add(i)
		}
	}

	return imports
}

func (d DefaultCustom) Schema() []byte {
	if d.custom != nil && d.custom.SchemaDefinition != "" {
		return []byte(fmt.Sprintf("Default: %s,\n", d.custom.SchemaDefinition))
	}

	return nil
}
