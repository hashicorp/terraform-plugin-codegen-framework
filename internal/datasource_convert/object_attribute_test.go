// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

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
		"attribute-type-bool": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
		},
		"attribute-type-string": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
		},
		"attribute-type-list-string": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
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
				AttributeTypes: specschema.ObjectAttributeTypes{
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
			input: &datasource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
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
				AttributeTypes: specschema.ObjectAttributeTypes{
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
			input: &datasource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name:   "obj_str",
											String: &specschema.StringType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name:   "obj_str",
											String: &specschema.StringType{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"attribute-type-object-string": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
		},
		"attribute-type-object-list-string": {
			input: &datasource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
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
				AttributeTypes: specschema.ObjectAttributeTypes{
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
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorObjectAttribute{
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
				Validators: specschema.ObjectValidators{
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
			expected: datasource_generate.GeneratorObjectAttribute{
				Validators: specschema.ObjectValidators{
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
