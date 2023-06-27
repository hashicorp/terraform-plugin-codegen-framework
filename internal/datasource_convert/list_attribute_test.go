// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func TestConvertListAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.ListAttribute
		expected      datasource_generate.GeneratorListAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.ListAttribute is nil"),
		},
		"element-type-bool": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
		},
		"element-type-string": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"element-type-list-string": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
		},
		"element-type-map-string": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
		},
		"element-type-list-object-string": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
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
			},
			expected: datasource_generate.GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
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
			},
		},
		"element-type-object-string": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
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
		"element-type-object-list-string": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
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
			},
			expected: datasource_generate.GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: []specschema.ObjectAttributeType{
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
			},
		},
		"computed": {
			input: &datasource.ListAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Computed: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"computed_optional": {
			input: &datasource.ListAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Computed: true,
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"optional": {
			input: &datasource.ListAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"required": {
			input: &datasource.ListAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Required: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"custom_type": {
			input: &datasource.ListAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"deprecation_message": {
			input: &datasource.ListAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					DeprecationMessage: "deprecation message",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"description": {
			input: &datasource.ListAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"sensitive": {
			input: &datasource.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					Sensitive: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"validators": {
			input: &datasource.ListAttribute{
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
			expected: datasource_generate.GeneratorListAttribute{
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
