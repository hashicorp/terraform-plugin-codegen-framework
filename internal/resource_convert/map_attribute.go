// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func convertMapAttribute(a *resource.MapAttribute) (resource_generate.GeneratorMapAttribute, error) {
	if a == nil {
		return resource_generate.GeneratorMapAttribute{}, fmt.Errorf("*resource.MapAttribute is nil")
	}

	return resource_generate.GeneratorMapAttribute{
		MapAttribute: schema.MapAttribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},
		CustomType:    a.CustomType,
		Default:       a.Default,
		ElementType:   a.ElementType,
		PlanModifiers: a.PlanModifiers,
		Validators:    a.Validators,
	}, nil
}
