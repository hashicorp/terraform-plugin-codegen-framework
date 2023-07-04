package datasource_generate

import (
	"fmt"
	"strings"

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
