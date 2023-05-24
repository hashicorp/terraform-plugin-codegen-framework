package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func convertStringAttribute(a *datasource.StringAttribute) (datasource_generate.GeneratorStringAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorStringAttribute{}, fmt.Errorf("*datasource.StringAttribute is nil")
	}

	return datasource_generate.GeneratorStringAttribute{
		StringAttribute: schema.StringAttribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}
