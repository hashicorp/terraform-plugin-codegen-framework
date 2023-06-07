// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func TestConvertListNestedBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.ListNestedBlock
		expected      resource_generate.GeneratorListNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.ListNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type is not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "bool_attribute",
							Bool: &resource.BoolAttribute{
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorListNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"bool_attribute": resource_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "list_attribute",
							List: &resource.ListAttribute{
								ComputedOptionalRequired: "optional",
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorListNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"list_attribute": resource_generate.GeneratorListAttribute{
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
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							ListNested: &resource.ListNestedAttribute{
								NestedObject: resource.NestedAttributeObject{
									Attributes: []resource.Attribute{
										{
											Name: "nested_bool",
											Bool: &resource.BoolAttribute{
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
			expected: resource_generate.GeneratorListNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorListNestedAttribute{
							NestedObject: resource_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]resource_generate.GeneratorAttribute{
									"nested_bool": resource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Computed: true,
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
		},
		"attributes-object-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "object_attribute",
							Object: &resource.ObjectAttribute{
								AttributeTypes: []specschema.ObjectAttributeType{
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
			expected: resource_generate.GeneratorListNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"object_attribute": resource_generate.GeneratorObjectAttribute{
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
		},
		"attributes-single-nested-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							SingleNested: &resource.SingleNestedAttribute{
								Attributes: []resource.Attribute{
									{
										Name: "nested_bool",
										Bool: &resource.BoolAttribute{
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
			expected: resource_generate.GeneratorListNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]resource_generate.GeneratorAttribute{
								"nested_bool": resource_generate.GeneratorBoolAttribute{
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
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Blocks: []resource.Block{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("block type is not defined: %+v", resource.Block{
				Name: "empty",
			}),
		},

		"blocks-list-nested-bool": {
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Blocks: []resource.Block{
						{
							Name: "nested_block",
							ListNested: &resource.ListNestedBlock{
								NestedObject: resource.NestedBlockObject{
									Attributes: []resource.Attribute{
										{
											Name: "nested_bool",
											Bool: &resource.BoolAttribute{
												ComputedOptionalRequired: "computed",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorListNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Blocks: map[string]resource_generate.GeneratorBlock{
						"nested_block": resource_generate.GeneratorListNestedBlock{
							NestedObject: resource_generate.GeneratorNestedBlockObject{
								Attributes: map[string]resource_generate.GeneratorAttribute{
									"bool_attribute": resource_generate.GeneratorBoolAttribute{
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
			input: &resource.ListNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Blocks: []resource.Block{
						{
							Name: "nested_block",
							SingleNested: &resource.SingleNestedBlock{
								Attributes: []resource.Attribute{
									{
										Name: "nested_bool",
										Bool: &resource.BoolAttribute{
											ComputedOptionalRequired: "computed",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorListNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Blocks: map[string]resource_generate.GeneratorBlock{
						"nested_block": resource_generate.GeneratorSingleNestedBlock{
							Attributes: map[string]resource_generate.GeneratorAttribute{
								"bool_attribute": resource_generate.GeneratorBoolAttribute{
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
			input: &resource.ListNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: resource_generate.GeneratorListNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.ListNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: resource_generate.GeneratorListNestedBlock{
				ListNestedBlock: schema.ListNestedBlock{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.ListNestedBlock{
				Description: pointer("description"),
			},
			expected: resource_generate.GeneratorListNestedBlock{
				ListNestedBlock: schema.ListNestedBlock{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"validators": {
			input: &resource.ListNestedBlock{
				Validators: []specschema.ListValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorListNestedBlock{
				Validators: []specschema.ListValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
		"plan-modifiers": {
			input: &resource.ListNestedBlock{
				PlanModifiers: []specschema.ListPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorListNestedBlock{
				PlanModifiers: []specschema.ListPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
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

			got, err := convertListNestedBlock(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
