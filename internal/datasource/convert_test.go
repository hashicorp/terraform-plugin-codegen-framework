// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

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
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: "optional",
										Sensitive:                pointer(true),
									},
								},
								{
									Name: "list_attribute",
									List: &datasource.ListAttribute{
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
									Map: &datasource.MapAttribute{
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
									Set: &datasource.SetAttribute{
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
									ListNested: &datasource.ListNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: "optional",
													},
												},
												{
													Name: "nested_list_attribute",
													List: &datasource.ListAttribute{
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
									Object: &datasource.ObjectAttribute{
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
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "nested_bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: "optional",
												},
											},
											{
												Name: "nested_list_attribute",
												List: &datasource.ListAttribute{
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
							Blocks: []datasource.Block{
								{
									Name: "list_nested_block",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: "optional",
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block",
									SingleNested: &datasource.SingleNestedBlock{
										Attributes: []datasource.Attribute{
											{
												Name: "nested_bool_attribute",
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
			},
			expectedSchema: map[string]generatorschema.GeneratorSchema{
				"example": {
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							Sensitive:                convert.NewSensitive(pointer(true)),
						},
						"list_attribute": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Computed: true,
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
								Computed: true,
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
								Computed: true,
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
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									},
									"nested_list_attribute": GeneratorListAttribute{
										ListAttribute: schema.ListAttribute{
											Computed: true,
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
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
								"nested_list_attribute": GeneratorListAttribute{
									ListAttribute: schema.ListAttribute{
										Computed: true,
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
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									},
								},
							},
						},
						"single_nested_block": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool_attribute": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
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
