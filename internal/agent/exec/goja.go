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

func InitJsRuntime() {
	registry := new(require.Registry) // registry 能夠被 多個 goja.Runtime 共用

	MyVm.JsRuntime = goja.New()

	// below script will get/set the variables in exec context
	MyVm.JsRuntime.Set("getDatapoolVariable", func(datapool, field, seq string) string {
		return fmt.Sprintf("getDatapoolVariable(%s, %s, %s)", datapool, field, seq)
	})

	MyVm.JsRuntime.Set("getEnvironmentVariable", func(name string) string {
		return fmt.Sprintf("getVariable(%s)", name)
	})
	MyVm.JsRuntime.Set("setEnvironmentVariable", func(name, val string) string {
		return fmt.Sprintf("setVariable(%s, %s)", name, val)
	})
	MyVm.JsRuntime.Set("clearEnvironmentVariable", func(name string) string {
		return fmt.Sprintf("clearEnvironmentVariable(%s)", name)
	})

	MyVm.JsRuntime.Set("getVariable", func(name string) string {
		return fmt.Sprintf("getVariable(%s)", name)
	})
	MyVm.JsRuntime.Set("setVariable", func(name, val string) string {
		return fmt.Sprintf("setVariable(%s, %s)", name, val)
	})
	MyVm.JsRuntime.Set("clearVariable", func(name string) string {
		return fmt.Sprintf("clearVariable(%s)", name)
	})

	// load global script
	MyRequire = registry.Enable(MyVm.JsRuntime)
	pth := filepath.Join(consts.TmpDir, "dp.js")
	fileUtils.WriteFile(pth, scriptHelper.GetScript(scriptHelper.ScriptGlobal))
	dp, err := MyRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	MyVm.JsRuntime.Set("dp", dp)
}

func ExecJs(script string, variables domain.Variables, datapools domain.Datapools) (ret goja.Value) {
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
