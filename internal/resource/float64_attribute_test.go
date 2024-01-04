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

func TestGeneratorFloat64Attribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.Float64Attribute
		expected      GeneratorFloat64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.Float64Attribute is nil"),
		},
		"computed": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorFloat64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorFloat64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorFloat64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorFloat64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.Float64Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorFloat64Attribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, nil),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, nil),
			},
		},
		"deprecation_message": {
			input: &resource.Float64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorFloat64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage:  convert.NewDeprecationMessage(pointer("deprecation message")),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.Float64Attribute{
				Description: pointer("description"),
			},
			expected: GeneratorFloat64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description:         convert.NewDescription(pointer("description")),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.Float64Attribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorFloat64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:           convert.NewSensitive(pointer(true)),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.Float64Attribute{
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, nil),
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
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{
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
			input: &resource.Float64Attribute{
				PlanModifiers: specschema.Float64PlanModifiers{
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
			expected: GeneratorFloat64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: specschema.Float64PlanModifiers{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.Float64Attribute{
				Default: &specschema.Float64Default{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(1.234),
				},
			},
			expected: GeneratorFloat64Attribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Default: &specschema.Float64Default{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(1.234),
				},
				DefaultFloat64: convert.NewDefaultFloat64(&specschema.Float64Default{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(1.234),
				}),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, specschema.CustomValidators{}),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				CustomTypePrimitive: convert.NewCustomTypePrimitive(
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
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorFloat64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorFloat64Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Computed: true,
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

		"validators": {
			input: GeneratorFloat64Attribute{
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeFloat64, []*specschema.CustomValidator{
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

		"plan-modifiers": {
			input: GeneratorFloat64Attribute{
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeFloat64, []*specschema.CustomPlanModifier{
					{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
PlanModifiers: []planmodifier.Float64{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorFloat64Attribute{
				DefaultFloat64: convert.NewDefaultFloat64(&specschema.Float64Default{
					Static: pointer(1.234),
				}),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Default: float64default.StaticFloat64(1.234),
},`,
		},

		"default-custom": {
			input: GeneratorFloat64Attribute{
				DefaultFloat64: convert.NewDefaultFloat64(&specschema.Float64Default{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_float64_default.Default()",
					},
				}),
			},
			expected: `"float64_attribute": schema.Float64Attribute{
Default: my_float64_default.Default(),
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
