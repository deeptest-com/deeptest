package main

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"log"
)

var (
	mockJsVm      JsVm
	mockJsRequire *require.RequireModule

	mockFunc func(p interface{}) interface{}
)

type JsVm struct {
	JsRuntime *goja.Runtime
}

func main() {
	registry := new(require.Registry)
	mockJsVm.JsRuntime = goja.New()
	mockJsVm.JsRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	mockJsRequire = registry.Enable(mockJsVm.JsRuntime)

	pth := "/Users/aaron/deeptest/tmp/1.js"
	md5, err := mockJsRequire.Require(pth)

	mockJsVm.JsRuntime.Set("md5", md5)

	script := `md5('123')`
	out, err := mockJsVm.JsRuntime.RunString(script)

	log.Println(out)

	if err != nil {
		log.Println(err.Error())
	}
}
