// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func convertBlock(b datasource.Block) (datasource_generate.GeneratorBlock, error) {
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
