// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestGeneratorProviderSchema_ModelObjectHelpersTemplate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorProviderSchema
		expected      []byte
		expectedError error
	}{
		"bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"bool": GeneratorBoolAttribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"float64": GeneratorFloat64Attribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"int64": GeneratorInt64Attribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Float64: &specschema.Float64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.Float64Type,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Int64: &specschema.Int64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.Int64Type,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Map: &specschema.MapType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Number: &specschema.NumberType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.NumberType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_object": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Set: &specschema.SetType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.SetType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.StringType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"list": GeneratorListAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_list_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute_outer": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"list_nested_attribute_inner": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute_outer": types.ListType{
ElemType: ListNestedAttributeOuterModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeOuterModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeOuterModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute_inner": types.ListType{
ElemType: ListNestedAttributeInnerModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedAttributeOuterModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeOuterModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeInnerModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeInnerModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"map": GeneratorMapAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_map_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_object": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"object": GeneratorObjectAttribute{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name: "bool",
											Bool: &specschema.BoolType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"set": GeneratorSetAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_set_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_single_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_attribute_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_bool": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_float64": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_int64": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_list": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"list": GeneratorListAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_list_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_list_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block_outer": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Blocks: GeneratorBlocks{
								"list_nested_block_inner": GeneratorListNestedBlock{
									NestedObject: GeneratorNestedBlockObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block_outer": types.ListType{
ElemType: ListNestedBlockOuterModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockOuterModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockOuterModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block_inner": types.ListType{
ElemType: ListNestedBlockInnerModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedBlockOuterModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockOuterModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedBlockInnerModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockInnerModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_map": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"map": GeneratorMapAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_map_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_number": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_object": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"object": GeneratorObjectAttribute{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name: "bool",
											Bool: &specschema.BoolType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_set": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"set": GeneratorSetAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_set_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_set_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Blocks: GeneratorBlocks{
								"set_nested_block": GeneratorSetNestedBlock{
									NestedObject: GeneratorNestedBlockObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_single_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_single_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Blocks: GeneratorBlocks{
								"single_nested_block": GeneratorSingleNestedBlock{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"list_nested_block_string": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"list_nested_block": GeneratorListNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Float64: &specschema.Float64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.Float64Type,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Int64: &specschema.Int64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.Int64Type,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Map: &specschema.MapType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Number: &specschema.NumberType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.NumberType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_object": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Set: &specschema.SetType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.SetType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.StringType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"list": GeneratorListAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_list_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"map": GeneratorMapAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_map_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute_outer": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"map_nested_attribute_inner": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute_outer": types.MapType{
ElemType: MapNestedAttributeOuterModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeOuterModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeOuterModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute_inner": types.MapType{
ElemType: MapNestedAttributeInnerModel{}.ObjectType(ctx),
},
}
}

func (m MapNestedAttributeOuterModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeOuterModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeInnerModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeInnerModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_object": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"object": GeneratorObjectAttribute{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name: "bool",
											Bool: &specschema.BoolType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"set": GeneratorSetAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_set_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_single_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"map_nested_attribute_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"number": GeneratorNumberAttribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:    "float64",
								Float64: &specschema.Float64Type{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"float64": types.Float64Type,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:  "int64",
								Int64: &specschema.Int64Type{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"int64": types.Int64Type,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "map",
								Map: &specschema.MapType{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "number",
								Number: &specschema.NumberType{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"number": types.NumberType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "set",
								Set: &specschema.SetType{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"object_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"object": GeneratorObjectAttribute{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "string",
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"string": types.StringType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Float64: &specschema.Float64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.Float64Type,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Int64: &specschema.Int64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.Int64Type,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Map: &specschema.MapType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Number: &specschema.NumberType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.NumberType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_object": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Set: &specschema.SetType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.SetType{
ElemType: types.BoolType,
},
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.StringType,
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"list": GeneratorListAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_list_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"map": GeneratorMapAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_map_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_object": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"object": GeneratorObjectAttribute{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name: "bool",
											Bool: &specschema.BoolType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"set": GeneratorSetAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_set_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute_outer": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"set_nested_attribute_inner": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute_outer": types.SetType{
ElemType: SetNestedAttributeOuterModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeOuterModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeOuterModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute_inner": types.SetType{
ElemType: SetNestedAttributeInnerModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedAttributeOuterModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeOuterModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeInnerModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeInnerModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_single_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_attribute_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: GeneratorAttributes{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_bool": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_float64": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_int64": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_list": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"list": GeneratorListAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_list_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_list_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Blocks: GeneratorBlocks{
								"list_nested_block": GeneratorListNestedBlock{
									NestedObject: GeneratorNestedBlockObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_map": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"map": GeneratorMapAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_map_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_number": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_object": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"object": GeneratorObjectAttribute{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name: "bool",
											Bool: &specschema.BoolType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_set": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"set": GeneratorSetAttribute{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_set_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_set_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block_outer": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Blocks: GeneratorBlocks{
								"set_nested_block_inner": GeneratorSetNestedBlock{
									NestedObject: GeneratorNestedBlockObject{
										Attributes: GeneratorAttributes{
											"bool": GeneratorBoolAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block_outer": types.SetType{
ElemType: SetNestedBlockOuterModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockOuterModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockOuterModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block_inner": types.SetType{
ElemType: SetNestedBlockInnerModel{}.ObjectType(ctx),
},
}
}

func (m SetNestedBlockOuterModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockOuterModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedBlockInnerModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockInnerModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_single_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_single_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Blocks: GeneratorBlocks{
								"single_nested_block": GeneratorSingleNestedBlock{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"set_nested_block_string": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"set_nested_block": GeneratorSetNestedBlock{
						NestedObject: GeneratorNestedBlockObject{
							Attributes: GeneratorAttributes{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_bool": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"bool": GeneratorBoolAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_float64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"float64": GeneratorFloat64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_int64": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"int64": GeneratorInt64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_list": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"list": GeneratorListAttribute{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_list_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"list_nested_attribute": GeneratorListNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_map": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"map": GeneratorMapAttribute{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_map_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"map_nested_attribute": GeneratorMapNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_number": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"number": GeneratorNumberAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_object": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"object": GeneratorObjectAttribute{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_set": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"set": GeneratorSetAttribute{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_set_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"set_nested_attribute": GeneratorSetNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_single_nested_attribute": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute_outer": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"single_nested_attribute_inner": GeneratorSingleNestedAttribute{
								Attributes: GeneratorAttributes{
									"bool": GeneratorBoolAttribute{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute_outer": types.ObjectType{
AttrTypes: SingleNestedAttributeOuterModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeOuterModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeOuterModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute_inner": types.ObjectType{
AttrTypes: SingleNestedAttributeInnerModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m SingleNestedAttributeOuterModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeOuterModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeInnerModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeInnerModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_attribute_string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: GeneratorAttributes{
							"string": GeneratorStringAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_bool": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"bool": GeneratorBoolAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_float64": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"float64": GeneratorFloat64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_int64": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"int64": GeneratorInt64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_list": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"list": GeneratorListAttribute{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_list_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"list_nested_attribute": GeneratorListNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_list_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Blocks: GeneratorBlocks{
							"list_nested_block": GeneratorListNestedBlock{
								NestedObject: GeneratorNestedBlockObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m ListNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ListNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_map": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"map": GeneratorMapAttribute{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_map_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"map_nested_attribute": GeneratorMapNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m MapNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m MapNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_number": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"number": GeneratorNumberAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_object": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"object": GeneratorObjectAttribute{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_set": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"set": GeneratorSetAttribute{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_set_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"set_nested_attribute": GeneratorSetNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_set_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Blocks: GeneratorBlocks{
							"set_nested_block": GeneratorSetNestedBlock{
								NestedObject: GeneratorNestedBlockObject{
									Attributes: GeneratorAttributes{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(ctx),
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SetNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SetNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_single_nested_attribute": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"single_nested_attribute": GeneratorSingleNestedAttribute{
								Attributes: GeneratorAttributes{
									"bool": GeneratorBoolAttribute{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedAttributeModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedAttributeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_single_nested_block": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block_outer": GeneratorSingleNestedBlock{
						Blocks: GeneratorBlocks{
							"single_nested_block_inner": GeneratorSingleNestedBlock{
								Attributes: GeneratorAttributes{
									"bool": GeneratorBoolAttribute{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block_outer": types.ObjectType{
AttrTypes: SingleNestedBlockOuterModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockOuterModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockOuterModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block_inner": types.ObjectType{
AttrTypes: SingleNestedBlockInnerModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m SingleNestedBlockOuterModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockOuterModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}

func (m SingleNestedBlockInnerModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockInnerModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"single_nested_block_string": {
			input: GeneratorProviderSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: GeneratorAttributes{
							"string": GeneratorStringAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(ctx),
},
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m SingleNestedBlockModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m SingleNestedBlockModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
		"string": {
			input: GeneratorProviderSchema{
				Attributes: GeneratorAttributes{
					"string": GeneratorStringAttribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ExampleModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}

func (m ExampleModel) ObjectNull(ctx context.Context) types.Object {
return types.ObjectNull(
m.ObjectAttributeTypes(ctx),
)
}

func (m ExampleModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
return types.ObjectValueFrom(
ctx,
m.ObjectAttributeTypes(ctx),
data,
)
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelObjectHelpersTemplate("example")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
