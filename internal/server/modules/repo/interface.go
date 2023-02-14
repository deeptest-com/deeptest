package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type InterfaceRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *InterfaceRepo) GetInterfaceTree(projectId int) (root *model.Interface, err error) {
	pos, err := r.ListByProject(projectId)

	if err != nil {
		return
	}

	root = pos[0]
	root.Slots = iris.Map{"icon": "icon"}
	r.makeTree(pos[1:], root)
	return
}

func (r *InterfaceRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.Interface
		r.DB.Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.Interface{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.Interface{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *InterfaceRepo) ListByProject(projectId int) (pos []*model.Interface, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *InterfaceRepo) Get(interfaceId uint) (field model.Interface, err error) {
	err = r.DB.
		Where("id=?", interfaceId).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *InterfaceRepo) GetDetail(interfId uint) (interf model.Interface, err error) {
	if interfId > 0 {
		interf, err = r.Get(interfId)

		interf.Params, _ = r.ListParams(interfId)
		interf.Headers, _ = r.ListHeaders(interfId)
		interf.BodyFormData, _ = r.ListBodyFormData(interfId)
		interf.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(interfId)

		interf.BasicAuth, _ = r.GetBasicAuth(interfId)
		interf.BearerToken, _ = r.GetBearerToken(interfId)
		interf.OAuth20, _ = r.GetOAuth20(interfId)
		interf.ApiKey, _ = r.GetApiKey(interfId)
	}

	return
}

func (r *InterfaceRepo) Save(interf *model.Interface) (err error) {
	err = r.DB.Save(interf).Error
	return
}

func (r *InterfaceRepo) Update(interf model.Interface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Updates(interf).Error
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.Params)
		if err != nil {
			return err
		}

		err = r.UpdateBodyFormData(interf.ID, interf.BodyFormData)
		if err != nil {
			return err
		}

		err = r.UpdateBodyFormUrlencoded(interf.ID, interf.BodyFormUrlencoded)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(interf.ID, interf.Headers)
		if err != nil {
			return err
		}

		err = r.UpdateBasicAuth(interf.ID, interf.BasicAuth)
		if err != nil {
			return err
		}

		err = r.UpdateBearerToken(interf.ID, interf.BearerToken)
		if err != nil {
			return err
		}

		err = r.UpdateOAuth20(interf.ID, interf.OAuth20)
		if err != nil {
			return err
		}

		err = r.UpdateApiKey(interf.ID, interf.ApiKey)
		if err != nil {
			return err
		}

		return err
	})

	return
}
func (r *InterfaceRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.Interface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *InterfaceRepo) UpdateHeaders(id uint, headers []model.InterfaceHeader) (err error) {
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

func (r *InterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceHeader{}, "").Error

	return
}

func (r *InterfaceRepo) UpdateParams(id uint, params []model.InterfaceParam) (err error) {
	err = r.RemoveParams(id)

	if len(params) == 0 {
		return
	}

	for idx, _ := range params {
		params[idx].ID = 0
		params[idx].InterfaceId = id
	}

	err = r.DB.Create(&params).Error

	return
}
func (r *InterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceParam{}, "").Error

	return
}

func (r *InterfaceRepo) UpdateCookies(id uint, cookies []model.InterfaceCookie) (err error) {
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

func (r *InterfaceRepo) RemoveCookie(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceCookie{}, "").Error

	return
}

func (r *InterfaceRepo) UpdateBodyFormData(id uint, items []model.InterfaceBodyFormDataItem) (err error) {
	err = r.RemoveBodyFormData(id)

	if len(items) == 0 {
		return
	}

	for idx, _ := range items {
		items[idx].ID = 0
		items[idx].InterfaceId = id
	}

	err = r.DB.Create(&items).Error

	return
}
func (r *InterfaceRepo) RemoveBodyFormData(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *InterfaceRepo) UpdateBodyFormUrlencoded(id uint, items []model.InterfaceBodyFormUrlEncodedItem) (err error) {
	err = r.RemoveBodyFormUrlencoded(id)

	if len(items) == 0 {
		return
	}

	for idx, _ := range items {
		items[idx].ID = 0
		items[idx].InterfaceId = id
	}

	err = r.DB.Create(&items).Error

	return
}
func (r *InterfaceRepo) RemoveBodyFormUrlencoded(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *InterfaceRepo) UpdateBasicAuth(id uint, payload model.InterfaceBasicAuth) (err error) {
	r.RemoveBasicAuth(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *InterfaceRepo) RemoveBasicAuth(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceBasicAuth{}, "").Error

	return
}

func (r *InterfaceRepo) UpdateBearerToken(id uint, payload model.InterfaceBearerToken) (err error) {
	r.RemoveBearerToken(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *InterfaceRepo) RemoveBearerToken(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceBearerToken{}, "").Error

	return
}

func (r *InterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.InterfaceOAuth20) (err error) {
	r.RemoveOAuth20(interfaceId)

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}
func (r *InterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.InterfaceOAuth20{}).Error

	return
}

func (r *InterfaceRepo) UpdateApiKey(id uint, payload model.InterfaceApiKey) (err error) {
	r.RemoveApiKey(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *InterfaceRepo) RemoveApiKey(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceApiKey{}, "").Error

	return
}

func (r *InterfaceRepo) makeTree(Data []*model.Interface, node *model.Interface) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.haveChild(Data, node) //判断节点是否有子节点并返回
	if children != nil {
		node.Children = append(node.Children, children[0:]...) //添加子节点
		for _, v := range children {                           //查询子节点的子节点，并添加到子节点
			_, has := r.haveChild(Data, v)
			if has {
				r.makeTree(Data, v) //递归添加节点
			}
		}
	}
}

func (r *InterfaceRepo) haveChild(Data []*model.Interface, node *model.Interface) (children []*model.Interface, yes bool) {
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

func (r *InterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Interface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	//field := model.UsedByInterface{}
	//field.ID = id
	//err = r.DB.Delete(field).SendErrorMsg

	return
}

func (r *InterfaceRepo) GetChildren(defId, fieldId uint) (children []*model.Interface, err error) {
	err = r.DB.Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *InterfaceRepo) SetIsRange(fieldId uint, b bool) (err error) {
	err = r.DB.Model(&model.Interface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *InterfaceRepo) UpdateOrdAndParent(interf model.Interface) (err error) {
	err = r.DB.Model(&interf).
		Updates(model.Interface{InterfaceBase: model.InterfaceBase{Ordr: interf.Ordr, ParentId: interf.ParentId}}).
		Error

	return
}

func (r *InterfaceRepo) SetOAuth2AccessToken(token string, interfaceId int) (err error) {
	err = r.DB.Model(&model.InterfaceOAuth20{}).
		Where("interface_id = ?", interfaceId).
		Update("access_token", token).Error

	return
}

func (r *InterfaceRepo) ListParams(interfaceId uint) (pos []model.InterfaceParam, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error
	return
}
func (r *InterfaceRepo) ListHeaders(interfaceId uint) (pos []model.InterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *InterfaceRepo) ListCookies(interfaceId uint) (pos []model.InterfaceCookie, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *InterfaceRepo) ListBodyFormData(interfaceId uint) (pos []model.InterfaceBodyFormDataItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *InterfaceRepo) ListBodyFormUrlencoded(interfaceId uint) (pos []model.InterfaceBodyFormUrlEncodedItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *InterfaceRepo) GetBasicAuth(id uint) (po model.InterfaceBasicAuth, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *InterfaceRepo) GetBearerToken(id uint) (po model.InterfaceBearerToken, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *InterfaceRepo) GetOAuth20(id uint) (po model.InterfaceOAuth20, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *InterfaceRepo) GetApiKey(id uint) (po model.InterfaceApiKey, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *InterfaceRepo) SaveInterfaces(interf model.Interface) (err error) {

	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.UpdateInterface(&interf)
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

func (r *InterfaceRepo) UpdateRequestBody(requestBody *model.InterfaceRequestBody) (err error) {
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

func (r *InterfaceRepo) removeRequestBodyItem(requestBodyId uint) (err error) {
	err = r.DB.
		Where("request_body_id = ?", requestBodyId).
		Delete(&model.InterfaceRequestBodyItem{}).Error

	return
}

func (r *InterfaceRepo) removeRequestBody(interId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interId).
		Delete(&model.InterfaceRequestBody{}).Error

	return
}

func (r *InterfaceRepo) UpdateResponseBodies(interfaceId uint, responseBodies []model.InterfaceResponseBody) (err error) {
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
		err = r.removeRequestBodyItem(responseBody.ID)
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

func (r *InterfaceRepo) removeResponseBodies(interId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interId).
		Delete(&model.InterfaceResponseBody{}).Error

	return
}

func (r *InterfaceRepo) removeResponseBodyItem(responseBodyId uint) (err error) {
	err = r.DB.
		Where("response_body_id = ?", responseBodyId).
		Delete(&model.InterfaceResponseBodyItem{}).Error

	return
}

func (r *InterfaceRepo) removeResponseBodyHeader(responseBodyId uint) (err error) {
	err = r.DB.
		Where("response_body_id = ?", responseBodyId).
		Delete(&model.InterfaceResponseBodyHeader{}).Error

	return
}

func (r *InterfaceRepo) UpdateInterface(interf *model.Interface) (err error) {
	err = r.BaseRepo.Save(interf.ID, interf)
	return
}

func (r *InterfaceRepo) GetByEndpointId(endpointId uint) (interfaces []model.Interface, err error) {

	interfaces, err = r.GetEndpointId(endpointId)
	for key, interf := range interfaces {
		interfaces[key].Params, _ = r.ListParams(interf.ID)
		interfaces[key].Headers, _ = r.ListHeaders(interf.ID)
		interfaces[key].Cookies, _ = r.ListCookies(interf.ID)
		interfaces[key].RequestBody, _ = r.ListRequestBody(interf.ID)
		interfaces[key].ResponseBodies, _ = r.ListResponseBodies(interf.ID)
	}

	return
}

func (r *InterfaceRepo) GetEndpointId(endpointId uint) (field []model.Interface, err error) {
	err = r.DB.
		Where("endpoint_id=?", endpointId).
		Where("NOT deleted").
		Find(&field).Error
	return
}

func (r *InterfaceRepo) ListRequestBody(interfaceId uint) (requestBody model.InterfaceRequestBody, err error) {
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
func (r *InterfaceRepo) ListRequestBodyItem(requestBodyId uint) (requestBodyItem model.InterfaceRequestBodyItem, err error) {
	//fmt.Println(requestBodyId, "+++++++++++++")
	err = r.DB.First(&requestBodyItem, "request_body_id = ?", requestBodyId).Error
	//fmt.Println(err)
	return
}

func (r *InterfaceRepo) ListResponseBodies(interfaceId uint) (responseBodies []model.InterfaceResponseBody, err error) {
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

func (r *InterfaceRepo) ListResponseBodyItem(requestBodyId uint) (responseBodyItem model.InterfaceResponseBodyItem, err error) {
	err = r.DB.First(&responseBodyItem, "response_body_id = ?", requestBodyId).Error
	//fmt.Println(err)
	return
}

func (r *InterfaceRepo) ListResponseBodyHeaders(requestBodyId uint) (responseBodyHeaders []model.InterfaceResponseBodyHeader, err error) {
	err = r.DB.Find(&responseBodyHeaders, "response_body_id = ?", requestBodyId).Error
	return
}
