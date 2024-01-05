package templating

import (
	"bytes"
	"fmt"
	"io/fs"
)

func (t *templator) hasDefaultDataSource() bool {
	return len(t.defaultDataSourceBytes) != 0
}

func (t *templator) ProcessDataSources(templateData map[string]DataSourceTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(templateData))

	for name, datasourceData := range templateData {
		templateBytes, err := fs.ReadFile(t.templateDir, fmt.Sprintf("%s_datasource.gotmpl", name))
		if err != nil {
			if !t.hasDefaultDataSource() {
				t.logger.Debug(fmt.Sprintf("no data source or default template found for %q, skipping", name))
				continue
			}
			t.logger.Debug(fmt.Sprintf("no data source template found for %q, using default template", name))
			templateBytes = t.defaultDataSourceBytes
		}

		tmpl, err := t.baseTemplate.Clone()
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error cloning base template with built-ins: %s, skipping", err))
			continue
		}

		datasourceTemplate, err := tmpl.Parse(string(templateBytes))
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error parsing  %q data source template: %s, skipping", name, err))
			continue
		}

		var buf bytes.Buffer
		err = datasourceTemplate.Execute(&buf, datasourceData)
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error executing %q data source template: %s, skipping", name, err))
			continue
		}

		outputName := fmt.Sprintf("%s_datasource_gen.go", name)
		outputData[outputName] = buf.Bytes()
	}

	return outputData, nil
}
