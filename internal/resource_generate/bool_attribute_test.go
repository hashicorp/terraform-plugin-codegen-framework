// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorBoolAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorBoolAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{},
		},
		"custom-type-without-import": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-nil": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer("github.com/myotherproject/myvalidators/validator"),
						},
					},
					{
						Custom: &specschema.CustomValidator{
							Import: pointer("github.com/myproject/myvalidators/validator"),
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.ValidatorImport:                    {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
		"plan-modifier-custom-nil": {
			input: GeneratorBoolAttribute{
				PlanModifiers: []specschema.BoolPlanModifier{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorBoolAttribute{
				PlanModifiers: []specschema.BoolPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifiers-custom-import-empty-string": {
			input: GeneratorBoolAttribute{
				PlanModifiers: []specschema.BoolPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import": {
			input: GeneratorBoolAttribute{
				PlanModifiers: []specschema.BoolPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: pointer("github.com/myotherproject/myplanmodifiers/planmodifier"),
						},
					},
					{
						Custom: &specschema.CustomPlanModifier{
							Import: pointer("github.com/myproject/myplanmodifiers/planmodifier"),
						},
					},
				}},
			expected: map[string]struct{}{
				planModifierImport: {},
				"github.com/myotherproject/myplanmodifiers/planmodifier": {},
				"github.com/myproject/myplanmodifiers/planmodifier":      {},
			},
		},
		"default-nil": {
			input:    GeneratorBoolAttribute{},
			expected: map[string]struct{}{},
		},
		"default-custom-and-static-nil": {
			input: GeneratorBoolAttribute{
				Default: &specschema.BoolDefault{},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import-nil": {
			input: GeneratorBoolAttribute{
				Default: &specschema.BoolDefault{
					Custom: &specschema.CustomDefault{},
				},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import-empty-string": {
			input: GeneratorBoolAttribute{
				Default: &specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import": {
			input: GeneratorBoolAttribute{
				Default: &specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer("github.com/myproject/mydefaults/default"),
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/myproject/mydefaults/default": {},
			},
		},
		"default-static": {
			input: GeneratorBoolAttribute{
				Default: &specschema.BoolDefault{
					Static: pointer(true),
				},
			},
			expected: map[string]struct{}{
				defaultBoolImport: {},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorBoolAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorBoolAttribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorBoolAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Required: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Optional: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Computed: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Sensitive: true,
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					Description: "description",
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorBoolAttribute{
				BoolAttribute: schema.BoolAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorBoolAttribute{
				Validators: []specschema.BoolValidator{
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
					},
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					},
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Validators: []validator.Bool{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorBoolAttribute{
				PlanModifiers: []specschema.BoolPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							SchemaDefinition: "my_plan_modifier.Modify()",
						},
					},
					{
						Custom: &specschema.CustomPlanModifier{
							SchemaDefinition: "my_other_plan_modifier.Modify()",
						},
					},
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
PlanModifiers: []planmodifier.Bool{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorBoolAttribute{
				Default: &specschema.BoolDefault{
					Static: pointer(true),
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Default: booldefault.StaticBool(true),
},`,
		},

		"default-custom": {
			input: GeneratorBoolAttribute{
				Default: &specschema.BoolDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_bool_default.Default()",
					},
				},
			},
			expected: `
"bool_attribute": schema.BoolAttribute{
Default: my_bool_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("bool_attribute")

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
