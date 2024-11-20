{{ define "Initial" }}
// Template for generating Terraform provider Initial code
// Needed data is as follows.
// ResourceName string

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

func (a *{{.ResourceName | ToCamelCase}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{.ProviderName}}_{{.ResourceName}}"
}

func (a *{{.ResourceName | ToCamelCase}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = {{.ResourceName | ToPascalCase}}ResourceSchema(ctx)
}

{{ end }}