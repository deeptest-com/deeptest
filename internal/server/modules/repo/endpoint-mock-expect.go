package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type EndpointMockExpectRepo struct {
	*BaseRepo `inject:""`
}

func (r *EndpointMockExpectRepo) ListByEndpointId(endpointId uint) (res []model.EndpointMockExpect, err error) {
	err = r.DB.Model(model.EndpointMockExpect{}).
		Where("endpoint_id = ?", endpointId).
		Where("NOT deleted").
		Order(" ordr").
		Find(&res).Error

	return
}

func (r *EndpointMockExpectRepo) GetExpectDetail(expectId uint) (expect model.EndpointMockExpect, err error) {
	expect, err = r.GetExpectById(expectId)
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
	if err != nil && err != gorm.ErrRecordNotFound {
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
	res = make(map[consts.ParamIn][]model.EndpointMockExpectRequest)
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

func (r *EndpointMockExpectRepo) Save(req model.EndpointMockExpect) (expectId uint, err error) {
	isCreate := false
	if req.ID == 0 {
		isCreate = true
		req.Ordr = r.GetMaxOrder(req.EndpointId)
	}
	if err = r.DB.Save(&req).Error; err != nil {
		return
	}

	req.ResponseBody.EndpointMockExpectId = req.ID
	if err = r.DB.Save(&req.ResponseBody).Error; err != nil {
		return
	}

	if isCreate {
		if err = r.CreateExpectRequest(req); err != nil {
			return 0, err
		}

		expectResponseHeaders := make([]model.EndpointMockExpectResponseHeader, 0)
		for _, v := range req.ResponseHeaders {
			v.EndpointMockExpectId = req.ID
			expectResponseHeaders = append(expectResponseHeaders, v)
		}
		if len(expectResponseHeaders) > 0 {
			if err = r.BatchCreateExpectResponseHeader(expectResponseHeaders); err != nil {
				return 0, err
			}
		}
	} else {
		if err = r.DeleteRequestByExpectIdAndSource(req.ID, consts.ParamInBody); err != nil {
			return 0, err
		}

		if err = r.UpdateExpectRequest(req); err != nil {
			return 0, err
		}

		for _, v := range req.ResponseHeaders {
			if err = r.DB.Save(&v).Error; err != nil {
				return 0, err
			}
		}
	}
	expectId = req.ID
	return
}

func (r *EndpointMockExpectRepo) CreateExpectRequest(req model.EndpointMockExpect) (err error) {
	expectRequest := make([]model.EndpointMockExpectRequest, 0)
	for _, header := range req.RequestHeaders {
		header.EndpointMockExpectId = req.ID
		expectRequest = append(expectRequest, header)
	}
	for _, body := range req.RequestBodies {
		body.EndpointMockExpectId = req.ID
		expectRequest = append(expectRequest, body)
	}
	for _, query := range req.RequestQueryParams {
		query.EndpointMockExpectId = req.ID
		expectRequest = append(expectRequest, query)
	}
	for _, path := range req.RequestPathParams {
		path.EndpointMockExpectId = req.ID
		expectRequest = append(expectRequest, path)
	}

	if len(expectRequest) > 0 {
		err = r.BatchCreateExpectRequest(expectRequest)
	}

	return
}

func (r *EndpointMockExpectRepo) UpdateExpectRequest(req model.EndpointMockExpect) (err error) {
	for _, header := range req.RequestHeaders {
		if err = r.DB.Save(&header).Error; err != nil {
			return err
		}
	}
	for _, body := range req.RequestBodies {
		if err = r.DB.Save(&body).Error; err != nil {
			return err
		}
	}
	for _, query := range req.RequestQueryParams {
		if err = r.DB.Save(&query).Error; err != nil {
			return err
		}
	}
	for _, path := range req.RequestPathParams {
		if err = r.DB.Save(&path).Error; err != nil {
			return err
		}
	}
	return
}

func (r *EndpointMockExpectRepo) BatchCreateExpectRequest(req []model.EndpointMockExpectRequest) (err error) {
	err = r.DB.Model(&model.EndpointMockExpectRequest{}).Create(req).Error

	return
}

func (r *EndpointMockExpectRepo) BatchCreateExpectResponseHeader(req []model.EndpointMockExpectResponseHeader) (err error) {
	err = r.DB.Model(&model.EndpointMockExpectResponseHeader{}).Create(req).Error

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

func (r *EndpointMockExpectRepo) Disable(endpointId uint) (err error) {
	err = r.DB.Model(&model.Endpoint{}).
		Where("id = ?", endpointId).
		Update("advanced_mock_disabled", gorm.Expr("NOT advanced_mock_disabled")).Error

	return
}

func (r *EndpointMockExpectRepo) DeleteDetailByExpectId(model interface{}, expectId uint) (err error) {
	err = r.DB.Model(&model).
		Where("endpoint_mock_expect_id = ?", expectId).
		Update("deleted", 1).Error

	return
}

func (r *EndpointMockExpectRepo) DeleteRequestByExpectIdAndSource(expectId uint, selectType consts.ParamIn) (err error) {
	err = r.DB.Model(&model.EndpointMockExpectRequest{}).
		Where("endpoint_mock_expect_id = ?", expectId).
		Where("source = ?", selectType).
		Update("deleted", 1).Error

	return
}

func (r *EndpointMockExpectRepo) SaveOrder(ids []uint) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for key, id := range ids {
			err = r.DB.Model(&model.EndpointMockExpect{}).Where("id=?", id).Update("ordr", key).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *EndpointMockExpectRepo) UpdateDisabledStatus(expectId uint, disabled bool) (err error) {
	err = r.DB.Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Update("disabled", disabled).Error

	return
}

func (r *EndpointMockExpectRepo) UpdateExpectName(expectId uint, name string) (err error) {
	err = r.DB.Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Update("name", name).Error

	return
}

func (r *EndpointMockExpectRepo) GetMaxOrder(endpointId uint) (order int) {
	expect := model.EndpointMockExpect{}

	err := r.DB.Model(&model.EndpointMockExpect{}).
		Where("endpoint_id = ?", endpointId).
		Order("ordr DESC").
		First(&expect).Error

	if err == nil {
		order = expect.Ordr + 1
	}

	return
}
