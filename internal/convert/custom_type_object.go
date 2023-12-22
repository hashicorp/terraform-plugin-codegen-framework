// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
)

type CustomTypeObject struct {
	customType string
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
	var customType string

	switch {
	case c != nil:
		customType = c.Type
	case a != nil:
		customType = fmt.Sprintf("%sType{\ntypes.ObjectType{\nAttrTypes: %sValue{}.AttributeTypes(ctx),\n},\n}", format.ToPascalCase(name), format.ToPascalCase(name))
	}

	return CustomTypeObject{
		customType: customType,
	}
}

func (a CustomTypeObject) Equal(other CustomTypeObject) bool {
	return a.customType == other.customType
}

func (a CustomTypeObject) Schema() []byte {
	if a.customType != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", a.customType))
	}

	return nil
}
