package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (r *ProjectRepo) Paginate(req serverDomain.ProjectReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.Project{}).Where("NOT deleted")

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("disabled = ?", commonUtils.IsDisable(req.Enabled))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count project error", zap.String("error:", err.Error()))
		return
	}

	projects := make([]*model.Project, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&projects).Error
	if err != nil {
		logUtils.Errorf("query project error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(projects, count, req.Page, req.PageSize)

	return
}

func (r *ProjectRepo) FindById(id uint) (model.Project, error) {
	project := model.Project{}
	err := r.DB.Model(&model.Project{}).Where("id = ?", id).First(&project).Error
	if err != nil {
		logUtils.Errorf("find project by id error", zap.String("error:", err.Error()))
		return project, err
	}

	return project, nil
}

func (r *ProjectRepo) FindByName(projectName string, id uint) (project serverDomain.ProjectResp, err error) {
	db := r.DB.Model(&model.Project{}).
		Where("name = ?", projectName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&project)

	return
}

func (r *ProjectRepo) Create(req serverDomain.ProjectReq) (id uint, bizErr *_domain.BizErr) {
	po, err := r.FindByName(req.Name, 0)
	if po.Name != "" {
		bizErr.Code = _domain.ErrNameExist.Code
		return
	}

	project := model.Project{ProjectBase: req.ProjectBase}

	err = r.DB.Model(&model.Project{}).Create(&project).Error
	if err != nil {
		logUtils.Errorf("add project error", zap.String("error:", err.Error()))
		bizErr.Code = _domain.ErrComm.Code

		return
	}

	id = project.ID

	return
}

func (r *ProjectRepo) Update(id uint, req serverDomain.ProjectReq) error {
	project := model.Project{ProjectBase: req.ProjectBase}
	err := r.DB.Model(&model.Project{}).Where("id = ?", req.Id).Updates(&project).Error
	if err != nil {
		logUtils.Errorf("update project error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectRepo) UpdateDefaultEnvironment(id, envId uint) (err error) {
	err = r.DB.Model(&model.Project{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"default_environment_id": envId}).Error

	if err != nil {
		logUtils.Errorf("update project environment error", err.Error())
		return err
	}

	return
}

func (r *ProjectRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete project by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ProjectRepo) DeleteChildren(ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.Project{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete project error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectRepo) GetChildrenIds(id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE project AS (
			SELECT * FROM biz_project WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_project child, project WHERE child.parent_id = project.id
		)
		SELECT id FROM project WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children project error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ProjectRepo) ListProjectByUser(userId uint) (projects []model.Project, err error) {
	var projectIds []uint
	r.DB.Model(&model.ProjectMember{}).
		Select("project_id").Where("user_id = ?", userId).Scan(&projectIds)

	err = r.DB.Model(&model.Project{}).
		Where("id IN (?)", projectIds).
		Find(&projects).Error

	return
}

func (r *ProjectRepo) GetCurrProjectByUser(userId uint) (currProject model.Project, err error) {
	var user model.SysUser
	err = r.DB.Preload("Profile").
		Where("id = ?", userId).
		First(&user).
		Error

	err = r.DB.Model(&model.Project{}).
		Where("id = ?", user.Profile.CurrProjectId).
		First(&currProject).Error

	return
}
