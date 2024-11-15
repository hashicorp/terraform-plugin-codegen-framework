{{ define "Main" }}

var (
	_ resource.Resource                = &{{.ResourceName | ToCamelCase}}Resource{}
	_ resource.ResourceWithConfigure   = &{{.ResourceName | ToCamelCase}}Resource{}
	_ resource.ResourceWithImportState = &{{.ResourceName | ToCamelCase}}Resource{}
)

func New{{.ResourceName | ToPascalCase}}Resource() resource.Resource {
	return &{{.ResourceName | ToCamelCase}}Resource{}
}

type {{.ResourceName | ToCamelCase}}Resource struct {
	config *conn.ProviderConfig
}

func (a *{{.ResourceName | ToCamelCase}}Resource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*conn.ProviderConfig)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ProviderConfig, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	a.config = config
}

func (a *{{.ResourceName | ToCamelCase}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

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

	err = waitResourceDeleted(ctx)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}
	
	tflog.Info(ctx, "Delete{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *{{.ResourceName | ToCamelCase}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_key" //임의로 space 넣어줘야 함 (_)가 필요하기 때문
}

func (a *{{.ResourceName | ToCamelCase}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan {{.DtoName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *{{.ResourceName | ToCamelCase}}Resource) Schema(context.Context, resource.SchemaRequest, *resource.SchemaResponse) {
	panic("unimplemented") // WIP
}

func (a *{{.ResourceName | ToCamelCase}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan {{.DtoName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := fmt.Sprintf(`{
		"isEnabled": %[1]s,
		"apiKeyDescription": %[2]s,
		"apiKeyName": %[3]s
	}`, plan.IsEnabled, plan.ApiKeyDescription, plan.ApiKeyName) // request에만 들어가있는 애들의 경우 dto에 없으므로 추가 수정 필요.

	tflog.Info(ctx, "Update{{.ResourceName | ToPascalCase}} reqParams="+reqBody)

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "{{.UpdateMethod}}", "{{.Endpoint}}"{{if .UpdatePathParams}}+plan.{{.UpdatePathParams | ToPascalCase}}.String(){{end}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-d", reqBody,
		)
	}

	response, err := util.Request(execFunc, reqBody)
	if err != nil {
		resp.Diagnostics.AddError("UPDATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("UPDATING ERROR", "response invalid")
		return
	}

	tflog.Info(ctx, "Update{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *{{.ResourceName | ToCamelCase}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{.DtoName | ToPascalCase}}Model

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := fmt.Sprintf(`{
			"apiKeyDescription": %[1]s,
			"apiKeyName": %[2]s
		}`, plan.ApiKeyDescription, plan.ApiKeyName) // 마찬가지로 수기 필요.

	tflog.Info(ctx, "Create{{.ResourceName | ToPascalCase}} reqParams="+reqBody)

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "{{.CreateMethod}}", "{{.Endpoint}}"{{if .CreatePathParams}}+plan.{{.CreatePathParams | ToPascalCase}}.String(){{end}},
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-d", reqBody,
		)
	}

	response, err := util.Request(execFunc, reqBody)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("CREATING ERROR", "response invalid")
		return
	}

	err = waitResourceCreated(ctx)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "Create{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

type {{.DtoName | ToPascalCase}}Model struct {
    {{.Model}}
}

func ConvertToFrameworkTypes(data map[string]interface{}) (*{{.DtoName | ToPascalCase}}Model, error) {
	var dto {{.DtoName | ToPascalCase}}Model

    {{.RefreshLogic}}

	return &dto, nil
}

func waitResourceCreated(ctx context.Context) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATING"},
		Target:  []string{"CREATED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
				return exec.Command("curl", "-X", "{{.ReadMethod}}", "{{.Endpoint}}"{{if .CreatePathParams}}+plan.{{.CreatePathParams | ToPascalCase}}.String(){{end}},
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
				)
			}

			response, err := util.Request(getExecFunc, "")
			if err != nil {
				return response, "CREATING", nil
			}
			if response == nil {
				return response, "CREATED", nil
			}

			return response, "CREATING", nil
		},
		Timeout:    conn.DefaultTimeout,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("error occured while waiting for resource to be created: %s", err)
	}
	return nil
}

func waitResourceDeleted(ctx context.Context) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"DELETING"},
		Target:  []string{"DELETED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
				return exec.Command("curl", "-X", "{{.ReadMethod}}", "{{.Endpoint}}"{{if .CreatePathParams}}+plan.{{.CreatePathParams | ToPascalCase}}.String(){{end}},
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
				)
			}

			response, err := util.Request(getExecFunc, "")
			if response != nil {
				return response, "DELETING", nil
			}

			if err != nil {
				return response, "DELETED", nil
			}

			return response, "DELETED", nil
		},
		Timeout:    conn.DefaultTimeout,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("error occured while waiting for resource to be deleted: %s", err)
	}
	return nil
}

func diagOff[V, T interface{}](input func(ctx context.Context, elementType T, elements any) (V, diag.Diagnostics), ctx context.Context, elementType T, elements any) V {
	var emptyReturn V

	v, diags := input(ctx, elementType, elements)

	if diags.HasError() {
		diags.AddError("REFRESING ERROR", "invalid diagOff operation")
		return emptyReturn
	}

	return v
}

func getAndRefresh(diagnostics diag.Diagnostics) *{{.DtoName | ToPascalCase}}Model {
	getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "{{.ReadMethod}}", "{{.Endpoint}}"{{if .CreatePathParams}}+plan.{{.CreatePathParams | ToPascalCase}}.String(){{end}},
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
	// WIP
	// err = waitBucketCreated(ctx, o.config, plan.BucketName.String())
	// if err != nil {
	// 	resp.Diagnostics.AddError("CREATING ERROR", err.Error())
	// 	return
	// }
	newPlan, err := ConvertToFrameworkTypes(response)
	if err != nil {
		diagnostics.AddError("CREATING ERROR", err.Error())
		return nil
	}

	return newPlan
}

{{ end }}