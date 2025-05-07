// Package workflow_test contains the xk6-workflow-test extension.
package workflow_test

import "go.k6.io/k6/js/modules"

type rootModule struct{}

func (*rootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &module{vu}
}

type module struct {
	vu modules.VU
}

func (m *module) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]any{
			"greeting":  m.greeting,
			"b32encode": m.b32encode,
			"b32decode": m.b32decode,
			"Random":    m.random,
		},
	}
}

var _ modules.Module = (*rootModule)(nil)
