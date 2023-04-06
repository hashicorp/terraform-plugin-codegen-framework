package generate

import (
	"bytes"
	"text/template"

	"github/hashicorp/terraform-provider-code-generator/internal/transform"
)

// DataSourcesSchema
// TODO: Setup different functions that can be called from the different commands.
// TODO: Handle writing of models.
// TODO: Consider adding writing of imports either with schema or as separate write.
// TODO: Handle processing of Provider and Resources schema
func DataSourcesSchema(ir transform.IntermediateRepresentation, output string) (map[string][]byte, error) {
	datasourceSchemaTemplate, err := template.ParseFiles(
		"internal/templates/datasource_schema.gotmpl",
		"internal/templates/attributes.gotmpl",
		"internal/templates/bool_attribute.gotmpl",
		"internal/templates/list_attribute.gotmpl",
		"internal/templates/validator.gotmpl",
		"internal/templates/element_type.gotmpl",
	)
	if err != nil {
		return nil, err
	}

	dataSourcesSchema := make(map[string][]byte, len(ir.DataSources))

	for _, d := range ir.DataSources {
		var buf bytes.Buffer

		err = datasourceSchemaTemplate.Execute(&buf, d)
		if err != nil {
			return nil, err
		}

		dataSourcesSchema[d.Name] = buf.Bytes()
	}

	return dataSourcesSchema, nil
}
