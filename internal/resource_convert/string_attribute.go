// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func convertStringAttribute(a *resource.StringAttribute) (resource_generate.GeneratorStringAttribute, error) {
	if a == nil {
		return resource_generate.GeneratorStringAttribute{}, fmt.Errorf("*resource.StringAttribute is nil")
	}

	return resource_generate.GeneratorStringAttribute{
		StringAttribute: schema.StringAttribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		AssociatedExternalType: generatorschema.NewAssocExtType(a.AssociatedExternalType),
		CustomType:             a.CustomType,
		Default:                a.Default,
		PlanModifiers:          a.PlanModifiers,
		Validators:             a.Validators,
	}, nil
}
