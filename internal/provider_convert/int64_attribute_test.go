// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
)

func TestConvertInt64Attribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.Int64Attribute
		expected      provider_generate.GeneratorInt64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.Int64Attribute is nil"),
		},
		"optional": {
			input: &provider.Int64Attribute{
				OptionalRequired: "optional",
			},
			expected: provider_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &provider.Int64Attribute{
				OptionalRequired: "required",
			},
			expected: provider_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &provider.Int64Attribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &provider.Int64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.Int64Attribute{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &provider.Int64Attribute{
				Sensitive: pointer(true),
			},
			expected: provider_generate.GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &provider.Int64Attribute{
				Validators: []specschema.Int64Validator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: provider_generate.GeneratorInt64Attribute{
				Validators: []specschema.Int64Validator{
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
