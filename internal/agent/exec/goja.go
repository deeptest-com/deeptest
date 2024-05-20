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
	"github.com/dop251/goja_nodejs/require"
	"path/filepath"
	"reflect"
)

func InitGojaRuntimeWithSession(session *ExecSession, vuNo int, tenantId consts.TenantId) {
	session.GojaRuntime, session.GojaRequire = GenerateGojaRuntime()

	jslibHelper.LoadChaiJslibs(session.GojaRuntime)

	defineJsFuncs(session, false)

	loadDeeptestScript(session)

	defineGoFuncs(session)

	// import other custom libs
	jslibHelper.RefreshRemoteAgentJslibs(session.GojaRuntime, session.GojaRequire,
		vuNo, tenantId, session.ProjectId,
		session.ServerUrl, session.ServerToken)

	return
}

func loadDeeptestScript(session *ExecSession) (err error) {
	script := scriptHelper.GetScript(scriptHelper.ScriptDeepTest)

	pth := filepath.Join(consts.TmpDir, fmt.Sprintf("deeptest-%s-%d.js", session.ExecUuid, session.VuNo))
	fileUtils.RmDir(pth)
	fileUtils.WriteFile(pth, script)

	dt, err := session.GojaRequire.Require(pth)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	err = session.GojaRuntime.Set("dt", dt)
	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	return
}

func defineJsFuncs(session *ExecSession, isSimple bool) (err error) {
	/* START: called by js */
	err = session.GojaRuntime.Set("getDatapoolVariable", func(dpName, field, seq string) (ret interface{}) {
		rowIndex := getDatapoolRow(dpName, seq, session.ExecScene.Datapools, session.ScenarioDebug.DatapoolCursor)

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

	if isSimple {
		return
	}

	err = session.GojaRuntime.Set("getVariable", func(name string) interface{} {
		var scopeId uint
		if session.ScenarioDebug.CurrProcessor != nil {
			scopeId = session.ScenarioDebug.CurrProcessor.ParentId
		}

		vari, err := GetVariable(name, scopeId, session)
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
	err = session.GojaRuntime.Set("setVariable", func(name string, val interface{}) {
		var scopeId uint
		if session.ScenarioDebug.CurrProcessor != nil {
			scopeId = session.ScenarioDebug.CurrProcessor.ParentId
		}

		ret, err := SetVariable(scopeId, name, val, commUtils.ValueType(val), consts.Public, session)

		if err == nil {
			session.AppendGojaVariables(ret)
		} else {
			session.AppendGojaLog(jsErrMsg(err.Error(), "setVariable", false))
		}

		return
	})
	err = session.GojaRuntime.Set("clearVariable", func(name string) {
		var scopeId uint
		if session.ScenarioDebug.CurrProcessor != nil {
			scopeId = session.ScenarioDebug.CurrProcessor.ParentId
		}

		err := ClearVariable(scopeId, name, session)
		if err != nil {
			session.AppendGojaLog(jsErrMsg(err.Error(), "clearVariable", false))
		}
	})

	// http request
	err = session.GojaRuntime.Set("sendRequest", func(data goja.Value, cb func(interface{}, interface{})) {
		req := gojaUtils.GenRequest(data, session.GojaRuntime)

		errOfCallbackParam := ""

		resp, err2 := Invoke(&req)
		if err2 != nil {
			// AppendGojaLog(execUuid, jsErrMsg(err2.Error(), "sendRequest", false))
			errOfCallbackParam = jsErrMsg(err2.Error(), "sendRequest", false)
		}

		cb(errOfCallbackParam, resp)
	})

	// log
	err = session.GojaRuntime.Set("log", func(value interface{}) {
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
	err = session.GojaRuntime.Set("getReqValueFromGoja", func(value domain.BaseRequest) {
		session.SetCurrRequest(value)
	})
	err = session.GojaRuntime.Set("getRespValueFromGoja", func(value domain.DebugResponse) {
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

func defineGoFuncs(session *ExecSession) {
	// set data
	script := `function _setData(name, val) {
					dt[name] = val
				}`
	_, err := session.GojaRuntime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	err = session.GojaRuntime.ExportTo(session.GojaRuntime.Get("_setData"), &session.GojaSetValueFunc)
}

func GenerateGojaRuntime() (execRuntime *goja.Runtime, execRequire *require.RequireModule) {
	execRuntime = goja.New()
	execRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用
	execRequire = registry.Enable(execRuntime)

	return
}
