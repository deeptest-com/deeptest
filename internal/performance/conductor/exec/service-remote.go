package conductorExec

import (
	"encoding/json"
	"fmt"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

type PerformanceRemoteService struct {
}

func (s *PerformanceRemoteService) GetPlanToExec(req ptdomain.PerformanceTestReq) (
	ret ptdomain.PerformanceTestData, err error) {

	url := "performanceTestPlans/exec/loadScenario"

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
		QueryParams: &[]domain.Param{
			{
				Name:  "planId",
				Value: fmt.Sprintf("%d", req.PlanId),
			},
			{
				Name:  "environmentId",
				Value: _stringUtils.IntToStr(req.EnvironmentId),
			},
		},
	}

	bytes, err := s.GetRequest(httpReq)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &ret)
	if err != nil {
		logUtils.Infof("get exec obj failed,err:%v", err.Error())
	}

	ret.ServerUrl = req.ServerUrl
	ret.Token = req.Token

	return
}

func (s *PerformanceRemoteService) GetRequest(httpReq domain.BaseRequest) (
	ret []byte, err error) {

	request, err := json.Marshal(httpReq)
	logUtils.Infof("get exec obj request, request: %s", string(request))

	resp, err := httpHelper.Get(httpReq)

	logUtils.Infof("get exec obj response, response: %s", resp.Content)

	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get exec obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof("get exec obj failed, err %v", err)
	}

	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	ret, err = json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	return
}
