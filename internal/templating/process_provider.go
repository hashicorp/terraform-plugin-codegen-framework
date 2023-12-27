package templating

import (
	"bytes"
	"io/fs"
	"text/template"
)

func (t *templator) ProcessProvider(templateData map[string]ProviderTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(templateData))

	// TODO: swap to single provider processing (everywhere else is a map currently)
	for _, providerData := range templateData {
		templateBytes, err := fs.ReadFile(t.templateDir, "provider.gotmpl")
		if err != nil {
			return nil, err
		}

		tmpl := template.New("provider")
		providerTemplate, err := tmpl.Parse(string(templateBytes))
		if err != nil {
			return nil, err
		}

		var buf bytes.Buffer
		err = providerTemplate.Execute(&buf, providerData)
		if err != nil {
			return nil, err
		}

		outputData["provider_gen.go"] = buf.Bytes()
	}

	return outputData, nil
}
