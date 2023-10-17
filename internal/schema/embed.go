// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	_ "embed"
)

// Bool From/To

//go:embed templates/bool_from.gotmpl
var BoolFromTemplate string

//go:embed templates/bool_to.gotmpl
var BoolToTemplate string

// Bool Type

//go:embed templates/bool_type_equal.gotmpl
var BoolTypeEqualTemplate string

//go:embed templates/bool_type_string.gotmpl
var BoolTypeStringTemplate string

//go:embed templates/bool_type_type.gotmpl
var BoolTypeTypeTemplate string

//go:embed templates/bool_type_typable.gotmpl
var BoolTypeTypableTemplate string

//go:embed templates/bool_type_value_from_bool.gotmpl
var BoolTypeValueFromBoolTemplate string

//go:embed templates/bool_type_value_from_terraform.gotmpl
var BoolTypeValueFromTerraformTemplate string

//go:embed templates/bool_type_value_type.gotmpl
var BoolTypeValueTypeTemplate string

// Bool Value

//go:embed templates/bool_value_equal.gotmpl
var BoolValueEqualTemplate string

//go:embed templates/bool_value_type.gotmpl
var BoolValueTypeTemplate string

//go:embed templates/bool_value_value.gotmpl
var BoolValueValueTemplate string

//go:embed templates/bool_value_valuable.gotmpl
var BoolValueValuableTemplate string

// Object

//go:embed templates/object_from.gotmpl
var ObjectFromTemplate string

//go:embed templates/object_to.gotmpl
var ObjectToTemplate string

// Object Type

//go:embed templates/object_type_equal.gotmpl
var ObjectTypeEqualTemplate string

//go:embed templates/object_type_string.gotmpl
var ObjectTypeStringTemplate string

//go:embed templates/object_type_typable.gotmpl
var ObjectTypeTypableTemplate string

//go:embed templates/object_type_type.gotmpl
var ObjectTypeTypeTemplate string

//go:embed templates/object_type_value.gotmpl
var ObjectTypeValueTemplate string

//go:embed templates/object_type_value_from_object.gotmpl
var ObjectTypeValueFromObjectTemplate string

//go:embed templates/object_type_value_from_terraform.gotmpl
var ObjectTypeValueFromTerraformTemplate string

//go:embed templates/object_type_value_must.gotmpl
var ObjectTypeValueMustTemplate string

//go:embed templates/object_type_value_null.gotmpl
var ObjectTypeValueNullTemplate string

//go:embed templates/object_type_value_type.gotmpl
var ObjectTypeValueTypeTemplate string

//go:embed templates/object_type_value_unknown.gotmpl
var ObjectTypeValueUnknownTemplate string

// Object Value

//go:embed templates/object_value_attribute_types.gotmpl
var ObjectValueAttributeTypesTemplate string

//go:embed templates/object_value_equal.gotmpl
var ObjectValueEqualTemplate string

//go:embed templates/object_value_is_null.gotmpl
var ObjectValueIsNullTemplate string

//go:embed templates/object_value_is_unknown.gotmpl
var ObjectValueIsUnknownTemplate string

//go:embed templates/object_value_string.gotmpl
var ObjectValueStringTemplate string

//go:embed templates/object_value_to_object_value.gotmpl
var ObjectValueToObjectValueTemplate string

//go:embed templates/object_value_to_terraform_value.gotmpl
var ObjectValueToTerraformValueTemplate string

//go:embed templates/object_value_type.gotmpl
var ObjectValueTypeTemplate string

//go:embed templates/object_value_valuable.gotmpl
var ObjectValueValuableTemplate string

//go:embed templates/object_value_value.gotmpl
var ObjectValueValueTemplate string

//go:embed templates/schema.gotmpl
var SchemaGoTemplate string
