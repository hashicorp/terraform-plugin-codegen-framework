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
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorStringAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.StringAttribute
		expected      GeneratorStringAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.StringAttribute is nil"),
		},
		"optional": {
			input: &provider.StringAttribute{
				OptionalRequired: "optional",
			},
			expected: GeneratorStringAttribute{
				OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &provider.StringAttribute{
				OptionalRequired: "required",
			},
			expected: GeneratorStringAttribute{
				OptionalRequired:    convert.NewOptionalRequired(specschema.Required),
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &provider.StringAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorStringAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				CustomTypePrimitive: convert.NewCustomTypePrimitive(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, nil, "name"),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeString, nil),
			},
		},
		"deprecation_message": {
			input: &provider.StringAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage:  convert.NewDeprecationMessage(pointer("deprecation message")),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &provider.StringAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description:         convert.NewDescription(pointer("description")),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &provider.StringAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:           convert.NewSensitive(pointer(true)),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &provider.StringAttribute{
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
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
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{
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
		name, testCase := name, testCase

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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
			},
			expected: `"string_attribute": schema.StringAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorStringAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
			},
			expected: `"string_attribute": schema.StringAttribute{
Optional: true,
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
				Validators: specschema.StringValidators{},
			},
			expected: `"string_attribute": schema.StringAttribute{
},`,
		},
		"validators": {
			input: GeneratorStringAttribute{
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeString, []*specschema.CustomValidator{
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
		name, testCase := name, testCase

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
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "string_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorStringAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
				},
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "StringAttributeValue",
				TfsdkName: "string_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorStringAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
				},
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "string_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
