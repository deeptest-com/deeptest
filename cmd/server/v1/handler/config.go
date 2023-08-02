package handler

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"os"
)

type ConfigCtrl struct {
	BaseCtrl
}

//Get
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
