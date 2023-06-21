// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_generate

import (
	_ "embed"
	"strconv"
	"text/template"
)

//go:embed templates/schema.gotmpl
var schemaGoTemplate string

//go:embed templates/bool_attribute.gotmpl
var boolAttributeGoTemplate string

//go:embed templates/float64_attribute.gotmpl
var float64AttributeGoTemplate string

//go:embed templates/int64_attribute.gotmpl
var int64AttributeGoTemplate string

//go:embed templates/list_attribute.gotmpl
var listAttributeGoTemplate string

//go:embed templates/list_nested_attribute.gotmpl
var listNestedAttributeGoTemplate string

//go:embed templates/map_attribute.gotmpl
var mapAttributeGoTemplate string

//go:embed templates/map_nested_attribute.gotmpl
var mapNestedAttributeGoTemplate string

//go:embed templates/number_attribute.gotmpl
var numberAttributeGoTemplate string

//go:embed templates/object_attribute.gotmpl
var objectAttributeGoTemplate string

//go:embed templates/set_attribute.gotmpl
var setAttributeGoTemplate string

//go:embed templates/set_nested_attribute.gotmpl
var setNestedAttributeGoTemplate string

//go:embed templates/single_nested_attribute.gotmpl
var singleNestedAttributeGoTemplate string

//go:embed templates/string_attribute.gotmpl
var stringAttributeGoTemplate string

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

// Models

//go:embed model_templates/bool_model.gotmpl
var boolModel string

//go:embed model_templates/datasource_model.gotmpl
var datasourceModel string

//go:embed model_templates/list_attribute.gotmpl
var listAttributeModel string

//go:embed model_templates/single_nested_attribute.gotmpl
var singleNestedAttributeModel string

//go:embed model_templates/single_nested_model.gotmpl
var singleNestedModel string
