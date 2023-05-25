package datasource_generate

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type GeneratorSchema interface {
	ToBytes() (map[string][]byte, error)
}

type GeneratorDataSourceSchemas struct {
	schemas map[string]GeneratorDataSourceSchema
	// TODO: Could add a field to hold custom templates that are used in calls toBytes() to allow overriding.
	// getAttributes() and getBlocks() funcs.
}

type GeneratorDataSourceSchema struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
}

func NewGeneratorDataSourceSchemas(schemas map[string]GeneratorDataSourceSchema) GeneratorDataSourceSchemas {
	return GeneratorDataSourceSchemas{
		schemas: schemas,
	}
}

func (g GeneratorDataSourceSchemas) ToBytes() (map[string][]byte, error) {
	schemasBytes := make(map[string][]byte, len(g.schemas))

	for k, s := range g.schemas {
		b, err := g.toBytes(k, s)

		if err != nil {
			return nil, err
		}

		schemasBytes[k] = b
	}

	return schemasBytes, nil
}

func (g GeneratorDataSourceSchemas) toBytes(name string, a GeneratorDataSourceSchema) ([]byte, error) {
	funcMap := template.FuncMap{
		"getImports":    getImports,
		"getAttributes": getAttributes,
		"getBlocks":     getBlocks,
	}

	t, err := template.New("datasource_schema").Funcs(funcMap).Parse(
		dataSourceSchemaGoTemplate,
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	attrib := map[string]GeneratorDataSourceSchema{
		name: a,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getImports(schema GeneratorDataSourceSchema) (string, error) {
	var s strings.Builder

	var imports = make(map[string]struct{})

	for _, v := range schema.Attributes {
		for k := range v.Imports() {
			imports[k] = struct{}{}
		}
	}

	for _, v := range schema.Blocks {
		for k := range v.Imports() {
			imports[k] = struct{}{}
		}
	}

	for a := range imports {
		s.WriteString(fmt.Sprintf("\"%s\"\n", a))
	}

	return s.String(), nil
}

func getAttributes(attributes map[string]GeneratorAttribute) (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	keys := make([]string, len(attributes))

	for k := range attributes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if attributes[k] == nil {
			continue
		}

		str, err := attributes[k].ToString(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

func getBlocks(blocks map[string]GeneratorBlock) (string, error) {
	var s strings.Builder

	// Using sorted keys to guarantee attribute order as maps are unordered in Go.
	keys := make([]string, len(blocks))

	for k := range blocks {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if blocks[k] == nil {
			continue
		}

		str, err := blocks[k].ToString(k)

		if err != nil {
			return "", err
		}

		s.WriteString(str)
	}

	return s.String(), nil
}

type GeneratorImport interface {
	Imports() map[string]struct{}
}

type GeneratorAttribute interface {
	Equal(GeneratorAttribute) bool
	ToString(string) (string, error)
	GeneratorImport
}

type GeneratorBlock interface {
	Equal(GeneratorBlock) bool
	ToString(string) (string, error)
	GeneratorImport
}

type GeneratorNestedAttributeObject struct {
	Attributes map[string]GeneratorAttribute
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

type GeneratorNestedBlockObject struct {
	Attributes map[string]GeneratorAttribute
	Blocks     map[string]GeneratorBlock
	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

func customTypeEqual(x, y *specschema.CustomType) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if x.Import == nil && y.Import != nil {
		return false
	}

	if x.Import != nil && y.Import == nil {
		return false
	}

	if x.Import != nil && y.Import != nil {
		if *x.Import != *y.Import {
			return false
		}
	}

	if x.Type != y.Type {
		return false
	}

	if x.ValueType != y.ValueType {
		return false
	}

	return true
}
