// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

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
							BoolAttribute: schema.BoolAttribute{
								Optional:  true,
								Sensitive: true,
							},
						},
						"list_attribute": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
						"map_attribute": GeneratorMapAttribute{
							MapAttribute: schema.MapAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								Map: &specschema.MapType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
						"set_attribute": GeneratorSetAttribute{
							SetAttribute: schema.SetAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								Set: &specschema.SetType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
						"list_nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
									"nested_list_attribute": GeneratorListAttribute{
										ListAttribute: schema.ListAttribute{
											Optional: true,
										},
										ElementType: specschema.ElementType{
											String: &specschema.StringType{},
										},
									},
								},
							},
							ListNestedAttribute: schema.ListNestedAttribute{
								Optional: true,
							},
						},
						"object_attribute": GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
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
						},
						"single_nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool_attribute": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
								"nested_list_attribute": GeneratorListAttribute{
									ListAttribute: schema.ListAttribute{
										Optional: true,
									},
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
							SingleNestedAttribute: schema.SingleNestedAttribute{
								Optional: true,
							},
						},
					},
					Blocks: generatorschema.GeneratorBlocks{
						"list_nested_block": GeneratorListNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool_attribute": GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
						"single_nested_block": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool_attribute": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
