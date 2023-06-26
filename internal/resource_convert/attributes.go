// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
)

func convertAttribute(a resource.Attribute) (resource_generate.GeneratorAttribute, error) {
	switch {
	case a.Bool != nil:
		return convertBoolAttribute(a.Bool)
	case a.Float64 != nil:
		return convertFloat64Attribute(a.Float64)
	case a.Int64 != nil:
		return convertInt64Attribute(a.Int64)
	case a.List != nil:
		return convertListAttribute(a.List)
	case a.ListNested != nil:
		return convertListNestedAttribute(a.ListNested)
	case a.Map != nil:
		return convertMapAttribute(a.Map)
	case a.MapNested != nil:
		return convertMapNestedAttribute(a.MapNested)
	case a.Number != nil:
		return convertNumberAttribute(a.Number)
	case a.Object != nil:
		return convertObjectAttribute(a.Object)
	case a.Set != nil:
		return convertSetAttribute(a.Set)
	case a.SetNested != nil:
		return convertSetNestedAttribute(a.SetNested)
	case a.SingleNested != nil:
		return convertSingleNestedAttribute(a.SingleNested)
	case a.String != nil:
		return convertStringAttribute(a.String)
	}

	return nil, fmt.Errorf("attribute type not defined: %+v", a)
}
