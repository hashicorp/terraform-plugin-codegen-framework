// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorSingleNestedBlock_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.SingleNestedBlock
		expected      GeneratorSingleNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.SingleNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", provider.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "bool_attribute",
						Bool: &provider.BoolAttribute{
							OptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"bool_attribute": GeneratorBoolAttribute{
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						CustomType:       convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
						Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "list_attribute",
						List: &provider.ListAttribute{
							OptionalRequired: "optional",
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
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						CustomType: convert.NewCustomTypeCollection(
							nil,
							nil,
							convert.CustomCollectionTypeList,
							"types.BoolType",
							"list_attribute",
						),
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
						ElementTypeCollection: convert.NewElementType(specschema.ElementType{
							Bool: &specschema.BoolType{},
						}),
						Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "nested_attribute",
						ListNested: &provider.ListNestedAttribute{
							NestedObject: provider.NestedAttributeObject{
								Attributes: []provider.Attribute{
									{
										Name: "nested_bool",
										Bool: &provider.BoolAttribute{
											OptionalRequired: "optional",
										},
									},
								},
							},
							OptionalRequired: "optional",
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
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
									CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
						},
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						NestedAttributeObject: convert.NewNestedAttributeObject(
							generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
									CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							nil,
							convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
							"nested_attribute",
						),
						Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "object_attribute",
						Object: &provider.ObjectAttribute{
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
							},
							OptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"object_attribute": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "obj_bool",
								Bool: &specschema.BoolType{},
							},
						},
						AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
							{
								Name: "obj_bool",
								Bool: &specschema.BoolType{},
							},
						}),
						CustomType:       convert.NewCustomTypeObject(nil, nil, "object_attribute"),
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						Validators:       convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "nested_attribute",
						SingleNested: &provider.SingleNestedAttribute{
							Attributes: []provider.Attribute{
								{
									Name: "nested_bool",
									Bool: &provider.BoolAttribute{
										OptionalRequired: "optional",
									},
								},
							},
							OptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: generatorschema.GeneratorAttributes{
							"nested_bool": GeneratorBoolAttribute{
								OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
								CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
								Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
							},
						},
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						CustomType:       convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
						Validators:       convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"blocks-nil": {
			input: &provider.SingleNestedBlock{
				Blocks: []provider.Block{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("block type not defined: %+v", provider.Block{
				Name: "empty",
			}),
		},
		"blocks-list-nested-bool": {
			input: &provider.SingleNestedBlock{
				Blocks: []provider.Block{
					{
						Name: "nested_block",
						ListNested: &provider.ListNestedBlock{
							NestedObject: provider.NestedBlockObject{
								Attributes: []provider.Attribute{
									{
										Name: "bool_attribute",
										Bool: &provider.BoolAttribute{
											OptionalRequired: "optional",
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
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
									CustomType:       convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
									Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
						},
						NestedBlockObject: convert.NewNestedBlockObject(
							generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
									CustomType:       convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
									Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							generatorschema.GeneratorBlocks{},
							nil,
							convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
							"nested_block",
						),
						Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"blocks-single-nested-bool": {
			input: &provider.SingleNestedBlock{
				Blocks: []provider.Block{
					{
						Name: "nested_block",
						SingleNested: &provider.SingleNestedBlock{
							Attributes: []provider.Attribute{
								{
									Name: "bool_attribute",
									Bool: &provider.BoolAttribute{
										OptionalRequired: "optional",
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
								OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
								CustomType:       convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
								Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
							},
						},
						CustomType: convert.NewCustomTypeNestedObject(nil, "nested_block"),
						Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &provider.SingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorSingleNestedBlock{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/",
						},
						Type:      "my_type",
						ValueType: "myvalue_type",
					},
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &provider.SingleNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorSingleNestedBlock{
				CustomType:         convert.NewCustomTypeNestedObject(nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				Validators:         convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &provider.SingleNestedBlock{
				Description: pointer("description"),
			},
			expected: GeneratorSingleNestedBlock{
				CustomType:  convert.NewCustomTypeNestedObject(nil, "name"),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &provider.SingleNestedBlock{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/.../myvalidator",
							},
						},
						SchemaDefinition: "myvalidator.Validate()",
					},
				}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorSingleNestedBlock("name", testCase.input)

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
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{},
					"",
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSingleNestedBlock{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					"",
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import": {
			input: GeneratorSingleNestedBlock{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					"",
				),
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
						CustomType: convert.NewCustomTypeCollection(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
							nil,
							convert.CustomCollectionTypeList,
							"",
							"",
						),
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
						CustomType: convert.NewCustomTypeCollection(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
							nil,
							convert.CustomCollectionTypeList,
							"",
							"",
						),
						ElementTypeCollection: convert.NewElementType(specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/bool",
									},
								},
							},
						}),
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
						CustomType: convert.NewCustomTypeObject(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
							nil,
							"",
						),
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
						AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
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
						}),
						CustomType: convert.NewCustomTypeObject(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
							nil,
							"",
						),
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
						CustomType: convert.NewCustomTypeNestedCollection(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_block",
								},
							},
						),
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
				Validators: convert.NewValidators(convert.ValidatorTypeObject, nil),
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
		"validator-custom-import-nil": {
			input: GeneratorSingleNestedBlock{
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{},
				}),
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
		"validator-custom-import-empty-string": {
			input: GeneratorSingleNestedBlock{
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "",
							},
						},
					},
				}),
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
		"validator-custom-import": {
			input: GeneratorSingleNestedBlock{
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/myotherproject/myvalidators/validator",
							},
						},
					},
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/myvalidators/validator",
							},
						},
					},
				}),
			},
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

	blockName := "single_nested_block"

	testCases := map[string]struct {
		input         GeneratorSingleNestedBlock
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"bool": GeneratorBoolAttribute{
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						ElementTypeCollection: convert.NewElementType(specschema.ElementType{
							String: &specschema.StringType{},
						}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
						NestedAttributeObject: convert.NewNestedAttributeObject(
							generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
								},
							},
							nil,
							convert.NewValidators(convert.ValidatorTypeObject, nil),
							"nested_list_nested",
						),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
						AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						}),
						OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
								OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							},
						},
						CustomType: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
						NestedBlockObject: convert.NewNestedBlockObject(
							generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
								},
							},
							generatorschema.GeneratorBlocks{},
							nil,
							convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
							"nested_list_nested",
						),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
								OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							},
						},
						CustomType: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					blockName,
				),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
CustomType: my_custom_type,
},`,
		},

		"description": {
			input: GeneratorSingleNestedBlock{
				CustomType:  convert.NewCustomTypeNestedObject(nil, blockName),
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
				CustomType:         convert.NewCustomTypeNestedObject(nil, blockName),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
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
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					"",
				),
			},
			expected: model.Field{
				Name:      "SingleNestedBlock",
				ValueType: "my_custom_value_type",
				TfsdkName: "single_nested_block",
			},
		},
	}

	for name, testCase := range testCases {

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
