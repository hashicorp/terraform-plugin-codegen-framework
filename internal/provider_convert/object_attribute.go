// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertObjectAttribute(o *provider.ObjectAttribute) (provider_generate.GeneratorObjectAttribute, error) {
	if o == nil {
		return provider_generate.GeneratorObjectAttribute{}, fmt.Errorf("*provider.ObjectAttribute is nil")
	}

	attrTypes, err := convertAttrTypes(o.AttributeTypes)
	if err != nil {
		return provider_generate.GeneratorObjectAttribute{}, err
	}

	return provider_generate.GeneratorObjectAttribute{
		ObjectAttribute: schema.ObjectAttribute{
			AttributeTypes:      attrTypes,
			Required:            isRequired(o.OptionalRequired),
			Optional:            isOptional(o.OptionalRequired),
			Sensitive:           isSensitive(o.Sensitive),
			Description:         description(o.Description),
			MarkdownDescription: description(o.Description),
			DeprecationMessage:  deprecationMessage(o.DeprecationMessage),
		},
		CustomType: o.CustomType,
		Validators: o.Validators,
	}, nil
}
