// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestToFromNestedObject_renderFrom(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		assocExtType  *AssocExtType
		fromFuncs     map[string]ToFromConversion
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"bool_attribute": {
					Default: "BoolPointerValue",
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

return ExampleValue{
BoolAttribute: types.BoolPointerValue(apiObject.BoolAttribute),
state: attr.ValueStateKnown,
}, diags
}
`),
		},
		"default-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"type": {
					Default: "BoolPointerValue",
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

return ExampleValue{
ExampleType: types.BoolPointerValue(apiObject.Type),
state: attr.ValueStateKnown,
}, diags
}
`),
		},
		"nested-assoc-ext-type": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"bool_attribute": {
					AssocExtType: &AssocExtType{
						AssociatedExternalType: &schema.AssociatedExternalType{
							Type: "*api.BoolAttribute",
						},
					},
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

boolAttributeVal, d := BoolAttributeValue{}.FromApiBoolAttribute(ctx, apiObject.BoolAttribute)

diags.Append(d...)

if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
BoolAttribute: boolAttributeVal,
state: attr.ValueStateKnown,
}, diags
}
`),
		},
		"nested-assoc-ext-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"type": {
					AssocExtType: &AssocExtType{
						AssociatedExternalType: &schema.AssociatedExternalType{
							Type: "*api.BoolAttribute",
						},
					},
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

typeVal, d := TypeValue{}.FromApiBoolAttribute(ctx, apiObject.Type)

diags.Append(d...)

if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
ExampleType: typeVal,
state: attr.ValueStateKnown,
}, diags
}
`),
		},
		"collection-type": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"bool_attribute": {
					CollectionType: CollectionFields{
						ElementType:   "types.BoolType",
						TypeValueFrom: "types.ListValueFrom",
					},
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

boolAttributeVal, d := types.ListValueFrom(ctx, types.BoolType, apiObject.BoolAttribute)

diags.Append(d...)

if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
BoolAttribute: boolAttributeVal,
state: attr.ValueStateKnown,
}, diags
}
`),
		},
		"collection-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"type": {
					CollectionType: CollectionFields{
						ElementType:   "types.BoolType",
						TypeValueFrom: "types.ListValueFrom",
					},
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

typeVal, d := types.ListValueFrom(ctx, types.BoolType, apiObject.Type)

diags.Append(d...)

if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
ExampleType: typeVal,
state: attr.ValueStateKnown,
}, diags
}
`),
		},
		"object-type": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"object_attribute": {
					ObjectType: map[FrameworkIdentifier]ObjectField{
						FrameworkIdentifier("bool"): {
							Type:     "types.BoolType",
							FromFunc: "BoolPointerValue",
						},
						FrameworkIdentifier("float64"): {
							Type:     "types.Float64Type",
							FromFunc: "Float64PointerValue",
						},
						FrameworkIdentifier("int64"): {
							Type:     "types.Int64Type",
							FromFunc: "Int64PointerValue",
						},
						FrameworkIdentifier("number"): {
							Type:     "types.NumberType",
							FromFunc: "NumberValue",
						},
						FrameworkIdentifier("string"): {
							Type:     "types.StringType",
							FromFunc: "StringPointerValue",
						},
					},
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

objectAttributeVal, d := basetypes.NewObjectValue(
map[string]attr.Type{
"bool": types.BoolType,
"float64": types.Float64Type,
"int64": types.Int64Type,
"number": types.NumberType,
"string": types.StringType,
}, map[string]attr.Value{
"bool": types.BoolPointerValue(apiObject.ObjectAttribute.Bool),
"float64": types.Float64PointerValue(apiObject.ObjectAttribute.Float64),
"int64": types.Int64PointerValue(apiObject.ObjectAttribute.Int64),
"number": types.NumberValue(apiObject.ObjectAttribute.Number),
"string": types.StringPointerValue(apiObject.ObjectAttribute.String),
})

diags.Append(d...)

if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
ObjectAttribute: objectAttributeVal,
state: attr.ValueStateKnown,
}, diags
}
`),
		},
		"object-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				AssociatedExternalType: &schema.AssociatedExternalType{
					Type: "*apisdk.Type",
				},
			},
			fromFuncs: map[string]ToFromConversion{
				"type": {
					ObjectType: map[FrameworkIdentifier]ObjectField{
						FrameworkIdentifier("bool"): {
							Type:     "types.BoolType",
							FromFunc: "BoolPointerValue",
						},
						FrameworkIdentifier("float64"): {
							Type:     "types.Float64Type",
							FromFunc: "Float64PointerValue",
						},
						FrameworkIdentifier("int64"): {
							Type:     "types.Int64Type",
							FromFunc: "Int64PointerValue",
						},
						FrameworkIdentifier("number"): {
							Type:     "types.NumberType",
							FromFunc: "NumberValue",
						},
						FrameworkIdentifier("string"): {
							Type:     "types.StringType",
							FromFunc: "StringPointerValue",
						},
					},
				},
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return NewExampleValueNull(), diags
}

typeVal, d := basetypes.NewObjectValue(
map[string]attr.Type{
"bool": types.BoolType,
"float64": types.Float64Type,
"int64": types.Int64Type,
"number": types.NumberType,
"string": types.StringType,
}, map[string]attr.Value{
"bool": types.BoolPointerValue(apiObject.Type.Bool),
"float64": types.Float64PointerValue(apiObject.Type.Float64),
"int64": types.Int64PointerValue(apiObject.Type.Int64),
"number": types.NumberValue(apiObject.Type.Number),
"string": types.StringPointerValue(apiObject.Type.String),
})

diags.Append(d...)

if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
ExampleType: typeVal,
state: attr.ValueStateKnown,
}, diags
}
`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromNestedObject(testCase.name, testCase.assocExtType, nil, testCase.fromFuncs)

			got, err := toFromObject.renderFrom()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestToFromNestedObject_renderTo(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		assocExtType  *AssocExtType
		toFuncs       map[string]ToFromConversion
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"bool_attribute": {
					Default: "ValueBoolPointer",
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

return &apisdk.Type{
BoolAttribute: v.BoolAttribute.ValueBoolPointer(),
}, diags
}`),
		},
		"default-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"type": {
					Default: "ValueBoolPointer",
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

return &apisdk.Type{
Type: v.ExampleType.ValueBoolPointer(),
}, diags
}`),
		},
		"nested-assoc-ext-type": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"bool_attribute": {
					AssocExtType: &AssocExtType{
						AssociatedExternalType: &schema.AssociatedExternalType{
							Type: "*api.BoolAttribute",
						},
					},
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

apiBoolAttribute, d := v.BoolAttribute.ToApiBoolAttribute(ctx)

diags.Append(d...)

if diags.HasError() {
return nil, diags
}

return &apisdk.Type{
BoolAttribute: apiBoolAttribute,
}, diags
}`),
		},
		"nested-assoc-ext-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"type": {
					AssocExtType: &AssocExtType{
						AssociatedExternalType: &schema.AssociatedExternalType{
							Type: "*api.BoolAttribute",
						},
					},
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

apiBoolAttribute, d := v.ExampleType.ToApiBoolAttribute(ctx)

diags.Append(d...)

if diags.HasError() {
return nil, diags
}

return &apisdk.Type{
Type: apiBoolAttribute,
}, diags
}`),
		},
		"collection-type": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"bool_attribute": {
					CollectionType: CollectionFields{
						GoType: "[]*bool",
					},
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

var boolAttributeField []*bool

d := v.BoolAttribute.ElementsAs(ctx, &boolAttributeField, false)

diags.Append(d...)

if diags.HasError() {
return nil, diags
}

return &apisdk.Type{
BoolAttribute: boolAttributeField,
}, diags
}`),
		},
		"collection-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"type": {
					CollectionType: CollectionFields{
						GoType: "[]*bool",
					},
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

var typeField []*bool

d := v.ExampleType.ElementsAs(ctx, &typeField, false)

diags.Append(d...)

if diags.HasError() {
return nil, diags
}

return &apisdk.Type{
Type: typeField,
}, diags
}`),
		},
		"object-type": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"object_attribute": {
					ObjectType: map[FrameworkIdentifier]ObjectField{
						FrameworkIdentifier("bool"): {
							GoType: "*bool",
							Type:   "types.Bool",
							ToFunc: "ValueBoolPointer",
						},
						FrameworkIdentifier("float64"): {
							GoType: "*float64",
							Type:   "types.Float64",
							ToFunc: "ValueFloat64Pointer",
						},
						FrameworkIdentifier("int64"): {
							GoType: "*int64",
							Type:   "types.Int64",
							ToFunc: "ValueInt64Pointer",
						},
						FrameworkIdentifier("number"): {
							GoType: "*big.Float",
							Type:   "types.Number",
							ToFunc: "ValueBigFloat",
						},
						FrameworkIdentifier("string"): {
							GoType: "*string",
							Type:   "types.String",
							ToFunc: "ValueStringPointer",
						},
					},
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

attributes := v.ObjectAttribute.Attributes()

objectAttributeFieldBool, ok := attributes["bool"].(types.Bool)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ObjectAttribute Field bool Is Wrong Type",
fmt.Sprintf(` + "`" + `ObjectAttribute field bool expected to be types.Bool, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

objectAttributeFieldFloat64, ok := attributes["float64"].(types.Float64)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ObjectAttribute Field float64 Is Wrong Type",
fmt.Sprintf(` + "`" + `ObjectAttribute field float64 expected to be types.Float64, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

objectAttributeFieldInt64, ok := attributes["int64"].(types.Int64)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ObjectAttribute Field int64 Is Wrong Type",
fmt.Sprintf(` + "`" + `ObjectAttribute field int64 expected to be types.Int64, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

objectAttributeFieldNumber, ok := attributes["number"].(types.Number)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ObjectAttribute Field number Is Wrong Type",
fmt.Sprintf(` + "`" + `ObjectAttribute field number expected to be types.Number, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

objectAttributeFieldString, ok := attributes["string"].(types.String)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ObjectAttribute Field string Is Wrong Type",
fmt.Sprintf(` + "`" + `ObjectAttribute field string expected to be types.String, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

return &apisdk.Type{
ObjectAttribute: struct {
Bool *bool
Float64 *float64
Int64 *int64
Number *big.Float
String *string
}{
Bool: objectAttributeFieldBool.ValueBoolPointer(),
Float64: objectAttributeFieldFloat64.ValueFloat64Pointer(),
Int64: objectAttributeFieldInt64.ValueInt64Pointer(),
Number: objectAttributeFieldNumber.ValueBigFloat(),
String: objectAttributeFieldString.ValueStringPointer(),
},
}, diags
}`),
		},
		"object-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			assocExtType: &AssocExtType{
				&schema.AssociatedExternalType{
					Import: &code.Import{
						Path: "example.com/apisdk",
					},
					Type: "*apisdk.Type",
				},
			},
			toFuncs: map[string]ToFromConversion{
				"type": {
					ObjectType: map[FrameworkIdentifier]ObjectField{
						FrameworkIdentifier("bool"): {
							GoType: "*bool",
							Type:   "types.Bool",
							ToFunc: "ValueBoolPointer",
						},
						FrameworkIdentifier("float64"): {
							GoType: "*float64",
							Type:   "types.Float64",
							ToFunc: "ValueFloat64Pointer",
						},
						FrameworkIdentifier("int64"): {
							GoType: "*int64",
							Type:   "types.Int64",
							ToFunc: "ValueInt64Pointer",
						},
						FrameworkIdentifier("number"): {
							GoType: "*big.Float",
							Type:   "types.Number",
							ToFunc: "ValueBigFloat",
						},
						FrameworkIdentifier("string"): {
							GoType: "*string",
							Type:   "types.String",
							ToFunc: "ValueStringPointer",
						},
					},
				},
			},
			expected: []byte(`func (v ExampleValue) ToApisdkType(ctx context.Context) (*apisdk.Type, diag.Diagnostics) {
var diags diag.Diagnostics

if v.IsNull() {
return nil, diags
}

if v.IsUnknown() {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue Value Is Unknown",
` + "`" + `"ExampleValue" is unknown.` + "`" + `,
))

return nil, diags
}

attributes := v.ExampleType.Attributes()

typeFieldBool, ok := attributes["bool"].(types.Bool)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleType Field bool Is Wrong Type",
fmt.Sprintf(` + "`" + `ExampleType field bool expected to be types.Bool, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

typeFieldFloat64, ok := attributes["float64"].(types.Float64)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleType Field float64 Is Wrong Type",
fmt.Sprintf(` + "`" + `ExampleType field float64 expected to be types.Float64, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

typeFieldInt64, ok := attributes["int64"].(types.Int64)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleType Field int64 Is Wrong Type",
fmt.Sprintf(` + "`" + `ExampleType field int64 expected to be types.Int64, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

typeFieldNumber, ok := attributes["number"].(types.Number)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleType Field number Is Wrong Type",
fmt.Sprintf(` + "`" + `ExampleType field number expected to be types.Number, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

typeFieldString, ok := attributes["string"].(types.String)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleType Field string Is Wrong Type",
fmt.Sprintf(` + "`" + `ExampleType field string expected to be types.String, was: %T` + "`" + `, attributes["bool"]),
))

return nil, diags
}

return &apisdk.Type{
Type: struct {
Bool *bool
Float64 *float64
Int64 *int64
Number *big.Float
String *string
}{
Bool: typeFieldBool.ValueBoolPointer(),
Float64: typeFieldFloat64.ValueFloat64Pointer(),
Int64: typeFieldInt64.ValueInt64Pointer(),
Number: typeFieldNumber.ValueBigFloat(),
String: typeFieldString.ValueStringPointer(),
},
}, diags
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromNestedObject(testCase.name, testCase.assocExtType, testCase.toFuncs, nil)

			got, err := toFromObject.renderTo()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
