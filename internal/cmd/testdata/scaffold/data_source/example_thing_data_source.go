package scaffold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ExampleThingDataSource)(nil)

func NewexampleThingDataSource() datasource.DataSource {
	return &ExampleThingDataSource{}
}

type ExampleThingDataSource struct{}

func (r *ExampleThingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example_thing"
}

func (r *ExampleThingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {

}

func (r *ExampleThingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

}
