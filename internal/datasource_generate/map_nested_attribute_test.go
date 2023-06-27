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

func TestGeneratorMapNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorMapNestedAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"custom-type-and-nested-object-custom-type-without-import": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                  {},
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"custom-type-with-import-with-nested-object-custom-type-with-import": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/nested_object",
						},
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute":     {},
				"github.com/my_account/my_project/nested_object": {},
			},
		},
		"nested-list": {
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                    {},
				"github.com/my_account/my_project/nested_list": {},
			},
		},
		"nested-list-with-custom-type-with-element-with-custom-type": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: &code.Import{
											Path: "github.com/my_account/my_project/bool",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                    {},
				"github.com/my_account/my_project/nested_list": {},
				"github.com/my_account/my_project/bool":        {},
			},
		},
		"nested-object": {
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"obj": GeneratorObjectAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                      {},
				"github.com/my_account/my_project/nested_object": {},
			},
		},
		"nested-object-with-custom-type-with-attribute-with-custom-type": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"obj": GeneratorObjectAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name: "bool",
									Bool: &specschema.BoolType{
										CustomType: &specschema.CustomType{
											Import: &code.Import{
												Path: "github.com/my_account/my_project/bool",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                      {},
				"github.com/my_account/my_project/nested_object": {},
				"github.com/my_account/my_project/bool":          {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorMapNestedAttribute{
				Validators: []specschema.MapValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorMapNestedAttribute{
				Validators: []specschema.MapValidator{
					{
						Custom: &specschema.CustomValidator{},
					},
				}},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorMapNestedAttribute{
				Validators: []specschema.MapValidator{
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
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"validator-custom-import": {
			input: GeneratorMapNestedAttribute{
				Validators: []specschema.MapValidator{
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
			expected: map[string]struct{}{
				generatorschema.TypesImport:                        {},
				generatorschema.ValidatorImport:                    {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
		"nested-object-validator-custom-nil": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: nil,
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"nested-object-validator-custom-import-nil": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: &specschema.CustomValidator{},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"nested-object-validator-custom-import-empty-string": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: &specschema.CustomValidator{
								Imports: []code.Import{
									{
										Path: "",
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport: {},
			},
		},
		"nested-object-validator-custom-import": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
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
					},
				},
			},
			expected: map[string]struct{}{
				generatorschema.TypesImport:                        {},
				generatorschema.ValidatorImport:                    {},
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

func TestGeneratorMapNestedAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorMapNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorMapNestedAttribute{
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
"map_nested_attribute": schema.MapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
"map_nested_attribute": schema.MapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"nested_list_nested": GeneratorMapNestedAttribute{
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
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.MapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
"map_nested_attribute": schema.MapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
"map_nested_attribute": schema.MapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Required: true,
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Optional: true,
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Computed: true,
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Sensitive: true,
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Description: "description",
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorMapNestedAttribute{
				Validators: []specschema.MapValidator{
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
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Validators: []validator.Map{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Type: "my_custom_type",
					},
				},
			},
			expected: `
"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorMapNestedAttribute{
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
"map_nested_attribute": schema.MapNestedAttribute{
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("map_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorMapNestedAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorMapNestedAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "MapNestedAttribute",
				ValueType: "types.Map",
				TfsdkName: "map_nested_attribute",
			},
		},
		"custom-type": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "MapNestedAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "map_nested_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("map_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
