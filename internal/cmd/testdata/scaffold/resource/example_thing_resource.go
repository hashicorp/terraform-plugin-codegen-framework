// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scaffold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = (*exampleThingResource)(nil)

func NewExampleThingResource() resource.Resource {
	return &exampleThingResource{}
}

type exampleThingResource struct{}

func (r *exampleThingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example_thing"
}

func (r *exampleThingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

}

func (r *exampleThingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

}

func (r *exampleThingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *exampleThingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *exampleThingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
