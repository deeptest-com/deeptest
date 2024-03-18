package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type EndpointRepo struct {
	*BaseRepo             `inject:""`
	EndpointInterfaceRepo *EndpointInterfaceRepo `inject:""`
	ServeRepo             *ServeRepo             `inject:""`
	ProjectRepo           *ProjectRepo           `inject:""`
	EndpointTagRepo       *EndpointTagRepo       `inject:""`
	EnvironmentRepo       *EnvironmentRepo       `inject:""`
	EndpointFavoriteRepo  *EndpointFavoriteRepo  `inject:""`
}

func (r *EndpointRepo) Paginate(tenantId consts.TenantId, req v1.EndpointReqPaginate) (ret _domain.PageData, err error) {
	//fmt.Println(r.GetDB(tenantId).Model(&modelRef.SysUser{}))
	//err = r.GetDB(tenantId).Where("id=?", id).Where("name=?", name).Find(&res).Error
	var count int64
	db := r.GetDB(tenantId).Model(&model.Endpoint{}).Where("project_id = ? AND NOT deleted AND NOT disabled", req.ProjectId)

	if req.Title != "" {
		db = db.Where("title LIKE ? or path LIKE ?", fmt.Sprintf("%%%s%%", req.Title), fmt.Sprintf("%%%s%%", req.Title))
	}
	if len(req.CreateUser) > 0 {
		db = db.Where("create_user in ?", req.CreateUser)
	}
	if len(req.Status) > 0 {
		db = db.Where("status in ?", req.Status)
	}
	if req.ServeId != 0 {
		db = db.Where("serve_id = ?", req.ServeId)
	}
	if len(req.ServeIds) > 0 {
		db = db.Where("serve_id in ?", req.ServeIds)
	}

	if req.ServeVersion != "" {
		if ids, err := r.ServeRepo.GetBindEndpointIds(tenantId, req.ServeId, req.ServeVersion); err == nil {
			db = db.Where("id in ?", ids)
		}
	}

	if req.IsFavorite {
		if ids, err := r.EndpointFavoriteRepo.GetEndpointIds(tenantId, req.UserId); err == nil {
			db = db.Where("id in ?", ids)
		}
	}

	if req.CategoryId > 0 {
		var categoryIds []uint
		categoryIds, err = r.BaseRepo.GetDescendantIds(tenantId, uint(req.CategoryId), model.Category{}.TableName(),
			serverConsts.EndpointCategory, int(req.ProjectId))
		if err != nil {
			return
		}
		if len(categoryIds) > 0 {
			db.Where("category_id IN(?)", categoryIds)
		}
	} else if req.CategoryId == -1 {
		db.Where("category_id IN(-1)")
	}

	if len(req.TagNames) != 0 {
		endpointIds, err := r.EndpointTagRepo.GetEndpointIdsByTagNames(tenantId, req.TagNames, req.ProjectId)
		if err != nil {
			return ret, err
		}
		db.Where("id IN (?)", endpointIds)
	}

	db = db.Order("created_at desc")
	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.Endpoint, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}

	serveNames := map[uint]string{}

	for key, result := range results {
		var versions []model.EndpointVersion
		r.GetDB(tenantId).Find(&versions, "endpoint_id=?", result.ID).Order("version desc")
		results[key].Versions = versions
		if len(versions) > 0 {
			results[key].Version = versions[0].Version
		}

		if _, ok := serveNames[result.ServeId]; !ok {
			var serve model.Serve
			r.GetDB(tenantId).Find(&serve, "id=?", result.ServeId)
			serveNames[result.ServeId] = serve.Name
		}
		results[key].ServeName = serveNames[result.ServeId]

		results[key].Tags, err = r.EndpointTagRepo.GetTagNamesByEndpointId(tenantId, result.ID, result.ProjectId)
	}

	r.CombineMethodsForEndpoints(tenantId, results)

	ret.Populate(results, count, req.Page, req.PageSize)

	return
}
func (r *EndpointRepo) CombineMethodsForEndpoints(tenantId consts.TenantId, endpoints []*model.Endpoint) {
	endpointIds := make([]uint, 0)
	for _, v := range endpoints {
		endpointIds = append(endpointIds, v.ID)
	}
	if len(endpointIds) == 0 {
		return
	}

	interfaces, err := r.EndpointInterfaceRepo.BatchGetByEndpointIds(tenantId, endpointIds)
	if err != nil {
		return
	}

	endpointMethodsMap := make(map[uint][]consts.HttpMethod)
	for _, v := range interfaces {
		endpointMethodsMap[v.EndpointId] = append(endpointMethodsMap[v.EndpointId], v.Method)
	}

	for k, v := range endpoints {
		endpoints[k].Methods = endpointMethodsMap[v.ID]
	}
	return
}
func (r *EndpointRepo) SaveAll(tenantId consts.TenantId, endpoint *model.Endpoint) (err error) {
	r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {

		//更新终端
		err = r.saveEndpoint(tenantId, endpoint)
		if err != nil {
			return err
		}

		//创建version
		err = r.saveEndpointVersion(tenantId, endpoint)
		if err != nil {
			return err
		}

		//保存路径参数
		err = r.saveEndpointParams(tenantId, endpoint.ID, endpoint.PathParams)
		if err != nil {
			return err
		}

		//保存接口
		err = r.saveInterfaces(tenantId, endpoint.ID, endpoint.ProjectId, endpoint.Path, endpoint.Version, endpoint.Title, endpoint.Interfaces)
		if err != nil {
			return err
		}

		//更新调试接口
		//err = r.updateDebugInterfaceUrl(endpoint.ID, endpoint.Path)
		//if err != nil {
		//	return err
		//}

		return nil
	})
	return
}

// 保存终端信息
func (r *EndpointRepo) saveEndpoint(tenantId consts.TenantId, endpoint *model.Endpoint) (err error) {
	err = r.Save(tenantId, endpoint.ID, endpoint)
	if err != nil {
		return
	}

	err = r.UpdateSerialNumber(tenantId, endpoint.ID, uint(endpoint.ProjectId))
	return
}

func (r *EndpointRepo) UpdateSerialNumber(tenantId consts.TenantId, id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(tenantId, projectId)
	if err != nil {
		return
	}

	err = r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id=?", id).Update("serial_number", project.ShortName+"-I-"+strconv.Itoa(int(id))).Error
	return

}

func (r *EndpointRepo) saveEndpointVersion(tenantId consts.TenantId, endpoint *model.Endpoint) (err error) {
	if endpoint.Version == "" {
		endpoint.Version = "v0.1.0"
	}

	endpointVersion := model.EndpointVersion{EndpointId: endpoint.ID, Version: endpoint.Version}
	r.FindVersion(tenantId, &endpointVersion)
	if endpointVersion.ID == 0 {
		err = r.GetDB(tenantId).Create(&endpointVersion).Error
		if err != nil {
			endpoint.Version = endpointVersion.Version
		}
	}

	return
}

// 保存路径参数
func (r *EndpointRepo) saveEndpointParams(tenantId consts.TenantId, endpointId uint, params []model.EndpointPathParam) (err error) {
	err = r.removeEndpointParams(tenantId, endpointId)
	if err != nil {
		return
	}
	for _, item := range params {
		item.EndpointId = endpointId
		err = r.Save(tenantId, item.ID, &item)
		if err != nil {
			return
		}
	}
	return
}

func (r *EndpointRepo) removeEndpointParams(tenantId consts.TenantId, endpointId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("endpoint_id = ?", endpointId).
		Delete(&model.EndpointPathParam{}, "").Error

	return
}

// 保存接口信息
func (r *EndpointRepo) saveInterfaces(tenantId consts.TenantId, endpointId, projectId uint, path, version, title string, interfaces []model.EndpointInterface) (err error) {

	interfaceIds := make([]uint, 0)
	for _, v := range interfaces {
		if v.ID != 0 {
			interfaceIds = append(interfaceIds, v.ID)
		}
	}
	interfaceIdModelMap, err := r.EndpointInterfaceRepo.GetIdAndModelMap(tenantId, interfaceIds)

	err = r.removeInterfaces(tenantId, endpointId)
	if err != nil {
		return
	}

	for key, v := range interfaces {
		if interfaceModel, ok := interfaceIdModelMap[v.ID]; ok {
			interfaces[key].DebugInterfaceId = interfaceModel.DebugInterfaceId
		}

		interfaces[key].EndpointId = endpointId
		interfaces[key].Version = version
		interfaces[key].Url = path
		interfaces[key].ProjectId = projectId
		interfaces[key].Name = title

		err = r.EndpointInterfaceRepo.SaveInterfaces(tenantId, &interfaces[key])
		if err != nil {
			return err
		}
	}
	return
}

// 保存调试接口Url
func (r *EndpointRepo) updateDebugInterfaceUrl(tenantId consts.TenantId, endpointId uint, url string) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
		Where("endpoint_id = ?", endpointId).
		Update("url", url).Error

	return
}

func (r *EndpointRepo) removeInterfaces(tenantId consts.TenantId, endpointId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("endpoint_id = ?", endpointId).
		Delete(&model.EndpointInterface{}, "").Error

	return
}

func (r *EndpointRepo) GetAll(tenantId consts.TenantId, id uint, version string) (endpoint model.Endpoint, err error) {
	endpoint, err = r.Get(tenantId, id)
	if err != nil {
		return
	}

	endpoint.Tags, _ = r.EndpointTagRepo.GetTagNamesByEndpointId(tenantId, id, endpoint.ProjectId)
	endpoint.PathParams, _ = r.GetEndpointPathParams(tenantId, id)
	endpoint.Interfaces, _ = r.EndpointInterfaceRepo.ListByEndpointId(tenantId, id, version)
	endpoint.GlobalParams, _ = r.EnvironmentRepo.ListParamModel(tenantId, endpoint.ProjectId)

	return
}

func (r *EndpointRepo) GetWithInterface(tenantId consts.TenantId, id uint, version string) (endpoint model.Endpoint, err error) {
	endpoint, err = r.Get(tenantId, id)
	if err != nil {
		return
	}

	endpoint.Tags, _ = r.EndpointTagRepo.GetTagNamesByEndpointId(tenantId, id, endpoint.ProjectId)
	endpoint.PathParams, _ = r.GetEndpointPathParams(tenantId, id)
	endpoint.Interfaces, _ = r.EndpointInterfaceRepo.ListByEndpointId(tenantId, id, version)

	return
}

func (r *EndpointRepo) Get(tenantId consts.TenantId, id uint) (res model.Endpoint, err error) {
	err = r.GetDB(tenantId).First(&res, id).Error
	return
}

func (r *EndpointRepo) GetEndpointPathParams(tenantId consts.TenantId, endpointId uint) (pathParam []model.EndpointPathParam, err error) {
	err = r.GetDB(tenantId).Find(&pathParam, "endpoint_id=?", endpointId).Error
	return
}

func (r *EndpointRepo) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("id = ?", id).
		Update("deleted", 1).Error

	return
}
func (r *EndpointRepo) DeleteByIds(tenantId consts.TenantId, ids []uint) error {
	return r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id IN ?", ids).Update("deleted", 1).Error
}

func (r *EndpointRepo) DeleteByCategoryIds(tenantId consts.TenantId, categoryIds []int64) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("category_id IN ?", categoryIds).
		Update("deleted", 1).Error

	return
}

func (r *EndpointRepo) DisableById(tenantId consts.TenantId, id uint) error {
	return r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id = ?", id).Update("status", 4).Error
}

func (r *EndpointRepo) UpdateStatus(tenantId consts.TenantId, id uint, status int64) error {
	return r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id = ?", id).Update("status", status).Error
}

func (r *EndpointRepo) GetVersionsByEndpointId(tenantId consts.TenantId, endpointId uint) (res []model.EndpointVersion, err error) {
	err = r.GetDB(tenantId).Find(&res, "endpoint_id=?", endpointId).Error
	return
}

func (r *EndpointRepo) GetLatestVersion(tenantId consts.TenantId, endpointId uint) (res model.EndpointVersion, err error) {
	err = r.GetDB(tenantId).Take(&res, "endpoint_id=?", endpointId).Order("version desc").Error
	return
}
func (r *EndpointRepo) FindVersion(tenantId consts.TenantId, res *model.EndpointVersion) (err error) {
	err = r.GetDB(tenantId).Where("endpoint_id=? and version=?", res.EndpointId, res.Version).First(&res).Error
	return
}

func (r *EndpointRepo) GetFirstMethod(tenantId consts.TenantId, id uint) (res model.EndpointInterface, err error) {
	var interfs []model.EndpointInterface
	interfs, err = r.EndpointInterfaceRepo.ListByEndpointId(tenantId, id, "v0.1.0")
	if len(interfs) > 0 {
		res = interfs[0]
	}
	return
}

func (r *EndpointRepo) GetCountByServeId(tenantId consts.TenantId, serveId uint) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).Where("serve_id=? and NOT deleted", serveId).Count(&count).Error
	return
}

func (r *EndpointRepo) ListEndpointByCategory(tenantId consts.TenantId, categoryId uint) (ids []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Select("id").
		Where("category_id = ? AND NOT deleted", categoryId).
		Find(&ids).Error
	return
}

func (r *EndpointRepo) ListEndpointByCategories(tenantId consts.TenantId, categoryIds []uint) (ids []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Select("id").
		Where("category_id IN (?) AND NOT deleted", categoryIds).
		Find(&ids).Error
	return
}

func (r *EndpointRepo) CreateEndpointSample(tenantId consts.TenantId, serveId uint) (endpointId uint, err error) {

	return
}

func (r *EndpointRepo) GetCategoryCount(tenantId consts.TenantId, result interface{}, projectId uint) (err error) {
	err = r.GetDB(tenantId).Raw("select count(id) count, category_id from "+model.Endpoint{}.TableName()+" where not deleted and not disabled and project_id=? group by category_id", projectId).Scan(result).Error
	return
}

func (r *EndpointRepo) GetByProjectId(tenantId consts.TenantId, projectId uint, needDetail bool) (endpoints []*model.Endpoint, err error) {
	err = r.GetDB(tenantId).Find(&endpoints, "project_id = ? and not deleted and not disabled", projectId).Error
	r.GetByEndpoints(tenantId, endpoints, needDetail)
	return
}

func (r *EndpointRepo) GetByServeIds(tenantId consts.TenantId, serveIds []uint, needDetail bool) (endpoints []*model.Endpoint, err error) {
	err = r.GetDB(tenantId).Where("serve_id = ? and not deleted and not disabled", serveIds).Find(&endpoints).Error
	r.GetByEndpoints(tenantId, endpoints, needDetail)
	return
}

func (r *EndpointRepo) GetByEndpointIds(tenantId consts.TenantId, endpointIds []uint, needDetail bool) (endpoints []*model.Endpoint, err error) {
	err = r.GetDB(tenantId).Where("id in ? and not deleted and not disabled", endpointIds).Find(&endpoints).Error
	r.GetByEndpoints(tenantId, endpoints, needDetail)
	return
}

func (r *EndpointRepo) GetByEndpoints(tenantId consts.TenantId, endpoints []*model.Endpoint, needDetail bool) {
	var endpointIds []uint
	for _, endpoint := range endpoints {
		endpointIds = append(endpointIds, endpoint.ID)
	}

	interfaces, err := r.EndpointInterfaceRepo.GetInterfaces(tenantId, endpointIds, needDetail)
	if err != nil {
		return
	}

	for key, endpoint := range endpoints {
		endpoints[key].Interfaces = interfaces[endpoint.ID]
	}
	return
}

func (r *EndpointRepo) GetPathParams(tenantId consts.TenantId, endpointIds []uint) (err error, pathParams model.EndpointPathParam) {
	err = r.GetDB(tenantId).Find(&pathParams, "not disabled and not deleted and endpoint_id in ?", endpointIds).Error
	return
}

func (r *EndpointRepo) GetUsedCountByEndpointId(tenantId consts.TenantId, endpointId uint) (count int64, err error) {
	endpointInterfaceIds, _ := r.EndpointInterfaceRepo.ListIdByEndpoint(tenantId, endpointId)

	err = r.GetDB(tenantId).Model(&model.Processor{}).
		Where("NOT deleted and endpoint_interface_id IN (?)", endpointInterfaceIds).
		Count(&count).Error

	return
}

func (r *EndpointRepo) CreateEndpoints(tenantId consts.TenantId, endpoints []*model.Endpoint) error {
	return r.GetDB(tenantId).Create(endpoints).Error
}

func (r *EndpointRepo) BatchUpdateStatus(tenantId consts.TenantId, ids []uint, status int64) error {
	return r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id IN (?)", ids).Update("status", status).Error
}

func (r *EndpointRepo) BatchUpdateCategory(tenantId consts.TenantId, ids []uint, categoryId int64) error {
	return r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id IN (?)", ids).Update("category_id", categoryId).Error
}

func (r *EndpointRepo) BatchUpdate(tenantId consts.TenantId, ids []uint, data map[string]interface{}) error {
	return r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id IN (?)", ids).Updates(data).Error
}

func (r *EndpointRepo) GetByItem(tenantId consts.TenantId, sourceType consts.SourceType, projectId uint, path string, serveId uint, categoryId int64) (res model.Endpoint, err error) {
	db := r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("source_type = ?", sourceType).
		Where("project_id = ?", projectId).
		Where("path = ?", path).
		Where("serve_id = ? AND NOT deleted", serveId)

	if categoryId > 0 {
		categoryIds, err := r.BaseRepo.GetDescendantIds(tenantId, uint(categoryId), model.Category{}.TableName(), serverConsts.EndpointCategory, int(projectId))
		if err != nil {
			return res, err
		}
		if len(categoryIds) > 0 {
			db.Where("category_id IN (?)", categoryIds)
		}
	} else if categoryId == -1 {
		db.Where("category_id = -1")
	}

	err = db.First(&res).Error

	return

}

func (r *EndpointRepo) ListByProjectIdAndServeId(tenantId consts.TenantId, serveId uint, method consts.HttpMethod) (endpoints []*model.Endpoint, err error) {

	err = r.GetDB(tenantId).Model(&model.Endpoint{}).Select("biz_endpoint.*").Joins("left join biz_endpoint_interface on biz_endpoint_interface.endpoint_id = biz_endpoint.id").Where("not biz_endpoint.deleted and not biz_endpoint_interface.deleted and biz_endpoint.serve_id=? and biz_endpoint_interface.method=?", serveId, method).Scan(&endpoints).Error

	return
}

func (r *EndpointRepo) GetByPath(tenantId consts.TenantId, serveId uint, pth string, method consts.HttpMethod) (endpoints []*model.Endpoint, err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).Select("biz_endpoint.*").Joins("left join biz_endpoint_interface on biz_endpoint_interface.endpoint_id = biz_endpoint.id").Where("not biz_endpoint.deleted and not biz_endpoint_interface.deleted and biz_endpoint.serve_id=? and biz_endpoint_interface.method=? and biz_endpoint.path=?", serveId, method, pth).Scan(&endpoints).Error

	return
}

func (r *EndpointRepo) UpdateAdvancedMockDisabled(tenantId consts.TenantId, endpointId uint, advancedMockDisabled bool) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("id = ?", endpointId).
		Update("advanced_mock_disabled", advancedMockDisabled).Error

	return
}

func (r *EndpointRepo) GetByNameAndProject(tenantId consts.TenantId, name string, projectId uint) (res model.Endpoint, err error) {
	err = r.GetDB(tenantId).Where("title = ?", name).
		Where("project_id = ?", projectId).
		First(&res).Error
	return
}

func (r *EndpointRepo) UpdateBodyIsChanged(tenantId consts.TenantId, endpointId uint, changedStatus consts.ChangedStatus) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("id = ?", endpointId).
		Updates(map[string]interface{}{"changed_status": changedStatus, "Updated_at": time.Now()}).Error

	return
}

func (r *EndpointRepo) UpdateSnapshot(tenantId consts.TenantId, endpointId uint, snapshot string) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("id = ?", endpointId).
		UpdateColumns(map[string]interface{}{"changed_status": consts.Changed, "snapshot": snapshot, "changed_time": time.Now()}).Error

	return
}

func (r *EndpointRepo) UpdateName(tenantId consts.TenantId, id uint, name string) (err error) {
	return r.GetDB(tenantId).Model(&model.Endpoint{}).Where("id = ?", id).Update("title", name).Error
}

func (r *EndpointRepo) ChangeSnapShot(tenantId consts.TenantId, endpointId uint, snapshot string) (err error) {
	err = r.GetDB(tenantId).Model(&model.Endpoint{}).
		Where("id = ?", endpointId).
		UpdateColumn("snapshot", snapshot).Error

	return
}

func (r *EndpointRepo) GetByCategoryId(tenantId consts.TenantId, categoryId uint) (endpoints []model.Endpoint, err error) {
	err = r.GetDB(tenantId).Where("category_id = ?", categoryId).Order("created_at desc").Find(&endpoints).Error
	return endpoints, err
}

func (r *EndpointRepo) FavoriteList(tenantId consts.TenantId, projectId, userId uint) (endpoints []model.Endpoint, err error) {
	endpointIds, err := r.EndpointFavoriteRepo.GetEndpointIds(tenantId, userId)
	if err != nil {
		return
	}
	err = r.GetDB(tenantId).Where("project_id = ? and id in  ? and not deleted", projectId, endpointIds).Order("created_at desc").Find(&endpoints).Error

	return
}

func (r *EndpointRepo) GetEntity(tenantId consts.TenantId, id uint) (data map[string]interface{}, err error) {
	data = map[string]interface{}{}
	endpoint, err := r.Get(tenantId, id)
	if err != nil {
		return
	}
	data["id"] = endpoint.ID
	data["name"] = endpoint.Title
	data["method"], _ = r.EndpointInterfaceRepo.GetMethodsByEndpointId(tenantId, id)
	return
}
