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

// Float64 From/To

//go:embed templates/float64_from.gotmpl
var Float64FromTemplate string

//go:embed templates/float64_to.gotmpl
var Float64ToTemplate string

// Float64 Type

//go:embed templates/float64_type_equal.gotmpl
var Float64TypeEqualTemplate string

//go:embed templates/float64_type_string.gotmpl
var Float64TypeStringTemplate string

//go:embed templates/float64_type_type.gotmpl
var Float64TypeTypeTemplate string

//go:embed templates/float64_type_typable.gotmpl
var Float64TypeTypableTemplate string

//go:embed templates/float64_type_value_from_float64.gotmpl
var Float64TypeValueFromFloat64Template string

//go:embed templates/float64_type_value_from_terraform.gotmpl
var Float64TypeValueFromTerraformTemplate string

//go:embed templates/float64_type_value_type.gotmpl
var Float64TypeValueTypeTemplate string

// Float64 Value

//go:embed templates/float64_value_equal.gotmpl
var Float64ValueEqualTemplate string

//go:embed templates/float64_value_type.gotmpl
var Float64ValueTypeTemplate string

//go:embed templates/float64_value_value.gotmpl
var Float64ValueValueTemplate string

//go:embed templates/float64_value_valuable.gotmpl
var Float64ValueValuableTemplate string

// Int64 From/To

//go:embed templates/int64_from.gotmpl
var Int64FromTemplate string

//go:embed templates/int64_to.gotmpl
var Int64ToTemplate string

// Int64 Type

//go:embed templates/int64_type_equal.gotmpl
var Int64TypeEqualTemplate string

//go:embed templates/int64_type_string.gotmpl
var Int64TypeStringTemplate string

//go:embed templates/int64_type_type.gotmpl
var Int64TypeTypeTemplate string

//go:embed templates/int64_type_typable.gotmpl
var Int64TypeTypableTemplate string

//go:embed templates/int64_type_value_from_int64.gotmpl
var Int64TypeValueFromInt64Template string

//go:embed templates/int64_type_value_from_terraform.gotmpl
var Int64TypeValueFromTerraformTemplate string

//go:embed templates/int64_type_value_type.gotmpl
var Int64TypeValueTypeTemplate string

// Int64 Value

//go:embed templates/int64_value_equal.gotmpl
var Int64ValueEqualTemplate string

//go:embed templates/int64_value_type.gotmpl
var Int64ValueTypeTemplate string

//go:embed templates/int64_value_value.gotmpl
var Int64ValueValueTemplate string

//go:embed templates/int64_value_valuable.gotmpl
var Int64ValueValuableTemplate string

// Number From/To

//go:embed templates/number_from.gotmpl
var NumberFromTemplate string

//go:embed templates/number_to.gotmpl
var NumberToTemplate string

// Number Type

//go:embed templates/number_type_equal.gotmpl
var NumberTypeEqualTemplate string

//go:embed templates/number_type_string.gotmpl
var NumberTypeStringTemplate string

//go:embed templates/number_type_type.gotmpl
var NumberTypeTypeTemplate string

//go:embed templates/number_type_typable.gotmpl
var NumberTypeTypableTemplate string

//go:embed templates/number_type_value_from_number.gotmpl
var NumberTypeValueFromNumberTemplate string

//go:embed templates/number_type_value_from_terraform.gotmpl
var NumberTypeValueFromTerraformTemplate string

//go:embed templates/number_type_value_type.gotmpl
var NumberTypeValueTypeTemplate string

// Number Value

//go:embed templates/number_value_equal.gotmpl
var NumberValueEqualTemplate string

//go:embed templates/number_value_type.gotmpl
var NumberValueTypeTemplate string

//go:embed templates/number_value_value.gotmpl
var NumberValueValueTemplate string

//go:embed templates/number_value_valuable.gotmpl
var NumberValueValuableTemplate string

// Object From/To

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

// String From/To

//go:embed templates/string_from.gotmpl
var StringFromTemplate string

//go:embed templates/string_to.gotmpl
var StringToTemplate string

// String Type

//go:embed templates/string_type_equal.gotmpl
var StringTypeEqualTemplate string

//go:embed templates/string_type_string.gotmpl
var StringTypeStringTemplate string

//go:embed templates/string_type_type.gotmpl
var StringTypeTypeTemplate string

//go:embed templates/string_type_typable.gotmpl
var StringTypeTypableTemplate string

//go:embed templates/string_type_value_from_string.gotmpl
var StringTypeValueFromStringTemplate string

//go:embed templates/string_type_value_from_terraform.gotmpl
var StringTypeValueFromTerraformTemplate string

//go:embed templates/string_type_value_type.gotmpl
var StringTypeValueTypeTemplate string

// String Value

//go:embed templates/string_value_equal.gotmpl
var StringValueEqualTemplate string

//go:embed templates/string_value_type.gotmpl
var StringValueTypeTemplate string

//go:embed templates/string_value_value.gotmpl
var StringValueValueTemplate string

//go:embed templates/string_value_valuable.gotmpl
var StringValueValuableTemplate string
