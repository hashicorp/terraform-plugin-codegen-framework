package resource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func convertSetAttribute(a *resource.SetAttribute) (resource_generate.GeneratorSetAttribute, error) {
	if a == nil {
		return resource_generate.GeneratorSetAttribute{}, fmt.Errorf("*resource.SetAttribute is nil")
	}

	elemType, err := convertElementType(a.ElementType)
	if err != nil {
		return resource_generate.GeneratorSetAttribute{}, err
	}

	return resource_generate.GeneratorSetAttribute{
		SetAttribute: schema.SetAttribute{
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
