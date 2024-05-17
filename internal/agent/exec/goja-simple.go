package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

func ExecJsFuncSimple(expr string, tenantId consts.TenantId, projectId uint, execUuid string, loadCustom bool) (ret string) {
	InitJsRuntimeSimple(tenantId, projectId, execUuid, loadCustom)
	execRuntime, _ := jslibHelper.GetGojaRuntime(tenantId, projectId)

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

	// load buildin funcs
	content := scriptHelper.GetScript(scriptHelper.ScriptFuncs)

	_, err := execRuntime.RunString(content)
	if err != nil {
		logUtils.Infof("goja require buildin funcs failed, path: %s, err: %s.", scriptHelper.ScriptFuncs, err.Error())
	}

	// import other custom libs
	if loadCustom {
		jslibHelper.RefreshRemoteAgentJslibs(execRuntime, execRequire, tenantId, projectId, GetServerUrl(execUuid), GetServerToken(execUuid))
	}
}
