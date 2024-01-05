// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func NewSchemas(spec spec.Specification) (map[string]generatorschema.GeneratorSchema, error) {
	resourceSchemas := make(map[string]generatorschema.GeneratorSchema, len(spec.Resources))

	for _, v := range spec.Resources {
		s, err := NewSchema(v)
		if err != nil {
			return nil, err
		}

		resourceSchemas[v.Name] = s
	}

	return resourceSchemas, nil
}

func NewSchema(d resource.Resource) (generatorschema.GeneratorSchema, error) {
	var s generatorschema.GeneratorSchema

	attributes := make(generatorschema.GeneratorAttributes, len(d.Schema.Attributes))
	blocks := make(generatorschema.GeneratorBlocks, len(d.Schema.Blocks))

	for _, v := range d.Schema.Attributes {
		a, err := NewAttribute(v)

		if err != nil {
			return s, err
		}

		attributes[v.Name] = a
	}

	s.Attributes = attributes

	for _, v := range d.Schema.Blocks {
		b, err := NewBlock(v)

		if err != nil {
			return s, err
		}

		blocks[v.Name] = b
	}

	s.Blocks = blocks

	return s, nil
}

func NewAttribute(a resource.Attribute) (generatorschema.GeneratorAttribute, error) {
	switch {
	case a.Bool != nil:
		return NewGeneratorBoolAttribute(a.Name, a.Bool)
	case a.Float64 != nil:
		return NewGeneratorFloat64Attribute(a.Name, a.Float64)
	case a.Int64 != nil:
		return NewGeneratorInt64Attribute(a.Name, a.Int64)
	case a.List != nil:
		return NewGeneratorListAttribute(a.Name, a.List)
	case a.ListNested != nil:
		return NewGeneratorListNestedAttribute(a.ListNested)
	case a.Map != nil:
		return NewGeneratorMapAttribute(a.Name, a.Map)
	case a.MapNested != nil:
		return NewGeneratorMapNestedAttribute(a.MapNested)
	case a.Number != nil:
		return NewGeneratorNumberAttribute(a.Name, a.Number)
	case a.Object != nil:
		return NewGeneratorObjectAttribute(a.Object)
	case a.Set != nil:
		return NewGeneratorSetAttribute(a.Name, a.Set)
	case a.SetNested != nil:
		return NewGeneratorSetNestedAttribute(a.SetNested)
	case a.SingleNested != nil:
		return NewGeneratorSingleNestedAttribute(a.SingleNested)
	case a.String != nil:
		return NewGeneratorStringAttribute(a.Name, a.String)
	}

	return nil, fmt.Errorf("attribute type not defined: %+v", a)
}

func NewBlock(b resource.Block) (generatorschema.GeneratorBlock, error) {
	switch {
	case b.ListNested != nil:
		return NewGeneratorListNestedBlock(b.ListNested)
	case b.SetNested != nil:
		return NewGeneratorSetNestedBlock(b.SetNested)
	case b.SingleNested != nil:
		return NewGeneratorSingleNestedBlock(b.SingleNested)
	}

	return nil, fmt.Errorf("block type not defined: %+v", b)
}
