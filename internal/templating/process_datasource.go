package templating

import (
	"fmt"
)

func (t *templator) ProcessDataSources(templateData map[string]DataSourceTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(templateData))

	for name, dataSourceData := range templateData {
		// Process data source template
		dataSourceTemplate := fmt.Sprintf("%s_datasource", name)

		dataSourceBytes, err := t.processTemplateWithDefault(dataSourceTemplate, dataSourceData, t.defaultDataSourceBytes)
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error processing %q data source template: %s, skipping", name, err))
			continue
		}
		dataSourceOutputName := fmt.Sprintf("%s_datasource_gen.go", name)
		outputData[dataSourceOutputName] = dataSourceBytes

		// Process data source test template
		testTemplate := fmt.Sprintf("%s_datasource_test", name)

		testBytes, err := t.processTemplateWithDefault(testTemplate, dataSourceData, t.defaultDataSourceTestBytes)
		if err != nil {
			t.logger.Debug(fmt.Sprintf("error processing %q data source test template: %s, skipping", name, err))
			continue
		}
		testOutputName := fmt.Sprintf("%s_datasource_gen_test.go", name)
		outputData[testOutputName] = testBytes
	}

	return outputData, nil
}
