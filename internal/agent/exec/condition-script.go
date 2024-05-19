package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/kataras/iris/v12"
)

func ExecScript(scriptObj *domain.ScriptBase, session *ExecSession) (err error) {
	execRuntime := session.GojaRuntime

	session.ResetGojaVariables()
	session.ResetGojaLogs()

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

		logs = *session.GojaLogs

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

//func InitJsRuntime(tenantId consts.TenantId, projectId uint, execUuid string) {
//	jslibHelper.InitProjectGojaRuntime(tenantId, projectId)
//	execRuntime, execRequire := jslibHelper.GetProjectGojaRuntime(tenantId, projectId)
//
//	jslibHelper.LoadChaiJslibs(execRuntime)
//
//	defineJsFuncs(execUuid, tenantId, projectId)
//	defineGoFuncs(tenantId, projectId)
//
//	// load global script
//	tmpPath := fmt.Sprintf("%s/deeptest.js", consts.TmpDirRelativeAgent)
//	tmpContent := scriptHelper.GetScript(scriptHelper.ScriptDeepTest)
//	fileUtils.WriteFileIfNotExist(tmpPath, tmpContent)
//
//	dt, err := execRequire.Require("./" + tmpPath)
//	if err != nil {
//		logUtils.Infof("goja require failed, path: %s, err: %s.", tmpPath, err.Error())
//	}
//
//	execRuntime.Set("dt", dt)
//
//	// import other custom libs
//	jslibHelper.RefreshRemoteAgentJslibs(execRuntime, execRequire, tenantId, projectId, GetServerUrl(execUuid), GetServerToken(execUuid))
//}

func GetReqValueFromGoja(execUuid string, tenantId consts.TenantId, projectId uint) (err error) {
	execRuntime, _ := jslibHelper.GetProjectGojaRuntime(tenantId, projectId)
	_, err = execRuntime.RunString(fmt.Sprintf("getReqValueFromGoja('%s', dt.request);", execUuid))
	return
}
func GetRespValueFromGoja(execUuid string, tenantId consts.TenantId, projectId uint) (err error) {
	execRuntime, _ := jslibHelper.GetProjectGojaRuntime(tenantId, projectId)
	_, err = execRuntime.RunString(fmt.Sprintf("getRespValueFromGoja('%s', dt.response);", execUuid))
	return
}

func SetReqValueToGoja(req *domain.BaseRequest, session *ExecSession) {
	session.GojaSetValueFunc("request", req)
}
func SetRespValueToGoja(resp *domain.DebugResponse) {
	// set resp.Data to json object for goja edit
	if httpHelper.IsJsonResp(*resp) {
		var data interface{}
		err := _commUtils.JsonDecode(resp.Content, &data)
		//err := json.Unmarshal([]byte(resp.Content), &data)
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
