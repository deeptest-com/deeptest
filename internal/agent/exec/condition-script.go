package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"path/filepath"
	"strings"
)

var (
	execVm      JsVm
	execRequire *require.RequireModule

	VariableSettings []domain.ExecVariable

	logs []string
)

type JsVm struct {
	JsRuntime *goja.Runtime
}

func ExecScript(scriptObj *domain.ScriptBase) (err error) {
	if execVm.JsRuntime == nil {
		InitJsRuntime()
	}

	VariableSettings = []domain.ExecVariable{}

	if scriptObj.Content == "" {
		return
	}

	logs = nil
	resultVal, err := execVm.JsRuntime.RunString(scriptObj.Content)

	result := fmt.Sprintf("%v", resultVal)
	if result == "undefined" {
		result = "空"
	}

	output := strings.Join(logs, "; ")

	if err != nil {
		scriptObj.ResultStatus = consts.Fail
		scriptObj.Output = fmt.Sprintf("RESULT: %v; OUTPUT: %s; ERROR: %s", result, output, err.Error())
		logUtils.Error(scriptObj.Output)

	} else {
		scriptObj.ResultStatus = consts.Pass
		scriptObj.Output = fmt.Sprintf("%s", output)
	}

	return
}

func InitJsRuntime() {
	if execVm.JsRuntime != nil {
		jslibHelper.LoadAgentJslibs(execVm.JsRuntime, execRequire, ServerUrl, ServerToken)
		return
	}

	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用

	execVm.JsRuntime = goja.New()
	execVm.JsRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	defineJsFuncs()
	defineGoFuncs()

	// load global script
	execRequire = registry.Enable(execVm.JsRuntime)
	pth := filepath.Join(consts.TmpDir, "deeptest.js")
	fileUtils.WriteFile(pth, scriptHelper.GetScript(scriptHelper.ScriptDeepTest))
	dt, err := execRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	execVm.JsRuntime.Set("dt", dt)

	// import other custom libs
	jslibHelper.LoadAgentJslibs(execVm.JsRuntime, execRequire, ServerUrl, ServerToken)
}

func GetReqValueFromGoja() (err error) {
	_, err = execVm.JsRuntime.RunString("getReqValueFromGoja(dt.request);")
	return
}
func GetRespValueFromGoja() (err error) {
	_, err = execVm.JsRuntime.RunString("getRespValueFromGoja(dt.response);")
	return
}

func defineJsFuncs() {
	execVm.JsRuntime.Set("getDatapoolVariable", func(dpName, field, seq string) (ret interface{}) {
		rowIndex := getDatapoolRow(dpName, seq, ExecScene.Datapools)

		if ExecScene.Datapools[dpName] == nil {
			ret = "DATAPOOL_NOT_FOUND: " + dpName
			return
		}

		if rowIndex > len(ExecScene.Datapools[dpName])-1 {
			ret = "DATAPOOL_INDEX_OUT_OF_RANGE"
			return
		}

		ret = ExecScene.Datapools[dpName][rowIndex][field]
		if ret == nil {
			ret = "DATAPOOL_VARIABLE_NOT_FOUND: " + field
		}

		return
	})

	execVm.JsRuntime.Set("getVariable", func(name string) interface{} {
		var scopeId uint
		if CurrScenarioProcessor != nil {
			scopeId = CurrScenarioProcessor.ParentId
		}
		vari, _ := GetVariable(scopeId, name)
		return vari.Value
	})
	execVm.JsRuntime.Set("setVariable", func(name, val string) {
		var scopeId uint
		if CurrScenarioProcessor != nil {
			scopeId = CurrScenarioProcessor.ParentId
		}
		ret, err := SetVariable(scopeId, name, val, consts.Public)

		if err == nil {
			VariableSettings = append(VariableSettings, ret)
		}

		return
	})
	execVm.JsRuntime.Set("clearVariable", func(name string) {
		var scopeId uint
		if CurrScenarioProcessor != nil {
			scopeId = CurrScenarioProcessor.ParentId
		}
		ClearVariable(scopeId, name)
	})

	execVm.JsRuntime.Set("getReqValueFromGoja", func(value domain.BaseRequest) {
		CurrRequest = value
	})
	execVm.JsRuntime.Set("getRespValueFromGoja", func(value domain.DebugResponse) {
		if httpHelper.IsJsonResp(value) {
			bytes, _ := json.Marshal(value.Data)
			value.Content = string(bytes)
			CurrResponse = value
		} else {
			value.Content = value.Data.(string)
			CurrResponse = value
		}
	})

	execVm.JsRuntime.Set("log", func(value interface{}) {
		bytes, _ := json.Marshal(value)
		logs = append(logs, string(bytes))
	})
}

var (
	_setValueFunc func(name string, value interface{})
)

func defineGoFuncs() {
	// set data
	script := `function _setData(name, val) {
					dt[name] = val
				}`
	_, err := execVm.JsRuntime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	err = execVm.JsRuntime.ExportTo(execVm.JsRuntime.Get("_setData"), &_setValueFunc)
}

func SetReqValueToGoja(req domain.BaseRequest) {
	SetValueToGoja("request", req)
}
func SetRespValueToGoja(resp domain.DebugResponse) {
	// set resp.Data to json object for goja edit
	if httpHelper.IsJsonResp(resp) {
		var data interface{}
		err := json.Unmarshal([]byte(resp.Content), &data)
		if err == nil {
			resp.Data = data
		} else {
			resp.Data = resp.Content
		}
	} else {
		resp.Data = resp.Content
	}

	SetValueToGoja("response", resp)
}
func SetValueToGoja(name string, value interface{}) {
	_setValueFunc(name, value)
}
