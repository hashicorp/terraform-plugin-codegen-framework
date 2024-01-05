package templating

import (
	"bytes"
	"fmt"
	"io/fs"
)

func (t *templator) hasDefaultResource() bool {
	return len(t.defaultResourceBytes) != 0
}

func (t *templator) ProcessResources(templateData map[string]ResourceTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(templateData))

	for name, resourceData := range templateData {
		templateBytes, err := fs.ReadFile(t.templateDir, fmt.Sprintf("%s_resource.gotmpl", name))
		if err != nil {
			if !t.hasDefaultResource() {
				t.logger.Debug(fmt.Sprintf("no resource or default template found for %q, skipping", name))
				continue
			}
			t.logger.Debug(fmt.Sprintf("no resource template found for %q, using default template", name))
			templateBytes = t.defaultResourceBytes
		}

		tmpl, err := t.baseTemplate.Clone()
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error cloning base template with built-ins: %s, skipping", err))
			continue
		}

		resourceTemplate, err := tmpl.Parse(string(templateBytes))
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error parsing  %q resource template: %s, skipping", name, err))
			continue
		}

		var buf bytes.Buffer
		err = resourceTemplate.Execute(&buf, resourceData)
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error executing %q resource template: %s, skipping", name, err))
			continue
		}

		outputName := fmt.Sprintf("%s_resource_gen.go", name)
		outputData[outputName] = buf.Bytes()
	}

	return outputData, nil
}
