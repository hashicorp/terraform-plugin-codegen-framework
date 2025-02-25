// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
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
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "bool_attribute",
								Bool: &provider.BoolAttribute{
									OptionalRequired: "optional",
									Sensitive:        pointer(true),
								},
							},
							{
								Name: "list_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: "optional",
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
								Map: &provider.MapAttribute{
									OptionalRequired: "optional",
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
								Set: &provider.SetAttribute{
									OptionalRequired: "optional",
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
								ListNested: &provider.ListNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: "optional",
												},
											},
											{
												Name: "nested_list_attribute",
												List: &provider.ListAttribute{
													OptionalRequired: "optional",
													ElementType: specschema.ElementType{
														String: &specschema.StringType{},
													},
												},
											},
										},
									},
									OptionalRequired: "optional",
								},
							},
							{
								Name: "object_attribute",
								Object: &provider.ObjectAttribute{
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
									OptionalRequired: "optional",
								},
							},
							{
								Name: "single_nested_attribute",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "nested_bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: "optional",
											},
										},
										{
											Name: "nested_list_attribute",
											List: &provider.ListAttribute{
												OptionalRequired: "optional",
												ElementType: specschema.ElementType{
													String: &specschema.StringType{},
												},
											},
										},
									},
									OptionalRequired: "optional",
								},
							},
						},
						Blocks: []provider.Block{
							{
								Name: "list_nested_block",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: "optional",
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_block",
								SingleNested: &provider.SingleNestedBlock{
									Attributes: []provider.Attribute{
										{
											Name: "nested_bool_attribute",
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
			expectedSchema: map[string]generatorschema.GeneratorSchema{
				"example": {
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							CustomType:       convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							Sensitive:        convert.NewSensitive(pointer(true)),
							Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
						"list_attribute": GeneratorListAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
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
							Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
						"map_attribute": GeneratorMapAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
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
							Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
						},
						"set_attribute": GeneratorSetAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
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
							Validators: convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
						},
						"list_nested_attribute": GeneratorListNestedAttribute{
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
										CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
									"nested_list_attribute": GeneratorListAttribute{
										OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
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
										Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
									},
								},
							},
							NestedAttributeObject: convert.NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
										CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
									"nested_list_attribute": GeneratorListAttribute{
										OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
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
										Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"list_nested_attribute",
							),
							Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
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
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							CustomType:       convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							Validators:       convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
						"single_nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool_attribute": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
									CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
									Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
								"nested_list_attribute": GeneratorListAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
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
									Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
								},
							},
							OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
							CustomType:       convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
							Validators:       convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					Blocks: generatorschema.GeneratorBlocks{
						"list_nested_block": GeneratorListNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
										CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							NestedBlockObject: convert.NewNestedBlockObject(
								generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
										CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
										Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								generatorschema.GeneratorBlocks{},
								nil,
								convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"list_nested_block",
							),
							Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
						"single_nested_block": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool_attribute": GeneratorBoolAttribute{
									OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
									CustomType:       convert.NewCustomTypePrimitive(nil, nil, "nested_bool_attribute"),
									Validators:       convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_block"),
							Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
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
