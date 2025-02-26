// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustomListType_renderEqual(t *testing.T) {
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

return t.ListType.Equal(other.ListType)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListType := NewCustomListType(testCase.name)

			got, err := customListType.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListType_renderString(t *testing.T) {
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

			customListType := NewCustomListType(testCase.name)

			got, err := customListType.renderString()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListType_renderTypable(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:     "Example",
			expected: []byte(`var _ basetypes.ListTypable = ExampleType{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListType := NewCustomListType(testCase.name)

			got, err := customListType.renderTypable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListType_renderType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`type ExampleType struct {
basetypes.ListType
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListType := NewCustomListType(testCase.name)

			got, err := customListType.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListType_renderValueFromList(t *testing.T) {
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
				"bool_attribute": "basetypes.ListValue",
			},
			expected: []byte(`
func (t ExampleType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
return ExampleValue{
ListValue: in,
}, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListType := NewCustomListType(testCase.name)

			got, err := customListType.renderValueFromList()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListType_renderValueFromTerraform(t *testing.T) {
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
attrValue, err := t.ListType.ValueFromTerraform(ctx, in)

if err != nil {
return nil, err
}

listValue, ok := attrValue.(basetypes.ListValue)

if !ok {
return nil, fmt.Errorf("unexpected value type of %T", attrValue)
}

listValuable, diags := t.ValueFromList(ctx, listValue)

if diags.HasError() {
return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
}

return listValuable, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListType := NewCustomListType(testCase.name)

			got, err := customListType.renderValueFromTerraform()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListType_renderValueType(t *testing.T) {
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

			customListType := NewCustomListType(testCase.name)

			got, err := customListType.renderValueType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListValue_renderEqual(t *testing.T) {
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

return v.ListValue.Equal(other.ListValue)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListValue := NewCustomListValue(testCase.name, testCase.elementType)

			got, err := customListValue.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListValue_renderType(t *testing.T) {
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
ListType: basetypes.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListValue := NewCustomListValue(testCase.name, testCase.elementType)

			got, err := customListValue.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListValue_renderValuable(t *testing.T) {
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
			expected:    []byte(`var _ basetypes.ListValuable = ExampleValue{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListValue := NewCustomListValue(testCase.name, testCase.elementType)

			got, err := customListValue.renderValuable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomListValue_renderValue(t *testing.T) {
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
basetypes.ListValue
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customListValue := NewCustomListValue(testCase.name, testCase.elementType)

			got, err := customListValue.renderValue()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
