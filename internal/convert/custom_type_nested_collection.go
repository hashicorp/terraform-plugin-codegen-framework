// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type CustomTypeNestedCollection struct {
	customType *specschema.CustomType
}

// NewCustomTypeNestedCollection constructs an CustomTypeNestedCollection which is used to determine whether a CustomType
// should be assigned to a nested attribute in the schema.
//
// If a CustomType has been declared in the spec, then the CustomType.Type will be used as
// the CustomType in the schema.
//
// If the spec CustomType is nil, the generator will create custom Type and Value types using the attribute
// name, and the generated custom Type type will be used as the CustomType in the schema.
func NewCustomTypeNestedCollection(c *specschema.CustomType) CustomTypeNestedCollection {
	return CustomTypeNestedCollection{
		customType: c,
	}
}

func (a CustomTypeNestedCollection) Equal(other CustomTypeNestedCollection) bool {
	return a.customType.Equal(other.customType)
}

func (a CustomTypeNestedCollection) Schema() []byte {
	if a.customType != nil && a.customType.Type != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", a.customType.Type))
	}

	return nil
}
