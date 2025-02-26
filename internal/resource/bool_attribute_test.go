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

func TestGeneratorBoolAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.BoolAttribute
		expected      GeneratorBoolAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.BoolAttribute is nil"),
		},
		"computed": {
			input: &resource.BoolAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorBoolAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.BoolAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorBoolAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.BoolAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorBoolAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.BoolAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorBoolAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.BoolAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
				Validators:    convert.NewValidators(convert.ValidatorTypeBool, nil),
			},
		},
		"deprecation_message": {
			input: &resource.BoolAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorBoolAttribute{
				CustomType:         convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				PlanModifiers:      convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Validators:         convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.BoolAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorBoolAttribute{
				CustomType:    convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description:   convert.NewDescription(pointer("description")),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.BoolAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorBoolAttribute{
				CustomType:    convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Sensitive:     convert.NewSensitive(pointer(true)),
				Validators:    convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.BoolAttribute{
				Validators: specschema.BoolValidators{
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
			expected: GeneratorBoolAttribute{
				CustomType:    convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
				Validators: convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{
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
			input: &resource.BoolAttribute{
				PlanModifiers: specschema.BoolPlanModifiers{
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
			expected: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.BoolAttribute{
				Default: &specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(true),
				},
			},
			expected: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Default: convert.NewDefaultBool(&specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(true),
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorBoolAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorBoolAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorBoolAttribute
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
			input: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{}, nil, ""),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorBoolAttribute{
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
			input: GeneratorBoolAttribute{
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
			input: GeneratorBoolAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeBool, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorBoolAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{
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
			input: GeneratorBoolAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{
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
			input: GeneratorBoolAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{
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
			input: GeneratorBoolAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorBoolAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{
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
			input: GeneratorBoolAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{
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
			input: GeneratorBoolAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{
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
			input: GeneratorBoolAttribute{},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-and-static-nil": {
			input: GeneratorBoolAttribute{
				Default: convert.NewDefaultBool(&specschema.BoolDefault{}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-nil": {
			input: GeneratorBoolAttribute{
				Default: convert.NewDefaultBool(&specschema.BoolDefault{
					Custom: &specschema.CustomDefault{},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-empty-string": {
			input: GeneratorBoolAttribute{
				Default: convert.NewDefaultBool(&specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
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
		"default-custom-import": {
			input: GeneratorBoolAttribute{
				Default: convert.NewDefaultBool(&specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/mydefaults/default",
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
					Path: "github.com/myproject/mydefaults/default",
				},
			},
		},
		"default-static": {
			input: GeneratorBoolAttribute{
				Default: convert.NewDefaultBool(&specschema.BoolDefault{
					Static: pointer(true),
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault",
				},
			},
		},
		"associated-external-type": {
			input: GeneratorBoolAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.BoolAttribute",
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
			input: GeneratorBoolAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.BoolAttribute",
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
			input: GeneratorBoolAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.BoolAttribute",
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

func TestGeneratorBoolAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorBoolAttribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					"bool_attribute",
				),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ExtBool",
					},
					"bool_attribute",
				),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
CustomType: BoolAttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ExtBool",
					},
					"bool_attribute",
				),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorBoolAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorBoolAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorBoolAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorBoolAttribute{
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorBoolAttribute{
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorBoolAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorBoolAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeBool, []*specschema.CustomValidator{
					{
						SchemaDefinition: "my_validator.Validate()",
					},
					{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Validators: []validator.Bool{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorBoolAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeBool, []*specschema.CustomPlanModifier{
					{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
PlanModifiers: []planmodifier.Bool{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorBoolAttribute{
				Default: convert.NewDefaultBool(&specschema.BoolDefault{
					Static: pointer(true),
				}),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Default: booldefault.StaticBool(true),
},`,
		},

		"default-custom": {
			input: GeneratorBoolAttribute{
				Default: convert.NewDefaultBool(&specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_bool_default.Default()",
					},
				}),
			},
			expected: `"bool_attribute": schema.BoolAttribute{
Default: my_bool_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("bool_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorBoolAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorBoolAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "BoolAttribute",
				ValueType: "types.Bool",
				TfsdkName: "bool_attribute",
			},
		},
		"custom-type": {
			input: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					"",
				),
			},
			expected: model.Field{
				Name:      "BoolAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "bool_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.BoolAttribute",
					},
					"bool_attribute",
				),
			},
			expected: model.Field{
				Name:      "BoolAttribute",
				ValueType: "BoolAttributeValue",
				TfsdkName: "bool_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorBoolAttribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.BoolAttribute",
					},
					"",
				),
			},
			expected: model.Field{
				Name:      "BoolAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "bool_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("bool_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func pointer[T any](in T) *T {
	return &in
}

var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return x.Error() == y.Error()
})
