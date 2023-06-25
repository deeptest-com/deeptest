package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DebugInterfaceRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *DebugInterfaceRepo) Tested(id uint) (res bool, err error) {
	var count int64
	err = r.DB.Model(&model.DebugInterface{}).Where("id=?", id).Count(&count).Error
	if err != nil {
		return
	}
	res = count > 0
	return
}

func (r *DebugInterfaceRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.DebugInterface
		r.DB.Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.DebugInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.DebugInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *DebugInterfaceRepo) ListByProject(projectId int) (pos []*model.DebugInterface, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *DebugInterfaceRepo) Get(id uint) (po model.DebugInterface, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *DebugInterfaceRepo) GetDetail(interfId uint) (interf model.DebugInterface, err error) {
	if interfId <= 0 {
		return
	}

	interf, err = r.Get(interfId)

	interf.QueryParams, interf.PathParams, _ = r.ListParams(interfId)
	interf.Headers, _ = r.ListHeaders(interfId)

	interf.BodyFormData, _ = r.ListBodyFormData(interfId)
	interf.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(interfId)

	interf.BasicAuth, _ = r.GetBasicAuth(interfId)
	interf.BearerToken, _ = r.GetBearerToken(interfId)
	interf.OAuth20, _ = r.GetOAuth20(interfId)
	interf.ApiKey, _ = r.GetApiKey(interfId)

	return
}

func (r *DebugInterfaceRepo) Save(interf *model.DebugInterface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Save(interf).Error
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.QueryParams, interf.PathParams)
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
func (r *DebugInterfaceRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *DebugInterfaceRepo) UpdateHeaders(id uint, headers []model.DebugInterfaceHeader) (err error) {
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

func (r *DebugInterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceHeader{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateParams(id uint, queryParams, pathParams []model.DebugInterfaceParam) (err error) {
	err = r.RemoveParams(id)

	var params []model.DebugInterfaceParam

	for _, p := range queryParams {

		if p.Name == "" {
			continue
		}

		p.ID = 0
		p.InterfaceId = id
		p.ParamIn = consts.ParamInQuery
		params = append(params, p)
	}

	for _, p := range pathParams {
		if p.Name == "" {
			continue
		}
		p.ID = 0
		p.InterfaceId = id
		p.ParamIn = consts.ParamInPath
		params = append(params, p)
	}

	if len(params) == 0 {
		return
	}

	err = r.DB.Create(&params).Error

	return
}
func (r *DebugInterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceParam{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateCookies(id uint, cookies []model.DebugInterfaceCookie) (err error) {
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

func (r *DebugInterfaceRepo) RemoveCookie(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceCookie{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBodyFormData(id uint, items []model.DebugInterfaceBodyFormDataItem) (err error) {
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
func (r *DebugInterfaceRepo) RemoveBodyFormData(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBodyFormUrlencoded(id uint, items []model.DebugInterfaceBodyFormUrlEncodedItem) (err error) {
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
func (r *DebugInterfaceRepo) RemoveBodyFormUrlencoded(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBasicAuth(id uint, payload model.DebugInterfaceBasicAuth) (err error) {
	r.RemoveBasicAuth(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBasicAuth(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBasicAuth{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBearerToken(id uint, payload model.DebugInterfaceBearerToken) (err error) {
	r.RemoveBearerToken(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBearerToken(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBearerToken{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.DebugInterfaceOAuth20) (err error) {
	r.RemoveOAuth20(interfaceId)

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.DebugInterfaceOAuth20{}).Error

	return
}

func (r *DebugInterfaceRepo) UpdateApiKey(id uint, payload model.DebugInterfaceApiKey) (err error) {
	r.RemoveApiKey(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveApiKey(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceApiKey{}, "").Error

	return
}

//func (r *DebugInterfaceRepo) makeTree(Data []*modelRef.DebugInterface, node *modelRef.DebugInterface) { //参数为父节点，添加父节点的子节点指针切片
//	children, _ := r.haveChild(Data, node) //判断节点是否有子节点并返回
//	if children != nil {
//		node.Children = append(node.Children, children[0:]...) //添加子节点
//		for _, v := range children {                           //查询子节点的子节点，并添加到子节点
//			_, has := r.haveChild(Data, v)
//			if has {
//				r.makeTree(Data, v) //递归添加节点
//			}
//		}
//	}
//}
//
//func (r *DebugInterfaceRepo) haveChild(Data []*modelRef.DebugInterface, node *modelRef.DebugInterface) (children []*modelRef.DebugInterface, yes bool) {
//	for _, v := range Data {
//		if v.ParentId == node.ID {
//			v.Slots = iris.Map{"icon": "icon"}
//			children = append(children, v)
//		}
//	}
//	if children != nil {
//		yes = true
//	}
//	return
//}

func (r *DebugInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	//field := modelRef.InterfaceDebug{}
	//field.ID = id
	//err = r.DB.Remove(field).SendErrorMsg

	return
}

func (r *DebugInterfaceRepo) GetChildren(defId, fieldId uint) (children []*model.DebugInterface, err error) {
	err = r.DB.Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *DebugInterfaceRepo) SetIsRange(fieldId uint, b bool) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *DebugInterfaceRepo) UpdateOrdAndParent(interf model.DebugInterface) (err error) {
	err = r.DB.Model(&interf).
		Updates(model.DebugInterface{InterfaceBase: model.InterfaceBase{Ordr: interf.Ordr, ParentId: interf.ParentId}}).
		Error

	return
}

func (r *DebugInterfaceRepo) SetOAuth2AccessToken(token string, interfaceId int) (err error) {
	err = r.DB.Model(&model.DebugInterfaceOAuth20{}).
		Where("interface_id = ?", interfaceId).
		Update("access_token", token).Error

	return
}

func (r *DebugInterfaceRepo) ListParams(interfaceId uint) (
	queryParams []model.DebugInterfaceParam, pathParams []model.DebugInterfaceParam, err error) {

	pos := []model.DebugInterfaceParam{}

	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	for _, po := range pos {
		if po.ParamIn == consts.ParamInQuery {
			queryParams = append(queryParams, po)
		} else if po.ParamIn == consts.ParamInPath {
			pathParams = append(pathParams, po)
		}
	}

	return
}
func (r *DebugInterfaceRepo) ListHeaders(interfaceId uint) (pos []model.DebugInterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListCookies(interfaceId uint) (pos []model.DebugInterfaceCookie, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListBodyFormData(interfaceId uint) (pos []model.DebugInterfaceBodyFormDataItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListBodyFormUrlencoded(interfaceId uint) (pos []model.DebugInterfaceBodyFormUrlEncodedItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *DebugInterfaceRepo) GetBasicAuth(id uint) (po model.DebugInterfaceBasicAuth, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetBearerToken(id uint) (po model.DebugInterfaceBearerToken, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetOAuth20(id uint) (po model.DebugInterfaceOAuth20, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetApiKey(id uint) (po model.DebugInterfaceApiKey, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *DebugInterfaceRepo) SaveInterfaces(interf *model.DebugInterface) (err error) {

	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.UpdateInterface(interf)
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.QueryParams, interf.PathParams)
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

		return err
	})

	return
}

func (r *DebugInterfaceRepo) UpdateInterface(interf *model.DebugInterface) (err error) {
	err = r.BaseRepo.Save(interf.ID, interf)
	return
}

//func (r *DebugInterfaceRepo) HasEndpointInterfaceDebugRecord(endpointInterfaceId uint) (id uint, err error) {
//	po, err := r.GetByOwner(endpointInterfaceId, consts.InterfaceDebug)
//	id = po.ID
//
//	return
//}
//func (r *DebugInterfaceRepo) HasScenarioInterfaceDebugRecord(scenarioInterfaceId uint) (id uint, err error) {
//	po, err := r.GetByOwner(scenarioInterfaceId, consts.ScenarioDebug)
//	id = po.ID
//
//	return
//}
//func (r *DebugInterfaceRepo) HasTestInterfaceDebugRecord(testInterfaceId uint) (id uint, err error) {
//	po, err := r.GetByOwner(testInterfaceId, consts.TestDebug)
//	id = po.ID
//
//	return
//}
//
//func (r *DebugInterfaceRepo) GetByOwner(ownerId uint, usedBy consts.UsedBy) (ret modelRef.DebugInterface, err error) {
//	db := r.DB.Where("NOT deleted")
//
//	if usedBy == consts.InterfaceDebug {
//		db.Where("endpoint_interface_id=?", ownerId)
//
//	} else if usedBy == consts.ScenarioDebug {
//		db.Where("scenario_interface_id=?", ownerId)
//
//	} else if usedBy == consts.TestDebug {
//		db.Where("test_interface_id=?", ownerId)
//
//	}
//
//	err = db.Find(&ret).Error
//
//	return
//}

func (r *DebugInterfaceRepo) PopulateProps(po *model.DebugInterface) (err error) {
	po.QueryParams, po.PathParams, _ = r.ListParams(po.ID)
	po.Headers, _ = r.ListHeaders(po.ID)
	po.BodyFormData, _ = r.ListBodyFormData(po.ID)
	po.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(po.ID)
	po.BasicAuth, _ = r.GetBasicAuth(po.ID)
	po.BearerToken, _ = r.GetBearerToken(po.ID)
	po.OAuth20, _ = r.GetOAuth20(po.ID)
	po.ApiKey, _ = r.GetApiKey(po.ID)

	return
}

func (r *DebugInterfaceRepo) UpdateDebugInfo(id uint, values map[string]interface{}) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id=?", id).
		Updates(values).
		Error

	return
}
