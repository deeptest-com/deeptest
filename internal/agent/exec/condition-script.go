package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	"github.com/kataras/iris/v12"
)

func ExecScript(session *ExecSession, scriptObj *domain.ScriptBase, projectId uint) (err error) {
	execRuntime := session.GojaRuntime

	gojaVariables := make([]domain.ExecVariable, 0)
	session.GojaVariables = &gojaVariables

	gojaLogs := make([]string, 0)
	session.GojaLogs = &gojaLogs

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

func GetReqValueFromGoja(session *ExecSession) (err error) {
	_, err = session.GojaRuntime.RunString("getReqValueFromGoja(dt.request);")
	return
}
func GetRespValueFromGoja(session *ExecSession) (err error) {
	_, err = session.GojaRuntime.RunString("getRespValueFromGoja(dt.response);")
	return
}

func SetReqValueToGoja(session *ExecSession, req *domain.BaseRequest) {
	session.GojaSetValueFunc("request", req)
}
func SetRespValueToGoja(session *ExecSession, resp *domain.DebugResponse) {
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

	session.GojaSetValueFunc("response", resp)
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
