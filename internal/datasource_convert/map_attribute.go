// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func convertMapAttribute(a *datasource.MapAttribute) (datasource_generate.GeneratorMapAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorMapAttribute{}, fmt.Errorf("*datasource.MapAttribute is nil")
	}

	elemType, err := convertElementType(a.ElementType)
	if err != nil {
		return datasource_generate.GeneratorMapAttribute{}, err
	}

	return datasource_generate.GeneratorMapAttribute{
		MapAttribute: schema.MapAttribute{
			ElementType:         elemType,
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
