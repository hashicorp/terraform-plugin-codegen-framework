package templating

import (
	"bytes"
	"io/fs"
	"text/template"
)

func (t *templator) ProcessProvider(templateData ProviderTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, 1)

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
	err = providerTemplate.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	outputData["provider_gen.go"] = buf.Bytes()

	return outputData, nil
}
