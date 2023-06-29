package handler

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"os"
)

type ConfigCtrl struct {
	BaseCtrl
}

func (c *ConfigCtrl) Get(ctx iris.Context) {
	data := iris.Map{
		"demoTestSite": os.Getenv("DEMO_TEST_SITE"),
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}
