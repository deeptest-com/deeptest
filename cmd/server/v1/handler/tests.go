package handler

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"log"
)

type TestsCtrl struct {
	BaseCtrl
}

func (c *TestsCtrl) Test(ctx iris.Context) {
	param1 := ctx.URLParams()

	log.Println(param1)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
