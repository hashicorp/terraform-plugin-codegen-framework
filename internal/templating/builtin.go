package templating

import (
	"embed"
	_ "embed"
	"text/template"
)

//go:embed _builtin/*.gotmpl
var builtinEmbeds embed.FS

func addBuiltInTemplates(base *template.Template) (*template.Template, error) {
	tmpl, err := base.ParseFS(builtinEmbeds, "_builtin/*.gotmpl")
	if err != nil {
		return tmpl, err
	}

	return tmpl, nil
}
