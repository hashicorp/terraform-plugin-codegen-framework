// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func TestConvertInt64Attribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.Int64Attribute
		expected      datasource_generate.GeneratorInt64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.Int64Attribute is nil"),
		},
		"computed": {
			input: &datasource.Int64Attribute{
				ComputedOptionalRequired: "computed",
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.Int64Attribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.Int64Attribute{
				ComputedOptionalRequired: "optional",
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.Int64Attribute{
				ComputedOptionalRequired: "required",
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.Int64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{},
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.Int64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.Int64Attribute{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.Int64Attribute{
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.Int64Attribute{
				Validators: specschema.Int64Validators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/.../myvalidator",
								},
							},
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorInt64Attribute{
				Validators: specschema.Int64Validators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/.../myvalidator",
								},
							},
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

			got, err := convertInt64Attribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
