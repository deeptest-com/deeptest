package gojaPlugin

import (
	_ "embed"
	"github.com/dop251/goja"
)

const mochaFilename = "mocha.js"

//go:embed mocha.js
var mochaSource string

type MochaRootModule struct {
	program *goja.Program
}

func NewMocha() Module {
	prog := mustCompile(mochaSource, mochaFilename, true)
	return &MochaRootModule{
		program: prog,
	}
}

func (root *MochaRootModule) NewModuleInstance(vu VU) Instance { // nolint:varnamelen
	exports := mustRequire(root.program, vu.Runtime())

	return &CommModule{exports: *exports}
}

var (
	_ Module = (*MochaRootModule)(nil)
)
