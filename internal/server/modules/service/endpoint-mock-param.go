package service

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12"
	"regexp"
	"strings"
)

type EndpointMockParamService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
}

func (s *EndpointMockParamService) GetRealRequestValues(tenantId consts.TenantId, ctx iris.Context,
	endpointInterface model.EndpointInterface, endpoint model.Endpoint) (
	headers []domain.Param, queryParams []domain.Param, pathParams []domain.Param,
	body string, bodyForm map[string][]string, cookies []domain.ExecCookie) {

	queryParams = s.getRealQueryParamValues(tenantId, ctx, endpointInterface)

	pathParams = s.getRealPathParamValues(tenantId, ctx, endpoint)

	headers = s.getRealHeaderParamValues(tenantId, ctx, endpointInterface)

	body, bodyForm = s.getRealBody(tenantId, ctx)

	cookies = s.getRealCookies(tenantId, ctx)

	return
}

func (s *EndpointMockParamService) getRealQueryParamValues(tenantId consts.TenantId, ctx iris.Context, endpointInterface model.EndpointInterface) (
	ret []domain.Param) {
	definedParams, _ := s.EndpointInterfaceRepo.ListParams(tenantId, endpointInterface.ID)

	definedParamTypeMap := map[string]string{}
	for _, definedParam := range definedParams {
		definedParamTypeMap[definedParam.Name] = definedParam.Type
	}

	realParams := ctx.URLParams()
	for key, realParam := range realParams {
		item := domain.Param{
			Name:  key,
			Type:  definedParamTypeMap[key],
			Value: realParam,
		}

		ret = append(ret, item)
	}

	return
}

func (s *EndpointMockParamService) getRealPathParamValues(tenantId consts.TenantId, ctx iris.Context, endpoint model.Endpoint) (
	ret []domain.Param) {
	definedParams, _ := s.EndpointRepo.GetEndpointPathParams(tenantId, endpoint.ID)

	definedParamTypeMap := map[string]string{}
	for _, definedParam := range definedParams {
		definedParamTypeMap[definedParam.Name] = definedParam.Type
	}

	mockPath := "/" + ctx.Params().Get("path")
	realParams, _ := s.MatchEndpointByMockPath(tenantId, mockPath, endpoint)
	for key, realParam := range realParams {
		item := domain.Param{
			Name:  key,
			Type:  definedParamTypeMap[key],
			Value: realParam,
		}

		ret = append(ret, item)
	}

	return
}

func (s *EndpointMockParamService) getRealHeaderParamValues(tenantId consts.TenantId, ctx iris.Context, endpointInterface model.EndpointInterface) (
	ret []domain.Param) {
	definedParams, _ := s.EndpointInterfaceRepo.ListHeaders(tenantId, endpointInterface.ID)

	definedParamTypeMap := map[string]string{}
	for _, definedParam := range definedParams {
		definedParamTypeMap[definedParam.Name] = definedParam.Type
	}

	realParams := ctx.Request().Header
	for key, realParam := range realParams {
		item := domain.Param{
			Name:  key,
			Type:  definedParamTypeMap[key],
			Value: strings.Join(realParam, ","),
		}

		ret = append(ret, item)
	}

	return
}

func (s *EndpointMockParamService) getRealBody(tenantId consts.TenantId, ctx iris.Context) (body string, bodyForm map[string][]string) {
	method := ctx.Method()
	if method != consts.POST.String() && method != consts.PUT.String() && method != consts.DELETE.String() {
		return
	}

	reqContentType := ctx.GetContentTypeRequested()

	if reqContentType == consts.ContentTypeJSON.String() {
		var req interface{}
		ctx.ReadJSON(&req)

		bodyBytes, _ := json.Marshal(req)
		body = string(bodyBytes)

	} else if reqContentType == consts.ContentTypeFormData.String() {
		bodyForm = ctx.FormValues()
	}

	return
}

func (s *EndpointMockParamService) MatchEndpointByMockPath(tenantId consts.TenantId, mockPath string, endpoint model.Endpoint) (
	pathParamsMap map[string]string, matched bool) {

	pathParams, _ := s.EndpointRepo.GetEndpointPathParams(tenantId, endpoint.ID)
	pathParamRegxMap := map[string]string{}
	for _, pathParam := range pathParams {
		paramRegxStr := ""
		if pathParam.Type == "number" || pathParam.Type == "integer" {
			paramRegxStr = "\\d+"
		} else if pathParam.Type == "boolean" {
			paramRegxStr = "true|false"
		} else {
			paramRegxStr = "[^/]+"
		}
		pathParamRegxMap[pathParam.Name] = fmt.Sprintf("(%s)", paramRegxStr)
	}

	pathRegxStr := endpoint.Path
	arr := regexp.MustCompile(`(?U)\{(.+)\}`).FindAllStringSubmatch(endpoint.Path, -1)
	for _, items := range arr {
		regxStr, ok := pathParamRegxMap[items[1]]
		if !ok {
			continue
		}
		pathRegxStr = strings.Replace(pathRegxStr, items[0], regxStr, 1)
	}

	pathRegxStr = "^" + strings.TrimSuffix(pathRegxStr, "/") + "/?$"
	arr1 := regexp.MustCompile(pathRegxStr).FindAllStringSubmatch(mockPath, -1)
	if len(arr1) > 0 {
		matched = true

		pathParamsMap = map[string]string{}
		for index, pathParam := range pathParams {
			pathParamsMap[pathParam.Name] = arr1[0][index+1]
		}
	}

	return
}

func (s *EndpointMockParamService) getRealCookies(tenantId consts.TenantId, ctx iris.Context) (ret []domain.ExecCookie) {
	ctx.VisitAllCookies(func(name string, value string) {
		item := domain.ExecCookie{
			Name:  name,
			Value: value,
		}

		cookie, err := ctx.GetRequestCookie(name)
		if err == nil && cookie != nil {
			item.Path = cookie.Path
			item.Domain = cookie.Domain
			item.ExpireTime = &cookie.Expires
		}

		ret = append(ret, item)
	})
	return
}
