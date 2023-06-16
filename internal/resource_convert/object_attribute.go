// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
)

func convertObjectAttribute(o *resource.ObjectAttribute) (resource_generate.GeneratorObjectAttribute, error) {
	if o == nil {
		return resource_generate.GeneratorObjectAttribute{}, fmt.Errorf("*resource.ObjectAttribute is nil")
	}

	attrTypes, err := convertAttrTypes(o.AttributeTypes)
	if err != nil {
		return resource_generate.GeneratorObjectAttribute{}, err
	}

	return resource_generate.GeneratorObjectAttribute{
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
		CustomType:    o.CustomType,
		Default:       o.Default,
		PlanModifiers: o.PlanModifiers,
		Validators:    o.Validators,
	}, nil
}
