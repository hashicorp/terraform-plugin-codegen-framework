// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
)

type CustomTypePrimitive struct {
	associatedExternalType *specschema.AssociatedExternalType
	customType             *specschema.CustomType
	name                   string
}

// NewCustomTypePrimitive constructs an CustomTypePrimitive which is used to determine whether a CustomType
// should be assigned to a primitive attribute in the schema.
//
// If a CustomType has been declared in the spec, then the CustomType.Type will be used as
// the CustomType in the Schema.
//
// If the spec CustomType is nil, and the spec AssociatedExternalType is not nil, the generator
// will create custom Type and Value types using the attribute name, and the generated custom
// Type type will be used as the CustomType in the schema.
func NewCustomTypePrimitive(c *specschema.CustomType, a *specschema.AssociatedExternalType, name string) CustomTypePrimitive {
	return CustomTypePrimitive{
		associatedExternalType: a,
		customType:             c,
		name:                   name,
	}
}

func (c CustomTypePrimitive) Equal(other CustomTypePrimitive) bool {
	if !c.associatedExternalType.Equal(other.associatedExternalType) {
		return false
	}

	if !c.customType.Equal(other.customType) {
		return false
	}

	return c.name == other.name
}

func (c CustomTypePrimitive) Schema() []byte {
	var customType string

	switch {
	case c.customType != nil:
		customType = c.customType.Type
	case c.associatedExternalType != nil:
		customType = fmt.Sprintf("%sType{}", format.ToPascalCase(c.name))
	}

	if customType != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", customType))
	}

	return nil
}
