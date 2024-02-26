package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type EndpointMockExpectRepo struct {
	*BaseRepo `inject:""`
}

func (r *EndpointMockExpectRepo) ListByEndpointId(tenantId consts.TenantId, endpointId uint) (res []model.EndpointMockExpect, err error) {
	err = r.GetDB(tenantId).Model(model.EndpointMockExpect{}).
		Where("endpoint_id = ?", endpointId).
		Where("NOT deleted").
		Order(" ordr").
		Find(&res).Error

	return
}

func (r *EndpointMockExpectRepo) GetExpectDetail(tenantId consts.TenantId, expectId uint) (expect model.EndpointMockExpect, err error) {
	expect, err = r.GetExpectById(tenantId, expectId)
	if err != nil {
		return
	}

	expectRequests, err := r.GetExpectRequest(tenantId, expectId)
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

	responseBody, err := r.GetExpectResponse(tenantId, expectId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	expect.ResponseBody = responseBody

	responseHeaders, err := r.GetExpectResponseHeaders(tenantId, expectId)
	if err != nil {
		return
	}
	expect.ResponseHeaders = responseHeaders
	return
}

func (r *EndpointMockExpectRepo) GetExpectById(tenantId consts.TenantId, expectId uint) (expect model.EndpointMockExpect, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Where("NOT deleted").
		First(&expect).Error

	return
}
func (r *EndpointMockExpectRepo) GetExpectRequest(tenantId consts.TenantId, expectId uint) (res map[consts.ParamIn][]model.EndpointMockExpectRequest, err error) {
	res = make(map[consts.ParamIn][]model.EndpointMockExpectRequest)
	allRequests := make([]model.EndpointMockExpectRequest, 0)
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpectRequest{}).
		Where("endpoint_mock_expect_id = ?", expectId).
		Where("NOT deleted AND NOT disabled").
		Find(&allRequests).Error

	for _, v := range allRequests {
		res[v.Source] = append(res[v.Source], v)
	}

	return
}

func (r *EndpointMockExpectRepo) GetExpectResponse(tenantId consts.TenantId, expectId uint) (response model.EndpointMockExpectResponse, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpectResponse{}).
		Where("endpoint_mock_expect_id = ?", expectId).
		Where("NOT deleted AND NOT disabled").
		First(&response).Error

	return
}

func (r *EndpointMockExpectRepo) GetExpectResponseHeaders(tenantId consts.TenantId, expectId uint) (responseHeaders []model.EndpointMockExpectResponseHeader, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpectResponseHeader{}).
		Where("endpoint_mock_expect_id = ?", expectId).
		Where("NOT deleted AND NOT disabled").
		Find(&responseHeaders).Error

	return
}

func (r *EndpointMockExpectRepo) Save(tenantId consts.TenantId, req model.EndpointMockExpect) (expectId uint, err error) {
	if req.ID == 0 {
		req.Ordr = r.GetMaxOrder(tenantId, req.EndpointId)
	} else {
		if err = r.DeleteDetail(tenantId, req.ID); err != nil {
			return 0, err
		}
	}
	if err = r.GetDB(tenantId).Save(&req).Error; err != nil {
		return
	}

	req.ResponseBody.EndpointMockExpectId = req.ID
	if err = r.GetDB(tenantId).Save(&req.ResponseBody).Error; err != nil {
		return
	}

	if err = r.CreateExpectRequest(tenantId, req); err != nil {
		return
	}

	if err = r.CreateExpectResponseHeaders(tenantId, req); err != nil {
		return 0, err
	}

	expectId = req.ID
	return
}

func (r *EndpointMockExpectRepo) CreateExpectRequest(tenantId consts.TenantId, req model.EndpointMockExpect) (err error) {
	expectRequest := make([]model.EndpointMockExpectRequest, 0)
	for _, header := range req.RequestHeaders {
		if header.Name == "" || header.CompareWay == "" {
			continue
		}
		header.EndpointMockExpectId = req.ID
		header.ID = 0
		expectRequest = append(expectRequest, header)
	}
	for _, body := range req.RequestBodies {
		if body.CompareWay == "" || (body.SelectType != consts.FullText && body.Name == "") {
			continue
		}
		body.EndpointMockExpectId = req.ID
		body.ID = 0
		expectRequest = append(expectRequest, body)
	}
	for _, query := range req.RequestQueryParams {
		if query.Name == "" || query.CompareWay == "" {
			continue
		}
		query.EndpointMockExpectId = req.ID
		query.ID = 0
		expectRequest = append(expectRequest, query)
	}
	for _, path := range req.RequestPathParams {
		if path.Name == "" || path.CompareWay == "" {
			continue
		}
		path.EndpointMockExpectId = req.ID
		path.ID = 0
		expectRequest = append(expectRequest, path)
	}

	if len(expectRequest) > 0 {
		err = r.BatchCreateExpectRequest(tenantId, expectRequest)
	}

	return
}

func (r *EndpointMockExpectRepo) CreateExpectResponseHeaders(tenantId consts.TenantId, req model.EndpointMockExpect) (err error) {
	expectResponseHeaders := make([]model.EndpointMockExpectResponseHeader, 0)
	for _, v := range req.ResponseHeaders {
		v.ID = 0
		v.EndpointMockExpectId = req.ID
		expectResponseHeaders = append(expectResponseHeaders, v)
	}

	if len(expectResponseHeaders) > 0 {
		err = r.BatchCreateExpectResponseHeader(tenantId, expectResponseHeaders)
	}

	return
}

func (r *EndpointMockExpectRepo) BatchCreateExpectRequest(tenantId consts.TenantId, req []model.EndpointMockExpectRequest) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpectRequest{}).Create(req).Error

	return
}

func (r *EndpointMockExpectRepo) BatchCreateExpectResponseHeader(tenantId consts.TenantId, req []model.EndpointMockExpectResponseHeader) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpectResponseHeader{}).Create(req).Error

	return
}

func (r *EndpointMockExpectRepo) DeleteById(tenantId consts.TenantId, expectId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Update("deleted", 1).Error
	if err != nil {
		return
	}

	err = r.DeleteDetail(tenantId, expectId)

	return
}

func (r *EndpointMockExpectRepo) DeleteDetail(tenantId consts.TenantId, expectId uint) (err error) {
	modelArr := []interface{}{
		model.EndpointMockExpectRequest{},
		model.EndpointMockExpectResponse{},
		model.EndpointMockExpectResponseHeader{},
	}
	for _, v := range modelArr {
		if err = r.DeleteDetailByExpectId(tenantId, v, expectId); err != nil {
			return err
		}
	}

	return
}
func (r *EndpointMockExpectRepo) Disable(tenantId consts.TenantId, endpointId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("id = ?", endpointId).
		Update("advanced_mock_disabled", gorm.Expr("NOT advanced_mock_disabled")).Error

	return
}

func (r *EndpointMockExpectRepo) DeleteDetailByExpectId(tenantId consts.TenantId, model interface{}, expectId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model).
		Where("endpoint_mock_expect_id = ?", expectId).
		Update("deleted", 1).Error

	return
}

func (r *EndpointMockExpectRepo) SaveOrder(tenantId consts.TenantId, ids []uint) (err error) {
	return r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		for key, id := range ids {
			err = r.GetDB(tenantId).Model(&model.EndpointMockExpect{}).Where("id=?", id).Update("ordr", key).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *EndpointMockExpectRepo) UpdateDisabledStatus(tenantId consts.TenantId, expectId uint, disabled bool) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Update("disabled", disabled).Error

	return
}

func (r *EndpointMockExpectRepo) UpdateExpectName(tenantId consts.TenantId, expectId uint, name string) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockExpect{}).
		Where("id = ?", expectId).
		Update("name", name).Error

	return
}

func (r *EndpointMockExpectRepo) GetMaxOrder(tenantId consts.TenantId, endpointId uint) (order int) {
	expect := model.EndpointMockExpect{}

	err := r.GetDB(tenantId).Model(&model.EndpointMockExpect{}).
		Where("endpoint_id = ?", endpointId).
		Order("ordr DESC").
		First(&expect).Error

	if err == nil {
		order = expect.Ordr + 1
	}

	return
}
