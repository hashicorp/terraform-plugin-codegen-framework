package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratorSingleNestedBlock_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		singleNestedBlock GeneratorSingleNestedBlock
		expected          string
		expectedError     error
	}{
		"attribute-bool": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"bool": GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},`,
		},

		"attribute-list": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						ListAttribute: schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
},`,
		},

		"attribute-list-nested": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"nested_list_nested": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"bool": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},`,
		},

		"attribute-object": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"object": GeneratorObjectAttribute{
						ObjectAttribute: schema.ObjectAttribute{
							AttributeTypes: map[string]attr.Type{
								"str": types.StringType,
							},
							Optional: true,
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
},`,
		},

		"attribute-single-nested-bool": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Attributes: map[string]GeneratorAttribute{
					"nested_single_nested": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"bool": GeneratorBoolAttribute{
								BoolAttribute: schema.BoolAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"nested_single_nested": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},`,
		},

		"block-list-nested-bool": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Blocks: map[string]GeneratorBlock{
					"nested_list_nested": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: map[string]GeneratorAttribute{
								"bool": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Blocks: map[string]schema.Block{
"nested_list_nested": schema.ListNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},`,
		},

		"block-single-nested-bool": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Blocks: map[string]GeneratorBlock{
					"nested_single_nested": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
							"bool": GeneratorBoolAttribute{
								BoolAttribute: schema.BoolAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Blocks: map[string]schema.Block{
"nested_single_nested": schema.SingleNestedBlock{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},`,
		},

		"custom-type": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
CustomType: my_custom_type,
},`,
		},

		"description": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					Description: "description",
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				SingleNestedBlock: schema.SingleNestedBlock{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			singleNestedBlock: GeneratorSingleNestedBlock{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
					},
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					},
				},
			},
			expected: `
"single_nested_block": schema.SingleNestedBlock{
Validators: []validator.Bool{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.singleNestedBlock.ToString("single_nested_block")

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
