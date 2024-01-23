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

func TestGeneratorMapNestedAttribute_New(t *testing.T) {
	t.Parallel()

	attributes, err := NewAttributes(resource.Attributes{})

	if err != nil {
		t.Error(err)
	}

	testCases := map[string]struct {
		input         *resource.MapNestedAttribute
		expected      GeneratorMapNestedAttribute
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
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
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
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"attributes-map-nested-bool": {
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
			expected: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorMapNestedAttribute{
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
							PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
						},
					},
				},
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorMapNestedAttribute{
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
							PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
							ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, specschema.CustomPlanModifiers{}),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, specschema.CustomValidators{}),
					"name",
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"attributes-object-bool": {
			input: &resource.MapNestedAttribute{
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
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
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
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"computed": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{})},
		},
		"required": {
			input: &resource.MapNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &resource.MapNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, nil),
			},
		},
		"deprecation_message": {
			input: &resource.MapNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, nil),
			},
		},
		"description": {
			input: &resource.MapNestedAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, nil),
			},
		},
		"sensitive": {
			input: &resource.MapNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, nil),
			},
		},
		"validators": {
			input: &resource.MapNestedAttribute{
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
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
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
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{
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
			input: &resource.MapNestedAttribute{
				PlanModifiers: specschema.MapPlanModifiers{
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
			expected: GeneratorMapNestedAttribute{
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
				PlanModifiers: specschema.MapPlanModifiers{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_planmodifier",
							},
						},
						SchemaDefinition: "my_planmodifier.Modify()",
					},
				}),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeMap, nil),
			},
		},
		"default": {
			input: &resource.MapNestedAttribute{
				Default: &specschema.MapDefault{
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
			expected: GeneratorMapNestedAttribute{
				Default: &specschema.MapDefault{
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
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{}),
				ValidatorsCustom:    convert.NewValidatorsCustom(convert.ValidatorTypeMap, nil),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorMapNestedAttribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorMapNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorMapNestedAttribute
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
		"nested-map": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"map": GeneratorMapAttribute{
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
		"nested-map-with-custom-type": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"map": GeneratorMapAttribute{
							CustomTypeCollection: convert.NewCustomTypeCollection(
								&specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/nested_map",
									},
								},
								nil,
								convert.CustomCollectionTypeMap,
								"",
								"map",
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
					Path: "github.com/my_account/my_project/nested_map",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-map-with-custom-type-with-element-with-custom-type": {
			input: GeneratorMapNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"map": GeneratorMapAttribute{
							CustomTypeCollection: convert.NewCustomTypeCollection(
								&specschema.CustomType{
									Import: &code.Import{
										Path: "github.com/my_account/my_project/nested_map",
									},
								},
								nil,
								convert.CustomCollectionTypeMap,
								"",
								"map",
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
					Path: "github.com/my_account/my_project/nested_map",
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
				Validators: specschema.MapValidators{
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
			input: GeneratorMapNestedAttribute{
				Validators: specschema.MapValidators{
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
			input: GeneratorMapNestedAttribute{
				Validators: specschema.MapValidators{
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
			input: GeneratorMapNestedAttribute{
				Validators: specschema.MapValidators{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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
			input: GeneratorMapNestedAttribute{
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

func TestGeneratorMapNestedAttribute_Schema(t *testing.T) {
	t.Parallel()

	attributeName := "map_nested_attribute"

	testCases := map[string]struct {
		input         GeneratorMapNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorMapNestedAttribute{
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
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorMapNestedAttribute{
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
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-map-nested": {
			input: GeneratorMapNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{
						"nested_map_nested": GeneratorMapNestedAttribute{
							NestedAttributeObject: NewNestedAttributeObject(
								generatorschema.GeneratorAttributes{
									"bool": GeneratorBoolAttribute{
										ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
									},
								},
								nil,
								convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
								convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
								"nested_map_nested",
							),
						},
					},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_map_nested": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: NestedMapNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedMapNestedValue{}.AttributeTypes(ctx),
},
},
},
},
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorMapNestedAttribute{
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
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorMapNestedAttribute{
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
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
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
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorMapNestedAttribute{
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
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorMapNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"map_nested_attribute",
				),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorMapNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"map_nested_attribute",
				),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorMapNestedAttribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"map_nested_attribute",
				),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorMapNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"map_nested_attribute",
				),
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorMapNestedAttribute{
				Description: convert.NewDescription(pointer("description")),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"map_nested_attribute",
				),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorMapNestedAttribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.PlanModifiersCustom{},
					convert.ValidatorsCustom{},
					"map_nested_attribute",
				),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorMapNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
				ValidatorsCustom: convert.NewValidatorsCustom(convert.ValidatorTypeMap, specschema.CustomValidators{
					&specschema.CustomValidator{
						SchemaDefinition: "my_validator.Validate()",
					},
					&specschema.CustomValidator{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Validators: []validator.Map{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorMapNestedAttribute{
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
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorMapNestedAttribute{
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
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
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
			input: GeneratorMapNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
				PlanModifiersCustom: convert.NewPlanModifiersCustom(convert.PlanModifierTypeMap, specschema.CustomPlanModifiers{
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_plan_modifier.Modify()",
					},
					&specschema.CustomPlanModifier{
						SchemaDefinition: "my_other_plan_modifier.Modify()",
					},
				}),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
PlanModifiers: []planmodifier.Map{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorMapNestedAttribute{
				NestedAttributeObject: NewNestedAttributeObject(
					generatorschema.GeneratorAttributes{},
					nil,
					convert.NewPlanModifiersCustom(convert.PlanModifierTypeObject, nil),
					convert.NewValidatorsCustom(convert.ValidatorTypeObject, nil),
					attributeName,
				),
				DefaultCustom: convert.NewDefaultCustom(&specschema.CustomDefault{
					SchemaDefinition: "my_map_default.Default()",
				}),
			},
			expected: `"map_nested_attribute": schema.MapNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: MapNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Default: my_map_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("map_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorMapNestedAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorMapNestedAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "MapNestedAttribute",
				ValueType: "types.Map",
				TfsdkName: "map_nested_attribute",
			},
		},
		"custom-type": {
			input: GeneratorMapNestedAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "MapNestedAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "map_nested_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("map_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
