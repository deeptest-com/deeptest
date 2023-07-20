package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"path/filepath"
)

var (
	MyVm      JsVm
	MyRequire *require.RequireModule
)

type JsVm struct {
	JsRuntime *goja.Runtime
}

func ExecScript(scriptObj *domain.ScriptBase) (err error) {
	if MyVm.JsRuntime == nil {
		InitJsRuntime()
	}

	if scriptObj.Content == "" {
		return
	}

	result, err := MyVm.JsRuntime.RunString(scriptObj.Content)
	if err != nil {
		scriptObj.ResultStatus = consts.Fail
		scriptObj.Output = fmt.Sprintf("%v, ERROR: %s", result, err.Error())
		logUtils.Error(scriptObj.Output)
	} else {
		scriptObj.ResultStatus = consts.Pass
		scriptObj.Output = fmt.Sprintf("%v", result)
	}

	return
}

func InitJsRuntime() {
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用

	MyVm.JsRuntime = goja.New()

	// below script will get/set the variables in exec context
	MyVm.JsRuntime.Set("getDatapoolVariable", func(dpName, field, seq string) (ret interface{}) {
		rowIndex := getDatapoolRow(dpName, seq, ExecScene.Datapools)

		ret = ExecScene.Datapools[dpName][rowIndex][field]
		if ret == nil {
			ret = "NOT_FOUND"
		}

		return
	})

	MyVm.JsRuntime.Set("getVariable", func(name string) interface{} {
		return getVariableValue(name)
	})
	MyVm.JsRuntime.Set("setVariable", func(name, val string) {
		SetVariable(CurrScenarioProcessorId, name, val, consts.Public)
	})
	MyVm.JsRuntime.Set("clearVariable", func(name string) {
		ClearVariable(CurrScenarioProcessorId, name)
	})

	// load global script
	MyRequire = registry.Enable(MyVm.JsRuntime)
	pth := filepath.Join(consts.TmpDir, "deeptest.js")
	fileUtils.WriteFile(pth, scriptHelper.GetScript(scriptHelper.ScriptDeepTest))
	dp, err := MyRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	MyVm.JsRuntime.Set("dp", dp)
}
