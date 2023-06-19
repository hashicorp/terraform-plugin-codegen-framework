// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorNumberAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorNumberAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{},
		},
		"custom-type-without-import": {
			input: GeneratorNumberAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorNumberAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorNumberAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorNumberAttribute{
				Validators: []specschema.NumberValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-nil": {
			input: GeneratorNumberAttribute{
				Validators: []specschema.NumberValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorNumberAttribute{
				Validators: []specschema.NumberValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import": {
			input: GeneratorNumberAttribute{
				Validators: []specschema.NumberValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer("github.com/myotherproject/myvalidators/validator"),
						},
					},
					{
						Custom: &specschema.CustomValidator{
							Import: pointer("github.com/myproject/myvalidators/validator"),
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.ValidatorImport:                    {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorNumberAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input             GeneratorNumberAttribute
		expectedAttribute string
		expectedError     error
	}{
		"custom-type": {
			input: GeneratorNumberAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expectedAttribute: `
"number_attribute": schema.NumberAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Required: true,
				},
			},
			expectedAttribute: `
"number_attribute": schema.NumberAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Optional: true,
				},
			},
			expectedAttribute: `
"number_attribute": schema.NumberAttribute{
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Sensitive: true,
				},
			},
			expectedAttribute: `
"number_attribute": schema.NumberAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Description: "description",
				},
			},
			expectedAttribute: `
"number_attribute": schema.NumberAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expectedAttribute: `
"number_attribute": schema.NumberAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorNumberAttribute{
				Validators: []specschema.NumberValidator{
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
					},
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					},
				},
			},
			expectedAttribute: `
"number_attribute": schema.NumberAttribute{
Validators: []validator.Number{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("number_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
