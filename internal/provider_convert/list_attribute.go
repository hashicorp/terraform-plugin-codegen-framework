// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertListAttribute(a *provider.ListAttribute) (provider_generate.GeneratorListAttribute, error) {
	if a == nil {
		return provider_generate.GeneratorListAttribute{}, fmt.Errorf("*provider.ListAttribute is nil")
	}

	return provider_generate.GeneratorListAttribute{
		ListAttribute: schema.ListAttribute{
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
