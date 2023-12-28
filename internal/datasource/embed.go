// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	_ "embed"
	"strconv"
	"text/template"
)

//go:embed templates/map_attribute.gotmpl
var mapAttributeTemplate string

//go:embed templates/map_nested_attribute.gotmpl
var mapNestedAttributeGoTemplate string

//go:embed templates/set_attribute.gotmpl
var setAttributeTemplate string

//go:embed templates/set_nested_attribute.gotmpl
var setNestedAttributeGoTemplate string

//go:embed templates/list_nested_block.gotmpl
var listNestedBlockGoTemplate string

//go:embed templates/set_nested_block.gotmpl
var setNestedBlockGoTemplate string

//go:embed templates/single_nested_block.gotmpl
var singleNestedBlockGoTemplate string

//go:embed templates/common_attribute.gotmpl
var commonAttributeGoTemplate string

func addCommonAttributeTemplate(t *template.Template) (*template.Template, error) {
	commonTemplateFuncs := template.FuncMap{
		"quote": strconv.Quote,
	}

	return t.New("common_attribute").Funcs(commonTemplateFuncs).Parse(commonAttributeGoTemplate)
}

//go:embed templates/common_block.gotmpl
var commonBlockGoTemplate string

func addCommonBlockTemplate(t *template.Template) (*template.Template, error) {
	commonTemplateFuncs := template.FuncMap{
		"quote": strconv.Quote,
	}

	return t.New("common_block").Funcs(commonTemplateFuncs).Parse(commonBlockGoTemplate)
}

//go:embed templates/attribute.gotmpl
var attributeTemplate string

func addAttributeTemplate(t *template.Template) (*template.Template, error) {
	templateFuncs := template.FuncMap{
		"quote": strconv.Quote,
	}

	return t.New("attribute").Funcs(templateFuncs).Parse(attributeTemplate)
}
