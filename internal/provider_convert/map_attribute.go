// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertMapAttribute(a *provider.MapAttribute) (provider_generate.GeneratorMapAttribute, error) {
	if a == nil {
		return provider_generate.GeneratorMapAttribute{}, fmt.Errorf("*provider.MapAttribute is nil")
	}

	return provider_generate.GeneratorMapAttribute{
		MapAttribute: schema.MapAttribute{
			Required:            isRequired(a.OptionalRequired),
			Optional:            isOptional(a.OptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		CustomType:  a.CustomType,
		ElementType: a.ElementType,
		Validators:  a.Validators,
	}, nil
}
