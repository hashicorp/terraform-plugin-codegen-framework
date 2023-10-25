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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
