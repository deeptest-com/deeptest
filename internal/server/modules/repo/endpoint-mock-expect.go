package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type EndpointMockExpectRepo struct {
	*BaseRepo `inject:""`
}

func (r *EndpointMockExpectRepo) ListByEndpointId(endpointId uint) (res []model.EndpointMockExpect, err error) {
	err = r.DB.Model(model.EndpointMockExpect{}).
		Where("endpoint_id = ?", endpointId).
		Where("NOT deleted AND NOT disabled").
		Group("method").
		Order("id desc").
		Find(&res).Error

	return
}

func (r *EndpointMockExpectRepo) GetExpectDetail(expectId uint) (res model.EndpointMockExpect, err error) {
	expect, err := r.GetExpectById(expectId)
	if err != nil {
		return
	}

	expectRequests, err := r.GetExpectRequest(expectId)
	if err != nil {
		return
	}
	if requestHeader, ok := expectRequests[consts.ParamInHeader]; ok {
		expect.RequestHeaders = requestHeader
	}
	if requestBody, ok := expectRequests[consts.ParamInBody]; ok {
		expect.RequestBodies = requestBody
	}
	if requestQuery, ok := expectRequests[consts.ParamInQuery]; ok {
		expect.RequestQueryParams = requestQuery
	}
	if requestPath, ok := expectRequests[consts.ParamInPath]; ok {
		expect.RequestPathParams = requestPath
	}

	responseBody, err := r.GetExpectResponse(expectId)
	if err != nil {
		return
	}
	expect.ResponseBody = responseBody

	responseHeaders, err := r.GetExpectResponseHeaders(expectId)
	if err != nil {
		return
	}
	expect.ResponseHeaders = responseHeaders
	return
}

func (r *EndpointMockExpectRepo) GetExpectById(expectId uint) (expect model.EndpointMockExpect, err error) {
	err = r.DB.Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Where("NOT deleted AND NOT disabled").
		First(&expect).Error

	return
}
func (r *EndpointMockExpectRepo) GetExpectRequest(expectId uint) (res map[consts.ParamIn][]model.EndpointMockExpectRequest, err error) {
	allRequests := make([]model.EndpointMockExpectRequest, 0)
	err = r.DB.Model(&model.EndpointMockExpectRequest{}).
		Where("endpoint_mock_expect_id = ?", expectId).
		Where("NOT deleted AND NOT disabled").
		Find(&allRequests).Error

	for _, v := range allRequests {
		res[v.Source] = append(res[v.Source], v)
	}

	return
}

func (r *EndpointMockExpectRepo) GetExpectResponse(expectId uint) (response model.EndpointMockExpectResponse, err error) {
	err = r.DB.Model(&model.EndpointMockExpectResponse{}).
		Where("endpoint_mock_expect_id = ?", expectId).
		Where("NOT deleted AND NOT disabled").
		First(&response).Error

	return
}

func (r *EndpointMockExpectRepo) GetExpectResponseHeaders(expectId uint) (responseHeaders []model.EndpointMockExpectResponseHeader, err error) {
	err = r.DB.Model(&model.EndpointMockExpectResponseHeader{}).
		Where("endpoint_mock_expect_id = ?", expectId).
		Where("NOT deleted AND NOT disabled").
		Find(&responseHeaders).Error

	return
}

func (r *EndpointMockExpectRepo) DeleteById(expectId uint) (err error) {
	err = r.DB.Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Update("deleted", 1).Error
	if err != nil {
		return
	}

	modelArr := []interface{}{
		model.EndpointMockExpectRequest{},
		model.EndpointMockExpectResponse{},
		model.EndpointMockExpectResponseHeader{},
	}
	for _, v := range modelArr {
		if err = r.DeleteDetailByExpectId(v, expectId); err != nil {
			return err
		}
	}
	return
}

func (r *EndpointMockExpectRepo) DeleteDetailByExpectId(model interface{}, expectId uint) (err error) {
	err = r.DB.Model(&model).
		Where("endpoint_mock_expect_id = ?", expectId).
		Update("deleted", 1).Error

	return
}
