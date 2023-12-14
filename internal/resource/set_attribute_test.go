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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorSetAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SetAttribute
		expected      GeneratorSetAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SetAttribute is nil"),
		},
		"element-type-bool": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
		},
		"element-type-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"element-type-list-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
		},
		"element-type-map-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
		},
		"element-type-list-object-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name:   "str",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name:   "str",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
		},
		"element-type-object-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
		},
		"element-type-object-list-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
		},
		"computed": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Computed: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"computed_optional": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Computed: true,
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"optional": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"required": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Required: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"custom_type": {
			input: &resource.SetAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"deprecation_message": {
			input: &resource.SetAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					DeprecationMessage: "deprecation message",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"description": {
			input: &resource.SetAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"sensitive": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Sensitive: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"validators": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.SetValidators{
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
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.SetValidators{
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
		},
		"plan-modifiers": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: specschema.SetPlanModifiers{
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
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: specschema.SetPlanModifiers{
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
		},
		"default": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.SetDefault{
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
			expected: GeneratorSetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.SetDefault{
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
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorSetAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			// TODO: This prevents misleading failure when ElementType for both got and expected are nil.
			// TODO: Could overwrite comparison using an option to cmp.Diff()?
			if got.SetAttribute.ElementType == nil && testCase.expected.SetAttribute.ElementType == nil {
				return
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetAttribute_Schema(t *testing.T) {
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
								AttributeTypes: specschema.ObjectAttributeTypes{
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
								AttributeTypes: specschema.ObjectAttributeTypes{
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
						AttributeTypes: specschema.ObjectAttributeTypes{
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
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "obj",
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
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
						AttributeTypes: specschema.ObjectAttributeTypes{
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
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorSetAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
CustomType: SetAttributeType{
types.SetType{
ElemType: types.StringType,
},
},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorSetAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"set_attribute": schema.SetAttribute{
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
				Validators: specschema.SetValidators{
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
				PlanModifiers: specschema.SetPlanModifiers{
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
						AttributeTypes: specschema.ObjectAttributeTypes{
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

			got, err := testCase.input.Schema("set_attribute")

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
		"associated-external-type": {
			input: GeneratorSetAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.BoolAttribute",
					},
				},
			},
			expected: model.Field{
				Name:      "SetAttribute",
				ValueType: "SetAttributeValue",
				TfsdkName: "set_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorSetAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.BoolAttribute",
					},
				},
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
