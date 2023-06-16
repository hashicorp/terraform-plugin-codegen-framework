// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
)

func convertBlock(b provider.Block) (provider_generate.GeneratorBlock, error) {
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
