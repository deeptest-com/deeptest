package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	mockHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/mock"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
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

func (s *MockAdvanceService) ByAdvanceMock(tenantId consts.TenantId, endpointInterface model.EndpointInterface, ctx iris.Context) (
	resp mockGenerator.Response, byAdvance bool) {

	endpoint, _ := s.EndpointRepo.Get(tenantId, endpointInterface.EndpointId)

	if endpoint.AdvancedMockDisabled && endpoint.ScriptMockDisabled {
		byAdvance = false
		return
	}

	var req mockGenerator.Request
	if !endpoint.AdvancedMockDisabled { // expect result
		req, resp, byAdvance = s.ByExpect(tenantId, endpointInterface, endpoint, ctx)
		if !byAdvance { // return if failed
			return
		}
	}

	if !endpoint.ScriptMockDisabled {
		s.ByScript(tenantId, endpoint, req, &resp)
	}

	return
}

func (s *MockAdvanceService) ByExpect(tenantId consts.TenantId, endpointInterface model.EndpointInterface, endpoint model.Endpoint, ctx iris.Context) (
	req mockGenerator.Request, resp mockGenerator.Response, byAdvance bool) {

	headerParams, queryParams, pathParams, body, bodyForm, cookies :=
		s.EndpointMockParamService.GetRealRequestValues(tenantId, ctx, endpointInterface, endpoint)

	req = s.genRequest(headerParams, queryParams, pathParams, body, bodyForm, cookies)
	req.Url = ctx.Path()
	req.Method = endpointInterface.Method

	expects, _ := s.EndpointMockExpectRepo.ListByEndpointId(tenantId, endpointInterface.EndpointId)

	for _, expect := range expects {
		if expect.Disabled || expect.Method != endpointInterface.Method {
			continue
		}

		expectRequestMap, _ := s.EndpointMockExpectRepo.GetExpectRequest(tenantId, expect.ID)

		if s.MatchExpect(expectRequestMap, endpointInterface, endpoint,
			headerParams, queryParams, pathParams, body, bodyForm, ctx) {

			respData, respHeaders := s.GetExpectResult(tenantId, expect)
			respDefine := s.EndpointInterfaceRepo.GetResponse(tenantId, endpointInterface.ID, respData.Code)

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

			expectResp, _ := s.EndpointMockExpectRepo.GetExpectResponse(tenantId, expect.ID)
			resp.DelayTime = expectResp.DelayTime

			if httpHelper.IsJsonRespType(resp.ContentType) && resp.Content != "" {
				json.Unmarshal([]byte(resp.Content), &resp.Data)
			} else {
				resp.Data = resp.Content
			}

			return
		}
	}

	return
}

func (s *MockAdvanceService) ByScript(tenantId consts.TenantId, endpoint model.Endpoint, req mockGenerator.Request, resp *mockGenerator.Response) {
	script, err := s.EndpointMockScriptRepo.Get(tenantId, endpoint.ID)
	if err != nil || script.Disabled || script.Content == "" {
		return
	}

	mockHelper.InitJsRuntime(tenantId, endpoint.ProjectId)
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
	endpointInterface model.EndpointInterface, endpoint model.Endpoint,
	headerParams []domain.Param, queryParams []domain.Param, pathParams []domain.Param,
	body string, bodyForm map[string][]string, ctx iris.Context) (ret bool) {

	if len(expectRequestMap) == 0 {
		return false
	}

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

						if s.EndpointMockCompareService.CompareString(actualVal, expectVal, compareWay) {
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

						if s.EndpointMockCompareService.CompareString(actualVal, expectVal, compareWay) {
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

						if s.EndpointMockCompareService.CompareString(actualVal, expectVal, compareWay) {
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

func (s *MockAdvanceService) GetExpectResult(tenantId consts.TenantId, expect model.EndpointMockExpect) (
	respContent model.EndpointMockExpectResponse, respHeaders []domain.Param) {

	respContent, _ = s.EndpointMockExpectRepo.GetExpectResponse(tenantId, expect.ID)
	headers, _ := s.EndpointMockExpectRepo.GetExpectResponseHeaders(tenantId, expect.ID)

	for _, item := range headers {
		header := domain.Param{
			Name:  item.Name,
			Value: item.Value,
		}
		respHeaders = append(respHeaders, header)
	}

	return
}

func (s *MockAdvanceService) genRequest(
	headerParams []domain.Param, queryParams []domain.Param,
	pathParams []domain.Param, body string, form map[string][]string, cookies []domain.ExecCookie) (
	req mockGenerator.Request) {

	req.Cookies = cookies
	req.Body = body

	for _, item := range headerParams {
		req.Headers = append(req.Headers, domain.Param{
			Name:  item.Name,
			Value: item.Value,
			Type:  item.Type,
		})
	}

	for _, item := range queryParams {
		req.QueryParams = append(req.QueryParams, domain.Param{
			Name:  item.Name,
			Value: item.Value,
			Type:  item.Type,
		})
	}

	for _, item := range pathParams {
		req.PathParams = append(req.PathParams, domain.Param{
			Name:  item.Name,
			Value: item.Value,
			Type:  item.Type,
		})
	}

	for key, val := range form {
		req.FormData = append(req.FormData, domain.BodyFormDataItem{
			Name:  key,
			Value: strings.Join(val, ","),
		})
	}

	return
}
