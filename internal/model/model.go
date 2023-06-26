// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package model

import (
	"fmt"
	"strings"
)

const (
	BoolValueType    = "types.Bool"
	Float64ValueType = "types.Float64"
	Int64ValueType   = "types.Int64"
	ListValueType    = "types.List"
	MapValueType     = "types.Map"
	NumberValueType  = "types.Number"
	ObjectValueType  = "types.Object"
	SetValueType     = "types.Set"
	StringValueType  = "types.String"
)

type StructField struct {
	Name      string
	TfsdkName string
	ValueType string
}

func (f StructField) String() string {
	return fmt.Sprintf("%s %s `tfsdk:%q`", f.Name, f.ValueType, f.TfsdkName)
}

type Model struct {
	Name   string
	Fields string
}

func (m Model) String() string {
	return fmt.Sprintf("type %sModel struct {\n%s\n}", m.Name, m.Fields)
}

// SnakeCaseToCamelCase relies on the convention of using snake-case
// names in configuration.
// TODO: A more robust approach is likely required here.
func SnakeCaseToCamelCase(input string) string {
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
