// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestGeneratorSingleNestedBlock_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSingleNestedBlock
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{},
		},
		"custom-type-without-import": {
			input: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"nested-attribute-list": {
			input: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: map[string]struct{}{

				typesImport: {},
			},
		},
		"nested-attribute-list-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						CustomType: &specschema.CustomType{
							Import: pointer("github.com/my_account/my_project/nested_list"),
						},
					},
				},
			},
			expected: map[string]struct{}{

				"github.com/my_account/my_project/nested_list": {},
			},
		},
		"nested-attribute-list-with-custom-type-with-element-with-custom-type": {
			input: GeneratorSingleNestedBlock{
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
			expected: map[string]struct{}{

				"github.com/my_account/my_project/nested_list": {},
				"github.com/my_account/my_project/bool":        {},
			},
		},
		"nested-attribute-object": {
			input: GeneratorSingleNestedBlock{
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
			expected: map[string]struct{}{

				attrImport:  {},
				typesImport: {},
			},
		},
		"nested-attribute-object-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"obj": GeneratorObjectAttribute{
						CustomType: &specschema.CustomType{
							Import: pointer("github.com/my_account/my_project/nested_object"),
						},
					},
				},
			},
			expected: map[string]struct{}{

				"github.com/my_account/my_project/nested_object": {},
			},
		},
		"nested-attribute-object-with-custom-type-with-attribute-with-custom-type": {
			input: GeneratorSingleNestedBlock{
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
			expected: map[string]struct{}{

				"github.com/my_account/my_project/nested_object": {},
				"github.com/my_account/my_project/bool":          {},
			},
		},
		"nested-block-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Blocks: map[string]GeneratorBlock{
					"list-nested-block": GeneratorListNestedBlock{
						CustomType: &specschema.CustomType{
							Import: pointer("github.com/my_account/my_project/nested_block"),
						},
					},
				},
			},
			expected: map[string]struct{}{

				"github.com/my_account/my_project/nested_block": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorSingleNestedBlock{
				Validators: []specschema.ObjectValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-nil": {
			input: GeneratorSingleNestedBlock{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorSingleNestedBlock{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import": {
			input: GeneratorSingleNestedBlock{
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
				}},
			expected: map[string]struct{}{
				validatorImport: {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
		"plan-modifier-custom-nil": {
			input: GeneratorSingleNestedBlock{
				PlanModifiers: []specschema.ObjectPlanModifier{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorSingleNestedBlock{
				PlanModifiers: []specschema.ObjectPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifiers-custom-import-empty-string": {
			input: GeneratorSingleNestedBlock{
				PlanModifiers: []specschema.ObjectPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import": {
			input: GeneratorSingleNestedBlock{
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
				}},
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

func TestGeneratorSingleNestedBlock_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSingleNestedBlock
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"bool": GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorSingleNestedBlock{
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
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorSingleNestedBlock{
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
			expected: `
"single_nested_block": schema.SingleNestedBlock{
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
},`,
		},

		"attribute-object": {
			input: GeneratorSingleNestedBlock{
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
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorSingleNestedBlock{
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
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"nested_single_nested": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},`,
		},

		"block-list-nested-bool": {
			input: GeneratorSingleNestedBlock{
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
			expected: `
"single_nested_block": schema.SingleNestedBlock{
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
},`,
		},

		"block-single-nested-bool": {
			input: GeneratorSingleNestedBlock{
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
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Blocks: map[string]schema.Block{
"nested_single_nested": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
CustomType: my_custom_type,
},`,
		},

		"description": {
			input: GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					Description: "description",
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSingleNestedBlock{
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
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Validators: []validator.Bool{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorSingleNestedBlock{
				PlanModifiers: []specschema.ObjectPlanModifier{
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
"single_nested_block": schema.SingleNestedBlock{
PlanModifiers: []planmodifier.Object{
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

			got, err := testCase.input.ToString("single_nested_block")

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
