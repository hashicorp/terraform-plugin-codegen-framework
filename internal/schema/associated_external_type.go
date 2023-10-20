// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type AssocExtType struct {
	*schema.AssociatedExternalType
}

func NewAssocExtType(assocExtType *schema.AssociatedExternalType) *AssocExtType {
	if assocExtType == nil {
		return nil
	}

	return &AssocExtType{
		AssociatedExternalType: assocExtType,
	}
}

func (a *AssocExtType) Imports() *Imports {
	imports := NewImports()

	if a == nil {
		return imports
	}

	if a.AssociatedExternalType.Import == nil {
		return imports
	}

	if len(a.AssociatedExternalType.Import.Path) > 0 {
		imports.Add(*a.AssociatedExternalType.Import)

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

	return a.AssociatedExternalType.Type
}

func (a *AssocExtType) TypeReference() string {
	if a == nil {
		return ""
	}

	tr, _ := strings.CutPrefix(a.AssociatedExternalType.Type, "*")

	return tr
}

func (a *AssocExtType) Equal(other *AssocExtType) bool {
	if a == nil && other == nil {
		return true
	}

	if a == nil || other == nil {
		return false
	}

	return a.AssociatedExternalType.Equal(other.AssociatedExternalType)
}

func (a *AssocExtType) ToPascalCase() string {
	inputSplit := strings.Split(a.TypeReference(), ".")

	var ucName string

	for _, v := range inputSplit {
		if len(v) < 1 {
			continue
		}

		firstChar := v[0:1]
		ucFirstChar := strings.ToUpper(firstChar)

		if len(v) < 2 {
			ucName += ucFirstChar
			continue
		}

		ucName += ucFirstChar + v[1:]
	}

	return ucName
}

func (a *AssocExtType) ToCamelCase() string {
	pascal := a.ToPascalCase()

	// Grab first rune and lower case it
	firstLetter, size := utf8.DecodeRuneInString(pascal)
	if firstLetter == utf8.RuneError && size <= 1 {
		return pascal
	}

	return string(unicode.ToLower(firstLetter)) + pascal[size:]
}
