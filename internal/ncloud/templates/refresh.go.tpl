{{ define "Refresh" }}
// Template for generating Terraform provider Refresh operation code
// Needed data is as follows.
// DtoName string
// RefreshLogic string
// ReadMethod string
// Endpoint string
// ReadPathParams string, optional

func ConvertToFrameworkTypes(data map[string]interface{}) (*{{.DtoName | ToPascalCase}}Model, error) {
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

func getAndRefresh(diagnostics diag.Diagnostics, plan {{.DtoName | ToPascalCase}}Model) *{{.DtoName | ToPascalCase}}Model {
	getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "{{.ReadMethod}}", "{{.Endpoint}}"{{if .ReadPathParams}}+plan.{{.ReadPathParams | ToPascalCase}}.String(){{end}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
		)
	}

	response, err := util.Request(getExecFunc, "")
	if err != nil {
		diagnostics.AddError("UPDATING ERROR", err.Error())
		return nil
	}
	if response == nil {
		diagnostics.AddError("UPDATING ERROR", "response invalid")
		return nil
	}

	newPlan, err := ConvertToFrameworkTypes(response)
	if err != nil {
		diagnostics.AddError("CREATING ERROR", err.Error())
		return nil
	}

	return newPlan
}

{{ end }}