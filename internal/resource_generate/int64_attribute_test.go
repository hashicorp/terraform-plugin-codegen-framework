// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

func TestGeneratorInt64Attribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorInt64Attribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorInt64Attribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Required: true,
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Optional: true,
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Computed: true,
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Sensitive: true,
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					Description: "description",
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorInt64Attribute{
				Int64Attribute: schema.Int64Attribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorInt64Attribute{
				Validators: specschema.Int64Validators{
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
"int64_attribute": schema.Int64Attribute{
Validators: []validator.Int64{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorInt64Attribute{
				PlanModifiers: []specschema.Int64PlanModifier{
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
"int64_attribute": schema.Int64Attribute{
PlanModifiers: []planmodifier.Int64{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorInt64Attribute{
				Default: &specschema.Int64Default{
					Static: pointer(int64(1234)),
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
Default: int64default.StaticInt64(1234),
},`,
		},

		"default-custom": {
			input: GeneratorInt64Attribute{
				Default: &specschema.Int64Default{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_int64_default.Default()",
					},
				},
			},
			expected: `
"int64_attribute": schema.Int64Attribute{
Default: my_int64_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("int64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorInt64Attribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorInt64Attribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "types.Int64",
				TfsdkName: "int64_attribute",
			},
		},
		"custom-type": {
			input: GeneratorInt64Attribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "Int64Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "int64_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("int64_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
