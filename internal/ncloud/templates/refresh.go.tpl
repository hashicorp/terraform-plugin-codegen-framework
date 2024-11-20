{{ define "Refresh" }}
// Template for generating Terraform provider Refresh operation code
// Needed data is as follows.
// DtoName string
// RefreshLogic string
// ReadMethod string
// Endpoint string
// ReadPathParams string, optional

func ConvertToFrameworkTypes(data map[string]interface{}, rest []interface{}) (*{{.DtoName | ToPascalCase}}Model, error) {
	var dto {{.DtoName | ToPascalCase}}Model

    {{.RefreshLogic}}

	return &dto, nil
}

func diagOff[V, T interface{}](input func(ctx context.Context, elementType T, elements any) (V, diag.Diagnostics), ctx context.Context, elementType T, elements any) V {
	var emptyReturn V

	v, diags := input(ctx, elementType, elements)

	if diags.HasError() {
		diags.AddError("REFRESHING ERROR", "invalid diagOff operation")
		return emptyReturn
	}

	return v
}

func getAndRefresh(diagnostics diag.Diagnostics, id string, rest ...interface{}) *{{.DtoName | ToPascalCase}}Model {
	getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-s", "-X", "{{.ReadMethod}}", "{{.Endpoint}}"{{if .ReadPathParams}}+"/"+id{{end}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-H", "cache-control: no-cache",
			"-H", "pragma: no-cache",
		)
	}

	response, err := util.Request(getExecFunc, "{{.ReadMethod}}", "{{.Endpoint | ExtractPath}}"{{if .ReadPathParams}}+"/"+id{{end}}, os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
	if err != nil {
		diagnostics.AddError("UPDATING ERROR", err.Error())
		return nil
	}
	if response == nil {
		diagnostics.AddError("UPDATING ERROR", "response invalid")
		return nil
	}

	newPlan, err := ConvertToFrameworkTypes(util.ConvertKeys(response).(map[string]interface{}), rest)
	if err != nil {
		diagnostics.AddError("CREATING ERROR", err.Error())
		return nil
	}

	return newPlan
}

{{ end }}