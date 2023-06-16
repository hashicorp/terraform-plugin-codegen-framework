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

func TestConvertMapNestedAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.MapNestedAttribute
		expected      resource_generate.GeneratorMapNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.MapNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &resource.MapNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.MapNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "bool_attribute",
							Bool: &resource.BoolAttribute{
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"bool_attribute": resource_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &resource.MapNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "list_attribute",
							List: &resource.ListAttribute{
								ComputedOptionalRequired: "optional",
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"list_attribute": resource_generate.GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								ElementType: types.BoolType,
								Optional:    true,
							},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.MapNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							MapNested: &resource.MapNestedAttribute{
								NestedObject: resource.NestedAttributeObject{
									Attributes: []resource.Attribute{
										{
											Name: "nested_bool",
											Bool: &resource.BoolAttribute{
												ComputedOptionalRequired: "computed",
											},
										},
									},
								},
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorMapNestedAttribute{
							NestedObject: resource_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]resource_generate.GeneratorAttribute{
									"nested_bool": resource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
							MapNestedAttribute: schema.MapNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &resource.MapNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "object_attribute",
							Object: &resource.ObjectAttribute{
								AttributeTypes: []specschema.ObjectAttributeType{
									{
										Name: "obj_bool",
										Bool: &specschema.BoolType{},
									},
								},
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"object_attribute": resource_generate.GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								AttributeTypes: map[string]attr.Type{
									"obj_bool": types.BoolType,
								},
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.MapNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							SingleNested: &resource.SingleNestedAttribute{
								Attributes: []resource.Attribute{
									{
										Name: "nested_bool",
										Bool: &resource.BoolAttribute{
											ComputedOptionalRequired: "computed",
										},
									},
								},
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]resource_generate.GeneratorAttribute{
								"nested_bool": resource_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Computed: true,
									},
								},
							},
							SingleNestedAttribute: schema.SingleNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"computed": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &resource.MapNestedAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.MapNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.MapNestedAttribute{
				Description: pointer("description"),
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &resource.MapNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				MapNestedAttribute: schema.MapNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &resource.MapNestedAttribute{
				Validators: []specschema.MapValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				Validators: []specschema.MapValidator{
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
			input: &resource.MapNestedAttribute{
				PlanModifiers: []specschema.MapPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				PlanModifiers: []specschema.MapPlanModifier{
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
			input: &resource.MapNestedAttribute{
				Default: &specschema.MapDefault{
					Custom: &specschema.CustomDefault{
						Import:           pointer("github.com/.../my_default"),
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
			expected: resource_generate.GeneratorMapNestedAttribute{
				Default: &specschema.MapDefault{
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

			got, err := convertMapNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
