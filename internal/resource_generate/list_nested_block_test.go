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

func TestGeneratorListNestedBlock_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListNestedBlock
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{},
		},
		"custom-type-without-import": {
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-and-nested-object-custom-type-without-import": {
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{},
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-attribute-list": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-attribute-list-with-custom-type": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-attribute-list-with-custom-type-with-element-with-custom-type": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-attribute-object": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-attribute-object-with-custom-type": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-attribute-object-with-custom-type-with-attribute-with-custom-type": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-block-with-custom-type": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Blocks: map[string]GeneratorBlock{
						"list-nested-block": GeneratorListNestedBlock{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/nested_block"),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/nested_block": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorListNestedBlock{
				Validators: []specschema.ListValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-nil": {
			input: GeneratorListNestedBlock{
				Validators: []specschema.ListValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorListNestedBlock{
				Validators: []specschema.ListValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import": {
			input: GeneratorListNestedBlock{
				Validators: []specschema.ListValidator{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				PlanModifiers: []specschema.ListPlanModifier{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorListNestedBlock{
				PlanModifiers: []specschema.ListPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifiers-custom-import-empty-string": {
			input: GeneratorListNestedBlock{
				PlanModifiers: []specschema.ListPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import": {
			input: GeneratorListNestedBlock{
				PlanModifiers: []specschema.ListPlanModifier{
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
		"nested-object-plan-modifier-custom-nil": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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

func TestGeneratorListNestedBlock_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListNestedBlock
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: map[string]GeneratorAttribute{
						"nested_list_nested": GeneratorListNestedAttribute{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.ListNestedAttribute{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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

		"block-list-nested-bool": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Blocks: map[string]GeneratorBlock{
						"nested_list_nested": GeneratorListNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Blocks: map[string]schema.Block{
"nested_list_nested": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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

		"block-single-nested-bool": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Blocks: map[string]GeneratorBlock{
						"nested_single_nested": GeneratorSingleNestedBlock{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Blocks: map[string]schema.Block{
"nested_single_nested": schema.SingleNestedBlock{
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
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
},
CustomType: my_custom_type,
},`,
		},

		"description": {
			input: GeneratorListNestedBlock{
				ListNestedBlock: schema.ListNestedBlock{
					Description: "description",
				},
			},
			expected: `
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListNestedBlock{
				ListNestedBlock: schema.ListNestedBlock{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListNestedBlock{
				Validators: []specschema.ListValidator{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
},
Validators: []validator.List{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{
						Type: "my_custom_type",
					},
				},
			},
			expected: `
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorListNestedBlock{
				PlanModifiers: []specschema.ListPlanModifier{
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
"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
},
PlanModifiers: []planmodifier.List{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("list_nested_block")

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
