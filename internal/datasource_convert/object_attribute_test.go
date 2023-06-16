// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func TestConvertObjectAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.ObjectAttribute
		expected      datasource_generate.GeneratorObjectAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.ObjectAttribute is nil"),
		},
		"attribute-type-nil": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", specschema.ObjectAttributeType{
				Name: "empty",
			}),
		},
		"attribute-type-bool": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_bool": types.BoolType,
					},
				},
			},
		},
		"attribute-type-string": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_string": types.StringType,
					},
				},
			},
		},
		"attribute-type-list-string": {
			input: &datasource.ObjectAttribute{
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
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_list_string": types.ListType{
							ElemType: types.StringType,
						},
					},
				},
			},
		},
		"attribute-type-map-string": {
			input: &datasource.ObjectAttribute{
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
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_map_string": types.MapType{
							ElemType: types.StringType,
						},
					},
				},
			},
		},
		"attribute-type-list-object-string": {
			input: &datasource.ObjectAttribute{
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
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_list_object_string": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"obj_string": types.StringType,
								},
							},
						},
					},
				},
			},
		},
		"attribute-type-object-string": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_string": types.StringType,
					},
				},
			},
		},
		"attribute-type-object-list-string": {
			input: &datasource.ObjectAttribute{
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
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_list_string": types.ListType{
							ElemType: types.StringType,
						},
					},
				},
			},
		},
		"computed": {
			input: &datasource.ObjectAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.ObjectAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.ObjectAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.ObjectAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.ObjectAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.ObjectAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.ObjectAttribute{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.ObjectAttribute{
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.ObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
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
