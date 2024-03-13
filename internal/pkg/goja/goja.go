package gojaUtils

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitGojaRuntime() (execRuntime *goja.Runtime, execRequire *require.RequireModule) {
	execRuntime = goja.New()
	execRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用
	execRequire = registry.Enable(execRuntime)

	return
}
