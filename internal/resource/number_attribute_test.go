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

func TestGeneratorNumberAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.NumberAttribute
		expected      GeneratorNumberAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.NumberAttribute is nil"),
		},
		"computed": {
			input: &resource.NumberAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorNumberAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.NumberAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorNumberAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.NumberAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorNumberAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.NumberAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorNumberAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.NumberAttribute{
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, nil),
				Validators:    convert.NewValidators(convert.ValidatorTypeNumber, nil),
			},
		},
		"deprecation_message": {
			input: &resource.NumberAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorNumberAttribute{
				CustomType:         convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				PlanModifiers:      convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:         convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.NumberAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorNumberAttribute{
				CustomType:    convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description:   convert.NewDescription(pointer("description")),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.NumberAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorNumberAttribute{
				CustomType:    convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:     convert.NewSensitive(pointer(true)),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.NumberAttribute{
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
				CustomType:    convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, nil),
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
		"plan-modifiers": {
			input: &resource.NumberAttribute{
				PlanModifiers: specschema.NumberPlanModifiers{
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
			expected: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.NumberAttribute{
				Default: &specschema.NumberDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
			expected: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "github.com/.../my_default",
						},
					},
					SchemaDefinition: "my_default.Default()",
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{}),
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

func TestGeneratorNumberAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorNumberAttribute
		expected []code.Import
	}{
		"default": {
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"custom-type-without-import": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{}, nil, ""),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					nil,
					"",
				),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import": {
			input: GeneratorNumberAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					nil,
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
			},
		},
		"validator-custom-nil": {
			input: GeneratorNumberAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorNumberAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{
					&specschema.CustomValidator{},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorNumberAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "",
							},
						},
					},
				})},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import": {
			input: GeneratorNumberAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/myotherproject/myvalidators/validator",
							},
						},
					},
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/myvalidators/validator",
							},
						},
					},
				})},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.ValidatorImport,
				},
				{
					Path: "github.com/myotherproject/myvalidators/validator",
				},
				{
					Path: "github.com/myproject/myvalidators/validator",
				},
			},
		},
		"plan-modifier-custom-nil": {
			input: GeneratorNumberAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorNumberAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifiers-custom-import-empty-string": {
			input: GeneratorNumberAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "",
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifier-custom-import": {
			input: GeneratorNumberAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/myotherproject/myplanmodifiers/planmodifier",
							},
						},
					},
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/myplanmodifiers/planmodifier",
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.PlanModifierImport,
				},
				{
					Path: "github.com/myotherproject/myplanmodifiers/planmodifier",
				},
				{
					Path: "github.com/myproject/myplanmodifiers/planmodifier",
				},
			},
		},
		"default-nil": {
			input: GeneratorNumberAttribute{},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-nil": {
			input: GeneratorNumberAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-empty-string": {
			input: GeneratorNumberAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "",
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import": {
			input: GeneratorNumberAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "github.com/myproject/mydefaults/default",
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/myproject/mydefaults/default",
				},
			},
		},
		"associated-external-type": {
			input: GeneratorNumberAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.NumberAttribute",
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types",
				},
				{
					Path: "fmt",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/diag",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/attr",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-go/tftypes",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
				},
			},
		},
		"associated-external-type-with-import": {
			input: GeneratorNumberAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.NumberAttribute",
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types",
				},
				{
					Path: "fmt",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/diag",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/attr",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-go/tftypes",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
				},
				{
					Path: "github.com/api",
				},
			},
		},
		"associated-external-type-with-custom-type": {
			input: GeneratorNumberAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.NumberAttribute",
					},
				},
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					nil,
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: "fmt",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/diag",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/attr",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-go/tftypes",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
				},
				{
					Path: "github.com/api",
				},
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports().All()

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
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorNumberAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorNumberAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Computed: true,
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

		"validators": {
			input: GeneratorNumberAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeNumber, []*specschema.CustomValidator{
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

		"plan-modifiers": {
			input: GeneratorNumberAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeNumber, []*specschema.CustomPlanModifier{
					{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"number_attribute": schema.NumberAttribute{
PlanModifiers: []planmodifier.Number{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorNumberAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					SchemaDefinition: "my_number_default.Default()",
				}),
			},
			expected: `"number_attribute": schema.NumberAttribute{
Default: my_number_default.Default(),
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
