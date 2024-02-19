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
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/kataras/iris/v12"
	"path/filepath"
	"reflect"
)

func ExecScript(scriptObj *domain.ScriptBase, tenantId consts.TenantId, projectId uint, execUuid string) (err error) {
	//InitJsRuntime(projectId, execUuid)
	execRuntime, _ := jslibHelper.GetGojaRuntime(tenantId, projectId)
	ResetGojaVariables(execUuid)
	ResetGojaLogs(execUuid)

	var logs []string

	if scriptObj.Content == "" {
		scriptObj.ResultStatus = consts.Pass
		scriptObj.ResultMsg = ""
		logs = append(logs, "")

		return

	} else {
		resultVal, err := execRuntime.RunString(scriptObj.Content)

		result := fmt.Sprintf("%v", resultVal)
		if result == "undefined" {
			result = "空"
		}

		logs = *GetGojaLogs(execUuid)

		if err != nil {
			scriptObj.ResultStatus = consts.Fail
			logs = append(logs, err.Error())
		} else {
			scriptObj.ResultStatus = consts.Pass
		}
	}

	if logs != nil {
		bytes, _ := json.Marshal(logs)
		scriptObj.Output = string(bytes)
	} else {
		scriptObj.Output = ""
	}

	return
}

func InitJsRuntime(tenantId consts.TenantId, projectId uint, execUuid string) {
	jslibHelper.InitGojaRuntime(tenantId, projectId)
	execRuntime, execRequire := jslibHelper.GetGojaRuntime(tenantId, projectId)

	jslibHelper.LoadChaiJslibs(execRuntime)

	defineJsFuncs(execUuid, tenantId, projectId)
	defineGoFuncs(tenantId, projectId)

	// load global script
	pth := filepath.Join(consts.TmpDir, fmt.Sprintf("%v/deeptest.js", tenantId))
	//pth = fmt.Sprintf("/res/goja/module/%s", "deeptest.js")
	fileUtils.WriteFile(pth, scriptHelper.GetScript(scriptHelper.ScriptDeepTest))
	dt, err := execRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	execRuntime.Set("dt", dt)

	// import other custom libs
	jslibHelper.RefreshRemoteAgentJslibs(execRuntime, execRequire, tenantId, projectId, GetServerUrl(execUuid), GetServerToken(execUuid))
}

func defineJsFuncs(execUuid string, tenantId consts.TenantId, projectId uint) (err error) {
	execRuntime, _ := jslibHelper.GetGojaRuntime(tenantId, projectId)

	/* START: called by js */
	err = execRuntime.Set("getDatapoolVariable", func(dpName, field, seq string) (ret interface{}) {
		execScene := GetExecScene(execUuid)

		rowIndex := getDatapoolRow(dpName, seq, execScene.Datapools, execUuid)

		if execScene.Datapools[dpName] == nil {
			ret = consts.INVALID_VALUE
			AppendGojaLogs(execUuid,
				jsErrMsg("DATAPOOL_NOT_FOUND:"+dpName, "getDatapoolVariable", false))
			return
		}

		if rowIndex > len(execScene.Datapools[dpName])-1 {
			ret = consts.INVALID_VALUE
			AppendGojaLogs(execUuid,
				jsErrMsg("DATAPOOL_INDEX_OUT_OF_RANGE:"+dpName, "getDatapoolVariable", false))
			return
		}

		ret = execScene.Datapools[dpName][rowIndex][field]
		if ret == nil {
			ret = consts.INVALID_VALUE

			AppendGojaLogs(execUuid,
				jsErrMsg("DATAPOOL_VARIABLE_NOT_FOUND:"+field+"@"+dpName, "getDatapoolVariable", false))
			return
		}

		return
	})

	err = execRuntime.Set("getVariable", func(name string) interface{} {
		var scopeId uint
		if GetCurrScenarioProcessor(execUuid) != nil {
			scopeId = GetCurrScenarioProcessor(execUuid).ParentId
		}
		vari, err := GetVariable(scopeId, name, execUuid)
		if err != nil {
			vari.Value = consts.INVALID_VALUE

			AppendGojaLogs(execUuid,
				jsErrMsg(err.Error(), "getVariable", false))

			return vari.Value
		}

		vari.Value, err = commUtils.ConvertValueForUse(vari.Value, vari.ValueType)
		if err != nil {
			vari.Value = consts.INVALID_VALUE
			AppendGojaLogs(execUuid,
				jsErrMsg(err.Error(), "getVariable", false))

			return vari.Value
		}

		return vari.Value
	})
	err = execRuntime.Set("setVariable", func(name string, val interface{}) {
		var scopeId uint
		if GetCurrScenarioProcessor(execUuid) != nil {
			scopeId = GetCurrScenarioProcessor(execUuid).ParentId
		}
		ret, err := SetVariable(scopeId, name, val, commUtils.ValueType(val), consts.Public, execUuid)

		if err == nil {
			AppendGojaVariables(execUuid, ret)
		} else {
			AppendGojaLogs(execUuid,
				jsErrMsg(err.Error(), "setVariable", false))
		}

		return
	})
	err = execRuntime.Set("clearVariable", func(name string) {
		var scopeId uint
		if GetCurrScenarioProcessor(execUuid) != nil {
			scopeId = GetCurrScenarioProcessor(execUuid).ParentId
		}

		err := ClearVariable(scopeId, name, execUuid)
		if err != nil {
			AppendGojaLogs(execUuid, jsErrMsg(err.Error(), "clearVariable", false))
		}
	})

	// http request
	err = execRuntime.Set("sendRequest", func(data goja.Value, cb func(interface{}, interface{})) {
		req := gojaUtils.GenRequest(data, execRuntime)

		errOfCallbackParam := ""

		resp, err2 := Invoke(&req)
		if err2 != nil {
			// AppendGojaLogs(execUuid, jsErrMsg(err2.Error(), "sendRequest", false))
			errOfCallbackParam = jsErrMsg(err2.Error(), "sendRequest", false)
		}

		cb(errOfCallbackParam, resp)
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
	/* END: called by js */

	/* START: called by go */
	err = execRuntime.Set("getReqValueFromGoja", func(execUuid string, value domain.BaseRequest) {
		SetCurrRequest(execUuid, value)
	})
	err = execRuntime.Set("getRespValueFromGoja", func(execUuid string, value domain.DebugResponse) {
		if httpHelper.IsJsonResp(value) {
			bytes, _ := json.Marshal(value.Data)
			value.Content = string(bytes)
			SetCurrResponse(execUuid, value)
		} else {
			var ok bool
			if value.Content, ok = value.Data.(string); ok {

			}
			SetCurrResponse(execUuid, value)

		}
	})
	/* END: called by go */

	return
}

func GetReqValueFromGoja(execUuid string, tenantId consts.TenantId, projectId uint) (err error) {
	execRuntime, _ := jslibHelper.GetGojaRuntime(tenantId, projectId)
	_, err = execRuntime.RunString(fmt.Sprintf("getReqValueFromGoja('%s', dt.request);", execUuid))
	return
}
func GetRespValueFromGoja(execUuid string, tenantId consts.TenantId, projectId uint) (err error) {
	execRuntime, _ := jslibHelper.GetGojaRuntime(tenantId, projectId)
	_, err = execRuntime.RunString(fmt.Sprintf("getRespValueFromGoja('%s', dt.response);", execUuid))
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
func defineGoFuncs(tenantId consts.TenantId, projectId uint) {
	execRuntime, _ := jslibHelper.GetGojaRuntime(tenantId, projectId)

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

func jsErrMsg(msg string, category string, success bool) (ret string) {
	mp := iris.Map{
		"isCustomObj": true,
		"success":     success,
		"category":    category,
		"msg":         msg,
	}

	bytes, err := json.Marshal(mp)

	if err != nil {
		return err.Error()
	}

	ret = string(bytes)

	return
}
