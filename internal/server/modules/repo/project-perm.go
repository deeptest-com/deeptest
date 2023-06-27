package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type ProjectPermRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ProjectPermRepo) GetRecordByNameAndAction(name, action string) (projectPerm model.ProjectPerm, err error) {
	err = r.DB.Model(&model.ProjectPerm{}).
		Where("name = ?", name).
		Where("act = ?", action).
		First(&projectPerm).Error
	return
}

func (r *ProjectPermRepo) CreateIfNotExist(perms []model.ProjectPerm) (successCount int, failItems []string) {
	_ = r.DB.Delete(&model.ProjectPerm{}, "id > 0").Error
	for _, perm := range perms {
		err := r.DB.Model(&model.ProjectPerm{}).Create(&perm).Error
		if err != nil {
			logUtils.Errorf("添加权限%s失败，错误%s。", perm.Name, err.Error())
			failItems = append(failItems, fmt.Sprintf("添加权限%s失败，错误%s", perm.Name, err.Error()))
		} else {
			successCount++
		}
	}
	return
}
