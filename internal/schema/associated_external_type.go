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

func (a AssocExtType) Type() string {
	return a.typ
}

func (a AssocExtType) TypeReference() string {
	tr, _ := strings.CutPrefix(a.typ, "*")

	return tr
}
