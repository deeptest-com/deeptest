package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
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

		// get logs from js script execution
		logs = *session.GetGojaLogs()

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
	execRuntime := session.GojaRuntime

	_, err = execRuntime.RunString(fmt.Sprintf("getReqValueFromGoja('%s', dt.request);", session.ExecUuid))

	return
}
func GetRespValueFromGoja(session *ExecSession) (err error) {
	execRuntime := session.GojaRuntime

	_, err = execRuntime.RunString(fmt.Sprintf("getRespValueFromGoja('%s', dt.response);", session.ExecUuid))
	return
}

func SetReqValueToGoja(req *domain.BaseRequest, session *ExecSession) {
	session.GojaSetValueFunc("request", req)
}
func SetRespValueToGoja(resp *domain.DebugResponse, session *ExecSession) {
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
