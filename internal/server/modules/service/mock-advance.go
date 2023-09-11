package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12"
	"log"
)

type MockAdvanceService struct {
	EndpointService     *EndpointService          `inject:""`
	ProjectSettingsRepo *repo.ProjectSettingsRepo `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`

	EndpointMockExpectRepo    *repo.EndpointMockExpectRepo `inject:""`
	EndpointMockExpectService *EndpointMockExpectService   `inject:""`
	EndpointMockScriptService *EndpointMockScriptService   `inject:""`
}

func (s *MockAdvanceService) ByAdvanceMock(endpointInterface model.EndpointInterface, paramsMap map[string]string, ctx iris.Context) (resp mockGenerator.Response, byAdvance bool) {

	endpoint, _ := s.EndpointRepo.Get(endpointInterface.EndpointId)

	if endpoint.AdvancedMockDisabled && endpoint.ScriptMockDisabled {
		byAdvance = false
		return
	}

	if !endpoint.AdvancedMockDisabled { // expect result
		resp, byAdvance = s.ByExpect(endpointInterface, endpoint, ctx)
		if !byAdvance { // return if failed
			return
		}
	}

	if !endpoint.ScriptMockDisabled {
		s.ByScript(endpoint, &resp)
	}

	return
}

func (s *MockAdvanceService) ByExpect(endpointInterface model.EndpointInterface, endpoint model.Endpoint, ctx iris.Context) (resp mockGenerator.Response, byAdvance bool) {
	expects, _ := s.EndpointMockExpectRepo.ListByEndpointId(endpointInterface.EndpointId)

	for _, expect := range expects {
		if expect.Disabled {
			continue
		}

		expectRequestMap, _ := s.EndpointMockExpectRepo.GetExpectRequest(expect.ID)
		if s.MatchExpect(expectRequestMap, endpointInterface, endpoint, ctx) {
			respBody, respHeaders := s.GetExpectResult(expect)

			resp.ContentType = endpointInterface.BodyType
			resp.Content = respBody.Value
			resp.Headers = respHeaders

			resp.UseAdvMockMock = true
		}
	}

	return
}

func (s *MockAdvanceService) ByScript(endpoint model.Endpoint, resp *mockGenerator.Response) {

	return
}

func (s *MockAdvanceService) MatchExpect(expectRequestMap map[consts.ParamIn][]model.EndpointMockExpectRequest,
	endpointInterface model.EndpointInterface, endpoint model.Endpoint, ctx iris.Context) (ret bool) {
	headerParams, queryParams, pathParams, body, bodyForm := s.getRealRequestValues(ctx, endpointInterface, endpoint)

	ret = true
	for source, expectRequests := range expectRequestMap {
		if source == consts.ParamInHeader {
			for _, item := range expectRequests {
				result := false

				for _, param := range headerParams {
					if item.Name == param.Name {
						actualVal := param.Value
						expectVal := item.Value

						if actualVal == expectVal {
							result = true
						}
					}
				}

				if !result {
					return false
				}
			}

		} else if source == consts.ParamInQuery {
			for _, item := range expectRequests {
				result := false

				for _, param := range queryParams {
					if item.Name == param.Name {
						actualVal := param.Value
						expectVal := item.Value

						if actualVal == expectVal {
							result = true
						}
					}
				}

				if !result {
					return false
				}
			}

		} else if source == consts.ParamInPath {
			for _, item := range expectRequests {
				result := false

				for _, param := range pathParams {
					if item.Name == param.Name {
						actualVal := param.Value
						expectVal := item.Value

						if actualVal == expectVal {
							result = true
						}
					}
				}

				if !result {
					return false
				}
			}

		} else if source == consts.ParamInBody {
			for _, item := range expectRequests {
				actualVal := body
				expectVal := item.Value

				if actualVal != expectVal {
					return false
				}
			}
		}

		// TODO:
		log.Println(headerParams, queryParams, pathParams, body, bodyForm, expectRequests)
	}

	return
}

func (s *MockAdvanceService) compare(expect model.EndpointMockExpect) (ret bool) {
	return
}

func (s *MockAdvanceService) GetExpectResult(expect model.EndpointMockExpect) (
	respContent model.EndpointMockExpectResponse, respHeaders []model.EndpointMockExpectResponseHeader) {

	respContent, _ = s.EndpointMockExpectRepo.GetExpectResponse(expect.ID)
	respHeaders, _ = s.EndpointMockExpectRepo.GetExpectResponseHeaders(expect.ID)

	return
}

func (s *MockAdvanceService) getRealRequestValues(ctx iris.Context,
	endpointInterface model.EndpointInterface, endpoint model.Endpoint) (
	headers []model.InterfaceParamBase, queryParams []model.InterfaceParamBase, pathParams []model.InterfaceParamBase,
	body string, bodyForm map[string][]string) {

	headers = s.getRealHeaderParamValues(ctx, endpointInterface)

	queryParams = s.getRealQueryParamValues(ctx, endpointInterface)

	pathParams = s.getRealPathParamValues(ctx, endpoint)

	body, bodyForm = s.getBody(ctx)

	return
}

func (s *MockAdvanceService) getRealHeaderParamValues(ctx iris.Context, endpointInterface model.EndpointInterface) (ret []model.InterfaceParamBase) {
	definedParams, _ := s.EndpointInterfaceRepo.ListHeaders(endpointInterface.ID)

	realParams := map[string]string{}
	ctx.ReadHeaders(&realParams)

	for _, mockParam := range definedParams {
		item := model.InterfaceParamBase{
			Name: mockParam.Name,
			Type: mockParam.Type,
		}

		val, ok := realParams[mockParam.Name]
		if ok {
			item.Value = val
		}

		ret = append(ret, item)
	}

	return
}

func (s *MockAdvanceService) getRealQueryParamValues(ctx iris.Context, endpointInterface model.EndpointInterface) (ret []model.InterfaceParamBase) {
	definedParams, _ := s.EndpointInterfaceRepo.ListParams(endpointInterface.ID)

	realParams := map[string]string{}
	ctx.ReadQuery(&realParams)

	for _, mockParam := range definedParams {
		item := model.InterfaceParamBase{
			Name: mockParam.Name,
			Type: mockParam.Type,
		}

		val, ok := realParams[mockParam.Name]
		if ok {
			item.Value = val
		}

		ret = append(ret, item)
	}

	return
}

func (s *MockAdvanceService) getRealPathParamValues(ctx iris.Context, endpoint model.Endpoint) (ret []model.InterfaceParamBase) {
	definedParams, _ := s.EndpointRepo.GetEndpointPathParams(endpoint.ID)

	realParams := map[string]string{}
	ctx.ReadParams(&realParams)

	for _, mockParam := range definedParams {
		item := model.InterfaceParamBase{
			Name: mockParam.Name,
			Type: mockParam.Type,
		}

		val, ok := realParams[mockParam.Name]
		if ok {
			item.Value = val
		}

		ret = append(ret, item)
	}

	return
}

func (s *MockAdvanceService) getBody(ctx iris.Context) (body string, bodyForm map[string][]string) {
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
