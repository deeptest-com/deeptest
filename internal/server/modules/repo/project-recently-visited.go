package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
	"time"
)

type ProjectRecentlyVisitedRepo struct {
	DB *gorm.DB `inject:""`
}

func NewProjectRecentlyVisitedRepo() *ProjectRecentlyVisitedRepo {
	return &ProjectRecentlyVisitedRepo{}
}

func (r *ProjectRecentlyVisitedRepo) FindUserProjectToday(userId, projectId uint) (projectRecentlyVisited model.ProjectRecentlyVisited, err error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	db := r.DB.Model(&model.ProjectRecentlyVisited{}).Where("user_id = ?", userId).Where("project_id = ?", projectId).Where("created_at >= ?", today)

	err = db.First(&projectRecentlyVisited).Error
	return
}

func (r *ProjectRecentlyVisitedRepo) Create(req v1.ProjectRecentlyVisitedReq) (id uint, err error) {
	userProject, err := r.FindUserProjectToday(req.UserId, req.ProjectId)
	if userProject.ID != 0 {
		logUtils.Infof("项目访问记录今日已记录")
		return
	}

	projectRecentlyVisited := model.ProjectRecentlyVisited{ProjectRecentlyVisitedBase: req.ProjectRecentlyVisitedBase}
	err = r.DB.Create(&projectRecentlyVisited).Error
	if err != nil {
		logUtils.Errorf("创建项目最近访问记录失败%s", err.Error())
		return
	}
	id = projectRecentlyVisited.ID
	return
}
