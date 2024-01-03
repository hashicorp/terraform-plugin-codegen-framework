package provider

import (
	"context"
	"terraform-provider-petstore/internal/provider/datasource_order"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*orderDataSource)(nil)

func NewOrderDataSource() datasource.DataSource {
	return &orderDataSource{}
}

type orderDataSource struct{}

func (d *orderDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_order"
}

func (d *orderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_order.OrderDataSourceSchema(ctx)
}

func (d *orderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_order.OrderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
