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

func TestGeneratorStringAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorStringAttribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorStringAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Required: true,
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Optional: true,
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Computed: true,
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Sensitive: true,
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Description: "description",
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorStringAttribute{
				Validators: specschema.StringValidators{
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
"string_attribute": schema.StringAttribute{
Validators: []validator.String{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorStringAttribute{
				PlanModifiers: specschema.StringPlanModifiers{
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
"string_attribute": schema.StringAttribute{
PlanModifiers: []planmodifier.String{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-static": {
			input: GeneratorStringAttribute{
				Default: &specschema.StringDefault{
					Static: pointer("default string"),
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
Default: stringdefault.StaticString("default string"),
},`,
		},

		"default-custom": {
			input: GeneratorStringAttribute{
				Default: &specschema.StringDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_string_default.Default()",
					},
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
Default: my_string_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("string_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorStringAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorStringAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "types.String",
				TfsdkName: "string_attribute",
			},
		},
		"custom-type": {
			input: GeneratorStringAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "string_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("string_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
