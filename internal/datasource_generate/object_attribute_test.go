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

func TestGeneratorObjectAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorObjectAttribute
		expected map[string]struct{}
	}{
		"default": {
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"custom-type-without-import": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer(""),
				},
			},
			expected: map[string]struct{}{},
		},
		"custom-type-with-import": {
			input: GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import: pointer("github.com/my_account/my_project/attribute"),
				},
			},
			expected: map[string]struct{}{
				"github.com/my_account/my_project/attribute": {},
			},
		},
		"object-without-attribute-types": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"object-wit-empty-attribute-types": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{},
				},
			},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"object-with-attr-type-bool": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"bool": types.BoolType,
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
			input: GeneratorObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: nil,
					},
				}},
			expected: map[string]struct{}{
				datasourceSchemaImport: {},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorObjectAttribute{
				Validators: []specschema.ObjectValidator{
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
			input: GeneratorObjectAttribute{
				Validators: []specschema.ObjectValidator{
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
			input: GeneratorObjectAttribute{
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

func TestGeneratorObjectAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorObjectAttribute
		expected      string
		expectedError error
	}{
		"attr-type-bool": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"bool": types.BoolType,
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},`,
		},

		"attr-type-list": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"list": types.ListType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-list-list": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"list": types.ListType{
							ElemType: types.ListType{
								ElemType: types.BoolType,
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},
},`,
		},

		"attr-type-list-object": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"list": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"bool": types.BoolType,
								},
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},`,
		},

		"attr-type-map": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"map": types.MapType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-map-map": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"map": types.MapType{
							ElemType: types.MapType{
								ElemType: types.BoolType,
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},
},`,
		},

		"attr-type-map-object": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"map": types.MapType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"bool": types.BoolType,
								},
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},`,
		},

		"attr-type-object": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bool": types.BoolType,
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"attr-type-object-object": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"obj_obj": types.ObjectType{
									AttrTypes: map[string]attr.Type{
										"bool": types.BoolType,
									},
								},
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"obj_obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},
},`,
		},

		"attr-type-object-list": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"list": types.ListType{
									ElemType: types.BoolType,
								},
							},
						},
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},
},
},`,
		},

		"attr-type-string": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
},`,
		},

		"custom-type": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Required: true,
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Optional: true,
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Computed: true,
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Sensitive: true,
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Description: "description",
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
				},
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
			expected: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Validators: []validator.Bool{
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

			got, err := testCase.input.ToString("object_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
