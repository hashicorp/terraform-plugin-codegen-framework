// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestGeneratorFloat64Attribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorFloat64Attribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorFloat64Attribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Required: true,
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Optional: true,
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Computed: true,
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Sensitive: true,
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					Description: "description",
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorFloat64Attribute{
				Float64Attribute: schema.Float64Attribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorFloat64Attribute{
				Validators: []specschema.Float64Validator{
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
					},
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					},
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Validators: []validator.Float64{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorFloat64Attribute{
				PlanModifiers: []specschema.Float64PlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							SchemaDefinition: "my_plan_modifier.Modify()",
						},
					},
					{
						Custom: &specschema.CustomPlanModifier{
							SchemaDefinition: "my_other_plan_modifier.Modify()",
						},
					},
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
PlanModifiers: []planmodifier.Float64{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorFloat64Attribute{
				Default: &specschema.Float64Default{
					Static: pointer(1.234),
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Default: float64default.StaticFloat64(1.234),
},`,
		},

		"default-custom": {
			input: GeneratorFloat64Attribute{
				Default: &specschema.Float64Default{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_float64_default.Default()",
					},
				},
			},
			expected: `
"float64_attribute": schema.Float64Attribute{
Default: my_float64_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("float64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
