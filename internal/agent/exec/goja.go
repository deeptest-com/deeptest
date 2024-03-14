package agentExec

import (
	"encoding/json"
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
	"path/filepath"
	"reflect"
)

func InitGojaRuntime(session *ExecSession) {
	session.GojaRuntime, session.GojaRequire = gojaUtils.InitGojaRuntime()

	jslibHelper.LoadChaiJslibs(session.GojaRuntime)

	defineJsFuncs(session)
	defineGoFuncs(session)

	// load global script
	pth := filepath.Join(consts.TmpDir, "deeptest.js")
	fileUtils.WriteFile(pth, scriptHelper.GetScript(scriptHelper.ScriptDeepTest))
	dt, err := session.GojaRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	session.GojaRuntime.Set("dt", dt)

	// import other custom libs
	jslibHelper.RefreshRemoteAgentJslibs(session.GojaRuntime, session.GojaRequire, session.ProjectId, session.ServerUrl, session.ServerToken)

	return
}

func defineJsFuncs(session *ExecSession) (err error) {
	execRuntime := session.GojaRuntime

	/* START: called by js */
	err = execRuntime.Set("getDatapoolVariable", func(dpName, field, seq string) (ret interface{}) {
		rowIndex := getDatapoolRow(dpName, seq, session.ExecScene.Datapools, session.DatapoolCursor)

		if session.ExecScene.Datapools[dpName] == nil {
			ret = consts.INVALID_VALUE
			session.AppendGojaLog(jsErrMsg("DATAPOOL_NOT_FOUND:"+dpName, "getDatapoolVariable", false))
			return
		}

		if rowIndex > len(session.ExecScene.Datapools[dpName])-1 {
			ret = consts.INVALID_VALUE
			session.AppendGojaLog(jsErrMsg("DATAPOOL_INDEX_OUT_OF_RANGE:"+dpName, "getDatapoolVariable", false))
			return
		}

		ret = session.ExecScene.Datapools[dpName][rowIndex][field]
		if ret == nil {
			ret = consts.INVALID_VALUE

			session.AppendGojaLog(jsErrMsg("DATAPOOL_VARIABLE_NOT_FOUND:"+field+"@"+dpName, "getDatapoolVariable", false))
			return
		}

		return
	})

	err = execRuntime.Set("getVariable", func(name string) interface{} {
		var scopeId uint
		if session.CurrScenarioProcessor != nil {
			scopeId = session.CurrScenarioProcessor.ParentId
		}

		vari, err := GetVariable(session, scopeId, name)
		if err != nil {
			vari.Value = consts.INVALID_VALUE

			session.AppendGojaLog(jsErrMsg(err.Error(), "getVariable", false))

			return vari.Value
		}

		vari.Value, err = commUtils.ConvertValueForUse(vari.Value, vari.ValueType)
		if err != nil {
			vari.Value = consts.INVALID_VALUE
			session.AppendGojaLog(jsErrMsg(err.Error(), "getVariable", false))

			return vari.Value
		}

		return vari.Value
	})
	err = execRuntime.Set("setVariable", func(name string, val interface{}) {
		var scopeId uint
		if session.CurrScenarioProcessor != nil {
			scopeId = session.CurrScenarioProcessor.ParentId
		}

		ret, err := SetVariable(session, scopeId, name, val, commUtils.ValueType(val), consts.Public)

		if err == nil {
			session.AppendGojaVariables(ret)
		} else {
			session.AppendGojaLog(jsErrMsg(err.Error(), "setVariable", false))
		}

		return
	})
	err = execRuntime.Set("clearVariable", func(name string) {
		var scopeId uint
		if session.CurrScenarioProcessor != nil {
			scopeId = session.CurrScenarioProcessor.ParentId
		}

		err := ClearVariable(session, scopeId, name)
		if err != nil {
			session.AppendGojaLog(jsErrMsg(err.Error(), "clearVariable", false))
		}
	})

	// http request
	err = execRuntime.Set("sendRequest", func(data goja.Value, cb func(interface{}, interface{})) {
		req := gojaUtils.GenRequest(data, execRuntime)

		errOfCallbackParam := ""

		resp, err2 := Invoke(&req)
		if err2 != nil {
			// AppendGojaLog(execUuid, jsErrMsg(err2.Error(), "sendRequest", false))
			errOfCallbackParam = jsErrMsg(err2.Error(), "sendRequest", false)
		}

		cb(errOfCallbackParam, resp)
	})

	// log
	err = execRuntime.Set("log", func(value interface{}) {
		if value == nil {
			session.AppendGojaLog("空")
			return
		}

		typ := reflect.TypeOf(value).Name()

		if typ == "string" {
			session.AppendGojaLog(value.(string))
		} else {
			bytes, _ := json.Marshal(value)
			session.AppendGojaLog(string(bytes))
		}
	})
	/* END: called by js */

	/* START: called by go */
	err = execRuntime.Set("getReqValueFromGoja", func(value domain.BaseRequest) {
		session.SetCurrRequest(value)
	})
	err = execRuntime.Set("getRespValueFromGoja", func(value domain.DebugResponse) {
		if httpHelper.IsJsonResp(value) {
			bytes, _ := json.Marshal(value.Data)
			value.Content = string(bytes)
			session.SetCurrResponse(value)
		} else {
			var ok bool
			if value.Content, ok = value.Data.(string); ok {

			}
			session.SetCurrResponse(value)
		}
	})
	/* END: called by go */

	return
}

// we can call go SetValueToGoja as call js _setData
var (
	_setValueFunc func(name string, value interface{})
)

func SetValueToGoja(name string, value interface{}) {
	_setValueFunc(name, value)
}

func defineGoFuncs(session *ExecSession) {
	// set data
	script := `function _setData(name, val) {
					dt[name] = val
				}`
	_, err := session.GojaRuntime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	err = session.GojaRuntime.ExportTo(session.GojaRuntime.Get("_setData"), &_setValueFunc)
}
