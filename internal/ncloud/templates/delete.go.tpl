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
		return exec.Command("curl", "-s", "-X", "{{.DeleteMethod}}", "{{.Endpoint}}"{{.DeletePathParams}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-H", "cache-control: no-cache",
			"-H", "pragma: no-cache",
		)
	}

	_, err := util.Request(execFunc, "{{.DeleteMethod}}", "{{.Endpoint | ExtractPath}}"{{.DeletePathParams}}, os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}

	err = waitResourceDeleted(ctx, util.ClearDoubleQuote(plan.ID.String()))
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}
}

{{ end }}