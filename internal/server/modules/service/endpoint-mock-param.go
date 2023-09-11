package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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

func (s *EndpointMockParamService) GetRealRequestValues(ctx iris.Context,
	endpointInterface model.EndpointInterface, endpoint model.Endpoint) (
	headers []model.InterfaceParamBase, queryParams []model.InterfaceParamBase, pathParams []model.InterfaceParamBase,
	body string, bodyForm map[string][]string) {

	queryParams = s.getRealQueryParamValues(ctx, endpointInterface)

	pathParams = s.getRealPathParamValues(ctx, endpoint)

	headers = s.getRealHeaderParamValues(ctx, endpointInterface)

	body, bodyForm = s.getRealBody(ctx)

	return
}

func (s *EndpointMockParamService) getRealQueryParamValues(ctx iris.Context, endpointInterface model.EndpointInterface) (ret []model.InterfaceParamBase) {
	definedParams, _ := s.EndpointInterfaceRepo.ListParams(endpointInterface.ID)
	realParams := ctx.URLParams()

	for _, mockParam := range definedParams {
		item := model.InterfaceParamBase{
			Name: mockParam.Name,
			Type: mockParam.Type,
		}

		val, ok := realParams[mockParam.Name]
		if ok {
			item.Value = val
			ret = append(ret, item)
		}
	}

	return
}

func (s *EndpointMockParamService) getRealPathParamValues(ctx iris.Context, endpoint model.Endpoint) (ret []model.InterfaceParamBase) {
	definedParams, _ := s.EndpointRepo.GetEndpointPathParams(endpoint.ID)

	mockPath := "/" + ctx.Params().Get("path")
	realParams, _ := s.computerMockPathParam(mockPath, endpoint)

	for _, mockParam := range definedParams {
		item := model.InterfaceParamBase{
			Name: mockParam.Name,
			Type: mockParam.Type,
		}

		val, ok := realParams[mockParam.Name]
		if ok {
			item.Value = val
			ret = append(ret, item)
		}
	}

	return
}

func (s *EndpointMockParamService) getRealHeaderParamValues(ctx iris.Context, endpointInterface model.EndpointInterface) (ret []model.InterfaceParamBase) {
	definedParams, _ := s.EndpointInterfaceRepo.ListHeaders(endpointInterface.ID)

	realParams := ctx.Request().Header

	for _, mockParam := range definedParams {
		item := model.InterfaceParamBase{
			Name: mockParam.Name,
			Type: mockParam.Type,
		}

		val, ok := realParams[mockParam.Name]
		if ok {
			item.Value = strings.Join(val, ",")
			ret = append(ret, item)
		}
	}

	return
}

func (s *EndpointMockParamService) getRealBody(ctx iris.Context) (body string, bodyForm map[string][]string) {
	method := ctx.Method()
	if method != consts.POST.String() && method != consts.PUT.String() && method != consts.DELETE.String() {
		return
	}

	reqContentType := ctx.GetContentType()

	if reqContentType == consts.ContentTypeJSON.String() {
		var req interface{}
		ctx.ReadJSON(&req)

	} else if reqContentType == consts.ContentTypeFormData.String() {
		bodyForm = ctx.FormValues()
	}

	return
}

func (s *EndpointMockParamService) computerMockPathParam(mockPath string, endpoint model.Endpoint) (
	paramsMap map[string]string, matched bool) {

	pathParams, _ := s.EndpointRepo.GetEndpointPathParams(endpoint.ID)
	pathParamRegxMap := map[string]string{}
	for _, pathParam := range pathParams {
		paramRegxStr := ""
		if pathParam.Type == "number" || pathParam.Type == "interger" {
			paramRegxStr = "\\d"
		} else if pathParam.Type == "boolean" {
			paramRegxStr = "true|false"
		} else {
			paramRegxStr = ".+"
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

	arr1 := regexp.MustCompile("^"+pathRegxStr+"$").FindAllStringSubmatch(mockPath, -1)
	if len(arr1) > 0 {
		matched = true

		paramsMap = map[string]string{}
		for index, pathParam := range pathParams {
			paramsMap[pathParam.Name] = arr1[0][index+1]
		}
	}

	return
}
