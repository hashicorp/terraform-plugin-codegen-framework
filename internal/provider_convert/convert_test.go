// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func pointer[T any](in T) *T {
	return &in
}

func TestToGeneratorProviderSchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		spec           spec.Specification
		expectedSchema map[string]provider_generate.GeneratorProviderSchema
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
			expectedSchema: map[string]provider_generate.GeneratorProviderSchema{
				"example": {
					Attributes: map[string]provider_generate.GeneratorAttribute{
						"bool_attribute": provider_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional:  true,
								Sensitive: true,
							},
						},
						"list_attribute": provider_generate.GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
								ElementType: types.ListType{
									ElemType: types.StringType,
								},
							},
						},
						"list_nested_attribute": provider_generate.GeneratorListNestedAttribute{
							NestedObject: provider_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]provider_generate.GeneratorAttribute{
									"nested_bool_attribute": provider_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
									"nested_list_attribute": provider_generate.GeneratorListAttribute{
										ListAttribute: schema.ListAttribute{
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
							},
							ListNestedAttribute: schema.ListNestedAttribute{
								Optional: true,
							},
						},
						"object_attribute": provider_generate.GeneratorObjectAttribute{
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
						"single_nested_attribute": provider_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]provider_generate.GeneratorAttribute{
								"nested_bool_attribute": provider_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
								"nested_list_attribute": provider_generate.GeneratorListAttribute{
									ListAttribute: schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
							SingleNestedAttribute: schema.SingleNestedAttribute{
								Optional: true,
							},
						},
					},
					Blocks: map[string]provider_generate.GeneratorBlock{
						"list_nested_block": provider_generate.GeneratorListNestedBlock{
							NestedObject: provider_generate.GeneratorNestedBlockObject{
								Attributes: map[string]provider_generate.GeneratorAttribute{
									"nested_bool_attribute": provider_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
						"single_nested_block": provider_generate.GeneratorSingleNestedBlock{
							Attributes: map[string]provider_generate.GeneratorAttribute{
								"nested_bool_attribute": provider_generate.GeneratorBoolAttribute{
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

			got, err := c.ToGeneratorProviderSchema()

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
