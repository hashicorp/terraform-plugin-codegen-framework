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
		name          string
		assocExtType  *AssocExtType
		fromFuncs     map[string]string
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
			fromFuncs: map[string]string{
				"bool_attribute": "BoolPointerValue",
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromObject(testCase.name, testCase.assocExtType, nil, testCase.fromFuncs)

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
		name          string
		assocExtType  *AssocExtType
		toFuncs       map[string]string
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
			toFuncs: map[string]string{
				"bool_attribute": "ValueBoolPointer",
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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			toFromObject := NewToFromObject(testCase.name, testCase.assocExtType, testCase.toFuncs, nil)

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
