package scaffold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*exampleThingProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &exampleThingProvider{}
	}
}

type exampleThingProvider struct{}

func (p *exampleThingProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

}

func (p *exampleThingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *exampleThingProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "example_thing"
}

func (p *exampleThingProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *exampleThingProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
