package ncloud

import (
	_ "embed"
)

//go:embed templates/main.go.tpl
var MainTemplate string
