package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/jinzhu/copier"
	"strings"
)

type interfaces2debug struct {
	Inter model.EndpointInterface
	Serve model.Serve
}

func NewInterfaces2debug(inter model.EndpointInterface, serve model.Serve) *interfaces2debug {
	return &interfaces2debug{Inter: inter, Serve: serve}
}

func (i *interfaces2debug) Convert() (debugInterface *model.DebugInterface) {
	debugInterface = new(model.DebugInterface)

	copier.CopyWithOption(debugInterface, &i.Inter, copier.Option{DeepCopy: true})

	debugInterface.ID = 0
	debugInterface.BodyFormData = i.BodyFormData()
	debugInterface.BodyFormUrlencoded = i.BodyFormUrlencoded()
	debugInterface.Body = i.Body()
	debugInterface.BodyType = i.BodyType()
	debugInterface.AuthorizationType, debugInterface.ApiKey, debugInterface.OAuth20, debugInterface.BearerToken, debugInterface.BasicAuth = i.security()

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

func (i *interfaces2debug) params() {

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
