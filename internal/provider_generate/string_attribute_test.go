// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func TestGeneratorStringAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorStringAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorStringAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorStringAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorStringAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorStringAttribute{
				Validators: []specschema.StringValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorStringAttribute{
				Validators: []specschema.StringValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorStringAttribute{
				Validators: []specschema.StringValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"validator-custom-import": {
			input: GeneratorStringAttribute{
				Validators: []specschema.StringValidator{
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
				schemaImport:    {},
				validatorImport: {},
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

func TestGeneratorStringAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input             GeneratorStringAttribute
		expectedAttribute string
		expectedError     error
	}{
		"custom-type": {
			input: GeneratorStringAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expectedAttribute: `
"string_attribute": schema.StringAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Required: true,
				},
			},
			expectedAttribute: `
"string_attribute": schema.StringAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Optional: true,
				},
			},
			expectedAttribute: `
"string_attribute": schema.StringAttribute{
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Sensitive: true,
				},
			},
			expectedAttribute: `
"string_attribute": schema.StringAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Description: "description",
				},
			},
			expectedAttribute: `
"string_attribute": schema.StringAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expectedAttribute: `
"string_attribute": schema.StringAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorStringAttribute{
				Validators: []specschema.StringValidator{
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
"string_attribute": schema.StringAttribute{
Validators: []validator.String{
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

			got, err := testCase.input.ToString("string_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
