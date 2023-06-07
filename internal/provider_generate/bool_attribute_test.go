// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func TestGeneratorBoolAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorBoolAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
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
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
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
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
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

func TestGeneratorBoolAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		boolAttribute     GeneratorBoolAttribute
		expectedAttribute string
		expectedError     error
	}{
		"custom-type": {
			boolAttribute: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expectedAttribute: `
"bool_attribute": schema.BoolAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			boolAttribute: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Required: true,
				},
			},
			expectedAttribute: `
"bool_attribute": schema.BoolAttribute{
Required: true,
},`,
		},

		"optional": {
			boolAttribute: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Optional: true,
				},
			},
			expectedAttribute: `
"bool_attribute": schema.BoolAttribute{
Optional: true,
},`,
		},

		"sensitive": {
			boolAttribute: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Sensitive: true,
				},
			},
			expectedAttribute: `
"bool_attribute": schema.BoolAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			boolAttribute: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Description: "description",
				},
			},
			expectedAttribute: `
"bool_attribute": schema.BoolAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			boolAttribute: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expectedAttribute: `
"bool_attribute": schema.BoolAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			boolAttribute: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
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
"bool_attribute": schema.BoolAttribute{
Validators: []validator.Bool{
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

			got, err := testCase.boolAttribute.ToString("bool_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
