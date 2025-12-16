// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scaffold

import (
	"bytes"
	"os"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

// DataSourceBytes will create scaffolding Go code bytes for a Terraform Plugin Framework data source
func DataSourceBytes(dataSourceIdentifier schema.FrameworkIdentifier, packageName, templatePath string) ([]byte, error) {
	templateStr := dataSourceScaffoldGoTemplate

	if templatePath != "" {
		file, err := os.ReadFile(templatePath)
		if err != nil {
			return nil, err
		}

		templateStr = string(file)
	}

	t, err := template.New("data_source_scaffold").Parse(templateStr)
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
		NameCamel:   dataSourceIdentifier.ToCamelCase(),
		NamePascal:  dataSourceIdentifier.ToPascalCase(),
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
