package templating

import (
	"fmt"
	"io/fs"
	"log/slog"
	"text/template"
)

const (
	default_resource_template   = "resource_default.gotmpl"
	default_datasource_template = "datasource_default.gotmpl"
)

type templator struct {
	baseTemplate           *template.Template
	logger                 *slog.Logger
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
	Resources       []ResourceTemplateData
	DataSources     []DataSourceTemplateData
}

type Templator interface {
	ProcessResources(map[string]ResourceTemplateData) (map[string][]byte, error)
	ProcessDataSources(map[string]DataSourceTemplateData) (map[string][]byte, error)
	ProcessProvider(map[string]ProviderTemplateData, map[string]ResourceTemplateData, map[string]DataSourceTemplateData) (map[string][]byte, error)
}

func NewTemplator(logger *slog.Logger, templateDir fs.FS) Templator {
	templator := &templator{
		logger:      logger,
		templateDir: templateDir,
	}

	defaultResourceBytes, err := fs.ReadFile(templateDir, default_resource_template)
	if err != nil {
		templator.logger.Debug("no default resource template found")
	} else {
		templator.defaultResourceBytes = defaultResourceBytes
	}

	defaultDataSourceBytes, err := fs.ReadFile(templateDir, default_datasource_template)
	if err != nil {
		templator.logger.Debug("no default datasource template found")
	} else {
		templator.defaultDataSourceBytes = defaultDataSourceBytes
	}

	tmpl, err := addBuiltInTemplates(template.New("base"))
	if err != nil {
		templator.logger.Warn(fmt.Sprintf("error loading built-in templates: %s", err))
	}

	templator.baseTemplate = tmpl

	return templator
}
