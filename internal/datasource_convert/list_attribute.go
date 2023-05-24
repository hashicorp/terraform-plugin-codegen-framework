package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func convertListAttribute(a *datasource.ListAttribute) (datasource_generate.GeneratorListAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorListAttribute{}, fmt.Errorf("*datasource.ListAttribute is nil")
	}

	elemType, err := convertElementType(a.ElementType)
	if err != nil {
		return datasource_generate.GeneratorListAttribute{}, err
	}

	return datasource_generate.GeneratorListAttribute{
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

		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}
