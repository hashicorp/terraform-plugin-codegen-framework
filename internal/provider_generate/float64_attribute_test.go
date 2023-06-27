// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorFloat64Attribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorFloat64Attribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorFloat64Attribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorFloat64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorFloat64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorFloat64Attribute{
				Validators: []specschema.Float64Validator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorFloat64Attribute{
				Validators: []specschema.Float64Validator{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{},
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorFloat64Attribute{
				Validators: []specschema.Float64Validator{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "",
								},
							},
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import": {
			input: GeneratorFloat64Attribute{
				Validators: []specschema.Float64Validator{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myotherproject/myvalidators/validator",
								},
							},
						},
					},
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myproject/myvalidators/validator",
								},
							},
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.ValidatorImport:                    {},
				generatorschema.TypesImport:                        {},
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

func TestGeneratorFloat64Attribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input             GeneratorFloat64Attribute
		expectedAttribute string
		expectedError     error
	}{
		"custom-type": {
			input: GeneratorFloat64Attribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expectedAttribute: `
"float64_attribute": schema.Float64Attribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Required: true,
				},
			},
			expectedAttribute: `
"float64_attribute": schema.Float64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Optional: true,
				},
			},
			expectedAttribute: `
"float64_attribute": schema.Float64Attribute{
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Sensitive: true,
				},
			},
			expectedAttribute: `
"float64_attribute": schema.Float64Attribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Description: "description",
				},
			},
			expectedAttribute: `
"float64_attribute": schema.Float64Attribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					DeprecationMessage: "deprecated",
				},
			},
			expectedAttribute: `
"float64_attribute": schema.Float64Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorFloat64Attribute{
				Validators: []specschema.Float64Validator{
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
"float64_attribute": schema.Float64Attribute{
Validators: []validator.Float64{
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

			got, err := testCase.input.ToString("float64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorFloat64Attribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorFloat64Attribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "Float64Attribute",
				ValueType: "types.Float64",
				TfsdkName: "float64_attribute",
			},
		},
		"custom-type": {
			input: GeneratorFloat64Attribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "Float64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "float64_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("float64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
