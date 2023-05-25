package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratorMapAttribute_Imports(t *testing.T) {
	testCases := map[string]struct {
		input    GeneratorMapAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorMapAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorMapAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorMapAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"elem-type-bool": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.BoolType,
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
				typesImport:            {},
			},
		},
		"elem-type-object": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ObjectType{},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
				typesImport:            {},
			},
		},
		"elem-type-object-bool": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"b": types.BoolType,
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

func TestGeneratorMapAttribute_ToString(t *testing.T) {
	testCases := map[string]struct {
		input         GeneratorMapAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.BoolType,
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ListType{
						ElemType: types.BoolType,
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ListType{
						ElemType: types.ListType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ListType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bool": types.BoolType,
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"element-type-map": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.MapType{
						ElemType: types.BoolType,
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.MapType{
						ElemType: types.MapType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.MapType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bool": types.BoolType,
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"element-type-object": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"bool": types.BoolType,
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"obj": types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"bool": types.BoolType,
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},`,
		},

		"element-type-object-list": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"list": types.ListType{
								ElemType: types.BoolType,
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},
},`,
		},

		"element-type-string": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
					Required:    true,
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
					Computed:    true,
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
					Sensitive:   true,
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
					Description: "description",
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType:        types.StringType,
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					ElementType: types.StringType,
				},
				Validators: []specschema.MapValidator{
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
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Validators: []validator.Map{
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

			got, err := testCase.input.ToString("map_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
