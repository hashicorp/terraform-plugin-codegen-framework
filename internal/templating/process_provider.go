package templating

import (
	"bytes"
	"io/fs"
)

func (t *templator) ProcessProvider(templateData map[string]ProviderTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(templateData))

	// TODO: swap to single provider processing (everywhere else is a map currently)
	for _, providerData := range templateData {
		templateBytes, err := fs.ReadFile(t.templateDir, "provider.gotmpl")
		if err != nil {
			// TODO: log
			continue
		}

		tmpl, err := t.baseTemplate.Clone()
		if err != nil {
			// TODO: log
			continue
		}

		providerTemplate, err := tmpl.Parse(string(templateBytes))
		if err != nil {
			// TODO: log
			continue
		}

		var buf bytes.Buffer
		err = providerTemplate.Execute(&buf, providerData)
		if err != nil {
			// TODO: log
			continue
		}

		outputData["provider_gen.go"] = buf.Bytes()
	}

	return outputData, nil
}
