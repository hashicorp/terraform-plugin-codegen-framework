package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratorListNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListNestedAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"custom-type-and-nested-object-custom-type-without-import": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer(""),
					},
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
				datasourceSchemaImport:                       {},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer("github.com/my_account/my_project/attribute"),
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport:                       {},
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"custom-type-with-import-with-nested-object-custom-type-with-import": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: pointer("github.com/my_account/my_project/nested_object"),
					},
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute":     {},
				"github.com/my_account/my_project/nested_object": {},
			},
		},
		"nested-list": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								ElementType: types.BoolType,
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
				typesImport:            {},
			},
		},
		"nested-list-with-custom-type": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/nested_list"),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport:                         {},
				"github.com/my_account/my_project/nested_list": {},
			},
		},
		"nested-object": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"obj": GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								AttributeTypes: map[string]attr.Type{
									"bool": types.BoolType,
								},
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
				attrImport:             {},
				typesImport:            {},
			},
		},
		"nested-object-with-custom-type": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"obj": GeneratorObjectAttribute{
							CustomType: &specschema.CustomType{
								Import: pointer("github.com/my_account/my_project/nested_object"),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport:                           {},
				"github.com/my_account/my_project/nested_object": {},
			},
		},
		"validator-custom-nil": {
			input: GeneratorListNestedAttribute{
				Validators: []specschema.ListValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorListNestedAttribute{
				Validators: []specschema.ListValidator{
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
			input: GeneratorListNestedAttribute{
				Validators: []specschema.ListValidator{
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
			input: GeneratorListNestedAttribute{
				Validators: []specschema.ListValidator{
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
		"nested-object-validator-custom-nil": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: nil,
						},
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"nested-object-validator-custom-import-nil": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: &specschema.CustomValidator{
								Import: nil,
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"nested-object-validator-custom-import-empty-string": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
						{
							Custom: &specschema.CustomValidator{
								Import: pointer(""),
							},
						},
					},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"nested-object-validator-custom-import": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
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
					},
				},
			},
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

func TestGeneratorListNestedAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"bool": GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"list": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								ElementType: types.StringType,
								Optional:    true,
							},
						},
					},
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"nested_list_nested": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: map[string]GeneratorAttribute{
									"bool": GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"object": GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								AttributeTypes: map[string]attr.Type{
									"str": types.StringType,
								},
								Optional: true,
							},
						},
					},
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: map[string]GeneratorAttribute{
						"nested_single_nested": GeneratorSingleNestedAttribute{
							Attributes: map[string]GeneratorAttribute{
								"bool": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_single_nested": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Required: true,
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Optional: true,
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Computed: true,
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Sensitive: true,
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Description: "description",
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListNestedAttribute{
				Validators: []specschema.ListValidator{
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
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
},
Validators: []validator.List{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Type: "my_custom_type",
					},
				},
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: []specschema.ObjectValidator{
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
			},
			expected: `
"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ToString("list_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
