package scaffold

import (
	"bytes"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/util"
)

// ResourceBytes will create scaffolding Go code bytes for a Terraform Plugin Framework resource
func ResourceBytes(resourceIdentifier util.FrameworkIdentifer, packageName string) ([]byte, error) {
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
		NameCamel:   resourceIdentifier.ToPascalCase(),
		NamePascal:  resourceIdentifier.ToCamelCase(),
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
