package service

import (
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	repo "github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type MetricsService struct {
	MetricsRepo *repo.MetricsRepo `inject:""`
}

func (s *MetricsService) List(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (
	metrics []model.AiMetrics, err error) {

	metrics, err = s.MetricsRepo.List(tenantId, debugInterfaceId, endpointInterfaceId)

	return
}

func (s *MetricsService) Get(tenantId consts.TenantId, id uint) (condition model.AiMetrics, err error) {
	condition, err = s.MetricsRepo.Get(tenantId, id)

	return
}

func (s *MetricsService) Create(tenantId consts.TenantId, metrics *model.AiMetrics) (err error) {
	err = s.MetricsRepo.Save(tenantId, metrics)

	return
}

func (s *MetricsService) Update(tenantId consts.TenantId, metrics *model.AiMetrics) (err error) {
	err = s.MetricsRepo.Save(tenantId, metrics)

	return
}

func (s *MetricsService) Delete(tenantId consts.TenantId, reqId uint) (err error) {
	err = s.MetricsRepo.Delete(tenantId, reqId)

	return
}

func (s *MetricsService) Disable(tenantId consts.TenantId, reqId uint) (err error) {
	err = s.MetricsRepo.Disable(tenantId, reqId)

	return
}

func (s *MetricsService) Move(tenantId consts.TenantId, req serverDomain.ConditionMoveReq) (err error) {
	err = s.MetricsRepo.UpdateOrders(tenantId, req)

	return
}
