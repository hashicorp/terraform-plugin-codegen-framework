{{ define "Update" }}
// Template for generating Terraform provider Update operation code
// Needed data is as follows.
// ResourceName string
// DtoName string
// UpdateReqBody string
// UpdateMethod string
// Endpoint string
// UpdatePathParams string, optional

func (a *{{.ResourceName | ToCamelCase}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan {{.DtoName | ToPascalCase}}Model

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

	tflog.Info(ctx, "Update{{.ResourceName | ToPascalCase}} reqParams="+string(reqBody))

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "{{.UpdateMethod}}", "{{.Endpoint}}"{{if .UpdatePathParams}}+plan.{{.UpdatePathParams | ToPascalCase}}.String(){{end}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-d", string(reqBody),
		)
	}

	response, err := util.Request(execFunc, string(reqBody))
	if err != nil {
		resp.Diagnostics.AddError("UPDATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("UPDATING ERROR", "response invalid")
		return
	}

	tflog.Info(ctx, "Update{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{ end }}