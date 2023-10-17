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

func convertInt64Attribute(a *provider.Int64Attribute) (provider_generate.GeneratorInt64Attribute, error) {
	if a == nil {
		return provider_generate.GeneratorInt64Attribute{}, fmt.Errorf("*provider.Int64Attribute is nil")
	}

	return provider_generate.GeneratorInt64Attribute{
		Int64Attribute: schema.Int64Attribute{
			Required:            isRequired(a.OptionalRequired),
			Optional:            isOptional(a.OptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		AssociatedExternalType: generatorschema.NewAssocExtType(a.AssociatedExternalType),
		CustomType:             a.CustomType,
		Validators:             a.Validators,
	}, nil
}
