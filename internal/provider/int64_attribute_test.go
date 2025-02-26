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

func TestGeneratorInt64Attribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.Int64Attribute
		expected      GeneratorInt64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.Int64Attribute is nil"),
		},
		"optional": {
			input: &provider.Int64Attribute{
				OptionalRequired: "optional",
			},
			expected: GeneratorInt64Attribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
				CustomType:       convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:       convert.NewValidators(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &provider.Int64Attribute{
				OptionalRequired: "required",
			},
			expected: GeneratorInt64Attribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
				CustomType:       convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:       convert.NewValidators(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &provider.Int64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeInt64, nil),
			},
		},
		"deprecation_message": {
			input: &provider.Int64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorInt64Attribute{
				CustomType:         convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				Validators:         convert.NewValidators(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &provider.Int64Attribute{
				Description: pointer("description"),
			},
			expected: GeneratorInt64Attribute{
				CustomType:  convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &provider.Int64Attribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &provider.Int64Attribute{
				Validators: specschema.Int64Validators{
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
			expected: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeInt64, specschema.CustomValidators{
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

			got, err := NewGeneratorInt64Attribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorInt64Attribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorInt64Attribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					"int64_attribute",
				),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ExtInt64",
					},
					"int64_attribute",
				),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
CustomType: Int64AttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ExtInt64",
					},
					"int64_attribute",
				),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorInt64Attribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorInt64Attribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorInt64Attribute{
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorInt64Attribute{
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorInt64Attribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators-empty": {
			input: GeneratorInt64Attribute{
				Validators: convert.NewValidators(convert.ValidatorTypeInt64, nil),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
},`,
		},
		"validators": {
			input: GeneratorInt64Attribute{
				Validators: convert.NewValidators(convert.ValidatorTypeInt64, specschema.CustomValidators{
					{
						SchemaDefinition: "my_validator.Validate()",
					},
					{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Validators: []validator.Int64{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("int64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorInt64Attribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorInt64Attribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "types.Int64",
				TfsdkName: "int64_attribute",
			},
		},
		"custom-type": {
			input: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					"",
				),
			},
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "int64_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.Int64Attribute",
					},
					"int64_attribute",
				),
			},
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "Int64AttributeValue",
				TfsdkName: "int64_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorInt64Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.Int64Attribute",
					},
					"",
				),
			},
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "int64_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("int64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return x.Error() == y.Error()
})
