package gojaPlugin

import (
	_ "embed"
	"errors"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
)

type CommModule struct {
	exports Exports
}

func (mod *CommModule) Exports() Exports {
	return mod.exports
}

var (
	_ Instance = (*CommModule)(nil)
)

func mustCompile(source, name string, isESM bool) *goja.Program {
	prog, err := compile(source, name, isESM)
	if err != nil {
		panic(err)
	}

	return prog
}

func compile(source, name string, isESM bool) (prog *goja.Program, err error) {
	//opts := parser.WithDisableSourceMaps
	//ast, err := parser.ParseFile(nil, name, source, 0, opts)
	//pgm, err := goja.CompileAST(ast, true)

	comp := NewCompiler(logrus.StandardLogger())
	comp.Options.CompatibilityMode = CompatibilityModeExtended
	prog, _, err = comp.Compile(source, name, isESM)
	if err != nil {
		return nil, err
	}

	return
}

func mustRequire(prog *goja.Program, runtime *goja.Runtime) *Exports {
	exports, err := require(prog, runtime)
	if err != nil {
		panic(err)
	}

	return exports
}

func require(prog *goja.Program, runtime *goja.Runtime) (*Exports, error) {
	exports, err := execute(prog, runtime)
	if err != nil {
		return nil, err
	}

	named, assertOK := exports.Export().(map[string]interface{})
	if !assertOK {
		return nil, errInvalidModule
	}

	return &Exports{
		Default: exports.Get("default"),
		Named:   named,
	}, nil
}

func execute(prog *goja.Program, runtime *goja.Runtime) (*goja.Object, error) {
	module := runtime.NewObject()
	exports := runtime.NewObject()

	if err := module.Set("exports", exports); err != nil {
		return nil, err
	}

	value, err := runtime.RunProgram(prog)
	if err != nil {
		return nil, err
	}

	callable, assertOK := goja.AssertFunction(value)
	if !assertOK {
		return nil, errInvalidModule
	}

	_, err = callable(exports, module, exports)
	if err != nil {
		return nil, err
	}

	exports, assertOK = module.Get("exports").(*goja.Object)
	if !assertOK {
		return nil, errInvalidModule
	}

	return exports, nil
}

var errInvalidModule = errors.New("invalid module")

type Instance interface {
	Exports() Exports
}

type Exports struct {
	// Default is what will be the `default` export of a module
	Default interface{}
	// Named is the named exports of a module
	Named map[string]interface{}
}

type Module interface {
	// NewModuleInstance will get modules.VU that should provide the module with a way to interact with the VU
	// This method will be called for *each* require/import and should return an unique instance for each call
	NewModuleInstance(VU) Instance
}

type VU interface {
	Runtime() *goja.Runtime
}

type AgentVU struct {
	RuntimeField *goja.Runtime
}

func (m *AgentVU) Runtime() *goja.Runtime {
	return m.RuntimeField
}
