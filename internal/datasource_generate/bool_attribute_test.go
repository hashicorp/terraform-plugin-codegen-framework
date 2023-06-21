// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func pointer[T any](in T) *T {
	return &in
}

func TestGeneratorBoolAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorBoolAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				"github.com/hashicorp/terraform-plugin-framework/types": {},
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
				"github.com/hashicorp/terraform-plugin-framework/types": {},
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
				"github.com/hashicorp/terraform-plugin-framework/types": {},
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
				"github.com/hashicorp/terraform-plugin-framework/types": {},
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
				generatorschema.ValidatorImport:                         {},
				"github.com/hashicorp/terraform-plugin-framework/types": {},
				"github.com/myotherproject/myvalidators/validator":      {},
				"github.com/myproject/myvalidators/validator":           {},
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
		input         GeneratorBoolAttribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Required: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Optional: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Computed: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Sensitive: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Description: "description",
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators-empty": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
},`,
		},
		"validators": {
			input: GeneratorBoolAttribute{
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
			expected: `
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

			got, err := testCase.input.ToString("bool_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorBoolAttribute_ToModel(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorBoolAttribute
		expected      string
		expectedError error
	}{
		"default": {
			expected: "\nBoolAttribute types.Bool `tfsdk:\"bool_attribute\"`",
		},
		"custom-type-nil": {
			input: GeneratorBoolAttribute{
				CustomType: nil,
			},
			expected: "\nBoolAttribute types.Bool `tfsdk:\"bool_attribute\"`",
		},
		"custom-type-missing-value-type": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: "\nBoolAttribute types.Bool `tfsdk:\"bool_attribute\"`",
		},
		"custom-type-value-type-empty-string": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: "\nBoolAttribute types.Bool `tfsdk:\"bool_attribute\"`",
		},
		"custom-type": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: "\nBoolAttribute my_custom_value_type `tfsdk:\"bool_attribute\"`",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToModel("bool_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
