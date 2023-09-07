package repo

import "github.com/aaronchen2k/deeptest/internal/server/modules/model"

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
