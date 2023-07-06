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
				Attributes: map[string]GeneratorAttribute{
					"bool": GeneratorBoolAttribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"float64": GeneratorFloat64Attribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"int64": GeneratorInt64Attribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"list_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Float64: &specschema.Float64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.Float64Type,
},
}
}`),
		},
		"list_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Int64: &specschema.Int64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.Int64Type,
},
}
}`),
		},
		"list_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							Number: &specschema.NumberType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.NumberType,
},
}
}`),
		},
		"list_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"list": GeneratorListAttribute{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.StringType,
},
}
}`),
		},
		"list_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"list_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"list_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute_outer": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"list_nested_attribute_inner": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute_outer": types.ListType{
ElemType: ListNestedAttributeOuterModel{}.objectType(),
},
}
}

func (m ListNestedAttributeOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute_inner": types.ListType{
ElemType: ListNestedAttributeInnerModel{}.objectType(),
},
}
}

func (m ListNestedAttributeInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"list_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_attribute": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block_outer": types.ListType{
ElemType: ListNestedBlockOuterModel{}.objectType(),
},
}
}

func (m ListNestedBlockOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block_inner": types.ListType{
ElemType: ListNestedBlockInnerModel{}.objectType(),
},
}
}

func (m ListNestedBlockInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockInnerModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"map_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Float64: &specschema.Float64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.Float64Type,
},
}
}`),
		},
		"map_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Int64: &specschema.Int64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.Int64Type,
},
}
}`),
		},
		"map_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							Number: &specschema.NumberType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.NumberType,
},
}
}`),
		},
		"map_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"map": GeneratorMapAttribute{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.StringType,
},
}
}`),
		},
		"map_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"map_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"map_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute_outer": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"map_nested_attribute_inner": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute_outer": types.MapType{
ElemType: MapNestedAttributeOuterModel{}.objectType(),
},
}
}

func (m MapNestedAttributeOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute_inner": types.MapType{
ElemType: MapNestedAttributeInnerModel{}.objectType(),
},
}
}

func (m MapNestedAttributeInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"map_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_attribute": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"number": GeneratorNumberAttribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"object_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Float64: &specschema.Float64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.Float64Type,
},
}
}`),
		},
		"set_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Int64: &specschema.Int64Type{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.Int64Type,
},
}
}`),
		},
		"set_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							Number: &specschema.NumberType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.NumberType,
},
}
}`),
		},
		"set_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"set": GeneratorSetAttribute{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.StringType,
},
}
}`),
		},
		"set_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"bool": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"set_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"set_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"set_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute_outer": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"set_nested_attribute_inner": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute_outer": types.SetType{
ElemType: SetNestedAttributeOuterModel{}.objectType(),
},
}
}

func (m SetNestedAttributeOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute_inner": types.SetType{
ElemType: SetNestedAttributeInnerModel{}.objectType(),
},
}
}

func (m SetNestedAttributeInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_attribute": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"float64": GeneratorFloat64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"int64": GeneratorInt64Attribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"list_nested_attribute": GeneratorListNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"map_nested_attribute": GeneratorMapNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"number": GeneratorNumberAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"set_nested_attribute": GeneratorSetNestedAttribute{
									NestedObject: GeneratorNestedAttributeObject{
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
										Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block_outer": types.SetType{
ElemType: SetNestedBlockOuterModel{}.objectType(),
},
}
}

func (m SetNestedBlockOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block_inner": types.SetType{
ElemType: SetNestedBlockInnerModel{}.objectType(),
},
}
}

func (m SetNestedBlockInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockInnerModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"single_nested_attribute": GeneratorSingleNestedAttribute{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
							Attributes: map[string]GeneratorAttribute{
								"string": GeneratorStringAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"single_nested_attribute_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"bool": GeneratorBoolAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"float64": GeneratorFloat64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"single_nested_attribute_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"int64": GeneratorInt64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"single_nested_attribute_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_attribute_list_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"list_nested_attribute": GeneratorListNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_attribute_map_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"map_nested_attribute": GeneratorMapNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"number": GeneratorNumberAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"single_nested_attribute_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_attribute_set_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"set_nested_attribute": GeneratorSetNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute_outer": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"single_nested_attribute_inner": GeneratorSingleNestedAttribute{
								Attributes: map[string]GeneratorAttribute{
									"bool": GeneratorBoolAttribute{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute_outer": types.ObjectType{
AttrTypes: SingleNestedAttributeOuterModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute_inner": types.ObjectType{
AttrTypes: SingleNestedAttributeInnerModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_attribute_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_attribute": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"string": GeneratorStringAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_float64": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
							"float64": GeneratorFloat64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"single_nested_block_int64": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
							"int64": GeneratorInt64Attribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"single_nested_block_list": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
						Attributes: map[string]GeneratorAttribute{
							"list_nested_attribute": GeneratorListNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_attribute": types.ListType{
ElemType: ListNestedAttributeModel{}.objectType(),
},
}
}

func (m ListNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_block": types.ListType{
ElemType: ListNestedBlockModel{}.objectType(),
},
}
}

func (m ListNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_map": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
						Attributes: map[string]GeneratorAttribute{
							"map_nested_attribute": GeneratorMapNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_attribute": types.MapType{
ElemType: MapNestedAttributeModel{}.objectType(),
},
}
}

func (m MapNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_number": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
							"number": GeneratorNumberAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"single_nested_block_object": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
						Attributes: map[string]GeneratorAttribute{
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
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
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
						Attributes: map[string]GeneratorAttribute{
							"set_nested_attribute": GeneratorSetNestedAttribute{
								NestedObject: GeneratorNestedAttributeObject{
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_attribute": types.SetType{
ElemType: SetNestedAttributeModel{}.objectType(),
},
}
}

func (m SetNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
									Attributes: map[string]GeneratorAttribute{
										"bool": GeneratorBoolAttribute{},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_block": types.SetType{
ElemType: SetNestedBlockModel{}.objectType(),
},
}
}

func (m SetNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_single_nested_attribute": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
							"single_nested_attribute": GeneratorSingleNestedAttribute{
								Attributes: map[string]GeneratorAttribute{
									"bool": GeneratorBoolAttribute{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_attribute": types.ObjectType{
AttrTypes: SingleNestedAttributeModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedAttributeModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedAttributeModel) objectAttributeTypes() map[string]attr.Type {
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
								Attributes: map[string]GeneratorAttribute{
									"bool": GeneratorBoolAttribute{},
								},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block_outer": types.ObjectType{
AttrTypes: SingleNestedBlockOuterModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block_inner": types.ObjectType{
AttrTypes: SingleNestedBlockInnerModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_block_string": {
			input: GeneratorDataSourceSchema{
				Blocks: GeneratorBlocks{
					"single_nested_block": GeneratorSingleNestedBlock{
						Attributes: map[string]GeneratorAttribute{
							"string": GeneratorStringAttribute{},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_block": types.ObjectType{
AttrTypes: SingleNestedBlockModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedBlockModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedBlockModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"string": GeneratorStringAttribute{},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
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
