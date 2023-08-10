// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func convertListNestedAttribute(a *resource.ListNestedAttribute) (resource_generate.GeneratorListNestedAttribute, error) {
	if a == nil {
		return resource_generate.GeneratorListNestedAttribute{}, fmt.Errorf("*resource.ListNestedAttribute is nil")
	}

	attributes := make(generatorschema.GeneratorAttributes, len(a.NestedObject.Attributes))

	for _, v := range a.NestedObject.Attributes {
		var attribute generatorschema.GeneratorAttribute
		var err error

		switch {
		case v.Bool != nil:
			attribute, err = convertBoolAttribute(v.Bool)
		case v.Float64 != nil:
			attribute, err = convertFloat64Attribute(v.Float64)
		case v.Int64 != nil:
			attribute, err = convertInt64Attribute(v.Int64)
		case v.List != nil:
			attribute, err = convertListAttribute(v.List)
		case v.ListNested != nil:
			attribute, err = convertListNestedAttribute(v.ListNested)
		case v.Map != nil:
			attribute, err = convertMapAttribute(v.Map)
		case v.MapNested != nil:
			attribute, err = convertMapNestedAttribute(v.MapNested)
		case v.Number != nil:
			attribute, err = convertNumberAttribute(v.Number)
		case v.Object != nil:
			attribute, err = convertObjectAttribute(v.Object)
		case v.Set != nil:
			attribute, err = convertSetAttribute(v.Set)
		case v.SetNested != nil:
			attribute, err = convertSetNestedAttribute(v.SetNested)
		case v.SingleNested != nil:
			attribute, err = convertSingleNestedAttribute(v.SingleNested)
		case v.String != nil:
			attribute, err = convertStringAttribute(v.String)
		default:
			return resource_generate.GeneratorListNestedAttribute{}, fmt.Errorf("attribute type not defined: %+v", v)
		}

		if err != nil {
			return resource_generate.GeneratorListNestedAttribute{}, err
		}

		attributes[v.Name] = attribute
	}

	return resource_generate.GeneratorListNestedAttribute{
		ListNestedAttribute: schema.ListNestedAttribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},
		CustomType: a.CustomType,
		Default:    a.Default,
		NestedObject: resource_generate.GeneratorNestedAttributeObject{
			AssociatedExternalType: a.NestedObject.AssociatedExternalType,
			Attributes:             attributes,
			CustomType:             a.NestedObject.CustomType,
			PlanModifiers:          a.NestedObject.PlanModifiers,
			Validators:             a.NestedObject.Validators,
		},
		PlanModifiers: a.PlanModifiers,
		Validators:    a.Validators,
	}, nil
}
