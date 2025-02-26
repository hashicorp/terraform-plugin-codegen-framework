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

func TestGeneratorListNestedAttribute_New(t *testing.T) {
	t.Parallel()

	attributes, err := NewAttributes(datasource.Attributes{})

	if err != nil {
		t.Error(err)
	}

	testCases := map[string]struct {
		input         *datasource.ListNestedAttribute
		expected      GeneratorListNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.ListNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
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
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
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
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_attribute",
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
		"attributes-object-bool": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{})},
		},
		"required": {
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &datasource.ListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorListNestedAttribute{
				CustomType: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, nil),
			},
		},
		"deprecation_message": {
			input: &datasource.ListNestedAttribute{
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, nil),
			},
		},
		"description": {
			input: &datasource.ListNestedAttribute{
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					"name",
				),
				Validators: convert.NewValidators(convert.ValidatorTypeList, nil),
			},
		},
		"sensitive": {
			input: &datasource.ListNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: convert.NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					"name",
				),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeList, nil),
			},
		},
		"validators": {
			input: &datasource.ListNestedAttribute{
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
				CustomType: convert.NewCustomTypeNestedCollection(
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
				CustomType: convert.NewCustomTypeNestedCollection(
					&specschema.CustomType{},
				),
				NestedAttributeObject: convert.NewNestedAttributeObject(
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
			input: GeneratorListNestedAttribute{
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
				CustomType: convert.NewCustomTypeNestedCollection(
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
			input: GeneratorListNestedAttribute{
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
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
				CustomType: convert.NewCustomTypeNestedCollection(
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
			input: GeneratorListNestedAttribute{
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
			input: GeneratorListNestedAttribute{
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
			input: GeneratorListNestedAttribute{
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								String: &specschema.StringType{},
							}),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									},
								},
								nil,
								convert.NewValidators(convert.ValidatorTypeObject, nil),
								"nested_list_nested",
							),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
							CustomType: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
						},
					},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
				CustomType: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Type: "my_custom_type",
				}),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.Validators{},
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
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.Validators{},
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

		"computed": {
			input: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.Validators{},
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
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: convert.NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.Validators{},
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
					convert.Validators{},
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
					convert.Validators{},
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					attributeName,
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
					convert.NewValidators(convert.ValidatorTypeObject, nil),
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
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
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
				CustomType: convert.NewCustomTypeNestedCollection(
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
