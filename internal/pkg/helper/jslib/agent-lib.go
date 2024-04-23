package jslibHelper

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	gojaPlugin "github.com/aaronchen2k/deeptest/internal/pkg/goja/plugin"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"log"
	"time"
)

func LoadChaiJslibs(runtime *goja.Runtime) {
	// check method
	err := runtime.Set("check", func(ok bool, name, msg string) {
		log.Println(fmt.Sprintf("%t, %s, %s", ok, name, msg))
	})

	// test method
	script := `function test(name, cb) {
		try {
			cb();
		} catch(err){
			log('Assertion Failed [' + name + '] ' + err + '.')
			check(false, name, err)
			return
		}

		log('Assertion Pass [' + name + '].')
		check(true, name, '')
	}`
	_, err = runtime.RunString(script)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	// chai method
	runtime.Set("require", func(call goja.FunctionCall) goja.Value { return goja.Undefined() })
	runtime.Set("global", runtime.GlobalObject())

	agentVu := gojaPlugin.AgentVU{
		RuntimeField: runtime,
	}
	chaiModule := gojaPlugin.NewChai()
	chaiInst := chaiModule.NewModuleInstance(&agentVu)
	runtime.Set("expect", chaiInst.Exports().Named["expect"])
}

func RefreshRemoteAgentJslibs(runtime *goja.Runtime, require *require.RequireModule, tenantId consts.TenantId, projectId uint, serverUrl, token string) {
	libs := getJslibsFromServer(tenantId, projectId, serverUrl, token)

	for _, lib := range libs {
		id := lib.Id
		if tenantId == "" {
			tenantId = "NA"
		}

		jsFileName := fmt.Sprintf("%d-%s-%d.js", id, tenantId, lib.UpdatedAt.Unix())

		tmpPath := fmt.Sprintf("%s/%s", consts.TmpDirRelativeAgent, jsFileName)
		tmpContent := lib.Script
		fileUtils.WriteFileIfNotExist(tmpPath, tmpContent)

		module, err := require.Require("./" + tmpPath)
		if err != nil {
			logUtils.Infof("goja require failed, path: %s, err: %s.", tmpPath, err.Error())
		}

		err = runtime.Set(lib.Name, module)
		if err != nil {
			logUtils.Errorf(err.Error())
		}

		//SetAgentCache(projectId, id, lib.UpdatedAt)
		logUtils.Infof("更新第三方库，projectId：%v,id:%v,lib.Name:%v", projectId, id, lib.Name)
		//}
	}
}

func GetAgentCache(tenantId consts.TenantId, projectId uint, id uint) (val time.Time, ok bool) {
	mp := GetAgentLoadedLibs(tenantId, projectId)
	val, ok = (*mp)[id]

	return
}

func SetAgentCache(tenantId consts.TenantId, projectId uint, id uint, val time.Time) {
	mp := GetAgentLoadedLibs(tenantId, projectId)
	(*mp)[id] = val
}

func getJslibsFromServer(tenantId consts.TenantId, projectId uint, serverUrl, token string) (libs []Jslib) {
	url := fmt.Sprintf("snippets/getJslibsForAgent?projectId=%d", projectId)

	loadedLibs := &map[uint]time.Time{} // GetAgentLoadedLibs(projectId)

	body, err := json.Marshal(loadedLibs)

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + url,
		BodyType:          consts.ContentTypeJSON,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: token,
		},

		Body:    string(body),
		Headers: &[]domain.Header{{Name: "TenantId", Value: string(tenantId)}},
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get Jslibs failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get Jslibs obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 0 {
		logUtils.Infof("get interface obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get interface obj failed, response %v", resp.Content)
		return
	}

	json.Unmarshal(bytes, &libs)

	return
}
