// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestConvertSetNestedAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.SetNestedAttribute
		expected      provider_generate.GeneratorSetNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.SetNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &provider.SetNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
					Attributes: []provider.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", provider.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &provider.SetNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: provider_generate.GeneratorSetNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
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
			input: &provider.SetNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: provider_generate.GeneratorSetNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
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
		},
		"attributes-list-nested-bool": {
			input: &provider.SetNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: provider_generate.GeneratorSetNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": provider_generate.GeneratorSetNestedAttribute{
							NestedObject: provider_generate.GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
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
			input: &provider.SetNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: provider_generate.GeneratorSetNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
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
		},
		"attributes-single-nested-bool": {
			input: &provider.SetNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			expected: provider_generate.GeneratorSetNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": provider_generate.GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
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
		"computed": {
			input: &provider.SetNestedAttribute{
				OptionalRequired: "computed",
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Optional: true,
				},
			},
		},
		"computed_optional": {
			input: &provider.SetNestedAttribute{
				OptionalRequired: "computed_optional",
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Optional: true,
				},
			},
		},
		"optional": {
			input: &provider.SetNestedAttribute{
				OptionalRequired: "optional",
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &provider.SetNestedAttribute{
				OptionalRequired: "required",
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &provider.SetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
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
			input: &provider.SetNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.SetNestedAttribute{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &provider.SetNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: provider_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &provider.SetNestedAttribute{
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
			expected: provider_generate.GeneratorSetNestedAttribute{
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

			got, err := convertSetNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
