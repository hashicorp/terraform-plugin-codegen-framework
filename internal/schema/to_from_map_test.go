// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestToFromMap_renderFrom(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		assocExtType  *AssocExtType
		elemTypeType  string
		elemTypeValue string
		elemFrom      string
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
			elemTypeType:  "types.BoolType",
			elemTypeValue: "types.Bool",
			elemFrom:      "types.BoolPointerValue",
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return ExampleValue{
types.MapNull(types.BoolType),
}, diags
}

elems := make(map[string]types.Bool)

for k, e := range *apiObject {
elems[k] = types.BoolPointerValue(e)
}

l, d := basetypes.NewMapValueFrom(ctx, types.BoolType, elems)

diags.Append(d...)

if diags.HasError() {
return ExampleValue{
types.MapUnknown(types.BoolType),
}, diags
}

return ExampleValue{
l,
}, diags
}
`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromMap := NewToFromMap(testCase.name, testCase.assocExtType, testCase.elemTypeType, testCase.elemTypeValue, testCase.elemFrom)

			got, err := toFromMap.renderFrom()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestToFromMap_renderTo(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		assocExtType  *AssocExtType
		elemTypeType  string
		elemTypeValue string
		elemFrom      string
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
			elemTypeType:  "types.BoolType",
			elemTypeValue: "types.Bool",
			elemFrom:      "types.BoolPointerValue",
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

d := v.ElementsAs(ctx, &apisdkType, false)

diags.Append(d...)

if diags.HasError() {
return nil, diags
}

return &apisdkType, diags
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromMap := NewToFromMap(testCase.name, testCase.assocExtType, testCase.elemTypeType, testCase.elemTypeValue, testCase.elemFrom)

			got, err := toFromMap.renderTo()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
