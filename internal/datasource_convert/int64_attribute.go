// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func convertInt64Attribute(a *datasource.Int64Attribute) (datasource_generate.GeneratorInt64Attribute, error) {
	if a == nil {
		return datasource_generate.GeneratorInt64Attribute{}, fmt.Errorf("*datasource.Int64Attribute is nil")
	}

	return datasource_generate.GeneratorInt64Attribute{
		Int64Attribute: schema.Int64Attribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}
