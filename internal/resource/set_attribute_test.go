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

func TestGeneratorSetAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SetAttribute
		expected      GeneratorSetAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SetAttribute is nil"),
		},
		"element-type-bool": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.BoolType",
					"name",
				),
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"element-type-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"element-type-list-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"element-type-map-string": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"element-type-list-object-string": {
			input: &resource.SetAttribute{
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
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"element-type-object-string": {
			input: &resource.SetAttribute{
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
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"element-type-object-list-string": {
			input: &resource.SetAttribute{
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
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.SetAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.SetAttribute{
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
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/",
						},
						Type:      "my_type",
						ValueType: "myvalue_type",
					},
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &resource.SetAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.SetAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.SetValidators{
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
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators: convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{
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
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: specschema.SetPlanModifiers{
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
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"name",
				),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
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
			expected: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					nil,
					convert.CustomCollectionTypeSet,
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
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorSetAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSetAttribute
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
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{},
					nil,
					convert.CustomCollectionTypeSet,
					"",
					"",
				),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					nil,
					convert.CustomCollectionTypeSet,
					"",
					"",
				),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					nil,
					convert.CustomCollectionTypeSet,
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeSet, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorSetAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{
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
			input: GeneratorSetAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{
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
			input: GeneratorSetAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{
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
			input: GeneratorSetAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorSetAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{
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
			input: GeneratorSetAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{
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
			input: GeneratorSetAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{
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
			input: GeneratorSetAttribute{},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-nil": {
			input: GeneratorSetAttribute{
				Default: convert.NewDefaultCustom(nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-nil": {
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
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
			input: GeneratorSetAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.SetAttribute",
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
			input: GeneratorSetAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.SetAttribute",
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
			input: GeneratorSetAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.SetAttribute",
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
						Type: "*api.SetAttribute",
					},
					convert.CustomCollectionTypeSet,
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

func TestGeneratorSetAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
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
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					convert.CustomCollectionTypeSet,
					"",
					"",
				),
			},
			expected: `"set_attribute": schema.SetAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.SetAttribute",
					},
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"set_attribute",
				),
			},
			expected: `"set_attribute": schema.SetAttribute{
CustomType: SetAttributeType{
types.SetType{
ElemType: types.StringType,
},
},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.SetAttribute",
					},
					convert.CustomCollectionTypeSet,
					"types.StringType",
					"set_attribute",
				),
			},
			expected: `"set_attribute": schema.SetAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorSetAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorSetAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorSetAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorSetAttribute{
				Description: convert.NewDescription(pointer("description")),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSetAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeSet, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Validators: []validator.Set{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeSet, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
PlanModifiers: []planmodifier.Set{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorSetAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					SchemaDefinition: "my_set_default.Default()",
				}),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.StringType,
Default: my_set_default.Default(),
},`,
		},

		"element-type-bool-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Type: "boolCustomType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: boolCustomType,
},`,
		},

		"element-type-float64-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Float64: &specschema.Float64Type{
						CustomType: &specschema.CustomType{
							Type: "float64CustomType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: float64CustomType,
},`,
		},

		"element-type-int64-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Int64: &specschema.Int64Type{
						CustomType: &specschema.CustomType{
							Type: "int64CustomType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: int64CustomType,
},`,
		},

		"element-type-list-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					List: &specschema.ListType{
						CustomType: &specschema.CustomType{
							Type: "customListType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: customListType,
},`,
		},

		"element-type-map-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Map: &specschema.MapType{
						CustomType: &specschema.CustomType{
							Type: "customMapType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: customMapType,
},`,
		},

		"element-type-number-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Number: &specschema.NumberType{
						CustomType: &specschema.CustomType{
							Type: "customNumberType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: customNumberType,
},`,
		},

		"element-type-object-custom": {
			input: GeneratorSetAttribute{
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
			expected: `"set_attribute": schema.SetAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": customBoolType,
},
},
},`,
		},

		"element-type-set-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					Set: &specschema.SetType{
						CustomType: &specschema.CustomType{
							Type: "customSetType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: customSetType,
},`,
		},

		"element-type-string-custom": {
			input: GeneratorSetAttribute{
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{
						CustomType: &specschema.CustomType{
							Type: "stringCustomType",
						},
					},
				}),
			},
			expected: `"set_attribute": schema.SetAttribute{
ElementType: stringCustomType,
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("set_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "SetAttribute",
				ValueType: "types.Set",
				TfsdkName: "set_attribute",
			},
		},
		"custom-type": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					convert.CustomCollectionTypeSet,
					"",
					"",
				),
			},
			expected: model.Field{
				Name:      "SetAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "set_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.SetAttribute",
					},
					convert.CustomCollectionTypeSet,
					"",
					"set_attribute",
				),
			},
			expected: model.Field{
				Name:      "SetAttribute",
				ValueType: "SetAttributeValue",
				TfsdkName: "set_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorSetAttribute{
				CustomType: convert.NewCustomTypeCollection(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.SetAttribute",
					},
					convert.CustomCollectionTypeSet,
					"",
					"set_attribute",
				),
			},
			expected: model.Field{
				Name:      "SetAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "set_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("set_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
