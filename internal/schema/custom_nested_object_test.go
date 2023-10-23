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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

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

func TestCustomObjectType_renderValue(t *testing.T) {
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
				"bool_attribute": "basetypes.BoolValue",
			},
			expected: []byte(`
func NewExampleValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ExampleValue, diag.Diagnostics) {
var diags diag.Diagnostics

// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
ctx := context.Background()

for name, attributeType := range attributeTypes {
attribute, ok := attributes[name]

if !ok {
diags.AddError(
"Missing ExampleValue Attribute Value",
"While creating a ExampleValue value, a missing attribute value was detected. "+
"A ExampleValue must contain values for all attributes, even if null or unknown. "+
"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
fmt.Sprintf("ExampleValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
)

continue
}

if !attributeType.Equal(attribute.Type(ctx)) {
diags.AddError(
"Invalid ExampleValue Attribute Type",
"While creating a ExampleValue value, an invalid attribute value was detected. "+
"A ExampleValue must use a matching attribute type for the value. "+
"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
fmt.Sprintf("ExampleValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
fmt.Sprintf("ExampleValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
)
}
}

for name := range attributes {
_, ok := attributeTypes[name]

if !ok {
diags.AddError(
"Extra ExampleValue Attribute Value",
"While creating a ExampleValue value, an extra attribute value was detected. "+
"A ExampleValue must not contain values beyond the expected attribute types. "+
"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
fmt.Sprintf("Extra ExampleValue Attribute Name: %s", name),
)
}
}

if diags.HasError() {
return NewExampleValueUnknown(), diags
}


boolAttributeAttribute, ok := attributes["bool_attribute"]

if !ok {
diags.AddError(
"Attribute Missing",
` + "`bool_attribute is missing from object`" + `)

return NewExampleValueUnknown(), diags
}

boolAttributeVal, ok := boolAttributeAttribute.(basetypes.BoolValue)

if !ok {
diags.AddError(
"Attribute Wrong Type",
fmt.Sprintf(` + "`bool_attribute expected to be basetypes.BoolValue, was: %T`" + `, boolAttributeAttribute))
}


if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
BoolAttribute: boolAttributeVal,
state: attr.ValueStateKnown,
}, diags
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, testCase.attrValues)

			got, err := customObjectType.renderValue()

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
				"bool_attribute": "basetypes.BoolValue",
			},
			expected: []byte(`
func (t ExampleType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
var diags diag.Diagnostics

attributes := in.Attributes()


boolAttributeAttribute, ok := attributes["bool_attribute"]

if !ok {
diags.AddError(
"Attribute Missing",
` + "`bool_attribute is missing from object`" + `)

return nil, diags
}

boolAttributeVal, ok := boolAttributeAttribute.(basetypes.BoolValue)

if !ok {
diags.AddError(
"Attribute Wrong Type",
fmt.Sprintf(` + "`bool_attribute expected to be basetypes.BoolValue, was: %T`" + `, boolAttributeAttribute))
}


if diags.HasError() {
return nil, diags
}

return ExampleValue{
BoolAttribute: boolAttributeVal,
state: attr.ValueStateKnown,
}, diags
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, testCase.attrValues)

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
if in.Type() == nil {
return NewExampleValueNull(), nil
}

if !in.Type().Equal(t.TerraformType(ctx)) {
return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
}

if !in.IsKnown() {
return NewExampleValueUnknown(), nil
}

if in.IsNull() {
return NewExampleValueNull(), nil
}

attributes := map[string]attr.Value{}

val := map[string]tftypes.Value{}

err := in.As(&val)

if err != nil {
return nil, err
}

for k, v := range val {
a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

if err != nil {
return nil, err
}

attributes[k] = a
}

return NewExampleValueMust(t.AttrTypes, attributes), nil
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

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

func TestCustomObjectType_renderValueMust(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func NewExampleValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ExampleValue {
object, diags := NewExampleValue(attributeTypes, attributes)

if diags.HasError() {
// This could potentially be added to the diag package.
diagsStrings := make([]string, 0, len(diags))

for _, diagnostic := range diags {
diagsStrings = append(diagsStrings, fmt.Sprintf(
"%s | %s | %s",
diagnostic.Severity(),
diagnostic.Summary(),
diagnostic.Detail()))
}

panic("NewExampleValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
}

return object
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

			got, err := customObjectType.renderValueMust()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectType_renderValueNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func NewExampleValueNull() ExampleValue {
return ExampleValue{
state: attr.ValueStateNull,
}
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

			got, err := customObjectType.renderValueNull()

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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

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

func TestCustomObjectType_renderValueUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func NewExampleValueUnknown() ExampleValue {
return ExampleValue{
state: attr.ValueStateUnknown,
}
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectType := NewCustomNestedObjectType(testCase.name, nil)

			got, err := customObjectType.renderValueUnknown()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderAttributeTypes(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		attrTypes     map[string]string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			attrTypes: map[string]string{
				"bool_attribute": "basetypes.BoolType{}",
			},
			expected: []byte(`
func (v ExampleValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool_attribute": basetypes.BoolType{},
}
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, testCase.attrTypes, nil)

			got, err := customObjectValue.renderAttributeTypes()

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
		attrValues    map[string]string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			attrValues: map[string]string{
				"bool_attribute": "basetypes.BoolValue",
			},
			expected: []byte(`
func (v ExampleValue) Equal(o attr.Value) bool {
other, ok := o.(ExampleValue)

if !ok {
return false
}

if v.state != other.state {
return false
}

if v.state != attr.ValueStateKnown {
return true
}


if !v.BoolAttribute.Equal(other.BoolAttribute) {
return false
}


return true
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, testCase.attrValues)

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

func TestCustomObjectValue_renderIsNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (v ExampleValue) IsNull() bool {
return v.state == attr.ValueStateNull
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil)

			got, err := customObjectValue.renderIsNull()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderIsUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (v ExampleValue) IsUnknown() bool {
return v.state == attr.ValueStateUnknown
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil)

			got, err := customObjectValue.renderIsUnknown()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (v ExampleValue) String() string {
return "ExampleValue"
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil)

			got, err := customObjectValue.renderString()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderToObjectValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name           string
		attributeTypes map[string]string
		attrTypes      map[string]string
		expected       []byte
		expectedError  error
	}{
		"default": {
			name: "Example",
			attributeTypes: map[string]string{
				"bool_attribute": "Bool",
			},
			attrTypes: map[string]string{
				"bool_attribute": "basetypes.BoolType{}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
objVal, diags := types.ObjectValue(
map[string]attr.Type{
"bool_attribute": basetypes.BoolType{},
},
map[string]attr.Value{
"bool_attribute": v.BoolAttribute,
})

return objVal, diags
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, testCase.attributeTypes, testCase.attrTypes, nil)

			got, err := customObjectValue.renderToObjectValue()

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomObjectValue_renderToTerraformValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name          string
		attrTypes     map[string]string
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			attrTypes: map[string]string{
				"bool_attribute": "basetypes.BoolType{}",
			},
			expected: []byte(`func (v ExampleValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
attrTypes := make(map[string]tftypes.Type, 1)

var val tftypes.Value
var err error


attrTypes["bool_attribute"] = basetypes.BoolType{}.TerraformType(ctx)

objectType := tftypes.Object{AttributeTypes: attrTypes}

switch v.state {
case attr.ValueStateKnown:
vals := make(map[string]tftypes.Value, 1)


val, err = v.BoolAttribute.ToTerraformValue(ctx)

if err != nil {
return tftypes.NewValue(objectType, tftypes.UnknownValue), err
}

vals["bool_attribute"] = val



if err := tftypes.ValidateValue(objectType, vals); err != nil {
return tftypes.NewValue(objectType, tftypes.UnknownValue), err
}

return tftypes.NewValue(objectType, vals), nil
case attr.ValueStateNull:
return tftypes.NewValue(objectType, nil), nil
case attr.ValueStateUnknown:
return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
default:
panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
}
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, testCase.attrTypes, nil)

			got, err := customObjectValue.renderToTerraformValue()

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
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`
func (v ExampleValue) Type(ctx context.Context) attr.Type {
return ExampleType{
basetypes.ObjectType{
AttrTypes: v.AttributeTypes(ctx),
},
}
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil)

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
		expected      []byte
		expectedError error
	}{
		"default": {
			name:     "Example",
			expected: []byte(`var _ basetypes.ObjectValuable = ExampleValue{}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil)

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
		expected      []byte
		expectedError error
	}{
		"default": {
			name: "Example",
			expected: []byte(`type ExampleValue struct {
state attr.ValueState
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil)

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

var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return x.Error() == y.Error()
})
