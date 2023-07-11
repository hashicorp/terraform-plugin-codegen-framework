package scaffold

import (
	"bytes"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/util"
)

// DataSourceBytes will create scaffolding Go code bytes for a Terraform Plugin Framework data source
func DataSourceBytes(dataSourceIdentifier util.FrameworkIdentifer, packageName string) ([]byte, error) {
	t, err := template.New("data_source_scaffold").Parse(dataSourceScaffoldGoTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		PackageName string
		NameSnake   string
		NameCamel   string
		NamePascal  string
	}{
		PackageName: packageName,
		NameSnake:   string(dataSourceIdentifier),
		NameCamel:   dataSourceIdentifier.ToPascalCase(),
		NamePascal:  dataSourceIdentifier.ToCamelCase(),
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
