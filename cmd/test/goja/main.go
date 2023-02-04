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
	registry := new(require.Registry) // registry 能夠被 多個 goja.Runtime 共用
	// 創建 虛擬機
	vm := goja.New()
	vm.Set(`println`, func(args ...interface{}) {
		fmt.Println(args...)
	})
	vm.Set("login", func(name, password string) string {
		return fmt.Sprintf("%s-%s", name, password)
	})

	req := registry.Enable(vm) // 爲 Runtime 啓用模塊

	m, err := req.Require("./underscore-min.js")
	if err != nil {
		log.Panic(err)
	}
	vm.Set("_", m)

	shortid, err := req.Require("./shortid.js")
	if err != nil {
		log.Panic(err)
	}
	vm.Set("shortid", shortid)

	app, err := req.Require("./app.js")
	if err != nil {
		log.Panic(err)
	}

	ob := app.ToObject(vm)
	fmt.Println(ob.Get("filteruser").String())
	fmt.Println(ob.Get("id").String())
	fmt.Println(ob.Get("id2").String())

	fmt.Println(ob.Get("id2").String())

	val, err := vm.RunString("var a= 1111; a += shortid.generate(); a;")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(val)
}
