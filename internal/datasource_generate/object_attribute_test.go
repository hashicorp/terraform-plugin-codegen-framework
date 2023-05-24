package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratorObjectAttribute_ToString(t *testing.T) {
	testCases := map[string]struct {
		objectAttribute   GeneratorObjectAttribute
		expectedAttribute string
		expectedError     error
	}{
		"attr-type-bool": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"bool": types.BoolType,
					},
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},`,
		},

		"attr-type-list": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"list": types.ListType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-list-list": {
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"map": types.MapType{
							ElemType: types.BoolType,
						},
					},
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-map-map": {
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
},`,
		},

		"custom-type": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Required: true,
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Required: true,
},`,
		},

		"optional": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Optional: true,
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},`,
		},

		"computed": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Computed: true,
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Computed: true,
},`,
		},

		"sensitive": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Sensitive: true,
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Sensitive: true,
},`,
		},

		"description": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					Description: "description",
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			objectAttribute: GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"str": types.StringType,
					},
					DeprecationMessage: "deprecated",
				},
			},
			expectedAttribute: `
"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			objectAttribute: GeneratorObjectAttribute{
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
			expectedAttribute: `
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

			got, err := testCase.objectAttribute.ToString("object_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
