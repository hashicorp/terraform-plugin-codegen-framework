package templating

import (
	"fmt"
	"io/fs"
	"text/template"
)

const (
	default_resource_template   = "resource_default.gotmpl"
	default_datasource_template = "datasource_default.gotmpl"
)

type templator struct {
	baseTemplate           *template.Template
	templateDir            fs.FS
	defaultResourceBytes   []byte
	defaultDataSourceBytes []byte
}

type DataSourceTemplateData struct {
	SnakeName       string
	PascalName      string
	CamelName       string
	Package         string
	SchemaFunc      string
	SchemaModelType string
}

type ResourceTemplateData struct {
	SnakeName       string
	PascalName      string
	CamelName       string
	Package         string
	SchemaFunc      string
	SchemaModelType string
}

type ProviderTemplateData struct {
	SnakeName       string
	PascalName      string
	CamelName       string
	Package         string
	SchemaFunc      string
	SchemaModelType string
}

type Templator interface {
	ProcessResources(templateData map[string]ResourceTemplateData) (map[string][]byte, error)
	ProcessDataSources(templateData map[string]DataSourceTemplateData) (map[string][]byte, error)
	ProcessProvider(templateData map[string]ProviderTemplateData) (map[string][]byte, error)
}

func NewTemplator(templateDir fs.FS) Templator {
	templator := &templator{
		templateDir: templateDir,
	}

	// Check for defaults
	defaultResourceBytes, err := fs.ReadFile(templateDir, default_resource_template)
	if err != nil {
		// TODO: log
	} else {
		templator.defaultResourceBytes = defaultResourceBytes
	}

	// Check for defaults
	defaultDataSourceBytes, err := fs.ReadFile(templateDir, default_datasource_template)
	if err != nil {
		// TODO: log
	} else {
		templator.defaultDataSourceBytes = defaultDataSourceBytes
	}

	// Add built-in templates
	tmpl, err := addBuiltInTemplates(template.New("base"))
	if err != nil {
		// TODO: log
		fmt.Println(err)
	}

	templator.baseTemplate = tmpl

	return templator
}
