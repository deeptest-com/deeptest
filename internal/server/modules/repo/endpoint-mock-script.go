package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EndpointMockScriptRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *EndpointMockScriptRepo) Get(tenantId consts.TenantId, endpointId uint) (script model.EndpointMockScript, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointMockScript{}).
		Where("endpoint_id = ?", endpointId).
		First(&script).Error

	if err == gorm.ErrRecordNotFound {
		script = model.EndpointMockScript{
			EndpointId: endpointId,
			Content:    "",
		}
		err = r.GetDB(tenantId).Save(&script).Error
	}

	return
}

func (r *EndpointMockScriptRepo) Update(tenantId consts.TenantId, po model.EndpointMockScript) (err error) {
	values := map[string]interface{}{"content": po.Content}

	err = r.GetDB(tenantId).Model(&model.EndpointMockScript{}).
		Where("id = ?", po.ID).Updates(values).Error

	if err != nil {
		logUtils.Errorf("update EndpointMockScript error", zap.String("error:", err.Error()))
		return err
	}

	return
}

func (r *EndpointMockScriptRepo) Disable(tenantId consts.TenantId, endpointId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("id = ?", endpointId).
		Update("script_mock_disabled", gorm.Expr("NOT script_mock_disabled")).Error

	return
}
