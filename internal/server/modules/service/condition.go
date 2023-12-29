package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ConditionService struct {
	ConditionRepo   *repo.ConditionRepo   `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	CheckpointRepo  *repo.CheckpointRepo  `inject:""`
	DatabaseOptRepo *repo.DatabaseOptRepo `inject:""`
	ScriptRepo      *repo.ScriptRepo      `inject:""`
}

func (s *ConditionService) List(debugInterfaceId, endpointInterfaceId uint,
	category consts.ConditionCategory, usedBy consts.UsedBy, src consts.ConditionSrc) (conditions []model.DebugCondition, err error) {
	conditions, err = s.ConditionRepo.List(debugInterfaceId, endpointInterfaceId, category, usedBy, "", src)

	return
}

func (s *ConditionService) Get(id uint) (condition model.DebugCondition, err error) {
	condition, err = s.ConditionRepo.Get(id)

	return
}

func (s *ConditionService) Create(condition *model.DebugCondition) (err error) {
	err = s.ConditionRepo.Save(condition)

	var entityId uint

	if condition.EntityType == consts.ConditionTypeScript {
		po := s.ScriptRepo.CreateDefault(condition.ID, condition.ConditionSrc)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeDatabase {
		po := s.DatabaseOptRepo.CreateDefault(condition.ID, condition.ConditionSrc)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeExtractor {
		po := s.ExtractorRepo.CreateDefault(condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeCheckpoint {
		po := s.CheckpointRepo.CreateDefault(condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeResponseDefine {
		//保存定义结构体
	}

	err = s.ConditionRepo.UpdateEntityId(condition.ID, entityId)

	return
}

func (s *ConditionService) Delete(reqId uint) (err error) {
	err = s.ConditionRepo.Delete(reqId)

	return
}

func (s *ConditionService) Disable(reqId uint) (err error) {
	err = s.ConditionRepo.Disable(reqId)

	return
}

func (s *ConditionService) Move(req serverDomain.ConditionMoveReq) (err error) {
	err = s.ConditionRepo.UpdateOrders(req)

	return
}

func (s *ConditionService) ResetForCase(endpointInterfaceId, debugInterfaceId uint) (err error) {
	usedBy := consts.CaseDebug
	err = s.ConditionRepo.RemoveAllForBenchmarkCase(debugInterfaceId, endpointInterfaceId, usedBy, "")
	if err != nil {
		return
	}

	err = s.ConditionRepo.CloneAll(debugInterfaceId, endpointInterfaceId, debugInterfaceId, usedBy, usedBy, false)

	return
}
