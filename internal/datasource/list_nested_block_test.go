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

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorListNestedBlock_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.ListNestedBlock
		expected      GeneratorListNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.ListNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Attributes: []datasource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", datasource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &datasource.ListNestedBlock{
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
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								Bool: &specschema.BoolType{},
							}),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								Bool: &specschema.BoolType{},
							}),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_attribute",
							),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_attribute",
							),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypeNestedObject:   convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypeNestedObject:   convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"blocks-nil": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Blocks: []datasource.Block{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("block type not defined: %+v", datasource.Block{
				Name: "empty",
			}),
		},

		"blocks-list-nested-bool": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Blocks: generatorschema.GeneratorBlocks{
						"nested_block": GeneratorListNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							NestedBlockObject: convert.NewNestedBlockObject(
								generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								generatorschema.GeneratorBlocks{},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_block",
							),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							NestedBlockObject: convert.NewNestedBlockObject(
								generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								generatorschema.GeneratorBlocks{},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_block",
							),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"blocks-single-nested-bool": {
			input: &datasource.ListNestedBlock{
				NestedObject: datasource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Blocks: generatorschema.GeneratorBlocks{
						"nested_block": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
						},
					},
				},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
						"nested_block": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"custom_type": {
			input: &datasource.ListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &datasource.ListNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorListNestedBlock{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				NestedObject:       GeneratorNestedBlockObject{},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &datasource.ListNestedBlock{
				Description: pointer("description"),
			},
			expected: GeneratorListNestedBlock{
				Description:  convert.NewDescription(pointer("description")),
				NestedObject: GeneratorNestedBlockObject{},
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &datasource.ListNestedBlock{
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
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
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
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
		name, testCase := name, testCase

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
				CustomType: &specschema.CustomType{},
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
				CustomType: &specschema.CustomType{},
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
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
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
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
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import": {
			input: GeneratorListNestedBlock{
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
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
				},
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
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/nested_object",
						},
					},
				},
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
				NestedObject: GeneratorNestedBlockObject{
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
				NestedObject: GeneratorNestedBlockObject{
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
				NestedObject: GeneratorNestedBlockObject{
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
				NestedObject: GeneratorNestedBlockObject{
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
				NestedObject: GeneratorNestedBlockObject{
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
				Validators: specschema.ListValidators{
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
			input: GeneratorListNestedBlock{
				Validators: specschema.ListValidators{
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
			input: GeneratorListNestedBlock{
				Validators: specschema.ListValidators{
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
			input: GeneratorListNestedBlock{
				Validators: specschema.ListValidators{
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
		"nested-object-validator-custom-nil": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: nil,
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
		"nested-object-validator-custom-import-nil": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: &specschema.CustomValidator{},
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
		"nested-object-validator-custom-import-empty-string": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
		"nested-object-validator-custom-import": {
			input: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
					},
				},
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								String: &specschema.StringType{},
							}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									},
								},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
								"nested_list_nested",
							),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
							CustomTypeNestedObject: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									},
								},
								generatorschema.GeneratorBlocks{},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_list_nested",
							),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Type: "my_custom_type",
				}),
				NestedBlockObject: convert.NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					blockName,
				),
				Validators: specschema.ListValidators{
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
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{
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
		name, testCase := name, testCase

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
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "ListNestedBlock",
				ValueType: "my_custom_value_type",
				TfsdkName: "list_nested_block",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
