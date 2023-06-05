// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func TestGeneratorInt64Attribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input             GeneratorInt64Attribute
		expectedAttribute string
		expectedError     error
	}{
		"custom-type": {
			input: GeneratorInt64Attribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expectedAttribute: `
"int64_attribute": schema.Int64Attribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Required: true,
				},
			},
			expectedAttribute: `
"int64_attribute": schema.Int64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Optional: true,
				},
			},
			expectedAttribute: `
"int64_attribute": schema.Int64Attribute{
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Sensitive: true,
				},
			},
			expectedAttribute: `
"int64_attribute": schema.Int64Attribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Description: "description",
				},
			},
			expectedAttribute: `
"int64_attribute": schema.Int64Attribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					DeprecationMessage: "deprecated",
				},
			},
			expectedAttribute: `
"int64_attribute": schema.Int64Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorInt64Attribute{
				Validators: []specschema.Int64Validator{
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
"int64_attribute": schema.Int64Attribute{
Validators: []validator.Int64{
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

			got, err := testCase.input.ToString("int64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
