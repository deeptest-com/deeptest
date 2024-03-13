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
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"strings"
)

type EndpointInterfaceRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`

	EndpointRepo       *EndpointRepo       `inject:""`
	DebugInterfaceRepo *DebugInterfaceRepo `inject:""`
	EnvironmentRepo    *EnvironmentRepo    `inject:""`
}

func (r *EndpointInterfaceRepo) Paginate(tenantId consts.TenantId, req v1.EndpointInterfaceReqPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.GetDB(tenantId).Model(&model.EndpointInterface{}).
		Joins("LEFT JOIN biz_endpoint e ON biz_endpoint_interface.endpoint_id=e.id").
		Where("biz_endpoint_interface.project_id = ?", req.ProjectId).
		Where("NOT biz_endpoint_interface.deleted AND NOT biz_endpoint_interface.disabled")

	if req.Keywords != "" {
		db = db.Where("biz_endpoint_interface.name LIKE ? or biz_endpoint_interface.url LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords), fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if len(req.ServeIds) != 0 {
		db = db.Where("e.serve_id in ?", req.ServeIds)
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

	db = db.Order("biz_endpoint_interface.created_at desc")
	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count error %s", err.Error())
		return
	}

	results := make([]*model.EndpointInterface, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}

	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *EndpointInterfaceRepo) ListByProject(tenantId consts.TenantId, projectId int) (pos []*model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}
func (r *EndpointInterfaceRepo) ListIdByEndpoint(tenantId consts.TenantId, endpointId uint) (ids []uint, err error) {
	err = r.GetDB(tenantId).
		Model(model.EndpointInterface{}).
		Select("id").
		Where("endpoint_id=?", endpointId).
		Where("NOT deleted").
		Find(&ids).Error

	return
}

func (r *EndpointInterfaceRepo) ListIdByEndpoints(tenantId consts.TenantId, endpointIds []uint) (ids []uint, err error) {
	err = r.GetDB(tenantId).
		Model(model.EndpointInterface{}).
		Select("id").
		Where("endpoint_id IN (?)", endpointIds).
		Where("NOT deleted").
		Find(&ids).Error

	return
}

func (r *EndpointInterfaceRepo) Get(tenantId consts.TenantId, interfaceId uint) (po model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).
		Where("id=? AND NOT deleted", interfaceId).
		First(&po).Error
	return
}

func (r *EndpointInterfaceRepo) GetByMethod(tenantId consts.TenantId, endpointId uint, method consts.HttpMethod) (debugInterfaceId, endpointInterfaceId uint) {
	var po model.EndpointInterface

	r.GetDB(tenantId).Where("endpoint_id=? AND method=? AND  NOT deleted", endpointId, method).
		First(&po)

	endpointInterfaceId = po.ID
	debugInterfaceId = po.DebugInterfaceId

	return
}

func (r *EndpointInterfaceRepo) BatchGet(tenantId consts.TenantId, interfaceIds []uint) (fields []model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).Model(model.EndpointInterface{}).
		Where("id IN (?) AND NOT deleted", interfaceIds).
		Find(&fields).Error
	return
}

func (r *EndpointInterfaceRepo) GetIdAndModelMap(tenantId consts.TenantId, interfaceIds []uint) (res map[uint]model.EndpointInterface, err error) {
	res = make(map[uint]model.EndpointInterface)

	if len(interfaceIds) == 0 {
		return
	}

	interfaces, err := r.BatchGet(tenantId, interfaceIds)
	if err != nil {
		return
	}

	for _, v := range interfaces {
		res[v.ID] = v
	}
	return
}

func (r *EndpointInterfaceRepo) GetDetail(tenantId consts.TenantId, interfId uint) (interf model.EndpointInterface, err error) {
	if interfId > 0 {
		interf, err = r.Get(tenantId, interfId)
		interf.Params, _ = r.ListParams(tenantId, interfId)
		interf.Headers, _ = r.ListHeaders(tenantId, interfId)
		interf.Cookies, _ = r.ListCookies(tenantId, interfId)
		interf.RequestBody, _ = r.ListRequestBody(tenantId, interfId)
		interf.ResponseBodies, _ = r.ListResponseBodies(tenantId, interfId)
		interf.GlobalParams, _ = r.GetGlobalParams(tenantId, interfId, interf.ProjectId)
	}
	return
}

func (r *EndpointInterfaceRepo) Save(tenantId consts.TenantId, interf *model.EndpointInterface) (err error) {
	err = r.GetDB(tenantId).Save(interf).Error
	return
}

func (r *EndpointInterfaceRepo) SetDebugInterfaceId(tenantId consts.TenantId, endpointInterfaceId, debugInterfaceId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointInterface{}).
		Where("id = ?", endpointInterfaceId).
		Update("debug_interface_id", debugInterfaceId).Error

	return
}

func (r *EndpointInterfaceRepo) Update(tenantId consts.TenantId, interf model.EndpointInterface) (err error) {
	r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.GetDB(tenantId).Updates(interf).Error
		if err != nil {
			return err
		}

		err = r.UpdateParams(tenantId, interf.ID, interf.Params)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(tenantId, interf.ID, interf.Headers)
		if err != nil {
			return err
		}
		return err
	})

	return
}
func (r *EndpointInterfaceRepo) UpdateName(tenantId consts.TenantId, id int, name string) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointInterface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *EndpointInterfaceRepo) UpdateNameByEndpointId(tenantId consts.TenantId, endpointId uint, name string) (err error) {
	return r.GetDB(tenantId).Model(&model.EndpointInterface{}).Where("endpoint_id = ?", endpointId).Update("name", name).Error
}

func (r *EndpointInterfaceRepo) UpdateHeaders(tenantId consts.TenantId, id uint, headers []model.EndpointInterfaceHeader) (err error) {
	err = r.RemoveHeaders(tenantId, id)

	if len(headers) == 0 {
		return
	}

	for idx, _ := range headers {
		headers[idx].ID = 0
		headers[idx].InterfaceId = id
	}

	err = r.GetDB(tenantId).Create(&headers).Error

	return
}

func (r *EndpointInterfaceRepo) RemoveHeaders(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.EndpointInterfaceHeader{}, "").Error

	return
}

func (r *EndpointInterfaceRepo) UpdateParams(tenantId consts.TenantId, id uint, params []model.EndpointInterfaceParam) (err error) {
	err = r.RemoveParams(tenantId, id)

	if len(params) == 0 {
		return
	}

	for idx, _ := range params {
		params[idx].ID = 0
		params[idx].InterfaceId = id
		params[idx].Value = params[idx].Example
	}

	err = r.GetDB(tenantId).Create(&params).Error

	return
}
func (r *EndpointInterfaceRepo) RemoveParams(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.EndpointInterfaceParam{}, "").Error

	return
}

func (r *EndpointInterfaceRepo) UpdateCookies(tenantId consts.TenantId, id uint, cookies []model.EndpointInterfaceCookie) (err error) {
	err = r.RemoveCookie(tenantId, id)

	if len(cookies) == 0 {
		return
	}

	for idx, _ := range cookies {
		cookies[idx].ID = 0
		cookies[idx].InterfaceId = id
	}

	err = r.GetDB(tenantId).Create(&cookies).Error

	return
}

func (r *EndpointInterfaceRepo) RemoveCookie(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.EndpointInterfaceCookie{}, "").Error

	return
}

func (r *EndpointInterfaceRepo) haveChild(Data []*model.EndpointInterface, node *model.EndpointInterface) (children []*model.EndpointInterface, yes bool) {
	for _, v := range Data {
		if v.ParentId == node.ID {
			v.Slots = iris.Map{"icon": "icon"}
			children = append(children, v)
		}
	}
	if children != nil {
		yes = true
	}
	return
}

func (r *EndpointInterfaceRepo) GetChildren(tenantId consts.TenantId, defId, fieldId uint) (children []*model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *EndpointInterfaceRepo) SetIsRange(tenantId consts.TenantId, fieldId uint, b bool) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointInterface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *EndpointInterfaceRepo) ListParams(tenantId consts.TenantId, interfaceId uint) (pos []model.EndpointInterfaceParam, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error
	return
}
func (r *EndpointInterfaceRepo) ListHeaders(tenantId consts.TenantId, interfaceId uint) (pos []model.EndpointInterfaceHeader, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *EndpointInterfaceRepo) ListCookies(tenantId consts.TenantId, interfaceId uint) (pos []model.EndpointInterfaceCookie, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *EndpointInterfaceRepo) SaveInterfaces(tenantId consts.TenantId, interf *model.EndpointInterface) (err error) {

	r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.UpdateInterface(tenantId, interf)
		if err != nil {
			return err
		}

		err = r.UpdateParams(tenantId, interf.ID, interf.Params)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(tenantId, interf.ID, interf.Headers)
		if err != nil {
			return err
		}

		err = r.UpdateCookies(tenantId, interf.ID, interf.Cookies)
		if err != nil {
			return err
		}

		interf.RequestBody.InterfaceId = interf.ID
		err = r.UpdateRequestBody(tenantId, &interf.RequestBody)
		if err != nil {
			return err
		}

		err = r.UpdateResponseBodies(tenantId, interf.ID, interf.ResponseBodies)
		if err != nil {
			return err
		}

		err = r.saveEndpointGlobalParams(tenantId, interf.ID, interf.GlobalParams)
		if err != nil {
			return err
		}

		return err
	})

	return
}

func (r *EndpointInterfaceRepo) UpdateRequestBody(tenantId consts.TenantId, requestBody *model.EndpointInterfaceRequestBody) (err error) {
	err = r.removeRequestBody(tenantId, requestBody.InterfaceId)
	if err != nil {
		return
	}

	err = r.BaseRepo.Save(tenantId, requestBody.ID, requestBody)
	if err != nil {
		return
	}

	schemaItem := requestBody.SchemaItem
	schemaItem.RequestBodyId = requestBody.ID
	err = r.removeRequestBodyItem(tenantId, requestBody.ID)
	if err != nil {
		return
	}

	err = r.BaseRepo.Save(tenantId, schemaItem.ID, &schemaItem)

	return
}

func (r *EndpointInterfaceRepo) removeRequestBodyItem(tenantId consts.TenantId, requestBodyId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("request_body_id = ?", requestBodyId).
		Delete(&model.EndpointInterfaceRequestBodyItem{}).Error

	return
}

func (r *EndpointInterfaceRepo) removeRequestBody(tenantId consts.TenantId, interId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", interId).
		Delete(&model.EndpointInterfaceRequestBody{}).Error

	return
}

func (r *EndpointInterfaceRepo) UpdateResponseBodies(tenantId consts.TenantId, interfaceId uint, responseBodies []model.EndpointInterfaceResponseBody) (err error) {
	err = r.removeResponseBodies(tenantId, interfaceId)
	if err != nil {
		return
	}
	codes := map[string]bool{}
	for _, responseBody := range responseBodies {
		if _, ok := codes[responseBody.Code]; ok {
			continue
		}
		responseBody.InterfaceId = interfaceId

		err = r.BaseRepo.Save(tenantId, responseBody.ID, &responseBody)
		if err != nil {
			return
		}

		err = r.removeResponseBodyItem(tenantId, responseBody.ID)
		if err != nil {
			return
		}

		schemaItem := responseBody.SchemaItem
		schemaItem.ResponseBodyId = responseBody.ID
		err = r.BaseRepo.Save(tenantId, schemaItem.ID, &schemaItem)
		if err != nil {
			return
		}

		err = r.removeResponseBodyHeader(tenantId, responseBody.ID)
		if err != nil {
			return
		}

		responseBodyHeaders := responseBody.Headers
		for _, header := range responseBodyHeaders {
			header.ResponseBodyId = responseBody.ID
			err = r.BaseRepo.Save(tenantId, header.ID, &header)
			if err != nil {
				return
			}
		}
		codes[responseBody.Code] = true
	}
	return
}

func (r *EndpointInterfaceRepo) removeResponseBodies(tenantId consts.TenantId, interId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", interId).
		Delete(&model.EndpointInterfaceResponseBody{}).Error

	return
}

func (r *EndpointInterfaceRepo) removeResponseBodyItem(tenantId consts.TenantId, responseBodyId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("response_body_id = ?", responseBodyId).
		Delete(&model.EndpointInterfaceResponseBodyItem{}).Error

	return
}

func (r *EndpointInterfaceRepo) removeResponseBodyHeader(tenantId consts.TenantId, responseBodyId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("response_body_id = ?", responseBodyId).
		Delete(&model.EndpointInterfaceResponseBodyHeader{}).Error

	return
}

func (r *EndpointInterfaceRepo) UpdateInterface(tenantId consts.TenantId, interf *model.EndpointInterface) (err error) {
	err = r.BaseRepo.Save(tenantId, interf.ID, interf)
	return
}

func (r *EndpointInterfaceRepo) ListByEndpointId(tenantId consts.TenantId, endpointId uint, version string) (interfaces []model.EndpointInterface, err error) {
	interfaces, err = r.QueryByEndpointId(tenantId, endpointId, version)
	for key, interf := range interfaces {
		interfaces[key].Params, _ = r.ListParams(tenantId, interf.ID)
		interfaces[key].Headers, _ = r.ListHeaders(tenantId, interf.ID)
		interfaces[key].Cookies, _ = r.ListCookies(tenantId, interf.ID)
		interfaces[key].RequestBody, _ = r.ListRequestBody(tenantId, interf.ID)
		interfaces[key].ResponseBodies, _ = r.ListResponseBodies(tenantId, interf.ID)
		interfaces[key].GlobalParams, _ = r.GetGlobalParams(tenantId, interf.ID, interf.ProjectId)
		//interfaces[key].ResponseCodes = strings.Split(interf.ResponseCodes.(string), ",")
	}

	return
}

func (r *EndpointInterfaceRepo) QueryByEndpointId(tenantId consts.TenantId, endpointId uint, version string) (pos []model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).
		Where("endpoint_id=? and version=?", endpointId, version).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *EndpointInterfaceRepo) ListRequestBody(tenantId consts.TenantId, interfaceId uint) (requestBody model.EndpointInterfaceRequestBody, err error) {
	err = r.GetDB(tenantId).First(&requestBody, "interface_id = ?", interfaceId).Error
	if err != nil {
		return
	}

	requestBody.SchemaItem, err = r.ListRequestBodyItem(tenantId, requestBody.ID)
	if err != nil {
		//requestBody.SchemaItem.Content = _commUtils.JsonDecode(builtin.Interface2String(requestBody.SchemaItem.Content))
	}
	return
}
func (r *EndpointInterfaceRepo) ListRequestBodyItem(tenantId consts.TenantId, requestBodyId uint) (requestBodyItem model.EndpointInterfaceRequestBodyItem, err error) {
	err = r.GetDB(tenantId).First(&requestBodyItem, "request_body_id = ?", requestBodyId).Error
	return
}

func (r *EndpointInterfaceRepo) ListResponseBodies(tenantId consts.TenantId, interfaceId uint) (responseBodies []model.EndpointInterfaceResponseBody, err error) {
	err = r.GetDB(tenantId).Find(&responseBodies, "interface_id = ?", interfaceId).Error
	if err != nil {
		return
	}

	for key, responseBody := range responseBodies {
		responseBodies[key].SchemaItem, err = r.ListResponseBodyItem(tenantId, responseBody.ID)
		if err != nil {
			return
		}

		responseBodies[key].Headers, err = r.ListResponseBodyHeaders(tenantId, responseBody.ID)
		if err != nil {
			return
		}
	}

	return
}

func (r *EndpointInterfaceRepo) ListResponseBodyItem(tenantId consts.TenantId, requestBodyId uint) (responseBodyItem model.EndpointInterfaceResponseBodyItem, err error) {
	err = r.GetDB(tenantId).First(&responseBodyItem, "response_body_id = ?", requestBodyId).Error
	//fmt.Println(err)
	return
}

func (r *EndpointInterfaceRepo) ListResponseBodyHeaders(tenantId consts.TenantId, requestBodyId uint) (responseBodyHeaders []model.EndpointInterfaceResponseBodyHeader, err error) {
	err = r.GetDB(tenantId).Find(&responseBodyHeaders, "response_body_id = ?", requestBodyId).Error
	return
}

func (r *EndpointInterfaceRepo) GetRequestBody(tenantId consts.TenantId, interfaceId uint) (result model.EndpointInterfaceRequestBody, err error) {
	err = r.GetDB(tenantId).Find(&result, "interface_id = ? AND NOT deleted", interfaceId).Error
	if err != nil {
		return
	}

	return
}

func (r *EndpointInterfaceRepo) GetCountByRef(tenantId consts.TenantId, ref string) (count int64, err error) {

	models := []interface{}{&model.EndpointPathParam{}, &model.EndpointInterfaceParam{}, &model.EndpointInterfaceHeader{}, &model.EndpointInterfaceCookie{}}

	for _, model := range models {
		err = r.GetDB(tenantId).Model(&model).Where("ref=?", ref).Count(&count).Error
		if err != nil {
			return
		}
		if count > 0 {
			return
		}
	}

	return
}

func (r *EndpointInterfaceRepo) ImportEndpointData(tenantId consts.TenantId, req v1.ImportEndpointDataReq) (err error) {
	/*
		if req.OpenUrlImport {

		} else {
			data, err := ioutil.ReadFile(req.FilePath)
			if err != nil {
				logUtils.Errorf("load end point data err ", zap.String("错误:", err.Error()))
				return err
			}
			var endpointData interface{}
			switch req.DriverType {
			case convert.POSTMAN:
				endpointData = postman.Doc{}
			case convert.YAPI:
				endpointData = yapi.Doc{}
			case convert.SWAGGER:
				endpointData = swagger.Doc{}
			}
			err = json.Unmarshal(data, &endpointData)
			if err != nil {
				logUtils.Errorf("unmarshall endpoint data err ", zap.String("错误:", err.Error()))
				return err
			}
		}
	*/
	return
}

func (r *EndpointInterfaceRepo) GetInterfaces(tenantId consts.TenantId, endpointIds []uint, needDetail bool) (interfaces map[uint][]model.EndpointInterface, err error) {
	interfaces = map[uint][]model.EndpointInterface{}
	var result []model.EndpointInterface
	err = r.GetDB(tenantId).Where("endpoint_id in ?", endpointIds).Where("NOT deleted").Find(&result).Error

	if !needDetail {
		for key, item := range result {
			interfaces[item.EndpointId] = append(interfaces[item.EndpointId], result[key])
		}
		return
	}

	var interfaceIds []uint
	for _, item := range result {
		interfaceIds = append(interfaceIds, item.ID)
	}

	var params map[uint][]model.EndpointInterfaceParam
	params, err = r.GetQueryParams(tenantId, interfaceIds)
	if err != nil {
		return
	}

	var cookies map[uint][]model.EndpointInterfaceCookie
	cookies, err = r.GetCookies(tenantId, interfaceIds)
	if err != nil {
		return
	}

	var headers map[uint][]model.EndpointInterfaceHeader
	headers, err = r.GetHeaders(tenantId, interfaceIds)
	if err != nil {
		return
	}

	var requestBodies map[uint]model.EndpointInterfaceRequestBody
	requestBodies, err = r.GetRequestBodies(tenantId, interfaceIds)
	if err != nil {
		return
	}

	var responseBodies map[uint][]model.EndpointInterfaceResponseBody
	responseBodies, err = r.GetResponseBodies(tenantId, interfaceIds)

	var globalParams map[uint][]model.EndpointInterfaceGlobalParam
	globalParams, err = r.GetMapGlobalParams(tenantId, interfaceIds)

	for key, item := range result {
		result[key].Params = params[item.ID]
		result[key].Cookies = cookies[item.ID]
		result[key].Headers = headers[item.ID]
		result[key].RequestBody = requestBodies[item.ID]
		result[key].ResponseBodies = responseBodies[item.ID]
		result[key].GlobalParams = globalParams[item.ID]
		interfaces[item.EndpointId] = append(interfaces[item.EndpointId], result[key])
	}

	return

}

func (r *EndpointInterfaceRepo) GetQueryParams(tenantId consts.TenantId, interfaceIds []uint) (params map[uint][]model.EndpointInterfaceParam, err error) {
	var result []model.EndpointInterfaceParam
	err = r.GetDB(tenantId).Where("NOT deleted and interface_id in ?", interfaceIds).Find(&result).Error

	params = make(map[uint][]model.EndpointInterfaceParam)
	for key, item := range result {
		params[item.InterfaceId] = append(params[item.InterfaceId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) GetCookies(tenantId consts.TenantId, interfaceIds []uint) (cookies map[uint][]model.EndpointInterfaceCookie, err error) {
	var result []model.EndpointInterfaceCookie
	err = r.GetDB(tenantId).Where("NOT deleted and interface_id in ?", interfaceIds).Find(&result).Error

	cookies = make(map[uint][]model.EndpointInterfaceCookie)
	for key, item := range result {
		cookies[item.InterfaceId] = append(cookies[item.InterfaceId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) GetHeaders(tenantId consts.TenantId, interfaceIds []uint) (headers map[uint][]model.EndpointInterfaceHeader, err error) {
	var result []model.EndpointInterfaceHeader
	err = r.GetDB(tenantId).Where("NOT deleted and interface_id in ?", interfaceIds).Find(&result).Error

	headers = make(map[uint][]model.EndpointInterfaceHeader)
	for key, item := range result {
		headers[item.InterfaceId] = append(headers[item.InterfaceId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) GetRequestBodies(tenantId consts.TenantId, interfaceIds []uint) (requestBodies map[uint]model.EndpointInterfaceRequestBody, err error) {
	var result []model.EndpointInterfaceRequestBody
	err = r.GetDB(tenantId).Find(&result, "interface_id in ?", interfaceIds).Error
	if err != nil {
		return
	}

	var requestBodyIds []uint
	for _, item := range result {
		requestBodyIds = append(requestBodyIds, item.ID)
	}

	var requestBodyItems map[uint]model.EndpointInterfaceRequestBodyItem
	requestBodyItems, err = r.GetRequestBodyItems(tenantId, requestBodyIds)

	requestBodies = make(map[uint]model.EndpointInterfaceRequestBody)
	for key, item := range result {
		result[key].SchemaItem = requestBodyItems[item.ID]
		requestBodies[item.InterfaceId] = result[key]
	}

	return
}

func (r *EndpointInterfaceRepo) GetRequestBodyItems(tenantId consts.TenantId, requestBodyIds []uint) (requestBodyItems map[uint]model.EndpointInterfaceRequestBodyItem, err error) {
	var result []model.EndpointInterfaceRequestBodyItem
	err = r.GetDB(tenantId).Find(&result, "request_body_id in ?", requestBodyIds).Error

	requestBodyItems = make(map[uint]model.EndpointInterfaceRequestBodyItem)
	for key, item := range result {
		requestBodyItems[item.RequestBodyId] = result[key]
	}
	return
}

func (r *EndpointInterfaceRepo) GetRequestBodyItem(tenantId consts.TenantId, requestBodyId uint) (result model.EndpointInterfaceRequestBodyItem, err error) {
	err = r.GetDB(tenantId).Find(&result, "request_body_id = ?  AND NOT deleted", requestBodyId).Error
	if err != nil {
		return
	}

	return
}

func (r *EndpointInterfaceRepo) GetResponseBodies(tenantId consts.TenantId, interfaceIds []uint) (responseBodies map[uint][]model.EndpointInterfaceResponseBody, err error) {
	var result []model.EndpointInterfaceResponseBody
	err = r.GetDB(tenantId).Find(&result, "interface_id in ?", interfaceIds).Error
	if err != nil {
		return
	}

	var responseBodyIds []uint
	for _, item := range result {
		responseBodyIds = append(responseBodyIds, item.ID)
	}

	var responseBodyItems map[uint]model.EndpointInterfaceResponseBodyItem
	responseBodyItems, err = r.GetResponseBodyItems(tenantId, responseBodyIds)
	if err != nil {
		return
	}

	var responseBodyHeaders map[uint][]model.EndpointInterfaceResponseBodyHeader
	responseBodyHeaders, err = r.GetResponseBodyHeaders(tenantId, responseBodyIds)
	if err != nil {
		return
	}

	responseBodies = make(map[uint][]model.EndpointInterfaceResponseBody)
	for key, item := range result {
		result[key].SchemaItem = responseBodyItems[item.ID]
		result[key].Headers = responseBodyHeaders[item.ID]
		responseBodies[item.InterfaceId] = append(responseBodies[item.InterfaceId], result[key])
	}

	return

}

func (r *EndpointInterfaceRepo) GetResponseBodyItems(tenantId consts.TenantId, responseBodyIds []uint) (responseBodyItem map[uint]model.EndpointInterfaceResponseBodyItem, err error) {
	var result []model.EndpointInterfaceResponseBodyItem
	err = r.GetDB(tenantId).Find(&result, "response_body_id in ?", responseBodyIds).Error

	responseBodyItem = make(map[uint]model.EndpointInterfaceResponseBodyItem)
	for key, item := range result {
		responseBodyItem[item.ResponseBodyId] = result[key]
	}

	return
}

func (r *EndpointInterfaceRepo) GetResponseBodyHeaders(tenantId consts.TenantId, responseBodyIds []uint) (responseBodyHeaders map[uint][]model.EndpointInterfaceResponseBodyHeader, err error) {

	var result []model.EndpointInterfaceResponseBodyHeader
	err = r.GetDB(tenantId).Find(&result, "response_body_id in ?", responseBodyIds).Error

	responseBodyHeaders = make(map[uint][]model.EndpointInterfaceResponseBodyHeader)
	for key, item := range result {
		responseBodyHeaders[item.ResponseBodyId] = append(responseBodyHeaders[item.ResponseBodyId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) DeleteByEndpoint(tenantId consts.TenantId, endpointId uint) (err error) {
	ids, err := r.ListIdByEndpoint(tenantId, endpointId)

	err = r.DeleteBatch(tenantId, ids)

	return
}

func (r *EndpointInterfaceRepo) DeleteByEndpoints(tenantId consts.TenantId, endpointIds []uint) (err error) {
	ids, err := r.ListIdByEndpoints(tenantId, endpointIds)

	err = r.DeleteBatch(tenantId, ids)

	return
}

func (r *EndpointInterfaceRepo) DeleteBatch(tenantId consts.TenantId, ids []uint) (err error) {
	for _, id := range ids {
		err = r.Delete(tenantId, id)
	}
	return
}
func (r *EndpointInterfaceRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	endpointInterface, _ := r.Get(tenantId, id)
	if endpointInterface.DebugInterfaceId > 0 {
		r.DebugInterfaceRepo.Delete(tenantId, endpointInterface.DebugInterfaceId)
	}
	return
}

func (r *EndpointInterfaceRepo) RemoveAll(tenantId consts.TenantId, id uint) (err error) {
	err = r.RemoveParams(tenantId, id)
	err = r.RemoveHeaders(tenantId, id)
	err = r.RemoveCookie(tenantId, id)

	err = r.removeRequestBody(tenantId, id)
	err = r.removeResponseBodies(tenantId, id)

	requestBody, err := r.ListRequestBody(tenantId, id)
	err = r.removeRequestBodyItem(tenantId, requestBody.ID)

	responseBodies, err := r.ListResponseBodies(tenantId, id)
	for _, body := range responseBodies {
		err = r.removeResponseBodyItem(tenantId, body.ID)
		err = r.removeResponseBodyHeader(tenantId, body.ID)
	}
	return
}

func (r *EndpointInterfaceRepo) GetByItem(tenantId consts.TenantId, sourceType consts.SourceType, endpointId uint, method consts.HttpMethod) (res model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).First(&res, "not deleted AND source_type = ? AND endpoint_id=? and method=?", sourceType, endpointId, method).Error
	return
}

func (r *EndpointInterfaceRepo) GetResponseCodes(tenantId consts.TenantId, endpointInterfaceId uint) (codes []string) {

	responseBodies, err := r.GetResponseBodies(tenantId, []uint{endpointInterfaceId})
	if err != nil {
		return
	}
	responseBody := responseBodies[endpointInterfaceId]

	for _, item := range responseBody {
		codes = append(codes, item.Code)
	}

	return
}

func (r *EndpointInterfaceRepo) GetResponse(tenantId consts.TenantId, endpointInterfaceId uint, code string) (ret model.EndpointInterfaceResponseBody) {
	responseBodies, err := r.GetResponseBodies(tenantId, []uint{endpointInterfaceId})
	if err != nil {
		return
	}

	responseBody := responseBodies[endpointInterfaceId]
	for _, item := range responseBody {
		if code == item.Code {
			return item
		}
	}

	return
}

func (r *EndpointInterfaceRepo) BatchGetByEndpointIds(tenantId consts.TenantId, endpointIds []uint) (fields []model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).Model(model.EndpointInterface{}).
		Where("endpoint_id IN (?) AND NOT deleted", endpointIds).
		Find(&fields).Error
	return
}

func (r *EndpointInterfaceRepo) GetByMethodAndPathAndServeId(tenantId consts.TenantId, serveId uint, path string, method consts.HttpMethod) (endpointInterfaceId uint) {

	type result struct {
		Id uint
	}
	var data result

	err := r.GetDB(tenantId).Model(&model.EndpointInterface{}).Select("biz_endpoint_interface.id").Joins("left join biz_endpoint on biz_endpoint_interface.endpoint_id = biz_endpoint.id").Where("not biz_endpoint.deleted and not biz_endpoint_interface.deleted and biz_endpoint.serve_id=? and biz_endpoint.path=? and biz_endpoint_interface.method=?", serveId, path, method).Scan(&data).Error
	if err != nil {
		return 0
	}
	return data.Id

}

func (r *EndpointInterfaceRepo) GetResponseDefine(tenantId consts.TenantId, endpointId uint, method consts.HttpMethod, code string) (ret model.EndpointInterfaceResponseBodyItem, err error) {

	var endpointInterface model.EndpointInterface
	err = r.GetDB(tenantId).Where("endpoint_id=? AND method=? AND  NOT deleted", endpointId, method).First(&endpointInterface).Error
	if err != nil {
		return
	}

	response := r.GetResponse(tenantId, endpointInterface.ID, code)
	if response.ID == 0 {
		return
	}

	var bodyItems map[uint]model.EndpointInterfaceResponseBodyItem
	bodyItems, err = r.GetResponseBodyItems(tenantId, []uint{response.ID})
	if err != nil {
		return
	}

	return bodyItems[response.ID], nil

}

func (r *EndpointInterfaceRepo) GetByEndpointId(tenantId consts.TenantId, endpointId uint) (fields []model.EndpointInterface, err error) {
	err = r.GetDB(tenantId).Model(model.EndpointInterface{}).
		Where("endpoint_id = ? AND NOT deleted", endpointId).
		Find(&fields).Error
	return
}

func (r *EndpointInterfaceRepo) GetResponseBody(tenantId consts.TenantId, interfaceId uint, code string) (result model.EndpointInterfaceResponseBody, err error) {
	err = r.GetDB(tenantId).Find(&result, "interface_id = ? AND code = ?  AND NOT deleted", interfaceId, code).Error
	if err != nil {
		return
	}

	return
}

func (r *EndpointInterfaceRepo) GetResponseBodyItem(tenantId consts.TenantId, responseBodyId uint) (result model.EndpointInterfaceResponseBodyItem, err error) {
	err = r.GetDB(tenantId).Find(&result, "response_body_id = ?  AND NOT deleted", responseBodyId).Error
	if err != nil {
		return
	}

	return
}

func (r *EndpointInterfaceRepo) SaveResponseBody(tenantId consts.TenantId, responseBody *model.EndpointInterfaceResponseBody) (err error) {

	err = r.BaseRepo.Save(tenantId, responseBody.ID, responseBody)
	if err != nil {
		return
	}

	err = r.removeResponseBodyItem(tenantId, responseBody.ID)
	if err != nil {
		return
	}

	schemaItem := responseBody.SchemaItem
	schemaItem.ResponseBodyId = responseBody.ID
	err = r.BaseRepo.Save(tenantId, schemaItem.ID, &schemaItem)
	if err != nil {
		return
	}

	err = r.AddResponseCode(tenantId, responseBody.InterfaceId, responseBody.Code)

	return

}

func (r *EndpointInterfaceRepo) AddResponseCode(tenantId consts.TenantId, interfaceId uint, code string) (err error) {
	var endpointInterface model.EndpointInterface
	endpointInterface, err = r.Get(tenantId, interfaceId)
	if err != nil {
		return err
	}

	needAdd := true
	if endpointInterface.ResponseCodes == "" {
		endpointInterface.ResponseCodes = "200,301,302,401,402,500"
	}
	responseCodes := strings.Split(endpointInterface.ResponseCodes, ",")
	for _, responseCode := range responseCodes {
		if responseCode == code {
			needAdd = false
			break
		}
	}
	if needAdd {
		responseCodes = append(responseCodes, code)
	}
	err = r.GetDB(tenantId).Model(endpointInterface).Update("response_codes", strings.Join(responseCodes, ",")).Where("id", interfaceId).Error

	return

}

// 保存路径参数
func (r *EndpointInterfaceRepo) saveEndpointGlobalParams(tenantId consts.TenantId, interfaceId uint, params []model.EndpointInterfaceGlobalParam) (err error) {
	err = r.removeEndpointGlobalParams(tenantId, interfaceId)
	if err != nil {
		return
	}
	for key, _ := range params {
		params[key].InterfaceId = interfaceId
	}

	if len(params) == 0 {
		return
	}

	err = r.GetDB(tenantId).Create(params).Error
	return
}

func (r *EndpointInterfaceRepo) removeEndpointGlobalParams(tenantId consts.TenantId, interfaceId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", interfaceId).
		Delete(&model.EndpointInterfaceGlobalParam{}, "").Error

	return
}

func (r *EndpointInterfaceRepo) GetGlobalParams(tenantId consts.TenantId, id, projectId uint) (ret []model.EndpointInterfaceGlobalParam, err error) {
	var po []model.EndpointInterfaceGlobalParam
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Find(&po).Error

	if err != nil {
		return nil, err
	}

	if params, err := r.EnvironmentRepo.ListParamModel(tenantId, projectId); err == nil {
		for _, param := range params {
			var temp model.EndpointInterfaceGlobalParam
			copier.CopyWithOption(&temp, &param, copier.Option{IgnoreEmpty: true, DeepCopy: true})
			for _, item := range po {
				if param.Name == item.Name && param.In == item.In {
					temp.Disabled = item.Disabled
				}
			}
			ret = append(ret, temp)
		}
	}

	return

}

func (r *EndpointInterfaceRepo) GetMapGlobalParams(tenantId consts.TenantId, interfaceIds []uint) (params map[uint][]model.EndpointInterfaceGlobalParam, err error) {
	var result []model.EndpointInterfaceGlobalParam
	err = r.GetDB(tenantId).Where("interface_id in ?", interfaceIds).Find(&result).Error

	params = make(map[uint][]model.EndpointInterfaceGlobalParam)
	for key, item := range result {
		params[item.InterfaceId] = append(params[item.InterfaceId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) GetMethodsByEndpointId(tenantId consts.TenantId, endpointId uint) (ret []string, err error) {
	err = r.GetDB(tenantId).Model(model.EndpointInterface{}).Select("method").Where("endpoint_id = ? and not deleted", endpointId).Scan(&ret).Error
	return
}
