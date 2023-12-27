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

	// Process each set of resource template data
	for name, resourceData := range templateData {
		templateBytes, err := fs.ReadFile(t.templateDir, fmt.Sprintf("%s_resource.gotmpl", name))
		if err != nil {
			if !t.hasDefaultResource() {
				// TODO: no default, skipping - log
				continue
			}
			// TODO: found default, using - log
			templateBytes = t.defaultResourceBytes
		}

		tmpl, err := t.baseTemplate.Clone()
		if err != nil {
			// TODO: log
			continue
		}

		resourceTemplate, err := tmpl.Parse(string(templateBytes))
		if err != nil {
			// TODO: log
			continue
		}

		var buf bytes.Buffer
		err = resourceTemplate.Execute(&buf, resourceData)
		if err != nil {
			// TODO: log
			continue
		}

		outputName := fmt.Sprintf("%s_resource_gen.go", name)
		outputData[outputName] = buf.Bytes()
	}

	return outputData, nil
}
