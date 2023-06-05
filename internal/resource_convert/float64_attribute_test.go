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

func TestConvertFloat64Attribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.Float64Attribute
		expected      resource_generate.GeneratorFloat64Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.Float64Attribute is nil"),
		},
		"computed": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "computed",
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "optional",
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &resource.Float64Attribute{
				ComputedOptionalRequired: "required",
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &resource.Float64Attribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.Float64Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.Float64Attribute{
				Description: pointer("description"),
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &resource.Float64Attribute{
				Sensitive: pointer(true),
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &resource.Float64Attribute{
				Validators: []specschema.Float64Validator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Validators: []specschema.Float64Validator{
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
			input: &resource.Float64Attribute{
				PlanModifiers: []specschema.Float64PlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				PlanModifiers: []specschema.Float64PlanModifier{
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
			input: &resource.Float64Attribute{
				Default: &specschema.Float64Default{
					Custom: &specschema.CustomDefault{
						Import:           pointer("github.com/.../my_default"),
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(1.234),
				},
			},
			expected: resource_generate.GeneratorFloat64Attribute{
				Default: &specschema.Float64Default{
					Custom: &specschema.CustomDefault{
						Import:           pointer("github.com/.../my_default"),
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer(1.234),
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertFloat64Attribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
