// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertSetNestedBlock(b *provider.SetNestedBlock) (provider_generate.GeneratorSetNestedBlock, error) {
	if b == nil {
		return provider_generate.GeneratorSetNestedBlock{}, fmt.Errorf("*provider.SetNestedBlock is nil")
	}

	attributes := make(map[string]provider_generate.GeneratorAttribute, len(b.NestedObject.Attributes))

	for _, v := range b.NestedObject.Attributes {
		var attribute provider_generate.GeneratorAttribute
		var err error

		switch {
		case v.Bool != nil:
			attribute, err = convertBoolAttribute(v.Bool)
		case v.Float64 != nil:
			attribute, err = convertFloat64Attribute(v.Float64)
		case v.Int64 != nil:
			attribute, err = convertInt64Attribute(v.Int64)
		case v.List != nil:
			attribute, err = convertListAttribute(v.List)
		case v.ListNested != nil:
			attribute, err = convertListNestedAttribute(v.ListNested)
		case v.Map != nil:
			attribute, err = convertMapAttribute(v.Map)
		case v.MapNested != nil:
			attribute, err = convertMapNestedAttribute(v.MapNested)
		case v.Number != nil:
			attribute, err = convertNumberAttribute(v.Number)
		case v.Object != nil:
			attribute, err = convertObjectAttribute(v.Object)
		case v.Set != nil:
			attribute, err = convertSetAttribute(v.Set)
		case v.SetNested != nil:
			attribute, err = convertSetNestedAttribute(v.SetNested)
		case v.SingleNested != nil:
			attribute, err = convertSingleNestedAttribute(v.SingleNested)
		case v.String != nil:
			attribute, err = convertStringAttribute(v.String)
		default:
			return provider_generate.GeneratorSetNestedBlock{}, fmt.Errorf("attribute type is not defined: %+v", v)
		}

		if err != nil {
			return provider_generate.GeneratorSetNestedBlock{}, err
		}

		attributes[v.Name] = attribute
	}

	blocks := make(map[string]provider_generate.GeneratorBlock, len(b.NestedObject.Blocks))

	for _, v := range b.NestedObject.Blocks {
		var block provider_generate.GeneratorBlock
		var err error

		switch {
		case v.ListNested != nil:
			block, err = convertListNestedBlock(v.ListNested)
		case v.SetNested != nil:
			block, err = convertSetNestedBlock(v.SetNested)
		case v.SingleNested != nil:
			block, err = convertSingleNestedBlock(v.SingleNested)
		default:
			return provider_generate.GeneratorSetNestedBlock{}, fmt.Errorf("block type is not defined: %+v", v)
		}

		if err != nil {
			return provider_generate.GeneratorSetNestedBlock{}, err
		}

		blocks[v.Name] = block
	}

	return provider_generate.GeneratorSetNestedBlock{
		SetNestedBlock: schema.SetNestedBlock{
			Description:         description(b.Description),
			MarkdownDescription: description(b.Description),
			DeprecationMessage:  deprecationMessage(b.DeprecationMessage),
		},

		CustomType: b.CustomType,
		NestedObject: provider_generate.GeneratorNestedBlockObject{
			Attributes: attributes,
			Blocks:     blocks,
			CustomType: b.NestedObject.CustomType,
			Validators: b.NestedObject.Validators,
		},
		Validators: b.Validators,
	}, nil
}
