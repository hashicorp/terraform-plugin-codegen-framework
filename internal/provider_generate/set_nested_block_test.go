// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

func TestGeneratorSetNestedBlock_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetNestedBlock
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: GeneratorAttributes{
						"bool": GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: GeneratorAttributes{
						"list": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: GeneratorAttributes{
						"nested_list_nested": GeneratorSetNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: GeneratorAttributes{
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
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: GeneratorAttributes{
						"object": GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name:   "str",
									String: &specschema.StringType{},
								},
							},
						},
					},
				},
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Attributes: GeneratorAttributes{
						"nested_single_nested": GeneratorSingleNestedAttribute{
							Attributes: GeneratorAttributes{
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
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"nested_single_nested": schema.SingleNestedAttribute{
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

		"block-list-nested-bool": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Blocks: GeneratorBlocks{
						"nested_list_nested": GeneratorSetNestedBlock{
							NestedObject: GeneratorNestedBlockObject{
								Attributes: GeneratorAttributes{
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
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Blocks: map[string]schema.Block{
"nested_list_nested": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},
},`,
		},

		"block-single-nested-bool": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					Blocks: GeneratorBlocks{
						"nested_single_nested": GeneratorSingleNestedBlock{
							Attributes: GeneratorAttributes{
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
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Blocks: map[string]schema.Block{
"nested_single_nested": schema.SingleNestedBlock{
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

		"custom-type": {
			input: GeneratorSetNestedBlock{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
},
CustomType: my_custom_type,
},`,
		},

		"description": {
			input: GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					Description: "description",
				},
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSetNestedBlock{
				Validators: []specschema.SetValidator{
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
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
},
Validators: []validator.Set{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
					CustomType: &specschema.CustomType{
						Type: "my_custom_type",
					},
				},
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorSetNestedBlock{
				NestedObject: GeneratorNestedBlockObject{
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
			},
			expected: `
"set_nested_block": schema.SetNestedBlock{
NestedObject: schema.NestedBlockObject{
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("set_nested_block")

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetNestedBlock_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetNestedBlock
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "SetNestedBlock",
				ValueType: "types.Set",
				TfsdkName: "set_nested_block",
			},
		},
		"custom-type": {
			input: GeneratorSetNestedBlock{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "SetNestedBlock",
				ValueType: "my_custom_value_type",
				TfsdkName: "set_nested_block",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("set_nested_block")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
