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

func (s *MetricsService) Get(tenantId consts.TenantId, id uint) (condition model.DebugMetrics, err error) {
	condition, err = s.MetricsRepo.Get(tenantId, id)

	return
}

func (s *MetricsService) Create(tenantId consts.TenantId, condition *model.DebugMetrics) (err error) {
	err = s.MetricsRepo.Save(tenantId, condition)

	var entityId uint

	if condition.EntityType == consts.MetricsTypeScript {
		po := s.ScriptRepo.CreateDefault(tenantId, condition.ID, condition.MetricsSrc)
		entityId = po.ID

	} else if condition.EntityType == consts.MetricsTypeDatabase {
		po := s.DatabaseOptRepo.CreateDefault(tenantId, condition.ID, condition.MetricsSrc)
		entityId = po.ID

	} else if condition.EntityType == consts.MetricsTypeExtractor {
		po := s.ExtractorRepo.CreateDefault(tenantId, condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.MetricsTypeCheckpoint {
		po := s.CheckpointRepo.CreateDefault(tenantId, condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.MetricsTypeResponseDefine {
		//保存定义结构体
	}

	err = s.MetricsRepo.UpdateEntityId(tenantId, condition.ID, entityId)

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

func (s *MetricsService) Move(tenantId consts.TenantId, req serverDomain.MetricsMoveReq) (err error) {
	err = s.MetricsRepo.UpdateOrders(tenantId, req)

	return
}

func (s *MetricsService) GetEntity(tenantId consts.TenantId, id uint) (entity interface{}, err error) {
	entity, err = s.MetricsRepo.GetEntity(tenantId, id)

	return
}
