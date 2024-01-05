package templating

import (
	"fmt"
	"io/fs"
	"sort"
)

func (t *templator) ProcessProvider(providerTemplateData map[string]ProviderTemplateData, resourceTemplateData map[string]ResourceTemplateData, dataSourceTemplateData map[string]DataSourceTemplateData) (map[string][]byte, error) {
	outputData := make(map[string][]byte, len(providerTemplateData))

	sortedResources := mapToSortedSlice(resourceTemplateData)
	sortedDataSources := mapToSortedSlice(dataSourceTemplateData)

	// TODO: swap to single provider processing? (everywhere else is a map currently)
	for name, providerData := range providerTemplateData {
		providerData.Resources = sortedResources
		providerData.DataSources = sortedDataSources

		// Process provider template
		templateBytes, err := fs.ReadFile(t.templateDir, "provider.gotmpl")
		if err != nil {
			t.logger.Debug(fmt.Sprintf("no provider template found for %q, skipping", name))
			continue
		}

		providerBytes, err := t.processTemplate(templateBytes, providerData)
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error processing %q provider template: %s, skipping", name, err))
			continue
		}
		outputData["provider_gen.go"] = providerBytes

		// Process provider test template
		testTemplateBytes, err := fs.ReadFile(t.templateDir, "provider_test.gotmpl")
		if err != nil {
			t.logger.Debug(fmt.Sprintf("no provider test template found for %q, skipping", name))
			continue
		}

		testBytes, err := t.processTemplate(testTemplateBytes, providerData)
		if err != nil {
			t.logger.Warn(fmt.Sprintf("error processing %q provider test template: %s, skipping", name, err))
			continue
		}
		outputData["provider_gen_test.go"] = testBytes
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
