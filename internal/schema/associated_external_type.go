// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type AssocExtType struct {
	imp *code.Import
	typ string
}

func NewAssocExtType(assocExtType *schema.AssociatedExternalType) *AssocExtType {
	if assocExtType == nil {
		return nil
	}

	return &AssocExtType{
		imp: assocExtType.Import,
		typ: assocExtType.Type,
	}
}

func (a *AssocExtType) Imports() *Imports {
	imports := NewImports()

	if a == nil {
		return imports
	}

	if a.imp == nil {
		return imports
	}

	if len(a.imp.Path) > 0 {
		imports.Add(*a.imp)

		imports.Add(code.Import{
			Path: BaseTypesImport,
		})
	}

	return imports
}

func (a *AssocExtType) Type() string {
	if a == nil {
		return ""
	}

	return a.typ
}

func (a *AssocExtType) TypeReference() string {
	if a == nil {
		return ""
	}

	tr, _ := strings.CutPrefix(a.typ, "*")

	return tr
}
