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

func TestConvertListAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.ListAttribute
		expected      provider_generate.GeneratorListAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.ListAttribute is nil"),
		},
		"element-type-nil": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{},
			},
			expectedError: fmt.Errorf("element type is not defined: %+v", specschema.ElementType{}),
		},
		"element-type-bool": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.BoolType,
				},
			},
		},
		"element-type-string": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
			},
		},
		"element-type-list-string": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ListType{
						ElemType: types.StringType,
					},
				},
			},
		},
		"element-type-map-string": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.MapType{
						ElemType: types.StringType,
					},
				},
			},
		},
		"element-type-list-object-string": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: []specschema.ObjectAttributeType{
								{
									Name:   "str",
									String: &specschema.StringType{},
								},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ListType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"str": types.StringType,
							},
						},
					},
				},
			},
		},
		"element-type-object-string": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name:   "str",
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"str": types.StringType,
						},
					},
				},
			},
		},
		"element-type-object-list-string": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name: "list",
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									String: &specschema.StringType{},
								},
							},
						},
					},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"list": types.ListType{
								ElemType: types.StringType,
							},
						},
					},
				},
			},
		},
		"optional": {
			input: &provider.ListAttribute{
				OptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Optional:    true,
					ElementType: types.StringType,
				},
			},
		},
		"required": {
			input: &provider.ListAttribute{
				OptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Required:    true,
					ElementType: types.StringType,
				},
			},
		},
		"custom_type": {
			input: &provider.ListAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &provider.ListAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					DeprecationMessage: "deprecation message",
					ElementType:        types.StringType,
				},
			},
		},
		"description": {
			input: &provider.ListAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Description:         "description",
					MarkdownDescription: "description",
					ElementType:         types.StringType,
				},
			},
		},
		"sensitive": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Sensitive:   true,
					ElementType: types.StringType,
				},
			},
		},
		"validators": {
			input: &provider.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: []specschema.ListValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: provider_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
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

			got, err := convertListAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			// TODO: This prevents misleading failure when ElementType for both got and expected are nil.
			// TODO: Could overwrite comparison using an option to cmp.Diff()?
			if got.ListAttribute.ElementType == nil && testCase.expected.ListAttribute.ElementType == nil {
				return
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
