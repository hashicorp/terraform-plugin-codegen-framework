{{ define "Read" }}
// Template for generating Terraform provider Read operation code
// Needed data is as follows.
// ResourceName string

func (a *{{.ResourceName | ToCamelCase}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan {{.DtoName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan = *getAndRefresh(resp.Diagnostics, plan.ID.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{ end }}