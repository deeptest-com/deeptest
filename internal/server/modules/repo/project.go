package repo

import (
	"errors"
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/pkg/domain"
	_commUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	_fileUtils "github.com/deeptest-com/deeptest/pkg/lib/file"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	DB                         *gorm.DB                    `inject:""`
	RoleRepo                   *RoleRepo                   `inject:""`
	ProjectRoleRepo            *ProjectRoleRepo            `inject:""`
	EnvironmentRepo            *EnvironmentRepo            `inject:""`
	UserRepo                   *UserRepo                   `inject:""`
	ServeRepo                  *ServeRepo                  `inject:""`
	EndpointRepo               *EndpointRepo               `inject:""`
	EndpointInterfaceRepo      *EndpointInterfaceRepo      `inject:""`
	ProjectRecentlyVisitedRepo *ProjectRecentlyVisitedRepo `inject:""`
	ServeServerRepo            *ServeServerRepo            `inject:""`
	ScenarioRepo               *ScenarioRepo               `inject:""`
	ScenarioNodeRepo           *ScenarioNodeRepo           `inject:""`
	ScenarioProcessorRepo      *ScenarioProcessorRepo      `inject:""`
	PlanRepo                   *PlanRepo                   `inject:""`
	EndpointMockExpectRepo     *EndpointMockExpectRepo     `inject:""`
	CategoryRepo               *CategoryRepo               `inject:""`
	ScenarioInterfaceRepo      *ScenarioInterfaceRepo      `inject:""`
	EndpointCaseRepo           *EndpointCaseRepo           `inject:""`
	DebugInterfaceRepo         *DebugInterfaceRepo         `inject:""`
	*BaseRepo                  `inject:""`
}

func (r *ProjectRepo) Paginate(tenantId consts.TenantId, req v1.ProjectReqPaginate, userId uint) (data _domain.PageData, err error) {
	var count int64
	var projectIds []uint
	r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Select("project_id").Where("user_id = ?", userId).Scan(&projectIds)

	db := r.GetDB(tenantId).Model(&model.Project{}).Where("NOT deleted AND id IN (?)", projectIds)

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("disabled = ?", _commUtils.IsDisable(req.Enabled))
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
		user, _ := r.UserRepo.FindById(tenantId, project.AdminId)
		projects[key].AdminName = user.Name
	}

	data.Populate(projects, count, req.Page, req.PageSize)

	return
}

func (r *ProjectRepo) Get(tenantId consts.TenantId, id uint) (project model.Project, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).
		Where("id = ?", id).
		First(&project).Error

	return
}

func (r *ProjectRepo) GetByName(tenantId consts.TenantId, projectName string, id uint) (project model.Project, err error) {
	db := r.GetDB(tenantId).Model(&model.Project{}).
		Where("name = ? AND NOT deleted AND NOT disabled", projectName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	err = db.First(&project).Error

	return
}

func (r *ProjectRepo) GetByCode(tenantId consts.TenantId, shortName string, id uint) (ret model.Project, err error) {
	db := r.GetDB(tenantId).Model(&ret).
		Where("short_name = ? AND NOT deleted", shortName)

	if id > 0 {
		db.Where("id != ?", id)
	}
	err = db.First(&ret).Error

	return
}

func (r *ProjectRepo) GetBySpec(tenantId consts.TenantId, spec string) (project model.Project, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).
		Where("spec = ?", spec).
		First(&project).Error

	return
}

func (r *ProjectRepo) Save(tenantId consts.TenantId, po *model.Project) (err error) {
	err = r.GetDB(tenantId).Save(po).Error

	return
}

func (r *ProjectRepo) Create(tenantId consts.TenantId, req v1.ProjectReq, userId uint) (id uint, bizErr _domain.BizErr) {
	po, err := r.GetByName(tenantId, req.Name, 0)
	if po.Name != "" {
		bizErr = _domain.ErrNameExist
		return
	}

	po, err = r.GetByCode(tenantId, req.ShortName, 0)
	if po.ShortName != "" {
		bizErr = _domain.ErrShortNameExist
		return
	}

	// create project
	project := model.Project{ProjectBase: req.ProjectBase}
	err = r.GetDB(tenantId).Model(&model.Project{}).Create(&project).Error
	if err != nil {
		logUtils.Errorf("add project error", zap.String("error:", err.Error()))
		bizErr = _domain.SystemErr

		return
	}
	if req.AdminId != userId {
		err = r.AddProjectMember(tenantId, project.ID, req.AdminId, r.BaseRepo.GetAdminRoleName())
		if err != nil {
			logUtils.Errorf("添加项目角色错误", zap.String("错误:", err.Error()))
			bizErr = _domain.SystemErr
			return 0, bizErr
		}
	}
	err = r.CreateProjectRes(tenantId, project.ID, userId, req.IncludeExample)

	id = project.ID

	return
}

func (r *ProjectRepo) CreateProjectRes(tenantId consts.TenantId, projectId, userId uint, IncludeExample bool) (err error) {

	// create project member
	err = r.AddProjectMember(tenantId, projectId, userId, r.BaseRepo.GetAdminRoleName())
	if err != nil {
		logUtils.Errorf("添加项目角色错误", zap.String("错误:", err.Error()))
		return
	}

	// create project environment
	err = r.EnvironmentRepo.AddDefaultForProject(tenantId, projectId)
	if err != nil {
		logUtils.Errorf("添加项目默认环境错误", zap.String("错误:", err.Error()))
		return
	}

	// create project serve
	serve, err := r.AddProjectDefaultServe(tenantId, projectId, userId)
	if err != nil {
		logUtils.Errorf("添加默认服务错误", zap.String("错误:", err.Error()))
		return
	}

	// create project endpoint category
	categoryId, err := r.AddProjectRootEndpointCategory(tenantId, projectId)
	if err != nil {
		logUtils.Errorf("添加终端分类错误", zap.String("错误:", err.Error()))
		return
	}

	// create project test category
	/*
		err = r.ServeRepo.AddDefaultTestCategory(serve.ProjectId)
		if err != nil {
			logUtils.Errorf("添加终端分类错误", zap.String("错误:", err.Error()))
			return
		}
	*/

	// create project scenario category
	err = r.AddProjectRootScenarioCategory(tenantId, projectId)
	if err != nil {
		logUtils.Errorf("添加场景分类错误", zap.String("错误:", err.Error()))
		return
	}

	// create project plan category
	err = r.AddProjectRootPlanCategory(tenantId, projectId)
	if err != nil {
		logUtils.Errorf("添加计划分类错误", zap.String("错误:", err.Error()))
		return
	}

	// create project schema category
	err = r.AddProjectRootSchemaCategory(tenantId, projectId)
	if err != nil {
		logUtils.Errorf("添加组件分类错误", zap.String("错误:", err.Error()))
		return
	}

	//create sample
	if IncludeExample {
		err = r.CreateSample(tenantId, projectId, serve.ID, userId, categoryId)
		if err != nil {
			logUtils.Errorf("创建示例失败", zap.String("错误:", err.Error()))
			return
		}
	}

	return
}

func (r *ProjectRepo) Update(tenantId consts.TenantId, req v1.ProjectReq) error {
	po, _ := r.GetByName(tenantId, req.Name, req.Id)
	if po.Name != "" {
		return errors.New("同名记录已存在")
	}

	po, _ = r.GetByCode(tenantId, req.ShortName, req.Id)
	if po.ShortName != "" {
		return errors.New("英文缩写已存在")
	}

	project := model.Project{ProjectBase: req.ProjectBase}
	err := r.GetDB(tenantId).Model(&model.Project{}).Where("id = ?", req.Id).Updates(&project).Error
	if err != nil {
		logUtils.Errorf("update project error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectRepo) UpdateDefaultEnvironment(tenantId consts.TenantId, projectId, envId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).
		Where("id = ?", projectId).
		Updates(map[string]interface{}{"environment_id": envId}).Error

	if err != nil {
		logUtils.Errorf("update project environment error", err.Error())
		return err
	}

	return
}

func (r *ProjectRepo) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).Where("id = ?", id).
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

func (r *ProjectRepo) GetChildrenIds(tenantId consts.TenantId, id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE project AS (
			SELECT * FROM biz_project WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_project child, project WHERE child.parent_id = project.id
		)
		SELECT id FROM project WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.GetDB(tenantId).Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children project error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ProjectRepo) ListProjectByUser(tenantId consts.TenantId, userId uint) (res []model.ProjectMemberRole, err error) {
	projectRoleMap, err := r.GetProjectRoleMapByUser(tenantId, userId)

	if err != nil {
		return
	}

	projectIds := make([]uint, 0)
	for k, _ := range projectRoleMap {
		projectIds = append(projectIds, k)
	}

	projects, err := r.GetProjectsByIds(tenantId, projectIds)
	if err != nil {
		return
	}

	res, err = r.CombineRoleForProject(tenantId, projects, projectRoleMap)

	if err != nil {
		return
	}

	//db := r.GetDB(tenantId).Model(&model.ProjectMember{}).
	//	Joins("LEFT JOIN biz_project p ON biz_project_member.project_id=p.id").
	//	Joins("LEFT JOIN biz_project_role r ON biz_project_member.project_role_id=r.id").
	//	Select("p.*, r.id role_id, r.name role_name").
	//	Where("NOT biz_project_member.deleted")
	//
	//if !isAdminUser {
	//	db.Where("biz_project_member.user_id = ?", userId)
	//}
	//err = db.Group("biz_project_member.project_id").Find(&projects).Error
	return
}

func (r *ProjectRepo) GetProjectRoleMapByUser(tenantId consts.TenantId, userId uint) (res map[uint]uint, err error) {
	isAdminUser, err := r.UserRepo.IsAdminUser(tenantId, userId)
	if err != nil {
		return
	}

	var projectMembers []model.ProjectMember
	db := r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Select("project_id, project_role_id")
	if !isAdminUser {
		db.Where("user_id = ?", userId)
	}
	if err = db.Find(&projectMembers).Error; err != nil {
		return
	}

	res = make(map[uint]uint)
	for _, v := range projectMembers {
		res[v.ProjectId] = v.ProjectRoleId
	}

	return
}

func (r *ProjectRepo) GetProjectsByIds(tenantId consts.TenantId, ids []uint) (projects []model.Project, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).
		Where("id IN (?) AND NOT deleted AND NOT disabled ", ids).
		Find(&projects).Error
	return
}

func (r *ProjectRepo) CombineRoleForProject(tenantId consts.TenantId, projects []model.Project, projectRoleMap map[uint]uint) (res []model.ProjectMemberRole, err error) {
	roleIds := make([]uint, 0)
	for _, v := range projectRoleMap {
		roleIds = append(roleIds, v)
	}
	roleIds = _commUtils.ArrayRemoveUintDuplication(roleIds)

	roleIdNameMap, err := r.ProjectRoleRepo.GetRoleIdNameMap(tenantId, roleIds)
	if err != nil {
		return
	}

	for _, v := range projects {
		projectMemberRole := model.ProjectMemberRole{
			Project: v,
		}
		if roleId, ok := projectRoleMap[v.ID]; ok {
			projectMemberRole.RoleId = roleId
		}
		if projectMemberRole.RoleId == 0 {
			continue
		}
		if roleName, ok := roleIdNameMap[projectMemberRole.RoleId]; ok {
			projectMemberRole.RoleName = roleName
		}
		res = append(res, projectMemberRole)
	}

	return
}

func (r *ProjectRepo) GetCurrProjectByUser(tenantId consts.TenantId, userId uint) (currProject model.Project, err error) {
	var user model.SysUser
	err = r.GetDB(tenantId).Preload("Profile").
		Where("id = ?", userId).
		First(&user).
		Error

	err = r.GetDB(tenantId).Model(&model.Project{}).
		Where("id = ?", user.Profile.CurrProjectId).
		First(&currProject).Error

	return
}

func (r *ProjectRepo) ListProjectsRecentlyVisited(tenantId consts.TenantId, userId uint) (projects []model.Project, err error) {
	err = r.GetDB(tenantId).Raw(fmt.Sprintf("SELECT p.*,max( v.created_at ) visited_time FROM biz_project_recently_visited v,biz_project p WHERE v.project_id = p.id AND v.user_id = %d AND NOT p.deleted GROUP BY v.project_id ORDER BY visited_time DESC LIMIT 3", userId)).Find(&projects).Error
	return
}

func (r *ProjectRepo) ChangeProject(tenantId consts.TenantId, projectId, userId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.SysUserProfile{}).Where("user_id = ?", userId).
		Updates(map[string]interface{}{"curr_project_id": projectId}).Error

	return
}

func (r *ProjectRepo) AddProjectMember(tenantId consts.TenantId, projectId, userId uint, role consts.RoleType) (err error) {
	var projectRole model.ProjectRole
	projectRole, err = r.ProjectRoleRepo.FindByName(tenantId, role)
	if err != nil {
		return
	}

	projectMember := model.ProjectMember{UserId: userId, ProjectId: projectId, ProjectRoleId: projectRole.ID}
	err = r.GetDB(tenantId).Create(&projectMember).Error

	return
}

func (r *ProjectRepo) AddProjectRootEndpointCategory(tenantId consts.TenantId, projectId uint) (id uint, err error) {
	root := model.Category{
		Name:      "所有API",
		Type:      serverConsts.EndpointCategory,
		ProjectId: projectId,
		IsDir:     true,
	}
	err = r.GetDB(tenantId).Create(&root).Error

	return root.ID, err
}

func (r *ProjectRepo) AddProjectRootScenarioCategory(tenantId consts.TenantId, projectId uint) (err error) {
	root := model.Category{
		Name:      "分类",
		Type:      serverConsts.ScenarioCategory,
		ProjectId: projectId,
		IsDir:     true,
	}
	err = r.GetDB(tenantId).Create(&root).Error

	return
}

func (r *ProjectRepo) AddProjectRootPlanCategory(tenantId consts.TenantId, projectId uint) (err error) {
	root := model.Category{
		Name:      "分类",
		Type:      serverConsts.PlanCategory,
		ProjectId: projectId,
		IsDir:     true,
	}
	err = r.GetDB(tenantId).Create(&root).Error

	return
}

func (r *ProjectRepo) AddProjectRootSchemaCategory(tenantId consts.TenantId, projectId uint) (err error) {
	root := model.Category{
		Name:      "分类",
		Type:      serverConsts.SchemaCategory,
		ProjectId: projectId,
		IsDir:     true,
	}
	err = r.GetDB(tenantId).Create(&root).Error

	return
}

func (r *ProjectRepo) AddProjectRootTestCategory(tenantId consts.TenantId, projectId, serveId uint) (err error) {

	root := model.DiagnoseInterface{
		Title:     "根节点",
		ProjectId: projectId,
		IsDir:     true,
		Type:      "dir",
		ServeId:   serveId,
	}
	err = r.GetDB(tenantId).Create(&root).Error

	return
}

func (r *ProjectRepo) Members(tenantId consts.TenantId, req v1.ProjectReqPaginate, projectId int) (data _domain.PageData, err error) {
	req.Order = "sys_user.created_at"
	db := r.GetDB(tenantId).Model(&model.SysUser{}).
		Select("sys_user.id, sys_user.username, sys_user.email,sys_user.name, m.project_role_id, r.name as role_name").
		Joins("left join biz_project_member m on sys_user.id=m.user_id").
		Joins("left join biz_project_role r on m.project_role_id=r.id").
		Where("m.project_id = ?", projectId)
	if req.Keywords != "" {
		db = db.Where("sys_user.username LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	if config.CONFIG.System.SysEnv == "ly" {
		db = db.Where("sys_user.username != ?", serverConsts.AdminUserName)
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

func (r *ProjectRepo) RemoveMember(tenantId consts.TenantId, userId, projectId int) (err error) {
	/*
		err = r.GetDB(tenantId).Model(&modelRef.ProjectMember{}).
			Where("user_id = ? AND project_id = ?", userId, projectId).
			Updates(map[string]interface{}{"deleted": true}).Error
		if err != nil {
			return
		}
	*/
	err = r.GetDB(tenantId).
		Where("user_id = ? AND project_id=?", userId, projectId).
		Delete(&model.ProjectMember{}).Error

	return
}

func (r *ProjectRepo) FindRolesByUser(tenantId consts.TenantId, userId uint) (members []model.ProjectMember, err error) {

	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Joins("LEFT JOIN biz_project_role r ON biz_project_member.project_role_id=r.id").
		Select("biz_project_member.*, r.name project_role_name").
		Where("biz_project_member.user_id = ?", userId).
		Find(&members).Error

	return
}

func (r *ProjectRepo) GetProjectsAndRolesByUser(tenantId consts.TenantId, userId uint) (projectIds, roleIds []uint) {
	var members []model.ProjectMember
	r.GetDB(tenantId).Model(&model.ProjectMember{}).
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

func (r *ProjectRepo) AddProjectDefaultServe(tenantId consts.TenantId, projectId, userId uint) (serve model.Serve, err error) {
	serve = model.Serve{
		Name:      "默认服务",
		ProjectId: projectId,
	}

	err = r.GetDB(tenantId).Create(&serve).Error

	r.ServeRepo.SetCurrServeByUser(tenantId, serve.ID, userId)

	r.ServeRepo.AddDefaultServer(tenantId, serve.ProjectId, serve.ID)

	//调试目录不挂在目录下面
	//	r.ServeRepo.AddDefaultTestCategory(serve.ProjectId, serve.ID)

	return
}

func (r *ProjectRepo) FindRolesByProjectAndUser(tenantId consts.TenantId, projectId, userId uint) (projectMember model.ProjectMember, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Where("project_id = ?", projectId).
		Where("user_id = ?", userId).
		Scan(&projectMember).Error
	return
}

func (r *ProjectRepo) UpdateUserRole(tenantId consts.TenantId, req v1.UpdateProjectMemberReq) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Where("project_id = ?", req.ProjectId).
		Where("user_id = ?", req.UserId).
		Updates(map[string]interface{}{"project_role_id": req.ProjectRoleId}).Error

	if err != nil {
		logUtils.Errorf("update project user role error", err.Error())
		return err
	}

	return
}

func (r *ProjectRepo) GetCurrProjectMemberRoleByUser(tenantId consts.TenantId, userId uint) (ret model.ProjectMember, err error) {
	curProject, err := r.GetCurrProjectByUser(tenantId, userId)
	if err != nil {
		return
	}
	if curProject.ID == 0 {
		return ret, errors.New("current project is not existed")
	}
	return r.FindRolesByProjectAndUser(tenantId, curProject.ID, userId)
}

func (r *ProjectRepo) GetMembersByProject(tenantId consts.TenantId, projectId uint) (ret []model.ProjectMember, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Where("project_id = ?", projectId).
		Find(&ret).Error
	return
}

func (r *ProjectRepo) GetAuditList(tenantId consts.TenantId, req v1.AuditProjectPaginate) (data _domain.PageData, err error) {
	req.Field = "status asc,created_at"

	var count int64
	db := r.GetDB(tenantId).Model(&model.ProjectMemberAudit{})
	if req.Type == 0 {
		projectIds := r.GetProjectIdsByUserIdAndRole(tenantId, req.AuditUserId, r.BaseRepo.GetAdminRoleName())
		db = db.Where("project_id in ? and status = 0", projectIds)
	} else {
		db = db.Where("apply_user_id = ?", req.ApplyUserId)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count ProjectMemberAudit error", zap.String("error:", err.Error()))
		return
	}

	list := make([]*model.ProjectMemberAudit, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&list).Error
	if err != nil {
		logUtils.Errorf("query ProjectMemberAudit error", zap.String("error:", err.Error()))
		return
	}

	r.refUserName(tenantId, list)
	r.refProjectName(tenantId, list)

	data.Populate(list, count, req.Page, req.PageSize)

	return
}

func (r *ProjectRepo) refUserName(tenantId consts.TenantId, list []*model.ProjectMemberAudit) {
	names := make(map[uint]string)
	for key, item := range list {
		if _, ok := names[item.ApplyUserId]; !ok {
			user, _ := r.UserRepo.FindById(tenantId, item.ApplyUserId)
			names[item.ApplyUserId] = user.Name

		}
		if _, ok := names[item.AuditUserId]; !ok {
			user, _ := r.UserRepo.FindById(tenantId, item.AuditUserId)
			names[item.AuditUserId] = user.Name
		}

		list[key].AuditUserName = names[item.AuditUserId]
		list[key].ApplyUserName = names[item.ApplyUserId]

	}
}

func (r *ProjectRepo) refProjectName(tenantId consts.TenantId, list []*model.ProjectMemberAudit) {
	names := make(map[uint]string)
	for key, item := range list {
		if _, ok := names[item.ProjectId]; !ok {
			project, _ := r.Get(tenantId, item.ProjectId)
			names[item.ProjectId] = project.Name
		}

		list[key].ProjectName = names[item.ProjectId]
	}
}

func (r *ProjectRepo) GetAudit(tenantId consts.TenantId, id uint) (ret model.ProjectMemberAudit, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMemberAudit{}).
		Where("id = ?", id).
		First(&ret).Error
	return
}

func (r *ProjectRepo) UpdateAuditStatus(tenantId consts.TenantId, id, auditUserId uint, status consts.AuditStatus) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMemberAudit{}).
		Where("id=?", id).
		Updates(map[string]interface{}{"status": status, "audit_user_id": auditUserId}).Error
	return
}

func (r *ProjectRepo) SaveAudit(tenantId consts.TenantId, audit model.ProjectMemberAudit) (auditId uint, err error) {
	err = r.GetDB(tenantId).Save(&audit).Error
	auditId = audit.ID
	return
}

func (r *ProjectRepo) IfProjectMember(tenantId consts.TenantId, userId, projectId uint) (res bool, err error) {
	var count int64
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Where("user_id=? and project_id=?", userId, projectId).Count(&count).Error
	if err != nil {
		return
	}
	res = count > 0
	return
}

func (r *ProjectRepo) CreateSample(tenantId consts.TenantId, projectId, serveId, userId, categoryId uint) (err error) {
	//创建目录
	var category model.Category
	categoryJson := _fileUtils.ReadFile("./config/sample/category.json")
	_commUtils.JsonDecode(categoryJson, &category)

	var components []*model.ComponentSchema
	componentJson := _fileUtils.ReadFile("./config/sample/component.json")
	_commUtils.JsonDecode(componentJson, &components)

	//获取接口配置
	var endpoints []model.Endpoint
	endpointJson := _fileUtils.ReadFile("./config/sample/endpoint.json")
	_commUtils.JsonDecode(endpointJson, &endpoints)

	endpointMockExpectsMap := make(map[string][]model.EndpointMockExpect)
	endpointMockExpectsJson := _fileUtils.ReadFile("./config/sample/endpoint-mock-expect.json")
	_commUtils.JsonDecode(endpointMockExpectsJson, &endpointMockExpectsMap)

	endpointCaseMap := make(map[string][]map[string]model.DebugInterface)
	endpointCaseJson := _fileUtils.ReadFile("./config/sample/endpoint-case.json")
	_commUtils.JsonDecode(endpointCaseJson, &endpointCaseMap)

	user, _ := r.UserRepo.FindById(tenantId, userId)

	//获取场景配置
	var scenario model.Scenario
	scenarioJson := _fileUtils.ReadFile("./config/sample/scenario.json")
	_commUtils.JsonDecode(scenarioJson, &scenario)

	//获取执行器
	var root agentExec.Processor
	processorJson := _fileUtils.ReadFile("./config/sample/processor.json")
	_commUtils.JsonDecode(processorJson, &root)

	var processorEntity map[string]interface{}
	processorEntityJson := _fileUtils.ReadFile("./config/sample/processor-entity.json")
	_commUtils.JsonDecode(processorEntityJson, &processorEntity)

	var plan model.Plan
	planJson := _fileUtils.ReadFile("./config/sample/plan.json")
	_commUtils.JsonDecode(planJson, &plan)

	return r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {

		for _, component := range components {
			component.ServeId = int64(serveId)
			component.Ref = "#/components/schemas/" + component.Name
		}
		err := r.ServeRepo.SaveSchemas(tenantId, components)
		if err != nil {
			return err
		}

		category.ProjectId, category.ServeId, category.ParentId = projectId, serveId, int(categoryId)
		err = r.CategoryRepo.Save(tenantId, &category)
		if err != nil {
			return err
		}

		//创建接口
		interfaceIds := map[string]model.EndpointInterface{}
		for _, endpoint := range endpoints {
			endpoint.ServeId = serveId
			endpoint.ProjectId = projectId
			endpoint.CreateUser = user.Username
			endpoint.CategoryId = int64(category.ID)
			err = r.EndpointRepo.SaveAll(tenantId, &endpoint)
			if err != nil {
				return err
			}
			interfaceIds[endpoint.Interfaces[0].Name] = endpoint.Interfaces[0]
		}

		r.ServeServerRepo.SetUrl(tenantId, serveId, "http://192.168.5.224:50400")

		// 创建接口用例
		for endpointName, caseDebugs := range endpointCaseMap {
			endpoint, err := r.EndpointRepo.GetByNameAndProject(tenantId, endpointName, projectId)
			if err != nil {
				return err
			}

			for _, caseDebug := range caseDebugs {
				for caseName, debug := range caseDebug {
					endpointCase := model.EndpointCase{
						Name:           caseName,
						EndpointId:     endpoint.ID,
						ServeId:        serveId,
						ProjectId:      projectId,
						CreateUserId:   userId,
						CreateUserName: user.Username,
					}
					if err = r.EndpointCaseRepo.Save(tenantId, &endpointCase); err != nil {
						return err
					}

					debug.ProjectId = projectId
					debug.EndpointInterfaceId = interfaceIds[endpoint.Title].ID
					debug.ServeId = serveId
					debug.CaseInterfaceId = endpointCase.ID
					r.DebugInterfaceRepo.Save(tenantId, &debug)

					if err = r.EndpointCaseRepo.UpdateDebugInterfaceId(tenantId, debug.ID, endpointCase.ID); err != nil {
						return err
					}
				}
			}
		}

		// 创建Mock期望
		for endpointName, mockExpects := range endpointMockExpectsMap {
			for _, mockExpect := range mockExpects {
				if endpointInterface, ok := interfaceIds[endpointName]; ok {
					mockExpect.EndpointId = endpointInterface.EndpointId
					mockExpect.EndpointInterfaceId = endpointInterface.ID
					mockExpect.Method = endpointInterface.Method
					mockExpect.CreateUser = user.Username
					_, err = r.EndpointMockExpectRepo.Save(tenantId, mockExpect)
					if err != nil {
						return err
					}
				}
			}
		}

		//创建场景目录
		ScenarioCategory, err := r.CategoryRepo.GetByItem(tenantId, 0, serverConsts.ScenarioCategory, projectId, "分类")
		if err != nil {
			return err
		}
		ScenarioCategory.ParentId = int(ScenarioCategory.ID)
		ScenarioCategory.ID = 0
		ScenarioCategory.Name = "宠物商店"
		err = r.CategoryRepo.Save(tenantId, &ScenarioCategory)
		if err != nil {
			return err
		}

		//TODO 创建场景
		scenario.ProjectId = projectId
		scenario.CategoryId = int64(ScenarioCategory.ID)
		scenario.Status = consts.Draft
		scenario.CreateUserId = userId
		scenario.CreateUserName = user.Username
		scenario, err = r.ScenarioRepo.Create(tenantId, scenario)
		if err != nil {
			return err
		}

		//TODO 添加执行器
		err = r.createProcessorTree(tenantId, &root, interfaceIds, processorEntity, projectId, scenario.ID, 0, userId, serveId)
		if err != nil {
			return err
		}
		//TODO 创建计划
		plan.ProjectId = projectId
		plan.Status = consts.Draft
		plan.CreateUserId = userId
		plan, _ = r.PlanRepo.Create(tenantId, plan)

		//关联场景
		err = r.PlanRepo.AddScenarios(tenantId, plan.ID, []uint{scenario.ID})
		if err != nil {
			return err
		}

		return nil
	})

}

func (r *ProjectRepo) GetProjectIdsByUserIdAndRole(tenantId consts.TenantId, userId uint, roleName consts.RoleType) (projectIds []uint) {
	var projects []model.ProjectMember
	err := r.GetDB(tenantId).Model(model.ProjectMember{}).
		Joins("LEFT JOIN biz_project_role r ON biz_project_member.project_role_id=r.id").
		Where("biz_project_member.user_id=? and r.name=? and not biz_project_member.deleted and not biz_project_member.disabled", userId, roleName).
		Find(&projects).Error
	if err != nil {
		return
	}
	for _, project := range projects {
		projectIds = append(projectIds, project.ProjectId)
	}
	return
}

func (r *ProjectRepo) createProcessorTree(tenantId consts.TenantId, root *agentExec.Processor, interfaceIds map[string]model.EndpointInterface, processorEntity map[string]interface{}, projectId, scenarioId, parentId, userId, serveId uint) error {
	processor := model.Processor{
		Name:                  root.Name,
		EntityCategory:        root.EntityCategory,
		EntityType:            root.EntityType,
		EndpointInterfaceId:   interfaceIds[root.Name].ID,
		ParentId:              parentId,
		ScenarioId:            scenarioId,
		ProjectId:             projectId,
		ProcessorInterfaceSrc: root.ProcessorInterfaceSrc,
		CreatedBy:             userId,
	}
	processor.Ordr = r.ScenarioNodeRepo.GetMaxOrder(tenantId, processor.ParentId)
	err := r.ScenarioNodeRepo.Save(tenantId, &processor)
	if err != nil {
		return err
	}

	processorCategory := root.EntityCategory
	if item, ok := processorEntity[root.Name]; ok {
		if processorCategory == consts.ProcessorGroup {
			var entity model.ProcessorGroup

			//将 map 转换为指定的结构体
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveGroup(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)

		} else if processorCategory == consts.ProcessorLogic {
			var entity model.ProcessorLogic
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveLogic(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)

		} else if processorCategory == consts.ProcessorLoop {
			var entity model.ProcessorLoop
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveLoop(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)

		} else if processorCategory == consts.ProcessorTimer {
			var entity model.ProcessorTimer
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveTimer(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)

		} else if processorCategory == consts.ProcessorPrint {
			var entity model.ProcessorPrint
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SavePrint(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)
		} else if processorCategory == consts.ProcessorVariable {
			var entity model.ProcessorVariable
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveVariable(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)

		} else if processorCategory == consts.ProcessorCookie {
			var entity model.ProcessorCookie
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveCookie(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)

		} else if processorCategory == consts.ProcessorAssertion {
			var entity model.ProcessorAssertion
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveAssertion(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)

		} else if processorCategory == consts.ProcessorData {
			var entity model.ProcessorData
			_commUtils.Map2Struct(item, &entity)
			entity.ProcessorID = processor.ID
			entity.ParentID = parentId
			err = r.ScenarioProcessorRepo.SaveData(tenantId, &entity)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, entity.ID)
		} else if processorCategory == consts.ProcessorInterface {
			var debug model.DebugInterface
			_commUtils.Map2Struct(item, &debug)
			debug.EndpointInterfaceId = interfaceIds[root.Name].ID
			debug.ProjectId = projectId
			debug.ServeId = serveId
			debug.ScenarioProcessorId = processor.ID
			err = r.ScenarioInterfaceRepo.SaveDebugData(tenantId, &debug)
			r.ScenarioProcessorRepo.UpdateInterfaceId(tenantId, debug.ScenarioProcessorId, debug.ID)
			r.ScenarioProcessorRepo.UpdateMethod(tenantId, debug.ScenarioProcessorId, debug.Method)
			r.ScenarioNodeRepo.UpdateEntityId(tenantId, processor.ID, debug.ID)

		}

	}

	for _, child := range root.Children {
		r.createProcessorTree(tenantId, child, interfaceIds, processorEntity, projectId, scenarioId, processor.ID, userId, serveId)
	}

	return nil

}

func (r *ProjectRepo) GetAuditUsers(tenantId consts.TenantId, projectId uint) (users []model.SysUser, err error) {
	err = r.GetDB(tenantId).Model(model.SysUser{}).
		Joins("LEFT JOIN biz_project_member m ON m.user_id=sys_user.id").
		Joins("LEFT JOIN biz_project_role r ON m.project_role_id=r.id").
		Where("m.project_id=? and r.name=? and not m.deleted and not m.disabled", projectId, r.BaseRepo.GetAdminRoleName()).
		Find(&users).Error

	return
}

func (r *ProjectRepo) ListAll(tenantId consts.TenantId) (res []model.Project, err error) {
	err = r.GetDB(tenantId).Model(model.Project{}).
		Where("not disabled and not deleted").
		Find(&res).Error
	return
}

func (r *ProjectRepo) GetByShortName(tenantId consts.TenantId, shortName string) (project model.Project, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).
		Where("short_name = ? and not deleted", shortName).
		First(&project).Error

	return
}

func (r *ProjectRepo) GetUserIdsByProjectAnRole(tenantId consts.TenantId, projectId, roleId uint) (projectMembers []model.ProjectMember, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Where("project_id = ?", projectId).
		Where("project_role_id = ?", roleId).
		Find(&projectMembers).Error
	return
}

func (r *ProjectRepo) GetUsernamesByProjectAndRole(tenantId consts.TenantId, projectId, roleId uint, exceptUserName string) (imAccounts []string, err error) {
	conn := r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Joins("left join sys_user u on biz_project_member.user_id=u.id").
		Select("u.username")
	if projectId != 0 {
		conn = conn.Where("biz_project_member.project_id = ?", projectId)
	}
	if roleId != 0 {
		conn = conn.Where("biz_project_member.project_role_id = ?", roleId)
	}
	if exceptUserName != "" {
		conn = conn.Where("u.username != ?", exceptUserName)
	}
	err = conn.Where("not biz_project_member.deleted and not u.deleted").Find(&imAccounts).Error
	return
}

func (r *ProjectRepo) GetAuditByItem(tenantId consts.TenantId, projectId, ApplyUserId uint, auditStatus []consts.AuditStatus) (ret model.ProjectMemberAudit, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMemberAudit{}).
		Where("project_id = ? and apply_user_id = ? and status in ? ", projectId, ApplyUserId, auditStatus).
		Last(&ret).Error
	return
}

func (r *ProjectRepo) UpdateProjectSource(tenantId consts.TenantId, projectId uint, source serverConsts.ProjectSource) (err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).
		Where("id = ?", projectId).Update("source", source).Error
	return
}

func (r *ProjectRepo) ListByUsername(tenantId consts.TenantId, username string) (res []model.Project, err error) {
	err = r.GetDB(tenantId).Model(model.Project{}).
		Joins("LEFT JOIN biz_project_member m ON biz_project.id=m.project_id").
		Joins("LEFT JOIN sys_user u ON m.user_id=u.id").
		Where("u.username = ?", username).
		Where("not biz_project.disabled and not biz_project.deleted and not m.disabled and not m.deleted").
		Find(&res).Error
	return
}

func (r *ProjectRepo) BatchGetByShortNames(tenantId consts.TenantId, shortNames []string) (ret []model.Project, err error) {
	err = r.GetDB(tenantId).Model(&ret).
		Where("short_name IN (?) AND NOT deleted", shortNames).
		Find(&ret).Error

	return
}

func (r *ProjectRepo) AddMemberIfNotExisted(tenantId consts.TenantId, projectId, userId uint, role consts.RoleType) (err error) {
	isMember, err := r.IfProjectMember(tenantId, userId, projectId)
	if err != nil || isMember {
		return
	}

	err = r.AddProjectMember(tenantId, projectId, userId, role)
	return
}

func (r *ProjectRepo) FindRolesByProjectsAndUsername(tenantId consts.TenantId, username string, projectIds []uint) (members []model.ProjectMember, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Joins("LEFT JOIN biz_project_role r ON biz_project_member.project_role_id=r.id").
		Joins("LEFT JOIN sys_user u ON biz_project_member.user_id=u.id").
		Select("biz_project_member.*, r.name project_role_name").
		Where("u.username = ?", username).
		Where("biz_project_member.project_id IN (?)", projectIds).
		Find(&members).Error

	return
}

func (r *ProjectRepo) GetUserProjectRoleMap(tenantId consts.TenantId, username string, projectIds []uint) (res map[uint]consts.RoleType, err error) {
	projectRoles, err := r.FindRolesByProjectsAndUsername(tenantId, username, projectIds)
	if err != nil {
		return
	}

	res = make(map[uint]consts.RoleType)
	for _, v := range projectRoles {
		res[v.ProjectId] = v.ProjectRoleName
	}

	return
}

func (r *ProjectRepo) GetProjectMemberCount(tenantId consts.TenantId, projectId uint) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Where("project_id=?", projectId).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (r *ProjectRepo) GetProjectMemberList(tenantId consts.TenantId, projectId uint) (list []model.ProjectMember, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Where("project_id=?", projectId).Find(&list).Error
	if err != nil {
		return
	}
	return
}
