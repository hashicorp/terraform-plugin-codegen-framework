package model

import (
	"fmt"
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
