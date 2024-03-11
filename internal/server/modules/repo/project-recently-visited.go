package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
	"time"
)

type ProjectRecentlyVisitedRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ProjectRecentlyVisitedRepo) FindUserProjectToday(tenantId consts.TenantId, userId, projectId uint) (projectRecentlyVisited model.ProjectRecentlyVisited, err error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	db := r.GetDB(tenantId).Model(&model.ProjectRecentlyVisited{}).Where("user_id = ?", userId).Where("project_id = ?", projectId).Where("created_at >= ?", today)

	err = db.First(&projectRecentlyVisited).Error
	return
}

func (r *ProjectRecentlyVisitedRepo) Create(tenantId consts.TenantId, req v1.ProjectRecentlyVisitedReq) (id uint, err error) {
	userLastVisitedProject, err := r.FindUserLastRecord(tenantId, req.UserId)
	if userLastVisitedProject.ProjectId == req.ProjectId {
		logUtils.Infof(fmt.Sprintf("用户%+v最后一次访问的项目已经是%+v", req.UserId, req.ProjectId))
		return
	}

	projectRecentlyVisited := model.ProjectRecentlyVisited{ProjectRecentlyVisitedBase: req.ProjectRecentlyVisitedBase}
	err = r.GetDB(tenantId).Create(&projectRecentlyVisited).Error
	if err != nil {
		logUtils.Errorf("创建项目最近访问记录失败%s", err.Error())
		return
	}
	id = projectRecentlyVisited.ID
	return
}

func (r *ProjectRecentlyVisitedRepo) FindUserLastRecord(tenantId consts.TenantId, userId uint) (projectRecentlyVisited model.ProjectRecentlyVisited, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectRecentlyVisited{}).Where("user_id = ?", userId).Last(&projectRecentlyVisited).Error
	return
}

func (r *ProjectRecentlyVisitedRepo) FindUserLastDistinctProjects(tenantId consts.TenantId, userId uint, limit int) (res []model.ProjectRecentlyVisited, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectRecentlyVisited{}).
		Select("group_concat(distinct project_id) as project_id, created_at").
		Where("user_id = ?", userId).
		Group("created_at").
		Order("created_at desc").
		Limit(limit).
		Find(&res).Error
	return
}
