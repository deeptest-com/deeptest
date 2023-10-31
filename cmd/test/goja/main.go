package main

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"log"
)

var (
	Myvm      MyVM
	Myrequire *require.RequireModule
)

type MyVM struct {
	JsRuntime *goja.Runtime
	Script    string
}

func main() {
	registry := new(require.Registry)
	vm := goja.New()

	//req := registry.Enable(vm)
	//mock, err := req.Require("./cmd/test/goja/lib/mock.js")
	//vm.Set("mock", mock)
	//
	//script := `
	//	function Mock(str) {
	//		const obj = JSON.parse(str);
	//		var data = mock.mock(obj)
	//		return data;
	//	}
	//	`
	//_, err = vm.RunString(script)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//var Mock func(p interface{}) interface{}
	//err = vm.ExportTo(vm.Get("Mock"), &Mock)
	//if err != nil {
	//	panic(err)
	//}
	//
	//str := `{"list|1-10": [{"id|+1": 1}]}`
	//
	//out := Mock(str)
	//
	//fmt.Println(out)

	// test functions
	req := registry.Enable(vm)
	module, err := req.Require("./cmd/test/goja/lib/funcs.js")

	vm.Set("math", module)

	script := `math.add(1, 1)`
	out, err := vm.RunString(script)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(out)
}
