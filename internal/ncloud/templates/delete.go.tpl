{{ define "Delete" }}
// Template for generating Terraform provider Delete operation code
// Needed data is as follows.
// ResourceName string
// DeleteMethod string
// Endpoint string
// DeletePathParams string, optional

func (a *{{.ResourceName | ToCamelCase}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan {{.RefreshObjectName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := util.MakeRequest("{{.DeleteMethod}}", "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{.DeletePathParams}}, "")
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}

	err = waitResourceDeleted(ctx, clearDoubleQuote(plan.ID.String()), plan)
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}
}

{{ end }}