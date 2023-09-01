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

	//vm.Set(`println`, func(args ...interface{}) {
	//	fmt.Println(args...)
	//})
	//vm.Set("login", func(name, password string) string {
	//	return fmt.Sprintf("%s-%s", name, password)
	//})
	//
	//req := registry.Enable(vm) // 爲 Runtime 啓用模塊
	//
	//m, err := req.Require("./underscore-min.js")
	//if err != nil {
	//	log.Panic(err)
	//}
	//vm.Set("_", m)
	//
	//shortid, err := req.Require("./shortid.js")
	//if err != nil {
	//	log.Panic(err)
	//}
	//vm.Set("shortid", shortid)
	//
	//vm.Set("getDatapoolVariable", func(datapool, field, seq string) string {
	//	return fmt.Sprintf("getDatapoolVariable(%s, %s, %s)", datapool, field, seq)
	//})
	//
	//vm.Set("getEnvironmentVariable", func(name string) string {
	//	return fmt.Sprintf("getVariable(%s)", name)
	//})
	//vm.Set("setEnvironmentVariable", func(name, val string) string {
	//	return fmt.Sprintf("setVariable(%s, %s)", name, val)
	//})
	//vm.Set("clearEnvironmentVariable", func(name string) string {
	//	return fmt.Sprintf("clearEnvironmentVariable(%s)", name)
	//})
	//
	//vm.Set("getVariable", func(name string) string {
	//	return fmt.Sprintf("getVariable(%s)", name)
	//})
	//vm.Set("setVariable", func(name, val string) string {
	//	return fmt.Sprintf("setVariable(%s, %s)", name, val)
	//})
	//vm.Set("clearVariable", func(name string) string {
	//	return fmt.Sprintf("clearVariable(%s)", name)
	//})

	//dp, err := req.Require("./cmd/test/goja/lib/dp.js")
	//if err != nil {
	//	log.Panic(err)
	//}
	//vm.Set("dp", dp)

	//app, err := req.Require("./app.js")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//ob := app.ToObject(vm)
	//fmt.Println(ob.Get("filteruser").String())
	//fmt.Println(ob.Get("id").String())
	//fmt.Println(ob.Get("id2").String())

	req := registry.Enable(vm)
	mock, err := req.Require("./cmd/test/goja/lib/mock.js")
	vm.Set("mock", mock)

	script := `
		function Mock(str) {
			const obj = JSON.parse(str);
			var data = mock.mock(obj)
			return data;
		}
		`
	_, err = vm.RunString(script)
	if err != nil {
		log.Panic(err)
	}

	var Mock func(p interface{}) interface{}
	err = vm.ExportTo(vm.Get("Mock"), &Mock)
	if err != nil {
		panic(err)
	}

	str := `{"list|1-10": [{"id|+1": 1}]}`

	out := Mock(str)

	fmt.Println(out)
}
