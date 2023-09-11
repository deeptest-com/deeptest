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
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type EndpointInterfaceRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`

	EndpointRepo       *EndpointRepo       `inject:""`
	DebugInterfaceRepo *DebugInterfaceRepo `inject:""`
}

func (r *EndpointInterfaceRepo) Paginate(req v1.EndpointInterfaceReqPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.EndpointInterface{}).
		Joins("LEFT JOIN biz_endpoint e ON biz_endpoint_interface.endpoint_id=e.id").
		Where("biz_endpoint_interface.project_id = ?", req.ProjectId).
		Where("NOT biz_endpoint_interface.deleted AND NOT biz_endpoint_interface.disabled")

	if req.Keywords != "" {
		db = db.Where("biz_endpoint_interface.name LIKE ? or biz_endpoint_interface.url LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords), fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.ServeId != 0 {
		db = db.Where("e.serve_id = ?", req.ServeId)
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

func (r *EndpointInterfaceRepo) ListByProject(projectId int) (pos []*model.EndpointInterface, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}
func (r *EndpointInterfaceRepo) ListIdByEndpoint(endpointId uint) (ids []uint, err error) {
	err = r.DB.
		Model(model.EndpointInterface{}).
		Select("id").
		Where("endpoint_id=?", endpointId).
		Where("NOT deleted").
		Find(&ids).Error

	return
}

func (r *EndpointInterfaceRepo) ListIdByEndpoints(endpointIds []uint) (ids []uint, err error) {
	err = r.DB.
		Model(model.EndpointInterface{}).
		Select("id").
		Where("endpoint_id IN (?)", endpointIds).
		Where("NOT deleted").
		Find(&ids).Error

	return
}

func (r *EndpointInterfaceRepo) Get(interfaceId uint) (po model.EndpointInterface, err error) {
	err = r.DB.
		Where("id=? AND NOT deleted", interfaceId).
		First(&po).Error
	return
}

func (r *EndpointInterfaceRepo) GetByMethod(endpointId uint, method consts.HttpMethod) (debugInterfaceId, endpointInterfaceId uint) {
	var po model.EndpointInterface

	r.DB.Where("endpoint_id=? AND method=? AND  NOT deleted", endpointId, method).
		First(&po)

	endpointInterfaceId = po.ID
	debugInterfaceId = po.DebugInterfaceId

	return
}

func (r *EndpointInterfaceRepo) BatchGet(interfaceIds []uint) (fields []model.EndpointInterface, err error) {
	err = r.DB.Model(model.EndpointInterface{}).
		Where("id IN (?) AND NOT deleted", interfaceIds).
		Find(&fields).Error
	return
}

func (r *EndpointInterfaceRepo) GetIdAndModelMap(interfaceIds []uint) (res map[uint]model.EndpointInterface, err error) {
	res = make(map[uint]model.EndpointInterface)

	if len(interfaceIds) == 0 {
		return
	}

	interfaces, err := r.BatchGet(interfaceIds)
	if err != nil {
		return
	}

	for _, v := range interfaces {
		res[v.ID] = v
	}
	return
}

func (r *EndpointInterfaceRepo) GetDetail(interfId uint) (interf model.EndpointInterface, err error) {
	if interfId > 0 {
		interf, err = r.Get(interfId)
		interf.Params, _ = r.ListParams(interfId)
		interf.Headers, _ = r.ListHeaders(interfId)
		interf.Cookies, _ = r.ListCookies(interfId)
		interf.RequestBody, _ = r.ListRequestBody(interfId)
		interf.ResponseBodies, _ = r.ListResponseBodies(interfId)
	}
	return
}

func (r *EndpointInterfaceRepo) Save(interf *model.EndpointInterface) (err error) {
	err = r.DB.Save(interf).Error
	return
}

func (r *EndpointInterfaceRepo) SetDebugInterfaceId(endpointInterfaceId, debugInterfaceId uint) (err error) {
	err = r.DB.Model(&model.EndpointInterface{}).
		Where("id = ?", endpointInterfaceId).
		Update("debug_interface_id", debugInterfaceId).Error

	return
}

func (r *EndpointInterfaceRepo) Update(interf model.EndpointInterface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Updates(interf).Error
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.Params)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(interf.ID, interf.Headers)
		if err != nil {
			return err
		}
		return err
	})

	return
}
func (r *EndpointInterfaceRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.EndpointInterface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *EndpointInterfaceRepo) UpdateHeaders(id uint, headers []model.EndpointInterfaceHeader) (err error) {
	err = r.RemoveHeaders(id)

	if len(headers) == 0 {
		return
	}

	for idx, _ := range headers {
		headers[idx].ID = 0
		headers[idx].InterfaceId = id
	}

	err = r.DB.Create(&headers).Error

	return
}

func (r *EndpointInterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.EndpointInterfaceHeader{}, "").Error

	return
}

func (r *EndpointInterfaceRepo) UpdateParams(id uint, params []model.EndpointInterfaceParam) (err error) {
	err = r.RemoveParams(id)

	if len(params) == 0 {
		return
	}

	for idx, _ := range params {
		params[idx].ID = 0
		params[idx].InterfaceId = id
		params[idx].Value = params[idx].Example
	}

	err = r.DB.Create(&params).Error

	return
}
func (r *EndpointInterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.EndpointInterfaceParam{}, "").Error

	return
}

func (r *EndpointInterfaceRepo) UpdateCookies(id uint, cookies []model.EndpointInterfaceCookie) (err error) {
	err = r.RemoveCookie(id)

	if len(cookies) == 0 {
		return
	}

	for idx, _ := range cookies {
		cookies[idx].ID = 0
		cookies[idx].InterfaceId = id
	}

	err = r.DB.Create(&cookies).Error

	return
}

func (r *EndpointInterfaceRepo) RemoveCookie(id uint) (err error) {
	err = r.DB.
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

func (r *EndpointInterfaceRepo) GetChildren(defId, fieldId uint) (children []*model.EndpointInterface, err error) {
	err = r.DB.Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *EndpointInterfaceRepo) SetIsRange(fieldId uint, b bool) (err error) {
	err = r.DB.Model(&model.EndpointInterface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *EndpointInterfaceRepo) ListParams(interfaceId uint) (pos []model.EndpointInterfaceParam, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error
	return
}
func (r *EndpointInterfaceRepo) ListHeaders(interfaceId uint) (pos []model.EndpointInterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *EndpointInterfaceRepo) ListCookies(interfaceId uint) (pos []model.EndpointInterfaceCookie, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *EndpointInterfaceRepo) SaveInterfaces(interf *model.EndpointInterface) (err error) {

	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.UpdateInterface(interf)
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.Params)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(interf.ID, interf.Headers)
		if err != nil {
			return err
		}

		err = r.UpdateCookies(interf.ID, interf.Cookies)
		if err != nil {
			return err
		}

		interf.RequestBody.InterfaceId = interf.ID
		err = r.UpdateRequestBody(&interf.RequestBody)
		if err != nil {
			return err
		}

		err = r.UpdateResponseBodies(interf.ID, interf.ResponseBodies)
		if err != nil {
			return err
		}

		return err
	})

	return
}

func (r *EndpointInterfaceRepo) UpdateRequestBody(requestBody *model.EndpointInterfaceRequestBody) (err error) {
	err = r.removeRequestBody(requestBody.InterfaceId)
	if err != nil {
		return
	}

	err = r.BaseRepo.Save(requestBody.ID, requestBody)
	if err != nil {
		return
	}

	schemaItem := requestBody.SchemaItem
	schemaItem.RequestBodyId = requestBody.ID
	err = r.removeRequestBodyItem(requestBody.ID)
	if err != nil {
		return
	}

	err = r.BaseRepo.Save(schemaItem.ID, &schemaItem)

	return
}

func (r *EndpointInterfaceRepo) removeRequestBodyItem(requestBodyId uint) (err error) {
	err = r.DB.
		Where("request_body_id = ?", requestBodyId).
		Delete(&model.EndpointInterfaceRequestBodyItem{}).Error

	return
}

func (r *EndpointInterfaceRepo) removeRequestBody(interId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interId).
		Delete(&model.EndpointInterfaceRequestBody{}).Error

	return
}

func (r *EndpointInterfaceRepo) UpdateResponseBodies(interfaceId uint, responseBodies []model.EndpointInterfaceResponseBody) (err error) {
	err = r.removeResponseBodies(interfaceId)
	if err != nil {
		return
	}
	codes := map[string]bool{}
	for _, responseBody := range responseBodies {
		if _, ok := codes[responseBody.Code]; ok {
			continue
		}
		responseBody.InterfaceId = interfaceId

		err = r.BaseRepo.Save(responseBody.ID, &responseBody)
		if err != nil {
			return
		}

		err = r.removeResponseBodyItem(responseBody.ID)
		if err != nil {
			return
		}

		schemaItem := responseBody.SchemaItem
		schemaItem.ResponseBodyId = responseBody.ID
		err = r.BaseRepo.Save(schemaItem.ID, &schemaItem)
		if err != nil {
			return
		}

		err = r.removeResponseBodyHeader(responseBody.ID)
		if err != nil {
			return
		}

		responseBodyHeaders := responseBody.Headers
		for _, header := range responseBodyHeaders {
			header.ResponseBodyId = responseBody.ID
			err = r.BaseRepo.Save(header.ID, &header)
			if err != nil {
				return
			}
		}
		codes[responseBody.Code] = true
	}
	return
}

func (r *EndpointInterfaceRepo) removeResponseBodies(interId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interId).
		Delete(&model.EndpointInterfaceResponseBody{}).Error

	return
}

func (r *EndpointInterfaceRepo) removeResponseBodyItem(responseBodyId uint) (err error) {
	err = r.DB.
		Where("response_body_id = ?", responseBodyId).
		Delete(&model.EndpointInterfaceResponseBodyItem{}).Error

	return
}

func (r *EndpointInterfaceRepo) removeResponseBodyHeader(responseBodyId uint) (err error) {
	err = r.DB.
		Where("response_body_id = ?", responseBodyId).
		Delete(&model.EndpointInterfaceResponseBodyHeader{}).Error

	return
}

func (r *EndpointInterfaceRepo) UpdateInterface(interf *model.EndpointInterface) (err error) {
	err = r.BaseRepo.Save(interf.ID, interf)
	return
}

func (r *EndpointInterfaceRepo) ListByEndpointId(endpointId uint, version string) (interfaces []model.EndpointInterface, err error) {
	interfaces, err = r.QueryByEndpointId(endpointId, version)
	for key, interf := range interfaces {
		interfaces[key].Params, _ = r.ListParams(interf.ID)
		interfaces[key].Headers, _ = r.ListHeaders(interf.ID)
		interfaces[key].Cookies, _ = r.ListCookies(interf.ID)
		interfaces[key].RequestBody, _ = r.ListRequestBody(interf.ID)
		interfaces[key].ResponseBodies, _ = r.ListResponseBodies(interf.ID)
		//interfaces[key].ResponseCodes = strings.Split(interf.ResponseCodes.(string), ",")
	}

	return
}

func (r *EndpointInterfaceRepo) QueryByEndpointId(endpointId uint, version string) (pos []model.EndpointInterface, err error) {
	err = r.DB.
		Where("endpoint_id=? and version=?", endpointId, version).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *EndpointInterfaceRepo) ListRequestBody(interfaceId uint) (requestBody model.EndpointInterfaceRequestBody, err error) {
	err = r.DB.First(&requestBody, "interface_id = ?", interfaceId).Error
	if err != nil {
		return
	}

	requestBody.SchemaItem, err = r.ListRequestBodyItem(requestBody.ID)
	if err != nil {
		//requestBody.SchemaItem.Content = _commUtils.JsonDecode(builtin.Interface2String(requestBody.SchemaItem.Content))
	}
	return
}
func (r *EndpointInterfaceRepo) ListRequestBodyItem(requestBodyId uint) (requestBodyItem model.EndpointInterfaceRequestBodyItem, err error) {
	err = r.DB.First(&requestBodyItem, "request_body_id = ?", requestBodyId).Error
	return
}

func (r *EndpointInterfaceRepo) ListResponseBodies(interfaceId uint) (responseBodies []model.EndpointInterfaceResponseBody, err error) {
	err = r.DB.Find(&responseBodies, "interface_id = ?", interfaceId).Error
	if err != nil {
		return
	}

	for key, responseBody := range responseBodies {
		responseBodies[key].SchemaItem, err = r.ListResponseBodyItem(responseBody.ID)
		if err != nil {
			return
		}

		responseBodies[key].Headers, err = r.ListResponseBodyHeaders(responseBody.ID)
		if err != nil {
			return
		}
	}

	return
}

func (r *EndpointInterfaceRepo) ListResponseBodyItem(requestBodyId uint) (responseBodyItem model.EndpointInterfaceResponseBodyItem, err error) {
	err = r.DB.First(&responseBodyItem, "response_body_id = ?", requestBodyId).Error
	//fmt.Println(err)
	return
}

func (r *EndpointInterfaceRepo) ListResponseBodyHeaders(requestBodyId uint) (responseBodyHeaders []model.EndpointInterfaceResponseBodyHeader, err error) {
	err = r.DB.Find(&responseBodyHeaders, "response_body_id = ?", requestBodyId).Error
	return
}

func (r *EndpointInterfaceRepo) GetCountByRef(ref string) (count int64, err error) {

	models := []interface{}{&model.EndpointPathParam{}, &model.EndpointInterfaceParam{}, &model.EndpointInterfaceHeader{}, &model.EndpointInterfaceCookie{}}

	for _, model := range models {
		err = r.DB.Model(&model).Where("ref=?", ref).Count(&count).Error
		if err != nil {
			return
		}
		if count > 0 {
			return
		}
	}

	return
}

func (r *EndpointInterfaceRepo) ImportEndpointData(req v1.ImportEndpointDataReq) (err error) {
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

func (r *EndpointInterfaceRepo) GetInterfaces(endpointIds []uint, needDetail bool) (interfaces map[uint][]model.EndpointInterface, err error) {
	interfaces = map[uint][]model.EndpointInterface{}
	var result []model.EndpointInterface
	err = r.DB.Where("endpoint_id in ?", endpointIds).Where("NOT deleted").Find(&result).Error

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
	params, err = r.GetQueryParams(interfaceIds)
	if err != nil {
		return
	}

	var cookies map[uint][]model.EndpointInterfaceCookie
	cookies, err = r.GetCookies(interfaceIds)
	if err != nil {
		return
	}

	var headers map[uint][]model.EndpointInterfaceHeader
	headers, err = r.GetHeaders(interfaceIds)
	if err != nil {
		return
	}

	var requestBodies map[uint]model.EndpointInterfaceRequestBody
	requestBodies, err = r.GetRequestBodies(interfaceIds)
	if err != nil {
		return
	}

	var responseBodies map[uint][]model.EndpointInterfaceResponseBody
	responseBodies, err = r.GetResponseBodies(interfaceIds)

	for key, item := range result {
		result[key].Params = params[item.ID]
		result[key].Cookies = cookies[item.ID]
		result[key].Headers = headers[item.ID]
		result[key].RequestBody = requestBodies[item.ID]
		result[key].ResponseBodies = responseBodies[item.ID]
		interfaces[item.EndpointId] = append(interfaces[item.EndpointId], result[key])
	}

	return

}

func (r *EndpointInterfaceRepo) GetQueryParams(interfaceIds []uint) (params map[uint][]model.EndpointInterfaceParam, err error) {
	var result []model.EndpointInterfaceParam
	err = r.DB.Where("NOT deleted and interface_id in ?", interfaceIds).Find(&result).Error

	params = make(map[uint][]model.EndpointInterfaceParam)
	for key, item := range result {
		params[item.InterfaceId] = append(params[item.InterfaceId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) GetCookies(interfaceIds []uint) (cookies map[uint][]model.EndpointInterfaceCookie, err error) {
	var result []model.EndpointInterfaceCookie
	err = r.DB.Where("NOT deleted and interface_id in ?", interfaceIds).Find(&result).Error

	cookies = make(map[uint][]model.EndpointInterfaceCookie)
	for key, item := range result {
		cookies[item.InterfaceId] = append(cookies[item.InterfaceId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) GetHeaders(interfaceIds []uint) (headers map[uint][]model.EndpointInterfaceHeader, err error) {
	var result []model.EndpointInterfaceHeader
	err = r.DB.Where("NOT deleted and interface_id in ?", interfaceIds).Find(&result).Error

	headers = make(map[uint][]model.EndpointInterfaceHeader)
	for key, item := range result {
		headers[item.InterfaceId] = append(headers[item.InterfaceId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) GetRequestBodies(interfaceIds []uint) (requestBodies map[uint]model.EndpointInterfaceRequestBody, err error) {
	var result []model.EndpointInterfaceRequestBody
	err = r.DB.Find(&result, "interface_id in ?", interfaceIds).Error
	if err != nil {
		return
	}

	var requestBodyIds []uint
	for _, item := range result {
		requestBodyIds = append(requestBodyIds, item.ID)
	}

	var requestBodyItems map[uint]model.EndpointInterfaceRequestBodyItem
	requestBodyItems, err = r.GetRequestBodyItems(requestBodyIds)

	requestBodies = make(map[uint]model.EndpointInterfaceRequestBody)
	for key, item := range result {
		result[key].SchemaItem = requestBodyItems[item.ID]
		requestBodies[item.InterfaceId] = result[key]
	}

	return
}

func (r *EndpointInterfaceRepo) GetRequestBodyItems(requestBodyIds []uint) (requestBodyItems map[uint]model.EndpointInterfaceRequestBodyItem, err error) {
	var result []model.EndpointInterfaceRequestBodyItem
	err = r.DB.Find(&result, "request_body_id in ?", requestBodyIds).Error

	requestBodyItems = make(map[uint]model.EndpointInterfaceRequestBodyItem)
	for key, item := range result {
		requestBodyItems[item.RequestBodyId] = result[key]
	}
	return
}

func (r *EndpointInterfaceRepo) GetResponseBodies(interfaceIds []uint) (responseBodies map[uint][]model.EndpointInterfaceResponseBody, err error) {
	var result []model.EndpointInterfaceResponseBody
	err = r.DB.Find(&result, "interface_id in ?", interfaceIds).Error
	if err != nil {
		return
	}

	var responseBodyIds []uint
	for _, item := range result {
		responseBodyIds = append(responseBodyIds, item.ID)
	}

	var responseBodyItems map[uint]model.EndpointInterfaceResponseBodyItem
	responseBodyItems, err = r.GetResponseBodyItems(responseBodyIds)
	if err != nil {
		return
	}

	var responseBodyHeaders map[uint][]model.EndpointInterfaceResponseBodyHeader
	responseBodyHeaders, err = r.GetResponseBodyHeaders(responseBodyIds)
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

func (r *EndpointInterfaceRepo) GetResponseBodyItems(responseBodyIds []uint) (responseBodyItem map[uint]model.EndpointInterfaceResponseBodyItem, err error) {
	var result []model.EndpointInterfaceResponseBodyItem
	err = r.DB.Find(&result, "response_body_id in ?", responseBodyIds).Error

	responseBodyItem = make(map[uint]model.EndpointInterfaceResponseBodyItem)
	for key, item := range result {
		responseBodyItem[item.ResponseBodyId] = result[key]
	}

	return
}

func (r *EndpointInterfaceRepo) GetResponseBodyHeaders(responseBodyIds []uint) (responseBodyHeaders map[uint][]model.EndpointInterfaceResponseBodyHeader, err error) {

	var result []model.EndpointInterfaceResponseBodyHeader
	err = r.DB.Find(&result, "response_body_id in ?", responseBodyIds).Error

	responseBodyHeaders = make(map[uint][]model.EndpointInterfaceResponseBodyHeader)
	for key, item := range result {
		responseBodyHeaders[item.ResponseBodyId] = append(responseBodyHeaders[item.ResponseBodyId], result[key])
	}

	return
}

func (r *EndpointInterfaceRepo) DeleteByEndpoint(endpointId uint) (err error) {
	ids, err := r.ListIdByEndpoint(endpointId)

	err = r.DeleteBatch(ids)

	return
}

func (r *EndpointInterfaceRepo) DeleteByEndpoints(endpointIds []uint) (err error) {
	ids, err := r.ListIdByEndpoints(endpointIds)

	err = r.DeleteBatch(ids)

	return
}

func (r *EndpointInterfaceRepo) DeleteBatch(ids []uint) (err error) {
	for _, id := range ids {
		err = r.Delete(id)
	}
	return
}
func (r *EndpointInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.EndpointInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	endpointInterface, _ := r.Get(id)
	if endpointInterface.DebugInterfaceId > 0 {
		r.DebugInterfaceRepo.Delete(endpointInterface.DebugInterfaceId)
	}
	return
}

func (r *EndpointInterfaceRepo) RemoveAll(id uint) (err error) {
	err = r.RemoveParams(id)
	err = r.RemoveHeaders(id)
	err = r.RemoveCookie(id)

	err = r.removeRequestBody(id)
	err = r.removeResponseBodies(id)

	requestBody, err := r.ListRequestBody(id)
	err = r.removeRequestBodyItem(requestBody.ID)

	responseBodies, err := r.ListResponseBodies(id)
	for _, body := range responseBodies {
		err = r.removeResponseBodyItem(body.ID)
		err = r.removeResponseBodyHeader(body.ID)
	}
	return
}

func (r *EndpointInterfaceRepo) GetByItem(sourceType consts.SourceType, endpointId uint, method consts.HttpMethod) (res model.EndpointInterface, err error) {
	err = r.DB.First(&res, "not deleted AND source_type = ? AND endpoint_id=? and method=?", sourceType, endpointId, method).Error
	return
}

func (r *EndpointInterfaceRepo) GetResponseCodes(endpointInterfaceId uint) (codes []string) {

	responseBodies, err := r.GetResponseBodies([]uint{endpointInterfaceId})
	if err != nil {
		return
	}
	responseBody := responseBodies[endpointInterfaceId]

	for _, item := range responseBody {
		codes = append(codes, item.Code)
	}

	return
}

func (r *EndpointInterfaceRepo) GetResponse(endpointInterfaceId uint, code string) (ret model.EndpointInterfaceResponseBody) {
	responseBodies, err := r.GetResponseBodies([]uint{endpointInterfaceId})
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

func (r *EndpointInterfaceRepo) BatchGetByEndpointIds(endpointIds []uint) (fields []model.EndpointInterface, err error) {
	err = r.DB.Model(model.EndpointInterface{}).
		Where("endpoint_id IN (?) AND NOT deleted", endpointIds).
		Find(&fields).Error
	return
}

func (r *EndpointInterfaceRepo) GetByMethodAndPathAndServeId(serveId uint, path string, method consts.HttpMethod) (endpointInterfaceId uint) {

	type result struct {
		Id uint
	}
	var data result

	err := r.DB.Model(&model.EndpointInterface{}).Select("biz_endpoint_interface.id").Joins("left join biz_endpoint on biz_endpoint_interface.endpoint_id = biz_endpoint.id").Where("not biz_endpoint.deleted and not biz_endpoint_interface.deleted and biz_endpoint.serve_id=? and biz_endpoint.path=? and biz_endpoint_interface.method=?", serveId, path, method).Scan(&data).Error
	if err != nil {
		return 0
	}
	return data.Id

}
