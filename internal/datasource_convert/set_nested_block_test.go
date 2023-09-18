// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestConvertSetNestedBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.SetNestedBlock
		expected      datasource_generate.GeneratorSetNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.SetNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Attributes: []datasource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type is not defined: %+v", datasource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Attributes: []datasource.Attribute{
						{
							Name: "bool_attribute",
							Bool: &datasource.BoolAttribute{
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSetNestedBlock{
				NestedObject: datasource_generate.GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": datasource_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Attributes: []datasource.Attribute{
						{
							Name: "list_attribute",
							List: &datasource.ListAttribute{
								ComputedOptionalRequired: "optional",
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSetNestedBlock{
				NestedObject: datasource_generate.GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list_attribute": datasource_generate.GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Attributes: []datasource.Attribute{
						{
							Name: "nested_attribute",
							SetNested: &datasource.SetNestedAttribute{
								NestedObject: datasource.NestedAttributeObject{
									Attributes: []datasource.Attribute{
										{
											Name: "nested_bool",
											Bool: &datasource.BoolAttribute{
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
			expected: datasource_generate.GeneratorSetNestedBlock{
				NestedObject: datasource_generate.GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": datasource_generate.GeneratorSetNestedAttribute{
							NestedObject: datasource_generate.GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": datasource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
							SetNestedAttribute: schema.SetNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Attributes: []datasource.Attribute{
						{
							Name: "object_attribute",
							Object: &datasource.ObjectAttribute{
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
			expected: datasource_generate.GeneratorSetNestedBlock{
				NestedObject: datasource_generate.GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"object_attribute": datasource_generate.GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Attributes: []datasource.Attribute{
						{
							Name: "nested_attribute",
							SingleNested: &datasource.SingleNestedAttribute{
								Attributes: []datasource.Attribute{
									{
										Name: "nested_bool",
										Bool: &datasource.BoolAttribute{
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
			expected: datasource_generate.GeneratorSetNestedBlock{
				NestedObject: datasource_generate.GeneratorNestedBlockObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": datasource_generate.GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": datasource_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Computed: true,
									},
								},
							},
							SingleNestedAttribute: schema.SingleNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},

		"blocks-nil": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Blocks: []datasource.Block{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("block type is not defined: %+v", datasource.Block{
				Name: "empty",
			}),
		},

		"blocks-list-nested-bool": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Blocks: []datasource.Block{
						{
							Name: "nested_block",
							SetNested: &datasource.SetNestedBlock{
								NestedObject: datasource.NestedBlockObject{
									Attributes: []datasource.Attribute{
										{
											Name: "bool_attribute",
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
			expected: datasource_generate.GeneratorSetNestedBlock{
				NestedObject: datasource_generate.GeneratorNestedBlockObject{
					Blocks: generatorschema.GeneratorBlocks{
						"nested_block": datasource_generate.GeneratorSetNestedBlock{
							NestedObject: datasource_generate.GeneratorNestedBlockObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool_attribute": datasource_generate.GeneratorBoolAttribute{
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
		},

		"blocks-single-nested-bool": {
			input: &datasource.SetNestedBlock{
				NestedObject: datasource.NestedBlockObject{
					Blocks: []datasource.Block{
						{
							Name: "nested_block",
							SingleNested: &datasource.SingleNestedBlock{
								Attributes: []datasource.Attribute{
									{
										Name: "bool_attribute",
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
			expected: datasource_generate.GeneratorSetNestedBlock{
				NestedObject: datasource_generate.GeneratorNestedBlockObject{
					Blocks: generatorschema.GeneratorBlocks{
						"nested_block": datasource_generate.GeneratorSingleNestedBlock{
							Attributes: generatorschema.GeneratorAttributes{
								"bool_attribute": datasource_generate.GeneratorBoolAttribute{
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

		"custom_type": {
			input: &datasource.SetNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorSetNestedBlock{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.SetNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.SetNestedBlock{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"validators": {
			input: &datasource.SetNestedBlock{
				Validators: specschema.SetValidators{
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
			expected: datasource_generate.GeneratorSetNestedBlock{
				Validators: specschema.SetValidators{
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
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertSetNestedBlock(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
