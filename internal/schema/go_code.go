package schema

import "go/format"

const (
	// ExportSchemaFunc refers to the Go identifier of a generated function that accepts a context.Context and returns a *schema.Schema struct
	// for a resource, data source, or provider schema.
	ExportSchemaFunc NotableExport = "SchemaFunc"

	// ExportSchemaModelType refers to the Go identifier of a generated struct that represents the schema model type
	// for a resource, data source, or provider schema.
	ExportSchemaModelType NotableExport = "SchemaModelType"
)

type NotableExport string

// GoCode is a struct representing unformatted Go code.
type GoCode struct {
	// PackageName is the name of the Go package that can be used to access exports.
	PackageName string

	// NotableExports is a map where the key refers to the contents of the export, and the value is a Go identifier.
	NotableExports map[NotableExport]string

	// Bytes is the unformatted Go code data.
	Bytes []byte
}

func (g GoCode) Format() ([]byte, error) {
	return format.Source(g.Bytes)
}
