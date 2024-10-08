package handler

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	service2 "github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type AuthCtrl struct {
	AuthService      *service2.AuthService      `inject:""`
	WebSocketService *service2.WebSocketService `inject:""`
	BaseCtrl
}

// OAuth2Authorization
// @Tags	授权模块
// @summary	生成OAuth认证信息
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization			header	string						true	"Authentication header"
// @Param 	currProjectId			query	int							true	"当前项目ID"
// @Param 	DebugInterfaceOAuth20	body 	model.DebugInterfaceOAuth20 true 	"生成OAuth认证信息的请求体"
// @success	200	{object}	_domain.Response{data=object{url=string}}
// @Router	/api/v1/auth/oauth2Authorization	[post]
func (c *AuthCtrl) OAuth2Authorization(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := model.DebugInterfaceOAuth20{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.AuthService.OAuth2Authorization(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// GetOAuth2AccessToken
// @Tags	授权模块
// @summary	调用认证服务生成访问令牌
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	accessTokenURL	query 	string 	true 	"accessTokenURL"
// @Param 	clientId		query 	string 	true 	"clientId"
// @Param 	clientSecret	query 	string 	true 	"clientSecret"
// @Param 	code			query 	string 	true 	"code"
// @success	200	{object}	_domain.Response{data=iris.Map{}}
// @Router	/api/v1/auth/getOAuth2AccessToken	[post]
func (c *AuthCtrl) GetOAuth2AccessToken(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	accessTokenURL := ctx.URLParam("accessTokenURL")
	clientId := ctx.URLParam("clientId")
	clientSecret := ctx.URLParam("clientSecret")
	code := ctx.URLParam("code")

	data, err := c.AuthService.GenOAuth2AccessToken(tenantId, accessTokenURL, clientId, clientSecret, code)

	c.WebSocketService.SendMsg(
		consts.WsDefaultNamespace,
		consts.WsDefaultRoom,
		data)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// UseOAuth2AccessToken
// @Tags	授权模块
// @summary	加载访问令牌到接口
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	interfaceId		query 	int 	true 	"接口ID"
// @Param 	name			query 	string 	true 	"name"
// @Param 	token			query 	string 	true 	"token"
// @Param 	tokenType		query 	string 	true 	"tokenType"
// @success	200	{object}	_domain.Response{data=object{token=string,tokenType=string}}
// @Router	/api/v1/auth/useOAuth2AccessToken	[post]
func (c *AuthCtrl) UseOAuth2AccessToken(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	currProjectId, _ := ctx.URLParamInt("currProjectId")
	interfaceId, _ := ctx.URLParamInt("interfaceId")
	name := ctx.URLParam("name")
	token := ctx.URLParam("token")
	tokenType := ctx.URLParam("tokenType")

	err := c.AuthService.AddToken(tenantId, name, token, tokenType, interfaceId, currProjectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data := iris.Map{
		"token":     token,
		"tokenType": tokenType,
	}
	c.WebSocketService.SendMsg(
		consts.WsDefaultNamespace,
		consts.WsDefaultRoom,
		data)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// ListOAuth2Token
// @Tags	授权模块
// @summary	加载访问令牌到接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=[]model.Auth2Token}
// @Router	/api/v1/auth/listOAuth2Token	[get]
func (c *AuthCtrl) ListOAuth2Token(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")

	pos, err := c.AuthService.ListOAuth2Token(tenantId, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: pos})
}

// RemoveToken
// @Tags	授权模块
// @summary	删除访问令牌
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"token id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/auth/removeToken	[get]
func (c *AuthCtrl) RemoveToken(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	tokenId, _ := ctx.URLParamInt("id")

	err := c.AuthService.RemoveToken(tenantId, tokenId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}
