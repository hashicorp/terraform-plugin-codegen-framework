package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TestGeneratorInt64Attribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorInt64Attribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorInt64Attribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorInt64Attribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorInt64Attribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorInt64Attribute{
				Validators: []specschema.Int64Validator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorInt64Attribute{
				Validators: []specschema.Int64Validator{
					{
						Custom: &specschema.CustomValidator{
							Import: nil,
						},
					},
				}},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorInt64Attribute{
				Validators: []specschema.Int64Validator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer(""),
						},
					},
				}},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"validator-custom-import": {
			input: GeneratorInt64Attribute{
				Validators: []specschema.Int64Validator{
					{
						Custom: &specschema.CustomValidator{
							Import: pointer("github.com/myotherproject/myvalidators/validator"),
						},
					},
					{
						Custom: &specschema.CustomValidator{
							Import: pointer("github.com/myproject/myvalidators/validator"),
						},
					},
				}},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
				validatorImport:        {},
				"github.com/myotherproject/myvalidators/validator": {},
				"github.com/myproject/myvalidators/validator":      {},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

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
				Validators: []specschema.Int64Validator{
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

var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return x.Error() == y.Error()
})
