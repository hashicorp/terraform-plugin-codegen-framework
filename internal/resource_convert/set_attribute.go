// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
)

func convertSetAttribute(a *resource.SetAttribute) (resource_generate.GeneratorSetAttribute, error) {
	if a == nil {
		return resource_generate.GeneratorSetAttribute{}, fmt.Errorf("*resource.SetAttribute is nil")
	}

	return resource_generate.GeneratorSetAttribute{
		SetAttribute: schema.SetAttribute{
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
