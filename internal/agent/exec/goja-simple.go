package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
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

func NewGojaSimple() (ret *GojaSimple) {
	s := GojaSimple{}
	ret = &s

	return
}

func (e *GojaSimple) ExecJsFuncSimple(content string, session *ExecSession, loadCustom bool) (
	ret interface{}, params domain.VarKeyValuePair) {
	params = domain.VarKeyValuePair{}

	e.InitJsRuntimeSimple(session, loadCustom)

	// add variables to goja runtime
	variables := GetAllValidVariables(session)
	for _, variable := range variables {
		e.execRuntime.Set(variable.Name, variable.Value)

		params[variable.Name] = variable.Value
	}

	resultVal, err := e.execRuntime.RunString(content)
	if err != nil {
		logUtils.Info(err.Error())
		ret = err.Error()
		return
	}

	ret = resultVal.Export()
	if ret == nil {
		ret = "空"
	}

	return
}

func (e *GojaSimple) InitJsRuntimeSimple(session *ExecSession, loadCustom bool) {
	e.execRuntime, e.execRequire = GenerateGojaRuntime()

	// init e.execRuntime, not session.GojaRuntime
	defineJsFuncs(e.execRuntime, e.execRequire, session, true)
	loadDeeptestScript(e.execRuntime, e.execRequire, session, true)

	// load buildin functions
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
