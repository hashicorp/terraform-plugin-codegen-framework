// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type NestedBlockObject struct {
	attributes       schema.GeneratorAttributes
	blocks           schema.GeneratorBlocks
	customType       CustomTypeNestedObject
	validatorsCustom ValidatorsCustom
}

// NewNestedBlockObject constructs a NestedBlockObject which is used to generate a
// nested attribute block in the schema.
func NewNestedBlockObject(a schema.GeneratorAttributes, b schema.GeneratorBlocks, c *specschema.CustomType, v ValidatorsCustom, name string) NestedBlockObject {
	return NestedBlockObject{
		attributes:       a,
		blocks:           b,
		customType:       NewCustomTypeNestedObject(c, name),
		validatorsCustom: v,
	}
}

func (n NestedBlockObject) Equal(other NestedBlockObject) bool {
	if !n.attributes.Equal(other.attributes) {
		return false
	}

	if !n.blocks.Equal(other.blocks) {
		return false
	}

	if !n.customType.Equal(other.customType) {
		return false
	}

	return n.validatorsCustom.Equal(other.validatorsCustom)
}

func (n NestedBlockObject) Imports() *schema.Imports {
	imports := schema.NewImports()

	if n.customType.customType != nil {
		if n.customType.customType.HasImport() {
			imports.Add(*n.customType.customType.Import)
		}
	} else {
		imports.Add(code.Import{
			Path: schema.TypesImport,
		})
	}

	imports.Append(n.validatorsCustom.Imports())

	imports.Append(n.attributes.Imports())

	imports.Append(n.blocks.Imports())

	return imports
}

func (n NestedBlockObject) Schema() ([]byte, error) {
	var b bytes.Buffer

	attributesSchema, err := n.attributes.Schema()

	if err != nil {
		return nil, err
	}

	blocksSchema, err := n.blocks.Schema()

	if err != nil {
		return nil, err
	}

	b.WriteString("NestedObject: schema.NestedBlockObject{\n")
	if attributesSchema != "" {
		b.WriteString("Attributes: map[string]schema.Attribute{")
		b.WriteString(attributesSchema)
		b.WriteString("\n},\n")
	}
	if blocksSchema != "" {
		b.WriteString("Blocks: map[string]schema.Block{")
		b.WriteString(blocksSchema)
		b.WriteString("\n},\n")
	}
	b.Write(n.customType.Schema())
	b.Write(n.validatorsCustom.Schema())
	b.WriteString("},\n")

	return b.Bytes(), nil
}
