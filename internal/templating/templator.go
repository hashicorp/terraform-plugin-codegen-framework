package templating

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"text/template"
)

const (
	default_resource_template        = "resource_default.gotmpl"
	default_resource_test_template   = "resource_default_test.gotmpl"
	default_datasource_template      = "datasource_default.gotmpl"
	default_datasource_test_template = "datasource_default_test.gotmpl"
)

type templator struct {
	baseTemplate               *template.Template
	logger                     *slog.Logger
	templateDir                fs.FS
	defaultResourceBytes       []byte
	defaultResourceTestBytes   []byte
	defaultDataSourceBytes     []byte
	defaultDataSourceTestBytes []byte
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

	defaultResourceTestBytes, err := fs.ReadFile(templateDir, default_resource_test_template)
	if err != nil {
		templator.logger.Debug("no default resource test template found")
	} else {
		templator.defaultResourceTestBytes = defaultResourceTestBytes
	}

	defaultDataSourceBytes, err := fs.ReadFile(templateDir, default_datasource_template)
	if err != nil {
		templator.logger.Debug("no default datasource template found")
	} else {
		templator.defaultDataSourceBytes = defaultDataSourceBytes
	}

	defaultDataSourceTestBytes, err := fs.ReadFile(templateDir, default_datasource_test_template)
	if err != nil {
		templator.logger.Debug("no default datasource test template found")
	} else {
		templator.defaultDataSourceTestBytes = defaultDataSourceTestBytes
	}

	tmpl, err := addBuiltInTemplates(template.New("base"))
	if err != nil {
		templator.logger.Warn(fmt.Sprintf("error loading built-in templates: %s", err))
	}

	templator.baseTemplate = tmpl

	return templator
}

func (t *templator) processTemplateWithDefault(templateName string, data any, defaultTemplateBytes []byte) ([]byte, error) {
	templateBytes, err := fs.ReadFile(t.templateDir, fmt.Sprintf("%s.gotmpl", templateName))
	if err != nil {
		if len(defaultTemplateBytes) == 0 {
			return nil, errors.New("no user-defined or default template found")
		}

		templateBytes = defaultTemplateBytes
	}

	return t.processTemplate(templateBytes, data)
}

func (t *templator) processTemplate(templateBytes []byte, data any) ([]byte, error) {
	tmpl, err := t.baseTemplate.Clone()
	if err != nil {
		return nil, fmt.Errorf("error cloning base template with built-ins: %w", err)
	}

	resourceTemplate, err := tmpl.Parse(string(templateBytes))
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %w", err)
	}

	var buf bytes.Buffer
	err = resourceTemplate.Execute(&buf, data)
	if err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	return buf.Bytes(), nil
}
