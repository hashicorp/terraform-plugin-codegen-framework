// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var datasourceDataSourceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"list_attribute": schema.ListAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
		"list_list_attribute": schema.ListAttribute{
			ElementType: types.ListType{
				ElemType: types.StringType,
			},
			Computed: true,
		},
		"list_map_attribute": schema.ListAttribute{
			ElementType: types.MapType{
				ElemType: types.StringType,
			},
			Computed: true,
		},
		"list_nested_bool_attribute": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"bool_attribute": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			Computed: true,
		},
		"list_nested_list_nested_bool_attribute": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"list_nested_attribute": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bool_attribute": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						Computed: true,
					},
				},
			},
			Computed: true,
		},
		"list_nested_list_nested_list_attribute": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"list_nested_attribute": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"list_attribute": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
				},
			},
			Computed: true,
		},
		"list_object_attribute": schema.ListAttribute{
			ElementType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"obj_string_attr": types.StringType,
				},
			},
			Computed: true,
		},
		"list_object_object_attribute": schema.ListAttribute{
			ElementType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"obj_obj_attr": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"obj_obj_string_attr": types.StringType,
						},
					},
				},
			},
			Computed: true,
		},
		"object_attribute": schema.ObjectAttribute{
			AttributeTypes: map[string]attr.Type{
				"obj_string_attr": types.StringType,
			},
			Computed: true,
		},
		"object_list_attribute": schema.ObjectAttribute{
			AttributeTypes: map[string]attr.Type{
				"obj_list_attr": types.ListType{
					ElemType: types.StringType,
				},
			},
			Computed: true,
		},
		"object_list_object_attribute": schema.ObjectAttribute{
			AttributeTypes: map[string]attr.Type{
				"obj_list_attr": types.ListType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"obj_list_obj_attr": types.StringType,
						},
					},
				},
			},
			Computed: true,
		},
		"single_nested_bool_attribute": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"bool_attribute": schema.BoolAttribute{
					Computed: true,
				},
			},
			Computed: true,
		},
		"single_nested_single_nested_bool_attribute": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"single_nested_attribute": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"bool_attribute": schema.BoolAttribute{
							Computed: true,
						},
					},
					Computed: true,
				},
			},
			Computed: true,
		},
		"single_nested_single_nested_list_attribute": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"single_nested_attribute": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"list_attribute": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
					},
					Computed: true,
				},
			},
			Computed: true,
		},
	},
	Blocks: map[string]schema.Block{
		"list_nested_block_bool_attribute": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"bool_attribute": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
		"list_nested_block_object_attribute_list_nested_nested_block_list_attribute": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"object_attribute": schema.ObjectAttribute{
						AttributeTypes: map[string]attr.Type{
							"obj_string_attr": types.StringType,
						},
						Computed: true,
					},
				},
				Blocks: map[string]schema.Block{
					"list_nested_block": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"list_attribute": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
		"list_nested_list_nested_block_bool_attribute": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Blocks: map[string]schema.Block{
					"list_nested_block": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"bool_attribute": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
		"single_nested_block_bool_attribute": schema.SingleNestedBlock{
			Attributes: map[string]schema.Attribute{
				"bool_attribute": schema.BoolAttribute{
					Computed: true,
				},
			},
		},
		"single_nested_block_object_attribute_single_nested_list_nested_block_list_attribute": schema.SingleNestedBlock{
			Attributes: map[string]schema.Attribute{
				"object_attribute": schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"obj_string_attr": types.StringType,
					},
					Computed: true,
				},
			},
			Blocks: map[string]schema.Block{
				"list_nested_block": schema.ListNestedBlock{
					NestedObject: schema.NestedBlockObject{
						Attributes: map[string]schema.Attribute{
							"list_attribute": schema.ListAttribute{
								ElementType: types.StringType,
								Computed:    true,
							},
						},
					},
				},
			},
		},
		"single_nested_single_nested_block_bool_attribute": schema.SingleNestedBlock{
			Blocks: map[string]schema.Block{
				"single_nested_block": schema.SingleNestedBlock{
					Attributes: map[string]schema.Attribute{
						"bool_attribute": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	},
}
