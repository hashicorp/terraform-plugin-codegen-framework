// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorSchema_ModelObjectHelpersTemplate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         schema.GeneratorSchema
		expected      []byte
		expectedError error
	}{
		"list_nested_attribute": {
			input: schema.GeneratorSchema{
				Attributes: schema.GeneratorAttributes{
					"bool": GeneratorBoolAttribute{},
				},
			},
			expected: []byte(`
var _ basetypes.ObjectTypable = ListNestedAttributeType{}

type ListNestedAttributeType struct {
basetypes.ObjectType
}

func (t ListNestedAttributeType) Equal(o attr.Type) bool {
other, ok := o.(ListNestedAttributeType)

if !ok {
return false
}

return t.ObjectType.Equal(other.ObjectType)
}

func (t ListNestedAttributeType) String() string {
return "ListNestedAttributeType"
}

func (t ListNestedAttributeType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
var diags diag.Diagnostics

state := attr.ValueStateKnown

attributes := in.Attributes()


bool, ok := attributes["bool"]

if !ok {
diags.AddError(
"Attribute Missing",
` + "`" + `bool is missing from object` + "`" + `)

return nil, diags
}

boolVal, ok := bool.(basetypes.BoolValue)

if !ok {
diags.AddError(
"Attribute Wrong Type",
fmt.Sprintf(` + "`" + `bool expected to be basetypes.BoolValue, was: %T` + "`" + `, bool))
}

if boolVal.IsUnknown() {
state = attr.ValueStateUnknown
}


return ListNestedAttributeValue{
Bool: boolVal,
state: state,
}, diags
}

func NewListNestedAttributeValueNull() ListNestedAttributeValue {
return ListNestedAttributeValue{
state: attr.ValueStateNull,
}
}

func NewListNestedAttributeValueUnknown() ListNestedAttributeValue {
return ListNestedAttributeValue{
state: attr.ValueStateUnknown,
}
}

func NewListNestedAttributeValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ListNestedAttributeValue, diag.Diagnostics) {
var diags diag.Diagnostics

// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
ctx := context.Background()

for name, attributeType := range attributeTypes {
attribute, ok := attributes[name]

if !ok {
diags.AddError(
"Missing ListNestedAttributeValue Attribute Value",
"While creating a ListNestedAttributeValue value, a missing attribute value was detected. "+
"A ListNestedAttributeValue must contain values for all attributes, even if null or unknown. "+
"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
fmt.Sprintf("ListNestedAttributeValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
)

continue
}

if !attributeType.Equal(attribute.Type(ctx)) {
diags.AddError(
"Invalid ListNestedAttributeValue Attribute Type",
"While creating a ListNestedAttributeValue value, an invalid attribute value was detected. "+
"A ListNestedAttributeValue must use a matching attribute type for the value. "+
"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
fmt.Sprintf("ListNestedAttributeValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
fmt.Sprintf("ListNestedAttributeValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
)
}
}

for name := range attributes {
_, ok := attributeTypes[name]

if !ok {
diags.AddError(
"Extra ListNestedAttributeValue Attribute Value",
"While creating a ListNestedAttributeValue value, an extra attribute value was detected. "+
"A ListNestedAttributeValue must not contain values beyond the expected attribute types. "+
"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
fmt.Sprintf("Extra ListNestedAttributeValue Attribute Name: %s", name),
)
}
}

if diags.HasError() {
return NewListNestedAttributeValueUnknown(), diags
}

state := attr.ValueStateKnown


bool, ok := attributes["bool"]

if !ok {
diags.AddError(
"Attribute Missing",
` + "`" + `bool is missing from object` + "`" + `)

return NewListNestedAttributeValueNull(), diags
}

boolVal, ok := bool.(basetypes.BoolValue)

if !ok {
diags.AddError(
"Attribute Wrong Type",
fmt.Sprintf(` + "`" + `bool expected to be basetypes.BoolValue, was: %T` + "`" + `, bool))
}

if boolVal.IsUnknown() {
state = attr.ValueStateUnknown
}


return ListNestedAttributeValue{
Bool: boolVal,
state: state,
}, diags
}

func NewListNestedAttributeValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ListNestedAttributeValue {
object, diags := NewListNestedAttributeValue(attributeTypes, attributes)

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

panic("NewListNestedAttributeValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
}

return object
}

func (t ListNestedAttributeType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
if in.Type() == nil {
return NewListNestedAttributeValueNull(), nil
}

if !in.Type().Equal(t.TerraformType(ctx)) {
return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
}

if !in.IsKnown() {
return NewListNestedAttributeValueUnknown(), nil
}

if in.IsNull() {
return NewListNestedAttributeValueNull(), nil
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

return NewListNestedAttributeValueMust(t.AttrTypes, attributes), nil
}

func (t ListNestedAttributeType) ValueType(ctx context.Context) attr.Value {
return ListNestedAttributeValue{}
}

var _ basetypes.ObjectValuable = ListNestedAttributeValue{}

type ListNestedAttributeValue struct {
Bool basetypes.BoolValue ` + "`" + `tfsdk:"bool"` + "`" + `
state attr.ValueState
}

func (v ListNestedAttributeValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
attrTypes := make(map[string]tftypes.Type, 1)

var val tftypes.Value
var err error


attrTypes["bool"] = basetypes.BoolType{}.TerraformType(ctx)

objectType := tftypes.Object{AttributeTypes: attrTypes}

switch v.state {
case attr.ValueStateKnown:
vals := make(map[string]tftypes.Value, 1)


val, err = v.Bool.ToTerraformValue(ctx)

if err != nil {
return tftypes.NewValue(objectType, tftypes.UnknownValue), err
}

vals["bool"] = val



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
}

func (v ListNestedAttributeValue) IsNull() bool {
return v.state == attr.ValueStateNull
}

func (v ListNestedAttributeValue) IsUnknown() bool {
return v.state == attr.ValueStateUnknown
}

func (v ListNestedAttributeValue) String() string {
return "ListNestedAttributeValue"
}

func (v ListNestedAttributeValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
objVal, diags := types.ObjectValue(
map[string]attr.Type{
"bool": basetypes.BoolType{},
},
map[string]attr.Value{
"bool": v.Bool,
})

return objVal, diags
}

func (v ListNestedAttributeValue) Equal(o attr.Value) bool {
other, ok := o.(ListNestedAttributeValue)

if !ok {
return false
}

if v.state != other.state {
return false
}

if v.state != attr.ValueStateKnown {
return true
}


if !v.Bool.Equal(other.Bool) {
return false
}


return true
}

func (v ListNestedAttributeValue) Type(ctx context.Context) attr.Type {
return ListNestedAttributeType{
basetypes.ObjectType{
AttrTypes: v.AttributeTypes(ctx),
},
}
}

func (v ListNestedAttributeValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": basetypes.BoolType{},
}
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelObjectHelpersTemplate(name)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
