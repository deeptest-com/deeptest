package handler

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type DocumentCtrl struct {
	DocumentService service.DocumentService
}

func (c *DocumentCtrl) Index(ctx iris.Context) {
	var req domain.DocumentReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, _ := c.DocumentService.Content(req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: new(domain.DocumentRep)})
}
