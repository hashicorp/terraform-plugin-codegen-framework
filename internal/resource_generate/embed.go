package resource_generate

import _ "embed"

//go:embed templates/resource_schema.gotmpl
var resourceSchemaGoTemplate string

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

//go:embed templates/common_block.gotmpl
var commonBlockGoTemplate string
