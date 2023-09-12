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

func TestConvertFloat64Attribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.Float64Attribute
		expected      datasource_generate.GeneratorFloat64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.Float64Attribute is nil"),
		},
		"computed": {
			input: &datasource.Float64Attribute{
				ComputedOptionalRequired: "computed",
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.Float64Attribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.Float64Attribute{
				ComputedOptionalRequired: "optional",
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.Float64Attribute{
				ComputedOptionalRequired: "required",
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.Float64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{},
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
			input: &datasource.Float64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.Float64Attribute{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.Float64Attribute{
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.Float64Attribute{
				Validators: specschema.Float64Validators{
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
			expected: datasource_generate.GeneratorFloat64Attribute{
				Validators: specschema.Float64Validators{
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

			got, err := convertFloat64Attribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
