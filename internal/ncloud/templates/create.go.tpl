{{ define "Create" }}
// Template for generating Terraform provider Create operation code
// Needed data is as follows.
// ResourceName string
// DtoName string
// CreateReqBody string
// CreateMethod string
// Endpoint string
// CreatePathParams string, optional

func (a *{{.ResourceName | ToCamelCase}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{.DtoName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody, err := json.Marshal(map[string]string{
		{{.CreateReqBody}}
	})
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "Create{{.ResourceName | ToPascalCase}} reqParams="+strings.Replace(string(reqBody), `\"`, "", -1))

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-s", "-X", "{{.CreateMethod}}", "{{.Endpoint}}"{{if .CreatePathParams}}+plan.{{.CreatePathParams | ToPascalCase}}.String(){{end}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-H", "cache-control: no-cache",
			"-H", "pragma: no-cache",
			"-d", strings.Replace(string(reqBody), `\"`, "", -1),
		)
	}

	response, err := util.Request(execFunc, "{{.CreateMethod}}", "{{.Endpoint | ExtractPath}}"{{if .CreatePathParams}}+plan.{{.CreatePathParams | ToPascalCase}}.String(){{end}}, os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), strings.Replace(string(reqBody), `\"`, "", -1))
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("CREATING ERROR", "response invalid")
		return
	}

	err = waitResourceCreated(ctx, {{.IdGetter}})
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "Create{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, {{.IdGetter}})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{ end }}