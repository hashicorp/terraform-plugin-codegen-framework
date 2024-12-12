{{ define "ImportState" }}
// Template for generating Terraform provider Initial code
// Needed data is as follows.
// ResourceName string
// ImportStateLogic string

func (a *{{.ResourceName | ToCamelCase}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	{{.ImportStateLogic}}
}

{{ end }}