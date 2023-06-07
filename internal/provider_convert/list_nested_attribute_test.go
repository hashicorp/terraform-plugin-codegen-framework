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

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func TestConvertListNestedAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.ListNestedAttribute
		expected      provider_generate.GeneratorListNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.ListNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &provider.ListNestedAttribute{
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
			input: &provider.ListNestedAttribute{
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
			expected: provider_generate.GeneratorListNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
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
			input: &provider.ListNestedAttribute{
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
			expected: provider_generate.GeneratorListNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
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
		},
		"attributes-list-nested-bool": {
			input: &provider.ListNestedAttribute{
				NestedObject: provider.NestedAttributeObject{
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
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
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
		},
		"attributes-object-bool": {
			input: &provider.ListNestedAttribute{
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
			expected: provider_generate.GeneratorListNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
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
		},
		"attributes-single-nested-bool": {
			input: &provider.ListNestedAttribute{
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
			expected: provider_generate.GeneratorListNestedAttribute{
				NestedObject: provider_generate.GeneratorNestedAttributeObject{
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
		"computed": {
			input: &provider.ListNestedAttribute{
				OptionalRequired: "optional",
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Optional: true,
				},
			},
		},
		"computed_optional": {
			input: &provider.ListNestedAttribute{
				OptionalRequired: "computed_optional",
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Optional: true,
				},
			},
		},
		"optional": {
			input: &provider.ListNestedAttribute{
				OptionalRequired: "optional",
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &provider.ListNestedAttribute{
				OptionalRequired: "required",
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &provider.ListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &provider.ListNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.ListNestedAttribute{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &provider.ListNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: provider_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &provider.ListNestedAttribute{
				Validators: []specschema.ListValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: provider_generate.GeneratorListNestedAttribute{
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertListNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
