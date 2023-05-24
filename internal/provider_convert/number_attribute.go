package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertNumberAttribute(a *provider.NumberAttribute) (provider_generate.GeneratorNumberAttribute, error) {
	if a == nil {
		return provider_generate.GeneratorNumberAttribute{}, fmt.Errorf("*provider.NumberAttribute is nil")
	}

	return provider_generate.GeneratorNumberAttribute{
		NumberAttribute: schema.NumberAttribute{
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
