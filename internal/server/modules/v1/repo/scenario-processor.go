package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ScenarioProcessorRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.TestProcessor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *ScenarioProcessorRepo) Save(req serverDomain.ScenarioProcessorReq) (err error) {
	if req.EntityType == consts.ProcessorLogicIf {
		po := model.ProcessorLogic{
			Expression: req.Expression,
			BaseModel: model.BaseModel{
				ID: req.EntityId,
			},
			ProcessorBase: model.ProcessorBase{
				Comments: req.Comments,
			},
		}

		err = r.DB.Save(po).Error
		if req.EntityId == 0 {
			r.UpdateEntityId(uint(req.Id), po.ID)
		}
	}

	return
}

func (r *ScenarioProcessorRepo) UpdateEntityId(id, entityId uint) (err error) {
	err = r.DB.Model(&model.TestProcessor{}).
		Where("id = ?", id).
		Update("entityId", entityId).Error

	return
}
