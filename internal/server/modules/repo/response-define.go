package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ResponseDefineRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ResponseDefineRepo) Get(id uint) (responseDefine model.DebugConditionResponseDefine, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&responseDefine).Error
	return
}

func (r *ResponseDefineRepo) Save(responseDefine *model.DebugConditionResponseDefine) (err error) {

	err = r.DB.Save(responseDefine).Error
	return
}

func (r *ResponseDefineRepo) UpdateResult(responseDefine domain.ResponseDefineBase) (err error) {
	values := map[string]interface{}{
		"result_msg":    responseDefine.ResultMsg,
		"result_status": responseDefine.ResultStatus,
	}

	err = r.DB.Model(&model.DebugConditionResponseDefine{}).
		Where("id=?", responseDefine.ConditionEntityId).
		Updates(values).
		Error

	return
}

func (r *ResponseDefineRepo) Update(id uint, data map[string]interface{}) (err error) {
	err = r.DB.Model(&model.DebugConditionResponseDefine{}).
		Where("id=?", id).
		Updates(data).
		Error
	return
}
