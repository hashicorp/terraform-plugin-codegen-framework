// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type NestedAttributeObject struct {
	attributes          generatorschema.GeneratorAttributes
	customType          convert.CustomTypeNestedObject
	planModifiersCustom convert.PlanModifiers
	validatorsCustom    convert.Validators
}

// NewNestedAttributeObject constructs a NestedAttributeObject which is used to generate a
// nested attribute object in the schema.
func NewNestedAttributeObject(a generatorschema.GeneratorAttributes, c *specschema.CustomType, p convert.PlanModifiers, v convert.Validators, name string) NestedAttributeObject {
	return NestedAttributeObject{
		attributes:          a,
		customType:          convert.NewCustomTypeNestedObject(c, name),
		planModifiersCustom: p,
		validatorsCustom:    v,
	}
}

func (n NestedAttributeObject) Equal(other NestedAttributeObject) bool {
	if !n.attributes.Equal(other.attributes) {
		return false
	}

	if !n.customType.Equal(other.customType) {
		return false
	}

	if !n.planModifiersCustom.Equal(other.planModifiersCustom) {
		return false
	}

	return n.validatorsCustom.Equal(other.validatorsCustom)
}

func (n NestedAttributeObject) Imports() *generatorschema.Imports {
	imports := generatorschema.NewImports()

	imports.Append(n.customType.Imports())

	imports.Append(n.planModifiersCustom.Imports())

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
	b.Write(n.planModifiersCustom.Schema())
	b.Write(n.validatorsCustom.Schema())
	b.WriteString("},\n")

	return b.Bytes(), nil
}
