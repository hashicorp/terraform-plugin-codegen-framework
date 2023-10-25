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

// List From/To

//go:embed templates/list_from.gotmpl
var ListFromTemplate string

//go:embed templates/list_to.gotmpl
var ListToTemplate string

// List Type

//go:embed templates/list_type_equal.gotmpl
var ListTypeEqualTemplate string

//go:embed templates/list_type_string.gotmpl
var ListTypeStringTemplate string

//go:embed templates/list_type_type.gotmpl
var ListTypeTypeTemplate string

//go:embed templates/list_type_typable.gotmpl
var ListTypeTypableTemplate string

//go:embed templates/list_type_value_from_list.gotmpl
var ListTypeValueFromListTemplate string

//go:embed templates/list_type_value_from_terraform.gotmpl
var ListTypeValueFromTerraformTemplate string

//go:embed templates/list_type_value_type.gotmpl
var ListTypeValueTypeTemplate string

// List Value

//go:embed templates/list_value_equal.gotmpl
var ListValueEqualTemplate string

//go:embed templates/list_value_type.gotmpl
var ListValueTypeTemplate string

//go:embed templates/list_value_value.gotmpl
var ListValueValueTemplate string

//go:embed templates/list_value_valuable.gotmpl
var ListValueValuableTemplate string

// Map From/To

//go:embed templates/map_from.gotmpl
var MapFromTemplate string

//go:embed templates/map_to.gotmpl
var MapToTemplate string

// Map Type

//go:embed templates/map_type_equal.gotmpl
var MapTypeEqualTemplate string

//go:embed templates/map_type_string.gotmpl
var MapTypeStringTemplate string

//go:embed templates/map_type_type.gotmpl
var MapTypeTypeTemplate string

//go:embed templates/map_type_typable.gotmpl
var MapTypeTypableTemplate string

//go:embed templates/map_type_value_from_map.gotmpl
var MapTypeValueFromMapTemplate string

//go:embed templates/map_type_value_from_terraform.gotmpl
var MapTypeValueFromTerraformTemplate string

//go:embed templates/map_type_value_type.gotmpl
var MapTypeValueTypeTemplate string

// Map Value

//go:embed templates/map_value_equal.gotmpl
var MapValueEqualTemplate string

//go:embed templates/map_value_type.gotmpl
var MapValueTypeTemplate string

//go:embed templates/map_value_value.gotmpl
var MapValueValueTemplate string

//go:embed templates/map_value_valuable.gotmpl
var MapValueValuableTemplate string

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

// NestedObject From/To

//go:embed templates/nested_object_from.gotmpl
var NestedObjectFromTemplate string

//go:embed templates/nested_object_to.gotmpl
var NestedObjectToTemplate string

// NestedObject Type

//go:embed templates/nested_object_type_equal.gotmpl
var NestedObjectTypeEqualTemplate string

//go:embed templates/nested_object_type_string.gotmpl
var NestedObjectTypeStringTemplate string

//go:embed templates/nested_object_type_typable.gotmpl
var NestedObjectTypeTypableTemplate string

//go:embed templates/nested_object_type_type.gotmpl
var NestedObjectTypeTypeTemplate string

//go:embed templates/nested_object_type_value.gotmpl
var NestedObjectTypeValueTemplate string

//go:embed templates/nested_object_type_value_from_object.gotmpl
var NestedObjectTypeValueFromObjectTemplate string

//go:embed templates/nested_object_type_value_from_terraform.gotmpl
var NestedObjectTypeValueFromTerraformTemplate string

//go:embed templates/nested_object_type_value_must.gotmpl
var NestedObjectTypeValueMustTemplate string

//go:embed templates/nested_object_type_value_null.gotmpl
var NestedObjectTypeValueNullTemplate string

//go:embed templates/nested_object_type_value_type.gotmpl
var NestedObjectTypeValueTypeTemplate string

//go:embed templates/nested_object_type_value_unknown.gotmpl
var NestedObjectTypeValueUnknownTemplate string

// NestedObject Value

//go:embed templates/nested_object_value_attribute_types.gotmpl
var NestedObjectValueAttributeTypesTemplate string

//go:embed templates/nested_object_value_equal.gotmpl
var NestedObjectValueEqualTemplate string

//go:embed templates/nested_object_value_is_null.gotmpl
var NestedObjectValueIsNullTemplate string

//go:embed templates/nested_object_value_is_unknown.gotmpl
var NestedObjectValueIsUnknownTemplate string

//go:embed templates/nested_object_value_string.gotmpl
var NestedObjectValueStringTemplate string

//go:embed templates/nested_object_value_to_object_value.gotmpl
var NestedObjectValueToObjectValueTemplate string

//go:embed templates/nested_object_value_to_terraform_value.gotmpl
var NestedObjectValueToTerraformValueTemplate string

//go:embed templates/nested_object_value_type.gotmpl
var NestedObjectValueTypeTemplate string

//go:embed templates/nested_object_value_valuable.gotmpl
var NestedObjectValueValuableTemplate string

//go:embed templates/nested_object_value_value.gotmpl
var NestedObjectValueValueTemplate string

//go:embed templates/schema.gotmpl
var SchemaGoTemplate string

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

//go:embed templates/object_type_type.gotmpl
var ObjectTypeTypeTemplate string

//go:embed templates/object_type_typable.gotmpl
var ObjectTypeTypableTemplate string

//go:embed templates/object_type_value_from_object.gotmpl
var ObjectTypeValueFromObjectTemplate string

//go:embed templates/object_type_value_from_terraform.gotmpl
var ObjectTypeValueFromTerraformTemplate string

//go:embed templates/object_type_value_type.gotmpl
var ObjectTypeValueTypeTemplate string

// Object Value

//go:embed templates/object_value_attribute_types.gotmpl
var ObjectValueAttributeTypesTemplate string

//go:embed templates/object_value_equal.gotmpl
var ObjectValueEqualTemplate string

//go:embed templates/object_value_type.gotmpl
var ObjectValueTypeTemplate string

//go:embed templates/object_value_value.gotmpl
var ObjectValueValueTemplate string

//go:embed templates/object_value_valuable.gotmpl
var ObjectValueValuableTemplate string

// Set From/To

//go:embed templates/set_from.gotmpl
var SetFromTemplate string

//go:embed templates/set_to.gotmpl
var SetToTemplate string

// Set Type

//go:embed templates/set_type_equal.gotmpl
var SetTypeEqualTemplate string

//go:embed templates/set_type_string.gotmpl
var SetTypeStringTemplate string

//go:embed templates/set_type_type.gotmpl
var SetTypeTypeTemplate string

//go:embed templates/set_type_typable.gotmpl
var SetTypeTypableTemplate string

//go:embed templates/set_type_value_from_set.gotmpl
var SetTypeValueFromSetTemplate string

//go:embed templates/set_type_value_from_terraform.gotmpl
var SetTypeValueFromTerraformTemplate string

//go:embed templates/set_type_value_type.gotmpl
var SetTypeValueTypeTemplate string

// Set Value

//go:embed templates/set_value_equal.gotmpl
var SetValueEqualTemplate string

//go:embed templates/set_value_type.gotmpl
var SetValueTypeTemplate string

//go:embed templates/set_value_value.gotmpl
var SetValueValueTemplate string

//go:embed templates/set_value_valuable.gotmpl
var SetValueValuableTemplate string

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
