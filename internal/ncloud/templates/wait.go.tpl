{{ define "Wait" }}
// Template for generating Terraform provider Waiting operation code
// Needed data is as follows.
// RefreshObjectName string
// ReadMethod string
// Endpoint string
// ReadPathParams string, optional

func waitResourceCreated(ctx context.Context, id string, plan {{.RefreshObjectName | ToPascalCase}}Model) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATING"},
		Target:  []string{"CREATED"},
		Refresh: func() (interface{}, string, error) {
			response, err := util.MakeRequest("{{.ReadMethod}}", "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+clearDoubleQuote(id){{end}}, "")
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

func waitResourceDeleted(ctx context.Context, id string, plan {{.RefreshObjectName | ToPascalCase}}Model) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"DELETING"},
		Target:  []string{"DELETED"},
		Refresh: func() (interface{}, string, error) {
			response, _ := util.MakeRequest("{{.ReadMethod}}", "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+clearDoubleQuote(id){{end}}, "")
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