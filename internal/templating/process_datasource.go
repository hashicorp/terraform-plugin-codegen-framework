package templating

import (
	"bytes"
	"fmt"
	"io/fs"
	"text/template"
)

func (t *templator) hasDefaultDataSource() bool {
	return len(t.defaultDataSourceBytes) != 0
}

func (t *templator) ProcessDataSources(templateData map[string]DataSourceTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(templateData))

	// Process each set of data source template data
	for name, datasourceData := range templateData {
		templateBytes, err := fs.ReadFile(t.templateDir, fmt.Sprintf("%s_datasource.gotmpl", name))
		if err != nil {
			if !t.hasDefaultDataSource() {
				// TODO: no default, skipping - log
				continue
			}
			// TODO: found default, using - log
			templateBytes = t.defaultDataSourceBytes
		}

		tmpl := template.New(name)
		datasourceTemplate, err := tmpl.Parse(string(templateBytes))
		if err != nil {
			// TODO: log
			continue
		}

		var buf bytes.Buffer
		err = datasourceTemplate.Execute(&buf, datasourceData)
		if err != nil {
			// TODO: log
			continue
		}

		outputName := fmt.Sprintf("%s_datasource_gen.go", name)
		outputData[outputName] = buf.Bytes()
	}

	return outputData, nil
}
