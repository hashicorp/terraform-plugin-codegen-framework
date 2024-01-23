// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert

import (
	"bytes"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type NestedAttributeObject struct {
	attributes       schema.GeneratorAttributes
	customType       CustomTypeNestedObject
	validatorsCustom ValidatorsCustom
}

// NewNestedAttributeObject constructs a NestedAttributeObject which is used to generate a
// nested attribute object in the schema.
func NewNestedAttributeObject(a schema.GeneratorAttributes, c *specschema.CustomType, v ValidatorsCustom, name string) NestedAttributeObject {
	return NestedAttributeObject{
		attributes:       a,
		customType:       NewCustomTypeNestedObject(c, name),
		validatorsCustom: v,
	}
}

func (n NestedAttributeObject) Equal(other NestedAttributeObject) bool {
	if !n.attributes.Equal(other.attributes) {
		return false
	}

	if !n.customType.Equal(other.customType) {
		return false
	}

	return n.validatorsCustom.Equal(other.validatorsCustom)
}

func (n NestedAttributeObject) Imports() *schema.Imports {
	imports := schema.NewImports()

	imports.Append(n.customType.Imports())

	imports.Append(n.validatorsCustom.Imports())

	imports.Append(n.attributes.Imports())

	return imports
}

func (n NestedAttributeObject) Schema() ([]byte, error) {
	var b bytes.Buffer

	attributesSchema, err := n.attributes.Schema()

	if err != nil {
		return nil, err
	}

	b.WriteString("NestedObject: schema.NestedAttributeObject{\n")
	b.WriteString("Attributes: map[string]schema.Attribute{")
	b.WriteString(attributesSchema)
	b.WriteString("\n},\n")
	b.Write(n.customType.Schema())
	b.Write(n.validatorsCustom.Schema())
	b.WriteString("},\n")

	return b.Bytes(), nil
}
