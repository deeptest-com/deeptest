package repo

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type ProjectRepo struct {
	DB              *gorm.DB         `inject:""`
	RoleRepo        *RoleRepo        `inject:""`
	ProjectRoleRepo *ProjectRoleRepo `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
	UserRepo        *UserRepo        `inject:""`
	ServeRepo       *ServeRepo       `inject:""`
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (r *ProjectRepo) Paginate(req v1.ProjectReqPaginate, userId uint) (data _domain.PageData, err error) {
	var count int64

	var projectIds []uint
	r.DB.Model(&model.ProjectMember{}).
		Select("project_id").Where("user_id = ?", userId).Scan(&projectIds)

	db := r.DB.Model(&model.Project{}).Where("NOT deleted AND id IN (?)", projectIds)

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

	for key, project := range projects {
		user, _ := r.UserRepo.FindById(project.AdminId)
		projects[key].AdminName = user.Name
	}

	data.Populate(projects, count, req.Page, req.PageSize)

	return
}

func (r *ProjectRepo) Get(id uint) (project model.Project, err error) {
	err = r.DB.Model(&model.Project{}).
		Where("id = ?", id).
		First(&project).Error

	return
}

func (r *ProjectRepo) GetByName(projectName string, id uint) (project v1.ProjectResp, err error) {
	db := r.DB.Model(&model.Project{}).
		Where("name = ?", projectName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&project)

	return
}

func (r *ProjectRepo) GetBySpec(spec string) (project model.Project, err error) {
	err = r.DB.Model(&model.Project{}).
		Where("spec = ?", spec).
		First(&project).Error

	return
}

func (r *ProjectRepo) Save(po *model.Project) (err error) {
	err = r.DB.Save(po).Error

	return
}

func (r *ProjectRepo) Create(req v1.ProjectReq, userId uint) (id uint, bizErr *_domain.BizErr) {
	po, err := r.GetByName(req.Name, 0)
	if po.Name != "" {
		bizErr.Code = _domain.ErrNameExist.Code
		return
	}

	// create project
	project := model.Project{ProjectBase: req.ProjectBase}
	err = r.DB.Model(&model.Project{}).Create(&project).Error
	if err != nil {
		logUtils.Errorf("add project error", zap.String("error:", err.Error()))
		bizErr.Code = _domain.SystemErr.Code

		return
	}

	r.CreateProjectRes(project.ID, userId)

	id = project.ID

	return
}

func (r *ProjectRepo) CreateProjectRes(projectId, userId uint) (err error) {

	// create project member
	err = r.AddProjectMember(projectId, userId, "admin")
	if err != nil {
		logUtils.Errorf("添加项目角色错误", zap.String("错误:", err.Error()))
		return
	}

	// create project environment
	err = r.EnvironmentRepo.AddDefaultForProject(projectId)
	if err != nil {
		logUtils.Errorf("添加项目默认环境错误", zap.String("错误:", err.Error()))
		return
	}

	// create project serve
	serve, err := r.AddProjectDefaultServe(projectId, userId)
	if err != nil {
		logUtils.Errorf("添加默认服务错误", zap.String("错误:", err.Error()))
		return
	}

	// create project interface category
	err = r.AddProjectRootInterface(serve.ID, projectId)
	if err != nil {
		logUtils.Errorf("添加接口错误", zap.String("错误:", err.Error()))
		return
	}

	// create project scenario category
	err = r.AddProjectRootScenarioCategory(projectId)
	if err != nil {
		logUtils.Errorf("添加场景分类错误", zap.String("错误:", err.Error()))
		return
	}

	// create project plan category
	err = r.AddProjectRootPlanCategory(projectId)
	if err != nil {
		logUtils.Errorf("添加场景分类错误", zap.String("错误:", err.Error()))
		return
	}

	return
}

func (r *ProjectRepo) Update(id uint, req v1.ProjectReq) error {
	project := model.Project{ProjectBase: req.ProjectBase}
	err := r.DB.Model(&model.Project{}).Where("id = ?", req.Id).Updates(&project).Error
	if err != nil {
		logUtils.Errorf("update project error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectRepo) UpdateDefaultEnvironment(projectId, envId uint) (err error) {
	err = r.DB.Model(&model.Project{}).
		Where("id = ?", projectId).
		Updates(map[string]interface{}{"environment_id": envId}).Error

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
		Where("NOT deleted AND id IN (?)", projectIds).
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

func (r *ProjectRepo) ListProjectsRecentlyVisited(userId uint) (projects []model.Project, err error) {
	//TODO 时间临时变更
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()-2, 0, 0, time.Local)

	err = r.DB.Model(&model.Project{}).Joins("LEFT JOIN biz_project_recently_visited v ON biz_project.id=v.project_id").
		Where("v.user_id = ?", userId).
		Where("v.created_at >= ?", date).
		Order("v.created_at desc").
		Find(&projects).Error

	return
}

func (r *ProjectRepo) ChangeProject(projectId, userId uint) (err error) {
	err = r.DB.Model(&model.SysUserProfile{}).Where("user_id = ?", userId).
		Updates(map[string]interface{}{"curr_project_id": projectId}).Error

	return
}

func (r *ProjectRepo) AddProjectMember(projectId, userId uint, role consts.RoleType) (err error) {
	var projectRole model.ProjectRole
	projectRole, err = r.ProjectRoleRepo.FindByName(role)
	if err != nil {
		return
	}

	projectMember := model.ProjectMember{UserId: userId, ProjectId: projectId, ProjectRoleId: projectRole.ID}
	err = r.DB.Create(&projectMember).Error

	return
}

func (r *ProjectRepo) AddProjectRootInterface(serveId, projectId uint) (err error) {
	root := model.Category{
		Name:      "分类",
		Type:      serverConsts.EndpointCategory,
		ServeId:   serveId,
		ProjectId: projectId,
		IsLeaf:    false,
	}
	err = r.DB.Create(&root).Error

	return
}

func (r *ProjectRepo) AddProjectRootScenarioCategory(projectId uint) (err error) {
	root := model.Category{
		Name:      "分类",
		Type:      serverConsts.ScenarioCategory,
		ProjectId: projectId,
		IsLeaf:    false,
	}
	err = r.DB.Create(&root).Error

	return
}

func (r *ProjectRepo) AddProjectRootPlanCategory(projectId uint) (err error) {
	root := model.Category{
		Name:      "分类",
		Type:      serverConsts.PlanCategory,
		ProjectId: projectId,
		IsLeaf:    false,
	}
	err = r.DB.Create(&root).Error

	return
}

func (r *ProjectRepo) Members(req v1.ProjectReqPaginate, projectId int) (data _domain.PageData, err error) {
	req.Order = "sys_user.created_at"
	db := r.DB.Model(&model.SysUser{}).
		Select("sys_user.id, sys_user.username, sys_user.email, m.project_role_id, r.name").
		Joins("left join biz_project_member m on sys_user.id=m.user_id").
		Joins("left join biz_project_role r on m.project_role_id=r.id").
		Where("m.project_id = ?", projectId)
	if req.Keywords != "" {
		db = db.Where("sys_user.username LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count users error", zap.String("error:", err.Error()))
		return
	}

	users := make([]v1.MemberResp, 0)
	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, "", req.Order)).
		Scan(&users).Error
	if err != nil {
		logUtils.Errorf("query users error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(users, count, req.Page, req.PageSize)

	return
}

func (r *ProjectRepo) RemoveMember(userId, projectId int) (err error) {
	/*
		err = r.DB.Model(&model.ProjectMember{}).
			Where("user_id = ? AND project_id = ?", userId, projectId).
			Updates(map[string]interface{}{"deleted": true}).Error
		if err != nil {
			return
		}
	*/
	err = r.DB.
		Where("user_id = ? AND project_id=?", userId, projectId).
		Delete(&model.ProjectMember{}).Error

	return
}

func (r *ProjectRepo) FindRolesByUser(userId uint) (ret []model.ProjectMember, err error) {
	var members []model.ProjectMember

	r.DB.Model(&model.ProjectMember{}).
		Where("user_id = ?", userId).
		Find(&members)

	for _, member := range members {
		projectRole, _ := r.ProjectRoleRepo.FindById(member.ProjectRoleId)

		member.ProjectRoleName = projectRole.Name
		ret = append(ret, member)
	}

	return
}

func (r *ProjectRepo) GetProjectsAndRolesByUser(userId uint) (projectIds, roleIds []uint) {
	var members []model.ProjectMember
	r.DB.Model(&model.ProjectMember{}).
		Where("user_id = ?", userId).
		Find(&members)

	roleIdsMap := make(map[uint]uint)

	for _, member := range members {
		projectIds = append(projectIds, member.ProjectId)
		roleIdsMap[member.ProjectRoleId] = member.ProjectRoleId
	}
	for _, v := range roleIdsMap {
		roleIds = append(roleIds, v)
	}

	return
}

func (r *ProjectRepo) AddProjectDefaultServe(projectId, userId uint) (serve model.Serve, err error) {
	po := model.Serve{
		Name:      "默认服务",
		ProjectId: projectId,
	}

	err = r.DB.Create(&po).Error

	r.ServeRepo.SetCurrServeByUser(po.ID, userId)

	return
}

func (r *ProjectRepo) FindRolesByProjectAndUser(projectId, userId uint) (projectMember model.ProjectMember, err error) {
	err = r.DB.Model(&model.ProjectMember{}).
		Where("project_id = ?", projectId).
		Where("user_id = ?", userId).
		Scan(&projectMember).Error
	return
}

func (r *ProjectRepo) UpdateUserRole(req v1.UpdateProjectMemberReq) (err error) {
	err = r.DB.Model(&model.ProjectMember{}).
		Where("project_id = ?", req.ProjectId).
		Where("user_id = ?", req.UserId).
		Updates(map[string]interface{}{"project_role_id": req.ProjectRoleId}).Error

	if err != nil {
		logUtils.Errorf("update project user role error", err.Error())
		return err
	}

	return
}

func (r *ProjectRepo) GetCurrProjectMemberRoleByUser(userId uint) (ret model.ProjectMember, err error) {
	curProject, err := r.GetCurrProjectByUser(userId)
	if err != nil {
		return
	}
	if curProject.ID == 0 {
		return ret, errors.New("current project is not existed")
	}
	return r.FindRolesByProjectAndUser(curProject.ID, userId)
}

func (r *ProjectRepo) GetMembersByProject(projectId uint) (ret []model.ProjectMember, err error) {
	err = r.DB.Model(&model.ProjectMember{}).
		Where("project_id = ?", projectId).
		Find(&ret).Error
	return
}
