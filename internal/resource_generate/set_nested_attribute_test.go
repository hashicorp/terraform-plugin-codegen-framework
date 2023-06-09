// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	generatorschema "github/hashicorp/terraform-provider-code-generator/internal/schema"
)

func TestGeneratorSetNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSetNestedAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{},
		},
		"custom-type-without-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-and-nested-object-custom-type-without-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer("github.com/my_account/my_project/attribute"),
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"custom-type-with-import-with-nested-object-custom-type-with-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer("github.com/my_account/my_project/nested_object"),
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute":     {},
				"github.com/my_account/my_project/nested_object": {},
			},
		},
		"nested-list": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"nested-list-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/nested_list"),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/nested_list": {},
			},
		},
		"nested-list-with-custom-type-with-element-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/nested_list"),
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: pointer("github.com/my_account/my_project/bool"),
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/nested_list": {},
				"github.com/my_account/my_project/bool":        {},
			},
		},
		"nested-object": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"obj": GeneratorObjectAttribute{
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name: "bool",
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.AttrImport:  {},
				generatorschema.TypesImport: {},
			},
		},
		"nested-object-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"obj": GeneratorObjectAttribute{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/nested_object"),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/nested_object": {},
			},
		},
		"nested-object-with-custom-type-with-attribute-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"obj": GeneratorObjectAttribute{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/nested_object"),
							},
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name: "bool",
									Bool: &specschema.BoolType{
										CustomType: &specschema.CustomType{
											Import: pointer("github.com/my_account/my_project/bool"),
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/nested_object": {},
				"github.com/my_account/my_project/bool":          {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorSetNestedAttribute{
				Validators: []specschema.SetValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-nil": {
			input: GeneratorSetNestedAttribute{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import": {
			input: GeneratorSetNestedAttribute{
				Validators: []specschema.SetValidator{
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
		"nested-object-validator-custom-nil": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: nil,
						},
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-validator-custom-import-nil": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: &specschema.CustomValidator{
								Import: nil,
							},
						},
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-validator-custom-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: &specschema.CustomValidator{
								Import: pointer(""),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-validator-custom-import": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
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
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.ValidatorImport:                    {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
		"plan-modifier-custom-nil": {
			input: GeneratorSetNestedAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorSetNestedAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifiers-custom-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import": {
			input: GeneratorSetNestedAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
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
			input:    GeneratorSetNestedAttribute{},
			expected: map[string]struct{}{},
		},
		"default-custom-nil": {
			input: GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import-nil": {
			input: GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{},
				},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import": {
			input: GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer("github.com/myproject/mydefaults/default"),
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/myproject/mydefaults/default": {},
			},
		},
		"nested-object-plan-modifier-custom-nil": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					PlanModifiers: []specschema.ObjectPlanModifier{
						{
							Custom: nil,
						},
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-plan-modifier-custom-import-nil": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					PlanModifiers: []specschema.ObjectPlanModifier{
						{
							Custom: &specschema.CustomPlanModifier{
								Import: nil,
							},
						},
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-plan-modifiers-custom-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					PlanModifiers: []specschema.ObjectPlanModifier{
						{
							Custom: &specschema.CustomPlanModifier{
								Import: pointer(""),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-plan-modifier-custom-import": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					PlanModifiers: []specschema.ObjectPlanModifier{
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
					},
				},
			},
			expected: map[string]struct{}{
				planModifierImport: {},
				"github.com/myotherproject/myplanmodifiers/planmodifier": {},
				"github.com/myproject/myplanmodifiers/planmodifier":      {},
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

func TestGeneratorSetNestedAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"bool": GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"nested_list_nested": GeneratorSetNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: map[string]GeneratorAttribute{
									"bool": GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"object": GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name:   "str",
									String: &specschema.StringType{},
								},
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"nested_single_nested": GeneratorSingleNestedAttribute{
							Attributes: map[string]GeneratorAttribute{
								"bool": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_single_nested": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Required: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Optional: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Computed: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Sensitive: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Description: "description",
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSetNestedAttribute{
				Validators: []specschema.SetValidator{
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
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Validators: []validator.Set{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Type: "my_custom_type",
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
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
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},
},`,
		},
		"plan-modifiers": {
			input: GeneratorSetNestedAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
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
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
PlanModifiers: []planmodifier.Set{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_set_default.Default()",
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Default: my_set_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("set_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
