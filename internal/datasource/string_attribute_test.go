// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

func TestGeneratorStringAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.StringAttribute
		expected      GeneratorStringAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.StringAttribute is nil"),
		},
		"computed": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &datasource.StringAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &datasource.StringAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeString, nil),
			},
		},
		"deprecation_message": {
			input: &datasource.StringAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorStringAttribute{
				CustomType:         convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				Validators:         convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &datasource.StringAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorStringAttribute{
				CustomType:  convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &datasource.StringAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &datasource.StringAttribute{
				Validators: specschema.StringValidators{
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
			expected: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{
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

			got, err := NewGeneratorStringAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorStringAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorStringAttribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					"string_attribute",
				),
			},
			expected: `"string_attribute": schema.StringAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ExtString",
					},
					"string_attribute",
				),
			},
			expected: `"string_attribute": schema.StringAttribute{
CustomType: StringAttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ExtString",
					},
					"string_attribute",
				),
			},
			expected: `"string_attribute": schema.StringAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
			},
			expected: `"string_attribute": schema.StringAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
			},
			expected: `"string_attribute": schema.StringAttribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
			},
			expected: `"string_attribute": schema.StringAttribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorStringAttribute{
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"string_attribute": schema.StringAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorStringAttribute{
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"string_attribute": schema.StringAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorStringAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"string_attribute": schema.StringAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators-empty": {
			input: GeneratorStringAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{})},
			expected: `"string_attribute": schema.StringAttribute{
},`,
		},
		"validators": {
			input: GeneratorStringAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeString, specschema.CustomValidators{
					{
						SchemaDefinition: "my_validator.Validate()",
					},
					{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"string_attribute": schema.StringAttribute{
Validators: []validator.String{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("string_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorStringAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorStringAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "types.String",
				TfsdkName: "string_attribute",
			},
		},
		"custom-type": {
			input: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					"",
				),
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "string_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
					"string_attribute",
				),
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "StringAttributeValue",
				TfsdkName: "string_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorStringAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
					"",
				),
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "string_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("string_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
