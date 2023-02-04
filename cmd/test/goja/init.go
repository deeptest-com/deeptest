package main

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func init() {
	registry := new(require.Registry)
	// use custom srcloader
	//registry := require.NewRegistryWithLoader(func(path string) ([]byte, error) {
	//	return Asset(path)
	//})
	Myvm = MyVM{
		JsRuntime: goja.New(),
		Script:    `require("dalong.js")("dalong","ddddd")`,
	}
	Myrequire = registry.Enable(Myvm.JsRuntime)
	Myvm.init()
}
