package provider

import (
	"context"
	"terraform-provider-petstore/internal/provider/resource_pet"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)


var _ resource.Resource = (*petResource)(nil)

func NewPetResource() resource.Resource {
    return &petResource{}
}

type petResource struct{}

func (r *petResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_pet"
}

func (r *petResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = resource_pet.PetResourceSchema(ctx)
}

func (r *petResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *petResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *petResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *petResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }
}

