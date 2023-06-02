// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratorListAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorListAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorListAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorListAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"elem-type-bool": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.BoolType,
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
				typesImport:            {},
			},
		},
		"elem-type-object": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ObjectType{},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
				typesImport:            {},
			},
		},
		"elem-type-object-bool": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
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
		"validator-custom-nil": {
			input: GeneratorListAttribute{
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
			input: GeneratorListAttribute{
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
			input: GeneratorListAttribute{
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
			input: GeneratorListAttribute{
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

func TestGeneratorListAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.BoolType,
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ListType{
						ElemType: types.BoolType,
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ListType{
						ElemType: types.ListType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
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
"list_attribute": schema.ListAttribute{
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
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.MapType{
						ElemType: types.BoolType,
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.MapType{
						ElemType: types.MapType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
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
"list_attribute": schema.ListAttribute{
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
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"bool": types.BoolType,
						},
					},
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
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
"list_attribute": schema.ListAttribute{
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
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
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
"list_attribute": schema.ListAttribute{
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
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Required:    true,
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Computed:    true,
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Sensitive:   true,
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Description: "description",
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType:        types.StringType,
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
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
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Validators: []validator.List{
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

			got, err := testCase.input.ToString("list_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
