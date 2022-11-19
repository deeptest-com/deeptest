package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	service2 "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type AuthCtrl struct {
	AuthService      *service2.AuthService      `inject:""`
	WebSocketService *service2.WebSocketService `inject:""`
	BaseCtrl
}

// OAuth2Authorization
func (c *AuthCtrl) OAuth2Authorization(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	req := model.InterfaceOAuth20{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.AuthService.OAuth2Authorization(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// GetOAuth2AccessToken
func (c *AuthCtrl) GetOAuth2AccessToken(ctx iris.Context) {
	accessTokenURL := ctx.URLParam("accessTokenURL")
	clientId := ctx.URLParam("clientId")
	clientSecret := ctx.URLParam("clientSecret")
	code := ctx.URLParam("code")

	data, err := c.AuthService.GenOAuth2AccessToken(accessTokenURL, clientId, clientSecret, code)

	c.WebSocketService.SendMsg(
		consts.WsDefaultNameSpace,
		consts.WsDefaultRoom,
		data)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// UseOAuth2AccessToken
func (c *AuthCtrl) UseOAuth2AccessToken(ctx iris.Context) {
	currProjectId, _ := ctx.URLParamInt("currProjectId")
	interfaceId, _ := ctx.URLParamInt("interfaceId")
	name := ctx.URLParam("name")
	token := ctx.URLParam("token")
	tokenType := ctx.URLParam("tokenType")

	err := c.AuthService.AddToken(name, token, tokenType, interfaceId, currProjectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data := iris.Map{
		"token":     token,
		"tokenType": tokenType,
	}
	c.WebSocketService.SendMsg(
		consts.WsDefaultNameSpace,
		consts.WsDefaultRoom,
		data)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// ListOAuth2Token
func (c *AuthCtrl) ListOAuth2Token(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("currProjectId")

	pos, err := c.AuthService.ListOAuth2Token(projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: pos})
}

// RemoveToken
func (c *AuthCtrl) RemoveToken(ctx iris.Context) {
	tokenId, _ := ctx.URLParamInt("id")

	err := c.AuthService.RemoveToken(tokenId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}
