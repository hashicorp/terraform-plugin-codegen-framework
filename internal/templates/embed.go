// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	_ "embed"
)

//go:embed object_type_equal.gotmpl
var ObjectTypeEqualTemplate string

//go:embed object_type_string.gotmpl
var ObjectTypeStringTemplate string

//go:embed object_type_typable.gotmpl
var ObjectTypeTypableTemplate string

//go:embed object_type_type.gotmpl
var ObjectTypeTypeTemplate string

//go:embed object_type_value.gotmpl
var ObjectTypeValueTemplate string

//go:embed object_type_value_from_object.gotmpl
var ObjectTypeValueFromObjectTemplate string

//go:embed object_type_value_from_terraform.gotmpl
var ObjectTypeValueFromTerraformTemplate string

//go:embed object_type_value_must.gotmpl
var ObjectTypeValueMustTemplate string

//go:embed object_type_value_null.gotmpl
var ObjectTypeValueNullTemplate string

//go:embed object_type_value_type.gotmpl
var ObjectTypeValueTypeTemplate string

//go:embed object_type_value_unknown.gotmpl
var ObjectTypeValueUnknownTemplate string

//go:embed model_object_helpers.gotmpl
var ModelObjectHelpersTemplate string

//go:embed model_object_helpers_edited.gotmpl
var ModelObjectHelpersTemplateEdited string

//go:embed schema.gotmpl
var SchemaGoTemplate string

//go:embed to_from.gotmpl
var ToFromTemplate string
