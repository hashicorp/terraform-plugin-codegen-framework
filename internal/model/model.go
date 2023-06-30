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

type Field struct {
	Name      string
	TfsdkName string
	ValueType string
}

func (f Field) String() string {
	return fmt.Sprintf("%s %s `tfsdk:%q`", f.Name, f.ValueType, f.TfsdkName)
}

type Model struct {
	Name   string
	Fields []Field
}

func (m Model) String() string {
	var fieldsStr string

	for _, field := range m.Fields {
		fieldsStr += field.String() + "\n"
	}

	fieldsStrTrim := strings.TrimSuffix(fieldsStr, "\n")

	return fmt.Sprintf("type %sModel struct {\n%s\n}", m.Name, fieldsStrTrim)
}

type ObjectHelper struct {
	Name      string
	AttrTypes map[string]string
}

func (o ObjectHelper) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(objectType, o.Name))

	var attrTypes []string

	for k, v := range o.AttrTypes {
		attrTypes = append(attrTypes, fmt.Sprintf("%q: %s,", k, v))
	}

	objAttributeTypes := fmt.Sprintf(objectAttributeTypes, o.Name, strings.Join(attrTypes, "\n"))

	sb.WriteString(fmt.Sprintf("\n\n%s", objAttributeTypes))

	return sb.String()
}

var objectType = `func (m %sModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}`

var objectAttributeTypes = `func (m %sModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
%s
}
}`

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
