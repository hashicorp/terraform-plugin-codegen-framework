// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

func TestGeneratorMapAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.MapAttribute
		expected      GeneratorMapAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.MapAttribute is nil"),
		},
		"element-type-bool": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.BoolType",
					"name",
				),
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"element-type-string": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"element-type-list-string": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.ListType{\nElemType: types.StringType,\n}",
					"name",
				),
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"element-type-map-string": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.MapType{\nElemType: types.StringType,\n}",
					"name",
				),
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"element-type-list-object-string": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name:   "str",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.ListType{\nElemType: types.ObjectType{\nAttrTypes: map[string]attr.Type{\n\"str\": types.StringType,\n},\n},\n}",
					"name",
				),
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name:   "str",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name:   "str",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"element-type-object-string": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.ObjectType{\nAttrTypes: map[string]attr.Type{\n\"str\": types.StringType,\n},\n}",
					"name",
				),
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"element-type-object-list-string": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.ObjectType{\nAttrTypes: map[string]attr.Type{\n\"list\": types.ListType{\nElemType: types.StringType,\n},\n},\n}",
					"name",
				),
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &datasource.MapAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &datasource.MapAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &datasource.MapAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &datasource.MapAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &datasource.MapAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/",
						},
						Type:      "my_type",
						ValueType: "myvalue_type",
					},
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &datasource.MapAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &datasource.MapAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				Description: convert.NewDescription(pointer("description")),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &datasource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.MapValidators{
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
			expected: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/.../myvalidator",
							},
						},
						SchemaDefinition: "myvalidator.Validate()",
					},
				}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorMapAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorMapAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorMapAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
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
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Map: &specschema.MapType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
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
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "obj",
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name: "bool",
											Bool: &specschema.BoolType{},
										},
									},
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
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
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
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
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					convert.CustomCollectionTypeMap,
					"",
					"",
				),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					&specschema.AssociatedExternalType{Type: "*api.MapAttribute"},
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"map_attribute",
				),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
CustomType: MapAttributeType{
types.MapType{
ElemType: types.StringType,
},
},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{Type: "my_custom_type"},
					&specschema.AssociatedExternalType{Type: "*api.MapAttribute"},
					convert.CustomCollectionTypeMap,
					"types.StringType",
					"name",
				),
			},
			expected: `"map_attribute": schema.MapAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorMapAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorMapAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorMapAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorMapAttribute{
				Description: convert.NewDescription(pointer("description")),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorMapAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeMap, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},

					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Validators: []validator.Map{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"element-type-bool-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Type: "boolCustomType",
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: boolCustomType,
},`,
		},

		"element-type-float64-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Float64: &specschema.Float64Type{
						CustomType: &specschema.CustomType{
							Type: "float64CustomType",
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: float64CustomType,
},`,
		},

		"element-type-int64-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Int64: &specschema.Int64Type{
						CustomType: &specschema.CustomType{
							Type: "int64CustomType",
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: int64CustomType,
},`,
		},

		"element-type-list-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						CustomType: &specschema.CustomType{
							Type: "customListType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: customListType,
},`,
		},

		"element-type-map-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						CustomType: &specschema.CustomType{
							Type: "customMapType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: customMapType,
},`,
		},

		"element-type-number-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Number: &specschema.NumberType{
						CustomType: &specschema.CustomType{
							Type: "numberCustomType",
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: numberCustomType,
},`,
		},

		"element-type-object-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "customBoolType",
									},
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": customBoolType,
},
},
},`,
		},

		"element-type-set-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Set: &specschema.SetType{
						CustomType: &specschema.CustomType{
							Type: "customSetType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: customSetType,
},`,
		},

		"element-type-string-custom": {
			input: GeneratorMapAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{
						CustomType: &specschema.CustomType{
							Type: "stringCustomType",
						},
					},
				}),
			},
			expected: `"map_attribute": schema.MapAttribute{
ElementType: stringCustomType,
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("map_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorMapAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorMapAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "types.Map",
				TfsdkName: "map_attribute",
			},
		},
		"custom-type": {
			input: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					convert.CustomCollectionTypeMap,
					"",
					"",
				),
			},
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "map_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.MapAttribute",
					},
					convert.CustomCollectionTypeMap,
					"",
					"map_attribute",
				),
			},
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "MapAttributeValue",
				TfsdkName: "map_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorMapAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
					convert.CustomCollectionTypeMap,
					"",
					"map_attribute",
				),
			},
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "map_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("map_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
