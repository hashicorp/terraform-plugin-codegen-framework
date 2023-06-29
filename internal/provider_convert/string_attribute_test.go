// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
)

func TestConvertStringAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.StringAttribute
		expected      provider_generate.GeneratorStringAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.StringAttribute is nil"),
		},
		"optional": {
			input: &provider.StringAttribute{
				OptionalRequired: "optional",
			},
			expected: provider_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &provider.StringAttribute{
				OptionalRequired: "required",
			},
			expected: provider_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &provider.StringAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{},
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
			input: &provider.StringAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.StringAttribute{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &provider.StringAttribute{
				Sensitive: pointer(true),
			},
			expected: provider_generate.GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &provider.StringAttribute{
				Validators: []specschema.StringValidator{
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
			expected: provider_generate.GeneratorStringAttribute{
				Validators: []specschema.StringValidator{
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
