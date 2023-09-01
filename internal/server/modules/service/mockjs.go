package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"path/filepath"
)

var (
	MockJsVm      JsVm
	MockJsRequire *require.RequireModule

	MockFunc func(p interface{}) interface{}
)

type JsVm struct {
	JsRuntime *goja.Runtime
}

type MockJsService struct {
	MockJsRepo *repo.MockJsRepo `inject:""`
}

func (s *MockJsService) ListExpressions() (pos []serverDomain.MockJsExpression, err error) {
	pos, err = s.MockJsRepo.ListExpressions()

	return
}

func (s *MockJsService) EvaluateExpression(req serverDomain.MockJsExpression) (ret serverDomain.MockJsExpression, err error) {
	ret = req
	if MockJsVm.JsRuntime == nil {
		InitJsRuntime()
	}

	if req.Expression == "" {
		return
	}

	ret.Result = MockFunc(req.Expression)

	return
}

func InitJsRuntime() {
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用
	MockJsVm.JsRuntime = goja.New()

	MockJsRequire = registry.Enable(MockJsVm.JsRuntime)

	module := "mock.js"
	pth := filepath.Join(consts.WorkDir, module)
	fileUtils.WriteFile(pth, scriptHelper.GetModule(module))
	mock, err := MockJsRequire.Require(pth)

	MockJsVm.JsRuntime.Set("mock", mock)

	script := `function Mock(str) {
					let param = str
					if (str.indexOf('@') !== 0) {
						param = JSON.parse(str);
					}

					var data = mock.mock(param)
					return data;
				}`
	_, err = MockJsVm.JsRuntime.RunString(script)
	if err != nil {
		_logUtils.Infof(err.Error())
	}

	err = MockJsVm.JsRuntime.ExportTo(MockJsVm.JsRuntime.Get("Mock"), &MockFunc)

	if err != nil {
		_logUtils.Infof(err.Error())
	}
}
