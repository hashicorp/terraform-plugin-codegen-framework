// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestToFromBool_renderFrom(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		assocExtType  *AssocExtType
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
			expected: []byte(`
func (v ExampleValue) FromApisdkType(ctx context.Context, apiObject *apisdk.Type) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

if apiObject == nil {
return ExampleValue{
types.BoolNull(),
}, diags
}

return ExampleValue{
types.BoolPointerValue(*apiObject),
}, diags
}
`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromBool := NewToFromBool(testCase.name, testCase.assocExtType)

			got, err := toFromBool.renderFrom()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestToFromBool_renderTo(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		assocExtType  *AssocExtType
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

a := apisdk.Type(v.ValueBoolPointer())

return &a, diags
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromBool := NewToFromBool(testCase.name, testCase.assocExtType)

			got, err := toFromBool.renderTo()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
