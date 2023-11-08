package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type PostConditionService struct {
	PostConditionRepo *repo.PostConditionRepo `inject:""`
	ExtractorRepo     *repo.ExtractorRepo     `inject:""`
	CheckpointRepo    *repo.CheckpointRepo    `inject:""`
	DatabaseOptRepo   *repo.DatabaseOptRepo   `inject:""`
	ScriptRepo        *repo.ScriptRepo        `inject:""`
}

func (s *PostConditionService) List(debugInterfaceId, endpointInterfaceId uint,
	category consts.ConditionCategory, usedBy consts.UsedBy) (conditions []model.DebugPostCondition, err error) {
	conditions, err = s.PostConditionRepo.List(debugInterfaceId, endpointInterfaceId, category, usedBy)

	return
}

func (s *PostConditionService) Get(id uint) (condition model.DebugPostCondition, err error) {
	condition, err = s.PostConditionRepo.Get(id)

	return
}

func (s *PostConditionService) Create(condition *model.DebugPostCondition) (err error) {
	err = s.PostConditionRepo.Save(condition)

	var entityId uint

	if condition.EntityType == consts.ConditionTypeExtractor {
		po := s.ExtractorRepo.CreateDefault(condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeCheckpoint {
		po := s.CheckpointRepo.CreateDefault(condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeScript {
		po := s.ScriptRepo.CreateDefault(condition.ID, consts.ConditionSrcPost)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeDatabase {
		po := s.DatabaseOptRepo.CreateDefault(condition.ID)
		entityId = po.ID

	} else if condition.EntityType == consts.ConditionTypeResponseDefine {
		//保存定义结构体
	}

	err = s.PostConditionRepo.UpdateEntityId(condition.ID, entityId)

	return
}

func (s *PostConditionService) Delete(reqId uint) (err error) {
	err = s.PostConditionRepo.Delete(reqId)

	return
}

func (s *PostConditionService) Disable(reqId uint) (err error) {
	err = s.PostConditionRepo.Disable(reqId)

	return
}

func (s *PostConditionService) Move(req serverDomain.ConditionMoveReq) (err error) {
	err = s.PostConditionRepo.UpdateOrders(req)

	return
}
