// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type NestedBlockObject struct {
	attributes          generatorschema.GeneratorAttributes
	blocks              generatorschema.GeneratorBlocks
	customType          convert.CustomTypeNestedObject
	planModifiersCustom convert.PlanModifiersCustom
	validatorsCustom    convert.ValidatorsCustom
}

// NewNestedBlockObject constructs a NestedBlockObject which is used to generate a
// nested attribute block in the schema.
func NewNestedBlockObject(a generatorschema.GeneratorAttributes, b generatorschema.GeneratorBlocks, c *specschema.CustomType, p convert.PlanModifiersCustom, v convert.ValidatorsCustom, name string) NestedBlockObject {
	return NestedBlockObject{
		attributes:          a,
		blocks:              b,
		customType:          convert.NewCustomTypeNestedObject(c, name),
		planModifiersCustom: p,
		validatorsCustom:    v,
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

	if !n.planModifiersCustom.Equal(other.planModifiersCustom) {
		return false
	}

	return n.validatorsCustom.Equal(other.validatorsCustom)
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
	b.Write(n.planModifiersCustom.Schema())
	b.Write(n.validatorsCustom.Schema())
	b.WriteString("},\n")

	return b.Bytes(), nil
}
