package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertBoolAttribute(a *provider.BoolAttribute) (provider_generate.GeneratorBoolAttribute, error) {
	if a == nil {
		return provider_generate.GeneratorBoolAttribute{}, fmt.Errorf("*provider.BoolAttribute is nil")
	}

	return provider_generate.GeneratorBoolAttribute{
		BoolAttribute: schema.BoolAttribute{
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
