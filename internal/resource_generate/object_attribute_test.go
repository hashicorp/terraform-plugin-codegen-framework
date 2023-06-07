// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestGeneratorObjectAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorObjectAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"object-without-attribute-types": {
			input: GeneratorObjectAttribute{},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"object-with-empty-attribute-types": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{},
			},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"object-with-attr-type-bool": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
			expected: map[string]struct{}{
				schemaImport: {},
				attrImport:   {},
				typesImport:  {},
			},
		},
		"object-with-attr-type-bool-with-import": {
			input: GeneratorObjectAttribute{
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
			expected: map[string]struct{}{
				schemaImport: {},
				"github.com/my_account/my_project/element": {},
			},
		},
		"object-with-attr-type-bool-with-imports": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "bool",
						Bool: &specschema.BoolType{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/element"),
							},
						},
					},
					{
						Name: "list",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: pointer("github.com/my_account/my_project/another_element"),
									},
								},
							},
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/list"),
							},
						},
					},
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: map[string]struct{}{
				schemaImport: {},
				"github.com/my_account/my_project/element":         {},
				"github.com/my_account/my_project/another_element": {},
				"github.com/my_account/my_project/list":            {},
				attrImport:                                         {},
				typesImport:                                        {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{
				schemaImport: {},
			},
		},
		"validator-custom-import": {
			input: GeneratorObjectAttribute{
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
				schemaImport:    {},
				validatorImport: {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
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

func TestGeneratorObjectAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorObjectAttribute
		expected      string
		expectedError error
	}{
		"attr-type-bool": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},`,
		},

		"attr-type-list": {
			input: GeneratorObjectAttribute{
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
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-list-list": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "list",
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
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},
},`,
		},

		"attr-type-list-object": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "list",
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
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},`,
		},

		"attr-type-map": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "map",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-map-map": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "map",
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
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},
},`,
		},

		"attr-type-map-object": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "map",
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
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},`,
		},

		"attr-type-object": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
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
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"attr-type-object-object": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj",
						Object: []specschema.ObjectAttributeType{
							{
								Name: "obj_obj",
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
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"obj_obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},
},`,
		},

		"attr-type-object-list": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj",
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
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},
},
},`,
		},

		"attr-type-string": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
},`,
		},

		"custom-type": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Required: true,
				},
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorObjectAttribute{
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
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Computed: true,
				},
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Sensitive: true,
				},
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Description: "description",
				},
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					DeprecationMessage: "deprecated",
				},
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
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
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Validators: []validator.Bool{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
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
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
PlanModifiers: []planmodifier.Object{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
				Default: &specschema.ObjectDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_object_default.Default()",
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Default: my_object_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("object_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
