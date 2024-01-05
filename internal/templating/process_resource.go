package templating

import (
	"fmt"
)

func (t *templator) ProcessResources(templateData map[string]ResourceTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(templateData))

	for name, resourceData := range templateData {
		// Process resource template
		resourceTemplate := fmt.Sprintf("%s_resource", name)

		resourceBytes, err := t.processTemplateWithDefault(resourceTemplate, resourceData, t.defaultResourceBytes)
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error processing %q resource template: %s, skipping", name, err))
			continue
		}
		resourceOutputName := fmt.Sprintf("%s_resource_gen.go", name)
		outputData[resourceOutputName] = resourceBytes

		// Process resource test template
		testTemplate := fmt.Sprintf("%s_resource_test", name)

		testBytes, err := t.processTemplateWithDefault(testTemplate, resourceData, t.defaultResourceTestBytes)
		if err != nil {
			t.logger.Debug(fmt.Sprintf("error processing %q resource test template: %s, skipping", name, err))
			continue
		}
		testOutputName := fmt.Sprintf("%s_resource_gen_test.go", name)
		outputData[testOutputName] = testBytes
	}

	return outputData, nil
}
