// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func TestConvertNumberAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.NumberAttribute
		expected      datasource_generate.GeneratorNumberAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.NumberAttribute is nil"),
		},
		"computed": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.NumberAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.NumberAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.NumberAttribute{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.NumberAttribute{
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.NumberAttribute{
				Validators: []specschema.NumberValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorNumberAttribute{
				Validators: []specschema.NumberValidator{
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

			got, err := convertNumberAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
