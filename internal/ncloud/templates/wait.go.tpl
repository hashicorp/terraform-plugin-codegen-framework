{{ define "Wait" }}
// Template for generating Terraform provider Waiting operation code
// Needed data is as follows.
// DtoName string
// ReadMethod string
// Endpoint string
// ReadPathParams string, optional

func waitResourceCreated(ctx context.Context, id string) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATING"},
		Target:  []string{"CREATED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
			return exec.Command("curl", "-s", "-X", "{{.ReadMethod}}",  "{{.Endpoint}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+util.ClearDoubleQuote(id){{end}},
					"-H", "accept: application/json;charset=UTF-8",
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
					"-H", "cache-control: no-cache",
					"-H", "pragma: no-cache",
				)
			}

			response, err := util.Request(getExecFunc, "{{.ReadMethod}}","{{.Endpoint | ExtractPath}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+util.ClearDoubleQuote(id){{end}}, os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
			if err != nil {
				return response, "CREATING", nil
			}
			if response != nil {
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

func waitResourceDeleted(ctx context.Context, id string) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"DELETING"},
		Target:  []string{"DELETED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
			return exec.Command("curl", "-s", "-X", "{{.ReadMethod}}", "{{.Endpoint}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+util.ClearDoubleQuote(id){{end}},
					"-H", "accept: application/json;charset=UTF-8",
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
					"-H", "cache-control: no-cache",
					"-H", "pragma: no-cache",
				)
			}

			response, _ := util.Request(getExecFunc, "{{.ReadMethod}}", "{{.Endpoint | ExtractPath}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+util.ClearDoubleQuote(id){{end}}, os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
			if response["error"] != nil {
				return response, "DELETED", nil
			}

			return response, "DELETING", nil
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

{{ end }}