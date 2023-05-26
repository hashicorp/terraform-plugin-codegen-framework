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

func TestConvertBoolAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.BoolAttribute
		expected      datasource_generate.GeneratorBoolAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.BoolAttribute is nil"),
		},
		"computed": {
			input: &datasource.BoolAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.BoolAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.BoolAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.BoolAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.BoolAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.BoolAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.BoolAttribute{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.BoolAttribute{
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.BoolAttribute{
				Validators: []specschema.BoolValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
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

			got, err := convertBoolAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
