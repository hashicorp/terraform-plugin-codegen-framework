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
		"list_nested_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"list_nested_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"list_nested_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_list_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested_outer": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"list_nested_inner": GeneratorListNestedAttribute{
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
"list_nested_outer": types.ListType{
ElemType: ListNestedOuterModel{}.objectType(),
},
}
}

func (m ListNestedOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested_inner": types.ListType{
ElemType: ListNestedInnerModel{}.objectType(),
},
}
}

func (m ListNestedInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_map_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"map_nested": GeneratorMapNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"list_nested_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}`),
		},
		"list_nested_set": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"list_nested_set_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"set_nested": GeneratorSetNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_single_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"single_nested": GeneratorSingleNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"list_nested": GeneratorListNestedAttribute{
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
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
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
		"map_nested_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"map_nested_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"map_nested_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_list_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"list_nested": GeneratorListNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_map_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested_outer": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"map_nested_inner": GeneratorMapNestedAttribute{
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
"map_nested_outer": types.MapType{
ElemType: MapNestedOuterModel{}.objectType(),
},
}
}

func (m MapNestedOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested_inner": types.MapType{
ElemType: MapNestedInnerModel{}.objectType(),
},
}
}

func (m MapNestedInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"map_nested_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}`),
		},
		"map_nested_set": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"map_nested_set_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"set_nested": GeneratorSetNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_single_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"single_nested": GeneratorSingleNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"map_nested_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"map_nested": GeneratorMapNestedAttribute{
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
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
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
		"set_nested_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"set_nested_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"set_nested_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_list_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"list_nested": GeneratorListNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_map_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"map_nested": GeneratorMapNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"set_nested_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}`),
		},
		"set_nested_set": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"set_nested_set_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested_outer": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"set_nested_inner": GeneratorSetNestedAttribute{
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
"set_nested_outer": types.SetType{
ElemType: SetNestedOuterModel{}.objectType(),
},
}
}

func (m SetNestedOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested_inner": types.SetType{
ElemType: SetNestedInnerModel{}.objectType(),
},
}
}

func (m SetNestedInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_single_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"single_nested": GeneratorSingleNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"set_nested_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"set_nested": GeneratorSetNestedAttribute{
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
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"string": types.StringType,
}
}`),
		},
		"single_nested_bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_float64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"float64": types.Float64Type,
}
}`),
		},
		"single_nested_int64": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"int64": types.Int64Type,
}
}`),
		},
		"single_nested_list": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_list_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"list_nested": GeneratorListNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_map": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map": types.MapType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_map_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"map_nested": GeneratorMapNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"map_nested": types.MapType{
ElemType: MapNestedModel{}.objectType(),
},
}
}

func (m MapNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m MapNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_number": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"number": types.NumberType,
}
}`),
		},
		"single_nested_object": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"object": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
}
}`),
		},
		"single_nested_set": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set": types.SetType{
ElemType: types.BoolType,
},
}
}`),
		},
		"single_nested_set_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"set_nested": GeneratorSetNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"set_nested": types.SetType{
ElemType: SetNestedModel{}.objectType(),
},
}
}

func (m SetNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SetNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_single_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested_outer": GeneratorSingleNestedAttribute{
						Attributes: map[string]GeneratorAttribute{
							"single_nested_inner": GeneratorSingleNestedAttribute{
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
"single_nested_outer": types.ObjectType{
AttrTypes: SingleNestedOuterModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedOuterModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedOuterModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"single_nested_inner": types.ObjectType{
AttrTypes: SingleNestedInnerModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedInnerModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedInnerModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"single_nested_string": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"single_nested": GeneratorSingleNestedAttribute{
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
"single_nested": types.ObjectType{
AttrTypes: SingleNestedModel{}.objectAttributeTypes(),
},
}
}

func (m SingleNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m SingleNestedModel) objectAttributeTypes() map[string]attr.Type {
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
