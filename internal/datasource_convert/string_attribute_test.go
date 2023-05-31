package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func TestConvertStringAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.StringAttribute
		expected      datasource_generate.GeneratorStringAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.StringAttribute is nil"),
		},
		"computed": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.StringAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.StringAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.StringAttribute{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.StringAttribute{
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.StringAttribute{
				Validators: []specschema.StringValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorStringAttribute{
				Validators: []specschema.StringValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertStringAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
