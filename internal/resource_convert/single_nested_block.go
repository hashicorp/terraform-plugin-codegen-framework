package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func convertSingleNestedBlock(b *resource.SingleNestedBlock) (resource_generate.GeneratorSingleNestedBlock, error) {
	if b == nil {
		return resource_generate.GeneratorSingleNestedBlock{}, fmt.Errorf("*resource.SingleNestedBlock is nil")
	}

	attributes := make(map[string]resource_generate.GeneratorAttribute, len(b.Attributes))

	for _, v := range b.Attributes {
		var attribute resource_generate.GeneratorAttribute
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
			return resource_generate.GeneratorSingleNestedBlock{}, fmt.Errorf("attribute type is not defined: %+v", v)
		}

		if err != nil {
			return resource_generate.GeneratorSingleNestedBlock{}, err
		}

		attributes[v.Name] = attribute
	}

	blocks := make(map[string]resource_generate.GeneratorBlock, len(b.Blocks))

	for _, v := range b.Blocks {
		var block resource_generate.GeneratorBlock
		var err error

		switch {
		case v.ListNested != nil:
			block, err = convertListNestedBlock(v.ListNested)
		case v.SetNested != nil:
			block, err = convertSetNestedBlock(v.SetNested)
		case v.SingleNested != nil:
			block, err = convertSingleNestedBlock(v.SingleNested)
		default:
			return resource_generate.GeneratorSingleNestedBlock{}, fmt.Errorf("block type is not defined: %+v", v)
		}

		if err != nil {
			return resource_generate.GeneratorSingleNestedBlock{}, err
		}

		blocks[v.Name] = block
	}

	return resource_generate.GeneratorSingleNestedBlock{
		SingleNestedBlock: schema.SingleNestedBlock{
			Description:         description(b.Description),
			MarkdownDescription: description(b.Description),
			DeprecationMessage:  deprecationMessage(b.DeprecationMessage),
		},
		Attributes:    attributes,
		Blocks:        blocks,
		CustomType:    b.CustomType,
		PlanModifiers: b.PlanModifiers,
		Validators:    b.Validators,
	}, nil
}
