// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorResourceSchemas_ModelsBytes(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         map[string]generatorschema.GeneratorSchema
		expected      string
		expectedError error
	}{
		"all": {
			input: map[string]generatorschema.GeneratorSchema{
				"example": {
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{},
						"bool_attribute_custom": GeneratorBoolAttribute{
							CustomType: convert.NewCustomTypePrimitive(
								&schema.CustomType{
									ValueType: "my_bool_value_type",
								},
								nil,
								"",
							),
						},
						"float64_attribute": GeneratorFloat64Attribute{},
						"float64_attribute_custom": GeneratorFloat64Attribute{
							CustomType: convert.NewCustomTypePrimitive(
								&schema.CustomType{
									ValueType: "my_float64_value_type",
								},
								nil,
								"",
							),
						},
						"int64_attribute": GeneratorInt64Attribute{},
						"int64_attribute_custom": GeneratorInt64Attribute{
							CustomType: convert.NewCustomTypePrimitive(
								&schema.CustomType{
									ValueType: "my_int64_value_type",
								},
								nil,
								"",
							),
						},
						"list_attribute": GeneratorListAttribute{},
						"list_attribute_custom": GeneratorListAttribute{
							CustomType: convert.NewCustomTypeCollection(
								&schema.CustomType{
									ValueType: "my_list_value_type",
								},
								nil,
								convert.CustomCollectionTypeList,
								"",
								"",
							),
						},
						"list_nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"list_nested_attribute_custom": GeneratorListNestedAttribute{
							CustomType: convert.NewCustomTypeNestedCollection(&schema.CustomType{
								ValueType: "my_list_nested_value_type",
							}),
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"map_attribute": GeneratorMapAttribute{},
						"map_attribute_custom": GeneratorMapAttribute{
							CustomType: convert.NewCustomTypeCollection(
								&schema.CustomType{
									ValueType: "my_map_value_type",
								},
								nil,
								convert.CustomCollectionTypeList,
								"",
								"",
							),
						},
						"map_nested_attribute": GeneratorMapNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"map_nested_attribute_custom": GeneratorMapNestedAttribute{
							CustomType: convert.NewCustomTypeNestedCollection(&schema.CustomType{
								ValueType: "my_map_nested_value_type",
							}),
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"number_attribute": GeneratorNumberAttribute{},
						"number_attribute_custom": GeneratorNumberAttribute{
							CustomType: convert.NewCustomTypePrimitive(
								&schema.CustomType{
									ValueType: "my_number_value_type",
								},
								nil,
								"",
							),
						},
						"object_attribute": GeneratorObjectAttribute{},
						"object_attribute_custom": GeneratorObjectAttribute{
							CustomType: convert.NewCustomTypeObject(
								&schema.CustomType{
									ValueType: "my_object_value_type",
								},
								nil,
								"",
							),
						},
						"set_attribute": GeneratorSetAttribute{},
						"set_attribute_custom": GeneratorSetAttribute{
							CustomType: convert.NewCustomTypeCollection(
								&schema.CustomType{
									ValueType: "my_set_value_type",
								},
								nil,
								convert.CustomCollectionTypeList,
								"",
								"",
							),
						},
						"set_nested_attribute": GeneratorSetNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"set_nested_attribute_custom": GeneratorSetNestedAttribute{
							CustomType: convert.NewCustomTypeNestedCollection(&schema.CustomType{
								ValueType: "my_set_nested_value_type",
							}),
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"single_nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{},
							},
						},
						"single_nested_attribute_custom": GeneratorSingleNestedAttribute{
							CustomType: convert.NewCustomTypeNestedObject(
								&schema.CustomType{
									ValueType: "my_single_nested_value_type",
								},
								"",
							),
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{},
							},
						},
						"string_attribute": GeneratorStringAttribute{},
						"string_attribute_custom": GeneratorStringAttribute{
							CustomType: convert.NewCustomTypePrimitive(
								&schema.CustomType{
									ValueType: "my_string_value_type",
								},
								nil,
								"",
							),
						},
					},
					Blocks: generatorschema.GeneratorBlocks{
						"list_nested_block": GeneratorListNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"list_nested_block_custom": GeneratorListNestedBlock{
							CustomType: convert.NewCustomTypeNestedCollection(
								&schema.CustomType{
									ValueType: "my_list_nested_value_type",
								},
							),
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"set_nested_block": GeneratorSetNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"set_nested_block_custom": GeneratorSetNestedBlock{
							CustomType: convert.NewCustomTypeNestedCollection(
								&schema.CustomType{
									ValueType: "my_set_nested_value_type",
								},
							),
							NestedObject: GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": GeneratorBoolAttribute{},
								},
							},
						},
						"single_nested_block": GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{},
							},
						},
						"single_nested_block_custom": GeneratorSingleNestedBlock{
							CustomType: convert.NewCustomTypeNestedObject(
								&schema.CustomType{
									ValueType: "my_single_nested_value_type",
								},
								"",
							),
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: "testdata/model.txt",
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			g := generatorschema.NewGeneratorSchemas(testCase.input)
			got, err := g.Models()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			expectedBytes, err := os.ReadFile(testCase.expected)
			if err != nil {
				t.Errorf("unexpected error reading %s file:%s", testCase.expected, err)
			}

			if diff := cmp.Diff(got["example"], expectedBytes); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
