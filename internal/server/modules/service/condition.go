package service

import (
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	repo "github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type ConditionService struct {
	ConditionRepo   *repo.ConditionRepo   `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	CheckpointRepo  *repo.CheckpointRepo  `inject:""`
	DatabaseOptRepo *repo.DatabaseOptRepo `inject:""`
	ScriptRepo      *repo.ScriptRepo      `inject:""`
}

func (s *ConditionService) List(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint,
	category consts.ConditionCategory, usedBy consts.UsedBy, src consts.ConditionSrc) (conditions []model.DebugCondition, err error) {

	conditions, err = s.ConditionRepo.List(tenantId, debugInterfaceId, endpointInterfaceId, category, usedBy, "", src)

	return
}

func (s *ConditionService) Get(tenantId consts.TenantId, id uint) (condition model.DebugCondition, err error) {
	condition, err = s.ConditionRepo.Get(tenantId, id)

	return
}

func (s *ConditionService) Create(tenantId consts.TenantId, condition *model.DebugCondition) (err error) {
	err = s.ConditionRepo.Save(tenantId, condition)

	var entityId uint

	if condition.EntityType == consts.ConditionTypeScript {
		po := s.ScriptRepo.CreateDefault(tenantId, condition.ID, condition.ConditionSrc)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeDatabase {
		po := s.DatabaseOptRepo.CreateDefault(tenantId, condition.ID, condition.ConditionSrc)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeExtractor {
		po := s.ExtractorRepo.CreateDefault(tenantId, condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeCheckpoint {
		po := s.CheckpointRepo.CreateDefault(tenantId, condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeResponseDefine {
		//保存定义结构体
	}

	err = s.ConditionRepo.UpdateEntityId(tenantId, condition.ID, entityId)

	return
}

func (s *ConditionService) Delete(tenantId consts.TenantId, reqId uint) (err error) {
	err = s.ConditionRepo.Delete(tenantId, reqId)

	return
}

func (s *ConditionService) Disable(tenantId consts.TenantId, reqId uint) (err error) {
	err = s.ConditionRepo.Disable(tenantId, reqId)

	return
}

func (s *ConditionService) Move(tenantId consts.TenantId, req serverDomain.ConditionMoveReq) (err error) {
	err = s.ConditionRepo.UpdateOrders(tenantId, req)

	return
}

func (s *ConditionService) ResetForCase(tenantId consts.TenantId, endpointInterfaceId, debugInterfaceId uint) (err error) {
	usedBy := consts.CaseDebug
	err = s.ConditionRepo.RemoveAllForBenchmarkCase(tenantId, debugInterfaceId, endpointInterfaceId, usedBy, "")
	if err != nil {
		return
	}

	err = s.ConditionRepo.CloneAll(tenantId, debugInterfaceId, endpointInterfaceId, debugInterfaceId, usedBy, usedBy, "false")

	return
}
