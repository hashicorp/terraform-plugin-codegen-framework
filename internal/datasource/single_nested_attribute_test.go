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

func TestGeneratorSingleNestedAttribute_New(t *testing.T) {
	t.Parallel()

	attributes, err := NewAttributes(datasource.Attributes{})

	if err != nil {
		t.Error(err)
	}

	testCases := map[string]struct {
		input         *datasource.SingleNestedAttribute
		expected      GeneratorSingleNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.SingleNestedAttribute is nil"),
		},
		"attributes-nil": {
			input: &datasource.SingleNestedAttribute{
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
			input: &datasource.SingleNestedAttribute{
				Attributes: []datasource.Attribute{
					{
						Name: "bool_attribute",
						Bool: &datasource.BoolAttribute{
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"bool_attribute": GeneratorBoolAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
						Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &datasource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &datasource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &datasource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &datasource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, nil),
			},
		},
		"computed": {
			input: &datasource.SingleNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &datasource.SingleNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &datasource.SingleNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &datasource.SingleNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &datasource.SingleNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorSingleNestedAttribute{
				Attributes: attributes,
				CustomType: convert.NewCustomTypeNestedObject(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &datasource.SingleNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorSingleNestedAttribute{
				Attributes:         attributes,
				CustomType:         convert.NewCustomTypeNestedObject(nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				Validators:         convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &datasource.SingleNestedAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorSingleNestedAttribute{
				Attributes:  attributes,
				CustomType:  convert.NewCustomTypeNestedObject(nil, "name"),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &datasource.SingleNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorSingleNestedAttribute{
				Attributes: attributes,
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &datasource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
				Attributes: attributes,
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

			got, err := NewGeneratorSingleNestedAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSingleNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSingleNestedAttribute
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
			input: GeneratorSingleNestedAttribute{
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
		"validator-custom-nil": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					nil,
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
		"validator-custom-import-nil": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					specschema.CustomValidators{
						&specschema.CustomValidator{},
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
		"validator-custom-import-empty-string": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					specschema.CustomValidators{
						&specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "",
								},
							},
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
		"validator-custom-import": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					specschema.CustomValidators{
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
					},
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

func TestGeneratorSingleNestedAttribute_Schema(t *testing.T) {
	t.Parallel()

	attributeName := "single_nested_attribute"

	testCases := map[string]struct {
		input         GeneratorSingleNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"bool": GeneratorBoolAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"list": GeneratorListAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						ElementTypeCollection: convert.NewElementType(specschema.ElementType{
							String: &specschema.StringType{},
						}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
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
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
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
				CustomType: convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_single_nested": GeneratorSingleNestedAttribute{
						Attributes: generatorschema.GeneratorAttributes{
							"bool": GeneratorBoolAttribute{
								ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							},
						},
						CustomType: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
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
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"custom-type": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(&specschema.CustomType{
					Type: "my_custom_type",
				}, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypeNestedObject(nil, attributeName),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, attributeName),
				Sensitive:  convert.NewSensitive(pointer(true)),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorSingleNestedAttribute{
				CustomType:  convert.NewCustomTypeNestedObject(nil, attributeName),
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSingleNestedAttribute{
				CustomType:         convert.NewCustomTypeNestedObject(nil, attributeName),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, attributeName),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
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

			got, err := testCase.input.Schema("single_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSingleNestedAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSingleNestedAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "SingleNestedAttribute",
				ValueType: "SingleNestedAttributeValue",
				TfsdkName: "single_nested_attribute",
			},
		},
		"custom-type": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					"single_nested_attribute",
				),
			},
			expected: model.Field{
				Name:      "SingleNestedAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "single_nested_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("single_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
