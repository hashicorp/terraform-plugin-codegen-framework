// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorStringAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.StringAttribute
		expected      GeneratorStringAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.StringAttribute is nil"),
		},
		"computed": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &resource.StringAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &resource.StringAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{},
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
			input: &resource.StringAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.StringAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &resource.StringAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorStringAttribute{
				StringAttribute: schema.StringAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &resource.StringAttribute{
				Validators: specschema.StringValidators{
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
			expected: GeneratorStringAttribute{
				Validators: specschema.StringValidators{
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
		"plan-modifiers": {
			input: &resource.StringAttribute{
				PlanModifiers: specschema.StringPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: GeneratorStringAttribute{
				PlanModifiers: specschema.StringPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
		},
		"default": {
			input: &resource.StringAttribute{
				Default: &specschema.StringDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer("default string"),
				},
			},
			expected: GeneratorStringAttribute{
				Default: &specschema.StringDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
					Static: pointer("default string"),
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorStringAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorStringAttribute_Schema(t *testing.T) {
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

		"associated-external-type": {
			input: GeneratorStringAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
				},
			},
			expected: `
"string_attribute": schema.StringAttribute{
CustomType: StringAttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorStringAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
				},
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

			got, err := testCase.input.Schema("string_attribute")

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
		"associated-external-type": {
			input: GeneratorStringAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
				},
			},
			expected: model.Field{
				Name:      "StringAttribute",
				ValueType: "StringAttributeValue",
				TfsdkName: "string_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorStringAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.StringAttribute",
					},
				},
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
