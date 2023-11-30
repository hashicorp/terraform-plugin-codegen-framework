// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorNumberAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.NumberAttribute
		expected      GeneratorNumberAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.NumberAttribute is nil"),
		},
		"computed": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.NumberAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.NumberAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{},
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
			input: &datasource.NumberAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.NumberAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.NumberAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.NumberAttribute{
				Validators: specschema.NumberValidators{
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
			expected: GeneratorNumberAttribute{
				Validators: specschema.NumberValidators{
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

			got, err := NewGeneratorNumberAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorNumberAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorNumberAttribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorNumberAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorNumberAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.NumberAttribute",
					},
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
CustomType: NumberAttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorNumberAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.NumberAttribute",
					},
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Required: true,
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Optional: true,
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Computed: true,
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Sensitive: true,
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					Description: "description",
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorNumberAttribute{
				NumberAttribute: schema.NumberAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"number_attribute": schema.NumberAttribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorNumberAttribute{
				Validators: specschema.NumberValidators{
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
"number_attribute": schema.NumberAttribute{
Validators: []validator.Number{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("number_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorNumberAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorNumberAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "types.Number",
				TfsdkName: "number_attribute",
			},
		},
		"custom-type": {
			input: GeneratorNumberAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "number_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorNumberAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.NumberAttribute",
					},
				},
			},
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "NumberAttributeValue",
				TfsdkName: "number_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorNumberAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.NumberAttribute",
					},
				},
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "NumberAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "number_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("number_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
