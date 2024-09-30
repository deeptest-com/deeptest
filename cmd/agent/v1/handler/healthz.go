package handler

import (
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type HealthzCtrl struct {
}

func (c *HealthzCtrl) Get(ctx iris.Context) {
	ctx.JSON(_domain.Response{Code: 200, Msg: "health"})
}
