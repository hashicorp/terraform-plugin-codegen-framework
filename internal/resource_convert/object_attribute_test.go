// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
)

func TestConvertObjectAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.ObjectAttribute
		expected      resource_generate.GeneratorObjectAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.ObjectAttribute is nil"),
		},
		"attribute-type-nil": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", specschema.ObjectAttributeType{
				Name: "empty",
			}),
		},
		"attribute-type-bool": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_bool": types.BoolType,
					},
				},
			},
		},
		"attribute-type-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_string": types.StringType,
					},
				},
			},
		},
		"attribute-type-list-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_list_string": types.ListType{
							ElemType: types.StringType,
						},
					},
				},
			},
		},
		"attribute-type-map-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_map_string",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_map_string": types.MapType{
							ElemType: types.StringType,
						},
					},
				},
			},
		},
		"attribute-type-list-object-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: []specschema.ObjectAttributeType{
									{
										Name:   "obj_string",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_list_object_string": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"obj_string": types.StringType,
								},
							},
						},
					},
				},
			},
		},
		"attribute-type-object-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_string": types.StringType,
					},
				},
			},
		},
		"attribute-type-object-list-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: []specschema.ObjectAttributeType{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_list_string": types.ListType{
							ElemType: types.StringType,
						},
					},
				},
			},
		},
		"computed": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &resource.ObjectAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.ObjectAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.ObjectAttribute{
				Description: pointer("description"),
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &resource.ObjectAttribute{
				Sensitive: pointer(true),
			},
			expected: resource_generate.GeneratorObjectAttribute{
				ObjectAttribute: schema.ObjectAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &resource.ObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				Validators: []specschema.ObjectValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
		"plan-modifiers": {
			input: &resource.ObjectAttribute{
				PlanModifiers: []specschema.ObjectPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				PlanModifiers: []specschema.ObjectPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
		},
		"default": {
			input: &resource.ObjectAttribute{
				Default: &specschema.ObjectDefault{
					Custom: &specschema.CustomDefault{
						Import:           pointer("github.com/.../my_default"),
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
			expected: resource_generate.GeneratorObjectAttribute{
				Default: &specschema.ObjectDefault{
					Custom: &specschema.CustomDefault{
						Import:           pointer("github.com/.../my_default"),
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertObjectAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
