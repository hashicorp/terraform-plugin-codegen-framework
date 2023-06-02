// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func convertFloat64Attribute(a *resource.Float64Attribute) (resource_generate.GeneratorFloat64Attribute, error) {
	if a == nil {
		return resource_generate.GeneratorFloat64Attribute{}, fmt.Errorf("*resource.Float64Attribute is nil")
	}

	return resource_generate.GeneratorFloat64Attribute{
		Float64Attribute: schema.Float64Attribute{
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
		PlanModifiers: a.PlanModifiers,
		Validators:    a.Validators,
	}, nil
}
