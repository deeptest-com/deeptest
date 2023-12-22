package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"time"
)

type ServeRepo struct {
	*BaseRepo       `inject:""`
	CategoryRepo    *CategoryRepo    `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
}

func (r *ServeRepo) ListVersion(serveId uint) (res []model.ServeVersion, err error) {
	err = r.DB.Where("serve_id = ? AND NOT deleted AND not disabled", serveId).Find(&res).Error
	return
}

func (r *ServeRepo) ListByProject(projectId uint) (pos []model.Serve, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *ServeRepo) ListByProjects(projectIds []uint) (pos []model.Serve, err error) {
	err = r.DB.
		Where("project_id IN (?)", projectIds).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *ServeRepo) Paginate(req v1.ServeReqPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.Serve{}).Where("project_id = ? AND NOT deleted AND NOT disabled", req.ProjectId)

	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.Serve, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}

	//关联环境
	for key, result := range results {
		results[key].Servers, err = r.ListServer(result.ID)
		if err != nil {
			return
		}

	}

	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ServeRepo) PaginateVersion(req v1.ServeVersionPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.ServeVersion{}).Where("serve_id = ? AND NOT deleted AND NOT disabled", req.ServeId)

	if req.Version != "" {
		db = db.Where("value LIKE ?", fmt.Sprintf("%%%s%%", req.Version))
	}

	if req.CreateUser != "" {
		db = db.Where("creat_user = ?", fmt.Sprintf("%s", req.CreateUser))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.ServeVersion, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ServeRepo) PaginateSchema(req v1.ServeSchemaPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.ComponentSchema{}).Where("serve_id = ? AND NOT deleted AND NOT disabled", req.ServeId)

	if req.Type != "" {
		db.Where("type=?", req.Type)
	}
	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))

	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.ComponentSchema, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ServeRepo) PaginateSecurity(req v1.ServeSecurityPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.ComponentSchemaSecurity{}).Where("serve_id = ? AND NOT deleted AND NOT disabled", req.ServeId)

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.ComponentSchemaSecurity, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ServeRepo) Get(id uint) (res model.Serve, err error) {
	//err = r.DB.Where("NOT deleted AND not disabled").First(&res, id).Error
	err = r.DB.Where("NOT deleted").First(&res, id).Error
	return
}

func (r *ServeRepo) GetByCode(projectId uint, shortName string) (ret model.Serve, err error) {
	db := r.DB.Model(&ret).
		Where("short_name = ? AND NOT deleted", shortName)

	if projectId > 0 {
		db.Where("project_id = ?", projectId)
	}

	err = db.First(&ret).Error

	return
}

func (r *ServeRepo) GetSchema(id uint) (res model.ComponentSchema, err error) {
	err = r.DB.Where("NOT deleted AND not disabled").First(&res, id).Error
	return
}

func (r *ServeRepo) GetSchemasByServeId(serveId uint) (res []model.ComponentSchema, err error) {
	err = r.DB.Where("NOT deleted AND not disabled AND serve_id = ?", serveId).Find(&res).Error
	return
}

func (r *ServeRepo) DeleteById(id uint) error {
	return r.DB.Model(&model.Serve{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) DisableById(id uint) error {
	return r.DB.Model(&model.Serve{}).Where("id = ?", id).Update("status", 4).Error
}

func (r *ServeRepo) DeleteVersionById(id uint) error {
	return r.DB.Model(&model.ServeVersion{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) DisableVersionById(id uint) error {
	return r.DB.Model(&model.ServeVersion{}).Where("id = ?", id).Update("disabled", 1).Error
}

func (r *ServeRepo) ListServer(serveId uint) (servers model.ServeServerArr, err error) {
	err = r.DB.Where("serve_id = ? AND NOT deleted AND not disabled", serveId).
		Find(&servers).Error

	for index, server := range servers {
		var environment model.Environment
		environment, err = r.EnvironmentRepo.Get(server.EnvironmentId)
		if err != nil {
			return
		}

		servers[index].EnvironmentName = environment.Name
		servers[index].Sort = environment.Sort
	}

	sort.Sort(servers)

	return
}

func (r *ServeRepo) ListSecurity(serveId uint) (res []model.ComponentSchemaSecurity, err error) {
	err = r.DB.Where("serve_id = ? AND NOT deleted AND not disabled", serveId).Find(&res).Error
	return
}

func (r *ServeRepo) DeleteSchemaById(id uint) error {
	return r.DB.Model(&model.ComponentSchema{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) DeleteSecurityId(id uint) error {
	return r.DB.Model(&model.ComponentSchemaSecurity{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) SaveServeEndpointVersions(versions []model.ServeEndpointVersion) error {
	return r.DB.Create(&versions).Error
}

func (r *ServeRepo) DeleteSaveServeEndpointVersions(serveId int64, version string) error {
	return r.DB.Delete(&model.ServeEndpointVersion{}, "serve_id=? and serve_version=?", serveId, version).Error
}

func (r *ServeRepo) BindEndpoint(serveId int64, serveVersion string, serveEndpointVersion []model.ServeEndpointVersion) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DeleteSaveServeEndpointVersions(serveId, serveVersion)
		if err != nil {
			return err
		}
		err = r.SaveServeEndpointVersions(serveEndpointVersion)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (r *ServeRepo) SaveServer(environmentId uint, environmentName string, servers []model.ServeServer) (err error) {
	if len(servers) == 0 {
		return
	}
	err = r.DB.Delete(&model.ServeServer{}, "environment_id=?", environmentId).Error
	if err != nil {
		return err
	}

	for key, _ := range servers {
		//servers[key].ID = 0
		servers[key].EnvironmentId = environmentId
		servers[key].Description = environmentName
		err = r.Save(servers[key].ID, &servers[key])
		if err != nil {
			return err
		}
	}
	/*
		err = r.DB.CreateExpression(servers).Error
		if err != nil {
			return err
		}
	*/
	return
}

func (r *ServeRepo) SaveServerForServe(serveId uint, servers []model.ServeServer) (err error) {
	if len(servers) == 0 {
		return
	}
	err = r.DB.Delete(&model.ServeServer{}, "serve_id=?", serveId).Error
	if err != nil {
		return err
	}

	err = r.DB.Create(servers).Error
	return
}

func (r *ServeRepo) ServeExist(id, projectId uint, name string) (res bool) {
	var count int64
	err := r.DB.Model(&model.Serve{}).Where("id != ? and name = ? and project_id = ? AND NOT deleted", id, name, projectId).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0

}

func (r *ServeRepo) VersionExist(id, serveId uint, value string) (res bool) {
	var count int64
	err := r.DB.Model(&model.ServeVersion{}).Where("id != ? and value = ? and serve_id=? AND NOT deleted", id, value, serveId).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0

}

func (r *ServeRepo) SecurityExist(id, serveId uint, name string) (res bool) {
	var count int64
	err := r.DB.Model(&model.ComponentSchemaSecurity{}).Where("id != ? and name = ? and serve_id=? AND NOT deleted", id, name, serveId).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0

}

func (r *ServeRepo) SchemaExist(id, serveId uint, name string) (res bool) {
	var count int64
	err := r.DB.Model(&model.ComponentSchema{}).Where("id != ? and name = ? and serve_id=? AND NOT deleted", id, name, serveId).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0

}

func (r *ServeRepo) SaveVersion(id uint, version *model.ServeVersion) (err error) {
	if id == 0 {
		err = r.CopyEndpointsVersionRef(version)
		if err != nil {
			return err
		}
	}
	err = r.Save(id, &version)
	if err != nil {
		return err
	}
	return
}

func (r *ServeRepo) GetLatestVersion(serveId uint) (res model.ServeVersion, err error) {
	var version model.ServeVersion
	err = r.DB.Take(&version, "serve_id=?", serveId).Order("value desc").Error
	if err != nil {
		return
	}
	return
}

func (r *ServeRepo) GetBindEndpoints(serveId uint, version string) (endpoints []model.ServeEndpointVersion, err error) {
	err = r.DB.Find(&endpoints, "serve_id=? and serve_version=?", serveId, version).Error
	return
}

func (r *ServeRepo) CopyEndpoints(endpoints []model.ServeEndpointVersion, version string) (err error) {
	for key, _ := range endpoints {
		endpoints[key].ID = 0
		endpoints[key].ServeVersion = version
	}
	return r.DB.Create(endpoints).Error
}

func (r *ServeRepo) CopyEndpointsVersionRef(version *model.ServeVersion) (err error) {
	var latestVersion model.ServeVersion
	latestVersion, _ = r.GetLatestVersion(uint(version.ServeId))
	if latestVersion.Value != "" {
		var endpoints []model.ServeEndpointVersion
		endpoints, err = r.GetBindEndpoints(uint(version.ServeId), latestVersion.Value)
		if err != nil {
			err = r.CopyEndpoints(endpoints, version.Value)
			if err != nil {
				return
			}
		}
	}
	return err
}
func (r *ServeRepo) GetBindEndpointIds(serveId uint, version string) (ids []int64, err error) {
	var endpointVersions []model.ServeEndpointVersion
	endpointVersions, err = r.GetBindEndpoints(serveId, version)
	for _, endpointVersion := range endpointVersions {
		ids = append(ids, endpointVersion.EndpointId)
	}
	return
}

func (r *ServeRepo) ChangeServe(serveId, userId uint) (serve model.Serve, err error) {
	err = r.DB.Model(&model.SysUserProfile{}).Where("user_id = ?", userId).
		Updates(map[string]interface{}{"curr_serve_id": serveId}).Error

	serve, err = r.Get(serveId)

	return
}

func (r *ServeRepo) GetCurrServeByUser(userId uint) (currServe model.Serve, err error) {
	var user model.SysUser
	err = r.DB.Preload("Profile").
		Where("id = ?", userId).
		First(&user).
		Error

	// may be null
	r.DB.Model(&model.Serve{}).
		Where("id = ?", user.Profile.CurrServeId).
		First(&currServe)

	return
}

func (r *ServeRepo) SetCurrServeByUser(serveId, userId uint) (err error) {
	err = r.DB.Model(&model.SysUserProfile{}).
		Where("user_id = ?", userId).
		Update("curr_serve_id", serveId).Error

	return
}

func (r *ServeRepo) GetCurrServerByUser(projectId, serveId, userId uint) (currServer model.ServeServer, err error) {
	projectUserServer, err := r.EnvironmentRepo.GetProjectUserServer(projectId, userId)
	if err != nil {
		return
	}

	// may be null
	err = r.DB.Model(&model.ServeServer{}).
		Where("environment_id = ? and serve_id = ?", projectUserServer.ServerId, serveId).
		First(&currServer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if currServer.ID == 0 || err == gorm.ErrRecordNotFound {
		return currServer, nil
	}

	environment, err := r.EnvironmentRepo.Get(projectUserServer.ServerId)
	if err != nil {
		return
	}
	currServer.EnvironmentName = environment.Name
	currServer.Sort = environment.Sort

	return
}

func (r *ServeRepo) SetCurrServerByUser(serverId, userId uint) (err error) {
	err = r.DB.Model(&model.SysUserProfile{}).
		Where("user_id = ?", userId).
		Update("curr_server_id", serverId).Error

	return
}

func (r *ServeRepo) SaveServe(serve *model.Serve) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if serve.ID == 0 { //生成目录树跟节点
			/*
				err = r.AddDefaultCategory(serve.ProjectId)
				if err != nil {
					return err
				}
			*/

			//创建默认环境
			err = r.Save(serve.ID, &serve)
			if err != nil {
				return err
			}

			err = r.AddDefaultServer(serve.ProjectId, serve.ID)
			if err != nil {
				return err
			}
			//目录树默认不挂在根目录下面
			//			err = r.AddDefaultTestCategory(serve.ProjectId, serve.ID)

		} else {
			err = r.Save(serve.ID, &serve)
		}

		if err != nil {
			return err
		}
		return nil
	})
}

func (r *ServeRepo) GetSchemaByRef(projectId uint, ref string) (res model.ComponentSchema, err error) {
	err = r.DB.Where("project_id = ? AND NOT deleted AND not disabled and ref = ?", projectId, ref).Find(&res).Error
	return
}

func (r *ServeRepo) GetServerCountByEnvironmentId(environmentId uint) (count int64, err error) {
	err = r.DB.Model(&model.ServeServer{}).Where("environment_id = ? AND NOT deleted AND not disabled ", environmentId).Count(&count).Error
	return
}

func (r *ServeRepo) GetCountByProject(projectId uint) (count int64, err error) {
	err = r.DB.Model(&model.Serve{}).Where("project_id = ? AND NOT deleted AND not disabled ", projectId).Count(&count).Error
	return
}

func (r *ServeRepo) CreateServeSample(projectId uint) (serveId uint, err error) {
	return
}

func (r *ServeRepo) AddDefaultCategory(projectId uint) (err error) {
	category := model.Category{Name: "所属分类", ProjectId: projectId, Type: serverConsts.EndpointCategory}
	err = r.CategoryRepo.Save(&category)
	return
}

func (r *ServeRepo) AddDefaultServer(projectId, serveId uint) (err error) {
	//var defaultEnv model.Environment
	defaultEnvs, err := r.EnvironmentRepo.GetDefaultByProject(projectId)
	if err != nil {
		return
	}

	var defaultServer []model.ServeServer
	for _, v := range defaultEnvs {
		server := model.ServeServer{
			ServeId:       serveId,
			EnvironmentId: v.ID,
			Description:   v.Name,
			Url:           serverConsts.DefaultSever,
		}
		if v.Name == "Mock环境" {
			host, _ := cache.GetCacheString("host")
			server.Url = host + "/mocks/" + strconv.Itoa(int(serveId))
		}
		defaultServer = append(defaultServer, server)
	}
	err = r.SaveServerForServe(serveId, defaultServer)

	if err != nil {
		return
	}

	return
}

func (r *ServeRepo) AddDefaultTestCategory(projectId uint) (err error) {
	root := model.DiagnoseInterface{
		Title:     "根节点",
		ProjectId: projectId,
		IsDir:     true,
		Type:      "dir",
	}
	err = r.DB.Create(&root).Error
	return
}

func (r *ServeRepo) GetServesByIds(ids []uint) (serves []model.Serve, err error) {
	err = r.DB.Where("id in ? AND not deleted AND not disabled ", ids).Find(&serves).Error
	return
}

func (r *ServeRepo) GetSchemas(serveIs []uint) (res []model.ComponentSchema, err error) {
	err = r.DB.Where("NOT deleted AND not disabled and serve_id in ?", serveIs).Find(&res).Error
	return
}

func (r *ServeRepo) GetSecurities(serveIs []uint) (res []model.ComponentSchemaSecurity, err error) {
	err = r.DB.Where("NOT deleted AND not disabled and serve_id in ?", serveIs).Find(&res).Error
	return
}

func (r *ServeRepo) GetServers(serveIs []uint) (res []model.ServeServer, err error) {
	err = r.DB.Where("NOT deleted AND not disabled and serve_id in ?", serveIs).Find(&res).Error
	var environmentIds []uint
	for _, server := range res {
		environmentIds = append(environmentIds, server.EnvironmentId)
	}

	environments, _ := r.EnvironmentRepo.GetByIds(environmentIds)
	for key, server := range res {
		res[key].EnvironmentName = environments[server.EnvironmentId].Name
	}
	return
}

func (r *ServeRepo) CreateSchemas(schemas []*model.ComponentSchema) (err error) {
	return r.DB.Create(schemas).Error
}

func (r *ServeRepo) SaveSwaggerSync(sync *model.SwaggerSync) (err error) {
	return r.Save(sync.ID, sync)
}

func (r *ServeRepo) GetSwaggerSync(projectId uint) (sync model.SwaggerSync, err error) {
	err = r.DB.First(&sync, "project_id=?", projectId).Error
	return
}

func (r *ServeRepo) GetSwaggerSyncById(id uint) (sync model.SwaggerSync, err error) {
	err = r.DB.First(&sync, "id=?", id).Error
	return
}

func (r *ServeRepo) GetSwaggerSyncList() (res []model.SwaggerSync, err error) {
	err = r.DB.Find(&res).Error
	return
}

func (r *ServeRepo) GetDefault(projectId uint) (res model.Serve, err error) {
	err = r.DB.Where("NOT deleted and project_id=?", projectId).First(&res).Error
	return
}

func (r *ServeRepo) GetComponentByItem(sourceType consts.SourceType, serveId uint, ref string) (res model.ComponentSchema, err error) {
	err = r.DB.First(&res, "source_type=? AND serve_id=? AND ref=? AND NOT deleted", sourceType, serveId, ref).Error
	return
}

func (r *ServeRepo) SaveSchemas(schemas []*model.ComponentSchema) (err error) {
	return r.DB.Save(schemas).Error
}

func (r *ServeRepo) UpdateSwaggerSyncExecTimeById(id uint) (err error) {
	return r.DB.Model(&model.SwaggerSync{}).Where("id=?", id).Update("exec_time", time.Now()).Error
}

func (r *ServeRepo) ListAll() (res []model.Serve, err error) {
	err = r.DB.Model(&model.Serve{}).
		Where("not disabled and not deleted").
		Find(&res).Error
	return
}

func (r *ServeRepo) BatchUpdateSchemaProjectByServeId(serveIds []uint, projectId uint) (err error) {
	err = r.DB.Model(&model.ComponentSchema{}).Where("serve_id IN (?)", serveIds).Update("project_id", projectId).Error
	return
}
