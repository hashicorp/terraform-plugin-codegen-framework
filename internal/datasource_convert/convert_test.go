// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func pointer[T any](in T) *T {
	return &in
}

func TestToGeneratorDataSourceSchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		spec           spec.Specification
		expectedSchema map[string]datasource_generate.GeneratorDataSourceSchema
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
			expectedSchema: map[string]datasource_generate.GeneratorDataSourceSchema{
				"example": {
					Attributes: map[string]datasource_generate.GeneratorAttribute{
						"bool_attribute": datasource_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional:  true,
								Sensitive: true,
							},
						},
						"list_attribute": datasource_generate.GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Computed: true,
								ElementType: types.ListType{
									ElemType: types.StringType,
								},
							},
						},
						"list_nested_attribute": datasource_generate.GeneratorListNestedAttribute{
							NestedObject: datasource_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]datasource_generate.GeneratorAttribute{
									"nested_bool_attribute": datasource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
									"nested_list_attribute": datasource_generate.GeneratorListAttribute{
										ListAttribute: schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
									},
								},
							},
							ListNestedAttribute: schema.ListNestedAttribute{
								Optional: true,
							},
						},
						"object_attribute": datasource_generate.GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								AttributeTypes: map[string]attr.Type{
									"obj_bool": basetypes.BoolType{},
									"obj_list": basetypes.ListType{
										ElemType: types.StringType,
									},
								},
								Optional: true,
							},
						},
						"single_nested_attribute": datasource_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]datasource_generate.GeneratorAttribute{
								"nested_bool_attribute": datasource_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
								"nested_list_attribute": datasource_generate.GeneratorListAttribute{
									ListAttribute: schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
							},
							SingleNestedAttribute: schema.SingleNestedAttribute{
								Optional: true,
							},
						},
					},
					Blocks: map[string]datasource_generate.GeneratorBlock{
						"list_nested_block": datasource_generate.GeneratorListNestedBlock{
							NestedObject: datasource_generate.GeneratorNestedBlockObject{
								Attributes: map[string]datasource_generate.GeneratorAttribute{
									"nested_bool_attribute": datasource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
						"single_nested_block": datasource_generate.GeneratorSingleNestedBlock{
							Attributes: map[string]datasource_generate.GeneratorAttribute{
								"nested_bool_attribute": datasource_generate.GeneratorBoolAttribute{
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

			got, err := c.ToGeneratorDataSourceSchema()

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
