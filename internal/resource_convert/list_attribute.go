package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func convertListAttribute(a *resource.ListAttribute) (resource_generate.GeneratorListAttribute, error) {
	if a == nil {
		return resource_generate.GeneratorListAttribute{}, fmt.Errorf("*resource.ListAttribute is nil")
	}

	elemType, err := convertElementType(a.ElementType)
	if err != nil {
		return resource_generate.GeneratorListAttribute{}, err
	}

	return resource_generate.GeneratorListAttribute{
		ListAttribute: schema.ListAttribute{
			ElementType:         elemType,
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
