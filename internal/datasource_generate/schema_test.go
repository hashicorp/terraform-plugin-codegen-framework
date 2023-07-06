// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestGeneratorDataSourceSchema_ModelObjectHelpersTemplate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorDataSourceSchema
		expected      []byte
		expectedError error
	}{
		"bool": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"float64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"int64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_bool": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_float64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_int64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_list": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_map": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_number": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_set": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_string": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"list_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"list_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"list_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_nested_attribute_set": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"list_nested_block_bool": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_float64": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"list_nested_block_int64": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"list_nested_block_list": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_block_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_list_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_map": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_block_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_number": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"list_nested_block_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"list_nested_block_set": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_block_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_set_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_single_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_block_string": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"map_bool": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_float64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_int64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_list": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_map": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_number": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_set": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_string": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"map_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"map_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"map_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"map_nested_attribute_set": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"number": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_bool": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_float64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_int64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_list": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_map": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_number": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_set": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"object_string": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_bool": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_float64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_int64": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_list": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_map": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_number": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_set": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_string": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"set_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"set_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"set_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_nested_attribute_set": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"set_nested_block_bool": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_float64": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"set_nested_block_int64": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"set_nested_block_list": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_block_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_list_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_map": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_block_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_number": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"set_nested_block_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"set_nested_block_set": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_block_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_set_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_single_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_block_string": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"single_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"single_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"single_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"single_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"single_nested_attribute_set": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"single_nested_block_bool": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_float64": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"single_nested_block_int64": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"single_nested_block_list": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_block_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_list_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m ListNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_map": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_block_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m MapNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_number": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"single_nested_block_object": {
			input: GeneratorDataSourceSchema{
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
}`),
		},
		"single_nested_block_set": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_block_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_set_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m SetNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SetNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedAttributeModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_single_nested_block": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockInnerModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockInnerModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_string": {
			input: GeneratorDataSourceSchema{
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

func (m SingleNestedBlockModel) ObjectType(ctx context.Context) types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"string": {
			input: GeneratorDataSourceSchema{
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
