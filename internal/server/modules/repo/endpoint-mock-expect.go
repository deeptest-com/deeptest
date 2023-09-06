package repo

import "github.com/aaronchen2k/deeptest/internal/server/modules/model"

func (r *PlanRepo) ListByEndpointId(endpointId uint) (res []model.EndpointMockExpect, err error) {
	err = r.DB.Model(model.EndpointMockExpect{}).
		Where("endpoint_id = ?", endpointId).
		Where("NOT deleted AND NOT disabled").
		Group("method").
		Order("id desc").
		Find(&res).Error
	return
}
