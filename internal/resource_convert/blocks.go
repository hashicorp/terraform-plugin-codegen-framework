package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func convertBlock(b resource.Block) (resource_generate.GeneratorBlock, error) {
	switch {
	case b.ListNested != nil:
		return convertListNestedBlock(b.ListNested)
	case b.SetNested != nil:
		return convertSetNestedBlock(b.SetNested)
	case b.SingleNested != nil:
		return convertSingleNestedBlock(b.SingleNested)
	}

	return nil, fmt.Errorf("block type not defined: %+v", b)
}
