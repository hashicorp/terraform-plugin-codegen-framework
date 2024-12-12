{{ define "Update" }}
// Template for generating Terraform provider Update operation code
// Needed data is as follows.
// ResourceName string
// RefreshObjectName string
// UpdateReqBody string
// UpdateMethod string
// Endpoint string
// UpdatePathParams string, optional

func (a *{{.ResourceName | ToCamelCase}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan {{.RefreshObjectName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody, err := json.Marshal(map[string]string{
		{{.UpdateReqBody}}
	})
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "Update{{.ResourceName | ToPascalCase}} reqParams="+strings.Replace(string(reqBody), `\"`, "", -1))

	response, err := util.MakeRequest("{{.UpdateMethod}}",  "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{if .UpdatePathParams}}{{.UpdatePathParams}}{{end}}, strings.Replace(string(reqBody), `\"`, "", -1))
	if err != nil {
		resp.Diagnostics.AddError("UPDATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("UPDATING ERROR", "response invalid")
		return
	}

	tflog.Info(ctx, "Update{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan.refreshFromOutput(resp.Diagnostics, plan, plan.ID.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{ end }}