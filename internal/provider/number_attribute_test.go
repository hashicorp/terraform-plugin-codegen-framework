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

func TestGeneratorNumberAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.NumberAttribute
		expected      GeneratorNumberAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.NumberAttribute is nil"),
		},
		"optional": {
			input: &provider.NumberAttribute{
				OptionalRequired: "optional",
			},
			expected: GeneratorNumberAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
				CustomType:       convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:       convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &provider.NumberAttribute{
				OptionalRequired: "required",
			},
			expected: GeneratorNumberAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
				CustomType:       convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:       convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &provider.NumberAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, nil),
			},
		},
		"deprecation_message": {
			input: &provider.NumberAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorNumberAttribute{
				CustomType:         convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				Validators:         convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &provider.NumberAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorNumberAttribute{
				CustomType:  convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &provider.NumberAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &provider.NumberAttribute{
				Validators: specschema.NumberValidators{
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
			expected: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{
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

			got, err := NewGeneratorNumberAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorNumberAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorNumberAttribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					"number_attribute",
				),
			},
			expected: `"number_attribute": schema.NumberAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ExtNumber",
					},
					"number_attribute",
				),
			},
			expected: `"number_attribute": schema.NumberAttribute{
CustomType: NumberAttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ExtNumber",
					},
					"number_attribute",
				),
			},
			expected: `"number_attribute": schema.NumberAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorNumberAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorNumberAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorNumberAttribute{
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorNumberAttribute{
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorNumberAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"number_attribute": schema.NumberAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators-empty": {
			input: GeneratorNumberAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
			expected: `"number_attribute": schema.NumberAttribute{
},`,
		},
		"validators": {
			input: GeneratorNumberAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{
					{
						SchemaDefinition: "my_validator.Validate()",
					},
					{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Validators: []validator.Number{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("number_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
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
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					"",
				),
			},
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "number_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.NumberAttribute",
					},
					"number_attribute",
				),
			},
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "NumberAttributeValue",
				TfsdkName: "number_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.NumberAttribute",
					},
					"",
				),
			},
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "number_attribute",
			},
		},
	}

	for name, testCase := range testCases {

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
