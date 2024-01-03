package templating

import (
	"bytes"
	"io/fs"
	"sort"
)

func (t *templator) ProcessProvider(providerTemplateData map[string]ProviderTemplateData, resourceTemplateData map[string]ResourceTemplateData, dataSourceTemplateData map[string]DataSourceTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(providerTemplateData))

	sortedResources := mapToSortedSlice(resourceTemplateData)
	sortedDataSources := mapToSortedSlice(dataSourceTemplateData)

	// TODO: swap to single provider processing (everywhere else is a map currently)
	for _, providerData := range providerTemplateData {
		providerData.Resources = sortedResources
		providerData.DataSources = sortedDataSources

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

func mapToSortedSlice[T any](m map[string]T) []T {
	slice := make([]T, len(m))
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i, k := range keys {
		slice[i] = m[k]
	}

	return slice
}
