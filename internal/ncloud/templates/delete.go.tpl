{{ define "Delete" }}
// Template for generating Terraform provider Delete operation code
// Needed data is as follows.
// ResourceName string
// DeleteMethod string
// Endpoint string
// DeletePathParams string, optional

func (a *{{.ResourceName | ToCamelCase}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan {{.DtoName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "{{.DeleteMethod}}", "{{.Endpoint}}"{{if .DeletePathParams}}+plan.{{.DeletePathParams | ToPascalCase}}.String(){{end}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
		)
	}

	response, err := util.Request(execFunc, "")
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("DELETING ERROR", "response invalid")
		return
	}

	err = waitResourceDeleted(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "Delete{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{ end }}