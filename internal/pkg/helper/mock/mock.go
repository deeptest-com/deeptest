package mockHelper

import (
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	httpHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/http"
	jslibHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/jslib"
	mockGenerator "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	scriptHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/script"
	fileUtils "github.com/deeptest-com/deeptest/pkg/lib/file"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
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
func SetReqValueToGoja(req mockGenerator.Request) {
	SetValueToGoja("request", req)
}
func SetRespValueToGoja(resp mockGenerator.Response) {
	SetValueToGoja("response", resp)
}
func SetValueToGoja(name string, value interface{}) {
	_setValueFunc(name, value)
}

func InitJsRuntime(tenantId consts.TenantId, projectId uint) {
	if mockVm.JsRuntime != nil {
		jslibHelper.LoadServerJslibs(tenantId, mockVm.JsRuntime, mockRequire, projectId)
		return
	}

	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用

	mockVm.JsRuntime = goja.New()
	mockVm.JsRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	defineJsFuncs()
	defineGoFuncs()

	// load global script
	mockRequire = registry.Enable(mockVm.JsRuntime)

	// import deeptest lib
	tmpPath := fmt.Sprintf("%s/mock.js", consts.TmpDirRelativeServer)
	tmpContent := scriptHelper.GetScript(scriptHelper.ScriptMock)
	fileUtils.WriteFileIfNotExist(tmpPath, tmpContent)

	dt, err := mockRequire.Require("./" + tmpPath)
	if err != nil {
		logUtils.Infof("goja require failed, path: %s, err: %s.", tmpPath, err.Error())
	}

	mockVm.JsRuntime.Set("dt", dt)

	// import other custom libs
	jslibHelper.LoadServerJslibs(tenantId, mockVm.JsRuntime, mockRequire, projectId)
}
