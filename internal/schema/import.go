// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
)

const (
	AttrImport         = "github.com/hashicorp/terraform-plugin-framework/attr"
	BaseTypesImport    = "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ContextImport      = "context"
	DiagImport         = "github.com/hashicorp/terraform-plugin-framework/diag"
	FmtImport          = "fmt"
	MathBigImport      = "math/big"
	PlanModifierImport = "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	StringsImport      = "strings"
	TfTypesImport      = "github.com/hashicorp/terraform-plugin-go/tftypes"
	TypesImport        = "github.com/hashicorp/terraform-plugin-framework/types"
	ValidatorImport    = "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type Imports struct {
	imports []code.Import
	paths   map[string]struct{}
}

func NewImports() *Imports {
	return &Imports{
		imports: []code.Import{},
		paths:   make(map[string]struct{}),
	}
}

func (i *Imports) Add(c ...code.Import) {
	for _, imp := range c {
		if _, ok := i.paths[imp.Path]; ok {
			continue
		}

		i.imports = append(i.imports, imp)
		i.paths[imp.Path] = struct{}{}
	}
}

func (i *Imports) All() []code.Import {
	return i.imports
}

func (i *Imports) Append(imps ...*Imports) {
	for _, imp := range imps {
		for _, c := range imp.imports {
			if _, ok := i.paths[c.Path]; ok {
				continue
			}

			i.imports = append(i.imports, c)
			i.paths[c.Path] = struct{}{}
		}
	}
}

func AttrImports() *Imports {
	imports := NewImports()

	imports.Add(code.Import{
		Path: AttrImport,
	})

	return imports
}

func AssociatedExternalTypeImports() *Imports {
	imports := NewImports()

	imports.Add([]code.Import{
		{
			Path: FmtImport,
		},
		{
			Path: DiagImport,
		},
		{
			Path: AttrImport,
		},
		{
			Path: TfTypesImport,
		},
		{
			Path: BaseTypesImport,
		},
	}...)

	return imports
}
