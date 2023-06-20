// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func convertNumberAttribute(a *datasource.NumberAttribute) (datasource_generate.GeneratorNumberAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorNumberAttribute{}, fmt.Errorf("*datasource.NumberAttribute is nil")
	}

	return datasource_generate.GeneratorNumberAttribute{
		NumberAttribute: schema.NumberAttribute{
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
