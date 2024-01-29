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
	attributes    generatorschema.GeneratorAttributes
	blocks        generatorschema.GeneratorBlocks
	customType    convert.CustomTypeNestedObject
	planModifiers convert.PlanModifiers
	validators    convert.Validators
}

// NewNestedBlockObject constructs a NestedBlockObject which is used to generate a
// nested attribute block in the schema.
func NewNestedBlockObject(a generatorschema.GeneratorAttributes, b generatorschema.GeneratorBlocks, c *specschema.CustomType, p convert.PlanModifiers, v convert.Validators, name string) NestedBlockObject {
	return NestedBlockObject{
		attributes:    a,
		blocks:        b,
		customType:    convert.NewCustomTypeNestedObject(c, name),
		planModifiers: p,
		validators:    v,
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

	if !n.planModifiers.Equal(other.planModifiers) {
		return false
	}

	return n.validators.Equal(other.validators)
}

func (n NestedBlockObject) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(n.customType.Imports())

	imports.Append(n.planModifiers.Imports())

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
	b.Write(n.planModifiers.Schema())
	b.Write(n.validators.Schema())
	b.WriteString("},\n")

	return b.Bytes(), nil
}
