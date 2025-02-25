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

func TestGeneratorListNestedBlock_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.ListNestedBlock
		expected      GeneratorListNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.ListNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Attributes: []provider.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", provider.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &provider.ListNestedBlock{
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
			expected: GeneratorListNestedBlock{
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
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							CustomType:       convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							Validators:       convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							CustomType:       convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							Validators:       convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"blocks-nil": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Blocks: []provider.Block{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("block type not defined: %+v", provider.Block{
				Name: "empty",
			}),
		},

		"blocks-list-nested-bool": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
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
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"blocks-single-nested-bool": {
			input: &provider.ListNestedBlock{
				NestedObject: provider.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
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
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"custom_type": {
			input: &provider.ListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}),
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &provider.ListNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorListNestedBlock{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				NestedObject:       GeneratorNestedBlockObject{},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &provider.ListNestedBlock{
				Description: pointer("description"),
			},
			expected: GeneratorListNestedBlock{
				Description:  convert.NewDescription(pointer("description")),
				NestedObject: GeneratorNestedBlockObject{},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &provider.ListNestedBlock{
				Validators: specschema.ListValidators{
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
			expected: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
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

			got, err := NewGeneratorListNestedBlock("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListNestedBlock_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListNestedBlock
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
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{},
				)},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{},
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
		"custom-type-and-nested-object-custom-type-without-import": {
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{},
				),
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{},
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				),
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
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "",
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
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				),
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					convert.Validators{},
					"",
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-with-nested-object-custom-type-with-import": {
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
				),
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/nested_object",
						},
					},
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-list": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							ElementType: specschema.ElementType{
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
		"nested-list-with-custom-type": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					nil,
					nil,
					convert.Validators{},
					"",
				),
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
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					nil,
					nil,
					convert.Validators{},
					"",
				),
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
		"nested-object": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-object-with-custom-type": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					nil,
					nil,
					convert.Validators{},
					"",
				),
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
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					nil,
					nil,
					convert.Validators{},
					"",
				),
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
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					generatorschema.GeneratorBlocks{
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
					nil,
					convert.Validators{},
					"",
				),
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
			input: GeneratorListNestedBlock{
				Validators: convert.NewValidators(convert.ValidatorTypeList, nil),
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
			input: GeneratorListNestedBlock{
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
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
			input: GeneratorListNestedBlock{
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
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
			input: GeneratorListNestedBlock{
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
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
		"nested-object-validator-custom-nil": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					"",
				),
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
		"nested-object-validator-custom-import-nil": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
						&specschema.CustomValidator{},
					}),
					"",
				),
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
		"nested-object-validator-custom-import-empty-string": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
						&specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "",
								},
							},
						},
					}),
					"",
				),
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
		"nested-object-validator-custom-import": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
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
					"",
				),
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

func TestGeneratorListNestedBlock_Schema(t *testing.T) {
	t.Parallel()

	blockName := "list_nested_block"

	testCases := map[string]struct {
		input         GeneratorListNestedBlock
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"bool": GeneratorBoolAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								String: &specschema.StringType{},
							}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"nested_single_nested": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
								},
							},
							CustomType: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"block-list-nested-bool": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
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
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"block-single-nested-bool": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
						"nested_single_nested": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
								},
							},
							CustomType: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
							Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
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
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Type: "my_custom_type",
				}),
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
CustomType: my_custom_type,
},`,
		},

		"description": {
			input: GeneratorListNestedBlock{
				Description: convert.NewDescription(pointer("description")),
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListNestedBlock{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
},
},
Validators: []validator.List{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
						&specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
						&specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					}),
					blockName,
				),
			},
			expected: `"list_nested_block": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: ListNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedBlockValue{}.AttributeTypes(ctx),
},
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("list_nested_block")

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListNestedBlock_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListNestedBlock
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "ListNestedBlock",
				ValueType: "types.List",
				TfsdkName: "list_nested_block",
			},
		},
		"custom-type": {
			input: GeneratorListNestedBlock{
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
				),
			},
			expected: model.Field{
				Name:      "ListNestedBlock",
				ValueType: "my_custom_value_type",
				TfsdkName: "list_nested_block",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("list_nested_block")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
