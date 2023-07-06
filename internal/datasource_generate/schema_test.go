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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute_outer": types.ListType{
ElemType: ListNestedAttributeOuterModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeOuterModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeOuterModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute_inner": types.ListType{
ElemType: ListNestedAttributeInnerModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeInnerModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeInnerModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block_outer": types.ListType{
ElemType: ListNestedBlockOuterModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockOuterModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockOuterModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block_inner": types.ListType{
ElemType: ListNestedBlockInnerModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockInnerModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockInnerModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute_outer": types.MapType{
ElemType: MapNestedAttributeOuterModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeOuterModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeOuterModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute_inner": types.MapType{
ElemType: MapNestedAttributeInnerModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeInnerModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeInnerModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute_outer": types.SetType{
ElemType: SetNestedAttributeOuterModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeOuterModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeOuterModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute_inner": types.SetType{
ElemType: SetNestedAttributeInnerModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeInnerModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeInnerModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block_outer": types.SetType{
ElemType: SetNestedBlockOuterModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockOuterModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockOuterModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block_inner": types.SetType{
ElemType: SetNestedBlockInnerModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockInnerModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockInnerModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute_outer": types.ObjectType{
AttrTypes: SingleNestedAttributeOuterModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeOuterModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeOuterModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute_inner": types.ObjectType{
AttrTypes: SingleNestedAttributeInnerModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeInnerModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeInnerModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.ObjectType(),
},
}
}

func (m ListNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.ObjectType(),
},
}
}

func (m ListNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ListNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.ObjectType(),
},
}
}

func (m MapNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m MapNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.ObjectType(),
},
}
}

func (m SetNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.ObjectType(),
},
}
}

func (m SetNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SetNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedAttributeModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block_outer": types.ObjectType{
AttrTypes: SingleNestedBlockOuterModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockOuterModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockOuterModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block_inner": types.ObjectType{
AttrTypes: SingleNestedBlockInnerModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockInnerModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockInnerModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.ObjectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m SingleNestedBlockModel) ObjectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) ObjectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.ObjectAttributeTypes()}
}

func (m ExampleModel) ObjectAttributeTypes() map[string]attr.Type {
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
