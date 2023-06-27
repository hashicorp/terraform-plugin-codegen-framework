// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gen

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
)

type DataSourcesModelsGenerator struct {
	Templates []string
}

func NewDataSourcesModelsGenerator() DataSourcesModelsGenerator {
	return DataSourcesModelsGenerator{
		Templates: []string{
			"internal/templates/model/datasource_model.gotmpl",
			"internal/templates/model/attributes.gotmpl",
			"internal/templates/model/bool_attribute.gotmpl",
			"internal/templates/model/list_attribute.gotmpl",
			"internal/templates/model/single_nested_attribute.gotmpl",
			"internal/templates/model/single_nested_model.gotmpl",
		},
	}
}

func (d DataSourcesModelsGenerator) Process(ir spec.Specification) (map[string][]byte, error) {
	funcMap := template.FuncMap{
		"snakeCaseToCamelCase": snakeCaseToCamelCase,
	}

	datasourceModelTemplate, err := template.New("datasource_model.gotmpl").Funcs(funcMap).ParseFiles(d.Templates...)
	if err != nil {
		return nil, err
	}

	dataSourcesModels := make(map[string][]byte, len(ir.DataSources))

	for _, d := range ir.DataSources {
		var buf bytes.Buffer

		err = datasourceModelTemplate.Execute(&buf, d)
		if err != nil {
			return nil, err
		}

		dataSourcesModels[d.Name] = buf.Bytes()
	}

	return dataSourcesModels, nil
}

// snakeCaseToCamelCase relies on the convention of using snake-case
// names in configuration.
// TODO: A more robust approach is likely required here.
func snakeCaseToCamelCase(input string) string {
	inputSplit := strings.Split(input, "_")

	var ucName string

	for _, v := range inputSplit {
		if len(v) < 1 {
			continue
		}

		firstChar := v[0:1]
		ucFirstChar := strings.ToUpper(firstChar)

		if len(v) < 2 {
			ucName += ucFirstChar
			continue
		}

		ucName += ucFirstChar + v[1:]
	}

	return ucName
}

type DataSourcesHelpersGenerator struct {
	Templates []string
}

func NewDataSourcesHelpersGenerator() DataSourcesHelpersGenerator {
	return DataSourcesHelpersGenerator{
		Templates: []string{
			"internal/templates/helper/datasource_helper.gotmpl",
			"internal/templates/helper/attributes.gotmpl",
			"internal/templates/helper/bool_attribute.gotmpl",
			"internal/templates/helper/list_attribute.gotmpl",
			"internal/templates/helper/elem_type.gotmpl",
			"internal/templates/helper/single_nested_attribute.gotmpl",
			"internal/templates/helper/single_nested_helper.gotmpl",
		},
	}
}

func (d DataSourcesHelpersGenerator) Process(ir spec.Specification) (map[string][]byte, error) {
	funcMap := template.FuncMap{
		"snakeCaseToCamelCase":                snakeCaseToCamelCase,
		"hasNestedWithAssociatedExternalType": hasNestedWithAssociatedExternalType,
		"ucFirst":                             ucFirst,
	}

	datasourceHelperTemplate, err := template.New("datasource_helper.gotmpl").Funcs(funcMap).ParseFiles(d.Templates...)
	if err != nil {
		return nil, err
	}

	dataSourcesHelpers := make(map[string][]byte, len(ir.DataSources))

	for _, d := range ir.DataSources {
		var buf bytes.Buffer

		err = datasourceHelperTemplate.Execute(&buf, d)
		if err != nil {
			return nil, err
		}

		dataSourcesHelpers[d.Name] = buf.Bytes()
	}

	return dataSourcesHelpers, nil
}

// hasNestedWithAssociatedExternalType ranges over the "top-level" data source attributes
// to determine whether any of them are SingleNested, and if so, a check is then made
// to see if AssociatedExternalType is set on SingleNested. If this returns true then
// it indicates that additional helper functions should be generated for converting to/from
// external types < = > models.
func hasNestedWithAssociatedExternalType(attributes []datasource.Attribute) bool {
	for range attributes {

	}

	return false
}

func ucFirst(input string) string {
	if len(input) == 0 {
		return input
	}

	firstChar := input[0:1]

	ucFirstChar := strings.ToUpper(firstChar)

	if len(input) < 2 {
		return ucFirstChar
	}

	return ucFirstChar + input[1:]
}
