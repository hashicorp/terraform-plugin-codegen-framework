package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func convertSetAttribute(a *datasource.SetAttribute) (datasource_generate.GeneratorSetAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorSetAttribute{}, fmt.Errorf("*datasource.SetAttribute is nil")
	}

	elemType, err := convertElementType(a.ElementType)
	if err != nil {
		return datasource_generate.GeneratorSetAttribute{}, err
	}

	return datasource_generate.GeneratorSetAttribute{
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

		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}
