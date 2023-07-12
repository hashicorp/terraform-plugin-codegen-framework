// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package generated

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var datasourceDataSourceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"bool_attribute": schema.BoolAttribute{
			Computed: true,
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
		"list_nested_attribute_one": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"bool_attribute": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			Computed: true,
		},
		"list_nested_attribute_three": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"list_nested_attribute_three_list_nested_attribute_one": schema.ListNestedAttribute{
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
		"list_nested_attribute_two": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"list_nested_attribute_two_list_nested_attribute_one": schema.ListNestedAttribute{
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
		"single_nested_attribute_one": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"bool_attribute": schema.BoolAttribute{
					Computed: true,
				},
			},
			Computed: true,
		},
		"single_nested_attribute_three": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"single_nested_attribute_three_single_nested_attribute_one": schema.SingleNestedAttribute{
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
		"single_nested_attribute_two": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"single_nested_attribute_two_single_nested_attribute_one": schema.SingleNestedAttribute{
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
	},
	Blocks: map[string]schema.Block{
		"list_nested_block_one": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"bool_attribute": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
		"list_nested_block_three": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"object_attribute": schema.ObjectAttribute{
						AttributeTypes: map[string]attr.Type{
							"string_attribute_type": types.StringType,
						},
						Computed: true,
					},
				},
				Blocks: map[string]schema.Block{
					"list_nested_block_three_list_nested_block_one": schema.ListNestedBlock{
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
		"list_nested_block_two": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Blocks: map[string]schema.Block{
					"list_nested_block_two_list_nested_block_one": schema.ListNestedBlock{
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
		"single_nested_block_one": schema.SingleNestedBlock{
			Attributes: map[string]schema.Attribute{
				"bool_attribute": schema.BoolAttribute{
					Computed: true,
				},
			},
		},
		"single_nested_block_three": schema.SingleNestedBlock{
			Attributes: map[string]schema.Attribute{
				"object_attribute": schema.ObjectAttribute{
					AttributeTypes: map[string]attr.Type{
						"string_attribute_type": types.StringType,
					},
					Computed: true,
				},
			},
			Blocks: map[string]schema.Block{
				"single_nested_block_three_list_nested_block_one": schema.ListNestedBlock{
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
		"single_nested_block_two": schema.SingleNestedBlock{
			Blocks: map[string]schema.Block{
				"single_nested_block_two_single_nested_block_one": schema.SingleNestedBlock{
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

type DatasourceModel struct {
	BoolAttribute              types.Bool   `tfsdk:"bool_attribute"`
	ListListAttribute          types.List   `tfsdk:"list_list_attribute"`
	ListMapAttribute           types.List   `tfsdk:"list_map_attribute"`
	ListNestedAttributeOne     types.List   `tfsdk:"list_nested_attribute_one"`
	ListNestedAttributeThree   types.List   `tfsdk:"list_nested_attribute_three"`
	ListNestedAttributeTwo     types.List   `tfsdk:"list_nested_attribute_two"`
	ListObjectAttribute        types.List   `tfsdk:"list_object_attribute"`
	ListObjectObjectAttribute  types.List   `tfsdk:"list_object_object_attribute"`
	ObjectAttribute            types.Object `tfsdk:"object_attribute"`
	ObjectListAttribute        types.Object `tfsdk:"object_list_attribute"`
	ObjectListObjectAttribute  types.Object `tfsdk:"object_list_object_attribute"`
	SingleNestedAttributeOne   types.Object `tfsdk:"single_nested_attribute_one"`
	SingleNestedAttributeThree types.Object `tfsdk:"single_nested_attribute_three"`
	SingleNestedAttributeTwo   types.Object `tfsdk:"single_nested_attribute_two"`
	ListNestedBlockOne         types.List   `tfsdk:"list_nested_block_one"`
	ListNestedBlockThree       types.List   `tfsdk:"list_nested_block_three"`
	ListNestedBlockTwo         types.List   `tfsdk:"list_nested_block_two"`
	SingleNestedBlockOne       types.Object `tfsdk:"single_nested_block_one"`
	SingleNestedBlockThree     types.Object `tfsdk:"single_nested_block_three"`
	SingleNestedBlockTwo       types.Object `tfsdk:"single_nested_block_two"`
}

type ListNestedAttributeOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

type ListNestedAttributeThreeModel struct {
	ListNestedAttributeThreeListNestedAttributeOne types.List `tfsdk:"list_nested_attribute_three_list_nested_attribute_one"`
}

type ListNestedAttributeThreeListNestedAttributeOneModel struct {
	ListAttribute types.List `tfsdk:"list_attribute"`
}

type ListNestedAttributeTwoModel struct {
	ListNestedAttributeTwoListNestedAttributeOne types.List `tfsdk:"list_nested_attribute_two_list_nested_attribute_one"`
}

type ListNestedAttributeTwoListNestedAttributeOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

type SingleNestedAttributeOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

type SingleNestedAttributeThreeModel struct {
	SingleNestedAttributeThreeSingleNestedAttributeOne types.Object `tfsdk:"single_nested_attribute_three_single_nested_attribute_one"`
}

type SingleNestedAttributeThreeSingleNestedAttributeOneModel struct {
	ListAttribute types.List `tfsdk:"list_attribute"`
}

type SingleNestedAttributeTwoModel struct {
	SingleNestedAttributeTwoSingleNestedAttributeOne types.Object `tfsdk:"single_nested_attribute_two_single_nested_attribute_one"`
}

type SingleNestedAttributeTwoSingleNestedAttributeOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

type ListNestedBlockOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

type ListNestedBlockThreeModel struct {
	ObjectAttribute                        types.Object `tfsdk:"object_attribute"`
	ListNestedBlockThreeListNestedBlockOne types.List   `tfsdk:"list_nested_block_three_list_nested_block_one"`
}

type ListNestedBlockThreeListNestedBlockOneModel struct {
	ListAttribute types.List `tfsdk:"list_attribute"`
}

type ListNestedBlockTwoModel struct {
	ListNestedBlockTwoListNestedBlockOne types.List `tfsdk:"list_nested_block_two_list_nested_block_one"`
}

type ListNestedBlockTwoListNestedBlockOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

type SingleNestedBlockOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

type SingleNestedBlockThreeModel struct {
	ObjectAttribute                          types.Object `tfsdk:"object_attribute"`
	SingleNestedBlockThreeListNestedBlockOne types.List   `tfsdk:"single_nested_block_three_list_nested_block_one"`
}

type SingleNestedBlockThreeListNestedBlockOneModel struct {
	ListAttribute types.List `tfsdk:"list_attribute"`
}

type SingleNestedBlockTwoModel struct {
	SingleNestedBlockTwoSingleNestedBlockOne types.Object `tfsdk:"single_nested_block_two_single_nested_block_one"`
}

type SingleNestedBlockTwoSingleNestedBlockOneModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

func (m ListNestedAttributeOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m ListNestedAttributeOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedAttributeOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m ListNestedAttributeThreeModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeThreeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_nested_attribute_three_list_nested_attribute_one": types.ListType{
			ElemType: ListNestedAttributeThreeListNestedAttributeOneModel{}.ObjectType(ctx),
		},
	}
}

func (m ListNestedAttributeThreeModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedAttributeThreeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m ListNestedAttributeThreeListNestedAttributeOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeThreeListNestedAttributeOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_attribute": types.ListType{
			ElemType: types.StringType,
		},
	}
}

func (m ListNestedAttributeThreeListNestedAttributeOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedAttributeThreeListNestedAttributeOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m ListNestedAttributeTwoModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeTwoModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_nested_attribute_two_list_nested_attribute_one": types.ListType{
			ElemType: ListNestedAttributeTwoListNestedAttributeOneModel{}.ObjectType(ctx),
		},
	}
}

func (m ListNestedAttributeTwoModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedAttributeTwoModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m ListNestedAttributeTwoListNestedAttributeOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedAttributeTwoListNestedAttributeOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m ListNestedAttributeTwoListNestedAttributeOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedAttributeTwoListNestedAttributeOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m SingleNestedAttributeOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m SingleNestedAttributeOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedAttributeOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m SingleNestedAttributeThreeModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeThreeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"single_nested_attribute_three_single_nested_attribute_one": types.ObjectType{
			AttrTypes: SingleNestedAttributeThreeSingleNestedAttributeOneModel{}.ObjectAttributeTypes(ctx),
		},
	}
}

func (m SingleNestedAttributeThreeModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedAttributeThreeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m SingleNestedAttributeThreeSingleNestedAttributeOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeThreeSingleNestedAttributeOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_attribute": types.ListType{
			ElemType: types.StringType,
		},
	}
}

func (m SingleNestedAttributeThreeSingleNestedAttributeOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedAttributeThreeSingleNestedAttributeOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m SingleNestedAttributeTwoModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeTwoModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"single_nested_attribute_two_single_nested_attribute_one": types.ObjectType{
			AttrTypes: SingleNestedAttributeTwoSingleNestedAttributeOneModel{}.ObjectAttributeTypes(ctx),
		},
	}
}

func (m SingleNestedAttributeTwoModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedAttributeTwoModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m SingleNestedAttributeTwoSingleNestedAttributeOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedAttributeTwoSingleNestedAttributeOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m SingleNestedAttributeTwoSingleNestedAttributeOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedAttributeTwoSingleNestedAttributeOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m ListNestedBlockOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m ListNestedBlockOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedBlockOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m ListNestedBlockThreeModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockThreeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_nested_block_three_list_nested_block_one": types.ListType{
			ElemType: ListNestedBlockThreeListNestedBlockOneModel{}.ObjectType(ctx),
		},
		"object_attribute": types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"string_attribute_type": types.StringType,
			},
		},
	}
}

func (m ListNestedBlockThreeModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedBlockThreeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m ListNestedBlockThreeListNestedBlockOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockThreeListNestedBlockOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_attribute": types.ListType{
			ElemType: types.StringType,
		},
	}
}

func (m ListNestedBlockThreeListNestedBlockOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedBlockThreeListNestedBlockOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m ListNestedBlockTwoModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockTwoModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_nested_block_two_list_nested_block_one": types.ListType{
			ElemType: ListNestedBlockTwoListNestedBlockOneModel{}.ObjectType(ctx),
		},
	}
}

func (m ListNestedBlockTwoModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedBlockTwoModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m ListNestedBlockTwoListNestedBlockOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m ListNestedBlockTwoListNestedBlockOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m ListNestedBlockTwoListNestedBlockOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m ListNestedBlockTwoListNestedBlockOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m SingleNestedBlockOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m SingleNestedBlockOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedBlockOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m SingleNestedBlockThreeModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockThreeModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"object_attribute": types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"string_attribute_type": types.StringType,
			},
		},
		"single_nested_block_three_list_nested_block_one": types.ListType{
			ElemType: SingleNestedBlockThreeListNestedBlockOneModel{}.ObjectType(ctx),
		},
	}
}

func (m SingleNestedBlockThreeModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedBlockThreeModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m SingleNestedBlockThreeListNestedBlockOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockThreeListNestedBlockOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"list_attribute": types.ListType{
			ElemType: types.StringType,
		},
	}
}

func (m SingleNestedBlockThreeListNestedBlockOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedBlockThreeListNestedBlockOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
func (m SingleNestedBlockTwoModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockTwoModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"single_nested_block_two_single_nested_block_one": types.ObjectType{
			AttrTypes: SingleNestedBlockTwoSingleNestedBlockOneModel{}.ObjectAttributeTypes(ctx),
		},
	}
}

func (m SingleNestedBlockTwoModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedBlockTwoModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}

func (m SingleNestedBlockTwoSingleNestedBlockOneModel) ObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{AttrTypes: m.ObjectAttributeTypes(ctx)}
}

func (m SingleNestedBlockTwoSingleNestedBlockOneModel) ObjectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m SingleNestedBlockTwoSingleNestedBlockOneModel) ObjectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.ObjectAttributeTypes(ctx),
	)
}

func (m SingleNestedBlockTwoSingleNestedBlockOneModel) ObjectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.ObjectAttributeTypes(ctx),
		data,
	)
}
