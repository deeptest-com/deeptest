package mockHelper

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"path/filepath"
	"strings"
)

var (
	mockVm      JsVm
	mockRequire *require.RequireModule

	logs []string

	CurrResponse mockGenerator.Response
)

type JsVm struct {
	JsRuntime *goja.Runtime
}

func ExecScript(script string) (err error) {
	if mockVm.JsRuntime == nil {
		InitJsRuntime()
	}

	if script == "" {
		return
	}

	logs = nil
	resultVal, err := mockVm.JsRuntime.RunString(script)

	result := fmt.Sprintf("%v", resultVal)
	if result == "undefined" {
		result = "空"
	}

	output := strings.Join(logs, "; ")

	if err != nil {
		logUtils.Error(output + ", " + err.Error())
	}

	logUtils.Infof("%v", logs)

	return
}

var (
	_setValueFunc func(name string, value interface{})
)

func defineJsFuncs() {
	mockVm.JsRuntime.Set("log", func(value interface{}) {
		bytes, _ := json.Marshal(value)
		logs = append(logs, string(bytes))
	})

	mockVm.JsRuntime.Set("getRespValueFromGoja", func(value mockGenerator.Response) {
		if httpHelper.IsJsonRespType(value.ContentType) {
			bytes, _ := json.Marshal(value.Data)
			value.Content = string(bytes)
			CurrResponse = value
		} else {
			value.Content = value.Data.(string)
			CurrResponse = value
		}
	})
}
func defineGoFuncs() {
	// set data
	script := `function _setData(name, val) {
					dt[name] = val
				}`
	_, err := mockVm.JsRuntime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	err = mockVm.JsRuntime.ExportTo(mockVm.JsRuntime.Get("_setData"), &_setValueFunc)
}

func GetRespValueFromGoja() (err error) {
	_, err = mockVm.JsRuntime.RunString("getRespValueFromGoja(dt.response);")
	return
}
func SetRespValueToGoja(resp mockGenerator.Response) {
	SetValueToGoja("response", resp)
}
func SetValueToGoja(name string, value interface{}) {
	_setValueFunc(name, value)
}

func InitJsRuntime() {
	if mockVm.JsRuntime != nil {
		return
	}

	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用

	mockVm.JsRuntime = goja.New()
	mockVm.JsRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	defineJsFuncs()
	defineGoFuncs()

	// load global script
	mockRequire = registry.Enable(mockVm.JsRuntime)
	pth := filepath.Join(consts.WorkDir, "mock.js")

	script := scriptHelper.GetScript(scriptHelper.ScriptMock)
	fileUtils.WriteFile(pth, script)

	dt, err := mockRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	mockVm.JsRuntime.Set("dt", dt)
}
