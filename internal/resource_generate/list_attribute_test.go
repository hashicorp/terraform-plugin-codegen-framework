// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratorListAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{},
		},
		"custom-type-without-import": {
			input: GeneratorListAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorListAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorListAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"elem-type-bool": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: map[string]struct{}{

				typesImport: {},
			},
		},
		"elem-type-bool-with-import": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Import: pointer("github.com/my_account/my_project/element"),
						},
					},
				},
			},
			expected: map[string]struct{}{

				"github.com/my_account/my_project/element": {},
			},
		},
		"elem-type-list-bool": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
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
		"elem-type-list-bool-with-import": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Import: pointer("github.com/my_account/my_project/element"),
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{

				"github.com/my_account/my_project/element": {},
			},
		},
		"elem-type-object": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{},
				},
			},
			expected: map[string]struct{}{},
		},
		"elem-type-object-bool": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "b",
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: map[string]struct{}{

				attrImport:  {},
				typesImport: {},
			},
		},
		"elem-type-object-bool-with-import": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "bool",
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Import: pointer("github.com/my_account/my_project/element"),
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{

				"github.com/my_account/my_project/element": {},
			},
		},
		"elem-type-object-with-imports": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "b",
							Bool: &specschema.BoolType{},
						},
						{
							Name: "c",
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Import: pointer("github.com/my_account/my_project/element"),
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{

				attrImport:  {},
				typesImport: {},
				"github.com/my_account/my_project/element": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorListAttribute{
				Validators: []specschema.ListValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"validator-custom-import-nil": {
			input: GeneratorListAttribute{
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
			input: GeneratorListAttribute{
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
			input: GeneratorListAttribute{
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
				validatorImport: {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
		"plan-modifier-custom-nil": {
			input: GeneratorListAttribute{
				PlanModifiers: []specschema.ListPlanModifier{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorListAttribute{
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
			input: GeneratorListAttribute{
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
			input: GeneratorListAttribute{
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
		"default-nil": {
			input:    GeneratorListAttribute{},
			expected: map[string]struct{}{},
		},
		"default-custom-nil": {
			input: GeneratorListAttribute{
				Default: &specschema.ListDefault{},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import-nil": {
			input: GeneratorListAttribute{
				Default: &specschema.ListDefault{
					Custom: &specschema.CustomDefault{},
				},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import-empty-string": {
			input: GeneratorListAttribute{
				Default: &specschema.ListDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"default-custom-import": {
			input: GeneratorListAttribute{
				Default: &specschema.ListDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer("github.com/myproject/mydefaults/default"),
					},
				},
			},
			expected: map[string]struct{}{

				"github.com/myproject/mydefaults/default": {},
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

func TestGeneratorListAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: []specschema.ObjectAttributeType{
								{
									Name: "bool",
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"element-type-map": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Map: &specschema.MapType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Object: []specschema.ObjectAttributeType{
								{
									Name: "bool",
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"element-type-object": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "bool",
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "obj",
							Object: []specschema.ObjectAttributeType{
								{
									Name: "bool",
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},`,
		},

		"element-type-object-list": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "list",
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},
},`,
		},

		"element-type-string": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorListAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Required: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Computed: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Sensitive: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Description: "description",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType:        types.StringType,
					DeprecationMessage: "deprecated",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
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
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Validators: []validator.List{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
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
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
PlanModifiers: []planmodifier.List{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorListAttribute{
				Default: &specschema.ListDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_list_default.Default()",
					},
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Default: my_list_default.Default(),
},`,
		},

		"element-type-bool-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Type: "boolCustomType",
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: boolCustomType,
},`,
		},

		"element-type-float64-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Float64: &specschema.Float64Type{
						CustomType: &specschema.CustomType{
							Type: "float64CustomType",
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: float64CustomType,
},`,
		},

		"element-type-int64-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Int64: &specschema.Int64Type{
						CustomType: &specschema.CustomType{
							Type: "int64CustomType",
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: int64CustomType,
},`,
		},

		"element-type-list-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						CustomType: &specschema.CustomType{
							Type: "customListType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: customListType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-map-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						CustomType: &specschema.CustomType{
							Type: "customMapType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: customMapType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-number-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Number: &specschema.NumberType{
						CustomType: &specschema.CustomType{
							Type: "numberCustomType",
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: numberCustomType,
},`,
		},

		"element-type-object-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "bool",
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": customBoolType,
},
},
},`,
		},

		"element-type-set-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Set: &specschema.SetType{
						CustomType: &specschema.CustomType{
							Type: "customSetType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: customSetType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-string-custom": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{
						CustomType: &specschema.CustomType{
							Type: "stringCustomType",
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: stringCustomType,
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("list_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
