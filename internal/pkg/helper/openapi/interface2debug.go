package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/jinzhu/copier"
	"strings"
)

type interfaces2debug struct {
	Inter    model.EndpointInterface
	Endpoint model.Endpoint
	Serve    model.Serve
}

func NewInterfaces2debug(inter model.EndpointInterface, endpoint model.Endpoint, serve model.Serve) *interfaces2debug {
	return &interfaces2debug{Inter: inter, Endpoint: endpoint, Serve: serve}
}

func (i *interfaces2debug) Convert() (debugInterface *model.DebugInterface) {
	debugInterface = new(model.DebugInterface)

	copier.CopyWithOption(debugInterface, &i.Inter, copier.Option{DeepCopy: true})
	/*
		for _, param := range i.Endpoint.PathParams {
			debugInterfaceParam := model.DebugInterfaceParam{
				InterfaceParamBase: model.InterfaceParamBase{
					Name:  param.Name,
					Value: param.Value,
				},
			}
			debugInterface.PathParams = append(debugInterface.PathParams, debugInterfaceParam)
		}
	*/
	debugInterface.ID = 0
	debugInterface.BodyFormData = i.BodyFormData()
	debugInterface.BodyFormUrlencoded = i.BodyFormUrlencoded()
	debugInterface.Body = i.Body()
	debugInterface.BodyType = i.BodyType()
	debugInterface.AuthorizationType, debugInterface.ApiKey, debugInterface.OAuth20, debugInterface.BearerToken, debugInterface.BasicAuth = i.security()
	debugInterface.QueryParams, debugInterface.PathParams, debugInterface.Headers, debugInterface.Cookies = i.params()

	return
}

func (i *interfaces2debug) BodyFormData() (bodyFormData []model.DebugInterfaceBodyFormDataItem) {
	return
}

func (i *interfaces2debug) BodyFormUrlencoded() (bodyFormUrlencoded []model.DebugInterfaceBodyFormUrlEncodedItem) {
	return
}

func (i *interfaces2debug) Body() (body string) {
	var examples []map[string]string
	_commUtils.JsonDecode(i.Inter.RequestBody.Examples, &examples)
	for _, example := range examples {
		return strings.ReplaceAll(example["content"], "\r\n", "")
	}
	return
}

func (i *interfaces2debug) BodyType() (mediaType consts.HttpContentType) {
	if i.Inter.RequestBody.MediaType != "" {
		mediaType = consts.HttpContentType(i.Inter.RequestBody.MediaType)
	}
	return
}

func (i *interfaces2debug) params() (queryParams []model.DebugInterfaceParam, pathParams []model.DebugInterfaceParam, headers []model.DebugInterfaceHeader, cookies []model.DebugInterfaceCookie) {
	for _, item := range i.Inter.Params {
		queryParams = append(queryParams, model.DebugInterfaceParam{InterfaceParamBase: model.InterfaceParamBase{Name: item.Name, ParamIn: consts.ParamInQuery, Value: item.Default}})
	}
	for _, item := range i.Endpoint.PathParams {
		pathParams = append(pathParams, model.DebugInterfaceParam{InterfaceParamBase: model.InterfaceParamBase{Name: item.Name, ParamIn: consts.ParamInPath, Value: item.Default}})
	}
	for _, item := range i.Inter.Headers {
		headers = append(headers, model.DebugInterfaceHeader{InterfaceHeaderBase: model.InterfaceHeaderBase{Name: item.Name, Value: item.Default}})
	}
	for _, item := range i.Inter.Cookies {
		cookies = append(cookies, model.DebugInterfaceCookie{InterfaceCookieBase: model.InterfaceCookieBase{Name: item.Name, Value: item.Default}})
	}
	return
}

func (i *interfaces2debug) security() (authorizationType string, apiKey model.DebugInterfaceApiKey, oAuth20 model.DebugInterfaceOAuth20, bearerToken model.DebugInterfaceBearerToken, basicAuth model.DebugInterfaceBasicAuth) {
	security := i.Inter.Security
	var securityInfo model.ComponentSchemaSecurity
	for _, item := range i.Serve.Securities {
		if security == "" && item.Default {
			security = item.Name
		}
		if security == item.Name {
			securityInfo = item
		}
	}

	authorizationType = securityInfo.Type
	switch securityInfo.Type {
	case "apiKey":
		apiKey = i.apiKey(securityInfo)
	case "bearerToken":
		bearerToken = i.bearerToken(securityInfo)
	case "basicAuth":
		basicAuth = i.basicAuth(securityInfo)
	}

	return
}

func (i *interfaces2debug) apiKey(securityInfo model.ComponentSchemaSecurity) (apiKey model.DebugInterfaceApiKey) {
	apiKey = model.DebugInterfaceApiKey{}
	apiKey.Key = securityInfo.Key
	apiKey.Value = securityInfo.Value
	return
}

func (i *interfaces2debug) oAuth20(securityInfo model.ComponentSchemaSecurity) (oAuth20 model.DebugInterfaceOAuth20) {
	return
}
func (i *interfaces2debug) bearerToken(securityInfo model.ComponentSchemaSecurity) (bearerToken model.DebugInterfaceBearerToken) {
	bearerToken = model.DebugInterfaceBearerToken{}
	bearerToken.Token = securityInfo.Token
	return
}
func (i *interfaces2debug) basicAuth(securityInfo model.ComponentSchemaSecurity) (basicAuth model.DebugInterfaceBasicAuth) {
	basicAuth = model.DebugInterfaceBasicAuth{}
	basicAuth.Username = securityInfo.Username
	basicAuth.Password = securityInfo.Password
	return
}
