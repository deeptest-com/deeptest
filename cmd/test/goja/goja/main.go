package main

import (
	"encoding/json"
	"fmt"
	gojaPlugin "github.com/aaronchen2k/deeptest/internal/pkg/goja/plugin"
	"github.com/dop251/goja"
	"log"
)

func main() {
	vm := goja.New()
	agentVu := gojaPlugin.AgentVU{
		RuntimeField: vm,
	}

	agentVu.Runtime().Set("require", func(call goja.FunctionCall) goja.Value { return goja.Undefined() })
	agentVu.Runtime().Set("global", agentVu.Runtime().GlobalObject())

	//mochaModule := gojaPlugin.NewMocha()
	//mochaInst := mochaModule.NewModuleInstance(&agentVu)
	//log.Println(mochaInst)

	chaiModule := gojaPlugin.NewChai()
	chaiInst := chaiModule.NewModuleInstance(&agentVu)
	log.Println(chaiInst)

	//vm.Set("describe", mochaInst.Exports().Named["describe"])
	vm.Set("expect", chaiInst.Exports().Named["expect"])

	err := vm.Set("log", func(value interface{}) {
		bytes, _ := json.Marshal(value)
		log.Println("=== ", bytes)
	})

	err = vm.Set("check", func(ok bool, name, msg string) {
		log.Println(fmt.Sprintf("%t, %s, %s", ok, name, msg)) // add to assert
	})

	script := `
		function test(name, cb) {
			try {
				cb();
			} catch(err){
				check(false, name, err)
				return
			}

			check(true, name, '')
		}

	  	test('get request', () => {
			var r1 = expect(200, 'status').to.equal(200);
			log(r1)
			var r2 = expect(-1, 'code').to.equal(0);
			log(r2)
	  	});

		test('post request', () => {
			var r1 = expect(200, 'status').to.equal(200);
			log(r1)
			var r2 = expect(0, 'code').to.equal(0);
			log(r2)
	  	});
	`
	out, err := agentVu.Runtime().RunString(script)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(out)
}
