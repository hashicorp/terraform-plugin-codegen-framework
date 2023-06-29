// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

func TestConvertListNestedAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.ListNestedAttribute
		expected      datasource_generate.GeneratorListNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.ListNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
					Attributes: []datasource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", datasource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
					Attributes: []datasource.Attribute{
						{
							Name: "bool_attribute",
							Bool: &datasource.BoolAttribute{
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				NestedObject: datasource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]datasource_generate.GeneratorAttribute{
						"bool_attribute": datasource_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
					Attributes: []datasource.Attribute{
						{
							Name: "list_attribute",
							List: &datasource.ListAttribute{
								ComputedOptionalRequired: "optional",
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				NestedObject: datasource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]datasource_generate.GeneratorAttribute{
						"list_attribute": datasource_generate.GeneratorListAttribute{
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
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
					Attributes: []datasource.Attribute{
						{
							Name: "nested_attribute",
							ListNested: &datasource.ListNestedAttribute{
								NestedObject: datasource.NestedAttributeObject{
									Attributes: []datasource.Attribute{
										{
											Name: "nested_bool",
											Bool: &datasource.BoolAttribute{
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
			expected: datasource_generate.GeneratorListNestedAttribute{
				NestedObject: datasource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]datasource_generate.GeneratorAttribute{
						"nested_attribute": datasource_generate.GeneratorListNestedAttribute{
							NestedObject: datasource_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]datasource_generate.GeneratorAttribute{
									"nested_bool": datasource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
							ListNestedAttribute: schema.ListNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
					Attributes: []datasource.Attribute{
						{
							Name: "object_attribute",
							Object: &datasource.ObjectAttribute{
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
			expected: datasource_generate.GeneratorListNestedAttribute{
				NestedObject: datasource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]datasource_generate.GeneratorAttribute{
						"object_attribute": datasource_generate.GeneratorObjectAttribute{
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
			input: &datasource.ListNestedAttribute{
				NestedObject: datasource.NestedAttributeObject{
					Attributes: []datasource.Attribute{
						{
							Name: "nested_attribute",
							SingleNested: &datasource.SingleNestedAttribute{
								Attributes: []datasource.Attribute{
									{
										Name: "nested_bool",
										Bool: &datasource.BoolAttribute{
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
			expected: datasource_generate.GeneratorListNestedAttribute{
				NestedObject: datasource_generate.GeneratorNestedAttributeObject{
					Attributes: map[string]datasource_generate.GeneratorAttribute{
						"nested_attribute": datasource_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]datasource_generate.GeneratorAttribute{
								"nested_bool": datasource_generate.GeneratorBoolAttribute{
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
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &datasource.ListNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &datasource.ListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
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
			input: &datasource.ListNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &datasource.ListNestedAttribute{
				Description: pointer("description"),
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &datasource.ListNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorListNestedAttribute{
				ListNestedAttribute: schema.ListNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &datasource.ListNestedAttribute{
				Validators: []specschema.ListValidator{
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
			expected: datasource_generate.GeneratorListNestedAttribute{
				Validators: []specschema.ListValidator{
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertListNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
