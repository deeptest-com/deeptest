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
	"path/filepath"
	"sync"
	"time"
)

var (
	AgentLoadedLibs sync.Map
)

func LoadChaiJslibs(runtime *goja.Runtime) {
	// check method
	err := runtime.Set("check", func(ok bool, name, msg string) {
		log.Println(fmt.Sprintf("%t, %s, %s", ok, name, msg)) // TODO: add to assert
	})

	// test method
	script := `function test(name, cb) {
		try {
			cb();
		} catch(err){
			log('TestCase \'' + name + '\' failed, ' + err + '.')
			check(false, name, err)
			return
		}

		log('TestCase \'' + name + '\' success.')
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

func RefreshRemoteAgentJslibs(runtime *goja.Runtime, require *require.RequireModule, projectId uint, serverUrl, token string) {
	libs := getJslibsFromServer(projectId, serverUrl, token)

	for _, lib := range libs {
		id := lib.Id

		updateTime, ok := GetAgentCache(id)

		if !ok || updateTime.Before(lib.UpdatedAt) {
			pth := filepath.Join(consts.TmpDir, fmt.Sprintf("%d.js", id))
			fileUtils.WriteFile(pth, lib.Script)
			module, err := require.Require(pth)
			if err != nil {
				logUtils.Info(err.Error())
			}

			runtime.Set(lib.Name, module)

			SetAgentCache(id, lib.UpdatedAt)
		}
	}
}

func GetAgentCache(id uint) (val time.Time, ok bool) {
	inf, ok := AgentLoadedLibs.Load(id)

	if ok {
		val = inf.(time.Time)
	}

	return
}

func SetAgentCache(id uint, val time.Time) {
	AgentLoadedLibs.Store(id, val)
}

func getJslibsFromServer(projectId uint, serverUrl, token string) (libs []Jslib) {
	url := fmt.Sprintf("snippets/getJslibsForAgent?projectId=%d", projectId)

	loadedLibs := map[uint]time.Time{}
	AgentLoadedLibs.Range(func(key, value interface{}) bool {
		loadedLibs[key.(uint)] = value.(time.Time)
		return true
	})
	body, err := json.Marshal(loadedLibs)

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + url,
		BodyType:          consts.ContentTypeJSON,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: token,
		},

		Body: string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get Jslibs failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
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
