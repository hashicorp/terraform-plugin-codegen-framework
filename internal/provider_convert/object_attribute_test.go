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

func TestConvertObjectAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.ObjectAttribute
		expected      provider_generate.GeneratorObjectAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.ObjectAttribute is nil"),
		},
		"attribute-type-bool": {
			input: &provider.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
		},
		"attribute-type-string": {
			input: &provider.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
		},
		"attribute-type-list-string": {
			input: &provider.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
		},
		"attribute-type-map-string": {
			input: &provider.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_map_string",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_map_string",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
		},
		"attribute-type-list-object-string": {
			input: &provider.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: []specschema.ObjectAttributeType{
									{
										Name:   "obj_string",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: []specschema.ObjectAttributeType{
									{
										Name:   "obj_string",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
		},
		"attribute-type-object-string": {
			input: &provider.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
		},
		"attribute-type-object-list-string": {
			input: &provider.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
		},
		"optional": {
			input: &provider.ObjectAttribute{
				OptionalRequired: "optional",
			},
			expected: provider_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &provider.ObjectAttribute{
				OptionalRequired: "required",
			},
			expected: provider_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &provider.ObjectAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &provider.ObjectAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: provider_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &provider.ObjectAttribute{
				Description: pointer("description"),
			},
			expected: provider_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &provider.ObjectAttribute{
				Sensitive: pointer(true),
			},
			expected: provider_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &provider.ObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: provider_generate.GeneratorObjectAttribute{
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

			got, err := convertObjectAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
