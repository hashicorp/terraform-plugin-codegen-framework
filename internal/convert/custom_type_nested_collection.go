// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
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

func (c CustomTypeNestedCollection) Equal(other CustomTypeNestedCollection) bool {
	return c.customType.Equal(other.customType)
}

func (c CustomTypeNestedCollection) Imports() *schema.Imports {
	imports := schema.NewImports()

	if c.customType != nil {
		if c.customType.HasImport() {
			imports.Add(*c.customType.Import)
		}
	} else {
		imports.Add(code.Import{
			Path: schema.TypesImport,
		})
	}

	return imports
}

func (c CustomTypeNestedCollection) Schema() []byte {
	if c.customType != nil && c.customType.Type != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", c.customType.Type))
	}

	return nil
}

func (c CustomTypeNestedCollection) ValueType() string {
	if c.customType != nil {
		return c.customType.ValueType
	}

	return ""
}
