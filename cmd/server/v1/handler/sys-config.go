package handler

import (
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"os"
)

type ConfigCtrl struct {
	BaseCtrl
	ConfigService *service.ConfigService `inject:""`
}

const token = "a1bc**2d&&423qvdw"

// Get
// @Tags	配置
// @summary	获取服务端配置
// @accept 	application/json
// @Produce application/json
// @success	200	{object}	_domain.Response{data=object{demoTestSite=string}}
// @Router	/api/v1/configs	[get]
func (c *ConfigCtrl) Get(ctx iris.Context) {
	data := iris.Map{
		"demoTestSite": os.Getenv("DEMO_TEST_SITE"),
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *ConfigCtrl) GetValue(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	key := ctx.URLParam("key")
	if key == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	value, err := c.ConfigService.Get(tenantId, key)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: value})
}

func (c *ConfigCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	headerToken := ctx.Request().Header.Get("token")
	if headerToken != token {
		ctx.JSON(_domain.Response{Code: _domain.AuthActionErr.Code, Msg: _domain.AuthActionErr.Msg})
		return
	}
	req := model.SysConfig{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ConfigService.Save(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
