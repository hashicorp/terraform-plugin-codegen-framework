// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
)

type CustomTypeNested struct {
	customType string
}

// NewCustomTypeNested constructs an CustomTypeNested which is used to determine whether a CustomType
// should be assigned to a nested attribute in the schema.
//
// If a CustomType has been declared in the spec, then the CustomType.Type will be used as
// the CustomType in the schema.
//
// If the spec CustomType is nil, the generator will create custom Type and Value types using the attribute
// name, and the generated custom Type type will be used as the CustomType in the schema.
func NewCustomTypeNested(c *specschema.CustomType, name string) CustomTypeNested {
	var customType string

	switch {
	case c != nil:
		customType = c.Type
	default:
		customType = fmt.Sprintf("%sType{\nObjectType: types.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n},\n}", format.ToPascalCase(name), format.ToPascalCase(name))
	}

	return CustomTypeNested{
		customType: customType,
	}
}

func (a CustomTypeNested) Equal(other CustomTypeNested) bool {
	return a.customType == other.customType
}

func (a CustomTypeNested) Schema() []byte {
	if a.customType != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", a.customType))
	}

	return nil
}
