// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func convertSingleNestedAttribute(a *provider.SingleNestedAttribute) (provider_generate.GeneratorSingleNestedAttribute, error) {
	if a == nil {
		return provider_generate.GeneratorSingleNestedAttribute{}, fmt.Errorf("*provider.SingleNestedAttribute is nil")
	}

	attributes := make(generatorschema.GeneratorAttributes, len(a.Attributes))

	for _, v := range a.Attributes {
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
			return provider_generate.GeneratorSingleNestedAttribute{}, fmt.Errorf("attribute type not defined: %+v", v)
		}

		if err != nil {
			return provider_generate.GeneratorSingleNestedAttribute{}, err
		}

		attributes[v.Name] = attribute
	}

	return provider_generate.GeneratorSingleNestedAttribute{
		SingleNestedAttribute: schema.SingleNestedAttribute{
			Required:            isRequired(a.OptionalRequired),
			Optional:            isOptional(a.OptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},
		AssociatedExternalType: generatorschema.NewAssocExtType(a.AssociatedExternalType),
		Attributes:             attributes,
		CustomType:             a.CustomType,
		Validators:             a.Validators,
	}, nil
}
