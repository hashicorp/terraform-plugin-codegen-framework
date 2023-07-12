// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func pointer[T any](in T) *T {
	return &in
}

func TestToGeneratorProviderSchema(t *testing.T) {
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
										AttributeTypes: []specschema.ObjectAttributeType{
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
					Attributes: map[string]generatorschema.GeneratorAttribute{
						"bool_attribute": resource_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional:  true,
								Sensitive: true,
							},
						},
						"list_attribute": resource_generate.GeneratorListAttribute{
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
						"map_attribute": resource_generate.GeneratorMapAttribute{
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
						"set_attribute": resource_generate.GeneratorSetAttribute{
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
						"list_nested_attribute": resource_generate.GeneratorListNestedAttribute{
							NestedObject: resource_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]generatorschema.GeneratorAttribute{
									"nested_bool_attribute": resource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
									"nested_list_attribute": resource_generate.GeneratorListAttribute{
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
						"object_attribute": resource_generate.GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
							AttributeTypes: []specschema.ObjectAttributeType{
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
						"single_nested_attribute": resource_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]generatorschema.GeneratorAttribute{
								"nested_bool_attribute": resource_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
								"nested_list_attribute": resource_generate.GeneratorListAttribute{
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
					Blocks: map[string]generatorschema.GeneratorBlock{
						"list_nested_block": resource_generate.GeneratorListNestedBlock{
							NestedObject: resource_generate.GeneratorNestedBlockObject{
								Attributes: map[string]generatorschema.GeneratorAttribute{
									"nested_bool_attribute": resource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
						"single_nested_block": resource_generate.GeneratorSingleNestedBlock{
							Attributes: map[string]generatorschema.GeneratorAttribute{
								"nested_bool_attribute": resource_generate.GeneratorBoolAttribute{
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

			c := NewConverter(testCase.spec)

			got, err := c.ToGeneratorResourceSchema()

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expectedSchema); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return x.Error() == y.Error()
})
