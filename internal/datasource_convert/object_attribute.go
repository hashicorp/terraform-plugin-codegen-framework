// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func convertObjectAttribute(o *datasource.ObjectAttribute) (datasource_generate.GeneratorObjectAttribute, error) {
	if o == nil {
		return datasource_generate.GeneratorObjectAttribute{}, fmt.Errorf("*datasource.ObjectAttribute is nil")
	}

	attrTypes, err := convertAttrTypes(o.AttributeTypes)
	if err != nil {
		return datasource_generate.GeneratorObjectAttribute{}, err
	}

	return datasource_generate.GeneratorObjectAttribute{
		ObjectAttribute: schema.ObjectAttribute{
			AttributeTypes:      attrTypes,
			Required:            isRequired(o.ComputedOptionalRequired),
			Optional:            isOptional(o.ComputedOptionalRequired),
			Computed:            isComputed(o.ComputedOptionalRequired),
			Sensitive:           isSensitive(o.Sensitive),
			Description:         description(o.Description),
			MarkdownDescription: description(o.Description),
			DeprecationMessage:  deprecationMessage(o.DeprecationMessage),
		},

		CustomType: o.CustomType,
		Validators: o.Validators,
	}, nil
}
