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
	"log"
	"path/filepath"
	"reflect"
	"strings"
)

func ExecScript(scriptObj *domain.ScriptBase, projectId uint, execUuid string) (err error) {
	execRuntime, _ := GetGojaRuntime(execUuid)

	if execRuntime == nil {
		InitJsRuntime(projectId, execUuid)
	}

	SetGojaVariables(execUuid, []domain.ExecVariable{})

	if scriptObj.Content == "" {
		return
	}

	SetGojaLogs(execUuid, nil)
	resultVal, err := execRuntime.RunString(scriptObj.Content)

	result := fmt.Sprintf("%v", resultVal)
	if result == "undefined" {
		result = "空"
	}

	output := strings.Join(GetGojaLogs(execUuid), "; ")

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

func InitJsRuntime(projectId uint, execUuid string) {
	execRuntime, execRequire := GetGojaRuntime(execUuid)

	if execRuntime != nil { // just load new project's Jslibs if needed
		jslibHelper.RefreshRemoteAgentJslibs(execRuntime, execRequire, projectId, GetServerUrl(execUuid), GetServerToken(execUuid))
		return
	}

	InitGojaRuntime(execUuid)
	execRuntime, execRequire = GetGojaRuntime(execUuid)

	jslibHelper.LoadChaiJslibs(execRuntime)

	defineJsFuncs(execUuid)
	defineGoFuncs(execUuid)

	// load global script
	pth := filepath.Join(consts.TmpDir, "deeptest.js")
	fileUtils.WriteFile(pth, scriptHelper.GetScript(scriptHelper.ScriptDeepTest))
	dt, err := execRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	execRuntime.Set("dt", dt)

	// import other custom libs
	jslibHelper.RefreshRemoteAgentJslibs(execRuntime, execRequire, 0, GetServerUrl(execUuid), GetServerToken(execUuid))
}

func GetReqValueFromGoja(execUuid string) (err error) {
	execRuntime, _ := GetGojaRuntime(execUuid)
	_, err = execRuntime.RunString("getReqValueFromGoja(dt.request);")
	return
}
func GetRespValueFromGoja(execUuid string) (err error) {
	execRuntime, _ := GetGojaRuntime(execUuid)
	_, err = execRuntime.RunString("getRespValueFromGoja(dt.response);")
	return
}

func defineJsFuncs(execUuid string) (err error) {
	execRuntime, _ := GetGojaRuntime(execUuid)

	err = execRuntime.Set("getDatapoolVariable", func(dpName, field, seq string) (ret interface{}) {
		execScene := GetExecScene(execUuid)

		rowIndex := getDatapoolRow(dpName, seq, execScene.Datapools, execUuid)

		if execScene.Datapools[dpName] == nil {
			ret = "DATAPOOL_NOT_FOUND: " + dpName
			return
		}

		if rowIndex > len(execScene.Datapools[dpName])-1 {
			ret = "DATAPOOL_INDEX_OUT_OF_RANGE"
			return
		}

		ret = execScene.Datapools[dpName][rowIndex][field]
		if ret == nil {
			ret = "DATAPOOL_VARIABLE_NOT_FOUND: " + field
		}

		return
	})

	err = execRuntime.Set("getVariable", func(name string) interface{} {
		var scopeId uint
		if GetCurrScenarioProcessor(execUuid) != nil {
			scopeId = GetCurrScenarioProcessor(execUuid).ParentId
		}
		vari, _ := GetVariable(scopeId, name, execUuid)
		return vari.Value
	})
	err = execRuntime.Set("setVariable", func(name string, val interface{}) {
		var scopeId uint
		if GetCurrScenarioProcessor(execUuid) != nil {
			scopeId = GetCurrScenarioProcessor(execUuid).ParentId
		}
		ret, err := SetVariable(scopeId, name, val, consts.ExtractorResultTypeObject, consts.Public, execUuid)

		if err == nil {
			AppendGojaVariables(execUuid, ret)
		}

		return
	})
	err = execRuntime.Set("clearVariable", func(name string) {
		var scopeId uint
		if GetCurrScenarioProcessor(execUuid) != nil {
			scopeId = GetCurrScenarioProcessor(execUuid).ParentId
		}
		ClearVariable(scopeId, name, execUuid)
	})

	err = execRuntime.Set("getReqValueFromGoja", func(value domain.BaseRequest) {
		SetCurrRequest(execUuid, value)
	})
	err = execRuntime.Set("getRespValueFromGoja", func(value domain.DebugResponse) {
		if httpHelper.IsJsonResp(value) {
			bytes, _ := json.Marshal(value.Data)
			value.Content = string(bytes)
			SetCurrResponse(execUuid, value)
		} else {
			value.Content = value.Data.(string)
			SetCurrResponse(execUuid, value)
		}
	})

	// http request
	err = execRuntime.Set("sendRequest", func(data goja.Value, cb func(interface{}, interface{})) {
		req := gojaUtils.GenRequest(data, execRuntime)

		resp, err2 := Invoke(&req)
		cb(err2, resp)

		log.Println("result")
	})

	// log
	err = execRuntime.Set("log", func(value interface{}) {
		if value == nil {
			AppendGojaLogs(execUuid, "空")
			return
		}

		typ := reflect.TypeOf(value).Name()

		if typ == "string" {
			AppendGojaLogs(execUuid, value.(string))
		} else {
			bytes, _ := json.Marshal(value)
			AppendGojaLogs(execUuid, string(bytes))
		}
	})

	return
}

func SetReqValueToGoja(req *domain.BaseRequest) {
	SetValueToGoja("request", req)
}
func SetRespValueToGoja(resp *domain.DebugResponse) {
	// set resp.Data to json object for goja edit
	if httpHelper.IsJsonResp(*resp) {
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
func defineGoFuncs(execUuid string) {
	execRuntime, _ := GetGojaRuntime(execUuid)

	// set data
	script := `function _setData(name, val) {
					dt[name] = val
				}`
	_, err := execRuntime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	err = execRuntime.ExportTo(execRuntime.Get("_setData"), &_setValueFunc)
}
