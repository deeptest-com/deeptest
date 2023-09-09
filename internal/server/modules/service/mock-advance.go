package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12"
)

type MockAdvanceService struct {
	EndpointService     *EndpointService          `inject:""`
	ProjectSettingsRepo *repo.ProjectSettingsRepo `inject:""`
	EndpointRepo        *repo.EndpointRepo        `inject:""`

	EndpointMockExpectRepo    *repo.EndpointMockExpectRepo `inject:""`
	EndpointMockExpectService *EndpointMockExpectService   `inject:""`
	EndpointMockScriptService *EndpointMockScriptService   `inject:""`
}

func (s *MockAdvanceService) ByAdvanceMock(endpointInterface model.EndpointInterface, ctx iris.Context) (
	resp mockGenerator.Response, byAdvance bool) {

	endpoint, _ := s.EndpointRepo.Get(endpointInterface.EndpointId)

	if endpoint.AdvancedMockDisabled && endpoint.ScriptMockDisabled {
		byAdvance = false
		return
	}

	if !endpoint.AdvancedMockDisabled {
		resp, byAdvance = s.ByExpect(endpointInterface, endpoint)
		if !byAdvance {
			return
		}
	}

	if !endpoint.ScriptMockDisabled {
		s.ByScript(endpoint, &resp)
	}

	return
}

func (s *MockAdvanceService) ByExpect(endpointInterface model.EndpointInterface, endpoint model.Endpoint) (
	resp mockGenerator.Response, byAdvance bool) {
	expects, _ := s.EndpointMockExpectRepo.ListByEndpointId(endpointInterface.EndpointId)

	for _, expect := range expects {
		if expect.Disabled {
			continue
		}

		expectRequestMap, _ := s.EndpointMockExpectRepo.GetExpectRequest(endpointInterface.EndpointId)
		if s.MatchExpect(expectRequestMap, endpoint) {
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
	endpoint model.Endpoint) (ret bool) {
	//for source, expectRequests := range expectRequestMap {
	//
	//}

	return
}

func (s *MockAdvanceService) GetExpectResult(expect model.EndpointMockExpect) (
	respContent model.EndpointMockExpectResponse, respHeaders []model.EndpointMockExpectResponseHeader) {

	respContent, _ = s.EndpointMockExpectRepo.GetExpectResponse(expect.ID)
	respHeaders, _ = s.EndpointMockExpectRepo.GetExpectResponseHeaders(expect.ID)

	return
}

func (s *MockAdvanceService) getReqValues(ctx iris.Context, endpoint model.Endpoint) (
	headers map[string]string, queryParams map[string]string, pathParams map[string]string,
	body string, bodyForm map[string][]string) {

	ctx.ReadHeaders(&headers)

	ctx.ReadParams(queryParams)

	pathParams = s.getPathParamValues(ctx, endpoint)

	body, bodyForm = s.getBody(ctx)

	return
}

func (s *MockAdvanceService) getPathParamValues(ctx iris.Context, endpoint model.Endpoint) (pathParams map[string]string) {
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
