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

func TestGeneratorListNestedBlock_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.ListNestedBlock
		expected      GeneratorListNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.ListNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.ListNestedBlock{
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
			expected: GeneratorListNestedBlock{
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
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
				},
				NestedBlockObject: NewNestedBlockObject(
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
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
									},
								},
							},
							NestedAttributeObject: NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
									},
								},
								nil,
								convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
								convert.NewValidators(convert.ValidatorTypeObject, nil),
								"nested_attribute",
							),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, nil),
							Validators:    convert.NewValidators(convert.ValidatorTypeList, nil),
						},
					},
				},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
									},
								},
							},
							NestedAttributeObject: NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
									},
								},
								nil,
								convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
								convert.NewValidators(convert.ValidatorTypeObject, nil),
								"nested_attribute",
							),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, nil),
							Validators:    convert.NewValidators(convert.ValidatorTypeList, nil),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
							CustomType:               convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: NewNestedBlockObject(
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
							PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"blocks-nil": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Blocks: []resource.Block{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("block type not defined: %+v", resource.Block{
				Name: "empty",
			}),
		},

		"blocks-list-nested-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
								nil,
								nil,
								convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
								convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_block",
							),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
				},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
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
								nil,
								nil,
								convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
								convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_block",
							),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"blocks-single-nested-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
			},
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
				},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
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
					}, nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},

		"custom_type": {
			input: &resource.ListNestedBlock{
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
				NestedObject: GeneratorNestedBlockObject{},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &resource.ListNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorListNestedBlock{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				NestedObject:       GeneratorNestedBlockObject{},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.ListNestedBlock{
				Description: pointer("description"),
			},
			expected: GeneratorListNestedBlock{
				Description:  convert.NewDescription(pointer("description")),
				NestedObject: GeneratorNestedBlockObject{},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.ListNestedBlock{
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
				NestedObject: GeneratorNestedBlockObject{},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
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
		"plan-modifiers": {
			input: &resource.ListNestedBlock{
				PlanModifiers: specschema.ListPlanModifiers{
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
			expected: GeneratorListNestedBlock{
				NestedObject: GeneratorNestedBlockObject{},
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{},
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/nested_object",
						},
					},
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
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
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
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
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
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
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
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
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
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
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
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
		"nested-object-plan-modifier-custom-import": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: NewNestedBlockObject(
					nil,
					nil,
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
						&specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/myotherproject/myvalidators/validator",
								},
							},
						},
						&specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/myproject/myvalidators/validator",
								},
							},
						},
					}),
					convert.NewValidators(convert.ValidatorTypeObject, nil),
					"",
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.PlanModifierImport,
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

	testCases := map[string]struct {
		input         GeneratorListNestedBlock
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorListNestedBlock{
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
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
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
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{
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
								"nested_list_nested"),
						},
					},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
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
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
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
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
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
					}, nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{
						"nested_single_nested": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
							CustomType:    convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					}, nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
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
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
						&specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
						&specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					}),
					"list_nested_block",
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

		"plan-modifiers": {
			input: GeneratorListNestedBlock{
				NestedBlockObject: NewNestedBlockObject(
					generatorschema.GeneratorAttributes{},
					generatorschema.GeneratorBlocks{},
					nil,
					convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"list_nested_block",
				),
				PlanModifiers: convert.NewPlanModifiers(
					convert.PlanModifierTypeList,
					specschema.CustomPlanModifiers{
						&specschema.CustomPlanModifier{
							SchemaDefinition: "my_plan_modifier.Modify()",
						},
						&specschema.CustomPlanModifier{
							SchemaDefinition: "my_other_plan_modifier.Modify()",
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
PlanModifiers: []planmodifier.List{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
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
