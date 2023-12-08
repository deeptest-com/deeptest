package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	gojaUtils "github.com/aaronchen2k/deeptest/internal/pkg/goja"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"log"
	"path/filepath"
	"reflect"
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

func ExecScript(scriptObj *domain.ScriptBase, projectId uint) (err error) {
	if execVm.JsRuntime == nil {
		InitJsRuntime(projectId)
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

func InitJsRuntime(projectId uint) {
	if execVm.JsRuntime != nil {
		jslibHelper.RefreshRemoteAgentJslibs(execVm.JsRuntime, execRequire, projectId, ServerUrl, ServerToken)
		return
	}

	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用

	execVm.JsRuntime = goja.New()
	execVm.JsRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	jslibHelper.LoadChaiJslibs(execVm.JsRuntime)

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
	jslibHelper.RefreshRemoteAgentJslibs(execVm.JsRuntime, execRequire, 0, ServerUrl, ServerToken)
}

func GetReqValueFromGoja() (err error) {
	_, err = execVm.JsRuntime.RunString("getReqValueFromGoja(dt.request);")
	return
}
func GetRespValueFromGoja() (err error) {
	_, err = execVm.JsRuntime.RunString("getRespValueFromGoja(dt.response);")
	return
}

func defineJsFuncs() (err error) {
	err = execVm.JsRuntime.Set("getDatapoolVariable", func(dpName, field, seq string) (ret interface{}) {
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

	err = execVm.JsRuntime.Set("getVariable", func(name string) interface{} {
		var scopeId uint
		if CurrScenarioProcessor != nil {
			scopeId = CurrScenarioProcessor.ParentId
		}
		vari, _ := GetVariable(scopeId, name)
		return vari.Value
	})
	err = execVm.JsRuntime.Set("setVariable", func(name string, val interface{}) {
		var scopeId uint
		if CurrScenarioProcessor != nil {
			scopeId = CurrScenarioProcessor.ParentId
		}
		ret, err := SetVariable(scopeId, name, val, consts.ExtractorResultTypeObject, consts.Public)

		if err == nil {
			VariableSettings = append(VariableSettings, ret)
		}

		return
	})
	err = execVm.JsRuntime.Set("clearVariable", func(name string) {
		var scopeId uint
		if CurrScenarioProcessor != nil {
			scopeId = CurrScenarioProcessor.ParentId
		}
		ClearVariable(scopeId, name)
	})

	err = execVm.JsRuntime.Set("getReqValueFromGoja", func(value domain.BaseRequest) {
		CurrRequest = value
	})
	err = execVm.JsRuntime.Set("getRespValueFromGoja", func(value domain.DebugResponse) {
		if httpHelper.IsJsonResp(value) {
			bytes, _ := json.Marshal(value.Data)
			value.Content = string(bytes)
			CurrResponse = value
		} else {
			value.Content = value.Data.(string)
			CurrResponse = value
		}
	})

	// http request
	err = execVm.JsRuntime.Set("sendRequest", func(data goja.Value, cb func(interface{}, interface{})) {
		req := gojaUtils.GenRequest(data, execVm.JsRuntime)

		resp, err2 := Invoke(&req)
		cb(err2, resp)

		log.Println("result")
	})

	// log
	err = execVm.JsRuntime.Set("log", func(value interface{}) {
		if value == nil {
			logs = append(logs, "空")
			return
		}

		typ := reflect.TypeOf(value).Name()

		if typ == "string" {
			logs = append(logs, value.(string))
		} else {
			bytes, _ := json.Marshal(value)
			logs = append(logs, string(bytes))
		}
	})

	return
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

// call go SetValueToGoja = call js _setData
var (
	_setValueFunc func(name string, value interface{})
)

func SetValueToGoja(name string, value interface{}) {
	_setValueFunc(name, value)
}
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
