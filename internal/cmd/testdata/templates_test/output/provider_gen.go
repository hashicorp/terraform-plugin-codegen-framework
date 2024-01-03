package provider

import (
	"context"
	"terraform-provider-petstore/internal/provider/provider_petstore"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*petstoreProvider)(nil)

func NewPetstoreProvider() func() provider.Provider {
	return func() provider.Provider {
		return &petstoreProvider{}
	}
}

type petstoreProvider struct{}

func (p *petstoreProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "petstore"
}

func (p *petstoreProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = provider_petstore.PetstoreProviderSchema(ctx)
}

func (p *petstoreProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *petstoreProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOrderDataSource,
	}
}

func (p *petstoreProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPetResource,
		NewUserResource,
	}
}
