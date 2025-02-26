// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustomObjectType_renderEqual(t *testing.T) {
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

return t.ObjectType.Equal(other.ObjectType)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomObjectType(testCase.name)

			got, err := customObjectType.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectType_renderString(t *testing.T) {
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

			customObjectType := NewCustomObjectType(testCase.name)

			got, err := customObjectType.renderString()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectType_renderTypable(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name:     "Example",
			expected: []byte(`var _ basetypes.ObjectTypable = ExampleType{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomObjectType(testCase.name)

			got, err := customObjectType.renderTypable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectType_renderType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`type ExampleType struct {
basetypes.ObjectType
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomObjectType(testCase.name)

			got, err := customObjectType.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectType_renderValueFromObject(t *testing.T) {
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
				"bool_attribute": "basetypes.ObjectValue",
			},
			expected: []byte(`
func (t ExampleType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
return ExampleValue{
ObjectValue: in,
}, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomObjectType(testCase.name)

			got, err := customObjectType.renderValueFromObject()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectType_renderValueFromTerraform(t *testing.T) {
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
attrValue, err := t.ObjectType.ValueFromTerraform(ctx, in)

if err != nil {
return nil, err
}

objectValue, ok := attrValue.(basetypes.ObjectValue)

if !ok {
return nil, fmt.Errorf("unexpected value type of %T", attrValue)
}

objectValuable, diags := t.ValueFromObject(ctx, objectValue)

if diags.HasError() {
return nil, fmt.Errorf("unexpected error converting ObjectValue to ObjectValuable: %v", diags)
}

return objectValuable, nil
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomObjectType(testCase.name)

			got, err := customObjectType.renderValueFromTerraform()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectType_renderValueType(t *testing.T) {
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

			customObjectType := NewCustomObjectType(testCase.name)

			got, err := customObjectType.renderValueType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderEqual(t *testing.T) {
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

return v.ObjectValue.Equal(other.ObjectValue)
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomObjectValue(testCase.name, testCase.elementType)

			got, err := customObjectValue.renderEqual()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderType(t *testing.T) {
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
ObjectType: basetypes.ObjectType{
AttrTypes: v.AttributeTypes(ctx),
},
}
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomObjectValue(testCase.name, testCase.elementType)

			got, err := customObjectValue.renderType()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderValuable(t *testing.T) {
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
			expected:    []byte(`var _ basetypes.ObjectValuable = ExampleValue{}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomObjectValue(testCase.name, testCase.elementType)

			got, err := customObjectValue.renderValuable()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderValue(t *testing.T) {
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
basetypes.ObjectValue
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomObjectValue(testCase.name, testCase.elementType)

			got, err := customObjectValue.renderValue()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
