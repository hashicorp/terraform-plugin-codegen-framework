// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorSingleNestedBlock_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.SingleNestedBlock
		expected      GeneratorSingleNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.SingleNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", datasource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "bool_attribute",
						Bool: &datasource.BoolAttribute{
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"bool_attribute": GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "list_attribute",
						List: &datasource.ListAttribute{
							ComputedOptionalRequired: "optional",
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"list_attribute": GeneratorListAttribute{
						ListAttribute: schema.ListAttribute{
							Optional: true,
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "nested_attribute",
						ListNested: &datasource.ListNestedAttribute{
							NestedObject: datasource.NestedAttributeObject{
								Attributes: []datasource.Attribute{
									{
										Name: "nested_bool",
										Bool: &datasource.BoolAttribute{
											ComputedOptionalRequired: "computed",
										},
									},
								},
							},
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Computed: true,
									},
								},
							},
						},
						ListNestedAttribute: schema.ListNestedAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "object_attribute",
						Object: &datasource.ObjectAttribute{
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
							},
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"object_attribute": GeneratorObjectAttribute{
						ObjectAttribute: schema.ObjectAttribute{
							Optional: true,
						},
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "obj_bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "nested_attribute",
						SingleNested: &datasource.SingleNestedAttribute{
							Attributes: []datasource.Attribute{
								{
									Name: "nested_bool",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: "computed",
									},
								},
							},
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: generatorschema.GeneratorAttributes{
							"nested_bool": GeneratorBoolAttribute{
								BoolAttribute: schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						SingleNestedAttribute: schema.SingleNestedAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"blocks-nil": {
			input: &datasource.SingleNestedBlock{
				Blocks: []datasource.Block{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("block type not defined: %+v", datasource.Block{
				Name: "empty",
			}),
		},
		"blocks-list-nested-bool": {
			input: &datasource.SingleNestedBlock{
				Blocks: []datasource.Block{
					{
						Name: "nested_block",
						ListNested: &datasource.ListNestedBlock{
							NestedObject: datasource.NestedBlockObject{
								Attributes: []datasource.Attribute{
									{
										Name: "bool_attribute",
										Bool: &datasource.BoolAttribute{
											ComputedOptionalRequired: "optional",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Blocks: generatorschema.GeneratorBlocks{
					"nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{
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
		"blocks-single-nested-bool": {
			input: &datasource.SingleNestedBlock{
				Blocks: []datasource.Block{
					{
						Name: "nested_block",
						SingleNested: &datasource.SingleNestedBlock{
							Attributes: []datasource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: "optional",
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Blocks: generatorschema.GeneratorBlocks{
					"nested_block": GeneratorSingleNestedBlock{
						Attributes: generatorschema.GeneratorAttributes{
							"bool_attribute": GeneratorBoolAttribute{
								BoolAttribute: schema.BoolAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"custom_type": {
			input: &datasource.SingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.SingleNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.SingleNestedBlock{
				Description: pointer("description"),
			},
			expected: GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"validators": {
			input: &datasource.SingleNestedBlock{
				Validators: specschema.ObjectValidators{
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
			expected: GeneratorSingleNestedBlock{
				Validators: specschema.ObjectValidators{
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorSingleNestedBlock(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSingleNestedBlock_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSingleNestedBlock
		expected []code.Import
	}{
		"default": {
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-without-import": {
			input: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{},
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import": {
			input: GeneratorSingleNestedBlock{
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
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-attribute-list": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
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
		"nested-attribute-list-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"list": GeneratorListAttribute{
						CustomType: &specschema.CustomType{
							Import: &code.Import{
								Path: "github.com/my_account/my_project/nested_list",
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
					Path: "github.com/my_account/my_project/nested_list",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-list-with-custom-type-with-element-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
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
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_list",
				},
				{
					Path: "github.com/my_account/my_project/bool",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-attribute-object": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"obj": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{},
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
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-attribute-object-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"obj": GeneratorObjectAttribute{
						CustomType: &specschema.CustomType{
							Import: &code.Import{
								Path: "github.com/my_account/my_project/nested_object",
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
					Path: "github.com/my_account/my_project/nested_object",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-with-custom-type-with-attribute-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"obj": GeneratorObjectAttribute{
						CustomType: &specschema.CustomType{
							Import: &code.Import{
								Path: "github.com/my_account/my_project/nested_object",
							},
						},
						AttributeTypes: specschema.ObjectAttributeTypes{
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
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				},
				{
					Path: "github.com/my_account/my_project/bool",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-block-with-custom-type": {
			input: GeneratorSingleNestedBlock{
				Blocks: generatorschema.GeneratorBlocks{
					"list-nested-block": GeneratorListNestedBlock{
						CustomType: &specschema.CustomType{
							Import: &code.Import{
								Path: "github.com/my_account/my_project/nested_block",
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
					Path: "github.com/my_account/my_project/nested_block",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-nil": {
			input: GeneratorSingleNestedBlock{
				Validators: specschema.ObjectValidators{
					{
						Custom: nil,
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorSingleNestedBlock{
				Validators: specschema.ObjectValidators{
					{
						Custom: &specschema.CustomValidator{},
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorSingleNestedBlock{
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
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-import": {
			input: GeneratorSingleNestedBlock{
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
				{
					Path: generatorschema.AttrImport,
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

func TestGeneratorSingleNestedBlock_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSingleNestedBlock
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
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
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
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
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_list_nested": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: generatorschema.GeneratorAttributes{
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
CustomType: NestedListNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedListNestedValue{}.AttributeTypes(ctx),
},
},
},
},
},
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						ObjectAttribute: schema.ObjectAttribute{
							Optional: true,
						},
						AttributeTypes: specschema.ObjectAttributeTypes{
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
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_single_nested": GeneratorSingleNestedAttribute{
						Attributes: generatorschema.GeneratorAttributes{
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
CustomType: NestedSingleNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedSingleNestedValue{}.AttributeTypes(ctx),
},
},
},
},
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"block-list-nested-bool": {
			input: GeneratorSingleNestedBlock{
				Blocks: generatorschema.GeneratorBlocks{
					"nested_list_nested": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: generatorschema.GeneratorAttributes{
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
CustomType: NestedListNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedListNestedValue{}.AttributeTypes(ctx),
},
},
},
},
},
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"block-single-nested-bool": {
			input: GeneratorSingleNestedBlock{
				Blocks: generatorschema.GeneratorBlocks{
					"nested_single_nested": GeneratorSingleNestedBlock{
						Attributes: generatorschema.GeneratorAttributes{
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
CustomType: NestedSingleNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedSingleNestedValue{}.AttributeTypes(ctx),
},
},
},
},
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
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
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
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
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSingleNestedBlock{
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
"single_nested_block": schema.SingleNestedBlock{
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("single_nested_block")

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSingleNestedBlock_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSingleNestedBlock
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "SingleNestedBlock",
				ValueType: "SingleNestedBlockValue",
				TfsdkName: "single_nested_block",
			},
		},
		"custom-type": {
			input: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "SingleNestedBlock",
				ValueType: "my_custom_value_type",
				TfsdkName: "single_nested_block",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("single_nested_block")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
