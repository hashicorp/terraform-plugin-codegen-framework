{{ define "Create" }}
// Template for generating Terraform provider Create operation code
// Needed data is as follows.
// ResourceName string
// RefreshObjectName string
// CreateReqBody string
// CreateMethod string
// Endpoint string
// CreatePathParams string, optional

func (a *{{.ResourceName | ToCamelCase}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{.RefreshObjectName | ToPascalCase}}Model

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

	response, err := util.MakeRequest("{{.CreateMethod}}", "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{if .CreatePathParams}}{{.CreatePathParams}}{{end}}, strings.Replace(string(reqBody), `\"`, "", -1))
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("CREATING ERROR", "response invalid")
		return
	}

	err = waitResourceCreated(ctx, {{.IdGetter}}, plan)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "Create{{.ResourceName | ToPascalCase}} response="+common.MarshalUncheckedString(response))

	plan.refreshFromOutput(resp.Diagnostics, plan, {{.IdGetter}})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{ end }}