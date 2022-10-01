package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	model2 "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	DB              *gorm.DB         `inject:""`
	RoleRepo        *RoleRepo        `inject:""`
	ProjectRoleRepo *ProjectRoleRepo `inject:""`
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (r *ProjectRepo) Paginate(req v1.ProjectReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model2.Project{}).Where("NOT deleted")

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

	projects := make([]*model2.Project, 0)

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

func (r *ProjectRepo) FindById(id uint) (model2.Project, error) {
	project := model2.Project{}
	err := r.DB.Model(&model2.Project{}).Where("id = ?", id).First(&project).Error
	if err != nil {
		logUtils.Errorf("find project by id error", zap.String("error:", err.Error()))
		return project, err
	}

	return project, nil
}

func (r *ProjectRepo) FindByName(projectName string, id uint) (project v1.ProjectResp, err error) {
	db := r.DB.Model(&model2.Project{}).
		Where("name = ?", projectName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&project)

	return
}

func (r *ProjectRepo) Create(req v1.ProjectReq, userId uint) (id uint, bizErr *_domain.BizErr) {
	po, err := r.FindByName(req.Name, 0)
	if po.Name != "" {
		bizErr.Code = _domain.ErrNameExist.Code
		return
	}

	project := model2.Project{ProjectBase: req.ProjectBase}

	err = r.DB.Model(&model2.Project{}).Create(&project).Error
	if err != nil {
		logUtils.Errorf("add project error", zap.String("error:", err.Error()))
		bizErr.Code = _domain.SystemErr.Code

		return
	}

	err = r.AddProjectMember(project.ID, userId)
	if err != nil {
		logUtils.Errorf("添加项目角色错误", zap.String("错误:", err.Error()))
		return
	}

	err = r.AddProjectRootInterface(project.ID)
	if err != nil {
		logUtils.Errorf("添加接口错误", zap.String("错误:", err.Error()))
		return
	}

	id = project.ID

	return
}

func (r *ProjectRepo) Update(id uint, req v1.ProjectReq) error {
	project := model2.Project{ProjectBase: req.ProjectBase}
	err := r.DB.Model(&model2.Project{}).Where("id = ?", req.Id).Updates(&project).Error
	if err != nil {
		logUtils.Errorf("update project error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectRepo) UpdateDefaultEnvironment(projectId, envId uint) (err error) {
	err = r.DB.Model(&model2.Project{}).
		Where("id = ?", projectId).
		Updates(map[string]interface{}{"environment_id": envId}).Error

	if err != nil {
		logUtils.Errorf("update project environment error", err.Error())
		return err
	}

	return
}

func (r *ProjectRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model2.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete project by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ProjectRepo) DeleteChildren(ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model2.Project{}).Where("id IN (?)", ids).
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

func (r *ProjectRepo) ListProjectByUser(userId uint) (projects []model2.Project, err error) {
	var projectIds []uint
	r.DB.Model(&model2.ProjectMember{}).
		Select("project_id").Where("user_id = ?", userId).Scan(&projectIds)

	err = r.DB.Model(&model2.Project{}).
		Where("NOT deleted AND id IN (?)", projectIds).
		Find(&projects).Error

	return
}

func (r *ProjectRepo) GetCurrProjectByUser(userId uint) (currProject model2.Project, err error) {
	var user model2.SysUser
	err = r.DB.Preload("Profile").
		Where("id = ?", userId).
		First(&user).
		Error

	err = r.DB.Model(&model2.Project{}).
		Where("id = ?", user.Profile.CurrProjectId).
		First(&currProject).Error

	return
}

func (r *ProjectRepo) ChangeProject(projectId, userId uint) (err error) {
	err = r.DB.Model(&model2.SysUserProfile{}).Where("user_id = ?", userId).
		Updates(map[string]interface{}{"curr_project_id": projectId}).Error

	return
}

func (r *ProjectRepo) AddProjectMember(projectId, userId uint) (err error) {
	projectRole, _ := r.ProjectRoleRepo.GetFirstOne()
	projectMember := model2.ProjectMember{UserId: userId, ProjectId: projectId, ProjectRoleId: projectRole.ID}
	err = r.DB.Create(&projectMember).Error

	return
}

func (r *ProjectRepo) AddProjectRootInterface(projectId uint) (err error) {
	interf := model2.Interface{Name: "所有接口", ProjectId: projectId, IsDir: true}
	err = r.DB.Create(&interf).Error

	return
}
