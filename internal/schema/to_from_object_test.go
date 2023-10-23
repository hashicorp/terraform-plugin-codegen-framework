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
types.ObjectNull(v.AttributeTypes(ctx)),
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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromObject(testCase.name, testCase.assocExtType, testCase.attrTypesFromFuncs)

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

var apisdkType apisdk.Type

d := v.As(ctx, &apisdkType, basetypes.ObjectAsOptions{})

diags.Append(d...)

if diags.HasError() {
return nil, diags
}

return &apisdkType, diags
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromObject(testCase.name, testCase.assocExtType, testCase.attrTypesFromFuncs)

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
