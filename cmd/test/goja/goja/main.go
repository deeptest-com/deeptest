package main

import (
	"encoding/json"
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
		log.Println(bytes)
	})

	script := `
		//import { describe, expect } from "chai.js";
		//log(expect);

	  	//describe('Fetch a list of public crocodiles', () => {
			const response = {status: 2001};
			var r = expect(response.status, 'response status').to.equal(200);
			log(r);
	  	//});
	`
	out, err := agentVu.Runtime().RunString(script)
	if err != nil {
		log.Println(err)
	}

	log.Println(out)
}
