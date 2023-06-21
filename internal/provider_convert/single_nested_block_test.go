// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
)

func TestConvertSingleNestedBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.SingleNestedBlock
		expected      provider_generate.GeneratorSingleNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.SingleNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type is not defined: %+v", provider.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "bool_attribute",
						Bool: &provider.BoolAttribute{
							OptionalRequired: "optional",
						},
					},
				},
			},
			expected: provider_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]provider_generate.GeneratorAttribute{
					"bool_attribute": provider_generate.GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &provider.SingleNestedBlock{
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
			expected: provider_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]provider_generate.GeneratorAttribute{
					"list_attribute": provider_generate.GeneratorListAttribute{
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
		"attributes-list-nested-bool": {
			input: &provider.SingleNestedBlock{
				Attributes: []provider.Attribute{
					{
						Name: "nested_attribute",
						ListNested: &provider.ListNestedAttribute{
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
			expected: provider_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]provider_generate.GeneratorAttribute{
					"nested_attribute": provider_generate.GeneratorListNestedAttribute{
						NestedObject: provider_generate.GeneratorNestedAttributeObject{
							Attributes: map[string]provider_generate.GeneratorAttribute{
								"nested_bool": provider_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
						ListNestedAttribute: schema.ListNestedAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &provider.SingleNestedBlock{
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
			expected: provider_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]provider_generate.GeneratorAttribute{
					"object_attribute": provider_generate.GeneratorObjectAttribute{
						ObjectAttribute: schema.ObjectAttribute{
							Optional: true,
						},
						AttributeTypes: []specschema.ObjectAttributeType{
							{
								Name: "obj_bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &provider.SingleNestedBlock{
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
			expected: provider_generate.GeneratorSingleNestedBlock{
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
		"blocks-nil": {
			input: &provider.SingleNestedBlock{
				Blocks: []provider.Block{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("block type is not defined: %+v", provider.Block{
				Name: "empty",
			}),
		},
		"blocks-list-nested-bool": {
			input: &provider.SingleNestedBlock{
				Blocks: []provider.Block{
					{
						Name: "nested_block",
						ListNested: &provider.ListNestedBlock{
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
			expected: provider_generate.GeneratorSingleNestedBlock{
				Blocks: map[string]provider_generate.GeneratorBlock{
					"nested_block": provider_generate.GeneratorListNestedBlock{
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
		"blocks-single-nested-bool": {
			input: &provider.SingleNestedBlock{
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
			expected: provider_generate.GeneratorSingleNestedBlock{
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
		"custom_type": {
			input: &provider.SingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &provider.SingleNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.SingleNestedBlock{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"validators": {
			input: &provider.SingleNestedBlock{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: provider_generate.GeneratorSingleNestedBlock{
				Validators: []specschema.ObjectValidator{
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

			got, err := convertSingleNestedBlock(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
