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

func TestGeneratorObjectAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.ObjectAttribute
		expected      GeneratorObjectAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.ObjectAttribute is nil"),
		},
		"attribute-type-bool": {
			input: &resource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
			},
			expected: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj_bool",
						Bool: &specschema.BoolType{},
					},
				}),
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attribute-type-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				}),
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attribute-type-list-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
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
			expected: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				}),
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attribute-type-map-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
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
			expected: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_map_string",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj_map_string",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				}),
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attribute-type-list-object-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name:   "obj_str",
											String: &specschema.StringType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name:   "obj_str",
											String: &specschema.StringType{},
										},
									},
								},
							},
						},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_object_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name:   "obj_str",
											String: &specschema.StringType{},
										},
									},
								},
							},
						},
					},
				}),
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attribute-type-object-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
			},
			expected: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "obj_string",
						String: &specschema.StringType{},
					},
				}),
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attribute-type-object-list-string": {
			input: &resource.ObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
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
			expected: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj_list_string",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				}),
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorObjectAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorObjectAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType:               convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorObjectAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{})},
		},
		"required": {
			input: &resource.ObjectAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorObjectAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.ObjectAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
					nil,
					"name",
				),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{})},
		},
		"deprecation_message": {
			input: &resource.ObjectAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorObjectAttribute{
				CustomType:         convert.NewCustomTypeObject(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				PlanModifiers:      convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:         convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{})},
		},
		"description": {
			input: &resource.ObjectAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorObjectAttribute{
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				Description:   convert.NewDescription(pointer("description")),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.ObjectAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorObjectAttribute{
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				Sensitive:     convert.NewSensitive(pointer(true)),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.ObjectAttribute{
				Validators: specschema.ObjectValidators{
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
			expected: GeneratorObjectAttribute{
				CustomType:    convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
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
			input: &resource.ObjectAttribute{
				PlanModifiers: specschema.ObjectPlanModifiers{
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
			expected: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(nil, nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"default": {
			input: &resource.ObjectAttribute{
				Default: &specschema.ObjectDefault{
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
			expected: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(nil, nil, "name"),
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "github.com/.../my_default",
						},
					},
					SchemaDefinition: "my_default.Default()",
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorObjectAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorObjectAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorObjectAttribute
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
			input: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{},
					nil,
					"",
				),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					nil,
					"",
				),
			},
			expected: []code.Import{},
		},
		"custom-type-with-import": {
			input: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					nil,
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
			},
		},
		"object-without-attribute-types": {
			input: GeneratorObjectAttribute{},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"object-with-empty-attribute-types": {
			input: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"object-with-attr-type-bool": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "bool",
						Bool: &specschema.BoolType{},
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
		// verifies that math/big is imported when object has
		// attribute type that is number type.
		"object-with-attr-type-number": {
			input: GeneratorObjectAttribute{
				AttributeTypes: specschema.ObjectAttributeTypes{
					{
						Name:   "number",
						Number: &specschema.NumberType{},
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "number",
						Number: &specschema.NumberType{},
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
					Path: generatorschema.MathBigImport,
				},
			},
		},
		"object-with-attr-type-bool-with-import": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
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
		"object-with-attr-type-bool-with-imports": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
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
					{
						Name: "list",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: &code.Import{
											Path: "github.com/my_account/my_project/another_element",
										},
									},
								},
							},
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/list",
								},
							},
						},
					},
					{
						Name:   "str",
						String: &specschema.StringType{},
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
				{
					Path: "github.com/my_account/my_project/list",
				},
				{
					Path: "github.com/my_account/my_project/another_element",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-nil": {
			input: GeneratorObjectAttribute{
				Validators: convert.NewValidators(convert.ValidatorTypeList, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorObjectAttribute{
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
			input: GeneratorObjectAttribute{
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
			input: GeneratorObjectAttribute{
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
			input: GeneratorObjectAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"plan-modifier-custom-import-nil": {
			input: GeneratorObjectAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
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
			input: GeneratorObjectAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
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
			input: GeneratorObjectAttribute{
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
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
			input: GeneratorObjectAttribute{},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-nil": {
			input: GeneratorObjectAttribute{
				Default: convert.NewDefaultCustom(nil),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-nil": {
			input: GeneratorObjectAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{}),
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
			},
		},
		"default-custom-import-empty-string": {
			input: GeneratorObjectAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "",
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
		"default-custom-import": {
			input: GeneratorObjectAttribute{
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "github.com/myproject/mydefaults/default",
						},
					},
				}),
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
			input: GeneratorObjectAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.ObjectAttribute",
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
			input: GeneratorObjectAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.ObjectAttribute",
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
			input: GeneratorObjectAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.ObjectAttribute",
					},
				},
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					&specschema.AssociatedExternalType{
						Import: &code.Import{
							Path: "github.com/api",
						},
						Type: "*api.ObjectAttribute",
					},
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
		// verifies that math/big is not imported when associated external type
		// is specified.
		"associated-external-type-with-number-attribute": {
			input: GeneratorObjectAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.ObjectAttribute",
					},
				},
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "number",
						Number: &specschema.NumberType{},
					},
				}),
			},
			expected: []code.Import{
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/attr",
				},
				{
					Path: "fmt",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/diag",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-go/tftypes",
				},
				{
					Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
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

func TestGeneratorObjectAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorObjectAttribute
		expected      string
		expectedError error
	}{
		"attr-type-bool": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "bool",
						Bool: &specschema.BoolType{},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},`,
		},

		"attr-type-list": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "list",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-list-list": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "list",
						List: &specschema.ListType{
							ElementType: specschema.ElementType{
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
			expected: `"object_attribute": schema.ObjectAttribute{
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
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "list",
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
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
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
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "map",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"attr-type-map-map": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "map",
						Map: &specschema.MapType{
							ElementType: specschema.ElementType{
								Map: &specschema.MapType{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
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
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "map",
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
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
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
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
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
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"attr-type-obj-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj",
						Object: &specschema.ObjectType{
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "bool",
									Bool: &specschema.BoolType{},
								},
							},
							CustomType: &specschema.CustomType{
								Type: "objCustomType",
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"obj": objCustomType,
},
},`,
		},

		"attr-type-object-object": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj",
						Object: &specschema.ObjectType{
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "obj_obj",
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
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
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
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "obj",
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
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
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
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
},`,
		},

		"custom-type": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					"",
				),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					specschema.ObjectAttributeType{
						Name: "bool",
						Bool: &specschema.BoolType{},
					},
				}),
				CustomType: convert.NewCustomTypeObject(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ObjectAttribute",
					},
					"object_attribute",
				),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
CustomType: ObjectAttributeType{
types.ObjectType{
AttrTypes: ObjectAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					specschema.ObjectAttributeType{
						Name: "bool",
						Bool: &specschema.BoolType{},
					},
				}),
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ObjectAttribute",
					},
					"object_attribute",
				),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				Validators:         convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},

					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
PlanModifiers: []planmodifier.Object{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name:   "str",
						String: &specschema.StringType{},
					},
				}),
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					SchemaDefinition: "my_object_default.Default()",
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Default: my_object_default.Default(),
},`,
		},

		"attr-type-bool-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "bool",
						Bool: &specschema.BoolType{
							CustomType: &specschema.CustomType{
								Type: "boolCustomType",
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"bool": boolCustomType,
},
},`,
		},

		"attr-type-float64-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "float64",
						Float64: &specschema.Float64Type{
							CustomType: &specschema.CustomType{
								Type: "float64CustomType",
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"float64": float64CustomType,
},
},`,
		},

		"attr-type-int64-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "int64",
						Int64: &specschema.Int64Type{
							CustomType: &specschema.CustomType{
								Type: "int64CustomType",
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"int64": int64CustomType,
},
},`,
		},

		"attr-type-list-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "list",
						List: &specschema.ListType{
							CustomType: &specschema.CustomType{
								Type: "listCustomType",
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "boolCustomType",
									},
								},
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"list": listCustomType,
},
},`,
		},

		"attr-type-map-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "map",
						Map: &specschema.MapType{
							CustomType: &specschema.CustomType{
								Type: "mapCustomType",
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "boolCustomType",
									},
								},
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"map": mapCustomType,
},
},`,
		},

		"attr-type-number-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "number",
						Number: &specschema.NumberType{
							CustomType: &specschema.CustomType{
								Type: "numberCustomType",
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"number": numberCustomType,
},
},`,
		},

		"attr-type-object-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "object",
						Object: &specschema.ObjectType{
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "bool",
									Bool: &specschema.BoolType{
										CustomType: &specschema.CustomType{
											Type: "boolCustomType",
										},
									},
								},
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": boolCustomType,
},
},
},
},`,
		},

		"attr-type-set-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "set",
						Set: &specschema.SetType{
							CustomType: &specschema.CustomType{
								Type: "setCustomType",
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "boolCustomType",
									},
								},
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"set": setCustomType,
},
},`,
		},

		"attr-type-string-custom": {
			input: GeneratorObjectAttribute{
				AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
					{
						Name: "string",
						Number: &specschema.NumberType{
							CustomType: &specschema.CustomType{
								Type: "stringCustomType",
							},
						},
					},
				}),
			},
			expected: `"object_attribute": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"string": stringCustomType,
},
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("object_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorObjectAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorObjectAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "ObjectAttribute",
				ValueType: "types.Object",
				TfsdkName: "object_attribute",
			},
		},
		"custom-type": {
			input: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					"",
				),
			},
			expected: model.Field{
				Name:      "ObjectAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "object_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ObjectAttribute",
					},
					"object_attribute",
				),
			},
			expected: model.Field{
				Name:      "ObjectAttribute",
				ValueType: "ObjectAttributeValue",
				TfsdkName: "object_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorObjectAttribute{
				CustomType: convert.NewCustomTypeObject(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ObjectAttribute",
					},
					"object_attribute",
				),
			},
			expected: model.Field{
				Name:      "ObjectAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "object_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("object_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
