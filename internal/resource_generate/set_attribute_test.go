// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorSetAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSetAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorSetAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSetAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorSetAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"elem-type-bool": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"elem-type-bool-with-import": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Import: pointer("github.com/my_account/my_project/element"),
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                {},
				"github.com/my_account/my_project/element": {},
			},
		},
		"elem-type-list-bool": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"elem-type-list-bool-with-import": {
			input: GeneratorSetAttribute{
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
				generatorschema.TypesImport:                {},
				"github.com/my_account/my_project/element": {},
			},
		},
		"elem-type-object": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"elem-type-object-bool": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
							{
								Name: "b",
								Bool: &specschema.BoolType{},
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
		"elem-type-object-bool-with-import": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
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
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                {},
				"github.com/my_account/my_project/element": {},
			},
		},
		"elem-type-object-with-imports": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
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
			},
			expected: map[string]struct{}{
				generatorschema.AttrImport:                 {},
				generatorschema.TypesImport:                {},
				"github.com/my_account/my_project/element": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorSetAttribute{
				Validators: []specschema.SetValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorSetAttribute{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorSetAttribute{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import": {
			input: GeneratorSetAttribute{
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
				generatorschema.TypesImport:                        {},
				generatorschema.ValidatorImport:                    {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
		"plan-modifier-custom-nil": {
			input: GeneratorSetAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorSetAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"plan-modifiers-custom-import-empty-string": {
			input: GeneratorSetAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"plan-modifier-custom-import": {
			input: GeneratorSetAttribute{
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
				generatorschema.TypesImport: {},
				planModifierImport:          {},
				"github.com/myotherproject/myplanmodifiers/planmodifier": {},
				"github.com/myproject/myplanmodifiers/planmodifier":      {},
			},
		},
		"default-nil": {
			input: GeneratorSetAttribute{},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"default-custom-nil": {
			input: GeneratorSetAttribute{
				Default: &specschema.SetDefault{},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"default-custom-import-nil": {
			input: GeneratorSetAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"default-custom-import-empty-string": {
			input: GeneratorSetAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"default-custom-import": {
			input: GeneratorSetAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						Import: pointer("github.com/myproject/mydefaults/default"),
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:               {},
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

func TestGeneratorSetAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorSetAttribute{
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
"set_attribute": schema.SetAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
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
			},
			expected: `
"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorSetAttribute{
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
"set_attribute": schema.SetAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
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
			},
			expected: `
"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
							{
								Name: "bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
							{
								Name: "obj",
								Object: &specschema.ObjectType{
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
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
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
			},
			expected: `
"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorSetAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Required: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Computed: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Sensitive: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Description: "description",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					DeprecationMessage: "deprecated",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
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
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Validators: []validator.Set{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
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
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
PlanModifiers: []planmodifier.Set{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_set_default.Default()",
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Default: my_set_default.Default(),
},`,
		},

		"element-type-bool-custom": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Type: "boolCustomType",
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: boolCustomType,
},`,
		},

		"element-type-float64-custom": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Float64: &specschema.Float64Type{
						CustomType: &specschema.CustomType{
							Type: "float64CustomType",
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: float64CustomType,
},`,
		},

		"element-type-int64-custom": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Int64: &specschema.Int64Type{
						CustomType: &specschema.CustomType{
							Type: "int64CustomType",
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: int64CustomType,
},`,
		},

		"element-type-list-custom": {
			input: GeneratorSetAttribute{
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
"set_attribute": schema.SetAttribute{
ElementType: customListType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-map-custom": {
			input: GeneratorSetAttribute{
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
"set_attribute": schema.SetAttribute{
ElementType: customMapType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-number-custom": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Number: &specschema.NumberType{
						CustomType: &specschema.CustomType{
							Type: "numberCustomType",
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: numberCustomType,
},`,
		},

		"element-type-object-custom": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
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
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": customBoolType,
},
},
},`,
		},

		"element-type-set-custom": {
			input: GeneratorSetAttribute{
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
"set_attribute": schema.SetAttribute{
ElementType: customSetType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-string-custom": {
			input: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{
						CustomType: &specschema.CustomType{
							Type: "stringCustomType",
						},
					},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
ElementType: stringCustomType,
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("set_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "SetAttribute",
				ValueType: "types.Set",
				TfsdkName: "set_attribute",
			},
		},
		"custom-type": {
			input: GeneratorSetAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "SetAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "set_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("set_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
