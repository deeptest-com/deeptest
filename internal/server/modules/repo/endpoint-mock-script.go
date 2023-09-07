package repo

import (
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EndpointMockScriptRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *EndpointMockScriptRepo) Get(endpointId uint) (script model.EndpointMockScript, err error) {
	err = r.DB.Model(&model.EndpointMockScript{}).
		Where("endpoint_id = ?", endpointId).
		First(&script).Error

	if err == gorm.ErrRecordNotFound {
		script = model.EndpointMockScript{
			EndpointId: endpointId,
			Content:    "",
		}
		err = r.DB.Save(&script).Error
	}

	return
}

func (r *EndpointMockScriptRepo) Update(po model.EndpointMockScript) (err error) {
	values := map[string]interface{}{"content": po.Content}

	err = r.DB.Model(&model.EndpointMockScript{}).
		Where("id = ?", po.ID).Updates(values).Error

	if err != nil {
		logUtils.Errorf("update EndpointMockScript error", zap.String("error:", err.Error()))
		return err
	}

	return
}
