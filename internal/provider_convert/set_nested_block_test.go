// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func TestConvertSetNestedBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.SetNestedBlock
		expected      provider_generate.GeneratorSetNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.SetNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Attributes: []provider.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type is not defined: %+v", provider.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Attributes: []provider.Attribute{
						{
							Name: "bool_attribute",
							Bool: &provider.BoolAttribute{
								OptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				NestedObject: provider_generate.GeneratorNestedBlockObject{
					Attributes: map[string]provider_generate.GeneratorAttribute{
						"bool_attribute": provider_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Attributes: []provider.Attribute{
						{
							Name: "list_attribute",
							List: &provider.ListAttribute{
								OptionalRequired: "optional",
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				NestedObject: provider_generate.GeneratorNestedBlockObject{
					Attributes: map[string]provider_generate.GeneratorAttribute{
						"list_attribute": provider_generate.GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								ElementType: types.BoolType,
								Optional:    true,
							},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Attributes: []provider.Attribute{
						{
							Name: "nested_attribute",
							SetNested: &provider.SetNestedAttribute{
								NestedObject: provider.NestedAttributeObject{
									Attributes: []provider.Attribute{
										{
											Name: "nested_bool",
											Bool: &provider.BoolAttribute{
												OptionalRequired: "optional",
											},
										},
									},
								},
								OptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				NestedObject: provider_generate.GeneratorNestedBlockObject{
					Attributes: map[string]provider_generate.GeneratorAttribute{
						"nested_attribute": provider_generate.GeneratorSetNestedAttribute{
							NestedObject: provider_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]provider_generate.GeneratorAttribute{
									"nested_bool": provider_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
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
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Attributes: []provider.Attribute{
						{
							Name: "object_attribute",
							Object: &provider.ObjectAttribute{
								AttributeTypes: []specschema.ObjectAttributeType{
									{
										Name: "obj_bool",
										Bool: &specschema.BoolType{},
									},
								},
								OptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				NestedObject: provider_generate.GeneratorNestedBlockObject{
					Attributes: map[string]provider_generate.GeneratorAttribute{
						"object_attribute": provider_generate.GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								AttributeTypes: map[string]attr.Type{
									"obj_bool": types.BoolType,
								},
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Attributes: []provider.Attribute{
						{
							Name: "nested_attribute",
							SingleNested: &provider.SingleNestedAttribute{
								Attributes: []provider.Attribute{
									{
										Name: "nested_bool",
										Bool: &provider.BoolAttribute{
											OptionalRequired: "optional",
										},
									},
								},
								OptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				NestedObject: provider_generate.GeneratorNestedBlockObject{
					Attributes: map[string]provider_generate.GeneratorAttribute{
						"nested_attribute": provider_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]provider_generate.GeneratorAttribute{
								"nested_bool": provider_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
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
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Blocks: []provider.Block{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("block type is not defined: %+v", provider.Block{
				Name: "empty",
			}),
		},

		"blocks-list-nested-bool": {
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Blocks: []provider.Block{
						{
							Name: "nested_block",
							SetNested: &provider.SetNestedBlock{
								NestedObject: provider.NestedBlockObject{
									Attributes: []provider.Attribute{
										{
											Name: "nested_bool",
											Bool: &provider.BoolAttribute{
												OptionalRequired: "computed",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				NestedObject: provider_generate.GeneratorNestedBlockObject{
					Blocks: map[string]provider_generate.GeneratorBlock{
						"nested_block": provider_generate.GeneratorSetNestedBlock{
							NestedObject: provider_generate.GeneratorNestedBlockObject{
								Attributes: map[string]provider_generate.GeneratorAttribute{
									"bool_attribute": provider_generate.GeneratorBoolAttribute{
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
			input: &provider.SetNestedBlock{
				NestedObject: provider.NestedBlockObject{
					Blocks: []provider.Block{
						{
							Name: "nested_block",
							SingleNested: &provider.SingleNestedBlock{
								Attributes: []provider.Attribute{
									{
										Name: "nested_bool",
										Bool: &provider.BoolAttribute{
											OptionalRequired: "computed",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				NestedObject: provider_generate.GeneratorNestedBlockObject{
					Blocks: map[string]provider_generate.GeneratorBlock{
						"nested_block": provider_generate.GeneratorSingleNestedBlock{
							Attributes: map[string]provider_generate.GeneratorAttribute{
								"bool_attribute": provider_generate.GeneratorBoolAttribute{
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
			input: &provider.SetNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &provider.SetNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.SetNestedBlock{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"validators": {
			input: &provider.SetNestedBlock{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: provider_generate.GeneratorSetNestedBlock{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
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
