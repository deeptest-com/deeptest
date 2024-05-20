package agentExec

import (
	"fmt"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type GojaSimple struct {
	execRuntime *goja.Runtime
	execRequire *require.RequireModule

	_setValueFuncSimple func(name string, value interface{})
}

func (e *GojaSimple) ExecJsFuncSimple(content string, session *ExecSession, loadCustom bool) (
	ret string) {

	e.InitJsRuntimeSimple(session, loadCustom)

	// add variables to goja runtime
	variables := GetAllValidVariables(session)
	for _, variable := range variables {
		e.execRuntime.Set(variable.Name, variable.Value)
	}

	resultVal, err := e.execRuntime.RunString(content)
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

func (e *GojaSimple) InitJsRuntimeSimple(session *ExecSession, loadCustom bool) {
	e.execRuntime, e.execRequire = GenerateGojaRuntime()

	// load buildin funcs
	content := scriptHelper.GetScript(scriptHelper.ScriptCustom)

	_, err := e.execRuntime.RunString(content)
	if err != nil {
		logUtils.Infof("goja require buildin funcs failed, path: %s, err: %s.", scriptHelper.ScriptCustom, err.Error())
	}

	// import other custom libs
	if loadCustom {
		jslibHelper.RefreshRemoteAgentJslibs(e.execRuntime, e.execRequire, session.VuNo, session.TenantId, session.ProjectId, session.ServerUrl, session.ServerToken)
	}
}
