// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

func TestGeneratorNumberAttribute_Schema(t *testing.T) {
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
				Validators: specschema.NumberValidators{
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

			got, err := testCase.input.Schema("number_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorNumberAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorNumberAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "types.Number",
				TfsdkName: "number_attribute",
			},
		},
		"custom-type": {
			input: GeneratorNumberAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "number_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("number_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
