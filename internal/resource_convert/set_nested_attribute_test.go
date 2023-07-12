// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestConvertSetNestedAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SetNestedAttribute
		expected      resource_generate.GeneratorSetNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SetNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &resource.SetNestedAttribute{
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
			input: &resource.SetNestedAttribute{
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
			expected: resource_generate.GeneratorSetNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]generatorschema.GeneratorAttribute{
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
			input: &resource.SetNestedAttribute{
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
			expected: resource_generate.GeneratorSetNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]generatorschema.GeneratorAttribute{
						"list_attribute": resource_generate.GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.SetNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							SetNested: &resource.SetNestedAttribute{
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
			expected: resource_generate.GeneratorSetNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]generatorschema.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorSetNestedAttribute{
							NestedObject: resource_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]generatorschema.GeneratorAttribute{
									"nested_bool": resource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
							SetNestedAttribute: schema.SetNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &resource.SetNestedAttribute{
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
			expected: resource_generate.GeneratorSetNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]generatorschema.GeneratorAttribute{
						"object_attribute": resource_generate.GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
							AttributeTypes: []specschema.ObjectAttributeType{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.SetNestedAttribute{
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
			expected: resource_generate.GeneratorSetNestedAttribute{
				NestedObject: resource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]generatorschema.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]generatorschema.GeneratorAttribute{
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
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &resource.SetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.SetNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.SetNestedAttribute{
				Description: pointer("description"),
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &resource.SetNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &resource.SetNestedAttribute{
				Validators: []specschema.SetValidator{
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
			expected: resource_generate.GeneratorSetNestedAttribute{
				Validators: []specschema.SetValidator{
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
		},
		"plan-modifiers": {
			input: &resource.SetNestedAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
		},
		"default": {
			input: &resource.SetNestedAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
			expected: resource_generate.GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
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

			got, err := convertSetNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
