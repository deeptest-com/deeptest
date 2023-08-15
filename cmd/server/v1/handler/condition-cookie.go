package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type CookieCtrl struct {
	CookieService *service.CookieService `inject:""`
	BaseCtrl
}

// Get 详情
// @Tags	Cookie
// @summary	Cookie详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"CookieID"
// @success	200	{object}	_domain.Response{data=model.DebugConditionCookie}
// @Router	/api/v1/extractors/{id}	[get]
func (c *CookieCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	extractor, err := c.CookieService.Get(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: extractor})
}

// Update 更新
// @Tags	Cookie
// @summary	更新Cookie
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string							true	"Authentication header"
// @Param 	currProjectId			query	int								true	"当前项目ID"
// @Param 	DebugConditionCookie	body	model.DebugConditionCookie	true	"更新Cookie的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/extractors	[put]
func (c *CookieCtrl) Update(ctx iris.Context) {
	var extractor model.DebugConditionCookie
	err := ctx.ReadJSON(&extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.CookieService.Update(&extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
// @Tags	Cookie
// @summary	删除Cookie
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"CookieID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/extractors/{id}	[delete]
func (c *CookieCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.CookieService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
