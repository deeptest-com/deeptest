package main

import (
	"fmt"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	gojaUtils "github.com/aaronchen2k/deeptest/internal/pkg/goja"
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
	//registry := new(require.Registry)
	vm := goja.New()

	//req := registry.Enable(vm)

	err := vm.Set("log", func(val interface{}) {
		log.Println(val)
	})

	// http request
	err = vm.Set("sendRequest", func(data goja.Value, cb func(interface{}, interface{})) {
		req := gojaUtils.GenRequest(data, vm)

		resp, err2 := agentExec.Invoke(&req)
		cb(err2, resp)

		log.Println("result")
	})

	if err != nil {
		log.Println(err)
	}

	script := `
sendRequest("http://111.231.16.35:9000/get", function (err, resp) {
	log(err ? err : resp);
});
		`
	out, err := vm.RunString(script)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(out)

	script2 := `
const postRequest = {
  url: 'https://postman-echo.com/post',
  method: 'POST',
  params: {
    'p1': '1',
    'p2': '2'
  },
  headers: {
    'Content-Type': 'application/json',
    'X-Foo': 'bar'
  },
  cookies: {},
  body: {
    mode: 'raw',
    raw: JSON.stringify({ key: 'this is json' })
  },

  basicAuth: {'username': 'admin', 'password': 'pass'},
  bearerToken: {'token': 'abc123'}
};
sendRequest(postRequest, (err, resp) => {
  log(err ? err : resp);
});`

	out, err = vm.RunString(script2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(out)
}
