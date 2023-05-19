package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type EndpointInterfaceRepo struct {
	*BaseRepo    `inject:""`
	DB           *gorm.DB      `inject:""`
	EndpointRepo *EndpointRepo `inject:""`
}

func (r *EndpointInterfaceRepo) Paginate(req v1.EndpointInterfaceReqPaginate) (ret _domain.PageData, err error) {
	/*
		endpointIds, err := r.EndpointRepo.ListEndpointByCategory(req.CategoryId)
		if err != nil {
			return
		}

	*/

	var count int64
	db := r.DB.Model(&model.EndpointInterface{}).
		Where("project_id = ? AND NOT deleted AND NOT disabled", req.ProjectId)

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	db = db.Order("created_at desc")
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

func (r *EndpointInterfaceRepo) Get(interfaceId uint) (field model.EndpointInterface, err error) {
	err = r.DB.
		Where("id=?", interfaceId).
		Where("NOT deleted").
		First(&field).Error
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

func (r *EndpointInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.EndpointInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	//field := model.UsedByInterface{}
	//field.ID = id
	//err = r.DB.Delete(field).SendErrorMsg

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
	for _, responseBody := range responseBodies {
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

func (r *EndpointInterfaceRepo) GetByEndpointId(endpointId uint, version string) (interfaces []model.EndpointInterface, err error) {

	interfaces, err = r.GetEndpointId(endpointId, version)
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

func (r *EndpointInterfaceRepo) GetEndpointId(endpointId uint, version string) (field []model.EndpointInterface, err error) {
	err = r.DB.
		Where("endpoint_id=? and version=?", endpointId, version).
		Where("NOT deleted").
		Find(&field).Error
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
	//fmt.Println(requestBodyId, "+++++++++++++")
	err = r.DB.First(&requestBodyItem, "request_body_id = ?", requestBodyId).Error
	//fmt.Println(err)
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
