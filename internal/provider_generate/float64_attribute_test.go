// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorFloat64Attribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.Float64Attribute
		expected      GeneratorFloat64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.Float64Attribute is nil"),
		},
		"optional": {
			input: &provider.Float64Attribute{
				OptionalRequired: "optional",
			},
			expected: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &provider.Float64Attribute{
				OptionalRequired: "required",
			},
			expected: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &provider.Float64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorFloat64Attribute{
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
			input: &provider.Float64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.Float64Attribute{
				Description: pointer("description"),
			},
			expected: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &provider.Float64Attribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &provider.Float64Attribute{
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
			expected: GeneratorFloat64Attribute{
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

			got, err := NewGeneratorFloat64Attribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorFloat64Attribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorFloat64Attribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorFloat64Attribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorFloat64Attribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.Float64Attribute",
					},
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
CustomType: Float64AttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorFloat64Attribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.Float64Attribute",
					},
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
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
			expected: `
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
			expected: `
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
			expected: `
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
			expected: `
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
			expected: `
"float64_attribute": schema.Float64Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorFloat64Attribute{
				Validators: specschema.Float64Validators{
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

			got, err := testCase.input.Schema("float64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
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
		"associated-external-type": {
			input: GeneratorFloat64Attribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.Float64Attribute",
					},
				},
			},
			expected: model.Field{
				Name:      "Float64Attribute",
				ValueType: "Float64AttributeValue",
				TfsdkName: "float64_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorFloat64Attribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.Float64Attribute",
					},
				},
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
