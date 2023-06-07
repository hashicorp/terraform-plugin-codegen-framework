// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func TestConvertMapAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.MapAttribute
		expected      resource_generate.GeneratorMapAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.MapAttribute is nil"),
		},
		"element-type-bool": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
		},
		"element-type-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"element-type-list-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
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
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
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
			input: &resource.MapAttribute{
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
			expected: resource_generate.GeneratorMapAttribute{
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
		},
		"element-type-object-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name:   "str",
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
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
		"element-type-object-list-string": {
			input: &resource.MapAttribute{
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
			expected: resource_generate.GeneratorMapAttribute{
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
		},
		"computed": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Computed: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"computed_optional": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Computed: true,
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"optional": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"required": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Required: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"custom_type": {
			input: &resource.MapAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
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
			input: &resource.MapAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					DeprecationMessage: "deprecation message",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"description": {
			input: &resource.MapAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"sensitive": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: resource_generate.GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Sensitive: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"validators": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: []specschema.MapValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: []specschema.MapValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
		"plan-modifiers": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: []specschema.MapPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: []specschema.MapPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
		},
		"default": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.MapDefault{
					Custom: &specschema.CustomDefault{
						Import:           pointer("github.com/.../my_default"),
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
			expected: resource_generate.GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.MapDefault{
					Custom: &specschema.CustomDefault{
						Import:           pointer("github.com/.../my_default"),
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertMapAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			// TODO: This prevents misleading failure when ElementType for both got and expected are nil.
			// TODO: Could overwrite comparison using an option to cmp.Diff()?
			if got.MapAttribute.ElementType == nil && testCase.expected.MapAttribute.ElementType == nil {
				return
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
