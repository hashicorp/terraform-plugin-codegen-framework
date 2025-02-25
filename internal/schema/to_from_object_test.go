// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestToFromObject_renderFrom(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name               string
		assocExtType       *AssocExtType
		attrTypesFromFuncs map[string]string
		expected           []byte
		expectedError      error
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
			attrTypesFromFuncs: map[string]string{
				"bool":    "types.BoolPointerValue",
				"float64": "types.Float64PointerValue",
				"int64":   "types.Int64PointerValue",
				"number":  "types.NumberValue",
				"string":  "types.StringPointerValue",
			},
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return ExampleValue{
types.ObjectNull(v.AttributeTypes(ctx)),
}, diags
}

o, d := basetypes.NewObjectValue(v.AttributeTypes(ctx), map[string]attr.Value{
"bool": types.BoolPointerValue(apiObject.Bool),
"float64": types.Float64PointerValue(apiObject.Float64),
"int64": types.Int64PointerValue(apiObject.Int64),
"number": types.NumberValue(apiObject.Number),
"string": types.StringPointerValue(apiObject.String),
})

diags.Append(d...)

if diags.HasError() {
return ExampleValue{
types.ObjectUnknown(v.AttributeTypes(ctx)),
}, diags
}

return ExampleValue{
o,
}, diags
}
`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromObject(testCase.name, testCase.assocExtType, nil, testCase.attrTypesFromFuncs)

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

func TestToFromObject_renderTo(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name               string
		assocExtType       *AssocExtType
		attrTypesToFuncs   map[string]AttrTypesToFuncs
		attrTypesFromFuncs map[string]string
		expected           []byte
		expectedError      error
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
			attrTypesToFuncs: map[string]AttrTypesToFuncs{
				"bool": {
					AttrValue: "types.Bool",
					ToFunc:    "ValueBoolPointer",
				},
				"float64": {
					AttrValue: "types.Float64",
					ToFunc:    "ValueFloat64Pointer",
				},
				"int64": {
					AttrValue: "types.Int64",
					ToFunc:    "ValueInt64Pointer",
				},

				"number": {
					AttrValue: "types.Number",
					ToFunc:    "ValueBigFloat",
				},
				"string": {
					AttrValue: "types.String",
					ToFunc:    "ValueStringPointer",
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

attributes := v.Attributes()

boolAttribute, ok := attributes["bool"].(types.Bool)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue bool is unexpected type",
fmt.Sprintf(` + "`" + `"ExampleValue" bool is type of %T".` + "`" + `, attributes["bool"]),
))
}

float64Attribute, ok := attributes["float64"].(types.Float64)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue float64 is unexpected type",
fmt.Sprintf(` + "`" + `"ExampleValue" float64 is type of %T".` + "`" + `, attributes["float64"]),
))
}

int64Attribute, ok := attributes["int64"].(types.Int64)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue int64 is unexpected type",
fmt.Sprintf(` + "`" + `"ExampleValue" int64 is type of %T".` + "`" + `, attributes["int64"]),
))
}

numberAttribute, ok := attributes["number"].(types.Number)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue number is unexpected type",
fmt.Sprintf(` + "`" + `"ExampleValue" number is type of %T".` + "`" + `, attributes["number"]),
))
}

stringAttribute, ok := attributes["string"].(types.String)

if !ok {
diags.Append(diag.NewErrorDiagnostic(
"ExampleValue string is unexpected type",
fmt.Sprintf(` + "`" + `"ExampleValue" string is type of %T".` + "`" + `, attributes["string"]),
))
}

if diags.HasError() {
return nil, diags
}

apisdkType := apisdk.Type {
Bool: boolAttribute.ValueBoolPointer(),
Float64: float64Attribute.ValueFloat64Pointer(),
Int64: int64Attribute.ValueInt64Pointer(),
Number: numberAttribute.ValueBigFloat(),
String: stringAttribute.ValueStringPointer(),
}

return &apisdkType, diags
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromObject(testCase.name, testCase.assocExtType, testCase.attrTypesToFuncs, nil)

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
