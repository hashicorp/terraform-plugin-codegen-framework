// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustomMapType_renderEqual(t *testing.T) {
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

return t.MapType.Equal(other.MapType)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapType := NewCustomMapType(testCase.name)

			got, err := customMapType.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapType_renderString(t *testing.T) {
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

			customMapType := NewCustomMapType(testCase.name)

			got, err := customMapType.renderString()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapType_renderTypable(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:     "Example",
			expected: []byte(`var _ basetypes.MapTypable = ExampleType{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapType := NewCustomMapType(testCase.name)

			got, err := customMapType.renderTypable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapType_renderType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`type ExampleType struct {
basetypes.MapType
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapType := NewCustomMapType(testCase.name)

			got, err := customMapType.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapType_renderValueFromMap(t *testing.T) {
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
				"bool_attribute": "basetypes.MapValue",
			},
			expected: []byte(`
func (t ExampleType) ValueFromMap(ctx context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
return ExampleValue{
MapValue: in,
}, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapType := NewCustomMapType(testCase.name)

			got, err := customMapType.renderValueFromMap()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapType_renderValueFromTerraform(t *testing.T) {
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
attrValue, err := t.MapType.ValueFromTerraform(ctx, in)

if err != nil {
return nil, err
}

mapValue, ok := attrValue.(basetypes.MapValue)

if !ok {
return nil, fmt.Errorf("unexpected value type of %T", attrValue)
}

mapValuable, diags := t.ValueFromMap(ctx, mapValue)

if diags.HasError() {
return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
}

return mapValuable, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapType := NewCustomMapType(testCase.name)

			got, err := customMapType.renderValueFromTerraform()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapType_renderValueType(t *testing.T) {
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

			customMapType := NewCustomMapType(testCase.name)

			got, err := customMapType.renderValueType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapValue_renderEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		elementType   string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:        "Example",
			elementType: "types.BoolType",
			expected: []byte(`
func (v ExampleValue) Equal(o attr.Value) bool {
other, ok := o.(ExampleValue)

if !ok {
return false
}

return v.MapValue.Equal(other.MapValue)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapValue := NewCustomMapValue(testCase.name, testCase.elementType)

			got, err := customMapValue.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapValue_renderType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		elementType   string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:        "Example",
			elementType: "types.BoolType",
			expected: []byte(`
func (v ExampleValue) Type(ctx context.Context) attr.Type {
return ExampleType{
MapType: basetypes.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapValue := NewCustomMapValue(testCase.name, testCase.elementType)

			got, err := customMapValue.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapValue_renderValuable(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		elementType   string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:        "Example",
			elementType: "types.BoolType",
			expected:    []byte(`var _ basetypes.MapValuable = ExampleValue{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapValue := NewCustomMapValue(testCase.name, testCase.elementType)

			got, err := customMapValue.renderValuable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomMapValue_renderValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		elementType   string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:        "Example",
			elementType: "types.BoolType",
			expected: []byte(`type ExampleValue struct {
basetypes.MapValue
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customMapValue := NewCustomMapValue(testCase.name, testCase.elementType)

			got, err := customMapValue.renderValue()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
