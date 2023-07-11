package scaffold

import (
	_ "embed"
)

//go:embed templates/resource_scaffold.gotmpl
var resourceScaffoldGoTemplate string

//go:embed templates/data_source_scaffold.gotmpl
var dataSourceScaffoldGoTemplate string
