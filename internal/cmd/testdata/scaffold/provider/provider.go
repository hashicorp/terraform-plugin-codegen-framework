package scaffold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*examplecloudProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &examplecloudProvider{}
	}
}

type examplecloudProvider struct{}

func (p *examplecloudProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

}

func (p *examplecloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *examplecloudProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "examplecloud"
}

func (p *examplecloudProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *examplecloudProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
