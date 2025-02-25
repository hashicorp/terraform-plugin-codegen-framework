// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustomNestedObjectType_renderEqual(t *testing.T) {
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

func TestCustomNestedObjectType_renderString(t *testing.T) {
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

func TestCustomNestedObjectType_renderTypable(t *testing.T) {
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

func TestCustomNestedObjectType_renderType(t *testing.T) {
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

func TestCustomNestedObjectType_renderValue(t *testing.T) {
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
		"attribute-name-same-as-generated-method-name": {
			name: "Example",
			attrValues: map[string]string{
				"type": "basetypes.BoolValue",
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


typeAttribute, ok := attributes["type"]

if !ok {
diags.AddError(
"Attribute Missing",
` + "`type is missing from object`" + `)

return NewExampleValueUnknown(), diags
}

typeVal, ok := typeAttribute.(basetypes.BoolValue)

if !ok {
diags.AddError(
"Attribute Wrong Type",
fmt.Sprintf(` + "`type expected to be basetypes.BoolValue, was: %T`" + `, typeAttribute))
}


if diags.HasError() {
return NewExampleValueUnknown(), diags
}

return ExampleValue{
ExampleType: typeVal,
state: attr.ValueStateKnown,
}, diags
}`),
		},
	}

	for name, testCase := range testCases {

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

func TestCustomNestedObjectType_renderValueFromObject(t *testing.T) {
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
		"attribute-name-same-as-generated-method-name": {
			name: "Example",
			attrValues: map[string]string{
				"type": "basetypes.BoolValue",
			},
			expected: []byte(`
func (t ExampleType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
var diags diag.Diagnostics

attributes := in.Attributes()


typeAttribute, ok := attributes["type"]

if !ok {
diags.AddError(
"Attribute Missing",
` + "`type is missing from object`" + `)

return nil, diags
}

typeVal, ok := typeAttribute.(basetypes.BoolValue)

if !ok {
diags.AddError(
"Attribute Wrong Type",
fmt.Sprintf(` + "`type expected to be basetypes.BoolValue, was: %T`" + `, typeAttribute))
}


if diags.HasError() {
return nil, diags
}

return ExampleValue{
ExampleType: typeVal,
state: attr.ValueStateKnown,
}, diags
}`),
		},
		"no-attributes": {
			name: "Example",
			expected: []byte(`
func (t ExampleType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
var diags diag.Diagnostics



if diags.HasError() {
return nil, diags
}

return ExampleValue{
state: attr.ValueStateKnown,
}, diags
}`),
		},
	}

	for name, testCase := range testCases {

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

func TestCustomNestedObjectType_renderValueFromTerraform(t *testing.T) {
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

return NewExampleValueMust(ExampleValue{}.AttributeTypes(ctx), attributes), nil
}`),
		},
	}

	for name, testCase := range testCases {

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

func TestCustomNestedObjectType_renderValueMust(t *testing.T) {
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

func TestCustomNestedObjectType_renderValueNull(t *testing.T) {
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

func TestCustomNestedObjectType_renderValueType(t *testing.T) {
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

func TestCustomNestedObjectType_renderValueUnknown(t *testing.T) {
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

func TestCustomNestedObjectValue_renderAttributeTypes(t *testing.T) {
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, testCase.attrTypes, nil, nil)

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

func TestCustomNestedObjectValue_renderEqual(t *testing.T) {
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
		"attribute-name-same-as-generated-method-name": {
			name: "Example",
			attrValues: map[string]string{
				"type": "basetypes.ListValue",
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


if !v.ExampleType.Equal(other.ExampleType) {
return false
}


return true
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, testCase.attrValues, nil)

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

func TestCustomNestedObjectValue_renderIsNull(t *testing.T) {
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil, nil)

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

func TestCustomNestedObjectValue_renderIsUnknown(t *testing.T) {
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil, nil)

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

func TestCustomNestedObjectValue_renderString(t *testing.T) {
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil, nil)

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

func TestCustomNestedObjectValue_renderToObjectValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name            string
		attributeTypes  map[string]string
		attrTypes       map[string]string
		collectionTypes map[string]map[string]string
		expected        []byte
		expectedError   error
	}{
		"non-nested": {
			name: "Example",
			attributeTypes: map[string]string{
				"bool_attribute": "Bool",
			},
			attrTypes: map[string]string{
				"bool_attribute": "basetypes.BoolType{}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

attributeTypes := map[string]attr.Type{
"bool_attribute": basetypes.BoolType{},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"bool_attribute": v.BoolAttribute,
})

return objVal, diags
}`),
		},
		"list-nested": {
			name: "Example",
			attributeTypes: map[string]string{
				"list_nested_attribute": "ListNested",
			},
			attrTypes: map[string]string{
				"list_nested_attribute": "basetypes.ListType{}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

listNestedAttribute := types.ListValueMust(
ListNestedAttributeType{
basetypes.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
v.ListNestedAttribute.Elements(),
)

if v.ListNestedAttribute.IsNull() {
listNestedAttribute = types.ListNull(
ListNestedAttributeType{
basetypes.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
)
}

if v.ListNestedAttribute.IsUnknown() {
listNestedAttribute = types.ListUnknown(
ListNestedAttributeType{
basetypes.ObjectType{
AttrTypes: ListNestedAttributeValue{}.AttributeTypes(ctx),
},
},
)
}


attributeTypes := map[string]attr.Type{
"list_nested_attribute": basetypes.ListType{},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"list_nested_attribute": listNestedAttribute,
})

return objVal, diags
}`),
		},
		"map-nested": {
			name: "Example",
			attributeTypes: map[string]string{
				"map_nested_attribute": "MapNested",
			},
			attrTypes: map[string]string{
				"map_nested_attribute": "basetypes.MapType{}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

mapNestedAttribute := types.MapValueMust(
MapNestedAttributeType{
basetypes.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
v.MapNestedAttribute.Elements(),
)

if v.MapNestedAttribute.IsNull() {
mapNestedAttribute = types.MapNull(
MapNestedAttributeType{
basetypes.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
)
}

if v.MapNestedAttribute.IsUnknown() {
mapNestedAttribute = types.MapUnknown(
MapNestedAttributeType{
basetypes.ObjectType{
AttrTypes: MapNestedAttributeValue{}.AttributeTypes(ctx),
},
},
)
}


attributeTypes := map[string]attr.Type{
"map_nested_attribute": basetypes.MapType{},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"map_nested_attribute": mapNestedAttribute,
})

return objVal, diags
}`),
		},
		"set-nested": {
			name: "Example",
			attributeTypes: map[string]string{
				"set_nested_attribute": "SetNested",
			},
			attrTypes: map[string]string{
				"set_nested_attribute": "basetypes.SetType{}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

setNestedAttribute := types.SetValueMust(
SetNestedAttributeType{
basetypes.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
v.SetNestedAttribute.Elements(),
)

if v.SetNestedAttribute.IsNull() {
setNestedAttribute = types.SetNull(
SetNestedAttributeType{
basetypes.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
)
}

if v.SetNestedAttribute.IsUnknown() {
setNestedAttribute = types.SetUnknown(
SetNestedAttributeType{
basetypes.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
)
}


attributeTypes := map[string]attr.Type{
"set_nested_attribute": basetypes.SetType{},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"set_nested_attribute": setNestedAttribute,
})

return objVal, diags
}`),
		},
		"non-nested-attribute-name-same-as-generated-method-name": {
			name: "Example",
			attributeTypes: map[string]string{
				"type": "Bool",
			},
			attrTypes: map[string]string{
				"type": "basetypes.BoolType{}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

attributeTypes := map[string]attr.Type{
"type": basetypes.BoolType{},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"type": v.ExampleType,
})

return objVal, diags
}`),
		},
		"nested-attribute-name-same-as-generated-method-name": {
			name: "Example",
			attributeTypes: map[string]string{
				"type": "ListNested",
			},
			attrTypes: map[string]string{
				"type": "basetypes.ListType{}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

exampleType := types.ListValueMust(
TypeType{
basetypes.ObjectType{
AttrTypes: TypeValue{}.AttributeTypes(ctx),
},
},
v.ExampleType.Elements(),
)

if v.ExampleType.IsNull() {
exampleType = types.ListNull(
TypeType{
basetypes.ObjectType{
AttrTypes: TypeValue{}.AttributeTypes(ctx),
},
},
)
}

if v.ExampleType.IsUnknown() {
exampleType = types.ListUnknown(
TypeType{
basetypes.ObjectType{
AttrTypes: TypeValue{}.AttributeTypes(ctx),
},
},
)
}


attributeTypes := map[string]attr.Type{
"type": basetypes.ListType{},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"type": exampleType,
})

return objVal, diags
}`),
		},
		"collection-type-list": {
			name: "Example",
			attributeTypes: map[string]string{
				"list_attribute": "List",
			},
			attrTypes: map[string]string{
				"list_attribute": "basetypes.ListType{\nElemType: types.BoolType,\n}",
			},
			collectionTypes: map[string]map[string]string{
				"list_attribute": {
					"ElementType":   "types.BoolType",
					"TypeValueFunc": "types.ListValue",
				},
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

var listAttributeVal basetypes.ListValue
switch {
case v.ListAttribute.IsUnknown():
listAttributeVal = types.ListUnknown(types.BoolType)
case v.ListAttribute.IsNull():
listAttributeVal = types.ListNull(types.BoolType)
default:
var d diag.Diagnostics
listAttributeVal, d = types.ListValue(types.BoolType, v.ListAttribute.Elements())
diags.Append(d...)
}

if diags.HasError() {
return types.ObjectUnknown(map[string]attr.Type{
"list_attribute": basetypes.ListType{
ElemType: types.BoolType,
},
}), diags
}

attributeTypes := map[string]attr.Type{
"list_attribute": basetypes.ListType{
ElemType: types.BoolType,
},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"list_attribute": listAttributeVal,
})

return objVal, diags
}`),
		},
		"collection-type-map": {
			name: "Example",
			attributeTypes: map[string]string{
				"map_attribute": "Map",
			},
			attrTypes: map[string]string{
				"map_attribute": "basetypes.MapType{\nElemType: types.StringType,\n}",
			},
			collectionTypes: map[string]map[string]string{
				"map_attribute": {
					"ElementType":   "types.StringType",
					"TypeValueFunc": "types.MapValue",
				},
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

var mapAttributeVal basetypes.MapValue
switch {
case v.MapAttribute.IsUnknown():
mapAttributeVal = types.MapUnknown(types.StringType)
case v.MapAttribute.IsNull():
mapAttributeVal = types.MapNull(types.StringType)
default:
var d diag.Diagnostics
mapAttributeVal, d = types.MapValue(types.StringType, v.MapAttribute.Elements())
diags.Append(d...)
}

if diags.HasError() {
return types.ObjectUnknown(map[string]attr.Type{
"map_attribute": basetypes.MapType{
ElemType: types.StringType,
},
}), diags
}

attributeTypes := map[string]attr.Type{
"map_attribute": basetypes.MapType{
ElemType: types.StringType,
},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"map_attribute": mapAttributeVal,
})

return objVal, diags
}`),
		},
		"collection-type-set": {
			name: "Example",
			attributeTypes: map[string]string{
				"set_attribute": "Set",
			},
			attrTypes: map[string]string{
				"set_attribute": "basetypes.SetType{\nElemType: types.Int64Type,\n}",
			},
			collectionTypes: map[string]map[string]string{
				"set_attribute": {
					"ElementType":   "types.Int64Type",
					"TypeValueFunc": "types.SetValue",
				},
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

var setAttributeVal basetypes.SetValue
switch {
case v.SetAttribute.IsUnknown():
setAttributeVal = types.SetUnknown(types.Int64Type)
case v.SetAttribute.IsNull():
setAttributeVal = types.SetNull(types.Int64Type)
default:
var d diag.Diagnostics
setAttributeVal, d = types.SetValue(types.Int64Type, v.SetAttribute.Elements())
diags.Append(d...)
}

if diags.HasError() {
return types.ObjectUnknown(map[string]attr.Type{
"set_attribute": basetypes.SetType{
ElemType: types.Int64Type,
},
}), diags
}

attributeTypes := map[string]attr.Type{
"set_attribute": basetypes.SetType{
ElemType: types.Int64Type,
},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"set_attribute": setAttributeVal,
})

return objVal, diags
}`),
		},
		"collection-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			attributeTypes: map[string]string{
				"type": "List",
			},
			attrTypes: map[string]string{
				"type": "basetypes.ListType{\nElemType: types.BoolType,\n}",
			},
			collectionTypes: map[string]map[string]string{
				"type": {
					"ElementType":   "types.BoolType",
					"TypeValueFunc": "types.ListValue",
				},
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

var typeVal basetypes.ListValue
switch {
case v.ExampleType.IsUnknown():
typeVal = types.ListUnknown(types.BoolType)
case v.ExampleType.IsNull():
typeVal = types.ListNull(types.BoolType)
default:
var d diag.Diagnostics
typeVal, d = types.ListValue(types.BoolType, v.ExampleType.Elements())
diags.Append(d...)
}

if diags.HasError() {
return types.ObjectUnknown(map[string]attr.Type{
"type": basetypes.ListType{
ElemType: types.BoolType,
},
}), diags
}

attributeTypes := map[string]attr.Type{
"type": basetypes.ListType{
ElemType: types.BoolType,
},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"type": typeVal,
})

return objVal, diags
}`),
		},
		"object-type": {
			name: "Example",
			attributeTypes: map[string]string{
				"object_attribute": "Object",
			},
			attrTypes: map[string]string{
				"object_attribute": "basetypes.ObjectType{\nAttrTypes: map[string]attr.Type{\n\"bool\": types.BoolType,\n\"float64\": types.Float64Type,\n\"int64\": types.Int64Type,\n\"number\": types.NumberType,\n\"string\": types.StringType,\n},\n}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

objectAttributeVal, d := types.ObjectValue(v.ObjectAttribute.AttributeTypes(ctx), v.ObjectAttribute.Attributes())

diags.Append(d...)

if d.HasError() {
return types.ObjectUnknown(map[string]attr.Type{
"object_attribute": basetypes.ObjectType{
AttrTypes: v.ObjectAttribute.AttributeTypes(ctx),
},
}), diags
}

attributeTypes := map[string]attr.Type{
"object_attribute": basetypes.ObjectType{
AttrTypes: v.ObjectAttribute.AttributeTypes(ctx),
},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"object_attribute": objectAttributeVal,
})

return objVal, diags
}`),
		},
		"object-type-attribute-name-same-as-generated-method-name": {
			name: "Example",
			attributeTypes: map[string]string{
				"type": "Object",
			},
			attrTypes: map[string]string{
				"type": "basetypes.ObjectType{\nAttrTypes: map[string]attr.Type{\n\"bool\": types.BoolType,\n\"float64\": types.Float64Type,\n\"int64\": types.Int64Type,\n\"number\": types.NumberType,\n\"string\": types.StringType,\n},\n}",
			},
			expected: []byte(`
func (v ExampleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
var diags diag.Diagnostics

typeVal, d := types.ObjectValue(v.ExampleType.AttributeTypes(ctx), v.ExampleType.Attributes())

diags.Append(d...)

if d.HasError() {
return types.ObjectUnknown(map[string]attr.Type{
"type": basetypes.ObjectType{
AttrTypes: v.ExampleType.AttributeTypes(ctx),
},
}), diags
}

attributeTypes := map[string]attr.Type{
"type": basetypes.ObjectType{
AttrTypes: v.ExampleType.AttributeTypes(ctx),
},
}

if v.IsNull() {
return types.ObjectNull(attributeTypes), diags
}

if v.IsUnknown() {
return types.ObjectUnknown(attributeTypes), diags
}

objVal, diags := types.ObjectValue(
attributeTypes,
map[string]attr.Value{
"type": typeVal,
})

return objVal, diags
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, testCase.attributeTypes, testCase.attrTypes, nil, testCase.collectionTypes)

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

func TestCustomNestedObjectValue_renderToTerraformValue(t *testing.T) {
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
		"attribute-name-same-as-generated-method-name": {
			name: "Example",
			attrTypes: map[string]string{
				"type": "basetypes.BoolType{}",
			},
			expected: []byte(`func (v ExampleValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
attrTypes := make(map[string]tftypes.Type, 1)

var val tftypes.Value
var err error


attrTypes["type"] = basetypes.BoolType{}.TerraformType(ctx)

objectType := tftypes.Object{AttributeTypes: attrTypes}

switch v.state {
case attr.ValueStateKnown:
vals := make(map[string]tftypes.Value, 1)


val, err = v.ExampleType.ToTerraformValue(ctx)

if err != nil {
return tftypes.NewValue(objectType, tftypes.UnknownValue), err
}

vals["type"] = val



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
		"no-attributes": {
			name: "Example",
			expected: []byte(`func (v ExampleValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
attrTypes := make(map[string]tftypes.Type, 0)



objectType := tftypes.Object{AttributeTypes: attrTypes}

switch v.state {
case attr.ValueStateKnown:
vals := make(map[string]tftypes.Value, 0)



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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, testCase.attrTypes, nil, nil)

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

func TestCustomNestedObjectValue_renderType(t *testing.T) {
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil, nil)

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

func TestCustomNestedObjectValue_renderValuable(t *testing.T) {
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, nil, nil)

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

func TestCustomNestedObjectValue_renderValue(t *testing.T) {
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
			expected: []byte(`type ExampleValue struct {
BoolAttribute basetypes.BoolValue ` + "`" + `tfsdk:"bool_attribute"` + "`" + `
state attr.ValueState
}`),
		},
		"attribute-name-same-as-generated-method-name": {
			name: "Example",
			attrValues: map[string]string{
				"type": "basetypes.BoolValue",
			},
			expected: []byte(`type ExampleValue struct {
ExampleType basetypes.BoolValue ` + "`" + `tfsdk:"type"` + "`" + `
state attr.ValueState
}`),
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			customObjectValue := NewCustomNestedObjectValue(testCase.name, nil, nil, testCase.attrValues, nil)

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
