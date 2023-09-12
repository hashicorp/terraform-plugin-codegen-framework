// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorObjectAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorObjectAttribute
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
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: []code.Import{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
			},
			expected: []code.Import{},
		},
		"custom-type-with-import": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
			},
		},
		"object-without-attribute-types": {
			input: GeneratorObjectAttribute{},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"object-with-empty-attribute-types": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
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
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"object-with-attr-type-bool-with-import": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "bool",
						Bool: &specschema.BoolType{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/element",
								},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/element",
				},
			},
		},
		"object-with-attr-type-bool-with-imports": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "bool",
						Bool: &specschema.BoolType{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/element",
								},
							},
						},
					},
					{
						Name: "list",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: &code.Import{
											Path: "github.com/my_account/my_project/another_element",
										},
									},
								},
							},
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/list",
								},
							},
						},
					},
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/element",
				},
				{
					Path: "github.com/my_account/my_project/list",
				},
				{
					Path: "github.com/my_account/my_project/another_element",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-nil": {
			input: GeneratorObjectAttribute{
				Validators: specschema.ObjectValidators{
					{
						Custom: nil,
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorObjectAttribute{
				Validators: specschema.ObjectValidators{
					{
						Custom: &specschema.CustomValidator{},
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorObjectAttribute{
				Validators: specschema.ObjectValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "",
								},
							},
						},
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import": {
			input: GeneratorObjectAttribute{
				Validators: specschema.ObjectValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myotherproject/myvalidators/validator",
								},
							},
						},
					},
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myproject/myvalidators/validator",
								},
							},
						},
					},
				}},
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports().All()

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

		"attr-type-obj-custom": {
			input: GeneratorObjectAttribute{
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
							CustomType: &specschema.CustomType{
								Type: "objCustomType",
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": objCustomType{
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
						Object: &specschema.ObjectType{
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name: "obj_obj",
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
				Validators: specschema.ObjectValidators{
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

		"attr-type-bool-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "bool",
						Bool: &specschema.BoolType{
							CustomType: &specschema.CustomType{
								Type: "boolCustomType",
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"bool": boolCustomType,
},
},`,
		},

		"attr-type-float64-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "float64",
						Float64: &specschema.Float64Type{
							CustomType: &specschema.CustomType{
								Type: "float64CustomType",
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"float64": float64CustomType,
},
},`,
		},

		"attr-type-int64-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "int64",
						Int64: &specschema.Int64Type{
							CustomType: &specschema.CustomType{
								Type: "int64CustomType",
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"int64": int64CustomType,
},
},`,
		},

		"attr-type-list-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "list",
						List: &specschema.ListType{
							CustomType: &specschema.CustomType{
								Type: "listCustomType",
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "boolCustomType",
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
"list": listCustomType{
ElemType: boolCustomType,
},
},
},`,
		},

		"attr-type-map-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "map",
						Map: &specschema.MapType{
							CustomType: &specschema.CustomType{
								Type: "mapCustomType",
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "boolCustomType",
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
"map": mapCustomType{
ElemType: boolCustomType,
},
},
},`,
		},

		"attr-type-number-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "number",
						Number: &specschema.NumberType{
							CustomType: &specschema.CustomType{
								Type: "numberCustomType",
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"number": numberCustomType,
},
},`,
		},

		"attr-type-object-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "object",
						Object: &specschema.ObjectType{
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name: "bool",
									Bool: &specschema.BoolType{
										CustomType: &specschema.CustomType{
											Type: "boolCustomType",
										},
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
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": boolCustomType,
},
},
},
},`,
		},

		"attr-type-set-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "set",
						Set: &specschema.SetType{
							CustomType: &specschema.CustomType{
								Type: "setCustomType",
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "boolCustomType",
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
"set": setCustomType{
ElemType: boolCustomType,
},
},
},`,
		},

		"attr-type-string-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "string",
						Number: &specschema.NumberType{
							CustomType: &specschema.CustomType{
								Type: "stringCustomType",
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"string": stringCustomType,
},
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

func TestGeneratorObjectAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorObjectAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "ObjectAttribute",
				ValueType: "types.Object",
				TfsdkName: "object_attribute",
			},
		},
		"custom-type": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "ObjectAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "object_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("object_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
