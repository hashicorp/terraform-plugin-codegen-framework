// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratorListAttribute_ToString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		listAttribute     GeneratorListAttribute
		expectedAttribute string
		expectedError     error
	}{
		"element-type-bool": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.BoolType,
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ListType{
						ElemType: types.BoolType,
					},
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ListType{
						ElemType: types.ListType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			listAttribute: GeneratorListAttribute{
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
			expectedAttribute: `
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
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.MapType{
						ElemType: types.BoolType,
					},
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.MapType{
						ElemType: types.MapType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			listAttribute: GeneratorListAttribute{
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
			expectedAttribute: `
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
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"bool": types.BoolType,
						},
					},
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			listAttribute: GeneratorListAttribute{
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
			expectedAttribute: `
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
			listAttribute: GeneratorListAttribute{
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
			expectedAttribute: `
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
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
CustomType: my_custom_type,
},`,
		},

		"required": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Required:    true,
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"sensitive": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Sensitive:   true,
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType: types.StringType,
					Description: "description",
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			listAttribute: GeneratorListAttribute{
				ListAttribute: schema.ListAttribute{
					ElementType:        types.StringType,
					DeprecationMessage: "deprecated",
				},
			},
			expectedAttribute: `
"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			listAttribute: GeneratorListAttribute{
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
			expectedAttribute: `
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

			got, err := testCase.listAttribute.ToString("list_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
