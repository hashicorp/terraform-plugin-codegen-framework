// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	_ "embed"
	"strconv"
	"text/template"
)

//go:embed templates/single_nested_block.gotmpl
var singleNestedBlockGoTemplate string

//go:embed templates/common_block.gotmpl
var commonBlockGoTemplate string

func addCommonBlockTemplate(t *template.Template) (*template.Template, error) {
	commonTemplateFuncs := template.FuncMap{
		"quote": strconv.Quote,
	}

	return t.New("common_block").Funcs(commonTemplateFuncs).Parse(commonBlockGoTemplate)
}
