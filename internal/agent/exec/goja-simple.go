package agentExec

import (
	"fmt"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
)

func ExecJsFuncSimple(nameWithProp string, session *ExecSession, loadCustom bool) (
	ret string) {

	InitProjectJsRuntimeSimple(session, loadCustom)
	execRuntime := session.GojaRuntime

	defineGoFuncsSimple(execRuntime)

	variables, _ := GetAllVariables(nameWithProp, session)

	for _, variable := range variables {
		// add variable to goja runtime
		setValueToGojaSimple(variable.Name, variable.Value)
	}

	resultVal, err := execRuntime.RunString(nameWithProp)
	if err != nil {
		ret = err.Error()
		return
	}

	ret = fmt.Sprintf("%v", resultVal.Export())
	if ret == "undefined" {
		ret = "空"
	}

	return
}

func InitProjectJsRuntimeSimple(session *ExecSession, loadCustom bool) {
	gojaRuntime, execRequire := GenerateGojaRuntime()

	// load buildin funcs
	content := scriptHelper.GetScript(scriptHelper.ScriptCustom)

	_, err := gojaRuntime.RunString(content)
	if err != nil {
		logUtils.Infof("goja require buildin funcs failed, path: %s, err: %s.", scriptHelper.ScriptCustom, err.Error())
	}

	// import other custom libs
	if loadCustom {
		jslibHelper.RefreshRemoteAgentJslibs(gojaRuntime, execRequire, session.VuNo, session.TenantId, session.ProjectId, session.ServerUrl, session.ServerToken)
	}
}

var (
	_setValueFuncSimple func(name string, value interface{})
)

func setValueToGojaSimple(name string, value interface{}) {
	_setValueFuncSimple(name, value)
}

func defineGoFuncsSimple(execRuntime *goja.Runtime) {
	script := ` function _setDataSimple(name, val) {
					dt_temp[name] = val
				}`
	_, err := execRuntime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	err = execRuntime.ExportTo(execRuntime.Get("_setDataSimple"), &_setValueFuncSimple)
}
