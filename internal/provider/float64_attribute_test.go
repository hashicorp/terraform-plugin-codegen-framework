// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
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
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
				CustomType:       convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:       convert.NewValidators(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &provider.Float64Attribute{
				OptionalRequired: "required",
			},
			expected: GeneratorFloat64Attribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
				CustomType:       convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:       convert.NewValidators(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
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
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeFloat64, nil),
			},
		},
		"deprecation_message": {
			input: &provider.Float64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorFloat64Attribute{
				CustomType:         convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				Validators:         convert.NewValidators(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &provider.Float64Attribute{
				Description: pointer("description"),
			},
			expected: GeneratorFloat64Attribute{
				CustomType:  convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &provider.Float64Attribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorFloat64Attribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
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
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeFloat64, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/.../myvalidator",
							},
						},
						SchemaDefinition: "myvalidator.Validate()",
					},
				}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorFloat64Attribute("name", testCase.input)

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
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					"float64_attribute",
				),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorFloat64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ExtFloat64",
					},
					"float64_attribute",
				),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
CustomType: Float64AttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorFloat64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ExtFloat64",
					},
					"float64_attribute",
				),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorFloat64Attribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorFloat64Attribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorFloat64Attribute{
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorFloat64Attribute{
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorFloat64Attribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators-empty": {
			input: GeneratorFloat64Attribute{
				Validators: convert.NewValidators(convert.ValidatorTypeFloat64, nil)},
			expected: `"float64_attribute": schema.Float64Attribute{
},`,
		},
		"validators": {
			input: GeneratorFloat64Attribute{
				Validators: convert.NewValidators(convert.ValidatorTypeFloat64, specschema.CustomValidators{
					{
						SchemaDefinition: "my_validator.Validate()",
					},
					{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Validators: []validator.Float64{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {

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
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					"",
				),
			},
			expected: model.Field{
				Name:      "Float64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "float64_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorFloat64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.Float64Attribute",
					},
					"float64_attribute",
				),
			},
			expected: model.Field{
				Name:      "Float64Attribute",
				ValueType: "Float64AttributeValue",
				TfsdkName: "float64_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorFloat64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.Float64Attribute",
					},
					"",
				),
			},
			expected: model.Field{
				Name:      "Float64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "float64_attribute",
			},
		},
	}

	for name, testCase := range testCases {

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
