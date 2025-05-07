package workflow_test

import "go.k6.io/k6/js/modules"

const importPath = "k6/x/workflow_test"

func init() {
	modules.Register(importPath, new(rootModule))
}
