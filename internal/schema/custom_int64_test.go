// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustomInt64Type_renderEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`func (t ExampleType) Equal(o attr.Type) bool {
other, ok := o.(ExampleType)

if !ok {
return false
}

return t.Int64Type.Equal(other.Int64Type)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Type := NewCustomInt64Type(testCase.name)

			got, err := customInt64Type.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Type_renderString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (t ExampleType) String() string {
return "ExampleType"
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Type := NewCustomInt64Type(testCase.name)

			got, err := customInt64Type.renderString()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Type_renderTypable(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:     "Example",
			expected: []byte(`var _ basetypes.Int64Typable = ExampleType{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Type := NewCustomInt64Type(testCase.name)

			got, err := customInt64Type.renderTypable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Type_renderType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`type ExampleType struct {
basetypes.Int64Type
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Type := NewCustomInt64Type(testCase.name)

			got, err := customInt64Type.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Type_renderValueFromInt64(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		attrValues    map[string]string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			attrValues: map[string]string{
				"bool_attribute": "basetypes.Int64Value",
			},
			expected: []byte(`
func (t ExampleType) ValueFromInt64(ctx context.Context, in basetypes.Int64Value) (basetypes.Int64Valuable, diag.Diagnostics) {
return ExampleValue{
Int64Value: in,
}, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Type := NewCustomInt64Type(testCase.name)

			got, err := customInt64Type.renderValueFromInt64()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Type_renderValueFromTerraform(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (t ExampleType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
attrValue, err := t.Int64Type.ValueFromTerraform(ctx, in)

if err != nil {
return nil, err
}

boolValue, ok := attrValue.(basetypes.Int64Value)

if !ok {
return nil, fmt.Errorf("unexpected value type of %T", attrValue)
}

boolValuable, diags := t.ValueFromInt64(ctx, boolValue)

if diags.HasError() {
return nil, fmt.Errorf("unexpected error converting Int64Value to Int64Valuable: %v", diags)
}

return boolValuable, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Type := NewCustomInt64Type(testCase.name)

			got, err := customInt64Type.renderValueFromTerraform()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Type_renderValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (t ExampleType) ValueType(ctx context.Context) attr.Value {
return ExampleValue{}
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Type := NewCustomInt64Type(testCase.name)

			got, err := customInt64Type.renderValueType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Value_renderEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		attrValues    map[string]string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			attrValues: map[string]string{
				"bool_attribute": "basetypes.Int64Value",
			},
			expected: []byte(`
func (v ExampleValue) Equal(o attr.Value) bool {
other, ok := o.(ExampleValue)

if !ok {
return false
}

return v.Int64Value.Equal(other.Int64Value)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Value := NewCustomInt64Value(testCase.name)

			got, err := customInt64Value.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Value_renderType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (v ExampleValue) Type(ctx context.Context) attr.Type {
return ExampleType{
}
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Value := NewCustomInt64Value(testCase.name)

			got, err := customInt64Value.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Value_renderValuable(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:     "Example",
			expected: []byte(`var _ basetypes.Int64Valuable = ExampleValue{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Value := NewCustomInt64Value(testCase.name)

			got, err := customInt64Value.renderValuable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomInt64Value_renderValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`type ExampleValue struct {
basetypes.Int64Value
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customInt64Value := NewCustomInt64Value(testCase.name)

			got, err := customInt64Value.renderValue()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
