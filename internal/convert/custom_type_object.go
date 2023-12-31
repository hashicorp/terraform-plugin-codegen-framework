// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
)

type CustomTypeObject struct {
	associatedExternalType *specschema.AssociatedExternalType
	customType             *specschema.CustomType
	name                   string
}

// NewCustomTypeObject constructs an CustomTypeObject which is used to determine whether a CustomType
// should be assigned to an object attribute in the schema.
//
// If a CustomType has been declared in the spec, then the CustomType.Type will be used as
// the CustomType in the schema.
//
// If the spec CustomType is nil, and the spec AssociatedExternalType is not nil, the generator
// will create custom Type and Value types using the attribute name, and the generated custom
// Type type will be used as the CustomType in the schema.
func NewCustomTypeObject(c *specschema.CustomType, a *specschema.AssociatedExternalType, name string) CustomTypeObject {
	return CustomTypeObject{
		associatedExternalType: a,
		customType:             c,
		name:                   name,
	}
}

func (c CustomTypeObject) Equal(other CustomTypeObject) bool {
	if !c.associatedExternalType.Equal(other.associatedExternalType) {
		return false
	}

	if !c.customType.Equal(other.customType) {
		return false
	}

	return c.name == other.name
}

func (c CustomTypeObject) Schema() []byte {
	var customType string

	switch {
	case c.customType != nil:
		customType = c.customType.Type
	case c.associatedExternalType != nil:
		customType = fmt.Sprintf("%sType{\ntypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n},\n}", format.ToPascalCase(c.name), format.ToPascalCase(c.name))
	}

	if customType != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", customType))
	}

	return nil
}
