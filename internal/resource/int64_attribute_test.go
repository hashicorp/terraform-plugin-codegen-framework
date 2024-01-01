// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorInt64Attribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.Int64Attribute
		expected      GeneratorInt64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.Int64Attribute is nil"),
		},
		"computed": {
			input: &resource.Int64Attribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorInt64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.Int64Attribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorInt64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.Int64Attribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorInt64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.Int64Attribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorInt64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.Int64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorInt64Attribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeInt64, nil),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeInt64, nil),
			},
		},
		"deprecation_message": {
			input: &resource.Int64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorInt64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage:  convert.NewDeprecationMessage(pointer("deprecation message")),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.Int64Attribute{
				Description: pointer("description"),
			},
			expected: GeneratorInt64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description:         convert.NewDescription(pointer("description")),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.Int64Attribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorInt64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:           convert.NewSensitive(pointer(true)),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.Int64Attribute{
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeInt64, nil),
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
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{
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
		"plan-modifiers": {
			input: &resource.Int64Attribute{
				PlanModifiers: specschema.Int64PlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: GeneratorInt64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: specschema.Int64PlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeInt64, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.Int64Attribute{
				Default: &specschema.Int64Default{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(int64(1234)),
				},
			},
			expected: GeneratorInt64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Default: &specschema.Int64Default{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(int64(1234)),
				},
				DefaultInt64: convert.NewDefaultInt64(&specschema.Int64Default{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(int64(1234)),
				}),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeInt64, specschema.CustomValidators{}),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorInt64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorInt64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Computed: true,
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

		"validators": {
			input: GeneratorInt64Attribute{
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeInt64, []*specschema.CustomValidator{
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

		"plan-modifiers": {
			input: GeneratorInt64Attribute{
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeInt64, []*specschema.CustomPlanModifier{
					{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
PlanModifiers: []planmodifier.Int64{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorInt64Attribute{
				DefaultInt64: convert.NewDefaultInt64(&specschema.Int64Default{
					Static: pointer(int64(1234)),
				}),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Default: int64default.StaticInt64(1234),
},`,
		},

		"default-custom": {
			input: GeneratorInt64Attribute{
				DefaultInt64: convert.NewDefaultInt64(&specschema.Int64Default{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_int64_default.Default()",
					},
				}),
			},
			expected: `"int64_attribute": schema.Int64Attribute{
Default: my_int64_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "int64_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorInt64Attribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.Int64Attribute",
					},
				},
			},
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "Int64AttributeValue",
				TfsdkName: "int64_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorInt64Attribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.Int64Attribute",
					},
				},
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "int64_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
