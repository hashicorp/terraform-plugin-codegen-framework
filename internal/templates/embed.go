// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	_ "embed"
)

//go:embed model_object_helpers.gotmpl
var ModelObjectHelpersTemplate string

//go:embed model_object_to_from.gotmpl
var ModelObjectToFromTemplate string

//go:embed primitive_to_from.gotmpl
var PrimitiveToFromTemplate string

//go:embed schema.gotmpl
var SchemaGoTemplate string
