// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func convertSetNestedAttribute(a *datasource.SetNestedAttribute) (datasource_generate.GeneratorSetNestedAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorSetNestedAttribute{}, fmt.Errorf("*datasource.SetNestedAttribute is nil")
	}

	attributes := make(map[string]datasource_generate.GeneratorAttribute, len(a.NestedObject.Attributes))

	for _, v := range a.NestedObject.Attributes {
		var attribute datasource_generate.GeneratorAttribute
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
			return datasource_generate.GeneratorSetNestedAttribute{}, fmt.Errorf("attribute type not defined: %+v", v)
		}

		if err != nil {
			return datasource_generate.GeneratorSetNestedAttribute{}, err
		}

		attributes[v.Name] = attribute
	}

	return datasource_generate.GeneratorSetNestedAttribute{
		SetNestedAttribute: schema.SetNestedAttribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		CustomType: a.CustomType,
		NestedObject: datasource_generate.GeneratorNestedAttributeObject{
			Attributes: attributes,
			CustomType: a.NestedObject.CustomType,
			Validators: a.NestedObject.Validators,
		},
		Validators: a.Validators,
	}, nil
}
