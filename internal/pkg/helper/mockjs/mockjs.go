package mockjsHelper

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

var (
	mockJsVm      JsVm
	mockJsRequire *require.RequireModule

	mockFunc func(p interface{}) interface{}
)

type JsVm struct {
	JsRuntime *goja.Runtime
}

func EvaluateExpression(req serverDomain.MockJsExpression) (ret serverDomain.MockJsExpression, err error) {
	ret = req
	if mockJsVm.JsRuntime == nil {
		initJsRuntime()
	}

	if req.Expression == "" {
		return
	}

	ret.Result = mockFunc(req.Expression)

	return
}

func initJsRuntime() {
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用
	mockJsVm.JsRuntime = goja.New()
	mockJsVm.JsRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	mockJsRequire = registry.Enable(mockJsVm.JsRuntime)

	moduleName := scriptHelper.ModuleMockJs
	tmpPath := fmt.Sprintf("%s/%s", consts.TmpDirRelatedServer, "tmp/"+moduleName)
	fileUtils.WriteFileIfNotExist(tmpPath, scriptHelper.GetModule(moduleName))

	mock, err := mockJsRequire.Require("./" + tmpPath)
	if err != nil {
		_logUtils.Infof("goja require failed, path: %s, err: %s.", tmpPath, err.Error())
	}

	mockJsVm.JsRuntime.Set("mock", mock)

	script := `function Mock(str) {
					let param = str
					if (str.indexOf('@') !== 0) {
						param = JSON.parse(str);
					}

					var data = mock.mock(param)
					return data;
				}`
	_, err = mockJsVm.JsRuntime.RunString(script)
	if err != nil {
		_logUtils.Infof(err.Error())
	}

	err = mockJsVm.JsRuntime.ExportTo(mockJsVm.JsRuntime.Get("Mock"), &mockFunc)

	if err != nil {
		_logUtils.Infof(err.Error())
	}
}
