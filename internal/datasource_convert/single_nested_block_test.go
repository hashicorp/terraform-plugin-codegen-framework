package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func TestConvertSingleNestedBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.SingleNestedBlock
		expected      datasource_generate.GeneratorSingleNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.SingleNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type is not defined: %+v", datasource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "bool_attribute",
						Bool: &datasource.BoolAttribute{
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]datasource_generate.GeneratorAttribute{
					"bool_attribute": datasource_generate.GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &datasource.SingleNestedBlock{
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
			expected: datasource_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]datasource_generate.GeneratorAttribute{
					"list_attribute": datasource_generate.GeneratorListAttribute{
						ListAttribute: schema.ListAttribute{
							ElementType: types.BoolType,
							Optional:    true,
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "nested_attribute",
						ListNested: &datasource.ListNestedAttribute{
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
			expected: datasource_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]datasource_generate.GeneratorAttribute{
					"nested_attribute": datasource_generate.GeneratorListNestedAttribute{
						NestedObject: datasource_generate.GeneratorNestedAttributeObject{
							Attributes: map[string]datasource_generate.GeneratorAttribute{
								"nested_bool": datasource_generate.GeneratorBoolAttribute{
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
		"attributes-object-bool": {
			input: &datasource.SingleNestedBlock{
				Attributes: []datasource.Attribute{
					{
						Name: "object_attribute",
						Object: &datasource.ObjectAttribute{
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
			expected: datasource_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]datasource_generate.GeneratorAttribute{
					"object_attribute": datasource_generate.GeneratorObjectAttribute{
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
		"attributes-single-nested-bool": {
			input: &datasource.SingleNestedBlock{
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
			expected: datasource_generate.GeneratorSingleNestedBlock{
				Attributes: map[string]datasource_generate.GeneratorAttribute{
					"nested_attribute": datasource_generate.GeneratorSingleNestedAttribute{
						Attributes: map[string]datasource_generate.GeneratorAttribute{
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
		"blocks-nil": {
			input: &datasource.SingleNestedBlock{
				Blocks: []datasource.Block{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("block type is not defined: %+v", datasource.Block{
				Name: "empty",
			}),
		},
		"blocks-list-nested-bool": {
			input: &datasource.SingleNestedBlock{
				Blocks: []datasource.Block{
					{
						Name: "nested_block",
						ListNested: &datasource.ListNestedBlock{
							NestedObject: datasource.NestedBlockObject{
								Attributes: []datasource.Attribute{
									{
										Name: "nested_bool",
										Bool: &datasource.BoolAttribute{
											ComputedOptionalRequired: "computed",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSingleNestedBlock{
				Blocks: map[string]datasource_generate.GeneratorBlock{
					"nested_block": datasource_generate.GeneratorListNestedBlock{
						NestedObject: datasource_generate.GeneratorNestedBlockObject{
							Attributes: map[string]datasource_generate.GeneratorAttribute{
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
		"blocks-single-nested-bool": {
			input: &datasource.SingleNestedBlock{
				Blocks: []datasource.Block{
					{
						Name: "nested_block",
						SingleNested: &datasource.SingleNestedBlock{
							Attributes: []datasource.Attribute{
								{
									Name: "nested_bool",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: "computed",
									},
								},
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSingleNestedBlock{
				Blocks: map[string]datasource_generate.GeneratorBlock{
					"nested_block": datasource_generate.GeneratorSingleNestedBlock{
						Attributes: map[string]datasource_generate.GeneratorAttribute{
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
		"custom_type": {
			input: &datasource.SingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.SingleNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.SingleNestedBlock{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"validators": {
			input: &datasource.SingleNestedBlock{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSingleNestedBlock{
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
