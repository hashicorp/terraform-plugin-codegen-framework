// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func Test_NewSchemas(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		spec           spec.Specification
		expectedSchema map[string]generatorschema.GeneratorSchema
	}{
		"success": {
			spec: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: "optional",
										Sensitive:                pointer(true),
									},
								},
								{
									Name: "list_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: "computed",
										ElementType: specschema.ElementType{
											List: &specschema.ListType{
												ElementType: specschema.ElementType{
													String: &specschema.StringType{},
												},
											},
										},
									},
								},
								{
									Name: "map_attribute",
									Map: &resource.MapAttribute{
										ComputedOptionalRequired: "computed",
										ElementType: specschema.ElementType{
											Map: &specschema.MapType{
												ElementType: specschema.ElementType{
													String: &specschema.StringType{},
												},
											},
										},
									},
								},
								{
									Name: "set_attribute",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: "computed",
										ElementType: specschema.ElementType{
											Set: &specschema.SetType{
												ElementType: specschema.ElementType{
													String: &specschema.StringType{},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_attribute",
									ListNested: &resource.ListNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: "optional",
													},
												},
												{
													Name: "nested_list_attribute",
													List: &resource.ListAttribute{
														ComputedOptionalRequired: "computed",
														ElementType: specschema.ElementType{
															String: &specschema.StringType{},
														},
													},
												},
											},
										},
										ComputedOptionalRequired: "optional",
									},
								},
								{
									Name: "object_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: specschema.ObjectAttributeTypes{
											{
												Name: "obj_bool",
												Bool: &specschema.BoolType{},
											},
											{
												Name: "obj_list",
												List: &specschema.ListType{
													ElementType: specschema.ElementType{
														String: &specschema.StringType{},
													},
												},
											},
										},
										ComputedOptionalRequired: "optional",
									},
								},
								{
									Name: "single_nested_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "nested_bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: "optional",
												},
											},
											{
												Name: "nested_list_attribute",
												List: &resource.ListAttribute{
													ComputedOptionalRequired: "computed",
													ElementType: specschema.ElementType{
														String: &specschema.StringType{},
													},
												},
											},
										},
										ComputedOptionalRequired: "optional",
									},
								},
							},
							Blocks: []resource.Block{
								{
									Name: "list_nested_block",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: "optional",
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block",
									SingleNested: &resource.SingleNestedBlock{
										Attributes: []resource.Attribute{
											{
												Name: "nested_bool_attribute",
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
			},
			expectedSchema: map[string]generatorschema.GeneratorSchema{
				"example": {
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							Sensitive:                convert.NewSensitive(pointer(true)),
							PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
							Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
						"list_attribute": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
							CustomType: convert.NewCustomTypeCollection(
								nil,
								nil,
								convert.CustomCollectionTypeList,
								"types.ListType{\nElemType: types.StringType,\n}",
								"list_attribute",
							),
							ElementType: specschema.ElementType{
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							}),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
						"map_attribute": GeneratorMapAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
							CustomType: convert.NewCustomTypeCollection(
								nil,
								nil,
								convert.CustomCollectionTypeMap,
								"types.MapType{\nElemType: types.StringType,\n}",
								"map_attribute",
							),
							ElementType: specschema.ElementType{
								Map: &specschema.MapType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								Map: &specschema.MapType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							}),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
						},
						"set_attribute": GeneratorSetAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
							CustomType: convert.NewCustomTypeCollection(
								nil,
								nil,
								convert.CustomCollectionTypeSet,
								"types.SetType{\nElemType: types.StringType,\n}",
								"set_attribute",
							),
							ElementType: specschema.ElementType{
								Set: &specschema.SetType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								Set: &specschema.SetType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							}),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
						},
						"list_nested_attribute": GeneratorListNestedAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
									"nested_list_attribute": GeneratorListAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType: convert.NewCustomTypeCollection(
											nil,
											nil,
											convert.CustomCollectionTypeList,
											"types.StringType",
											"nested_list_attribute",
										),
										ElementType: specschema.ElementType{
											String: &specschema.StringType{},
										},
										ElementTypeCollection: convert.NewElementType(specschema.ElementType{
											String: &specschema.StringType{},
										}),
										PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
										Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
									},
								},
							},
							NestedAttributeObject: NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
									"nested_list_attribute": GeneratorListAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomType: convert.NewCustomTypeCollection(
											nil,
											nil,
											convert.CustomCollectionTypeList,
											"types.StringType",
											"nested_list_attribute",
										),
										ElementType: specschema.ElementType{
											String: &specschema.StringType{},
										},
										ElementTypeCollection: convert.NewElementType(specschema.ElementType{
											String: &specschema.StringType{},
										}),
										PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
										Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
								convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"list_nested_attribute",
							),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
						"object_attribute": GeneratorObjectAttribute{
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
								{
									Name: "obj_list",
									List: &specschema.ListType{
										ElementType: specschema.ElementType{
											String: &specschema.StringType{},
										},
									},
								},
							},
							AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
								{
									Name: "obj_list",
									List: &specschema.ListType{
										ElementType: specschema.ElementType{
											String: &specschema.StringType{},
										},
									},
								},
							}),
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
						"single_nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool_attribute": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
								"nested_list_attribute": GeneratorListAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType: convert.NewCustomTypeCollection(
										nil,
										nil,
										convert.CustomCollectionTypeList,
										"types.StringType",
										"nested_list_attribute",
									),
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
									ElementTypeCollection: convert.NewElementType(specschema.ElementType{
										String: &specschema.StringType{},
									}),
									PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
									Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomType:               convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
							PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					Blocks: generatorschema.GeneratorBlocks{
						"list_nested_block": GeneratorListNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							NestedBlockObject: NewNestedBlockObject(
								generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
										CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
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
						"single_nested_block": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool_attribute": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							CustomType:    convert.NewCustomTypeNestedObject(nil, "single_nested_block"),
							PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewSchemas(testCase.spec)

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expectedSchema); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
