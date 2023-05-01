package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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

func InitJsRuntime() {
	registry := new(require.Registry) // registry 能夠被 多個 goja.Runtime 共用

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
		SetVariable(CurrProcessorId, name, val, consts.Public)
	})
	MyVm.JsRuntime.Set("clearVariable", func(name string) {
		ClearVariable(CurrProcessorId, name)
	})

	// load global script
	MyRequire = registry.Enable(MyVm.JsRuntime)
	pth := filepath.Join(consts.TmpDir, "dp.js")
	fileUtils.WriteFile(pth, scriptHelper.GetScript(scriptHelper.ScriptDp))
	dp, err := MyRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	MyVm.JsRuntime.Set("dp", dp)
}

func ExecJs(script string) (ret goja.Value) {
	if MyVm.JsRuntime == nil {
		InitJsRuntime()
	}

	if script == "" {
		return
	}

	ret, err := MyVm.JsRuntime.RunString(script)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	return
}
