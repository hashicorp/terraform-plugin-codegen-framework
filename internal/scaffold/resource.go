// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scaffold

import (
	"bytes"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

// ResourceBytes will create scaffolding Go code bytes for a Terraform Plugin Framework resource
func ResourceBytes(resourceIdentifier schema.FrameworkIdentifier, packageName string) ([]byte, error) {
	t, err := template.New("resource_scaffold").Parse(resourceScaffoldGoTemplate)
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
		NameSnake:   string(resourceIdentifier),
		NameCamel:   resourceIdentifier.ToCamelCase(),
		NamePascal:  resourceIdentifier.ToPascalCase(),
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
