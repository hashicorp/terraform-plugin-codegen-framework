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

func TestGeneratorSingleNestedAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SingleNestedAttribute
		expected      GeneratorSingleNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SingleNestedAttribute is nil"),
		},
		"attributes-nil": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "bool_attribute",
						Bool: &resource.BoolAttribute{
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"bool_attribute": GeneratorBoolAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						CustomType:               convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
						PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
						Validators:               convert.NewValidators(convert.ValidatorTypeBool, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &resource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"list_attribute": GeneratorListAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						CustomType: convert.NewCustomTypeCollection(
							nil,
							nil,
							convert.CustomCollectionTypeList,
							"types.BoolType",
							"list_attribute",
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
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "nested_attribute",
						ListNested: &resource.ListNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_attribute": GeneratorListNestedAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
								},
							},
						},
						NestedAttributeObject: NewNestedAttributeObject(
							generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
									Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
								},
							},
							nil,
							convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
							convert.NewValidators(convert.ValidatorTypeObject, nil),
							"nested_attribute",
						),
						PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeList, nil),
						Validators:    convert.NewValidators(convert.ValidatorTypeList, nil),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "object_attribute",
						Object: &resource.ObjectAttribute{
							AttributeTypes: specschema.ObjectAttributeTypes{
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
			expected: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"object_attribute": GeneratorObjectAttribute{
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
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						CustomType:               convert.NewCustomTypeObject(nil, nil, "object_attribute"),
						PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
						Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: generatorschema.GeneratorAttributes{
							"nested_bool": GeneratorBoolAttribute{
								ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
								CustomType:               convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
								PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeBool, nil),
								Validators:               convert.NewValidators(convert.ValidatorTypeBool, nil),
							},
						},
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						CustomType:               convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
						PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
						Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					},
				},
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers:            convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:               convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.SingleNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &resource.SingleNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorSingleNestedAttribute{
				CustomType:         convert.NewCustomTypeNestedObject(nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				PlanModifiers:      convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:         convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &resource.SingleNestedAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorSingleNestedAttribute{
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				Description:   convert.NewDescription(pointer("description")),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &resource.SingleNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorSingleNestedAttribute{
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
				Sensitive:     convert.NewSensitive(pointer(true)),
				Validators:    convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &resource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
				CustomType:    convert.NewCustomTypeNestedObject(nil, "name"),
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
			input: &resource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
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
			input: &resource.SingleNestedAttribute{
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
			expected: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, "name"),
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

			got, err := NewGeneratorSingleNestedAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSingleNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSingleNestedAttribute
		expected []code.Import
	}{
		"default": {
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-without-import": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{},
					"",
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
					"",
				),
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
					"",
				),
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-attribute-list": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"list": GeneratorListAttribute{
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
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-attribute-list-with-custom-type": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"list": GeneratorListAttribute{
						CustomType: convert.NewCustomTypeCollection(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
							nil,
							convert.CustomCollectionTypeList,
							"",
							"",
						),
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_list",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-list-with-custom-type-with-element-with-custom-type": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"list": GeneratorListAttribute{
						CustomType: convert.NewCustomTypeCollection(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
							nil,
							convert.CustomCollectionTypeList,
							"",
							"",
						),
						ElementTypeCollection: convert.NewElementType(specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/bool",
									},
								},
							},
						}),
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_list",
				},
				{
					Path: "github.com/my_account/my_project/bool",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-attribute-object": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"obj": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
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
		"nested-attribute-object-with-custom-type": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"obj": GeneratorObjectAttribute{
						CustomType: convert.NewCustomTypeObject(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
							nil,
							"",
						),
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-with-custom-type-with-attribute-with-custom-type": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"obj": GeneratorObjectAttribute{
						AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: &code.Import{
											Path: "github.com/my_account/my_project/bool",
										},
									},
								},
							},
						}),
						CustomType: convert.NewCustomTypeObject(
							&specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
							nil,
							"",
						),
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				},
				{
					Path: "github.com/my_account/my_project/bool",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-nil": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					nil,
				),
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
		"validator-custom-import-nil": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					specschema.CustomValidators{
						&specschema.CustomValidator{},
					},
				),
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
		"validator-custom-import-empty-string": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					specschema.CustomValidators{
						&specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "",
								},
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
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-import": {
			input: GeneratorSingleNestedAttribute{
				Validators: convert.NewValidators(
					convert.ValidatorTypeObject,
					specschema.CustomValidators{
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
					},
				),
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
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"plan-modifier-custom-import": {
			input: GeneratorSingleNestedAttribute{
				PlanModifiers: convert.NewPlanModifiers(
					convert.PlanModifierTypeObject,
					specschema.CustomPlanModifiers{
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
					},
				),
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
				{
					Path: generatorschema.AttrImport,
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

func TestGeneratorSingleNestedAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSingleNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"bool": GeneratorBoolAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"list": GeneratorListAttribute{
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						ElementTypeCollection: convert.NewElementType(specschema.ElementType{
							String: &specschema.StringType{},
						}),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_list_nested": GeneratorListNestedAttribute{
						NestedAttributeObject: NewNestedAttributeObject(
							generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
							nil,
							convert.NewPlanModifiers(convert.PlanModifierTypeObject, nil),
							convert.NewValidators(convert.ValidatorTypeObject, nil),
							"nested_list_nested"),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: NestedListNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedListNestedValue{}.AttributeTypes(ctx),
},
},
},
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypesObject: convert.NewObjectAttributeTypes(specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						}),
						ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_single_nested": GeneratorSingleNestedAttribute{
						Attributes: generatorschema.GeneratorAttributes{
							"bool": GeneratorBoolAttribute{
								ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							},
						},
						CustomType: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
					},
				},
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"nested_single_nested": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: NestedSingleNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedSingleNestedValue{}.AttributeTypes(ctx),
},
},
},
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},`,
		},

		"custom-type": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(&specschema.CustomType{
					Type: "my_custom_type",
				}, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorSingleNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
				Sensitive:  convert.NewSensitive(pointer(true)),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorSingleNestedAttribute{
				CustomType:  convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSingleNestedAttribute{
				CustomType:         convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
				Validators: convert.NewValidators(convert.ValidatorTypeObject, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
				PlanModifiers: convert.NewPlanModifiers(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
PlanModifiers: []planmodifier.Object{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(nil, "single_nested_attribute"),
				Default: convert.NewDefaultCustom(&specschema.CustomDefault{
					SchemaDefinition: "my_object_default.Default()",
				}),
			},
			expected: `"single_nested_attribute": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
},
CustomType: SingleNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SingleNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Default: my_object_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("single_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSingleNestedAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSingleNestedAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "SingleNestedAttribute",
				ValueType: "SingleNestedAttributeValue",
				TfsdkName: "single_nested_attribute",
			},
		},
		"custom-type": {
			input: GeneratorSingleNestedAttribute{
				CustomType: convert.NewCustomTypeNestedObject(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					"single_nested_attribute",
				),
			},
			expected: model.Field{
				Name:      "SingleNestedAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "single_nested_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("single_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
