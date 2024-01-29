// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type NestedBlockObject struct {
	attributes schema.GeneratorAttributes
	blocks     schema.GeneratorBlocks
	customType CustomTypeNestedObject
	validators Validators
}

// NewNestedBlockObject constructs a NestedBlockObject which is used to generate a
// nested attribute block in the schema.
func NewNestedBlockObject(a schema.GeneratorAttributes, b schema.GeneratorBlocks, c *specschema.CustomType, v Validators, name string) NestedBlockObject {
	return NestedBlockObject{
		attributes: a,
		blocks:     b,
		customType: NewCustomTypeNestedObject(c, name),
		validators: v,
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

	return n.validators.Equal(other.validators)
}

func (n NestedBlockObject) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(n.customType.Imports())

	imports.Append(n.validators.Imports())

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
	b.Write(n.validators.Schema())
	b.WriteString("},\n")

	return b.Bytes(), nil
}
