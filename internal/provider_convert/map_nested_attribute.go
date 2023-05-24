package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertMapNestedAttribute(a *provider.MapNestedAttribute) (provider_generate.GeneratorMapNestedAttribute, error) {
	if a == nil {
		return provider_generate.GeneratorMapNestedAttribute{}, fmt.Errorf("*provider.MapNestedAttribute is nil")
	}

	attributes := make(map[string]provider_generate.GeneratorAttribute, len(a.NestedObject.Attributes))

	for _, v := range a.NestedObject.Attributes {
		var attribute provider_generate.GeneratorAttribute
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
			return provider_generate.GeneratorMapNestedAttribute{}, fmt.Errorf("attribute type not defined: %+v", v)
		}

		if err != nil {
			return provider_generate.GeneratorMapNestedAttribute{}, err
		}

		attributes[v.Name] = attribute
	}

	return provider_generate.GeneratorMapNestedAttribute{
		MapNestedAttribute: schema.MapNestedAttribute{
			Required:            isRequired(a.OptionalRequired),
			Optional:            isOptional(a.OptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		CustomType: a.CustomType,
		NestedObject: provider_generate.GeneratorNestedAttributeObject{
			Attributes: attributes,
			CustomType: a.NestedObject.CustomType,
			Validators: a.NestedObject.Validators,
		},
		Validators: a.Validators,
	}, nil
}
