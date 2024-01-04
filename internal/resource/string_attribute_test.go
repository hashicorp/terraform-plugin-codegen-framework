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

func TestGeneratorStringAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.StringAttribute
		expected      GeneratorStringAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.StringAttribute is nil"),
		},
		"computed": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorStringAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.StringAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, nil),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, nil),
			},
		},
		"deprecation_message": {
			input: &resource.StringAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage:  convert.NewDeprecationMessage(pointer("deprecation message")),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.StringAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description:         convert.NewDescription(pointer("description")),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.StringAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:           convert.NewSensitive(pointer(true)),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.StringAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, nil),
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
		"plan-modifiers": {
			input: &resource.StringAttribute{
				PlanModifiers: specschema.StringPlanModifiers{
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
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: specschema.StringPlanModifiers{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.StringAttribute{
				Default: &specschema.StringDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer("str"),
				},
			},
			expected: GeneratorStringAttribute{
				CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Default: &specschema.StringDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer("str"),
				},
				DefaultString: convert.NewDefaultString(&specschema.StringDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer("str"),
				}),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeString, specschema.CustomValidators{}),
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

		"plan-modifiers": {
			input: GeneratorStringAttribute{
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeString, []*specschema.CustomPlanModifier{
					{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"string_attribute": schema.StringAttribute{
PlanModifiers: []planmodifier.String{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorStringAttribute{
				DefaultString: convert.NewDefaultString(&specschema.StringDefault{
					Static: pointer("str"),
				}),
			},
			expected: `"string_attribute": schema.StringAttribute{
Default: stringdefault.StaticString("str"),
},`,
		},

		"default-custom": {
			input: GeneratorStringAttribute{
				DefaultString: convert.NewDefaultString(&specschema.StringDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_string_default.Default()",
					},
				}),
			},
			expected: `"string_attribute": schema.StringAttribute{
Default: my_string_default.Default(),
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
