package gojaPlugin

import (
	_ "embed"
	"github.com/dop251/goja"
)

const chaiFilename = "chai.js"

//go:embed chai.js
var chaiSource string

type ChaiRootModule struct {
	program *goja.Program
}

func NewChai() Module {
	prog := mustCompile(chaiSource, chaiFilename, false)
	return &ChaiRootModule{
		program: prog,
	}
}

func (root *ChaiRootModule) NewModuleInstance(vu VU) Instance { // nolint:varnamelen
	exports := mustRequire(root.program, vu.Runtime())

	return &CommModule{exports: *exports}
}

var (
	_ Module = (*ChaiRootModule)(nil)
)
