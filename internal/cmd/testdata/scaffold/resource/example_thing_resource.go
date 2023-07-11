package scaffold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = (*ExampleThingResource)(nil)

func NewexampleThingResource() resource.Resource {
	return &ExampleThingResource{}
}

type ExampleThingResource struct{}

func (r *ExampleThingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example_thing"
}

func (r *ExampleThingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

}

func (r *ExampleThingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

}

func (r *ExampleThingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *ExampleThingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *ExampleThingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
