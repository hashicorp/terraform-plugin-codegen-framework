// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	_ "embed"
)

//go:embed schema.gotmpl
var SchemaGoTemplate string

//go:embed to_from.gotmpl
var ToFromTemplate string
