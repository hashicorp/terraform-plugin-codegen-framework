// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorListAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *provider.ListAttribute
		expected      GeneratorListAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*provider.ListAttribute is nil"),
		},
		"element-type-bool": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-string": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-list-string": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-map-string": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-list-object-string": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-object-string": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"element-type-object-list-string": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &provider.ListAttribute{
				OptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &provider.ListAttribute{
				OptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorListAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"deprecation_message": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &provider.ListAttribute{
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
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &provider.ListAttribute{
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
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeList, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &provider.ListAttribute{
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
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorListAttribute{
				CustomType: convert.NewCustomTypeCollection(
					nil,
					&specschema.AssociatedExternalType{Type: "*api.ListAttribute"},
					convert.CustomCollectionTypeList,
					"types.StringType",
					"list_attribute",
				),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
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
					&specschema.CustomType{Type: "my_custom_type"},
					&specschema.AssociatedExternalType{Type: "*api.ListAttribute"},
					convert.CustomCollectionTypeList,
					"types.StringType",
					"name",
				),
			},
			expected: `"list_attribute": schema.ListAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorListAttribute{
				OptionalRequired: convert.NewOptionalRequired(specschema.Required),
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
				OptionalRequired: convert.NewOptionalRequired(specschema.Optional),
				ElementTypeCollection: convert.NewElementType(specschema.ElementType{
					String: &specschema.StringType{},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
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
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
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
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
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
							Type: "numberCustomType",
						},
					},
				}),
			},
			expected: `"list_attribute": schema.ListAttribute{
ElementType: numberCustomType,
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
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
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
