package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertSetAttribute(a *provider.SetAttribute) (provider_generate.GeneratorSetAttribute, error) {
	if a == nil {
		return provider_generate.GeneratorSetAttribute{}, fmt.Errorf("*provider.SetAttribute is nil")
	}

	elemType, err := convertElementType(a.ElementType)
	if err != nil {
		return provider_generate.GeneratorSetAttribute{}, err
	}

	return provider_generate.GeneratorSetAttribute{
		SetAttribute: schema.SetAttribute{
			ElementType:         elemType,
			Required:            isRequired(a.OptionalRequired),
			Optional:            isOptional(a.OptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},
		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}
