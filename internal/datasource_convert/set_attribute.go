// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func convertSetAttribute(a *datasource.SetAttribute) (datasource_generate.GeneratorSetAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorSetAttribute{}, fmt.Errorf("*datasource.SetAttribute is nil")
	}

	return datasource_generate.GeneratorSetAttribute{
		SetAttribute: schema.SetAttribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
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
