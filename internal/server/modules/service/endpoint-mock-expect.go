package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"sort"
)

type EndpointMockExpectService struct {
	EndpointMockExpectRepo *repo.EndpointMockExpectRepo `inject:""`
	EndpointInterfaceRepo  *repo.EndpointInterfaceRepo  `inject:""`
	EndpointRepo           *repo.EndpointRepo           `inject:""`
}

func (s *EndpointMockExpectService) List(endpointId uint) (res []model.EndpointMockExpect, err error) {
	res, err = s.EndpointMockExpectRepo.ListByEndpointId(endpointId)
	return
}

func (s *EndpointMockExpectService) GetDetail(expectId uint) (res model.EndpointMockExpect, err error) {
	res, err = s.EndpointMockExpectRepo.GetExpectDetail(expectId)
	return
}

func (s *EndpointMockExpectService) Save(req model.EndpointMockExpect) (expectId uint, err error) {
	expectId, err = s.EndpointMockExpectRepo.Save(req)
	return
}

func (s *EndpointMockExpectService) Copy(expectId, endpointId uint, username string) (id uint, err error) {
	expectDetail, err := s.GetDetail(expectId)
	if err != nil {
		return
	}

	s.InitExpectId(&expectDetail, endpointId)
	expectDetail.CreateUser = username
	expectDetail.UpdateUser = ""

	id, err = s.Save(expectDetail)

	return
}

func (s *EndpointMockExpectService) InitExpectId(expect *model.EndpointMockExpect, endpointId uint) {
	expect.ID = 0
	if endpointId != 0 {
		expect.EndpointId = endpointId
	} else {
		expect.Name = "copy-" + expect.Name
	}

	expect.ResponseBody.ID = 0

	for k, v := range expect.RequestHeaders {
		v.ID = 0
		expect.RequestHeaders[k] = v
	}

	for k, v := range expect.RequestBodies {
		v.ID = 0
		expect.RequestBodies[k] = v
	}

	for k, v := range expect.RequestQueryParams {
		v.ID = 0
		expect.RequestQueryParams[k] = v
	}

	for k, v := range expect.RequestPathParams {
		v.ID = 0
		expect.RequestPathParams[k] = v
	}

	for k, v := range expect.ResponseHeaders {
		v.ID = 0
		expect.ResponseHeaders[k] = v
	}
}

func (s *EndpointMockExpectService) DeleteById(expectId uint) (err error) {
	err = s.EndpointMockExpectRepo.DeleteById(expectId)
	return
}

func (s *EndpointMockExpectService) Disable(endpointId uint) (err error) {
	err = s.EndpointMockExpectRepo.Disable(endpointId)
	return
}

func (s *EndpointMockExpectService) SaveOrder(req v1.MockExpectIdsReq) (err error) {
	err = s.EndpointMockExpectRepo.SaveOrder(req)
	return
}

func (s *EndpointMockExpectService) UpdateExpectDisabled(expectId uint, disabled bool) (err error) {
	err = s.EndpointMockExpectRepo.UpdateDisabledStatus(expectId, disabled)
	return
}

func (s *EndpointMockExpectService) UpdateExpectName(expectId uint, name string) (err error) {
	err = s.EndpointMockExpectRepo.UpdateExpectName(expectId, name)
	return
}

func (s *EndpointMockExpectService) GetExpectRequestOptions(endpointId, endpointInterfaceId uint) (ret v1.MockExpectRequestOptions, err error) {
	ret = make(v1.MockExpectRequestOptions)

	headerOptions, err := s.GetExpectRequestHeaderOptions(endpointInterfaceId)
	if err != nil {
		return
	}
	ret[consts.ParamInHeader] = headerOptions

	return
}

func (s *EndpointMockExpectService) GetExpectRequestHeaderOptions(endpointInterfaceId uint) (options []v1.MockExpectRequestOption, err error) {
	headers, err := s.EndpointInterfaceRepo.ListHeaders(endpointInterfaceId)
	if err != nil {
		return
	}

	headerOptionArr := make([]string, 0)
	headerOptionMap := make(map[string]string)
	for _, v := range headers {
		headerOptionArr = append(headerOptionArr, v.Name)
		headerOptionMap[v.Name] = v.Type
	}

	commonOptions := append(headerOptionArr, consts.HeaderOptions...)
	sort.Strings(commonOptions)

	for _, v := range commonOptions {
		option := v1.MockExpectRequestOption{
			Name: v,
		}
		if optionType, ok := headerOptionMap[v]; ok {
			option.Type = optionType
		}
		options = append(options, option)
	}

	return
}

func (s *EndpointMockExpectService) GetExpectRequestBodyOptions(endpointInterfaceId uint) (options []string, err error) {
	body, err := s.EndpointInterfaceRepo.ListRequestBody(endpointInterfaceId)
	if err != nil {
		return
	}

	if body.SchemaItem.Content != "" {
		content := ReqBodyOther{}
		err = json.Unmarshal([]byte(body.SchemaItem.Content), &content)
		for name, _ := range content.Properties {
			options = append(options, name)
		}
	}

	return
}

func (s *EndpointMockExpectService) GetExpectRequestQueryOptions(endpointInterfaceId uint) (options []v1.MockExpectRequestOption, err error) {
	queries, err := s.EndpointInterfaceRepo.ListParams(endpointInterfaceId)
	if err != nil {
		return
	}

	for _, v := range queries {
		options = append(options, v1.MockExpectRequestOption{Name: v.Name, Type: v.Type})
	}

	return
}

func (s *EndpointMockExpectService) GetExpectRequestPathOptions(endpointId uint) (options []v1.MockExpectRequestOption, err error) {
	paths, err := s.EndpointRepo.GetEndpointPathParams(endpointId)
	if err != nil {
		return
	}

	for _, v := range paths {
		options = append(options, v1.MockExpectRequestOption{Name: v.Name, Type: v.Type})
	}

	return
}
