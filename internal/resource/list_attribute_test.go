// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorListAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.ListAttribute
		expected      GeneratorListAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.ListAttribute is nil"),
		},
		"element-type-bool": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.BoolType",
					"name",
				),
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-string": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-list-string": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-map-string": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-list-object-string": {
			input: &resource.ListAttribute{
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
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-object-string": {
			input: &resource.ListAttribute{
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
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-object-list-string": {
			input: &resource.ListAttribute{
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
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &resource.ListAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.ListAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.ListAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.ListAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.ListAttribute{
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
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/",
						},
						Type:      "my_type",
						ValueType: "myvalue_type",
					},
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &resource.ListAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.ListAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Sensitive:     convert.NewSensitive(pointer(true)),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.ListValidators{
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
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
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
		"plan-modifiers": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: specschema.ListPlanModifiers{
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
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.ListAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.ListDefault{
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
			expected: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "github.com/.../my_default",
						},
					},
					SchemaDefinition: "my_default.Default()",
				}),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorListAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListAttribute
		expected []code.Import
	}{
		"default": {
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"custom-type-without-import": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{},
					nil,
					convert.CustomCollectionTypeList,
					"",
					"",
				),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					nil,
					convert.CustomCollectionTypeList,
					"",
					"",
				),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					nil,
					convert.CustomCollectionTypeList,
					"",
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
			},
		},
		"elem-type-bool": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"elem-type-bool-with-import": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Import: &code.Import{
								Path: "github.com/my_account/my_project/element",
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/element",
				},
			},
		},
		"elem-type-list-bool": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"elem-type-list-bool-with-import": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/element",
									},
								},
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/element",
				},
			},
		},
		"elem-type-object": {
			input: GeneratorListAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"elem-type-object-bool": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "b",
								Bool: &specschema.BoolType{},
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"elem-type-object-bool-with-import": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: &code.Import{
											Path: "github.com/my_account/my_project/element",
										},
									},
								},
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/element",
				},
			},
		},
		"elem-type-object-with-imports": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "b",
								Bool: &specschema.BoolType{},
							},
							{
								Name: "c",
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: &code.Import{
											Path: "github.com/my_account/my_project/element",
										},
									},
								},
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
				{
					Path: "github.com/my_account/my_project/element",
				},
			},
		},
		"validator-custom-nil": {
			input: GeneratorListAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeList, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorListAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
					&specschema.CustomValidator{},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorListAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "",
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import": {
			input: GeneratorListAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/myotherproject/myvalidators/validator",
							},
						},
					},
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/myvalidators/validator",
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.ValidatorImport,
				},
				{
					Path: "github.com/myotherproject/myvalidators/validator",
				},
				{
					Path: "github.com/myproject/myvalidators/validator",
				},
			},
		},
		"plan-modifier-custom-nil": {
			input: GeneratorListAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorListAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifiers-custom-import-empty-string": {
			input: GeneratorListAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "",
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifier-custom-import": {
			input: GeneratorListAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/myotherproject/myplanmodifiers/planmodifier",
							},
						},
					},
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/myplanmodifiers/planmodifier",
							},
						},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.PlanModifierImport,
				},
				{
					Path: "github.com/myotherproject/myplanmodifiers/planmodifier",
				},
				{
					Path: "github.com/myproject/myplanmodifiers/planmodifier",
				},
			},
		},
		"default-nil": {
			input: GeneratorListAttribute{},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-nil": {
			input: GeneratorListAttribute{
				Default: convert.NewDefaultCustom(nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-nil": {
			input: GeneratorListAttribute{
				Default: convert.NewDefaultCustom(
					&specschema.CustomDefault{},
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-empty-string": {
			input: GeneratorListAttribute{
				Default: convert.NewDefaultCustom(
					&specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "",
							},
						},
					},
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import": {
			input: GeneratorListAttribute{
				Default: convert.NewDefaultCustom(
					&specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/myproject/mydefaults/default",
							},
						},
					},
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/myproject/mydefaults/default",
				},
			},
		},
		"associated-external-type": {
			input: GeneratorListAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types",
				},
				{
					Path: "fmt",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/diag",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/attr",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-go/tftypes",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
				},
			},
		},
		"associated-external-type-with-import": {
			input: GeneratorListAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.ListAttribute",
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types",
				},
				{
					Path: "fmt",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/diag",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/attr",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-go/tftypes",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
				},
				{
					Path: "github.com/api",
				},
			},
		},
		"associated-external-type-with-custom-type": {
			input: GeneratorListAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.ListAttribute",
					},
				},
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					&specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.ListAttribute",
					},
					convert.CustomCollectionTypeList,
					"",
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: "fmt",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/diag",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/attr",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-go/tftypes",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
				},
				{
					Path: "github.com/api",
				},
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports().All()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
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
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
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
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					convert.CustomCollectionTypeList,
					"",
					"",
				),
			},
			expected: `"list_attribute": schema.ListAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
					convert.CustomCollectionTypeList,
					"types.StringType",
					"list_attribute",
				),
			},
			expected: `"list_attribute": schema.ListAttribute{
CustomType: ListAttributeType{
types.ListType{
ElemType: types.StringType,
},
},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
					convert.CustomCollectionTypeList,
					"types.StringType",
					"list_attribute",
				),
			},
			expected: `"list_attribute": schema.ListAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorListAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorListAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorListAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorListAttribute{
				Description: convert.NewDescription(pointer("description")),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Validators: []validator.List{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
PlanModifiers: []planmodifier.List{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorListAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					SchemaDefinition: "my_list_default.Default()",
				}),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Default: my_list_default.Default(),
},`,
		},

		"element-type-bool-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Type: "boolCustomType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: boolCustomType,
},`,
		},

		"element-type-float64-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Float64: &specschema.Float64Type{
						CustomType: &specschema.CustomType{
							Type: "float64CustomType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: float64CustomType,
},`,
		},

		"element-type-int64-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Int64: &specschema.Int64Type{
						CustomType: &specschema.CustomType{
							Type: "int64CustomType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: int64CustomType,
},`,
		},

		"element-type-list-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						CustomType: &specschema.CustomType{
							Type: "customListType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: customListType,
},`,
		},

		"element-type-map-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						CustomType: &specschema.CustomType{
							Type: "customMapType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: customMapType,
},`,
		},

		"element-type-number-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Number: &specschema.NumberType{
						CustomType: &specschema.CustomType{
							Type: "customNumberType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: customNumberType,
},`,
		},

		"element-type-object-custom": {
			input: GeneratorListAttribute{
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
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": customBoolType,
},
},
},`,
		},

		"element-type-set-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Set: &specschema.SetType{
						CustomType: &specschema.CustomType{
							Type: "customSetType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: customSetType,
},`,
		},

		"element-type-string-custom": {
			input: GeneratorListAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{
						CustomType: &specschema.CustomType{
							Type: "stringCustomType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: stringCustomType,
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("list_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "ListAttribute",
				ValueType: "types.List",
				TfsdkName: "list_attribute",
			},
		},
		"custom-type": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					convert.CustomCollectionTypeList,
					"",
					"",
				),
			},
			expected: model.Field{
				Name:      "ListAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "list_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
					convert.CustomCollectionTypeList,
					"",
					"list_attribute",
				),
			},
			expected: model.Field{
				Name:      "ListAttribute",
				ValueType: "ListAttributeValue",
				TfsdkName: "list_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ListAttribute",
					},
					convert.CustomCollectionTypeList,
					"",
					"list_attribute",
				),
			},
			expected: model.Field{
				Name:      "ListAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "list_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("list_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
