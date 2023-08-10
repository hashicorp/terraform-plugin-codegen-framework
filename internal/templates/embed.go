// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	_ "embed"
)

//go:embed model_object_helpers.gotmpl
var ModelObjectHelpersTemplate string

//go:embed schema.gotmpl
var SchemaGoTemplate string

//go:embed single_nested_object_to_from.gotmpl
var SingleNestedObjectToFromTemplate string
