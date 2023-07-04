package datasource_generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

type GeneratorDataSourceSchema struct {
	Attributes GeneratorAttributes
	Blocks     GeneratorBlocks
}

func (g GeneratorDataSourceSchema) ImportsString() (string, error) {
	imports := schema.NewImports()

	for _, v := range g.Attributes {
		imports.Add(v.Imports().All()...)
	}

	for _, v := range g.Blocks {
		imports.Add(v.Imports().All()...)
	}

	var sb strings.Builder

	for _, i := range imports.All() {
		var alias string

		if i.Alias != nil {
			alias = *i.Alias + " "
		}

		sb.WriteString(fmt.Sprintf("%s%q\n", alias, i.Path))
	}

	return sb.String(), nil
}

func (g GeneratorDataSourceSchema) SchemaBytes(name, packageName string) ([]byte, error) {
	funcMap := template.FuncMap{
		"ImportsString":    g.ImportsString,
		"AttributesString": g.Attributes.String,
		"BlocksString":     g.Blocks.String,
	}

	t, err := template.New("schema").Funcs(funcMap).Parse(
		schemaGoTemplate,
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		Name string
		GeneratorDataSourceSchema
		PackageName string
	}{
		Name:                      name,
		GeneratorDataSourceSchema: g,
		PackageName:               packageName,
	}

	err = t.Execute(&buf, templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
