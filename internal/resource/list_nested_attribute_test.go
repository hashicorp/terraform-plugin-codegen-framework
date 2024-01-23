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

func TestGeneratorListNestedAttribute_New(t *testing.T) {
	t.Parallel()

	attributes, err := NewAttributes(resource.Attributes{})

	if err != nil {
		t.Error(err)
	}

	testCases := map[string]struct {
		input         *resource.ListNestedAttribute
		expected      GeneratorListNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.ListNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &resource.ListNestedAttribute{
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
			input: &resource.ListNestedAttribute{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "bool_attribute"),
							PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-bool": {
			input: &resource.ListNestedAttribute{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypeCollection: convert.NewCustomTypeCollection(
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
							PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypeCollection: convert.NewCustomTypeCollection(
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
							PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.ListNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
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
			},
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedAttributeObject: NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_attribute",
							),
							PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorListNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							NestedAttributeObject: NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
										CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
										PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
										ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
									},
								},
								nil,
								convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
								"nested_attribute",
							),
							PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &resource.ListNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
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
			},
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
							CustomTypeObject:         convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
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
							CustomTypeObject:         convert.NewCustomTypeObject(nil, nil, "object_attribute"),
							PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.ListNestedAttribute{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypeNestedObject:   convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
									CustomTypePrimitive:      convert.NewCustomTypePrimitive(nil, nil, "nested_bool"),
									PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeBool, specschema.CustomPlanModifiers{}),
									ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeBool, specschema.CustomValidators{}),
								},
							},
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							CustomTypeNestedObject:   convert.NewCustomTypeNestedObject(nil, "nested_attribute"),
							PlanModifiersCustom:      convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:         convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &resource.ListNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.ListNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.ListNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{})},
		},
		"required": {
			input: &resource.ListNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.ListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"deprecation_message": {
			input: &resource.ListNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorListNestedAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"description": {
			input: &resource.ListNestedAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorListNestedAttribute{
				Description: convert.NewDescription(pointer("description")),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"sensitive": {
			input: &resource.ListNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				Sensitive:           convert.NewSensitive(pointer(true)),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"validators": {
			input: &resource.ListNestedAttribute{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
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
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
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
			input: &resource.ListNestedAttribute{
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
			expected: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
		"default": {
			input: &resource.ListNestedAttribute{
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
			expected: GeneratorListNestedAttribute{
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
				DefaultCustom: convert.NewDefaultCustom(&specschema.CustomDefault{
					Imports: []code.Import{
						{
							Path: "github.com/.../my_default",
						},
					},
					SchemaDefinition: "my_default.Default()",
				}),
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: attributes,
				},
				NestedAttributeObject: NewNestedAttributeObject(
					attributes,
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeList, nil),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorListNestedAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorListNestedAttribute
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
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
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
		"custom-type-and-nested-object-custom-type-without-import": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-with-nested-object-custom-type-with-import": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/nested_object",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-list": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-list-with-custom-type": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							CustomTypeCollection: convert.NewCustomTypeCollection(
								&specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/nested_list",
									},
								},
								nil,
								convert.CustomCollectionTypeList,
								"",
								"list",
							),
						},
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
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							CustomTypeCollection: convert.NewCustomTypeCollection(
								&specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/nested_list",
									},
								},
								nil,
								convert.CustomCollectionTypeList,
								"",
								"list",
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
		"nested-object": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
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
		"nested-object-with-custom-type": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"obj": GeneratorObjectAttribute{
							CustomTypeObject: convert.NewCustomTypeObject(
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
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"obj": GeneratorObjectAttribute{
							CustomTypeObject: convert.NewCustomTypeObject(
								&specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/nested_object",
									},
								},
								nil,
								"",
							),
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
						},
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
			input: GeneratorListNestedAttribute{
				Validators: specschema.ListValidators{
					{
						Custom: nil,
					},
				}},
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
			input: GeneratorListNestedAttribute{
				Validators: specschema.ListValidators{
					{
						Custom: &specschema.CustomValidator{},
					},
				}},
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
			input: GeneratorListNestedAttribute{
				Validators: specschema.ListValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "",
								},
							},
						},
					},
				}},
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
			input: GeneratorListNestedAttribute{
				Validators: specschema.ListValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myotherproject/myvalidators/validator",
								},
							},
						},
					},
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myproject/myvalidators/validator",
								},
							},
						},
					},
				}},
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
		"nested-object-validator-custom-nil": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: nil,
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
		"nested-object-validator-custom-import-nil": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: &specschema.CustomValidator{},
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
		"nested-object-validator-custom-import-empty-string": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: &specschema.CustomValidator{
								Imports: []code.Import{
									{
										Path: "",
									},
								},
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
		"nested-object-validator-custom-import": {
			input: GeneratorListNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: &specschema.CustomValidator{
								Imports: []code.Import{
									{
										Path: "github.com/myotherproject/myvalidators/validator",
									},
								},
							},
						},
						{
							Custom: &specschema.CustomValidator{
								Imports: []code.Import{
									{
										Path: "github.com/myproject/myvalidators/validator",
									},
								},
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports().All()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListNestedAttribute_Schema(t *testing.T) {
	t.Parallel()

	attributeName := "list_nested_attribute"

	testCases := map[string]struct {
		input         GeneratorListNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"bool": GeneratorBoolAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
							ElementTypeCollection: convert.NewElementType(specschema.ElementType{
								String: &specschema.StringType{},
							}),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_list_nested": GeneratorListNestedAttribute{
							NestedAttributeObject: NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									},
								},
								nil,
								convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
								"nested_list_nested",
							),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
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
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
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
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_single_nested": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
								},
							},
							CustomTypeNestedObject: convert.NewCustomTypeNestedObject(nil, "nested_single_nested"),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
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
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorListNestedAttribute{
				CustomTypeNestedCollection: convert.NewCustomTypeNestedCollection(&specschema.CustomType{
					Type: "my_custom_type",
				}),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorListNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorListNestedAttribute{
				Description: convert.NewDescription(pointer("description")),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorListNestedAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"list_nested_attribute",
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeList, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Validators: []validator.List{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{
						&specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
						&specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					}),
					attributeName,
				),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeList, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
PlanModifiers: []planmodifier.List{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorListNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
				DefaultCustom: convert.NewDefaultCustom(&specschema.CustomDefault{
					SchemaDefinition: "my_list_default.Default()",
				}),
			},
			expected: `"list_nested_attribute": schema.ListNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: ListNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Default: my_list_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("list_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorListNestedAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorListNestedAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "ListNestedAttribute",
				ValueType: "types.List",
				TfsdkName: "list_nested_attribute",
			},
		},
		"custom-type": {
			input: GeneratorListNestedAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "ListNestedAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "list_nested_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("list_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
