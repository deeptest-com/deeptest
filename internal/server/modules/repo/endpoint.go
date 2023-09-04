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
)

type EndpointRepo struct {
	*BaseRepo             `inject:""`
	EndpointInterfaceRepo *EndpointInterfaceRepo `inject:""`
	ServeRepo             *ServeRepo             `inject:""`
	ProjectRepo           *ProjectRepo           `inject:""`
	EndpointTagRepo       *EndpointTagRepo       `inject:""`
}

func (r *EndpointRepo) Paginate(req v1.EndpointReqPaginate) (ret _domain.PageData, err error) {
	//fmt.Println(r.DB.Model(&modelRef.SysUser{}))
	//err = r.DB.Where("id=?", id).Where("name=?", name).Find(&res).Error
	var count int64
	db := r.DB.Model(&model.Endpoint{}).Where("project_id = ? AND NOT deleted AND NOT disabled", req.ProjectId)

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
	if req.ServeVersion != "" {
		if ids, err := r.ServeRepo.GetBindEndpointIds(req.ServeId, req.ServeVersion); err != nil {
			db = db.Where("id in ?", ids)
		}
	}

	if req.CategoryId > 0 {
		var categoryIds []uint
		categoryIds, err = r.BaseRepo.GetDescendantIds(uint(req.CategoryId), model.Category{}.TableName(),
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
		endpointIds, err := r.EndpointTagRepo.GetEndpointIdsByTagNames(req.TagNames, req.ProjectId)
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
		r.DB.Find(&versions, "endpoint_id=?", result.ID).Order("version desc")
		results[key].Versions = versions
		if len(versions) > 0 {
			results[key].Version = versions[0].Version
		}

		if _, ok := serveNames[result.ServeId]; !ok {
			var serve model.Serve
			r.DB.Find(&serve, "id=?", result.ServeId)
			serveNames[result.ServeId] = serve.Name
		}
		results[key].ServeName = serveNames[result.ServeId]

		results[key].Tags, err = r.EndpointTagRepo.GetTagNamesByEndpointId(result.ID, result.ProjectId)
	}

	r.CombineMethodsForEndpoints(results)

	ret.Populate(results, count, req.Page, req.PageSize)

	return
}
func (r *EndpointRepo) CombineMethodsForEndpoints(endpoints []*model.Endpoint) {
	endpointIds := make([]uint, 0)
	for _, v := range endpoints {
		endpointIds = append(endpointIds, v.ID)
	}
	if len(endpointIds) == 0 {
		return
	}

	interfaces, err := r.EndpointInterfaceRepo.BatchGetByEndpointIds(endpointIds)
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
func (r *EndpointRepo) SaveAll(endpoint *model.Endpoint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {

		//更新终端
		err = r.saveEndpoint(endpoint)
		if err != nil {
			return err
		}

		//创建version
		err = r.saveEndpointVersion(endpoint)
		if err != nil {
			return err
		}

		//保存路径参数
		err = r.saveEndpointParams(endpoint.ID, endpoint.PathParams)
		if err != nil {
			return err
		}

		//保存接口
		err = r.saveInterfaces(endpoint.ID, endpoint.ProjectId, endpoint.Path, endpoint.Version, endpoint.Interfaces)
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

//保存终端信息
func (r *EndpointRepo) saveEndpoint(endpoint *model.Endpoint) (err error) {
	err = r.Save(endpoint.ID, endpoint)
	if err != nil {
		return
	}

	err = r.UpdateSerialNumber(endpoint.ID, uint(endpoint.ProjectId))
	return
}

func (r *EndpointRepo) UpdateSerialNumber(id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}

	err = r.DB.Model(&model.Endpoint{}).Where("id=?", id).Update("serial_number", project.ShortName+"-I-"+strconv.Itoa(int(id))).Error
	return

}

func (r *EndpointRepo) saveEndpointVersion(endpoint *model.Endpoint) (err error) {
	if endpoint.Version == "" {
		endpoint.Version = "v0.1.0"
	}

	endpointVersion := model.EndpointVersion{EndpointId: endpoint.ID, Version: endpoint.Version}
	r.FindVersion(&endpointVersion)
	if endpointVersion.ID == 0 {
		err = r.DB.Create(&endpointVersion).Error
		if err != nil {
			endpoint.Version = endpointVersion.Version
		}
	}

	return
}

//保存路径参数
func (r *EndpointRepo) saveEndpointParams(endpointId uint, params []model.EndpointPathParam) (err error) {
	err = r.removeEndpointParams(endpointId)
	if err != nil {
		return
	}
	for _, item := range params {
		item.EndpointId = endpointId
		err = r.Save(item.ID, &item)
		if err != nil {
			return
		}
	}
	return
}

func (r *EndpointRepo) removeEndpointParams(endpointId uint) (err error) {
	err = r.DB.
		Where("endpoint_id = ?", endpointId).
		Delete(&model.EndpointPathParam{}, "").Error

	return
}

//保存接口信息
func (r *EndpointRepo) saveInterfaces(endpointId, projectId uint, path, version string, interfaces []model.EndpointInterface) (err error) {
	interfaceIds := make([]uint, 0)
	for _, v := range interfaces {
		if v.ID != 0 {
			interfaceIds = append(interfaceIds, v.ID)
		}
	}
	interfaceIdModelMap, err := r.EndpointInterfaceRepo.GetIdAndModelMap(interfaceIds)

	err = r.removeInterfaces(endpointId)
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

		err = r.EndpointInterfaceRepo.SaveInterfaces(&interfaces[key])
		if err != nil {
			return err
		}
	}
	return
}

//保存调试接口Url
func (r *EndpointRepo) updateDebugInterfaceUrl(endpointId uint, url string) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("endpoint_id = ?", endpointId).
		Update("url", url).Error

	return
}

func (r *EndpointRepo) removeInterfaces(endpointId uint) (err error) {
	err = r.DB.
		Where("endpoint_id = ?", endpointId).
		Delete(&model.EndpointInterface{}, "").Error

	return
}

func (r *EndpointRepo) GetAll(id uint, version string) (endpoint model.Endpoint, err error) {
	endpoint, err = r.Get(id)
	if err != nil {
		return
	}

	endpoint.Tags, _ = r.EndpointTagRepo.GetTagNamesByEndpointId(id, endpoint.ProjectId)
	endpoint.PathParams, _ = r.GetEndpointParams(id)
	endpoint.Interfaces, _ = r.EndpointInterfaceRepo.ListByEndpointId(id, version)

	return
}

func (r *EndpointRepo) GetWithInterface(id uint, version string) (endpoint model.Endpoint, err error) {
	endpoint, err = r.Get(id)
	if err != nil {
		return
	}

	endpoint.Tags, _ = r.EndpointTagRepo.GetTagNamesByEndpointId(id, endpoint.ProjectId)
	endpoint.PathParams, _ = r.GetEndpointParams(id)
	endpoint.Interfaces, _ = r.EndpointInterfaceRepo.ListByEndpointId(id, version)

	return
}

func (r *EndpointRepo) Get(id uint) (res model.Endpoint, err error) {
	err = r.DB.First(&res, id).Error
	return
}

func (r *EndpointRepo) GetEndpointParams(endpointId uint) (pathParam []model.EndpointPathParam, err error) {
	err = r.DB.Find(&pathParam, "endpoint_id=?", endpointId).Error
	return
}

func (r *EndpointRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.Endpoint{}).
		Where("id = ?", id).
		Update("deleted", 1).Error

	return
}
func (r *EndpointRepo) DeleteByIds(ids []uint) error {
	return r.DB.Model(&model.Endpoint{}).Where("id IN ?", ids).Update("deleted", 1).Error
}

func (r *EndpointRepo) DeleteByCategoryIds(categoryIds []int64) (err error) {
	err = r.DB.Model(&model.Endpoint{}).
		Where("category_id IN ?", categoryIds).
		Update("deleted", 1).Error

	return
}

func (r *EndpointRepo) DisableById(id uint) error {
	return r.DB.Model(&model.Endpoint{}).Where("id = ?", id).Update("status", 4).Error
}

func (r *EndpointRepo) UpdateStatus(id uint, status int64) error {
	return r.DB.Model(&model.Endpoint{}).Where("id = ?", id).Update("status", status).Error
}

func (r *EndpointRepo) GetVersionsByEndpointId(endpointId uint) (res []model.EndpointVersion, err error) {
	err = r.DB.Find(&res, "endpoint_id=?", endpointId).Error
	return
}

func (r *EndpointRepo) GetLatestVersion(endpointId uint) (res model.EndpointVersion, err error) {
	err = r.DB.Take(&res, "endpoint_id=?", endpointId).Order("version desc").Error
	return
}
func (r *EndpointRepo) FindVersion(res *model.EndpointVersion) (err error) {
	err = r.DB.Where("endpoint_id=? and version=?", res.EndpointId, res.Version).First(&res).Error
	return
}

func (r *EndpointRepo) GetFirstMethod(id uint) (res model.EndpointInterface, err error) {
	var interfs []model.EndpointInterface
	interfs, err = r.EndpointInterfaceRepo.ListByEndpointId(id, "v0.1.0")
	if len(interfs) > 0 {
		res = interfs[0]
	}
	return
}

func (r *EndpointRepo) GetCountByServeId(serveId uint) (count int64, err error) {
	err = r.DB.Model(&model.Endpoint{}).Where("serve_id=? and NOT deleted", serveId).Count(&count).Error
	return
}

func (r *EndpointRepo) ListEndpointByCategory(categoryId uint) (ids []uint, err error) {
	err = r.DB.Model(&model.Endpoint{}).
		Select("id").
		Where("category_id = ? AND NOT deleted", categoryId).
		Find(&ids).Error
	return
}

func (r *EndpointRepo) ListEndpointByCategories(categoryIds []uint) (ids []uint, err error) {
	err = r.DB.Model(&model.Endpoint{}).
		Select("id").
		Where("category_id IN (?) AND NOT deleted", categoryIds).
		Find(&ids).Error
	return
}

func (r *EndpointRepo) CreateEndpointSample(serveId uint) (endpointId uint, err error) {

	return
}

func (r *EndpointRepo) GetCategoryCount(result interface{}, projectId uint) (err error) {
	err = r.DB.Raw("select count(id) count, category_id from "+model.Endpoint{}.TableName()+" where not deleted and not disabled and project_id=? group by category_id", projectId).Scan(result).Error
	return
}

func (r *EndpointRepo) GetByProjectId(projectId uint, needDetail bool) (endpoints []*model.Endpoint, err error) {
	err = r.DB.Find(&endpoints, "project_id = ? and not deleted and not disabled", projectId).Error
	r.GetByEndpoints(endpoints, needDetail)
	return
}

func (r *EndpointRepo) GetByServeIds(serveIds []uint, needDetail bool) (endpoints []*model.Endpoint, err error) {
	err = r.DB.Where("serve_id = ? and not deleted and not disabled", serveIds).Find(&endpoints).Error
	r.GetByEndpoints(endpoints, needDetail)
	return
}

func (r *EndpointRepo) GetByEndpointIds(endpointIds []uint, needDetail bool) (endpoints []*model.Endpoint, err error) {
	err = r.DB.Where("id in ? and not deleted and not disabled", endpointIds).Find(&endpoints).Error
	r.GetByEndpoints(endpoints, needDetail)
	return
}

func (r *EndpointRepo) GetByEndpoints(endpoints []*model.Endpoint, needDetail bool) {
	var endpointIds []uint
	for _, endpoint := range endpoints {
		endpointIds = append(endpointIds, endpoint.ID)
	}

	interfaces, err := r.EndpointInterfaceRepo.GetInterfaces(endpointIds, needDetail)
	if err != nil {
		return
	}

	for key, endpoint := range endpoints {
		endpoints[key].Interfaces = interfaces[endpoint.ID]
	}
	return
}

func (r *EndpointRepo) GetPathParams(endpointIds []uint) (err error, pathParams model.EndpointPathParam) {
	err = r.DB.Find(&pathParams, "not disabled and not deleted and endpoint_id in ?", endpointIds).Error
	return
}

func (r *EndpointRepo) GetUsedCountByEndpointId(endpointId uint) (count int64, err error) {
	endpointInterfaceIds, _ := r.EndpointInterfaceRepo.ListIdByEndpoint(endpointId)

	err = r.DB.Model(&model.Processor{}).
		Where("NOT deleted and endpoint_interface_id IN (?)", endpointInterfaceIds).
		Count(&count).Error

	return
}

func (r *EndpointRepo) CreateEndpoints(endpoints []*model.Endpoint) error {
	return r.DB.Create(endpoints).Error
}

func (r *EndpointRepo) BatchUpdateStatus(ids []uint, status int64) error {
	return r.DB.Model(&model.Endpoint{}).Where("id IN (?)", ids).Update("status", status).Error
}

func (r *EndpointRepo) BatchUpdateCategory(ids []uint, categoryId int64) error {
	return r.DB.Model(&model.Endpoint{}).Where("id IN (?)", ids).Update("category_id", categoryId).Error
}

func (r *EndpointRepo) GetByItem(sourceType consts.SourceType, projectId uint, path string, serveId uint, title string) (res model.Endpoint, err error) {

	err = r.DB.First(&res, "source_type=? and project_id=? AND path = ? AND serve_id = ? AND title = ?", sourceType, projectId, path, serveId, title).Error

	return

}

func (r *EndpointRepo) ListByProjectIdAndServeId(projectId, serveId uint, needDetail bool) (endpoints []*model.Endpoint, err error) {
	err = r.DB.Where("project_id = ? and serve_id = ? and not deleted and not disabled", projectId, serveId).Find(&endpoints).Error
	//r.GetByEndpoints(endpoints, needDetail)
	return
}

func (r *EndpointRepo) GetByPath(serveId uint, pth string) (ret model.Endpoint, err error) {
	db := r.DB.Model(&ret).
		Where("path = ? AND NOT deleted", pth)

	if serveId > 0 {
		db.Where("serve_id = ?", serveId)
	}

	err = db.First(&ret).Error

	return
}
