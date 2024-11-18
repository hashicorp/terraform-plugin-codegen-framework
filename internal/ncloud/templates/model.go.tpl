{{ define "Model" }}
// Template for generating Terraform provider Model code
// Needed data is as follows.
// DtoName string
// Model string

type {{.DtoName | ToPascalCase}}Model struct {
    {{.Model}}
}

{{ end }}