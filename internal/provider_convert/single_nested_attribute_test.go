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

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
)

func TestConvertSingleNestedAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.SingleNestedAttribute
		expected      provider_generate.GeneratorSingleNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.SingleNestedAttribute is nil"),
		},
		"attributes-nil": {
			input: &provider.SingleNestedAttribute{
				Attributes: []provider.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", provider.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &provider.SingleNestedAttribute{
				Attributes: []provider.Attribute{
					{
						Name: "bool_attribute",
						Bool: &provider.BoolAttribute{
							OptionalRequired: "optional",
						},
					},
				},
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
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
			input: &provider.SingleNestedAttribute{
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
			expected: provider_generate.GeneratorSingleNestedAttribute{
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
		"attributes-list-nested-bool": {
			input: &provider.SingleNestedAttribute{
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
			expected: provider_generate.GeneratorSingleNestedAttribute{
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
			input: &provider.SingleNestedAttribute{
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
			expected: provider_generate.GeneratorSingleNestedAttribute{
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
		"attributes-single-nested-bool": {
			input: &provider.SingleNestedAttribute{
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
			expected: provider_generate.GeneratorSingleNestedAttribute{
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
		"optional": {
			input: &provider.SingleNestedAttribute{
				OptionalRequired: "optional",
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &provider.SingleNestedAttribute{
				OptionalRequired: "required",
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &provider.SingleNestedAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &provider.SingleNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.SingleNestedAttribute{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &provider.SingleNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &provider.SingleNestedAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: provider_generate.GeneratorSingleNestedAttribute{
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

			got, err := convertSingleNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
