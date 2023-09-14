package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	mockHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/mock"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12"
	"strconv"
)

type MockAdvanceService struct {
	EndpointService     *EndpointService          `inject:""`
	ProjectSettingsRepo *repo.ProjectSettingsRepo `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`

	EndpointMockExpectRepo    *repo.EndpointMockExpectRepo `inject:""`
	EndpointMockScriptRepo    *repo.EndpointMockScriptRepo `inject:""`
	EndpointMockExpectService *EndpointMockExpectService   `inject:""`
	EndpointMockScriptService *EndpointMockScriptService   `inject:""`

	EndpointMockParamService   *EndpointMockParamService   `inject:""`
	EndpointMockCompareService *EndpointMockCompareService `inject:""`
}

func (s *MockAdvanceService) ByAdvanceMock(endpointInterface model.EndpointInterface, ctx iris.Context) (
	resp mockGenerator.Response, byAdvance bool) {

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
		req := mockGenerator.Request{
			Method: endpointInterface.Method,
		}
		s.ByScript(endpoint, req, &resp)
	}

	return
}

func (s *MockAdvanceService) ByExpect(endpointInterface model.EndpointInterface, endpoint model.Endpoint, ctx iris.Context) (
	resp mockGenerator.Response, byAdvance bool) {

	expects, _ := s.EndpointMockExpectRepo.ListByEndpointId(endpointInterface.EndpointId)

	for _, expect := range expects {
		if expect.Disabled || expect.Method != endpointInterface.Method {
			continue
		}

		expectRequestMap, _ := s.EndpointMockExpectRepo.GetExpectRequest(expect.ID)
		if s.MatchExpect(expectRequestMap, endpointInterface, endpoint, ctx) {
			respData, respHeaders := s.GetExpectResult(expect)
			respDefine := s.EndpointInterfaceRepo.GetResponse(endpointInterface.ID, respData.Code)

			codeInt, _ := strconv.ParseInt(respData.Code, 10, 64)
			resp.StatusCode = consts.HttpRespCode(codeInt)
			resp.Headers = respHeaders
			resp.Content = respData.Value

			resp.UseAdvMock = true
			byAdvance = true

			resp.ContentType = consts.HttpContentType(respDefine.MediaType)
			if resp.ContentType == "" {
				resp.ContentType = consts.ContentTypeJSON
			}

			expectResp, _ := s.EndpointMockExpectRepo.GetExpectResponse(expect.ID)
			resp.DelayTime = expectResp.DelayTime

			if httpHelper.IsJsonRespType(resp.ContentType) && resp.Content != "" {
				json.Unmarshal([]byte(resp.Content), &resp.Data)
			} else {
				resp.Data = resp.Content
			}
		}
	}

	return
}

func (s *MockAdvanceService) ByScript(endpoint model.Endpoint, req mockGenerator.Request, resp *mockGenerator.Response) {
	script, err := s.EndpointMockScriptRepo.Get(endpoint.ID)
	if err != nil || script.Disabled || script.Content == "" {
		return
	}

	mockHelper.InitJsRuntime()
	mockHelper.SetReqValueToGoja(req)
	mockHelper.SetRespValueToGoja(*resp)
	mockHelper.ExecScript(script.Content)
	mockHelper.GetRespValueFromGoja()

	if mockHelper.CurrResponse.Data != nil {
		*resp = mockHelper.CurrResponse
	}

	return
}

func (s *MockAdvanceService) MatchExpect(expectRequestMap map[consts.ParamIn][]model.EndpointMockExpectRequest,
	endpointInterface model.EndpointInterface, endpoint model.Endpoint, ctx iris.Context) (ret bool) {

	if len(expectRequestMap) == 0 {
		return false
	}

	headerParams, queryParams, pathParams, body, bodyForm :=
		s.EndpointMockParamService.GetRealRequestValues(ctx, endpointInterface, endpoint)

	ret = true
	for source, expectRequests := range expectRequestMap {
		if source == consts.ParamInQuery {
			for _, item := range expectRequests {
				result := false

				for _, param := range queryParams {
					if item.Name == param.Name {
						actualVal := param.Value
						expectVal := item.Value
						compareWay := item.CompareWay

						if s.EndpointMockCompareService.CompareValue(actualVal, expectVal, compareWay) {
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
						compareWay := item.CompareWay

						if s.EndpointMockCompareService.CompareValue(actualVal, expectVal, compareWay) {
							result = true
						}
					}
				}

				if !result {
					return false
				}
			}

		} else if source == consts.ParamInHeader {
			for _, item := range expectRequests {
				result := false

				for _, param := range headerParams {
					if item.Name == param.Name {
						actualVal := param.Value
						expectVal := item.Value
						compareWay := item.CompareWay

						if s.EndpointMockCompareService.CompareValue(actualVal, expectVal, compareWay) {
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
				contentType := endpointInterface.BodyType

				ret = s.EndpointMockCompareService.CompareBody(item, contentType, body, bodyForm)

				if !ret {
					return
				}
			}
		}
	}

	return
}

func (s *MockAdvanceService) GetExpectResult(expect model.EndpointMockExpect) (
	respContent model.EndpointMockExpectResponse, respHeaders []model.EndpointMockExpectResponseHeader) {

	respContent, _ = s.EndpointMockExpectRepo.GetExpectResponse(expect.ID)
	respHeaders, _ = s.EndpointMockExpectRepo.GetExpectResponseHeaders(expect.ID)

	return
}
