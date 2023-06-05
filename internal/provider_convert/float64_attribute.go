// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertFloat64Attribute(a *provider.Float64Attribute) (provider_generate.GeneratorFloat64Attribute, error) {
	if a == nil {
		return provider_generate.GeneratorFloat64Attribute{}, fmt.Errorf("*provider.Float64Attribute is nil")
	}

	return provider_generate.GeneratorFloat64Attribute{
		Float64Attribute: schema.Float64Attribute{
			Required:            isRequired(a.OptionalRequired),
			Optional:            isOptional(a.OptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},
		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}
