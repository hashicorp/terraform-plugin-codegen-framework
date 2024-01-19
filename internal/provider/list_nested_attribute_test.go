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

func TestGeneratorListNestedAttribute_New(t *testing.T) {
	t.Parallel()

	attributes, err := NewAttributes(provider.Attributes{})

	if err != nil {
		t.Error(err)
	}

	testCases := map[string]struct {
		input         *provider.ListNestedAttribute
		expected      GeneratorListNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.ListNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &provider.ListNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			input: &provider.ListNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
							CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
							CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &provider.ListNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							CustomTypeCollection: convert.NewCustomTypeCollection(
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
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							CustomTypeCollection: convert.NewCustomTypeCollection(
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
		"attributes-list-nested-bool": {
			input: &provider.ListNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
										CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
										CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
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
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
										CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
										CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_attribute",
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
		"attributes-object-bool": {
			input: &provider.ListNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
							CustomTypeObject: convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
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
							CustomTypeObject: convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &provider.ListNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
									CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							OptionalRequired:       convert.NewOptionalRequired(specschema.Optional),
							CustomTypeNestedObject: convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							ValidatorsCustom:       convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									OptionalRequired:    convert.NewOptionalRequired(specschema.Optional),
									CustomTypePrimitive: convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							OptionalRequired:       convert.NewOptionalRequired(specschema.Optional),
							CustomTypeNestedObject: convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							ValidatorsCustom:       convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &provider.ListNestedAttribute{
				OptionalRequired: "optional",
			},
			expected: GeneratorListNestedAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{})},
		},
		"required": {
			input: &provider.ListNestedAttribute{
				OptionalRequired: "required",
			},
			expected: GeneratorListNestedAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &provider.ListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"deprecation_message": {
			input: &provider.ListNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorListNestedAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"description": {
			input: &provider.ListNestedAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorListNestedAttribute{
				Description: convert.NewDescription(pointer("description")),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"sensitive": {
			input: &provider.ListNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				Sensitive:        convert.NewSensitive(pointer(true)),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"validators": {
			input: &provider.ListNestedAttribute{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
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

			got, err := NewGeneratorListNestedAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListNestedAttribute
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
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{},
				)},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{},
				),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					nil,
					&specschema.CustomType{},
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(
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
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(
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
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
				),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/nested_object",
						},
					},
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-list": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-list-with-custom-type": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							CustomTypeCollection: convert.NewCustomTypeCollection(
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							CustomTypeCollection: convert.NewCustomTypeCollection(
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"obj": GeneratorObjectAttribute{
							CustomTypeObject: convert.NewCustomTypeObject(
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
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
							CustomTypeObject: convert.NewCustomTypeObject(
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
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
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
		"validator-custom-nil": {
			input: GeneratorListNestedAttribute{
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
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
			input: GeneratorListNestedAttribute{
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
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
			input: GeneratorListNestedAttribute{
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
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
			input: GeneratorListNestedAttribute{
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
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
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					nil,
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{
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

func TestGeneratorListNestedAttribute_Schema(t *testing.T) {
	t.Parallel()

	attributeName := "list_nested_attribute"

	testCases := map[string]struct {
		input         GeneratorListNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"bool": GeneratorBoolAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								String: &specschema.StringType{},
							}),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_list_nested": GeneratorListNestedAttribute{
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"bool": GeneratorBoolAttribute{
										OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
									},
								},
								nil,
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
								"nested_list_nested",
							),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
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
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
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
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_single_nested": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
								},
							},
							CustomTypeNestedObject: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
						},
					},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
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
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Type: "my_custom_type",
				}),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorListNestedAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorListNestedAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Optional: true,
},`,
		},

		"sensitive": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorListNestedAttribute{
				Description: convert.NewDescription(pointer("description")),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListNestedAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
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
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{
						&specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
						&specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					}),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
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

			got, err := testCase.input.Schema("list_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListNestedAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListNestedAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "ListNestedAttribute",
				ValueType: "types.List",
				TfsdkName: "list_nested_attribute",
			},
		},
		"custom-type": {
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
				),
			},
			expected: model.Field{
				Name:      "ListNestedAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "list_nested_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("list_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
