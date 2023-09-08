package repo

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type MockJsRepo struct {
	*BaseRepo   `inject:""`
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
	ServeRepo   *ServeRepo   `inject:""`
}

func (r *MockJsRepo) ListExpressions() (tos []serverDomain.MockJsExpression, err error) {
	err = r.DB.Model(&model.MockJsExpression{}).
		Where("NOT deleted and expression != ''").
		Order("ordr ASC").
		Find(&tos).Error
	return
}

func (r *MockJsRepo) BatchCreateExpression(pos []model.MockJsExpression) (successCount int, failItems []string, err error) {
	for _, po := range pos {
		_, err := r.CreateExpression(po)
		if err != nil {
			failItems = append(failItems, po.Name)
			continue
		}
		successCount++
	}
	return
}

func (r *MockJsRepo) CreateExpression(po model.MockJsExpression) (id uint, err error) {
	menu, err := r.FindExpressionByName(po.Name)
	/*
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logUtils.Errorf("创建mockjs变量失败%s", err.Error())
			return
		}
	*/

	if menu.ID != 0 {
		po.ID = menu.ID
		err = r.DB.Save(&po).Error
		if err != nil {
			logUtils.Errorf("更新mockjs变量失败%s", err.Error())
			return
		}
		return
	}

	err = r.DB.Create(&po).Error
	if err != nil {
		logUtils.Errorf("创建mockjs变量失败%s", err.Error())
		return
	}
	id = po.ID
	return
}

func (r *MockJsRepo) FindExpressionByName(name string) (po model.MockJsExpression, err error) {
	db := r.DB.Model(&model.MockJsExpression{}).Where("name = ?", name)

	err = db.First(&po).Error
	return
}
