package repo

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type ProjectPermRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ProjectPermRepo) GetRecordByNameAndAction(tenantId consts.TenantId, name, action string) (projectPerm model.ProjectPerm, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectPerm{}).
		Where("name = ?", name).
		Where("act = ?", action).
		First(&projectPerm).Error
	return
}

func (r *ProjectPermRepo) CreateIfNotExist(tenantId consts.TenantId, perms []model.ProjectPerm) (successCount int, failItems []string) {
	_ = r.GetDB(tenantId).Delete(&model.ProjectPerm{}, "id > 0").Error
	for _, perm := range perms {
		err := r.GetDB(tenantId).Model(&model.ProjectPerm{}).Create(&perm).Error
		if err != nil {
			logUtils.Errorf("添加权限%s失败，错误%s。", perm.Name, err.Error())
			failItems = append(failItems, fmt.Sprintf("添加权限%s失败，错误%s", perm.Name, err.Error()))
		} else {
			successCount++
		}
	}
	return
}
