// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/format"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

const (
	CustomCollectionTypeList CustomCollectionTypes = "List"
	CustomCollectionTypeMap  CustomCollectionTypes = "Map"
	CustomCollectionTypeSet  CustomCollectionTypes = "Set"
)

type CustomCollectionTypes string

func (c CustomCollectionTypes) Equal(other CustomCollectionTypes) bool {
	return c == other
}

type CustomTypeCollection struct {
	associatedExternalType *specschema.AssociatedExternalType
	customCollectionType   CustomCollectionTypes
	customType             *specschema.CustomType
	elementType            string
	name                   string
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
	return CustomTypeCollection{
		associatedExternalType: a,
		customCollectionType:   cct,
		customType:             c,
		elementType:            elemType,
		name:                   name,
	}
}

func (c CustomTypeCollection) Equal(other CustomTypeCollection) bool {
	if !c.associatedExternalType.Equal(other.associatedExternalType) {
		return false
	}

	if !c.customCollectionType.Equal(other.customCollectionType) {
		return false
	}

	if !c.customType.Equal(other.customType) {
		return false
	}

	if c.elementType != other.elementType {
		return false
	}

	return c.name == other.name
}

func (c CustomTypeCollection) Imports() *schema.Imports {
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

func (c CustomTypeCollection) Schema() []byte {
	var customType string

	switch {
	case c.customType != nil:
		customType = c.customType.Type
	case c.associatedExternalType != nil:
		customType = fmt.Sprintf("%sType{\ntypes.%sType{\nElemType: %s,\n},\n}", format.ToPascalCase(c.name), c.customCollectionType, c.elementType)
	}

	if customType != "" {
		return []byte(fmt.Sprintf("CustomType: %s,\n", customType))
	}

	return nil
}

func (c CustomTypeCollection) ValueType() string {
	switch {
	case c.customType != nil:
		return c.customType.ValueType
	case c.associatedExternalType != nil:
		return fmt.Sprintf("%sValue", format.ToPascalCase(c.name))
	}

	return ""
}
