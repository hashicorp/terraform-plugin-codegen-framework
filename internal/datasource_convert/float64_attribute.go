package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func convertFloat64Attribute(a *datasource.Float64Attribute) (datasource_generate.GeneratorFloat64Attribute, error) {
	if a == nil {
		return datasource_generate.GeneratorFloat64Attribute{}, fmt.Errorf("*datasource.Float64Attribute is nil")
	}

	return datasource_generate.GeneratorFloat64Attribute{
		Float64Attribute: schema.Float64Attribute{
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
