// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorSingleNestedBlock_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SingleNestedBlock
		expected      GeneratorSingleNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SingleNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &resource.SingleNestedBlock{
				Attributes: []resource.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.SingleNestedBlock{
				Attributes: []resource.Attribute{
					{
						Name: "bool_attribute",
						Bool: &resource.BoolAttribute{
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				Attributes: generatorschema.GeneratorAttributes{
					"bool_attribute": GeneratorBoolAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
						PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
						Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &resource.SingleNestedBlock{
				Attributes: []resource.Attribute{
					{
						Name: "list_attribute",
						List: &resource.ListAttribute{
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
						PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
						Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.SingleNestedBlock{
				Attributes: []resource.Attribute{
					{
						Name: "nested_attribute",
						ListNested: &resource.ListNestedAttribute{
							NestedObject: resource.NestedAttributeObject{
								Attributes: []resource.Attribute{
									{
										Name: "nested_bool",
										Bool: &resource.BoolAttribute{
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
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
						},
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						NestedAttributeObject: NewNestedAttributeObject(
							generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							nil,
							convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
							"nested_attribute",
						),
						PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
						Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &resource.SingleNestedBlock{
				Attributes: []resource.Attribute{
					{
						Name: "object_attribute",
						Object: &resource.ObjectAttribute{
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
						CustomType:               convert.NewCustomTypeObject(nil, nil, "object_attribute"),
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
						Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.SingleNestedBlock{
				Attributes: []resource.Attribute{
					{
						Name: "nested_attribute",
						SingleNested: &resource.SingleNestedAttribute{
							Attributes: []resource.Attribute{
								{
									Name: "nested_bool",
									Bool: &resource.BoolAttribute{
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
								ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
								CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
								PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
								Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
							},
						},
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						CustomType:               convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
						PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
						Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"blocks-nil": {
			input: &resource.SingleNestedBlock{
				Blocks: []resource.Block{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("block type not defined: %+v", resource.Block{
				Name: "empty",
			}),
		},
		"blocks-list-nested-bool": {
			input: &resource.SingleNestedBlock{
				Blocks: []resource.Block{
					{
						Name: "nested_block",
						ListNested: &resource.ListNestedBlock{
							NestedObject: resource.NestedBlockObject{
								Attributes: []resource.Attribute{
									{
										Name: "bool_attribute",
										Bool: &resource.BoolAttribute{
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
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
						},
						NestedBlockObject: NewNestedBlockObject(
							generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							generatorschema.GeneratorBlocks{},
							nil,
							convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
							"nested_block",
						),
						PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
						Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"blocks-single-nested-bool": {
			input: &resource.SingleNestedBlock{
				Blocks: []resource.Block{
					{
						Name: "nested_block",
						SingleNested: &resource.SingleNestedBlock{
							Attributes: []resource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &resource.BoolAttribute{
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
								ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
								PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
								Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
							},
						},
						CustomType:    convert.NewCustomTypeNestedObject(nil, "nested_block"),
						PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
						Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.SingleNestedBlock{
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &resource.SingleNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorSingleNestedBlock{
				CustomType:         convert.NewCustomTypeNestedObject(nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				PlanModifiers:      convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:         convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.SingleNestedBlock{
				Description: pointer("description"),
			},
			expected: GeneratorSingleNestedBlock{
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				Description:   convert.NewDescription(pointer("description")),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.SingleNestedBlock{
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
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
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
		"plan-modifiers": {
			input: &resource.SingleNestedBlock{
				PlanModifiers: specschema.ObjectPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: GeneratorSingleNestedBlock{
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
		"plan-modifier-custom-import": {
			input: GeneratorSingleNestedBlock{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/myotherproject/myplanmodifiers/planmodifier",
							},
						},
					},
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/myplanmodifiers/planmodifier",
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
					Path: generatorschema.PlanModifierImport,
				},
				{
					Path: "github.com/myotherproject/myplanmodifiers/planmodifier",
				},
				{
					Path: "github.com/myproject/myplanmodifiers/planmodifier",
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
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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
						NestedAttributeObject: NewNestedAttributeObject(
							generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
							nil,
							convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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
								ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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
						NestedBlockObject: NewNestedBlockObject(
							generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
							generatorschema.GeneratorBlocks{},
							nil,
							convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
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
								ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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

		"plan-modifiers": {
			input: GeneratorSingleNestedBlock{
				CustomType: convert.NewCustomTypeNestedObject(nil, blockName),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"single_nested_block": schema.SingleNestedBlock{
CustomType: SingleNestedBlockType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedBlockValue{}.AttributeTypes(ctx),
},
},
PlanModifiers: []planmodifier.Object{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
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
