package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"strings"
)

func ExecJsFuncSimple(expr string, tenantId consts.TenantId, projectId uint, execUuid string,
	variables []string, loadCustom bool) (ret string) {

	InitJsRuntimeSimple(tenantId, projectId, execUuid, loadCustom)
	execRuntime, _ := jslibHelper.GetGojaRuntime(tenantId, projectId)

	defineGoFuncsSimple(execRuntime)

	for _, varName := range variables {
		// replace placeholder
		expr = strings.Replace(expr, fmt.Sprintf("#[%s]", varName),
			fmt.Sprintf("dt_temp.%s", varName), 1)

		// add variable to goja runtime
		varValue := getPlaceholderVariableValue(strings.TrimLeft(varName, "+"), execUuid)
		setValueToGojaSimple(varName, varValue)
	}

	resultVal, err := execRuntime.RunString(expr)
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

func InitJsRuntimeSimple(tenantId consts.TenantId, projectId uint, execUuid string, loadCustom bool) {
	jslibHelper.InitGojaRuntime(tenantId, projectId)
	execRuntime, execRequire := jslibHelper.GetGojaRuntime(tenantId, projectId)

	// define a temp obj
	script := `const dt_temp = {}`
	_, err := execRuntime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	// load buildin funcs
	content := scriptHelper.GetScript(scriptHelper.ScriptCustom)

	_, err = execRuntime.RunString(content)
	if err != nil {
		logUtils.Infof("goja require buildin funcs failed, path: %s, err: %s.", scriptHelper.ScriptCustom, err.Error())
	}

	// import other custom libs
	if loadCustom {
		jslibHelper.RefreshRemoteAgentJslibs(execRuntime, execRequire, tenantId, projectId, GetServerUrl(execUuid), GetServerToken(execUuid))
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
