package scaffold

import (
	_ "embed"
)

//go:embed templates/resource_scaffold.gotmpl
var resourceScaffoldGoTemplate string
