// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
)

const (
	CustomCollectionTypeList CustomCollectionTypes = "List"
	CustomCollectionTypeMap  CustomCollectionTypes = "Map"
	CustomCollectionTypeSet  CustomCollectionTypes = "Set"
)

type CustomCollectionTypes string

type CustomTypeCollection struct {
	customType string
}

// NewCustomTypeCollection constructs an CustomTypeCollection which is used to determine whether a CustomType
// should be assigned to a collection attribute in the schema.
//
// If a CustomType has been declared in the spec, then the CustomType.Type will be used as
// the CustomType in the schema.
//
// If the spec CustomType is nil, and the spec AssociatedExternalType is not nil, the generator
// will create custom Type and Value types using the attribute name, and the generated custom
// Type type will be used as the CustomType in the schema.
func NewCustomTypeCollection(c *specschema.CustomType, a *specschema.AssociatedExternalType, cct CustomCollectionTypes, elemType, name string) CustomTypeCollection {
	var customType string

	switch {
	case c != nil:
		customType = c.Type
	case a != nil:
		customType = fmt.Sprintf("%sType{\ntypes.%sType{\nElemType: %s,\n},\n}", format.ToPascalCase(name), cct, elemType)
	}

	return CustomTypeCollection{
		customType: customType,
	}
}

func (a CustomTypeCollection) Equal(other CustomTypeCollection) bool {
	return a.customType == other.customType
}

func (a CustomTypeCollection) Schema() []byte {
	if a.customType != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", a.customType))
	}

	return nil
}
