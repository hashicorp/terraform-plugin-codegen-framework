package scaffold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*exampleThingDataSource)(nil)

func NewExampleThingDataSource() datasource.DataSource {
	return &exampleThingDataSource{}
}

type exampleThingDataSource struct{}

func (d *exampleThingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example_thing"
}

func (d *exampleThingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {

}

func (d *exampleThingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

}
